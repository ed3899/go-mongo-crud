package routes

import (
	"context"
	"fmt"
	"go-quickstart/config"
	"go-quickstart/db"
	"go-quickstart/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbDisconnect func(ctx context.Context) error
var mongoClient *mongo.Client
var router *gin.Engine

func init() {
	mongoClient, dbDisconnect = db.Connect()
	db := mongoClient.Database(config.DB_NAME)

	router = gin.Default()
	v1 := router.Group("/v1")

	{
		v1.POST("/listing", func(ctx *gin.Context) {
			var newListing = new(models.Listing)
			var newListingResponse = new(models.ListingResponse)
			coll := db.Collection("listingsAndReviews")
			uuid := uuid.NewV4().String()

			if err := ctx.ShouldBindJSON(newListing); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			newListing.ListingID = uuid
			_, err := coll.InsertOne(context.TODO(), newListing)

			log.Printf("after insertion %#v", newListing)
			if err != nil {
				err := fmt.Errorf("something went wrong while inserting a new listing: %#v", err)
				log.Print(err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			log.Printf("using the uuid: %v", uuid)
			err = coll.FindOne(context.TODO(), bson.D{{"listing_id", uuid}}).Decode(newListingResponse)
			if err != nil {
				err := fmt.Errorf("something went wrong retrieving the last inserted listing: %#v", err)
				log.Print(err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			log.Printf("%#v", newListingResponse)

			ctx.JSON(http.StatusCreated, newListingResponse)
		})
	}

	{
		v1.GET("/listing/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			coll := db.Collection("listingsAndReviews")
			var ListingResponse = new(models.ListingResponse)

			if err := coll.FindOne(context.TODO(), bson.D{{"listing_id", id}}).Decode(ListingResponse); err != nil {
				err := fmt.Errorf("something went wrong retrieving the listing: %#v", err)
				log.Print(err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			ctx.JSON(http.StatusOK, ListingResponse)

		})

	}

}

func Serve(addr ...string) {
	defer func() {
		if err := dbDisconnect(context.TODO()); err != nil {
			log.Panicf("There was an error while disconnecting from the database: %#v", err)
		}
	}()
	router.Run(addr...)
}
