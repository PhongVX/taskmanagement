package auth

import (
	"context"
)

//Interfaces
type (
	ServiceInterface interface {
		Login(ctx context.Context, au *Auth) (Tokens, error)
		Refresh(ctx context.Context, t *Tokens) (Tokens, error)
		//Logout(ctx context.Context, t *Auth) error
	}
)

//Data Struct
type (
	//TODO Need to move this json format to common package
	User struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		UserName  string `json:"user_name"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
	}

	Auth struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}

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
