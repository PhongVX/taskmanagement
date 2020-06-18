package user

import (
	"context"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//Interfaces
type (
	Repository interface {
		Insert(ctx context.Context, t *User) error
		FindByID(ctx context.Context, id string) (*User, error)
		FindAll(ctx context.Context, r FindingRequestObject) ([]*User, error)
		Delete(cxt context.Context, id string) error
		Update(cxt context.Context, t *User) error
	}
)

//Data Struct
type (
	User struct {
		ID        bson.ObjectId `json:"id" bson:"_id"`
		Email     string        `json:"email" bson:"email,omitempty"`
		UserName  string        `json:"user_name" bson:"user_name,omitempty"`
		FirstName string        `json:"first_name" bson:"first_name,omitempty"`
		LastName  string        `json:"last_name" bson:"last_name,omitempty"`
	}

	Config struct {
		MaxPageSize int `envconfig:"USER_MAX_PAGE_SIZE" default:"50"`
	}

	FindingRequestObject struct {
		Offset int      `json:"offset"`
		Limit  int      `json:"limit"`
		SortBy []string `json:"sort_by"`
	}

	Handler struct {
		srv Service
		//Func Routes ==> handler_routes.go
		//Func FindAll ==> handler.go
	}

	Service struct {
		repo Repository
		conf Config
	}

	MongoDBRepository struct {
		session *mgo.Session
	}
)
