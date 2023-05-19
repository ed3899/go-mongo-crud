package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Access    string	`bson:"access,omitempty" json:"access,omitempty"`
	Address   *Address `json:"address,omitempty" bson:"address,omitempty"`
}

type ListingResponse struct {
	ObjectID primitive.ObjectID `bson:"_id,omitempty"`
	*Listing `bson:"inline"`
}
