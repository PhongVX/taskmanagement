package sprint

import (
	"net/http"

	"github.com/PhongVX/taskmanagement/internal/pkg/http/middleware"
	"github.com/PhongVX/taskmanagement/internal/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:        "/api/v1/sprints",
			Method:      http.MethodGet,
			Handler:     h.FindAll,
			Middlewares: []router.Middleware{middleware.EnsureAuthenticated},
		},
		{
			Path:        "/api/v1/sprints",
			Method:      http.MethodPost,
			Handler:     h.Insert,
			Middlewares: []router.Middleware{middleware.EnsureAuthenticated},
		},
		{
			Path:        "/api/v1/sprints",
			Method:      http.MethodPut,
			Handler:     h.Update,
			Middlewares: []router.Middleware{middleware.EnsureAuthenticated},
		},
		{
			Path:        "/api/v1/sprints/{id}",
			Method:      http.MethodGet,
			Handler:     h.FindByID,
			Middlewares: []router.Middleware{middleware.EnsureAuthenticated},
		},
		{
			Path:        "/api/v1/sprints/{id}",
			Method:      http.MethodDelete,
			Handler:     h.Delete,
			Middlewares: []router.Middleware{middleware.EnsureAuthenticated},
		},
	}
}
