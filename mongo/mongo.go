package mongodb

import (
	"context"
	"os"
	"strings"

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
	BacktestingQueue = lane.NewQueue()
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
				bson.EC.Double("volume", t.GetVolume()),
				bson.EC.Int32("side", pbAPI.Side_value[t.GetOrderSide().String()]),
				bson.EC.Int32("venue_type", pbAPI.VenueType_value[t.GetVenueType().String()]),
				bson.EC.Array("bids", arrBids),
				bson.EC.Array("asks", arrAsks),
			))
		if err != nil {
			logrus.Error("Problem to insert on mongo (trade)", err)
		}
	}
}

// WorkerBacktesting execute sequencial execution based on the received instructions
func WorkerBacktesting(item interface{}) {
	switch t := item.(type) {
	case *pbAPI.BacktestingReport:

		// Initialization       *ClientInitilization   `protobuf:"bytes,1,opt,name=initialization,proto3" json:"initialization,omitempty"`
		// Statistics           map[string]*Statistics `protobuf:"bytes,2,rep,name=statistics,proto3" json:"statistics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
		// Ticks                map[string]int32       `protobuf:"bytes,3,rep,name=ticks,proto3" json:"ticks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
		// Positions            map[string]*Position   `protobuf:"bytes,4,rep,name=positions,proto3" json:"positions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
		// Comment              string                 `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`

		// ############## Initialization ##############
		// Client init parameters
		subscriptionArr := bson.NewArray()
		for _, subscription := range t.GetInitialization().GetSubscription().GetSubscribe() {
			value := bson.VC.DocumentFromElements(
				bson.EC.Int32("venue", pbAPI.Venue_value[subscription.GetVenue().String()]),
				bson.EC.Int32("product", pbAPI.Product_value[subscription.GetProduct().String()]),
			)
			subscriptionArr.Append(value)
		}

		// Balance document
		venueBalArr := bson.NewArray()
		for _, balances := range t.Initialization.Balances {
			valueCurrency := bson.VC.DocumentFromElements(
				bson.EC.Double("available", balances.GetAvailable()),
				bson.EC.Int32("currency", pbAPI.Currency_value[balances.GetCurrency().String()]),
				bson.EC.Int32("venue", pbAPI.Venue_value[balances.GetVenue().String()]),
			)
			venueBalArr.Append(valueCurrency)
		}

		var initDocument *bson.Element
		initDocument = bson.EC.SubDocumentFromElements("initialization",
			bson.EC.Array("subscription", subscriptionArr),
			bson.EC.Array("balances", venueBalArr),
			bson.EC.Time("start_date", common.MakeTimestampFromInt64(t.Initialization.GetStartDate())),
			bson.EC.Time("end_date", common.MakeTimestampFromInt64(t.Initialization.GetEndDate())),
			bson.EC.String("candle_granularity", t.Initialization.GetCandleGranularity()),
			bson.EC.Int32("candle_group_by", pbAPI.CandleGroupBy_value[t.Initialization.GetCandleGroupBy().String()]),
		)
		// ############## - ##############

		// ############## Tick ##############
		ticksArrVal := bson.NewArray()
		for index, ticks := range t.GetTicks() {
			venueSymbol := strings.Split(index, ":")
			value := bson.VC.DocumentFromElements(
				bson.EC.Int32("venue", pbAPI.Venue_value[venueSymbol[0]]),
				bson.EC.Int32("product", pbAPI.Product_value[venueSymbol[1]]),
				bson.EC.Int32("total_ticks", ticks),
			)
			ticksArrVal.Append(value)
		}

		// ############## - ##############

		// ############## Positions ##############

		// ############## - ##############

		// ############## Positions ##############

		// ############## - ##############

		// Get current balances
		// symbols := strings.Split(execution.Request.GetProduct().String(), "_")
		// quoteCurrency, quoteExist := q.Account.Balances.Get(execution.Request.GetVenue().String() + ":" + pbAPI.Currency(pbAPI.Currency_value[symbols[1]]).String())
		// baseCurrency, baseExist := q.Account.Balances.Get(execution.Request.GetVenue().String() + ":" + pbAPI.Currency(pbAPI.Currency_value[symbols[0]]).String())
		// if !quoteExist || !baseExist {
		// 	logrus.Error()
		// 	return errors.New("Problem to update mongo, balances not found")
		// }
		// base := baseCurrency.(*Currency)
		// quote := quoteCurrency.(*Currency)

		// Positions
		// posittionsArrVal := bson.NewArray()
		// var posiDocument *bson.Element
		// for _, posi := range t.GetPositions() {
		// 	for _, order := range posi.Orders {
		// 		value := bson.VC.DocumentFromElements(
		// 			bson.EC.Int32("venue", pbAPI.Venue_value[order.GetVenue().String()]),
		// 			bson.EC.Int32("product", pbAPI.Product_value[order.GetProduct().String()]),
		// 			bson.EC.Double("volume", order.GetVolume()),
		// 			bson.EC.Double("left_volume", order.GetLeftVolume()),
		// 			bson.EC.Double("price", order.GetPrice()),
		// 			bson.EC.Int32("type", pbAPI.OrderType_value[order.GetType().String()]),
		// 			bson.EC.Int32("side", pbAPI.Side_value[order.GetSide().String()]),
		// 			bson.EC.Int32("state", pbAPI.OrderState_value[order.GetState().String()]),
		// 			bson.EC.Int32("entry_type", pbAPI.OrderEntryType_value[order.GetEntryType().String()]),
		// 			bson.EC.Int64("time_expiration", order.GetTimeExpiration()),
		// 			bson.EC.Int64("time_setup", order.GetTimeSetup()),
		// 			bson.EC.Int32("type_filling", pbAPI.OrderTypeFilling_value[order.GetTypeFilling().String()]),
		// 			bson.EC.Int32("type_time", pbAPI.OrderTypeTime_value[order.GetTypeTime().String()]),
		// 			bson.EC.Int32("reason", pbAPI.Reason_value[order.GetReason().String()]),
		// 			bson.EC.Double("fee", order.GetFee()),
		// 			bson.EC.String("comment", order.GetComment()),
		// 			// bson.EC.ArrayFromElements("balance",
		// 			// 	bson.VC.DocumentFromElements(
		// 			// 		bson.EC.Double("base_available", base.Available),
		// 			// 		bson.EC.Double("base_hold", base.Hold),
		// 			// 		bson.EC.Double("quote_available", quote.Available),
		// 			// 		bson.EC.Double("quote_hold", quote.Hold),
		// 			// 	),
		// 			// ),
		// 		)
		// 		posittionsArrVal.Append(value)
		// 	}
		// 	accountID, _ := objectid.FromHex(posi.GetAccountId())
		// 	posiDocument = bson.EC.SubDocumentFromElements("positions",
		// 		bson.EC.Int32("venue", pbAPI.Venue_value[posi.GetVenue().String()]),
		// 		bson.EC.ObjectID("account", accountID),
		// 		bson.EC.Int32("product", pbAPI.Product_value[posi.GetProduct().String()]),
		// 		bson.EC.Double("price_open", posi.GetPriceOpen()),
		// 		bson.EC.Double("weighted_price", posi.GetWeightedPrice()),
		// 		bson.EC.Double("volume", posi.GetVolume()),
		// 		bson.EC.Int64("position_time", posi.GetPositionTime()),
		// 		bson.EC.Int64("closing_time", posi.GetClosingTime()),
		// 		bson.EC.Int32("position_side", pbAPI.Side_value[posi.GetPositionSide().String()]),
		// 		bson.EC.Int32("position_reason", pbAPI.Reason_value[posi.GetPositionReason().String()]),
		// 		bson.EC.Double("sl", posi.GetSl()),
		// 		bson.EC.Double("tp", posi.GetTp()),
		// 		bson.EC.Double("swap", posi.GetSwap()),
		// 		bson.EC.Double("trailling_percent", posi.GetTraillingPercent()),
		// 		bson.EC.Double("profit_liquid", posi.GetProfitLiquid()),
		// 		bson.EC.Double("cumulative_fees", posi.GetCumulativeFees()),
		// 		bson.EC.String("comment", posi.GetComment()),
		// 		bson.EC.Array("orders", posittionsArrVal),
		// 	)
		// }
		coll := MongoDB.Collection("backtesting")
		_, err := coll.InsertOne(
			context.Background(),
			bson.NewDocument(
				bson.EC.Double("total_net_profit", 0.0),
				bson.EC.Array("ticks", ticksArrVal),
				initDocument,
				//	posiDocument,
				bson.EC.String("comment", t.GetComment()),
				bson.EC.SubDocumentFromElements("size",
					bson.EC.Int32("h", 28),
					bson.EC.Double("w", 35.5),
					bson.EC.String("uom", "cm"),
				),
			))
		if err != nil {
			logrus.Error("Problem to insert on mongo (backtesting) ", err)
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
