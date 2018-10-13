package candle

import (
	"fmt"
	"time"

	"github.com/maurodelazeri/concurrency-map-slice"
	number "github.com/maurodelazeri/go-number"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/sirupsen/logrus"
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
func NewCandles() *Candle {
	candle := new(Candle)
	candle.CandlesTime = utils.NewConcurrentMap()
	candle.CandlesTick = utils.NewConcurrentMap()
	candle.CandlesVolume = utils.NewConcurrentMap()
	candle.CandlesMoney = utils.NewConcurrentMap()
	return candle
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
				historyCandle := history.([]*pbAPI.Candle)[len(history.([]*pbAPI.Candle))-1]
				if currentPoint > historyCandle.Point {
					// New candle
					c.CandlesTime.Set(currentKey, append(history.([]*pbAPI.Candle), newCandle))
					continue
				}

				if currentPoint < historyCandle.Point {
					// candle our of order
					c.CandlesTime.Set(currentKey, append(history.([]*pbAPI.Candle), newCandle))
					for i := len(history.([]*pbAPI.Candle)) - 1; i > -1; i-- {
						found := false
						if history.([]*pbAPI.Candle)[i].Point == currentPoint {
							found = true
							historyCandle = history.([]*pbAPI.Candle)[i]
						}
						if !found {
							logrus.Warn("Candle out of order, ignoring")
						}
					}
					continue
				}
				// Just update
				historyCandle.Open = newCandle.GetOpen()
				if newCandle.GetHigh() > historyCandle.GetHigh() {
					historyCandle.High = newCandle.GetLow()
				}
				if newCandle.Low < historyCandle.Low {
					historyCandle.Low = newCandle.GetLow()
				}
				historyCandle.Volume = historyCandle.GetVolume() + volume.Float64()
				historyCandle.Total = historyCandle.GetTotal() + price.Mul(volume).Float64()
				historyCandle.TotalTrades = historyCandle.GetTotalTrades() + 1
				historyCandle.BuyTotal = historyCandle.GetBuyTotal() + int32(countBuy)
				historyCandle.SellTotal = historyCandle.GetSellTotal() + int32(countSell)
			} else {
				var arr []*pbAPI.Candle
				c.CandlesTime.Set(currentKey, append(arr, newCandle))
			}
		}
	}
}

// CreateOrUpdateCandleBarTick ...
// https://www.thebalance.com/tick-chart-or-1-minute-chart-for-day-trading-1030978
func (c *Candle) CreateOrUpdateCandleBarTick(venue pbAPI.Venue, product pbAPI.Product, price, volume number.Decimal, side pbAPI.Side, createdAt time.Time) {
	countBuy := 0
	countSell := 0
	if side == pbAPI.Side_BUY {
		countBuy = 1
	} else {
		countSell = 1
	}
	for _, g := range pbAPI.GranularityTick_value {
		currentPoint := createdAt.UTC().Truncate(time.Duration(g) * time.Second).Unix()
		if g != 0 {
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
			currentKey := fmt.Sprintf("%d:%s:%s", g, venue.String(), product.String())
			history, exist := c.CandlesTick.Get(currentKey)
			if exist {
				historyCandle := history.([]*pbAPI.Candle)[len(history.([]*pbAPI.Candle))-1]

				if currentPoint < historyCandle.Point {
					// candle our of order
					c.CandlesTime.Set(currentKey, append(history.([]*pbAPI.Candle), newCandle))
					for i := len(history.([]*pbAPI.Candle)) - 1; i > -1; i-- {
						found := false
						if history.([]*pbAPI.Candle)[i].Point == currentPoint {
							found = true
							historyCandle = history.([]*pbAPI.Candle)[i]
						}
						if !found {
							logrus.Warn("Candle out of order, ignoring")
						}
					}
					continue
				}

				if historyCandle.GetTotalTrades()+1 <= g {
					historyCandle.Open = newCandle.GetOpen()
					if newCandle.GetHigh() > historyCandle.GetHigh() {
						historyCandle.High = newCandle.GetLow()
					}
					if newCandle.Low < historyCandle.Low {
						historyCandle.Low = newCandle.GetLow()
					}
					historyCandle.Volume = historyCandle.GetVolume() + volume.Float64()
					historyCandle.Total = historyCandle.GetTotal() + price.Mul(volume).Float64()
					historyCandle.TotalTrades = historyCandle.GetTotalTrades() + 1
					historyCandle.BuyTotal = historyCandle.GetBuyTotal() + int32(countBuy)
					historyCandle.SellTotal = historyCandle.GetSellTotal() + int32(countSell)
				} else {
					// New Candle
					var arr []*pbAPI.Candle
					c.CandlesTime.Set(currentKey, append(arr, newCandle))
				}
			} else {
				// New Candle
				var arr []*pbAPI.Candle
				c.CandlesTime.Set(currentKey, append(arr, newCandle))
			}
		}
	}
}

