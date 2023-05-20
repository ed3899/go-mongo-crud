package models

import "go.mongodb.org/mongo-driver/mongo"

type CollectionHandler struct {
	*mongo.Collection
}
