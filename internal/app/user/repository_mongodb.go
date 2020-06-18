package user

import (
	"context"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func NewMongoDBRepository(session *mgo.Session) *MongoDBRepository {
	return &MongoDBRepository{
		session: session,
	}
}

// Find All user
func (r *MongoDBRepository) FindAll(ctx context.Context, rO FindingRequestObject) ([]*User, error) {
	findingField := bson.M{}
	users := make([]*User, 0)
	s := r.session.Clone()
	defer s.Close()
	if err := s.DB("").C(USER_COLLECTION_NAME).Find(findingField).Sort(rO.SortBy...).Skip(rO.Offset).Limit(rO.Limit).All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

// Create create new user
func (r *MongoDBRepository) Insert(ctx context.Context, t *User) error {
	s := r.session.Clone()
	defer s.Close()
	t.ID = bson.NewObjectId()
	if err := s.DB("").C(USER_COLLECTION_NAME).Insert(t); err != nil {
		return err
	}
	return nil
}

// Find a user by user id
func (r *MongoDBRepository) FindByID(ctx context.Context, id string) (*User, error) {
	s := r.session.Clone()
	defer s.Close()
	var t User
	if err := s.DB("").C(USER_COLLECTION_NAME).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&t); err != nil {
		return nil, err
	}
	return &t, nil
}

// Delete a task by task id
func (r *MongoDBRepository) Delete(cxt context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	return s.DB("").C(USER_COLLECTION_NAME).RemoveId(bson.ObjectIdHex(id))
}

// Update a task
func (r *MongoDBRepository) Update(cxt context.Context, t *User) error {
	s := r.session.Clone()
	defer s.Close()
	err := s.DB("").C(USER_COLLECTION_NAME).Update(bson.M{"_id": t.ID}, bson.M{"$set": &t})
	return err
}
