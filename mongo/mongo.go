package mongodb

import (
	"context"
	"os"

	"github.com/mongodb/mongo-go-driver/core/connstring"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/sirupsen/logrus"
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
	connectionString, err := connstring.Parse(os.Getenv("MONGODB_CONNECTION_URL"))
	if err != nil {
		logrus.Error("Mongo ", err)
		os.Exit(1)
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
		logrus.Error("Mongo ", err)
		os.Exit(1)
	}
	MongoDB = client.Database(dbname, nil)
}
