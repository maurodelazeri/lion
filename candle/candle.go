package candle

import (
	"fmt"
	"time"

	"github.com/maurodelazeri/concurrency-map-slice"
	number "github.com/maurodelazeri/go-number"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
)

// https://texlution.com/post/golang-lock-free-values-with-atomic-value/
var (
	// Candlestick ...
	SyncCandlestick = utils.NewConcurrentMap()

	// CandlesMap ...
	SyncCandlesMap = utils.NewConcurrentMap()
)

// CreateOrUpdateCandleTime ...
func CreateOrUpdateCandleTime(venue pbAPI.Venue, product pbAPI.Product, price, amount number.Decimal, side int32, createdAt time.Time) {
	buy := 0
	sell := 0
	if side == 0 {
		buy = 1
	} else {
		sell = 1
	}
	for _, g := range pbAPI.Granularity_value {
		if g != 0 {
			currentPoint := createdAt.UTC().Truncate(time.Duration(g) * time.Second).Unix()
			currentKey := fmt.Sprintf("%d:%d:%s:%s", g, currentPoint, venue.String(), product.String())
			history, exist := SyncCandlestick.Get(currentKey)
			newCandle := &pbAPI.Candle{
				Venue:       venue,
				Product:     product,
				Granularity: g,
				Point:       int32(currentPoint),
				Open:        price.Float64(),
				Close:       price.Float64(),
				High:        price.Float64(),
				Low:         price.Float64(),
				Volume:      amount.Float64(),
				Total:       price.Mul(amount).Float64(),
				TotalTrades: 1,
				BuyTotal:    int32(buy),
				SellTotal:   int32(sell),
			}
			if exist {
				historyCandle := history.(*pbAPI.Candle)
				historyCandle.Open = newCandle.Open
				if newCandle.High > historyCandle.High {
					historyCandle.High = newCandle.High
				}
				if newCandle.Low < historyCandle.Low {
					historyCandle.Low = newCandle.Low
				}
				historyCandle.Volume = historyCandle.Volume + newCandle.Volume
				historyCandle.Total = historyCandle.Total + newCandle.Total
				historyCandle.TotalTrades = newCandle.TotalTrades + 1
				historyCandle.BuyTotal = newCandle.BuyTotal + int32(buy)
				historyCandle.SellTotal = newCandle.SellTotal + int32(sell)
			} else {
				SyncCandlestick.Set(currentKey, newCandle)
				currentPointsMap, found := SyncCandlesMap.Get(fmt.Sprintf("%d:%s:%s", g, venue.String(), product.String()))
				if found {
					SyncCandlesMap.Set(fmt.Sprintf("%d:%s:%s", g, venue.String(), product.String()), append(currentPointsMap.([]int64), currentPoint))
				} else {
					SyncCandlesMap.Set(fmt.Sprintf("%d:%s:%s", g, venue.String(), product.String()), append([]int64{}, currentPoint))
				}
			}
		}
	}
}

// CreateOrUpdateCandleVolume ...
func CreateOrUpdateCandleVolume(venue, product string, price, amount number.Decimal, side int32, volume float64) {

}

// CreateOrUpdateCandleTrades ...
func CreateOrUpdateCandleTrades(venue, product string, price, amount number.Decimal, side int32, numTrades int64) {

}
