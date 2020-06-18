package user

import (
	"net/http"

	"github.com/PhongVX/taskmanagement/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/users",
			Method:  http.MethodGet,
			Handler: h.FindAll,
		},
		{
			Path:    "/api/v1/users",
			Method:  http.MethodPost,
			Handler: h.Insert,
		},
		{
			Path:    "/api/v1/users",
			Method:  http.MethodPut,
			Handler: h.Update,
		},
		{
			Path:    "/api/v1/users/{id}",
			Method:  http.MethodGet,
			Handler: h.FindByID,
		},
		{
			Path:    "/api/v1/users/{id}",
			Method:  http.MethodDelete,
			Handler: h.Delete,
		},
	}
}
