package mongodb

import (
	"context"
	"log"
	"os"

	"github.com/mongodb/mongo-go-driver/core/connstring"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	uri = os.Getenv("MONGODB_CONNECTION_URL")
)

var (
	// MongoDB ...
	MongoDB *mongo.Database
	client  *mongo.Client
)

func init() {
	InitEngine()
}

// InitEngine initializes our Database Connection
func InitEngine() {
	var err error
	// parse uri by using connstring.Parse()
	connectionString, err := connstring.Parse(os.Getenv("MONGODB_CONNECTION_URL"))
	if err != nil {
		log.Fatal(err)
	}
	// if database is not specified in connectionString
	// set database name
	dbname := connectionString.Database
	if dbname == "" {
		dbname = os.Getenv("MONGODB_DATABASE_NAME")
	}
	// connect to mongo
	client, err = mongo.Connect(context.Background(), uri, nil)
	if err != nil {
		log.Fatal(err)
	}
	MongoDB = client.Database(dbname, nil)
}
