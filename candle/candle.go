package candle

import (
	"math"
	"time"
)

// Bucket is the ohlc struck
type Bucket struct {
	Name           string
	Product        string
	OHLC           []OHLC
	NumMaxCandles  int
	DurationCandle time.Duration
	Initialized    bool
	LastUpdateTime time.Time
}

// OHLC type
type OHLC struct {
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume float64
	Last   float64
	Trades int64
}

// OHLCMessage handles the messages
type OHLCMessage struct {
	Size  float64
	Price float64
	Side  string
	Time  time.Time
}

// CandleByDuration create buckets by duration
func (s *Bucket) CandleByDuration(message OHLCMessage) []OHLC {
	t := message.Time.Truncate(s.DurationCandle)

	// If there are no buckets, start one.
	if !s.Initialized {
		s.OHLC = append(s.OHLC, OHLC{
			Open:   message.Price,
			Close:  message.Price,
			Low:    math.MaxFloat32,
			High:   0.0,
			Volume: message.Size,
			Last:   message.Price,
		})
		s.Initialized = true
	}

	ohlc := (s.OHLC)[len(s.OHLC)-1]
	if s.LastUpdateTime.Equal(t) {
		ohlc.Close = message.Price
		ohlc.Last = message.Price
		// only keep a predefined number max of buckets
		if s.NumMaxCandles > 0 {
			buckCurrent := len(s.OHLC)
			if buckCurrent > s.NumMaxCandles {
				removalTotal := buckCurrent - s.NumMaxCandles
				if removalTotal > 0 {
					s.OHLC = (s.OHLC)[:0+copy((s.OHLC)[0:], (s.OHLC)[removalTotal:])]
				}
			}
		}
	} else {
		// Time to start a new ohlc.
		s.OHLC = append(s.OHLC, OHLC{
			Open:   message.Price,
			Close:  message.Price,
			Low:    math.MaxFloat32,
			High:   0.0,
			Volume: message.Size,
			Last:   message.Price,
		})
		ohlc = (s.OHLC)[len(s.OHLC)-1]
	}
	ohlc.Trades++
	ohlc.High = math.Max(ohlc.High, message.Price)
	ohlc.Low = math.Min(ohlc.Low, message.Price)
	ohlc.Volume = ohlc.Volume + message.Size
	ohlc.Last = message.Price
	return s.OHLC
}
