package api

import "github.com/PhongVX/taskmanagement/internal/app/auth"

func newAuthHandler() (*auth.Handler, error) {
	srv := auth.NewService()
	handler := auth.NewHTTPHandler(srv)
	return handler, nil
}
