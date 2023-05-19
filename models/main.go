package models

type Location struct {
	Coordinates []int `json:"coordinates,omitempty" bson:"coordinates,omitempty"`
	IsLocationExact bool `json:"is_location_exact,omitempty" bson:"is_location_exact,omitempty"`
	Type string `json:"type,omitempty" bson:"type,omitempty"`
}

type Address struct {
	Country       string `json:"country,omitempty" bson:"country,omitempty"`
	CountryCode   string `json:"country_code,omitempty" bson:"country_code,omitempty"`
	GovermentArea string `json:"government_area,omitempty" bson:"government_area,omitempty"`
	Location *Location `json:"location,omitempty" bson:"location,omitempty"`
	Market string `json:"market,omitempty" bson:"market,omitempty"`
	Street string `json:"street,omitempty" bson:"street,omitempty"`
	Suburb string `json:"suburb,omitempty" bson:"suburb,omitempty"`
}

type ListingRequest struct {
	Id      string   `json:"id,omitempty" bson:"id,omitempty"`
	Access  string   `json:"access,omitempty" bson:"access,omitempty"`
	Address *Address `json:"address,omitempty" bson:"address,omitempty"`
}
