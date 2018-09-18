package candle

import (
	"fmt"
	"time"

	number "github.com/MixinNetwork/go-number"
)

func init() {
	Candles = make(map[string]*Candle)
}

var Candles map[string]*Candle

// CandleGranularity1M ...
const (
	CandleGranularity1M  = 60
	CandleGranularity2M  = 120
	CandleGranularity3M  = 180
	CandleGranularity4M  = 240
	CandleGranularity5M  = 300
	CandleGranularity6M  = 360
	CandleGranularity7M  = 420
	CandleGranularity8M  = 480
	CandleGranularity9M  = 540
	CandleGranularity10M = 600
	CandleGranularity15M = 900
	CandleGranularity20M = 1200
	CandleGranularity30M = 1800
	CandleGranularity40M = 2400
	CandleGranularity50M = 3000
	CandleGranularity1H  = 3600
	CandleGranularity2H  = 7200
	CandleGranularity3H  = 10800
	CandleGranularity4H  = 14400
	CandleGranularity5H  = 18000
	CandleGranularity6H  = 21600
	CandleGranularity7H  = 25200
	CandleGranularity8H  = 28800
	CandleGranularity9H  = 32400
	CandleGranularity10H = 36000
	CandleGranularity11H = 39600
	CandleGranularity12H = 43200
	CandleGranularity1D  = 86400
)

// Candle ...
type Candle struct {
	Granularity int64
	Point       int64
	Open        float64
	Close       float64
	High        float64
	Low         float64
	Volume      float64
	Total       float64
}

// CreateOrUpdateCandle ...
func CreateOrUpdateCandle(base, quote string, price, amount number.Decimal, createdAt time.Time) {
	var candles = make(map[string]*Candle)
	for _, g := range []int64{
		CandleGranularity1M,
		CandleGranularity2M,
		CandleGranularity3M,
		CandleGranularity4M,
		CandleGranularity5M,
		CandleGranularity6M,
		CandleGranularity7M,
		CandleGranularity8M,
		CandleGranularity9M,
		CandleGranularity10M,
		CandleGranularity15M,
		CandleGranularity20M,
		CandleGranularity30M,
		CandleGranularity40M,
		CandleGranularity50M,
		CandleGranularity1H,
		CandleGranularity2H,
		CandleGranularity3H,
		CandleGranularity4H,
		CandleGranularity5H,
		CandleGranularity6H,
		CandleGranularity7H,
		CandleGranularity8H,
		CandleGranularity9H,
		CandleGranularity10H,
		CandleGranularity11H,
		CandleGranularity12H,
		CandleGranularity1D,
	} {
		p := createdAt.UTC().Truncate(time.Duration(g) * time.Second).Unix()
		candles[fmt.Sprintf("%d:%d", g, p)] = &Candle{
			Granularity: g,
			Point:       p,
			Open:        price.Float64(),
			Close:       price.Float64(),
			High:        price.Float64(),
			Low:         price.Float64(),
			Volume:      amount.Float64(),
			Total:       price.Mul(amount).Float64(),
		}

		if c, ok := Candles[fmt.Sprintf("%d:%d", g, p)]; ok {
			n := candles[fmt.Sprintf("%d:%d", c.Granularity, c.Point)]
			n.Open = c.Open
			if c.High > n.High {
				n.High = c.High
			}
			if c.Low < n.Low {
				n.Low = c.Low
			}
			n.Volume = n.Volume + c.Volume
			n.Total = n.Total + c.Total
		}
	}
	Candles = candles
}
