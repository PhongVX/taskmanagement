package sprint

import (
	"context"

	"github.com/PhongVX/taskmanagement/internal/pkg/utils/timeutil"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func NewMongoDBRepository(session *mgo.Session) *MongoDBRepository {
	return &MongoDBRepository{
		session: session,
	}
}

// Find All task
func (r *MongoDBRepository) FindAll(ctx context.Context, rO FindingRequestObject) ([]*Sprint, error) {
	findingField := bson.M{}
	sprints := make([]*Sprint, 0)
	s := r.session.Clone()
	defer s.Close()
	if rO.CreatedByID != "" {
		findingField["created_by_id"] = rO.CreatedByID
	}
	if err := s.DB("").C(SPRINT_COLLECTION_NAME).Find(findingField).Sort(rO.SortBy...).Skip(rO.Offset).Limit(rO.Limit).All(&sprints); err != nil {
		return nil, err
	}
	return sprints, nil
}

// Create create new sprint
func (r *MongoDBRepository) Insert(ctx context.Context, t *Sprint) error {
	s := r.session.Clone()
	defer s.Close()
	t.ID = bson.NewObjectId()
	t.CreatedAt = timeutil.Now()
	t.UpdatedAt = t.CreatedAt
	if err := s.DB("").C(SPRINT_COLLECTION_NAME).Insert(t); err != nil {
		return err
	}
	return nil
}

// Find a sprint by sprint id
func (r *MongoDBRepository) FindByID(ctx context.Context, id string) (*Sprint, error) {
	s := r.session.Clone()
	defer s.Close()
	var t Sprint
	if err := s.DB("").C(SPRINT_COLLECTION_NAME).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&t); err != nil {
		return nil, err
	}
	return &t, nil
}

// Delete a sprint by sprint id
func (r *MongoDBRepository) Delete(cxt context.Context, id string) error {
	s := r.session.Clone()
	defer s.Close()
	return s.DB("").C(SPRINT_COLLECTION_NAME).RemoveId(bson.ObjectIdHex(id))
}

// Update a sprint
func (r *MongoDBRepository) Update(cxt context.Context, t *Sprint) error {
	s := r.session.Clone()
	defer s.Close()
	t.UpdatedAt = timeutil.Now()
	err := s.DB("").C(SPRINT_COLLECTION_NAME).Update(bson.M{"_id": t.ID}, bson.M{"$set": &t})
	return err
}
