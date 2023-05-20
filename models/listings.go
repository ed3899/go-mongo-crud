package models

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Location struct {
	Coordinates     []float64 `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
	IsLocationExact bool      `json:"is_location_exact,omitempty" bson:"is_location_exact,omitempty"`
	Type            string    `json:"type,omitempty" bson:"type,omitempty"`
}

type Address struct {
	Country       string    `json:"country,omitempty" bson:"country,omitempty"`
	CountryCode   string    `json:"country_code,omitempty" bson:"country_code,omitempty"`
	GovermentArea string    `json:"government_area,omitempty" bson:"government_area,omitempty"`
	Location      *Location `json:"location,omitempty" bson:"location,omitempty"`
	Market        string    `json:"market,omitempty" bson:"market,omitempty"`
	Street        string    `json:"street,omitempty" bson:"street,omitempty"`
	Suburb        string    `json:"suburb,omitempty" bson:"suburb,omitempty"`
}

type Listing struct {
	ListingID string   `bson:"listing_id,omitempty" json:"listing_id,omitempty"`
	Access    string   `bson:"access,omitempty" json:"access,omitempty"`
	Address   *Address `json:"address,omitempty" bson:"address,omitempty"`
}

type ListingResponse struct {
	ObjectID primitive.ObjectID `bson:"_id,omitempty"`
	*Listing `bson:"inline"`
}

func (collection *CollectionHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	var ListingResponse = new(ListingResponse)

	if err := collection.FindOne(context.TODO(), bson.D{{"listing_id", id}}).Decode(ListingResponse); err != nil {
		err := fmt.Errorf("something went wrong retrieving the listing: %#v", err)
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ListingResponse)
}

func (collection *CollectionHandler) GetAll(ctx *gin.Context) {
}

func (collection *CollectionHandler) Create(ctx *gin.Context) {
	var newListing = new(Listing)
	var newListingResponse = new(ListingResponse)
	uuid := uuid.NewV4().String()

	if err := ctx.ShouldBindJSON(newListing); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newListing.ListingID = uuid
	_, err := collection.InsertOne(context.TODO(), newListing)

	if err != nil {
		err := fmt.Errorf("something went wrong while inserting a new listing: %#v", err)
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("using the uuid: %v", uuid)

	// Code duplication TODO
	err = collection.FindOne(context.TODO(), bson.D{{"listing_id", uuid}}).Decode(newListingResponse)
	if err != nil {
		err := fmt.Errorf("something went wrong retrieving the last inserted listing: %#v", err)
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, newListingResponse)
}

func (h *CollectionHandler) Update(ctx *gin.Context) {
}

func (h *CollectionHandler) Delete(ctx *gin.Context) {
}

func GetListingsCollHandler(coll *mongo.Collection) *CollectionHandler {
	return &CollectionHandler{
		coll,
	}
}
