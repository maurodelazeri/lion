package poloniex

import (
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	utils "github.com/maurodelazeri/concurrency-map-slice"
	pbAPI "github.com/maurodelazeri/lion/protobuf/api"
	"github.com/maurodelazeri/lion/socket"
	venue "github.com/maurodelazeri/lion/venues"
	"github.com/maurodelazeri/lion/venues/config"
	"github.com/maurodelazeri/lion/venues/request"
	"github.com/sirupsen/logrus"
)

const (
	websocketURL             = "wss://api2.poloniex.com"
	wsAccountNotificationID  = 1000
	wsTickerDataID           = 1002
	ws24HourExchangeVolumeID = 1003
	wsHeartbeat              = 1010
	// Public endpoints
	// Authenticated endpoints
	// authenticated and unauthenticated limit rates
	authRate   = 1000
	unauthRate = 1000
)

// Poloniex internals
type Poloniex struct {
	venue.Base
	*request.Handler
}

// Websocket is the overarching type across the Poloniex package
type Websocket struct {
	base        *Poloniex
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

	OrderBookMAP        map[string]map[float64]float64
	subscribedPairs     []string
	pairsMapping        *utils.ConcurrentMap
	OrderbookTimestamps *utils.ConcurrentMap
}

// SetDefaults sets default values for the venue
func (r *Poloniex) SetDefaults() {
	r.Enabled = true
	r.LiveOrderBook = utils.NewConcurrentMap()
}

// Setup initialises the venue parameters with the current configuration
func (r *Poloniex) Setup(venueName string, config config.VenueConfig, streaming bool, maxLevelsOrderBook int, mode ...pbAPI.SystemMode) {
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
func (r *Poloniex) Start() {
	var dedicatedSocket, sharedSocket []string
	// Individual system order book for each product
	venueConf, ok := r.VenueConfig.Get(r.GetName())
	if ok {
		for product, value := range venueConf.(config.VenueConfig).Products {
			// Separate products that will use a exclusive connection from those sharing a connection
			if value.Enabled {
				if value.IndividualConnection {
					dedicatedSocket = append(dedicatedSocket, product)
				} else {
					sharedSocket = append(sharedSocket, product)
				}
			}
		}

		r.LiveOrderBook = utils.NewConcurrentMap()

		logrus.Infof("Initializing Socket Server")
		r.Base.SocketClient = socket.InitSocketEngine(os.Getenv("WINTER_CONTAINER_EVENT"), 0, "datafeed:winter")

		if len(dedicatedSocket) > 0 {
			for _, pair := range dedicatedSocket {
				socket := new(Websocket)
				socket.base = r
				socket.subscribedPairs = append(socket.subscribedPairs, pair)
				go socket.WebsocketClient()
			}
		}
		if len(sharedSocket) > 0 {
			socket := new(Websocket)
			socket.base = r
			socket.subscribedPairs = sharedSocket
			go socket.WebsocketClient()
		}
	}
}

// CurrencyPairID contains a list of IDS for currency pairs.
var CurrencyPairID = map[int]string{
	7:   "BTC_BCN",
	14:  "BTC_BTS",
	15:  "BTC_BURST",
	20:  "BTC_CLAM",
	25:  "BTC_DGB",
	27:  "BTC_DOGE",
	24:  "BTC_DASH",
	38:  "BTC_GAME",
	43:  "BTC_HUC",
	50:  "BTC_LTC",
	51:  "BTC_MAID",
	58:  "BTC_OMNI",
	61:  "BTC_NAV",
	64:  "BTC_NMC",
	69:  "BTC_NXT",
	75:  "BTC_PPC",
	89:  "BTC_STR",
	92:  "BTC_SYS",
	97:  "BTC_VIA",
	100: "BTC_VTC",
	108: "BTC_XCP",
	114: "BTC_XMR",
	116: "BTC_XPM",
	117: "BTC_XRP",
	112: "BTC_XEM",
	148: "BTC_ETH",
	150: "BTC_SC",
	153: "BTC_EXP",
	155: "BTC_FCT",
	160: "BTC_AMP",
	162: "BTC_DCR",
	163: "BTC_LSK",
	167: "BTC_LBC",
	168: "BTC_STEEM",
	170: "BTC_SBD",
	171: "BTC_ETC",
	174: "BTC_REP",
	177: "BTC_ARDR",
	178: "BTC_ZEC",
	182: "BTC_STRAT", // nolint: misspell
	184: "BTC_PASC",
	185: "BTC_GNT",
	187: "BTC_GNO",
	189: "BTC_BCH",
	192: "BTC_ZRX",
	194: "BTC_CVC",
	196: "BTC_OMG",
	198: "BTC_GAS",
	200: "BTC_STORJ",
	201: "BTC_EOS",
	204: "BTC_SNT",
	207: "BTC_KNC",
	210: "BTC_BAT",
	213: "BTC_LOOM",
	221: "BTC_QTUM",
	121: "USDT_BTC",
	216: "USDT_DOGE",
	122: "USDT_DASH",
	123: "USDT_LTC",
	124: "USDT_NXT",
	125: "USDT_STR",
	126: "USDT_XMR",
	127: "USDT_XRP",
	149: "USDT_ETH",
	219: "USDT_SC",
	218: "USDT_LSK",
	173: "USDT_ETC",
	175: "USDT_REP",
	180: "USDT_ZEC",
	217: "USDT_GNT",
	191: "USDT_BCH",
	220: "USDT_ZRX",
	203: "USDT_EOS",
	206: "USDT_SNT",
	209: "USDT_KNC",
	212: "USDT_BAT",
	215: "USDT_LOOM",
	223: "USDT_QTUM",
	129: "XMR_BCN",
	132: "XMR_DASH",
	137: "XMR_LTC",
	138: "XMR_MAID",
	140: "XMR_NXT",
	181: "XMR_ZEC",
	166: "ETH_LSK",
	169: "ETH_STEEM",
	172: "ETH_ETC",
	176: "ETH_REP",
	179: "ETH_ZEC",
	186: "ETH_GNT",
	188: "ETH_GNO",
	190: "ETH_BCH",
	193: "ETH_ZRX",
	195: "ETH_CVC",
	197: "ETH_OMG",
	199: "ETH_GAS",
	202: "ETH_EOS",
	205: "ETH_SNT",
	208: "ETH_KNC",
	211: "ETH_BAT",
	214: "ETH_LOOM",
	222: "ETH_QTUM",
	224: "USDC_BTC",
	226: "USDC_USDT",
	225: "USDC_ETH",
}
