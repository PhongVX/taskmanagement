package dbutil

import (
	"github.com/globalsign/mgo/bson"
)

// NewID return new id for database
func NewObjectId() bson.ObjectId {
	return bson.NewObjectId()
}
