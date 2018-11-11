package binance

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/maurodelazeri/concurrency-map-slice"
	"github.com/maurodelazeri/lion/common"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	venue "github.com/maurodelazeri/lion/venues"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/maurodelazeri/lion/venues/request"
)

const (
	apiURL       = "https://api.binance.com"
	mainWebsite  = "https://www.binance.com"
	websocketURL = "wss://stream.binance.com:9443/stream?streams="

	// Public endpoints
	exchangeInfo     = "/api/v1/exchangeInfo"
	orderBookDepth   = "/api/v1/depth"
	recentTrades     = "/api/v1/trades"
	historicalTrades = "/api/v1/historicalTrades"
	aggregatedTrades = "/api/v1/aggTrades"
	candleStick      = "/api/v1/klines"
	priceChange      = "/api/v1/ticker/24hr"
	symbolPrice      = "/api/v3/ticker/price"
	bestPrice        = "/api/v3/ticker/bookTicker"
	getAllAsset      = "/assetWithdraw/getAllAsset.html"

	// Authenticated endpoints
	newOrderTest = "/api/v3/order/test"
	newOrder     = "/api/v3/order"
	queryOrder   = "/api/v3/order"

	// binance authenticated and unauthenticated limit rates
	binanceAuthRate   = 1000
	binanceUnauthRate = 1000
)

// Binance internals
type Binance struct {
	venue.Base
	*request.Handler
}

// Websocket is the overarching type across the Binance package
type Websocket struct {
	base        *Binance
	nonce       int64
	isConnected bool
	mu          sync.Mutex
	*websocket.Conn
	dialer    *websocket.Dialer
	reqHeader http.Header
	httpResp  *http.Response
	dialErr   error

	// RecIntvlMin specifies the initial reconnecting interval,
	// default to 2 seconds
	RecIntvlMin time.Duration
	// RecIntvlMax specifies the maximum reconnecting interval,
	// default to 30 seconds
	RecIntvlMax time.Duration
	// RecIntvlFactor specifies the rate of increase of the reconnection
	// interval, default to 1.5
	RecIntvlFactor float64
	// HandshakeTimeout specifies the duration for the handshake to complete,
	// default to 2 seconds
	HandshakeTimeout time.Duration

	OrderBookMAP    map[string]map[float64]float64
	subscribedPairs []string
	pairsMapping    *utils.ConcurrentMap
	MessageType     []byte

	LockTillBookFetchToFinish map[string]string
}

// SetDefaults sets default values for the venue
func (r *Binance) SetDefaults() {
	r.Enabled = true
	r.Base.LiveOrderBook = utils.NewConcurrentMap()
}

// Setup initialises the venue parameters with the current configuration
func (r *Binance) Setup(venueName string, config config.VenueConfig, streaming bool, maxLevelsOrderBook int, mode ...pbAPI.SystemMode) {
	r.Streaming = streaming
	r.Name = venueName
	r.VenueConfig = utils.NewConcurrentMap()
	r.Base.MaxLevelsOrderBook = maxLevelsOrderBook
	r.VenueConfig.Set(venueName, config)
	if len(mode) > 0 {
		r.Mode = mode[0]
	}

	r.Handler = new(request.Handler)
	r.SetRequestHandler(r.Name, binanceAuthRate, binanceUnauthRate, new(http.Client))
}

// Start ...
func (r *Binance) Start() {
	var dedicatedSocket, sharedSocket []string
	// Individual system order book for each product
	venueConf, ok := r.VenueConfig.Get(r.GetName())
	if ok {
		for product, value := range venueConf.(config.VenueConfig).Products {
			// Separate products that will use a exclusive connection from those sharing a connection
			if value.IndividualConnection {
				dedicatedSocket = append(dedicatedSocket, product)
			} else {
				sharedSocket = append(sharedSocket, product)
			}
		}
		if len(dedicatedSocket) > 0 {
			for _, pair := range dedicatedSocket {
				socket := new(Websocket)
				socket.MessageType = make([]byte, 4)
				socket.base = r
				socket.subscribedPairs = append(socket.subscribedPairs, pair)
				socket.LockTillBookFetchToFinish = map[string]string{}
				go socket.WebsocketClient()
			}
		}
		if len(sharedSocket) > 0 {
			socket := new(Websocket)
			socket.MessageType = make([]byte, 4)
			socket.base = r
			socket.subscribedPairs = sharedSocket
			socket.LockTillBookFetchToFinish = map[string]string{}
			go socket.WebsocketClient()
		}
	}
}

// SendHTTPRequest sends an unauthenticated request
func (r *Binance) SendHTTPRequest(path string, result interface{}) error {
	return r.SendPayload("GET", path, nil, nil, result, false, r.Verbose)
}

// LoadOrderbook returns full orderbook information
//
// symbol: string of currency pair
// limit: returned limit amount
func (r *Binance) LoadOrderbook(symbol string, limit int64) (OrderBook, error) {
	resp := OrderBook{}
	params := url.Values{}
	params.Set("symbol", common.StringToUpper(symbol))
	params.Set("limit", strconv.FormatInt(limit, 10))
	path := fmt.Sprintf("%s%s?%s", apiURL, orderBookDepth, params.Encode())
	if err := r.SendHTTPRequest(path, &resp); err != nil {
		return resp, err
	}
	return resp, nil
}
