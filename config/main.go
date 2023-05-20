package config

import (
	"log"
	"os"
)

const (
	C1_DB_USERNAME = "C1_DB_USERNAME"
	C1_DB_PASSWORD = "C1_DB_PASSWORD"
	C1_DB_CLUSTER  = "C1_DB_CLUSTER"
	C1_DB_AIRBNB                     = "C1_DB_AIRBNB"
	C1_DB_AIRBNB_COLLEC_LISTINGS = "C1_DB_AIRBNB_COLLEC_LISTINGS"
)

var _Environment map[string]string

func init() {
	environmentVariables := [5]string{C1_DB_USERNAME, C1_DB_PASSWORD, C1_DB_CLUSTER, C1_DB_AIRBNB, C1_DB_AIRBNB_COLLEC_LISTINGS}
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
