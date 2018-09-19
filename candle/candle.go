package candle

import (
	"fmt"
	"sync"
	"time"

	number "github.com/maurodelazeri/go-number"
)

// SyncCandle ...
type SyncCandle struct {
	state map[string]*Candle
	mutex *sync.Mutex
}

// SyncMap ...
type SyncMap struct {
	state map[string][]int64
	mutex *sync.Mutex
}

// Candle ...
type Candle struct {
	Venue       string
	Product     string
	Granularity int64
	Point       int64
	Open        float64
	Close       float64
	High        float64
	Low         float64
	Volume      float64
	Total       float64
	TotalTrades int64
	BuySide     int
	SellSide    int
}

// https://texlution.com/post/golang-lock-free-values-with-atomic-value/
var (
	// Candlestick ...
	Candlestick     map[string]*Candle
	SyncCandlestick = NewSyncCandle()

	// CandlesMap ...
	CandlesMap     map[string][]int64
	SyncCandlesMap = NewSyncMap()
)

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

func init() {
	Candlestick = make(map[string]*Candle)
	CandlesMap = make(map[string][]int64)
}

// NewSyncCandle ...
func NewSyncCandle() *SyncCandle {
	s := &SyncCandle{
		state: make(map[string]*Candle),
		mutex: &sync.Mutex{},
	}
	return s
}

// NewSyncMap ...
func NewSyncMap() *SyncMap {
	s := &SyncMap{
		state: make(map[string][]int64),
		mutex: &sync.Mutex{},
	}
	return s
}

// Put ...
func (s *SyncCandle) Put(key string, value *Candle) {
	s.mutex.Lock()
	s.state[key] = value
	s.mutex.Unlock()
}

// Put ...
func (s *SyncMap) Put(key string, value []int64) {
	s.mutex.Lock()
	s.state[key] = value
	s.mutex.Unlock()
}

// Get ...
func (s *SyncCandle) Get(key string) *Candle {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.state[key]
}

// Get ...
func (s *SyncMap) Get(key string) []int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.state[key]
}

// Values ...
func (s *SyncCandle) Values() chan *Candle {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	values := make(chan *Candle, len(s.state))
	for _, value := range s.state {
		values <- value
	}
	close(values)
	return values
}

// Values ...
func (s *SyncMap) Values() chan []int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	values := make(chan []int64, len(s.state))
	for _, value := range s.state {
		values <- value
	}
	close(values)
	return values
}

// CreateOrUpdateCandleTime ...
func CreateOrUpdateCandleTime(venue, product string, price, amount number.Decimal, side int32, createdAt time.Time) {
	var candle = make(map[string]*Candle)
	buy := 0
	sell := 0
	if side == 0 {
		buy = 1
	} else {
		sell = 1
	}
	for _, g := range []int64{
		//CandleGranularity1M,
		//CandleGranularity2M,
		// CandleGranularity3M,
		// CandleGranularity4M,
		//CandleGranularity5M,
		// CandleGranularity6M,
		// CandleGranularity7M,
		// CandleGranularity8M,
		// CandleGranularity9M,
		// CandleGranularity10M,
		// CandleGranularity15M,
		// CandleGranularity20M,
		// CandleGranularity30M,
		// CandleGranularity40M,
		// CandleGranularity50M,
		CandleGranularity1H,
		// CandleGranularity2H,
		// CandleGranularity3H,
		// CandleGranularity4H,
		// CandleGranularity5H,
		// CandleGranularity6H,
		// CandleGranularity7H,
		// CandleGranularity8H,
		// CandleGranularity9H,
		// CandleGranularity10H,
		// CandleGranularity11H,
		// CandleGranularity12H,
		// CandleGranularity1D,
	} {
		currentPoint := createdAt.UTC().Truncate(time.Duration(g) * time.Second).Unix()
		currentKey := fmt.Sprintf("%d:%d:%s:%s", g, currentPoint, venue, product)
		candle[currentKey] = &Candle{
			Venue:       venue,
			Product:     product,
			Granularity: g,
			Point:       currentPoint,
			Open:        price.Float64(),
			Close:       price.Float64(),
			High:        price.Float64(),
			Low:         price.Float64(),
			Volume:      amount.Float64(),
			Total:       price.Mul(amount).Float64(),
			TotalTrades: 1,
			BuySide:     buy,
			SellSide:    sell,
		}

		// If exist, we need to update it
		if c, ok := Candlestick[currentKey]; ok {
			n := candle[currentKey]
			n.Open = c.Open
			if c.High > n.High {
				n.High = c.High
			}
			if c.Low < n.Low {
				n.Low = c.Low
			}
			n.Volume = n.Volume + c.Volume
			n.Total = n.Total + c.Total
			n.TotalTrades = c.TotalTrades + 1
			n.BuySide = c.BuySide + buy
			n.SellSide = c.SellSide + sell
		} else {
			CandlesMap[fmt.Sprintf("%d:%s:%s", g, venue, product)] = append(CandlesMap[fmt.Sprintf("%d:%s:%s", g, venue, product)], currentPoint)
			SyncCandlesMap.Put(fmt.Sprintf("%d:%s:%s", g, venue, product), CandlesMap[fmt.Sprintf("%d:%s:%s", g, venue, product)])
		}
		Candlestick[currentKey] = candle[currentKey]
		SyncCandlestick.Put(currentKey, candle[currentKey])
	}
}

// CreateOrUpdateCandleVolume ...
func CreateOrUpdateCandleVolume(venue, product string, price, amount number.Decimal, side int32, volume float64) {

}

// CreateOrUpdateCandleTrades ...
func CreateOrUpdateCandleTrades(venue, product string, price, amount number.Decimal, side int32, numTrades int64) {

}
