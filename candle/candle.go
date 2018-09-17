package candle

import (
	"math"
	"time"

	pbCandle "github.com/maurodelazeri/lion/protobuf/candle"
)

// Bucket is the ohlc struck
type Bucket struct {
	Name           string
	Product        string
	OHLC           *[]*OHLC
	NumMaxOHLCs    int
	DurationCandle time.Duration
}

// OHLC type
type OHLC struct {
	Open     float64
	Close    float64
	High     float64
	Low      float64
	Volume   float64
	Last     float64
	Trades   int64
	Start    time.Time
	Duration time.Duration
}

// OHLCMessage handles the messages
// type OHLCMessage struct {
// 	Size  float64
// 	Price float64
// 	Side  string
// 	Time  time.Time
// }

// New Initialize the ohlc
func New(bucket string, product string) *Bucket {
	buckets := []*OHLC{}
	s := &Bucket{
		Product:        product,
		Name:           bucket,
		OHLC:           &buckets,
		NumMaxOHLCs:    0,
		DurationCandle: 1,
	}
	return s
}

// CandleByDuration create buckets by duration
func (s *Bucket) CandleByDuration(message pbCandle.OHLCMessage, buckets *[]*OHLC) {
	t := message.Time.Truncate(s.DurationCandle)

	// If there are no buckets, start one.
	if len(*buckets) == 0 {
		*buckets = append(*buckets, &OHLC{
			Open:     message.Price,
			Close:    message.Price,
			Start:    t.UTC(),
			Low:      math.MaxFloat32,
			High:     0.0,
			Duration: s.DurationCandle,
			Volume:   message.Size,
			Last:     message.Price,
		})
	}

	ohlc := (*buckets)[len(*buckets)-1]
	if (*buckets)[len(*buckets)-1].Start.Equal(t) {
		ohlc.Close = message.Price
		ohlc.Last = message.Price
		// only keep a predefined number max of buckets
		if s.NumMaxOHLCs > 0 {
			buckCurrent := len(*buckets)
			if buckCurrent > s.NumMaxOHLCs {
				removalTotal := buckCurrent - s.NumMaxOHLCs
				if removalTotal > 0 {
					*buckets = (*buckets)[:0+copy((*buckets)[0:], (*buckets)[removalTotal:])]
				}
			}
		}
	} else {
		// Time to start a new ohlc.
		*buckets = append(*buckets, &OHLC{
			Open:     message.Price,
			Close:    message.Price,
			Start:    t.UTC(),
			Low:      math.MaxFloat32,
			High:     0.0,
			Duration: s.DurationCandle,
			Volume:   message.Size,
			Last:     message.Price,
		})
		ohlc = (*buckets)[len(*buckets)-1]
	}
	ohlc.Trades++
	ohlc.High = math.Max(ohlc.High, message.Price)
	ohlc.Low = math.Min(ohlc.Low, message.Price)
	ohlc.Volume = ohlc.Volume + message.Size
	ohlc.Last = message.Price
}

// // CandleByDeals create buckets by duration
// func (s *Bucket) CandleByDeals(message OHLCMessage, buckets *[]*OHLC) {
// 	t := message.Time.Truncate(s.DurationCandle)
// 	// If there are no buckets, start one.
// 	if len(*buckets) == 0 {
// 		*buckets = append(*buckets, &OHLC{
// 			Open:     message.Price,
// 			Close:    message.Price,
// 			Start:    t.UTC(),
// 			Low:      math.MaxFloat32,
// 			High:     0.0,
// 			Duration: s.DurationCandle,
// 			Volume:   message.Size,
// 			Last:     message.Price,
// 		})
// 	}

// 	ohlc := (*buckets)[len(*buckets)-1]
// 	if (*buckets)[len(*buckets)-1].Start.Equal(t) {
// 		ohlc.Close = message.Price
// 		ohlc.Last = message.Price
// 		// only keep a predefined number max of buckets
// 		if s.NumMaxOHLCs > 0 {
// 			buckCurrent := len(*buckets)
// 			if buckCurrent > s.NumMaxOHLCs {
// 				removalTotal := buckCurrent - s.NumMaxOHLCs
// 				if removalTotal > 0 {
// 					*buckets = (*buckets)[:0+copy((*buckets)[0:], (*buckets)[removalTotal:])]
// 				}
// 			}
// 		}
// 	} else {
// 		// Time to start a new ohlc.
// 		*buckets = append(*buckets, &OHLC{
// 			Open:     message.Price,
// 			Close:    message.Price,
// 			Start:    t.UTC(),
// 			Low:      math.MaxFloat32,
// 			High:     0.0,
// 			Duration: s.DurationCandle,
// 			Volume:   message.Size,
// 			Last:     message.Price,
// 		})
// 		ohlc = (*buckets)[len(*buckets)-1]
// 	}
// 	ohlc.Trades++
// 	ohlc.High = math.Max(ohlc.High, message.Price)
// 	ohlc.Low = math.Min(ohlc.Low, message.Price)
// 	ohlc.Volume = ohlc.Volume + message.Size
// 	ohlc.Last = message.Price
// }

// // CandleByVolume create buckets by duration
// func (s *Bucket) CandleByVolume(message OHLCMessage, buckets *[]*OHLC) {
// 	t := message.Time.Truncate(s.DurationCandle)
// 	// If there are no buckets, start one.
// 	if len(*buckets) == 0 {
// 		*buckets = append(*buckets, &OHLC{
// 			Open:     message.Price,
// 			Close:    message.Price,
// 			Start:    t.UTC(),
// 			Low:      math.MaxFloat32,
// 			High:     0.0,
// 			Duration: s.DurationCandle,
// 			Volume:   message.Size,
// 			Last:     message.Price,
// 		})
// 	}

// 	ohlc := (*buckets)[len(*buckets)-1]
// 	if (*buckets)[len(*buckets)-1].Start.Equal(t) {
// 		ohlc.Close = message.Price
// 		ohlc.Last = message.Price
// 		// only keep a predefined number max of buckets
// 		if s.NumMaxOHLCs > 0 {
// 			buckCurrent := len(*buckets)
// 			if buckCurrent > s.NumMaxOHLCs {
// 				removalTotal := buckCurrent - s.NumMaxOHLCs
// 				if removalTotal > 0 {
// 					*buckets = (*buckets)[:0+copy((*buckets)[0:], (*buckets)[removalTotal:])]
// 				}
// 			}
// 		}
// 	} else {
// 		// Time to start a new ohlc.
// 		*buckets = append(*buckets, &OHLC{
// 			Open:     message.Price,
// 			Close:    message.Price,
// 			Start:    t.UTC(),
// 			Low:      math.MaxFloat32,
// 			High:     0.0,
// 			Duration: s.DurationCandle,
// 			Volume:   message.Size,
// 			Last:     message.Price,
// 		})
// 		ohlc = (*buckets)[len(*buckets)-1]
// 	}
// 	ohlc.Trades++
// 	ohlc.High = math.Max(ohlc.High, message.Price)
// 	ohlc.Low = math.Min(ohlc.Low, message.Price)
// 	ohlc.Volume = ohlc.Volume + message.Size
// 	ohlc.Last = message.Price
// }
