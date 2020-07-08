package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/PhongVX/taskmanagement/internal/pkg/http/request"
	"github.com/PhongVX/taskmanagement/internal/pkg/jwt"
	"github.com/PhongVX/taskmanagement/internal/pkg/log"
	"github.com/PhongVX/taskmanagement/internal/pkg/mapstructure"
	"github.com/PhongVX/taskmanagement/internal/pkg/types/responsetype"
)

// NewService return a new auth2 service
func NewService() *Service {
	srv := &Service{}
	return srv
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)
	return state
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}

func (s *Service) GoogleLogin(w http.ResponseWriter, r *http.Request) {
	oauthState := generateStateOauthCookie(w)
	u := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func (s *Service) GoogleCallback(ctx context.Context, u *UserFromRequest) (Tokens, error) {
	var user User
	tokens := Tokens{}
	url := USER_SERVICE_URL + "/" + u.Email
	res := &responsetype.Base{}
	err := request.Get(url, res)
	if err != nil {
		return tokens, nil
	}
	result := res.Result
	if result == nil {
		//User doesn't existed
		userReuest := &User{}
		userReuest.Email = u.Email
		userReuest.VerifiedEmail = u.VerifiedEmail
		userReuest.Email = u.Email
		userReuest.Picture = u.Picture
		userReuest.FirstName = u.FirstName
		userReuest.LastName = u.LastName
		userReuest.IsOauthFirstLogin = true
		mapU, err := json.Marshal(userReuest)
		if err != nil {
			return tokens, err
		}
		//Create New
		res, err := request.Post(USER_SERVICE_URL, mapU)
		if err != nil {
			return tokens, err
		}
		log.WithContext(ctx).Infof("Created User:  %v", res)
		result = res.Result
		if result == "" {
			return tokens, nil
		}
	}
	if err := mapstructure.Decode(result, &user); err != nil {
		return tokens, err
	}
	td, err := jwt.CreateToken(string(user.ID))
	if err != nil {
		return tokens, err
	}
	tokens.AccessToken = td.AccessToken
	tokens.RefreshToken = td.RefreshToken

	return tokens, nil
}
