package main

import (
	"context"
	"go-quickstart/db"
	"go-quickstart/routes"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

var dbDisconnect func(ctx context.Context) error

func init() {
	dbDisconnect = db.Connect()
}

type Listing struct {
	Id     string `bson:"_id"`
	Access string `bson:"access"`
}

func main() {
	defer func() {
		if err := dbDisconnect(context.TODO()); err != nil {
			log.Panicf("There was an error while disconnecting from the database: %#v", err)
		}
	}()

	coll := db.MongoClient.Database(db.Db_name).Collection("listingsAndReviews")

	var result Listing
	err := coll.FindOne(context.TODO(), bson.D{}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	routes.Serve(":8080")
}
