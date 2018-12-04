package mongodb

import (
	"context"
	"os"

	"github.com/mongodb/mongo-go-driver/core/connstring"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/sirupsen/logrus"
)

var (
	// MongoDB ...
	MongoDB *mongo.Database
	// Client ...
	Client *mongo.Client
	// BacktestingQueue ...
	// BacktestingQueue *lane.Queue
)

func init() {
	InitEngine()
}

// InitEngine initializes our Database Connection
func InitEngine() {
	var err error
	var CS = connstring.ConnString{
		Hosts:    []string{os.Getenv("MONGODB_CONNECTION")},
		Username: os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASSWORD"),
	}
	Client, err = mongo.NewClientFromConnString(CS)
	if err != nil {
		logrus.Error("Problem to connect with mongo ", err)
		os.Exit(1)
	}
	// connect to mongo
	err = Client.Connect(context.Background())
	if err != nil {
		logrus.Error("Mongo ", err)
		os.Exit(1)
	}
	MongoDB = Client.Database(os.Getenv("MONGODB_DATABASE_NAME"), nil)
}

// InitQueueBacktesting ...
// func InitQueueBacktesting() {
// 	BacktestingQueue = lane.NewQueue()
// 	// Let's handle the clients asynchronously
// 	go func() {
// 		for {
// 			for BacktestingQueue.Head() != nil {
// 				item := BacktestingQueue.Dequeue()
// 				WorkerBacktesting(item)
// 			}
// 		}
// 	}()
// }

// WorkerBacktesting execute sequencial execution based on the received instructions
// func WorkerBacktesting(item interface{}) {
// 	switch t := item.(type) {
// 	case *pbAPI.BacktestingReport:

// 		// ############## Initialization ##############
// 		// Client init parameters
// 		subscriptionArr := bson.NewArray()
// 		for _, subscription := range t.GetInitialization().GetSubscription().GetSubscribe() {
// 			value := bson.VC.DocumentFromElements(
// 				bson.EC.Int32("venue", pbAPI.Venue_value[subscription.GetVenue().String()]),
// 				bson.EC.Int32("product", pbAPI.Product_value[subscription.GetProduct().String()]),
// 			)
// 			subscriptionArr.Append(value)
// 		}

// 		// Balance document
// 		venueBalArr := bson.NewArray()
// 		for _, balances := range t.Initialization.Balances {
// 			valueCurrency := bson.VC.DocumentFromElements(
// 				bson.EC.Double("available", balances.GetAvailable()),
// 				bson.EC.Int32("currency", pbAPI.Currency_value[balances.GetCurrency().String()]),
// 				bson.EC.Int32("venue", pbAPI.Venue_value[balances.GetVenue().String()]),
// 			)
// 			venueBalArr.Append(valueCurrency)
// 		}

// 		var initDocument *bson.Element
// 		initDocument = bson.EC.SubDocumentFromElements("initialization",
// 			bson.EC.Array("subscription", subscriptionArr),
// 			bson.EC.Array("balances", venueBalArr),
// 			bson.EC.Time("start_date", common.MakeTimestampFromInt64(t.Initialization.GetStartDate())),
// 			bson.EC.Time("end_date", common.MakeTimestampFromInt64(t.Initialization.GetEndDate())),
// 			bson.EC.String("candle_granularity", t.Initialization.GetCandleGranularity()),
// 			bson.EC.Int32("candle_group_by", pbAPI.CandleGroupBy_value[t.Initialization.GetCandleGroupBy().String()]),
// 		)
// 		// ############## - ##############

// 		// ############## Tick ##############
// 		ticksArrVal := bson.NewArray()
// 		for index, ticks := range t.GetTicks() {
// 			venueSymbol := strings.Split(index, ":")
// 			value := bson.VC.DocumentFromElements(
// 				bson.EC.Int32("venue", pbAPI.Venue_value[venueSymbol[0]]),
// 				bson.EC.Int32("product", pbAPI.Product_value[venueSymbol[1]]),
// 				bson.EC.Int32("total_ticks", ticks),
// 			)
// 			ticksArrVal.Append(value)
// 		}
// 		// ############## - ##############

