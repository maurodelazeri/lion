package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/maurodelazeri/lion/common"
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

	// TradesQueue ...
	TradesQueue *lane.Queue

	// OrderbookQueue ...
	OrderbookQueue *lane.Queue
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
	TradesQueue = lane.NewQueue()
	OrderbookQueue = lane.NewQueue()
	// Let's handle the clients asynchronously
	go func() {
		for {
			for TradesQueue.Head() != nil {
				item := TradesQueue.Dequeue()
				WorkerTrades(item)
			}
		}
	}()
	go func() {
		for {
			for OrderbookQueue.Head() != nil {
				item := OrderbookQueue.Dequeue()
				WorkerOrderbook(item)
			}
		}
	}()
}

// WorkerTrades execute sequencial execution based on the received instructions
func WorkerTrades(item interface{}) {
	switch t := item.(type) {
	case *pbAPI.Trade:
		coll := MongoDB.Collection("trades")
		layout := "2006-01-02T15:04:05.999999Z"
		timestamp, err := time.Parse(layout, t.GetTimestamp())
		if err != nil {
			return
		}
		_, err = coll.InsertOne(
			context.Background(),
			bson.NewDocument(
				bson.EC.Int32("venue", pbAPI.Venue_value[t.GetVenue().String()]),
				bson.EC.Int32("product", pbAPI.Product_value[t.GetProduct().String()]),
				bson.EC.Time("timestamp", timestamp),
				bson.EC.Double("price", t.GetPrice()),
				bson.EC.Double("size", t.GetSize()),
				bson.EC.Int32("side", pbAPI.Side_value[t.GetOrderSide().String()]),
				bson.EC.Int32("venue_type", pbAPI.VenueType_value[t.GetVenueType().String()]),
			))
		if err != nil {
			logrus.Error("Problem to insert on mongo ", err)
		}
	}
}

// WorkerOrderbook execute sequencial execution based on the received instructions
func WorkerOrderbook(item interface{}) {
	switch t := item.(type) {
	case *pbAPI.Orderbook:
		protobufByte, err := proto.Marshal(t)
		if err != nil {
			return
		}
		layout := "2006-01-02T15:04:05.999999Z"
		timestamp, err := time.Parse(layout, t.GetTimestamp())
		if err != nil {
			return
		}
		coll := MongoDB.Collection("orderbook")
		_, err = coll.InsertOne(
			context.Background(),
			bson.NewDocument(
				bson.EC.Int32("venue", pbAPI.Venue_value[t.GetVenue().String()]),
				bson.EC.Int32("product", pbAPI.Product_value[t.GetProduct().String()]),
				bson.EC.Timestamp("timestamp", timestamp),
				bson.EC.Binary("depth", common.CompressFlate(protobufByte)),
			))
		if err != nil {
			logrus.Error("Problem to insert on mongo ", err)
		}
	default:
		logrus.Error("Influx not found a correct type ", t)
	}
}

// fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05.999999Z"))
// layout := "2006-01-02T15:04:05.999999Z"
// str := "2018-10-10T18:47:00.945274Z"
// t, err := time.Parse(layout, str)
// if err != nil {
// 	fmt.Println(err)
// }
// fmt.Println(t)
