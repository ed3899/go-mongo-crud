package config

import (
	"log"
	"os"
)

const (
	DB_USERNAME = "DB_USERNAME"
	DB_PASSWORD = "DB_PASSWORD"
	DB_CLUSTER  = "DB_CLUSTER"
	DB_NAME     = "sample_airbnb"
)

var _Environment map[string]string

func init() {
	environmentVariables := [3]string{DB_USERNAME, DB_PASSWORD, DB_CLUSTER}
	_Environment = make(map[string]string)
	var present bool

	for _, env := range environmentVariables {
		_Environment[env], present = os.LookupEnv(env)
		if !present {
			log.Fatalf("Please provide the %v environment variable while executing the app", env)
		}
	}
}

func Get(key string) string {
	val, ok := _Environment[key]
	if !ok {
		log.Fatalf("The environment variable:%v is not present", val)
	}
	return val
}
