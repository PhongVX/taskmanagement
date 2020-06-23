package api

import (
	"os"
	"sync"

	"github.com/PhongVX/taskmanagement/internal/pkg/db/mongodb"

	"github.com/globalsign/mgo"
)

var (
	session     *mgo.Session
	sessionOnce sync.Once
)

func dialDefaultMongoDB() (*mgo.Session, error) {
	repoConf := mongodb.LoadConfigFromEnv()
	var err error
	sessionOnce.Do(func() {
		session, err = mongodb.Dial(repoConf)
	})
	if err != nil {
		return nil, err
	}
	s := session.Clone()
	return s, nil
}

func staticPath() string {
	pth := os.Getenv("STATIC_PATH")
	if pth != "" {
		return pth
	}
	return "web/dist"
}
