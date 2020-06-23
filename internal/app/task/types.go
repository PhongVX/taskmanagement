package task

import (
	"context"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//Interfaces
type (
	Repository interface {
		Insert(ctx context.Context, t *Task) error
		FindByID(ctx context.Context, id string) (*Task, error)
		FindAll(ctx context.Context, r FindingRequestObject) ([]*Task, error)
		Delete(cxt context.Context, id string) error
		Update(cxt context.Context, t *Task) error
	}
)

//Data Structs
type (
	Task struct {
		ID              bson.ObjectId `json:"id" bson:"_id"`
		Title           string        `json:"title" bson:"title,omitempty"`
		Description     string        `json:"description" bson:"description,omitempty"`
		Status          string        `json:"status" bson:"status" default:"TODO"`
		CreatedAt       *time.Time    `json:"created_at" bson:"created_at,omitempty"`
		UpdatedAt       *time.Time    `json:"updated_at" bson:"updated_at"`
		SprintID        string        `json:"sprint_id" bson:"sprint_id,omitempty"`
		CreatedByID     string        `json:"created_by_id,omitempty" bson:"created_by_id"`
		CreatedByName   string        `json:"created_by_name,omitempty" bson:"created_by_name"`
		CreatedByAvatar string        `json:"created_by_avatar,omitempty" bson:"created_by_avatar"`
	}

	FindingRequestObject struct {
		Offset      int      `json:"offset"`
		Limit       int      `json:"limit"`
		SprintID    []string `json:"sprint_id"`
		CreatedByID string   `json:"created_by_id"`
		SortBy      []string `json:"sort_by"`
	}

	Config struct {
		MaxPageSize int `envconfig:"TASK_MAX_PAGE_SIZE" default:"50"`
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
