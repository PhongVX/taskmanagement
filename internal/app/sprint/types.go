package sprint

import (
	"context"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//Interfaces
type (
	Repository interface {
		Insert(ctx context.Context, t *Sprint) error
		FindByID(ctx context.Context, id string) (*Sprint, error)
		FindAll(ctx context.Context, r FindingRequestObject) ([]*Sprint, error)
		Delete(cxt context.Context, id string) error
		Update(cxt context.Context, t *Sprint) error
	}
)

//Data Struct
type (
	Sprint struct {
		ID              bson.ObjectId `json:"id" bson:"_id"`
		Title           string        `json:"title" bson:"title,omitempty"`
		Description     string        `json:"description" bson:"description,omitempty"`
		Status          string        `json:"status" bson:"status" default:"TODO"`
		CreatedAt       *time.Time    `json:"created_at" bson:"created_at,omitempty"`
		UpdatedAt       *time.Time    `json:"updated_at" bson:"updated_at"`
		CreatedByID     string        `json:"created_by_id,omitempty" bson:"created_by_id"`
		CreatedByName   string        `json:"created_by_name,omitempty" bson:"created_by_name"`
		CreatedByAvatar string        `json:"created_by_avatar,omitempty" bson:"created_by_avatar"`
	}

	Config struct {
		MaxPageSize int `envconfig:"SPRINT_MAX_PAGE_SIZE" default:"50"`
	}

	FindingRequestObject struct {
		Offset      int      `json:"offset"`
		Limit       int      `json:"limit"`
		CreatedByID string   `json:"created_by_id"`
		SortBy      []string `json:"sort_by"`
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
