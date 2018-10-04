package mongodb

import (
	"context"
	"os"

	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
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
	InitQueue()
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
		// Create a point and add to batch
		// tags := map[string]string{
		// 	"venue":   t.GetVenue().String(),
		// 	"product": t.GetProduct().String(),
		// 	"side":    t.GetOrderSide().String(),
		// }
		// fields := map[string]interface{}{
		// 	"price":      t.GetPrice(),
		// 	"size":       t.GetSize(),
		// 	"side":       pbAPI.OrderType_value[t.GetOrderSide().String()],
		// 	"venue_type": pbAPI.VenueType_value[t.GetVenueType().String()],
		// }
		//InsertInflux("trade", tags, fields, time.Unix(0, int64(t.Timestamp)*int64(time.Nanosecond)))
	case *pbAPI.Orderbook:

	default:
		logrus.Error("Influx not found a correct type ", t)
	}
}

// InsertMongo ...
func InsertMongo() {
	// // Create a new point batch
	// bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
	// 	Precision: "ns",
	// })
	// pt, err := client.NewPoint(name, tags, fields, t)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// bp.AddPoint(pt)
	// // Write the batch
	// influxUDPClinet.Write(bp)
}
