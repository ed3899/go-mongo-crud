package main

import (
	"context"
	"go-quickstart/config"
	_ "go-quickstart/config"
	"go-quickstart/db"
	"go-quickstart/models"
	"go-quickstart/routes"
	"log"
)

func main() {
	connectionConfig := db.ConnectionConfig{
		DB_Username: config.Get(config.C1_DB_USERNAME),
		DB_Password: config.Get(config.C1_DB_PASSWORD),
		DB_Cluster: config.Get(config.C1_DB_CLUSTER),
		DB_Name: config.Get(config.C1_DB_AIRBNB),
		DB_Collection: config.Get(config.C1_DB_AIRBNB_COLLEC_LISTINGS),
	}
	listingsCollection, mongoDisconnect := db.Connect(connectionConfig)
	defer func() {
		if err := mongoDisconnect(context.TODO()); err != nil {
			log.Panicf("There was an error while disconnecting from the database: %#v", err)
		}
	}()

	listingsCollHandler := models.GetListingsCollHandler(listingsCollection)
	routes.SetBasicCRUD("api/v1", listingsCollHandler)
	routes.Serve(":8080")
}
