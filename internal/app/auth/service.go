package auth

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/PhongVX/taskmanagement/internal/pkg/hash/password"
	"github.com/PhongVX/taskmanagement/internal/pkg/http/request"
	"github.com/PhongVX/taskmanagement/internal/pkg/jwt"
	"github.com/PhongVX/taskmanagement/internal/pkg/log"
	"github.com/PhongVX/taskmanagement/internal/pkg/mapstructure"
	"github.com/PhongVX/taskmanagement/internal/pkg/types/responsetype"
)

// NewService return a new auth service
func NewService() *Service {
	srv := &Service{}
	return srv
}

func (s *Service) Login(ctx context.Context, au *Auth) (Tokens, error) {
	tokens := Tokens{}
	url := USER_SERVICE_URL + "/" + au.UserName
	//TODO Will using async request calling or message queue in future
	res := &responsetype.Base{}
	err := request.Get(url, res)
	if err != nil {
		return tokens, err
	}
	result := res.Result
	log.WithContext(ctx).Errorf("result, user:  %v", result)
	if result == nil {
		return tokens, errors.New("User doesn't exist")
	}
	var user User
	err = mapstructure.Decode(result, &user)
	if err != nil {
		return tokens, err
	}
	log.WithContext(ctx).Errorf("User, user:  %v", user)
	err = password.ComparePassword(user.Password, au.Password)
	if err != nil {
		log.WithContext(ctx).Errorf("Failed to comapre password, err:  %v", err)
		return tokens, errors.New("Password is not correct")
	}
	td, err := jwt.CreateToken(user.ID)
	tokens.AccessToken = td.AccessToken
	tokens.RefreshToken = td.RefreshToken
	return tokens, nil
}

func (s *Service) Refresh(ctx context.Context, t *Tokens) (Tokens, error) {
	//TODO Need to read from env file in future
	tokens := Tokens{}
	os.Setenv("REFRESH_SECRET", "my-refresh-secret")
	claims, err := jwt.VerifyToken(t.RefreshToken, os.Getenv("REFRESH_SECRET"))
	if err != nil {
		return tokens, err
	}
	td, err := jwt.CreateToken(fmt.Sprintf("%v", claims["user_id"]))
	tokens.AccessToken = td.AccessToken
	tokens.RefreshToken = td.RefreshToken
	return tokens, nil
}
