package api

import "github.com/PhongVX/taskmanagement/internal/app/sprint"

func newSprintHandler() (*sprint.Handler, error) {
	s, err := dialDefaultMongoDB()
	if err != nil {
		return nil, err
	}
	repo := sprint.NewMongoDBRepository(s)
	srv := sprint.NewService(repo)
	handler := sprint.NewHTTPHandler(*srv)
	return handler, nil
}
