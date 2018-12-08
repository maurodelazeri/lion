package shared

import "github.com/mongodb/mongo-go-driver/bson/objectid"

// MongoVenues ...
type MongoVenues struct {
	ID         objectid.ObjectID `bson:"_id"`
	Protobuf   int32             `bson:"protobuf"`
	Enabled    bool              `bson:"enabled"`
	APIKey     string            `bson:"api_key"`
	APISecret  string            `bson:"api_secret"`
	Passphrase string            `bson:"passphrase"`
	Name       string            `bson:"name"`
	Product    []struct {
		Product              int32   `bson:"product"`
		Enabled              bool    `bson:"enabled"`
		IndividualConnection bool    `bson:"individual_connection"`
		VenueName            string  `bson:"venue_name"`
		APIName              string  `bson:"api_name"`
		MinimumOrdersSize    float64 `bson:"minimum_orders_size"`
		StepSize             float64 `bson:"step_size"`
		MakerFee             float64 `bson:"maker_fee"`
		TakerFee             float64 `bson:"taker_fee"`
	} `json:"product"`
}
