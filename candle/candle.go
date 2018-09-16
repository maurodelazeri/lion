package candle

import (
	"math"
	"time"
	//pbmarket "github.com/maurodelazeri/lion/protobuf/marketdata"
)

var historyRatesInterval int
var intervalCandle int
var numMaxBuckets int

func init() {

}

// Venue is the bucket struck
type Venue struct {
	Name                 string
	Product              string
	Bucket               *[]*Bucket
	NumMaxBuckets        int
	IntervalCandle       int
	HistoryRatesInterval int
}

// Bucket type
type Bucket struct {
	Product  string
	Venue    string
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

// BucketMessage handles the messages
type BucketMessage struct {
	Size  float64
	Price float64
	Side  string
	Time  time.Time
}

// New Initialize the bucket
func New(venue string, product string) *Venue {
	buckets := []*Bucket{}
	s := &Venue{
		Product:              product,
		Name:                 venue,
		Bucket:               &buckets,
		NumMaxBuckets:        200,
		IntervalCandle:       1,
		HistoryRatesInterval: 60,
	}
	return s
}

// CandleByDuration create buckets by duration
func (s *Venue) CandleByDuration(message BucketMessage, buckets *[]*Bucket, candleSize time.Duration, numMaxBuckets int) {
	t := message.Time.Truncate(candleSize)
	// If there are no buckets, start one.
	if len(*buckets) == 0 {
		*buckets = append(*buckets, &Bucket{
			Open:     message.Price,
			Close:    message.Price,
			Start:    t.UTC(),
			Low:      math.MaxFloat32,
			High:     0.0,
			Duration: candleSize,
			Volume:   message.Size,
			Last:     message.Price,
		})
	}

	bucket := (*buckets)[len(*buckets)-1]
	if (*buckets)[len(*buckets)-1].Start.Equal(t) {
		bucket.Close = message.Price
		bucket.Last = message.Price
		// only keep a predefined number max of buckets
		buckCurrent := len(*buckets)
		if buckCurrent > numMaxBuckets {
			removalTotal := buckCurrent - numMaxBuckets
			if removalTotal > 0 {
				*buckets = (*buckets)[:0+copy((*buckets)[0:], (*buckets)[removalTotal:])]
			}
		}
	} else {
		// Time to start a new bucket.
		*buckets = append(*buckets, &Bucket{
			Open:     message.Price,
			Close:    message.Price,
			Start:    t.UTC(),
			Low:      math.MaxFloat32,
			High:     0.0,
			Duration: candleSize,
			Volume:   message.Size,
			Last:     message.Price,
		})
		bucket = (*buckets)[len(*buckets)-1]
	}
	bucket.Trades++
	bucket.High = math.Max(bucket.High, message.Price)
	bucket.Low = math.Min(bucket.Low, message.Price)
	bucket.Volume = bucket.Volume + message.Size
	bucket.Last = message.Price
}

// CandleByNumberOfTrades create buckets by duration
func (s *Venue) CandleByNumberOfTrades(message BucketMessage, buckets *[]*Bucket, candleSize time.Duration, numMaxBuckets int) {
	t := message.Time.Truncate(candleSize)
	// If there are no buckets, start one.
	if len(*buckets) == 0 {
		*buckets = append(*buckets, &Bucket{
			Open:     message.Price,
			Close:    message.Price,
			Start:    t.UTC(),
			Low:      math.MaxFloat32,
			High:     0.0,
			Duration: candleSize,
			Volume:   message.Size,
			Last:     message.Price,
		})
	}

	bucket := (*buckets)[len(*buckets)-1]
	if (*buckets)[len(*buckets)-1].Start.Equal(t) {
		bucket.Close = message.Price
		bucket.Last = message.Price
		// only keep a predefined number max of buckets
		buckCurrent := len(*buckets)
		if buckCurrent > numMaxBuckets {
			removalTotal := buckCurrent - numMaxBuckets
			if removalTotal > 0 {
				*buckets = (*buckets)[:0+copy((*buckets)[0:], (*buckets)[removalTotal:])]
			}
		}
	} else {
		// Time to start a new bucket.
		*buckets = append(*buckets, &Bucket{
			Open:     message.Price,
			Close:    message.Price,
			Start:    t.UTC(),
			Low:      math.MaxFloat32,
			High:     0.0,
			Duration: candleSize,
			Volume:   message.Size,
			Last:     message.Price,
		})
		bucket = (*buckets)[len(*buckets)-1]
	}
	bucket.Trades++
	bucket.High = math.Max(bucket.High, message.Price)
	bucket.Low = math.Min(bucket.Low, message.Price)
	bucket.Volume = bucket.Volume + message.Size
	bucket.Last = message.Price
}

// CandleByVolume create buckets by duration
func (s *Venue) CandleByVolume(message BucketMessage, buckets *[]*Bucket, candleSize time.Duration, numMaxBuckets int) {
	t := message.Time.Truncate(candleSize)
	// If there are no buckets, start one.
	if len(*buckets) == 0 {
		*buckets = append(*buckets, &Bucket{
			Open:     message.Price,
			Close:    message.Price,
			Start:    t.UTC(),
			Low:      math.MaxFloat32,
			High:     0.0,
			Duration: candleSize,
			Volume:   message.Size,
			Last:     message.Price,
		})
	}

	bucket := (*buckets)[len(*buckets)-1]
	if (*buckets)[len(*buckets)-1].Start.Equal(t) {
		bucket.Close = message.Price
		bucket.Last = message.Price
		// only keep a predefined number max of buckets
		buckCurrent := len(*buckets)
		if buckCurrent > numMaxBuckets {
			removalTotal := buckCurrent - numMaxBuckets
			if removalTotal > 0 {
				*buckets = (*buckets)[:0+copy((*buckets)[0:], (*buckets)[removalTotal:])]
			}
		}
	} else {
		// Time to start a new bucket.
		*buckets = append(*buckets, &Bucket{
			Open:     message.Price,
			Close:    message.Price,
			Start:    t.UTC(),
			Low:      math.MaxFloat32,
			High:     0.0,
			Duration: candleSize,
			Volume:   message.Size,
			Last:     message.Price,
		})
		bucket = (*buckets)[len(*buckets)-1]
	}
	bucket.Trades++
	bucket.High = math.Max(bucket.High, message.Price)
	bucket.Low = math.Min(bucket.Low, message.Price)
	bucket.Volume = bucket.Volume + message.Size
	bucket.Last = message.Price
}
