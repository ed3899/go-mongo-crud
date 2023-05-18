package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db_username string
	db_password string
	db_name     string
	db_cluster string
)

func init() {
	var present bool
	db_username, present = os.LookupEnv("DB_USERNAME")
	if !present {
		log.Fatalln("Please provide the DB_NAME environment variable while executing the app")
	}

	db_password, present = os.LookupEnv("DB_PASSWORD")
	if !present {
		log.Fatalln("Please provide the DB_PASSWORD environment variable while executing the app")
	}

	db_name, present = os.LookupEnv("DB_NAME")
	if !present {
		log.Fatalln("Please provide the DB_NAME environment variable while executing the app")
	}

	db_cluster, present = os.LookupEnv("DB_CLUSTER")
	if !present {
		log.Fatalln("Please provide the DB_CLUSTER environment variable while executing the app")
	}
}

var MongoClient *mongo.Client

func Connect() func(ctx context.Context) error {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	connection_URI := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", db_username, db_password, db_cluster)

	opts := options.Client().ApplyURI(connection_URI).SetServerAPIOptions(serverApi)

	var err error
	MongoClient, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Unable to connect to mongodb: %#v", err)
	}

	if err := MongoClient.Database(db_name).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	log.Printf("Succesfully pinged deployment. Connected to MongoDB!")

	return MongoClient.Disconnect
}
