package models

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	ListingID string   `json:"listing_id,omitempty" bson:"listing_id,omitempty"`
	Access    string   `json:"access,omitempty" bson:"access,omitempty" `
	Address   *Address `json:"address,omitempty" bson:"address,omitempty"`
}

type ListingResponse struct {
	_id      primitive.ObjectID `bson:"_id,omitempty"`
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

func (collection *CollectionHandler) GetByPage(ctx *gin.Context) {
	// Get query parameters
	const MAX_CAPACITY = 100
	limitQuery := ctx.DefaultQuery("limit", "10")
	pageQuery := ctx.DefaultQuery("page", "1")

	// Parse query parameters
	limit, err := strconv.ParseInt(limitQuery, 10, 64)
	switch {
	case err != nil:
		err := fmt.Errorf("there was an error parsing %s", limitQuery)
		log.Print(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	case limit > MAX_CAPACITY:
		err := fmt.Errorf("the maximum 'limit' is %d", MAX_CAPACITY)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, err := strconv.ParseInt(pageQuery, 10, 64)
	switch {
	case err != nil:
		err := fmt.Errorf("there was an error parsing '%d'", page)
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	case page <= 0:
		err := fmt.Errorf("page must be greater than '%d'", page)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set MongoDB query options. Pages are zero indexed.
	opts := options.Find().SetLimit(limit).SetSkip((page - 1) * limit)

	// Execute MongoDB query
	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		err := fmt.Errorf("there was an error retrieving the listings: %#v", err)
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Parse result from MongoDB query
	var results = make([]ListingResponse, 0, limit)
	if err := cursor.All(context.TODO(), &results); err != nil {
		err := fmt.Errorf("there was an error getting the resultls from the cursor: %#v", err)
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send result
	ctx.JSON(http.StatusOK, results)
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

func (collection *CollectionHandler) Update(ctx *gin.Context) {
	// Parse params
	id := ctx.Param("id")

	// Parse JSON body of update request
	var updateRequest Listing
	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		err := fmt.Errorf("there was an error parsing the json: %#v", err)
		ctx.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create MongoDB update filter and query
	filter := bson.D{{"listing_id", id}}
	updateQuery := bson.D{{"$set", updateRequest}}

	// Run MongoDB query
	_, err := collection.UpdateOne(context.TODO(), filter, updateQuery)
	if err != nil {
		err := fmt.Errorf("there was an error updating the listing: %#v", err)
		ctx.SecureJSON(http.StatusNotModified, err)
		return
	}

	// Send result
	ctx.SecureJSON(http.StatusOK, gin.H{"message": "resource updated"})
}

func (h *CollectionHandler) Delete(ctx *gin.Context) {
}

func GetListingsCollHandler(coll *mongo.Collection) *CollectionHandler {
	return &CollectionHandler{
		coll,
	}
}
