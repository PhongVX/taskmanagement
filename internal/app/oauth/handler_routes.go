package oauth

import (
	"net/http"

	"github.com/PhongVX/taskmanagement/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/oauth/google/login",
			Method:  http.MethodGet,
			Handler: h.GoogleLogin,
		},
		{
			Path:    "/api/v1/oauth/google/callback",
			Method:  http.MethodGet,
			Handler: h.GoogleCallback,
		},
	}
}
