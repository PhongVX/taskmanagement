package auth

import (
	"context"

	"github.com/globalsign/mgo/bson"
)

//Interfaces
type (
	ServiceInterface interface {
		Login(ctx context.Context, t *Auth) (map[string]string, error)
		//Logout(ctx context.Context, t *Auth) error
	}
)

//Data Struct
type (
	//TODO Need to move this json format to common package
	User struct {
		ID        bson.ObjectId `json:"id" bson:"_id"`
		Email     string        `json:"email" bson:"email,omitempty"`
		UserName  string        `json:"user_name" bson:"user_name,omitempty"`
		FirstName string        `json:"first_name" bson:"first_name,omitempty"`
		LastName  string        `json:"last_name" bson:"last_name,omitempty"`
		Password  string        `json:"password" bson:"password,omitempty"`
	}

	Auth struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}

	Handler struct {
		srv ServiceInterface
		//Func Routes ==> handler_routes.go
	}

	Service struct{}
)
