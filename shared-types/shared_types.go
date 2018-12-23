package shared

import "github.com/mongodb/mongo-go-driver/bson/objectid"

// MongoVenue ...
type MongoVenue struct {
	ID         objectid.ObjectID `bson:"_id"`
	Venue      int32             `bson:"venue_id"`
	Enabled    bool              `bson:"enabled"`
	APIKey     string            `bson:"api_key"`
	APISecret  string            `bson:"api_secret"`
	Passphrase string            `bson:"passphrase"`
	Name       string            `bson:"name"`
}

// MongoProduct ...
type MongoProduct struct {
	ID                   objectid.ObjectID `bson:"_id"`
	VenueID              int32             `bson:"venue_id"`
	Product              int32             `bson:"product"`
	Enabled              bool              `bson:"enabled"`
	IndividualConnection bool              `bson:"individual_connection"`
	VenueName            string            `bson:"venue_name"`
	MinimumOrdersSize    float64           `bson:"minimum_orders_size"`
	StepSize             float64           `bson:"step_size"`
	MakerFee             float64           `bson:"maker_fee"`
	TakerFee             float64           `bson:"taker_fee"`
}

// MongoUser ...
type MongoUser struct {
	ID       objectid.ObjectID `bson:"_id"`
	Email    string            `bson:"email"`
	Username string            `bson:"username"`
	Password string            `bson:"password"`
	Active   bool              `bson:"active"`
}

// Account ...
type Account struct {
	ID          objectid.ObjectID `bson:"_id"`
	UserID      objectid.ObjectID `bson:"user_id"`
	AccountMode int32             `bson:"account_mode"`
	AccountType int32             `bson:"account_type"`
	Active      bool              `bson:"active"`
}
