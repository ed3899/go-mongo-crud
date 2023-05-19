package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_ "go-quickstart/config"
	"go-quickstart/config"
)

func Connect() (*mongo.Client, func(ctx context.Context) error) {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	connection_URI := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", config.Get(config.DB_USERNAME), config.Get(config.DB_PASSWORD), config.Get(config.DB_CLUSTER))

	opts := options.Client().ApplyURI(connection_URI).SetServerAPIOptions(serverApi)

	var err error
	MongoClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Unable to connect to mongodb: %#v", err)
	}

	if err := MongoClient.Database(config.DB_NAME).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	log.Printf("Succesfully pinged deployment. Connected to MongoDB!")

	return MongoClient, MongoClient.Disconnect
}
