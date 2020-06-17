package mongodb

import (
	"github.com/PhongVX/taskmanagement/internal/pkg/config/envconfig"
	"github.com/PhongVX/taskmanagement/internal/pkg/log"
	"github.com/globalsign/mgo"
)

// LoadConfigFromEnv load mongodb configurations from environments
func LoadConfigFromEnv() *Config {
	var conf Config
	envconfig.Load(&conf)
	return &conf
}

// Dial dial to target server with Monotonic mode
func Dial(conf *Config) (*mgo.Session, error) {
	log.Infof("dialing to target MongoDB at: %v, database: %v", conf.Addrs, conf.Database)
	ms, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    conf.Addrs,
		Database: conf.Database,
		Username: conf.Username,
		Password: conf.Password,
		Timeout:  conf.Timeout,
	})
	if err != nil {
		return nil, err
	}
	ms.SetMode(mgo.Monotonic, true)
	log.Infof("successfully dialing to MongoDB at %v", conf.Addrs)
	return ms, nil
}

// DialInfo return dial info from config
func (conf *Config) DialInfo() *mgo.DialInfo {
	return &mgo.DialInfo{
		Addrs:    conf.Addrs,
		Database: conf.Database,
		Username: conf.Username,
		Password: conf.Password,
		Timeout:  conf.Timeout,
	}
}
