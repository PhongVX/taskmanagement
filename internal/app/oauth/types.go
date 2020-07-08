package oauth

import (
	"context"
	"net/http"

	"github.com/globalsign/mgo/bson"
)

//Interfaces
type (
	ServiceInterface interface {
		GoogleLogin(w http.ResponseWriter, r *http.Request)
		GoogleCallback(ctx context.Context, u *UserFromRequest) (Tokens, error)
	}
)

//Data Struct
type (
	//TODO Need to move this json format to common package
	User struct {
		ID                bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
		Email             string        `json:"email"`
		VerifiedEmail     bool          `json:"verified_email"`
		Picture           string        `json:"picture"`
		FirstName         string        `json:"first_name"`
		LastName          string        `json:"last_name"`
		IsOauthFirstLogin bool          `json:"is_oauth_first_login,omitempty"`
	}

	UserFromRequest struct {
		Email             string `json:"email"`
		VerifiedEmail     bool   `json:"verified_email"`
		Picture           string `json:"picture"`
		FirstName         string `json:"given_name"`
		LastName          string `json:"family_name"`
		IsOauthFirstLogin bool   `json:"is_oauth_first_login"`
	}

	// Auth struct {
	// 	UserName string `json:"user_name"`
	// 	Password string `json:"password"`
	// }

	Tokens struct {
		AccessToken  string `json:"access_token,omitempty"`
		RefreshToken string `json:"refresh_token,omitempty"`
	}

	Handler struct {
		srv ServiceInterface
		//Func Routes ==> handler_routes.go
	}

	Service struct{}
)
