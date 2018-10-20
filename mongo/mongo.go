package mongodb

import (
	"context"
	"os"

	"github.com/mongodb/mongo-go-driver/bson/objectid"

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

// Initialization       *ClientInitilization   `protobuf:"bytes,1,opt,name=initialization,proto3" json:"initialization,omitempty"`
// Statistics           map[string]*Statistics `protobuf:"bytes,2,rep,name=statistics,proto3" json:"statistics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
// Ticks                map[string]int32       `protobuf:"bytes,3,rep,name=ticks,proto3" json:"ticks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
// Positions            []string               `protobuf:"bytes,4,rep,name=positions,proto3" json:"positions,omitempty"`
// Comment              string                 `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`

// 		type ClientInitilization struct {
// 	SystemMode           SystemMode    `protobuf:"varint,1,opt,name=system_mode,json=systemMode,proto3,enum=api.SystemMode" json:"system_mode,omitempty"`
// 	Subscription         *Subscription `protobuf:"bytes,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
// 	StartDate            int64         `protobuf:"varint,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
// 	EndDate              int64         `protobuf:"varint,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
// 	CandleGranularity    string        `protobuf:"bytes,5,opt,name=candle_granularity,json=candleGranularity,proto3" json:"candle_granularity,omitempty"`
// 	CandleGroupBy        CandleGroupBy `protobuf:"varint,6,opt,name=candle_group_by,json=candleGroupBy,proto3,enum=api.CandleGroupBy" json:"candle_group_by,omitempty"`
// 	Verbose              bool          `protobuf:"varint,7,opt,name=verbose,proto3" json:"verbose,omitempty"`
// 	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
// 	XXX_unrecognized     []byte        `json:"-"`
// 	XXX_sizecache        int32         `json:"-"`
// }

// WorkerBacktesting execute sequencial execution based on the received instructions
func WorkerBacktesting(item interface{}) {
	switch t := item.(type) {
	case *pbAPI.BacktestingReport:
		arrVal := bson.NewArray()
		var posiDocumento *bson.Element
		for _, posi := range t.GetPositions() {
			for _, order := range posi.Orders {
				value := bson.VC.DocumentFromElements(
					bson.EC.Int32("venue", pbAPI.Venue_value[order.Venue.String()]),
					bson.EC.Int32("product", pbAPI.Product_value[order.Product.String()]),
					bson.EC.Double("volume", order.Volume),
					bson.EC.Double("left_volume", order.LeftVolume),
					bson.EC.Double("price", order.Price),
					bson.EC.Int32("type", pbAPI.OrderType_value[order.Type.String()]),
					bson.EC.Int32("side", pbAPI.Side_value[order.Side.String()]),
					bson.EC.Int32("state", pbAPI.OrderState_value[order.State.String()]),
					bson.EC.Int32("entry_type", pbAPI.OrderEntryType_value[order.EntryType.String()]),
					bson.EC.Int64("time_expiration", order.TimeExpiration),
					bson.EC.Int64("time_setup", order.TimeSetup),
					bson.EC.Int32("type_filling", pbAPI.OrderTypeFilling_value[order.TypeFilling.String()]),
					bson.EC.Int32("type_time", pbAPI.OrderTypeTime_value[order.TypeTime.String()]),
					bson.EC.Int32("reason", pbAPI.Reason_value[order.Reason.String()]),
					bson.EC.Double("fee", order.Fee),
					bson.EC.String("comment", order.Comment),
				)
				arrVal.Append(value)
			}
			accountID, _ := objectid.FromHex(posi.GetAccountId())
			posiDocumento = bson.EC.SubDocumentFromElements("positions",
				bson.EC.Int32("venue", pbAPI.Venue_value[posi.GetVenue().String()]),
				bson.EC.ObjectID("account", accountID),
				bson.EC.Int32("product", pbAPI.Product_value[posi.GetProduct().String()]),
				bson.EC.Double("price_open", posi.GetPriceOpen()),
				bson.EC.Double("weighted_price", posi.GetWeightedPrice()),
				bson.EC.Double("volume", posi.GetVolume()),
				bson.EC.Int64("position_time", posi.GetPositionTime()),
				bson.EC.Int64("closing_time", posi.GetClosingTime()),
				bson.EC.Int32("position_side", pbAPI.Side_value[posi.GetPositionSide().String()]),
				bson.EC.Int32("position_reason", pbAPI.Reason_value[posi.GetPositionReason().String()]),
				bson.EC.Double("sl", posi.GetSl()),
				bson.EC.Double("tp", posi.GetTp()),
				bson.EC.Double("swap", posi.GetSwap()),
				bson.EC.Double("trailling_percent", posi.GetTraillingPercent()),
				bson.EC.Double("profit_liquid", posi.GetProfitLiquid()),
				bson.EC.Double("cumulative_fees", posi.GetCumulativeFees()),
				bson.EC.String("comment", posi.GetComment()),
				bson.EC.Array("orders", arrVal),
			)
		}
		coll := MongoDB.Collection("trades")
		_, err := coll.InsertOne(
			context.Background(),
			bson.NewDocument(
				posiDocumento,
				bson.EC.SubDocumentFromElements("size",
					bson.EC.Int32("h", 28),
					bson.EC.Double("w", 35.5),
					bson.EC.String("uom", "cm"),
				),
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
