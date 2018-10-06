package mongodb

import (
	"context"
	"os"

	"github.com/gogo/protobuf/proto"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/core/connstring"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/oleiade/lane"
	"github.com/sirupsen/logrus"
)

var (
	uri = os.Getenv("MONGODB_CONNECTION_URL")
)

var (
	// MongoDB ...
	MongoDB *mongo.Database
	// Client ...
	Client *mongo.Client

	// MongoQueue ...
	MongoQueue *lane.Queue
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
	Client, err = mongo.Connect(context.Background(), uri, nil)
	if err != nil {
		logrus.Error("Mongo ", err)
		os.Exit(1)
	}
	MongoDB = Client.Database(dbname, nil)
}

// InitQueue the operations are in a queue to guarantee the correct execution order
func InitQueue() {
	MongoQueue = lane.NewQueue()
	// Let's handle the clients asynchronously
	go func() {
		for {
			for MongoQueue.Head() != nil {
				item := MongoQueue.Dequeue()
				Worker(item)
			}
		}
	}()
}

// Worker execute sequencial execution based on the received instructions
func Worker(item interface{}) {
	switch t := item.(type) {
	case *pbAPI.Trade:
		coll := MongoDB.Collection("trades")
		_, err := coll.InsertOne(
			context.Background(),
			bson.NewDocument(
				bson.EC.Int32("venue", pbAPI.Venue_value[t.GetVenue().String()]),
				bson.EC.Int32("product", pbAPI.Product_value[t.GetProduct().String()]),
				bson.EC.Int64("timestamp", t.GetTimestamp()),
				bson.EC.Double("price", t.GetPrice()),
				bson.EC.Double("size", t.GetSize()),
				bson.EC.Int32("side", pbAPI.Side_value[t.GetOrderSide().String()]),
				bson.EC.Int32("venue_type", pbAPI.VenueType_value[t.GetVenueType().String()]),
			))
		if err != nil {
			logrus.Error("Problem to insert on mongo ", err)
		}
	case *pbAPI.Orderbook:
		protobufByte, err := proto.Marshal(t)
		if err != nil {
			return
		}
		// compressed := common.CompressLZW(string(protobufByte))
		// fixByteArray := []byte{}
		// for i := range compressed {
		// 	n := compressed[len(compressed)-1-i]
		// 	fixByteArray = append(fixByteArray, byte(n))
		// }
		// for i, j := 0, len(fixByteArray)-1; i < j; i, j = i+1, j-1 {
		// 	fixByteArray[i], fixByteArray[j] = fixByteArray[j], fixByteArray[i]
		// }
		coll := MongoDB.Collection("orderbook")
		_, err = coll.InsertOne(
			context.Background(),
			bson.NewDocument(
				bson.EC.Int32("venue", pbAPI.Venue_value[t.GetVenue().String()]),
				bson.EC.Int32("product", pbAPI.Product_value[t.GetProduct().String()]),
				bson.EC.Int64("timestamp", t.GetTimestamp()),
				bson.EC.Binary("depth", protobufByte),
			))
		if err != nil {
			logrus.Error("Problem to insert on mongo ", err)
		}
	default:
		logrus.Error("Influx not found a correct type ", t)
	}
}
