package task

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

func (r *MongoDBRepository) FindAll(ctx context.Context, rO FindingRequestObject) ([]*Task, error) {
	findingField := bson.M{}
	tasks := make([]*Task, 0)
	s := r.session.Clone()
	defer s.Close()
	if err := s.DB("").C(TASK_COLLECTION_NAME).Find(findingField).Sort(rO.SortBy...).Skip(rO.Offset).Limit(rO.Limit).All(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

// Create create new task
func (r *MongoDBRepository) Insert(ctx context.Context, t *Task) error {
	s := r.session.Clone()
	defer s.Close()
	t.ID = bson.NewObjectId()
	t.CreatedAt = timeutil.Now()
	t.UpdatedAt = t.CreatedAt
	if err := s.DB("").C(TASK_COLLECTION_NAME).Insert(t); err != nil {
		return err
	}
	return nil
}

// Find a task by task id
func (r *MongoDBRepository) FindByID(ctx context.Context, id string) (*Task, error) {
	s := r.session.Clone()
	defer s.Close()
	var t Task
	if err := s.DB("").C(TASK_COLLECTION_NAME).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *MongoDBRepository) Delete(cxt context.Context, id string) error {

	return nil
}

func (r *MongoDBRepository) Update(cxt context.Context, t *Task) error {

	return nil
}
