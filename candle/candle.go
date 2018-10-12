package candle

import (
	"fmt"
	"time"

	"github.com/maurodelazeri/concurrency-map-slice"
	number "github.com/maurodelazeri/go-number"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
)

// http://www.blackarbs.com/blog/exploring-alternative-price-bars
// http://www.blackarbs.com/blog/

var (
	// SyncCandlestick ...
	SyncCandlestick *utils.ConcurrentMap
	// SyncCandlesMap ...
	SyncCandlesMap *utils.ConcurrentMap
)

func init() {
	// Candlestick ...
	SyncCandlestick = utils.NewConcurrentMap()
	// CandlesMap ...
	SyncCandlesMap = utils.NewConcurrentMap()
}

// CreateOrUpdateCandleTime ...
func CreateOrUpdateCandleTime(venue pbAPI.Venue, product pbAPI.Product, price, volume number.Decimal, side pbAPI.Side, createdAt time.Time) {
	countBuy := 0
	countSell := 0
	if side == pbAPI.Side_BUY {
		countBuy = 1
	} else {
		countSell = 1
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
				Volume:      volume.Float64(),
				Total:       price.Mul(volume).Float64(),
				TotalTrades: 1,
				BuyTotal:    0,
				SellTotal:   0,
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
				historyCandle.BuyTotal = newCandle.BuyTotal + int32(countBuy)
				historyCandle.SellTotal = newCandle.SellTotal + int32(countSell)
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

// CreateOrUpdateCandleTick ...
func CreateOrUpdateCandleTick(venue pbAPI.Venue, product pbAPI.Product, price, volume number.Decimal, side pbAPI.Side, ticks int32) {

}

// CreateOrUpdateCandleVolume ...
func CreateOrUpdateCandleVolume(venue pbAPI.Venue, product pbAPI.Product, price, volume number.Decimal, side pbAPI.Side, volLimit float64) {

}

// CreateOrUpdateCandleDollars ...
func CreateOrUpdateCandleDollars(venue pbAPI.Venue, product pbAPI.Product, price, volume number.Decimal, side pbAPI.Side, maxDollars float64) {

}
