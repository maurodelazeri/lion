package venue

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"

	centrifuge "github.com/centrifugal/centrifuge-go"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/maurodelazeri/concurrency-map-slice"
	number "github.com/maurodelazeri/go-number"
	"github.com/maurodelazeri/lion/common"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/maurodelazeri/summer/streaming"
)

// Base stores the individual venue information
type Base struct {
	Streaming          bool
	Name               string
	Verbose            bool
	Enabled            bool
	RefExecutionTime   int64
	Mode               pbAPI.SystemMode
	VenueConfig        *utils.ConcurrentMap
	LiveOrderBook      *utils.ConcurrentMap
	MaxLevelsOrderBook int
	SocketClient       *centrifuge.Client
	mutex              *sync.RWMutex
}

// Venues enforces standard functions for all venues supported in
type Venues interface {
	Setup(venue string, exch config.VenueConfig, streaming bool, maxLevelsOrderBook int, mode ...pbAPI.SystemMode)
	SetDefaults()
	Start()
	GetName() string
	IsEnabled() bool
	SetEnabled(bool)
	StartStreamingToStorage(mongo bool, influx bool)
	GetMakerFee(product pbAPI.Product, price, volume number.Decimal) (number.Decimal, error)
	GetTakerFee(product pbAPI.Product, price, volume number.Decimal) (number.Decimal, error)
	ValidateVolume(product pbAPI.Product, volume float64) error
	GetOrderBook(product pbAPI.Product) (*pbAPI.Orderbook, error)
	//GetProductDetail(product pbAPI.Product) (config.Product, error)
	SetBacktestingOrderbook(product pbAPI.Product, orderbook *pbAPI.Trade)
	GetExecutionTimestamp() int64
}

// GetName is a method that returns the name of the venue base
func (e *Base) GetName() string {
	return e.Name
}

// SetEnabled is a method that sets if the venue is enabled
func (e *Base) SetEnabled(enabled bool) {
	e.Enabled = enabled
}

// SetBacktestingOrderbook is a method that sets the current orderbook, used only on backtesting mode
func (e *Base) SetBacktestingOrderbook(product pbAPI.Product, trade *pbAPI.Trade) {
	//e.RefExecutionTime = trade.GetTimestamp()
	book := &pbAPI.Orderbook{
		Product: trade.GetProduct(),
		Venue:   trade.GetVenue(),
		Levels:  int64(len(trade.Asks)),
		//	Timestamp: trade.GetTimestamp(),
		Asks: trade.Asks,
		Bids: trade.Bids,
	}
	e.LiveOrderBook.Set(product.String(), book)
}

// GetExecutionTimestamp is a method that returns the current timestamp
func (e *Base) GetExecutionTimestamp() int64 {
	if e.Mode == pbAPI.SystemMode_BACKTESTING {
		return e.RefExecutionTime
	}
	return common.MakeTimestamp()
}

// IsEnabled is a method that returns if the current venue is enabled
func (e *Base) IsEnabled() bool {
	return e.Enabled
}

// StringInSlice returns position of string, of exist
func (e *Base) StringInSlice(a string, list []string) (int, bool) {
	for index, b := range list {
		if b == a {
			return index, true
		}
	}
	return 0, false
}

// FloatInSlice returns position of float64, of exist
func (e *Base) FloatInSlice(a float64, list []float64) (int, bool) {
	for index, b := range list {
		if b == a {
			return index, true
		}
	}
	return 0, false
}

// IntnSlice returns position of string, of exist
func (e *Base) IntnSlice(a int64, list []int64) (int, bool) {
	for index, b := range list {
		if b == a {
			return index, true
		}
	}
	return 0, false
}

// RemoveIndex removes an index from a float array
func (e *Base) RemoveIndex(s []float64, index int) []float64 {
	return append(s[:index], s[index+1:]...)
}

// FloatToString to convert a float number to a string
func (e *Base) FloatToString(number float64) string {
	return strconv.FormatFloat(number, 'f', -1, 64)
}

// PercentChange gives the difference in percentage from 2 numbers
func (e *Base) PercentChange(base float64, current float64) float64 {
	if base > 0 && current > 0 {
		absolute := base - current
		if absolute < 0 {
			absolute = absolute * -1
		}
		average := ((base + current) / 2) / 2
		result := absolute / average
		formated := fmt.Sprintf("%.4f", result)
		return e.Strfloat(formated)
	}
	return -1
}

