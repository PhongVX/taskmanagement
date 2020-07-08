package api

import (
	"github.com/PhongVX/taskmanagement/internal/app/oauth"
)

func newOAuthHandler() (*oauth.Handler, error) {
	srv := oauth.NewService()
	handler := oauth.NewHTTPHandler(srv)
	return handler, nil
}
