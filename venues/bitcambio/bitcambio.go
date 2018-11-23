package bitcambio

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/maurodelazeri/concurrency-map-slice"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	venue "github.com/maurodelazeri/lion/venues"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/maurodelazeri/lion/venues/request"
)

const (
	websocketURL = "wss://bitcambio_api.blinktrade.com/trade/"
	// Public endpoints
	// Authenticated endpoints
	// authenticated and unauthenticated limit rates
	authRate   = 1000
	unauthRate = 1000
)

// Bitcambio internals
type Bitcambio struct {
	venue.Base
	*request.Handler
}

// Websocket is the overarching type across the Bitcambio package
type Websocket struct {
	base        *Bitcambio
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
	OrderBookMAP     map[string]map[int64]BookItem
	subscribedPairs  []string
	pairsMapping     *utils.ConcurrentMap
	MessageType      []byte
}

// SetDefaults sets default values for the venue
func (r *Bitcambio) SetDefaults() {
	r.Enabled = true
	r.Base.LiveOrderBook = utils.NewConcurrentMap()
}

// Setup initialises the venue parameters with the current configuration
func (r *Bitcambio) Setup(venueName string, config config.VenueConfig, streaming bool, maxLevelsOrderBook int, mode ...pbAPI.SystemMode) {
	r.Streaming = streaming
	r.Name = venueName
	r.VenueConfig = utils.NewConcurrentMap()
	r.Base.MaxLevelsOrderBook = maxLevelsOrderBook
	r.VenueConfig.Set(venueName, config)
	if len(mode) > 0 {
		r.Mode = mode[0]
	}

	r.Handler = new(request.Handler)
	r.SetRequestHandler(r.Name, authRate, unauthRate, new(http.Client))
}

// Start ...
func (r *Bitcambio) Start() {
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

		r.LiveOrderBook = utils.NewConcurrentMap()

		if len(dedicatedSocket) > 0 {
			for _, pair := range dedicatedSocket {
				socket := new(Websocket)
				socket.MessageType = make([]byte, 4)
				socket.base = r
				socket.subscribedPairs = append(socket.subscribedPairs, pair)
				go socket.WebsocketClient()
			}
		}
		if len(sharedSocket) > 0 {
			socket := new(Websocket)
			socket.MessageType = make([]byte, 4)
			socket.base = r
			socket.subscribedPairs = sharedSocket
			go socket.WebsocketClient()
		}
	}
}
