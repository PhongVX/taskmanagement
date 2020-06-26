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

//Identity: userName, email, _id
func (r *MongoDBRepository) FindByIdentity(ctx context.Context, identity string) (*User, error) {
	s := r.session.Clone()
	defer s.Close()
	var u User
	var findingField = []bson.M{
		bson.M{"user_name": identity},
		bson.M{"email": identity},
	}
	if bson.IsObjectIdHex(identity) {
		findingField = append(findingField, bson.M{"_id": bson.ObjectIdHex(identity)})
	}
	if err := s.DB("").C(USER_COLLECTION_NAME).Find(bson.M{
		"$or": findingField,
	}).One(&u); err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *MongoDBRepository) FindByUserIdentity(ctx context.Context, u *User) (*User, error) {
	s := r.session.Clone()
	defer s.Close()
	var user User
	var findingField = []bson.M{
		bson.M{"user_name": u.UserName},
		bson.M{"email": u.Email},
	}
	if err := s.DB("").C(USER_COLLECTION_NAME).Find(bson.M{
		"$or": findingField,
	}).One(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// Delete a user by user id
func (r *MongoDBRepository) Delete(ctx context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	return s.DB("").C(USER_COLLECTION_NAME).RemoveId(bson.ObjectIdHex(id))
}

// Update a user
func (r *MongoDBRepository) Update(ctx context.Context, u *User) error {
	s := r.session.Clone()
	defer s.Close()
	err := s.DB("").C(USER_COLLECTION_NAME).Update(bson.M{"_id": u.ID}, bson.M{"$set": &u})
	return err
}