// 		// ############## Positions ##############
// 		posittionsArrVal := bson.NewArray()
// 		for _, posi := range t.GetPositions() {
// 			ordersOntoPositionArrVal := bson.NewArray()
// 			for _, order := range posi.Orders {
// 				value := bson.VC.DocumentFromElements(
// 					bson.EC.Int32("venue", pbAPI.Venue_value[order.GetVenue().String()]),
// 					bson.EC.Int32("product", pbAPI.Product_value[order.GetProduct().String()]),
// 					bson.EC.Double("volume", order.GetVolume()),
// 					bson.EC.Double("left_volume", order.GetLeftVolume()),
// 					bson.EC.Double("price", order.GetPrice()),
// 					bson.EC.Int32("type", pbAPI.OrderType_value[order.GetType().String()]),
// 					bson.EC.Int32("side", pbAPI.Side_value[order.GetSide().String()]),
// 					bson.EC.Int32("state", pbAPI.OrderState_value[order.GetState().String()]),
// 					bson.EC.Int32("entry_type", pbAPI.OrderEntryType_value[order.GetEntryType().String()]),
// 					bson.EC.Int64("time_expiration", order.GetTimeExpiration()),
// 					bson.EC.Int64("time_setup", order.GetTimeSetup()),
// 					bson.EC.Int32("type_filling", pbAPI.OrderTypeFilling_value[order.GetTypeFilling().String()]),
// 					bson.EC.Int32("type_time", pbAPI.OrderTypeTime_value[order.GetTypeTime().String()]),
// 					bson.EC.Int32("reason", pbAPI.Reason_value[order.GetReason().String()]),
// 					bson.EC.Double("fee", order.GetFee()),
// 					bson.EC.String("comment", order.GetComment()),
// 				)
// 				ordersOntoPositionArrVal.Append(value)
// 			}
// 			accountID, _ := objectid.FromHex(posi.GetAccountId())
// 			positionData := bson.VC.DocumentFromElements(
// 				bson.EC.Int32("venue", pbAPI.Venue_value[posi.GetVenue().String()]),
// 				bson.EC.ObjectID("account", accountID),
// 				bson.EC.Int32("product", pbAPI.Product_value[posi.GetProduct().String()]),
// 				bson.EC.Double("price_open", posi.GetPriceOpen()),
// 				bson.EC.Double("weighted_price", posi.GetWeightedPrice()),
// 				bson.EC.Double("volume", posi.GetVolume()),
// 				bson.EC.Int64("position_time", posi.GetPositionTime()),
// 				bson.EC.Int64("closing_time", posi.GetClosingTime()),
// 				bson.EC.Int32("position_side", pbAPI.Side_value[posi.GetPositionSide().String()]),
// 				bson.EC.Int32("position_reason", pbAPI.Reason_value[posi.GetPositionReason().String()]),
// 				bson.EC.Double("sl", posi.GetSl()),
// 				bson.EC.Double("tp", posi.GetTp()),
// 				bson.EC.Double("swap", posi.GetSwap()),
// 				bson.EC.Double("trailling_percent", posi.GetTraillingPercent()),
// 				bson.EC.Double("profit_liquid", posi.GetProfitLiquid()),
// 				bson.EC.Double("cumulative_fees", posi.GetCumulativeFees()),
// 				bson.EC.String("comment", posi.GetComment()),
// 				bson.EC.Array("orders", ordersOntoPositionArrVal),
// 			)
// 			posittionsArrVal.Append(positionData)
// 		}
// 		// ############## - ##############

