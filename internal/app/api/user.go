package api

import (
	"github.com/PhongVX/taskmanagement/internal/app/user"
)

func newUserHandler() (*user.Handler, error) {
	s, err := dialDefaultMongoDB()
	if err != nil {
		return nil, err
	}
	repo := user.NewMongoDBRepository(s)
	srv := user.NewService(repo)
	handler := user.NewHTTPHandler(srv)
	return handler, nil
}
