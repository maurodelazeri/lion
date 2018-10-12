package candle

import (
	"fmt"
	"time"

	"github.com/maurodelazeri/concurrency-map-slice"
	number "github.com/maurodelazeri/go-number"
	"github.com/maurodelazeri/lion/common"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
)

// http://www.blackarbs.com/blog/exploring-alternative-price-bars
// http://www.blackarbs.com/blog/

// Candle ...
type Candle struct {
	// CandlesTime ...
	CandlesTime *utils.ConcurrentMap
	// CandlesTick ...
	CandlesTick *utils.ConcurrentMap
	// CandlesVolume ...
	CandlesVolume *utils.ConcurrentMap
	// CandlesMoney ...
	CandlesMoney *utils.ConcurrentMap
}

// NewCandles ...
func NewCandles() {
	candle := new(Candle)
	candle.CandlesTime = utils.NewConcurrentMap()
	candle.CandlesTick = utils.NewConcurrentMap()
	candle.CandlesVolume = utils.NewConcurrentMap()
	candle.CandlesMoney = utils.NewConcurrentMap()
}

// CreateOrUpdateCandleBarTime ...
func (c *Candle) CreateOrUpdateCandleBarTime(venue pbAPI.Venue, product pbAPI.Product, price, volume number.Decimal, side pbAPI.Side, createdAt time.Time) {
	countBuy := 0
	countSell := 0
	if side == pbAPI.Side_BUY {
		countBuy = 1
	} else {
		countSell = 1
	}
	for _, g := range pbAPI.GranularityTime_value {
		if g != 0 {
			currentPoint := createdAt.UTC().Truncate(time.Duration(g) * time.Second).Unix()
			currentKey := fmt.Sprintf("%d:%s:%s", g, venue.String(), product.String())
			history, exist := c.CandlesTime.Get(currentKey)
			newCandle := &pbAPI.Candle{
				Venue:       venue,
				Product:     product,
				Granularity: g,
				Point:       currentPoint,
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
				// NEED A VALIDATION IN CASE A PRICE COMES OUT OF ORDER, based on the start of the candle (timestamp we can solve this)
				historyCandle := history.([]*pbAPI.Candle)[len(history.([]*pbAPI.Candle))-1]

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
				c.CandlesTime.Set(currentKey, append(history.([]*pbAPI.Candle), newCandle))
			}
		}
	}
}

// CreateOrUpdateCandleBarTick ...
// https://www.thebalance.com/tick-chart-or-1-minute-chart-for-day-trading-1030978
func (c *Candle) CreateOrUpdateCandleBarTick(venue pbAPI.Venue, product pbAPI.Product, price, volume number.Decimal, side pbAPI.Side) {
	countBuy := 0
	countSell := 0
	if side == pbAPI.Side_BUY {
		countBuy = 1
	} else {
		countSell = 1
	}
	for _, g := range pbAPI.GranularityTick_value {
		if g != 0 {
			newCandle := &pbAPI.Candle{
				Venue:       venue,
				Product:     product,
				Granularity: g,
				Point:       common.MakeTimestamp(),
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
			currentKey := fmt.Sprintf("%d:%s:%s", g, venue.String(), product.String())
			history, exist := c.CandlesTick.Get(currentKey)
			if exist {
				historyCandle := history.([]*pbAPI.Candle)[len(history.([]*pbAPI.Candle))-1]
				if historyCandle.TotalTrades <= g {
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
					// New Candle
					c.CandlesTick.Set(currentKey, append(history.([]*pbAPI.Candle), newCandle))
				}
			} else {
				// New Candle
				c.CandlesTick.Set(currentKey, append(history.([]*pbAPI.Candle), newCandle))
			}
		}
	}
}

// CreateOrUpdateCandleBarVolume ...
func (c *Candle) CreateOrUpdateCandleBarVolume(venue pbAPI.Venue, product pbAPI.Product, price, volume number.Decimal, side pbAPI.Side, volLimit float64) {
	countBuy := 0
	countSell := 0
	if side == pbAPI.Side_BUY {
		countBuy = 1
	} else {
		countSell = 1
	}
	for _, g := range pbAPI.GranularityVolume_value {
		if g != 0 {
			newCandle := &pbAPI.Candle{
				Venue:       venue,
				Product:     product,
				Granularity: g,
				Point:       common.MakeTimestamp(),
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
			currentKey := fmt.Sprintf("%d:%s:%s", g, venue.String(), product.String())
			history, exist := c.CandlesVolume.Get(currentKey)
			if exist {
				historyCandle := history.([]*pbAPI.Candle)[len(history.([]*pbAPI.Candle))-1]
				if int32(historyCandle.Total) <= g {
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
					// New Candle
					c.CandlesVolume.Set(currentKey, append(history.([]*pbAPI.Candle), newCandle))
				}
			} else {
				// New Candle
				c.CandlesVolume.Set(currentKey, append(history.([]*pbAPI.Candle), newCandle))
			}
		}
	}
}

// CreateOrUpdateCandleBarMoney ...
func (c *Candle) CreateOrUpdateCandleBarMoney(venue pbAPI.Venue, product pbAPI.Product, price, volume number.Decimal, side pbAPI.Side, maxDollars float64) {
	countBuy := 0
	countSell := 0
	if side == pbAPI.Side_BUY {
		countBuy = 1
	} else {
		countSell = 1
	}
	for _, g := range pbAPI.GranularityMoney_value {
		if g != 0 {
			newCandle := &pbAPI.Candle{
				Venue:       venue,
				Product:     product,
				Granularity: g,
				Point:       common.MakeTimestamp(),
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
			currentKey := fmt.Sprintf("%d:%s:%s", g, venue.String(), product.String())
			history, exist := c.CandlesMoney.Get(currentKey)
			if exist {
				historyCandle := history.([]*pbAPI.Candle)[len(history.([]*pbAPI.Candle))-1]
				if int32(historyCandle.Total) <= g {
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
					// New Candle
					c.CandlesMoney.Set(currentKey, append(history.([]*pbAPI.Candle), newCandle))
				}
			} else {
				// New Candle
				c.CandlesMoney.Set(currentKey, append(history.([]*pbAPI.Candle), newCandle))
			}
		}
	}
}
