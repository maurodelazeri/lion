package bitmex

import (
	"net/http"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/maurodelazeri/zion/config"
	"github.com/maurodelazeri/zion/redis"

	"github.com/gorilla/websocket"
	"github.com/maurodelazeri/zion/common"
	"github.com/maurodelazeri/zion/exchanges"
	"github.com/maurodelazeri/zion/exchanges/request"
	"github.com/maurodelazeri/zion/exchanges/ticker"
)

const (
	bitmexAPIURL = "https://api.bitmex.com/"

	bitmexAuthRate   = 0
	bitmexUnauthRate = 0
)

var sometin []string

// Bitmex is the overarching type across the bitmex package
type Bitmex struct {
	exchange.Base
	*request.Handler

	nonce       int64
	isConnected bool
	mu          sync.Mutex
	*websocket.Conn
	dialer    *websocket.Dialer
	url       string
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

	redisSession *redis.Client
}

// SetDefaults sets default values for the exchange
func (b *Bitmex) SetDefaults() {
	b.Name = "bitmex"
	b.Enabled = false
	b.Verbose = false
	b.TakerFee = 0.25
	b.MakerFee = 0
	b.Websocket = false
	b.RESTPollingDelay = 10
	b.AssetTypes = []string{ticker.Spot}
	b.APIUrl = bitmexAPIURL
	b.url = bitmexWebsocketURL
	b.Handler = new(request.Handler)
	b.SetRequestHandler(b.Name, bitmexAuthRate, bitmexUnauthRate, new(http.Client))
	b.redisSession = redisserver.GetRedisSession()

}

// Setup initialises the exchange parameters with the current configuration
func (b *Bitmex) Setup(exch config.ExchangeConfig) {
	if !exch.Enabled {
		b.SetEnabled(false)
	} else {
		b.Enabled = true
		b.AuthenticatedAPISupport = exch.AuthenticatedAPISupport
		b.SetAPIKeys(exch.APIKey, exch.APISecret, exch.ClientID, true)
		b.RESTPollingDelay = exch.RESTPollingDelay
		b.Verbose = exch.Verbose
		b.Websocket = exch.Websocket
		b.BaseCurrencies = common.SplitStrings(exch.BaseCurrencies, ",")
		b.APIEnabledPairs = exch.APIEnabledPairs
		b.ExchangeEnabledPairs = exch.ExchangeEnabledPairs
		b.MakerFee = exch.MakerFee
		b.TakerFee = exch.TakerFee
		b.InfluxStreaming = exch.InfluxStreaming
	}
}