// Let's use 20 and 30 as an example. We will need to divide the absolute difference by the average of those two numbers and express it as percentages.

// Strfloat find the absolute difference between the two: |20 - 30| = |-10| = 10
// find the average of those two numbers: (20 + 30) / 2 = 50 / 2 = 25
// divide those two: 10 / 25 = 0.4
// express it as percentages: 0.4 * 100 = 40%
func (e *Base) Strfloat(i string) float64 {
	f, _ := strconv.ParseFloat(i, 64)
	return f
}

// GetMakerFee ...
func (e *Base) GetMakerFee(product pbAPI.Product, price, volume number.Decimal) (number.Decimal, error) {
	mapping, exist := e.VenueConfig.Get(e.GetName())
	if exist {
		if prod, ok := mapping.(config.VenueConfig).Products[product.String()]; ok {
			return price.Mul(volume).Div(number.NewDecimal(100, 0)).Mul(number.FromFloat(prod.MakerFee)), nil
		}
	}
	return number.NewDecimal(0, 0), errors.New("Product not found (makerfee)")
}

// GetTakerFee ...
func (e *Base) GetTakerFee(product pbAPI.Product, price, volume number.Decimal) (number.Decimal, error) {
	mapping, exist := e.VenueConfig.Get(e.GetName())
	if exist {
		if prod, ok := mapping.(config.VenueConfig).Products[product.String()]; ok {
			return price.Mul(volume).Div(number.NewDecimal(100, 0)).Mul(number.FromFloat(prod.TakerFee)), nil
		}
	}
	return number.NewDecimal(0, 0), errors.New("Product not found (takerfee)")
}

// GetOrderBook ...
func (e *Base) GetOrderBook(product pbAPI.Product) (*pbAPI.Orderbook, error) {
	mapping, exist := e.LiveOrderBook.Get(product.String())
	if exist {
		book := mapping.(*pbAPI.Orderbook)
		if len(book.Bids) > 0 && len(book.Asks) > 0 {
			return book, nil
		}
		return new(pbAPI.Orderbook), errors.New("Bid/Asks with no values")
	}
	return new(pbAPI.Orderbook), errors.New("Product not found (GetOrderBook)")
}

// ValidateVolume ...
func (e *Base) ValidateVolume(product pbAPI.Product, volume float64) error {
	mapping, exist := e.VenueConfig.Get(e.GetName())
	if exist {
		if prod, ok := mapping.(config.VenueConfig).Products[product.String()]; ok {
			if volume <= 0 {
				return errors.New("Volume must be higher than 0")
			}
			if volume < prod.MinimumOrdersSize {
				return errors.New("Minimal volume size for this product is: " + e.FloatToString(prod.MinimumOrdersSize))
			}
			return nil
		}
	}
	return errors.New("Product not found")
}

// GetProductDetail ...
// func (e *Base) GetProductDetail(product pbAPI.Product) (config.Product, error) {
// 	var prodCondig config.Product
// 	mapping, exist := e.VenueConfig.Get(e.GetName())
// 	if exist {
// 		if prod, ok := mapping.(config.VenueConfig).Products[product.String()]; ok {
// 			return prod, nil
// 		}
// 	}
// 	return prodCondig, errors.New("Product not found")
// }

// StartStreamingToStorage load the streaming from kafka
func (e *Base) StartStreamingToStorage(mongo bool, influx bool) {
	venueConf, ok := e.VenueConfig.Get(e.GetName())
	if ok {
		pairs := []string{}
		for product := range venueConf.(config.VenueConfig).Products {
			pairs = append(pairs, product+"."+e.GetName()+".trade")
		}
		stream := streaming.InitKafkaConnection(pairs, e.GetName())
		stream.StartReading()
	}
}

// ConnToken ...
func (e *Base) ConnToken(user string, exp int64) string {
	// NOTE that JWT must be generated on backend side of your application!
	// Here we are generating it on client side only for example simplicity.
	claims := jwt.MapClaims{"sub": user}
	if exp > 0 {
		claims["exp"] = exp
	}
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SOCKET_SECRET")))
	if err != nil {
		panic(err)
	}
	return t
}
