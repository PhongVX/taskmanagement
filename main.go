package main

import (
	"net/http"
	"time"

	"github.com/PhongVX/taskmanagement/internal/app/api"
	"github.com/PhongVX/taskmanagement/internal/pkg/log"
)

func main() {
	log.Infof("Initializing HTTP routing...")
	r, err := api.NewRouter()
	if err != nil {
		log.Panicf("Failed to init routing, error %v", err)
	}
	log.Infof("Creating HTTP Server...")
	srv := &http.Server{
		Addr: "0.0.0.0:8585",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}
	log.Infof("Server is listening at port %s", ":8585")
	if err := srv.ListenAndServe(); err != nil {
		log.Panicf("Failed to init server, error %v", err)
	}

}
