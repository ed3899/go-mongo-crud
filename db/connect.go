package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectionConfig struct {
	DB_Username   string
	DB_Password   string
	DB_Cluster    string
	DB_Name       string
	DB_Collection string
}

type MongoDisconnect = func(ctx context.Context) error

func Connect[C ConnectionConfig](config C) (*mongo.Collection, MongoDisconnect) {
	// Set the options
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	cfg := ConnectionConfig(config)
	connection_URI := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", cfg.DB_Username, cfg.DB_Password, cfg.DB_Cluster)
	opts := options.Client().ApplyURI(connection_URI).SetServerAPIOptions(serverApi)

	// Establish connection
	mongoClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Unable to connect to mongodb: %#v", err)
	}

	// Ping the database
	db := mongoClient.Database(cfg.DB_Name)
	if err := db.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}
	log.Printf("Succesfully pinged deployment. Connected to MongoDB!")

	// Return connection instance and disconnect function
	return db.Collection(cfg.DB_Collection), mongoClient.Disconnect
}