// 		//############## Statistics ##############
// 		statistics := t.GetStatistics()
// 		statisticsArrVal := bson.NewArray()
// 		for _, stats := range statistics {
// 			balanceEvolutionArr := bson.NewArray()
// 			for index, value := range stats.GetBalanceEvolution() {
// 				venueSymbol := strings.Split(index, ":")
// 				balanceArr := bson.NewArray()
// 				for _, balance := range value.GetBalance() {
// 					balanceValue := bson.VC.DocumentFromElements(
// 						bson.EC.Double("available", balance),
// 					)
// 					balanceArr.Append(balanceValue)
// 				}
// 				value := bson.VC.DocumentFromElements(
// 					bson.EC.Int32("venue", pbAPI.Venue_value[venueSymbol[0]]),
// 					bson.EC.Int32("currency", pbAPI.Product_value[venueSymbol[1]]),
// 					bson.EC.Array("balance", balanceArr), //map[string]*BalanceEvolution
// 				)
// 				balanceEvolutionArr.Append(value)
// 			}
// 			statisticData := bson.VC.DocumentFromElements(
// 				bson.EC.Double("history_quality", stats.GetHistoryQuality()),
// 				bson.EC.Double("gross_profit", stats.GetGrossProfit()),
// 				bson.EC.Double("gross_loss", stats.GetGrossLoss()),
// 				bson.EC.Double("total_net_profit", stats.GetTotalNetProfit()),
// 				bson.EC.Double("balance_drawdown_absolute", stats.GetBalanceDrawdownAbsolute()),
// 				bson.EC.Double("balance_drawdown_maximal", stats.GetBalanceDrawdownMaximal()),
// 				bson.EC.Double("balance_drawdown_relative", stats.GetBalanceDrawdownRelative()),
// 				bson.EC.Double("equity_drawdown_absolute", stats.GetEquityDrawdownAbsolute()),
// 				bson.EC.Double("equity_drawdown_maximal", stats.GetEquityDrawdownMaximal()),
// 				bson.EC.Double("equity_drawdown_relative", stats.GetEquityDrawdownRelative()),
// 				bson.EC.Double("profit_factor", stats.GetProfitFactor()),
// 				bson.EC.Double("recovery_factor", stats.GetRecoveryFactor()),
// 				bson.EC.Double("ahpr", stats.GetAhpr()),
// 				bson.EC.Double("ghpr", stats.GetGhpr()),
// 				bson.EC.Double("expected_payoff", stats.GetExpectedPayoff()),
// 				bson.EC.Double("sharpe_ratio", stats.GetSharpeRatio()),
// 				bson.EC.Double("z_score", stats.GetZScore()),
// 				bson.EC.Double("total_trades", stats.GetTotalTrades()),
// 				bson.EC.Double("total_deals", stats.GetTotalDeals()),
// 				bson.EC.Double("short_trades_won", stats.GetShortTradesWon()),
// 				bson.EC.Double("long_trades_won", stats.GetLongTradesWon()),
// 				bson.EC.Double("profit_trades_total", stats.GetProfitTradesTotal()),
// 				bson.EC.Double("loss_trades_total", stats.GetLossTradesTotal()),
// 				bson.EC.Double("largest_profit_trade", stats.GetLargestProfitTrade()),
// 				bson.EC.Double("largest_loss_trade", stats.GetLargestLossTrade()),
// 				bson.EC.Double("average_profit_trade", stats.GetAverageProfitTrade()),
// 				bson.EC.Double("average_loss_trade", stats.GetAverageLossTrade()),
// 				bson.EC.Double("maximum_consecutive_wins", stats.GetMaximumConsecutiveWins()),
// 				bson.EC.Double("maximum_consecutive_wins_count", stats.GetMaximumConsecutiveWinsCount()),
// 				bson.EC.Double("maximum_consecutive_losses", stats.GetMaximumConsecutiveLosses()),
// 				bson.EC.Double("maximum_consecutive_losses_count", stats.GetMaximalConsecutiveLossCount()),
// 				bson.EC.Double("maximal_consecutive_profit", stats.GetMaximalConsecutiveProfit()),
// 				bson.EC.Double("maximal_consecutive_profit_count", stats.GetMaximalConsecutiveProfitCount()),
// 				bson.EC.Double("maximal_consecutive_loss", stats.GetMaximalConsecutiveLoss()),
// 				bson.EC.Double("maximal_consecutive_loss_count", stats.GetMaximalConsecutiveLossCount()),
// 				bson.EC.Double("average_consecutive_wins", stats.GetAverageConsecutiveWins()),
// 				bson.EC.Double("average_consecutive_losses", stats.GetAverageConsecutiveLosses()),
// 				bson.EC.Double("wins_total_series", stats.GetWinsTotalSeries()),
// 				bson.EC.Double("loss_total_series", stats.GetLossTotalSeries()),
// 				bson.EC.Double("correlation_mfe", stats.GetCorrelationMfe()),
// 				bson.EC.Double("correlation_mae", stats.GetCorrelationMae()),
// 				bson.EC.Int64("minimal_position_holding_time", stats.GetMinimalPositionHoldingTime()),
// 				bson.EC.Int64("maximal_position_holding_time", stats.GetMaximalPositionHoldingTime()),
// 				bson.EC.Int64("average_position_holding_time", stats.GetAveragePositionHoldingTime()),
// 				bson.EC.Double("lr_slope_balance", stats.GetLrSlopeBalance()),
// 				bson.EC.Double("lr_intercept_balance", stats.GetLrInterceptBalance()),
// 				bson.EC.Double("lr_r_squared_balance", stats.GetLrRSquaredBalance()),
// 				bson.EC.Double("lr_slope_standard_error_balance", stats.GetLrSlopeStandardErrorBalance()),
// 				bson.EC.Double("lr_intercept_standard_error_balance", stats.GetLrInterceptStandardErrorBalance()),
// 				bson.EC.Array("balance_evolution", balanceEvolutionArr),
// 			)
// 			statisticsArrVal.Append(statisticData)
// 		}
// 		// ############## - ##############

// 		coll := MongoDB.Collection("backtesting")
// 		_, err := coll.InsertOne(
// 			context.Background(),
// 			bson.NewDocument(
// 				initDocument,
// 				bson.EC.Array("ticks", ticksArrVal),
// 				bson.EC.Array("positions", posittionsArrVal),
// 				bson.EC.Array("statistics", statisticsArrVal),
// 				bson.EC.String("comment", t.GetComment()),
// 			))
// 		if err != nil {
// 			logrus.Error("Problem to insert on mongo (backtesting) ", err)
// 		}
// 	}
// }

// fmt.Println(time.Now().UTC().Format("2006-01-02T15:04:05.999999Z"))
// layout := "2006-01-02T15:04:05.999999Z"
// str := "2018-10-10T18:47:00.945274Z"
// t, err := time.Parse(layout, str)
// if err != nil {
// 	fmt.Println(err)
// }
// fmt.Println(t)
