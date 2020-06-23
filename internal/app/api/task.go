package api

import (
	"github.com/PhongVX/taskmanagement/internal/app/task"
)

func newTaskHandler() (*task.Handler, error) {
	s, err := dialDefaultMongoDB()
	if err != nil {
		return nil, err
	}
	repo := task.NewMongoDBRepository(s)
	srv := task.NewService(repo)
	handler := task.NewHTTPHandler(*srv)
	return handler, nil
}
