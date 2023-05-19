package config

import (
	"log"
	"os"
)

var (
	Db_username string
	Db_password string
	Db_cluster  string
)

const DB_NAME = "sample_airbnb"

func init() {
	var present bool
	Db_username, present = os.LookupEnv("DB_USERNAME")
	if !present {
		log.Fatalln("Please provide the DB_NAME environment variable while executing the app")
	}

	Db_password, present = os.LookupEnv("DB_PASSWORD")
	if !present {
		log.Fatalln("Please provide the DB_PASSWORD environment variable while executing the app")
	}

	Db_cluster, present = os.LookupEnv("DB_CLUSTER")
	if !present {
		log.Fatalln("Please provide the DB_CLUSTER environment variable while executing the app")
	}
}
