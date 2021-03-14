package mongodb

import "time"

type (
	// Config hold MongoDB configuration information
	Config struct {
		Addrs    []string      `envconfig:"MONGODB_ADDRS" default:"localhost:27017"`
		Database string        `envconfig:"MONGODB_DATABASE" default:"taskmanagement"`
		Username string        `envconfig:"MONGODB_USERNAME"`
		Password string        `envconfig:"MONGODB_PASSWORD"`
		Timeout  time.Duration `envconfig:"MONGODB_TIMEOUT" default:"10s"`
	}
)
