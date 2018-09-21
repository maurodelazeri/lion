package candle

import (
	"fmt"
	"sync"
	"time"

	number "github.com/maurodelazeri/go-number"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
)

// SyncCandle ...
type SyncCandle struct {
	state map[string]*pbAPI.Candle
	mutex *sync.Mutex
}

// SyncMap ...
type SyncMap struct {
	state map[string][]int64
	mutex *sync.Mutex
}

// https://texlution.com/post/golang-lock-free-values-with-atomic-value/
var (
	// Candlestick ...
	SyncCandlestick = NewSyncCandle()

	// CandlesMap ...
	SyncCandlesMap = NewSyncMap()
)

// NewSyncCandle ...
func NewSyncCandle() *SyncCandle {
	s := &SyncCandle{
		state: make(map[string]*pbAPI.Candle),
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
func (s *SyncCandle) Put(key string, value *pbAPI.Candle) {
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
func (s *SyncCandle) Get(key string) *pbAPI.Candle {
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
func (s *SyncCandle) Values() chan *pbAPI.Candle {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	values := make(chan *pbAPI.Candle, len(s.state))
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
			historyCandle := SyncCandlestick.Get(currentKey)
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
			if historyCandle.GetPoint() != 0 {
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
				SyncCandlestick.Put(currentKey, newCandle)
				currentPointsMap := SyncCandlesMap.Get(fmt.Sprintf("%d:%s:%s", g, venue.String(), product.String()))
				SyncCandlesMap.Put(fmt.Sprintf("%d:%s:%s", g, venue.String(), product.String()), append(currentPointsMap, currentPoint))
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
