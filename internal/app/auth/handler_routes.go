package auth

import (
	"net/http"

	"github.com/PhongVX/taskmanagement/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/auth/login",
			Method:  http.MethodPost,
			Handler: h.Login,
		},
	}
}
