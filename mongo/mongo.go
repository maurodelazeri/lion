package mongodb

import (
	"context"
	"os"

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

	// BacktestingQueue ...
	BacktestingQueue *lane.Queue
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

// InitQueueTrades ....
func InitQueueTrades() {
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
}

// InitQueueBacktesting ...
func InitQueueBacktesting() {
	TradesQueue = lane.NewQueue()
	OrderbookQueue = lane.NewQueue()
	// Let's handle the clients asynchronously
	go func() {
		for {
			for BacktestingQueue.Head() != nil {
				item := BacktestingQueue.Dequeue()
				WorkerBacktesting(item)
			}
		}
	}()
}

// WorkerTrades execute sequencial execution based on the received instructions
func WorkerTrades(item interface{}) {
	switch t := item.(type) {
	case *pbAPI.Trade:
		arrBids := bson.NewArray()
		for _, values := range t.GetBids() {
			value := bson.VC.DocumentFromElements(
				bson.EC.Double("price", values.GetPrice()),
				bson.EC.Double("volume", values.GetVolume()),
			)
			arrBids.Append(value)
		}
		arrAsks := bson.NewArray()
		for _, values := range t.GetAsks() {
			value := bson.VC.DocumentFromElements(
				bson.EC.Double("price", values.GetPrice()),
				bson.EC.Double("volume", values.GetVolume()),
			)
			arrAsks.Append(value)
		}
		coll := MongoDB.Collection("trades")
		_, err := coll.InsertOne(
			context.Background(),
			bson.NewDocument(
				bson.EC.Int32("venue", pbAPI.Venue_value[t.GetVenue().String()]),
				bson.EC.Int32("product", pbAPI.Product_value[t.GetProduct().String()]),
				bson.EC.Time("timestamp", common.MakeTimestampFromInt64(t.GetTimestamp())),
				bson.EC.Double("price", t.GetPrice()),
				bson.EC.Double("size", t.GetSize()),
				bson.EC.Int32("side", pbAPI.Side_value[t.GetOrderSide().String()]),
				bson.EC.Int32("venue_type", pbAPI.VenueType_value[t.GetVenueType().String()]),
				bson.EC.Array("bids", arrBids),
				bson.EC.Array("asks", arrAsks),
			))
		if err != nil {
			logrus.Error("Problem to insert on mongo ", err)
		}
	}
}

// WorkerBacktesting execute sequencial execution based on the received instructions
func WorkerBacktesting(item interface{}) {
	switch t := item.(type) {
	case *pbAPI.BacktestingReport:
		arrBids := bson.NewArray()
		for _, values := range t.GetPositions() {
			value := bson.VC.DocumentFromElements(
				bson.EC.Double("price", values.GetPrice()),
			)
			arrBids.Append(value)
		}
		arrAsks := bson.NewArray()
		for _, values := range t.GetAsks() {
			value := bson.VC.DocumentFromElements(
				bson.EC.Double("price", values.GetPrice()),
				bson.EC.Double("volume", values.GetVolume()),
			)
			arrAsks.Append(value)
		}
		coll := MongoDB.Collection("backtesting")
		_, err := coll.InsertOne(
			context.Background(),
			bson.NewDocument(
				bson.EC.Int32("venue", positRef.Orders[0].Venue),
				bson.EC.ObjectID("account", q.Account.ID),
				bson.EC.Int32("product", positRef.Orders[0].Product),
				bson.EC.Double("price_open", positRef.Orders[0].Price),
				bson.EC.Double("weighted_price", positRef.Orders[0].Price),
				bson.EC.Double("volume", positRef.Orders[0].Volume),
				bson.EC.Int64("position_time", positRef.Orders[0].TimeSetup),
				bson.EC.Int64("closing_time", 0),
				bson.EC.Int32("position_side", positRef.Orders[0].Side),
				bson.EC.Int32("position_reason", positRef.Orders[0].Reason),
				bson.EC.Double("sl", positRef.Orders[0].Sl),
				bson.EC.Double("tp", positRef.Orders[0].Tp),
				bson.EC.Double("swap", 0.0),
				bson.EC.Double("trailling_percent", positRef.Orders[0].TraillingPercent),
				bson.EC.Double("profit_liquid", 0.0),
				bson.EC.Double("cumulative_fees", execution.Fee),
				bson.EC.String("comment", positRef.Orders[0].Comment),
				bson.EC.Array("orders", arrVal),
			))
		if err != nil {
			logrus.Error("Problem to insert on mongo ", err)
		}
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
