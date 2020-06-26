package auth

import (
	"context"
	"errors"

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

func (s *Service) Login(ctx context.Context, au *Auth) (map[string]string, error) {
	url := USER_SERVICE_URL + "/" + au.UserName
	//TODO Will using async request calling or message queue in future
	res := &responsetype.Base{}
	err := request.Get(url, res)
	if err != nil {
		return nil, err
	}
	result := res.Result
	if result == "" {
		return nil, errors.New("User doesn't exist")
	}
	var user User
	err = mapstructure.Decode(result, &user)
	if err != nil {
		return nil, err
	}
	err = password.ComparePassword(user.Password, au.Password)
	if err != nil {
		log.WithContext(ctx).Errorf("Failed to comapre password, err:  %v", err)
		return nil, errors.New("Password is not correct")
	}

	td, err := jwt.CreateToken(user.ID.Hex())
	tokens := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
	}
	return tokens, nil
}
