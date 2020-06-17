package api

import (
	"net/http"

	"github.com/PhongVX/taskmanagement/internal/pkg/http/router"
	"github.com/PhongVX/taskmanagement/internal/pkg/log"
	"github.com/gorilla/mux"
)

func NewRouter() (http.Handler, error) {
	r := mux.NewRouter()
	taskHandler, err := newTaskHandler()
	if err != nil {
		return nil, err
	}
	routes := []router.Route{}
	routes = append(routes, taskHandler.Routes()...)
	//Routes
	for _, rt := range routes {
		var h http.Handler
		h = http.HandlerFunc(rt.Handler)
		r.Path(rt.Path).Methods(rt.Method).Handler(h)
	}
	//Middleware
	r.Use(log.NewHTTPContextHandler(log.Root()))

	return r, nil
}
