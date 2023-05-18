package main

import (
	"context"
	"go-quickstart/db"
	"log"
)

var dbDisconnect func(ctx context.Context) error

func init() {
	dbDisconnect = db.Connect()
}

func main() {
	defer func() {
		if err := dbDisconnect(context.TODO()); err != nil {
			log.Panicf("There was an error while disconnecting from the database: %#v", err)
		}
	}()
}