// CreateOrUpdateCandleBarVolume ...
func (c *Candle) CreateOrUpdateCandleBarVolume(venue pbAPI.Venue, product pbAPI.Product, price, volume number.Decimal, side pbAPI.Side, createdAt time.Time) {
	countBuy := 0
	countSell := 0
	if side == pbAPI.Side_BUY {
		countBuy = 1
	} else {
		countSell = 1
	}
	for _, g := range pbAPI.GranularityVolume_value {
		if g != 0 {
			currentPoint := createdAt.UTC().Truncate(time.Duration(g) * time.Second).Unix()
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
			currentKey := fmt.Sprintf("%d:%s:%s", g, venue.String(), product.String())
			history, exist := c.CandlesVolume.Get(currentKey)
			if exist {
				historyCandle := history.([]*pbAPI.Candle)[len(history.([]*pbAPI.Candle))-1]
				if currentPoint < historyCandle.Point {
					// candle our of order
					c.CandlesVolume.Set(currentKey, append(history.([]*pbAPI.Candle), newCandle))
					for i := len(history.([]*pbAPI.Candle)) - 1; i > -1; i-- {
						found := false
						if history.([]*pbAPI.Candle)[i].Point == currentPoint {
							found = true
							historyCandle = history.([]*pbAPI.Candle)[i]
						}
						if !found {
							logrus.Warn("Candle out of order, ignoring")
						}
					}
					continue
				}
				if historyCandle.GetVolume()+volume.Float64() <= float64(g) {
					historyCandle.Open = newCandle.GetOpen()
					if newCandle.GetHigh() > historyCandle.GetHigh() {
						historyCandle.High = newCandle.GetLow()
					}
					if newCandle.Low < historyCandle.Low {
						historyCandle.Low = newCandle.GetLow()
					}
					historyCandle.Volume = historyCandle.GetVolume() + volume.Float64()
					historyCandle.Total = historyCandle.GetTotal() + price.Mul(volume).Float64()
					historyCandle.TotalTrades = historyCandle.GetTotalTrades() + 1
					historyCandle.BuyTotal = historyCandle.GetBuyTotal() + int32(countBuy)
					historyCandle.SellTotal = historyCandle.GetSellTotal() + int32(countSell)
				} else {
					// New Candle
					var arr []*pbAPI.Candle
					c.CandlesVolume.Set(currentKey, append(arr, newCandle))
				}
			} else {
				// New Candle
				var arr []*pbAPI.Candle
				c.CandlesVolume.Set(currentKey, append(arr, newCandle))
			}
		}
	}
}

// CreateOrUpdateCandleBarMoney ...
func (c *Candle) CreateOrUpdateCandleBarMoney(venue pbAPI.Venue, product pbAPI.Product, price, volume number.Decimal, side pbAPI.Side, createdAt time.Time) {
	countBuy := 0
	countSell := 0
	if side == pbAPI.Side_BUY {
		countBuy = 1
	} else {
		countSell = 1
	}
	for _, g := range pbAPI.GranularityMoney_value {
		if g != 0 {
			currentPoint := createdAt.UTC().Truncate(time.Duration(g) * time.Second).Unix()
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
			currentKey := fmt.Sprintf("%d:%s:%s", g, venue.String(), product.String())
			history, exist := c.CandlesMoney.Get(currentKey)
			if exist {
				historyCandle := history.([]*pbAPI.Candle)[len(history.([]*pbAPI.Candle))-1]

				if currentPoint < historyCandle.Point {
					// candle our of order
					c.CandlesMoney.Set(currentKey, append(history.([]*pbAPI.Candle), newCandle))
					for i := len(history.([]*pbAPI.Candle)) - 1; i > -1; i-- {
						found := false
						if history.([]*pbAPI.Candle)[i].Point == currentPoint {
							found = true
							historyCandle = history.([]*pbAPI.Candle)[i]
						}
						if !found {
							logrus.Warn("Candle out of order, ignoring")
						}
					}
					continue
				}

				if historyCandle.GetTotal()+price.Mul(volume).Float64() <= float64(g) {
					historyCandle.Open = newCandle.GetOpen()
					if newCandle.GetHigh() > historyCandle.GetHigh() {
						historyCandle.High = newCandle.GetLow()
					}
					if newCandle.Low < historyCandle.Low {
						historyCandle.Low = newCandle.GetLow()
					}
					historyCandle.Volume = historyCandle.GetVolume() + volume.Float64()
					historyCandle.Total = historyCandle.GetTotal() + price.Mul(volume).Float64()
					historyCandle.TotalTrades = historyCandle.GetTotalTrades() + 1
					historyCandle.BuyTotal = historyCandle.GetBuyTotal() + int32(countBuy)
					historyCandle.SellTotal = historyCandle.GetSellTotal() + int32(countSell)
				} else {
					// New Candle
					var arr []*pbAPI.Candle
					c.CandlesMoney.Set(currentKey, append(arr, newCandle))
				}
			} else {
				// New Candle
				var arr []*pbAPI.Candle
				c.CandlesMoney.Set(currentKey, append(arr, newCandle))
			}
		}
	}
}
