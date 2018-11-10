package bitmex

import (
	"encoding/json"
	"errors"
	"time"
)

// ErrNotConnected is returned when the application read/writes
// a message and the connection is closed
var ErrNotConnected = errors.New("websocket: not connected")

type wsData struct {
	Table       string            `json:"table"`
	Action      string            `json:"action"`
	Keys        []string          `json:"keys"`
	Attributes  map[string]string `json:"attributes"`
	Types       map[string]string `json:"types"`
	ForeignKeys map[string]string `json:"foreignKeys"`
	Data        json.RawMessage
}

type wsSuccess struct {
	Success   bool              `json:"success"`
	Subscribe string            `json:"subscribe"`
	Request   map[string]string `json:"request"`
}

type wsInfo struct {
	Info      string    `json:"info"`
	Version   string    `json:"version"`
	Time      time.Time `json:"timestamp"`
	Docs      string    `json:"docs"`
	Heartbeat bool      `json:"heartbeatEnabled"`
}

//WSTrade - trade structure
type WSTrade struct {
	Size            float64   `json:"size"`
	Price           float64   `json:"price"`
	ForeignNotional float64   `json:"foreignNotional"`
	GrossValue      float64   `json:"grossValue"`
	HomeNotional    float64   `json:"homeNotional"`
	Symbol          string    `json:"symbol"`
	TickDirection   string    `json:"tickDirection"`
	Side            string    `json:"side"`
	TradeMatchID    string    `json:"trdMatchID"`
	Timestamp       time.Time `json:"timestamp"`
}

//WSQuote stores a quote structure
type WSQuote struct {
	AskPrice  float64 `json:"askPrice"`
	AskSize   int64   `json:"askSize"`
	BidPrice  float64 `json:"bidPrice"`
	BidSize   int64   `json:"bidSize"`
	Symbol    string  `json:"symbol"`
	Timestamp string  `json:"timestamp"`
}

type WSInstrument struct {
	AskPrice                       float64     `json:"askPrice"`
	BankruptLimitDownPrice         interface{} `json:"bankruptLimitDownPrice"`
	BankruptLimitUpPrice           interface{} `json:"bankruptLimitUpPrice"`
	BidPrice                       float64     `json:"bidPrice"`
	BuyLeg                         string      `json:"buyLeg"`
	CalcInterval                   interface{} `json:"calcInterval"`
	Capped                         bool        `json:"capped"`
	ClosingTimestamp               string      `json:"closingTimestamp"`
	Deleverage                     bool        `json:"deleverage"`
	Expiry                         interface{} `json:"expiry"`
	FairBasis                      float64     `json:"fairBasis"`
	FairBasisRate                  float64     `json:"fairBasisRate"`
	FairMethod                     string      `json:"fairMethod"`
	FairPrice                      float64     `json:"fairPrice"`
	Front                          string      `json:"front"`
	FundingBaseSymbol              string      `json:"fundingBaseSymbol"`
	FundingInterval                string      `json:"fundingInterval"`
	FundingPremiumSymbol           string      `json:"fundingPremiumSymbol"`
	FundingQuoteSymbol             string      `json:"fundingQuoteSymbol"`
	FundingRate                    float64     `json:"fundingRate"`
	FundingTimestamp               string      `json:"fundingTimestamp"`
	HasLiquidity                   bool        `json:"hasLiquidity"`
	HighPrice                      float64     `json:"highPrice"`
	ImpactAskPrice                 float64     `json:"impactAskPrice"`
	ImpactBidPrice                 float64     `json:"impactBidPrice"`
	ImpactMidPrice                 float64     `json:"impactMidPrice"`
	IndicativeFundingRate          float64     `json:"indicativeFundingRate"`
	IndicativeSettlePrice          float64     `json:"indicativeSettlePrice"`
	IndicativeTaxRate              int64       `json:"indicativeTaxRate"`
	InitMargin                     float64     `json:"initMargin"`
	InsuranceFee                   int64       `json:"insuranceFee"`
	InverseLeg                     string      `json:"inverseLeg"`
	IsInverse                      bool        `json:"isInverse"`
	IsQuanto                       bool        `json:"isQuanto"`
	LastChangePcnt                 float64     `json:"lastChangePcnt"`
	LastPrice                      float64     `json:"lastPrice"`
	LastPriceProtected             float64     `json:"lastPriceProtected"`
	LastTickDirection              string      `json:"lastTickDirection"`
	Limit                          interface{} `json:"limit"`
	LimitDownPrice                 interface{} `json:"limitDownPrice"`
	LimitUpPrice                   interface{} `json:"limitUpPrice"`
	Listing                        string      `json:"listing"`
	LotSize                        int64       `json:"lotSize"`
	LowPrice                       float64     `json:"lowPrice"`
	MaintMargin                    float64     `json:"maintMargin"`
	MakerFee                       float64     `json:"makerFee"`
	MarkMethod                     string      `json:"markMethod"`
	MarkPrice                      float64     `json:"markPrice"`
	MaxOrderQty                    float64     `json:"maxOrderQty"`
	MaxPrice                       float64     `json:"maxPrice"`
	MidPrice                       float64     `json:"midPrice"`
	Multiplier                     int64       `json:"multiplier"`
	OpenInterest                   int64       `json:"openInterest"`
	OpenValue                      int64       `json:"openValue"`
	OpeningTimestamp               string      `json:"openingTimestamp"`
	OptionMultiplier               interface{} `json:"optionMultiplier"`
	OptionStrikePcnt               interface{} `json:"optionStrikePcnt"`
	OptionStrikePrice              interface{} `json:"optionStrikePrice"`
	OptionStrikeRound              interface{} `json:"optionStrikeRound"`
	OptionUnderlyingPrice          interface{} `json:"optionUnderlyingPrice"`
	PositionCurrency               string      `json:"positionCurrency"`
	PrevClosePrice                 float64     `json:"prevClosePrice"`
	PrevPrice24h                   float64     `json:"prevPrice24h"`
	PrevTotalTurnover              int64       `json:"prevTotalTurnover"`
	PrevTotalVolume                int64       `json:"prevTotalVolume"`
	PublishInterval                interface{} `json:"publishInterval"`
	PublishTime                    interface{} `json:"publishTime"`
	QuoteCurrency                  string      `json:"quoteCurrency"`
	QuoteToSettleMultiplier        interface{} `json:"quoteToSettleMultiplier"`
	RebalanceInterval              interface{} `json:"rebalanceInterval"`
	RebalanceTimestamp             interface{} `json:"rebalanceTimestamp"`
	Reference                      string      `json:"reference"`
	ReferenceSymbol                string      `json:"referenceSymbol"`
	RelistInterval                 interface{} `json:"relistInterval"`
	RiskLimit                      int64       `json:"riskLimit"`
	RiskStep                       int64       `json:"riskStep"`
	RootSymbol                     string      `json:"rootSymbol"`
	SellLeg                        string      `json:"sellLeg"`
	SessionInterval                string      `json:"sessionInterval"`
	SettlCurrency                  string      `json:"settlCurrency"`
	Settle                         interface{} `json:"settle"`
	SettledPrice                   interface{} `json:"settledPrice"`
	SettlementFee                  int64       `json:"settlementFee"`
	State                          string      `json:"state"`
	Symbol                         string      `json:"symbol"`
	TakerFee                       float64     `json:"takerFee"`
	Taxed                          bool        `json:"taxed"`
	TickSize                       float64     `json:"tickSize"`
	Timestamp                      string      `json:"timestamp"`
	TotalTurnover                  int64       `json:"totalTurnover"`
	TotalVolume                    int64       `json:"totalVolume"`
	Turnover                       int64       `json:"turnover"`
	Turnover24h                    int64       `json:"turnover24h"`
	Typ                            string      `json:"typ"`
	Underlying                     string      `json:"underlying"`
	UnderlyingSymbol               string      `json:"underlyingSymbol"`
	UnderlyingToPositionMultiplier interface{} `json:"underlyingToPositionMultiplier"`
	UnderlyingToSettleMultiplier   int64       `json:"underlyingToSettleMultiplier"`
	Volume                         int64       `json:"volume"`
	Volume24h                      int64       `json:"volume24h"`
	Vwap                           float64     `json:"vwap"`
}

type wsError struct {
	Error string `json:"error"`
}

// ErrorCapture is a simple type for returned errors
type ErrorCapture struct {
	Message string `json:"message"`
}

// WebsocketSubscription takes in subscription information
type WebsocketSubscription struct {
	OP   string   `json:"op"`
	Args []string `json:"args"`
}

/*
{
	"table":"instrument",
	"keys":[
	   "symbol"
	],
	"types":{
	   "symbol":"symbol",
	   "rootSymbol":"symbol",
	   "state":"symbol",
	   "typ":"symbol",
	   "listing":"timestamp",
	   "front":"timestamp",
	   "expiry":"timestamp",
	   "settle":"timestamp",
	   "relistInterval":"timespan",
	   "inverseLeg":"symbol",
	   "sellLeg":"symbol",
	   "buyLeg":"symbol",
	   "optionStrikePcnt":"float",
	   "optionStrikeRound":"float",
	   "optionStrikePrice":"float",
	   "optionMultiplier":"float",
	   "positionCurrency":"symbol",
	   "underlying":"symbol",
	   "quoteCurrency":"symbol",
	   "underlyingSymbol":"symbol",
	   "reference":"symbol",
	   "referenceSymbol":"symbol",
	   "calcInterval":"timespan",
	   "publishInterval":"timespan",
	   "publishTime":"timespan",
	   "maxOrderQty":"long",
	   "maxPrice":"float",
	   "lotSize":"long",
	   "tickSize":"float",
	   "multiplier":"long",
	   "settlCurrency":"symbol",
	   "underlyingToPositionMultiplier":"long",
	   "underlyingToSettleMultiplier":"long",
	   "quoteToSettleMultiplier":"long",
	   "isQuanto":"boolean",
	   "isInverse":"boolean",
	   "initMargin":"float",
	   "maintMargin":"float",
	   "riskLimit":"long",
	   "riskStep":"long",
	   "limit":"float",
	   "capped":"boolean",
	   "taxed":"boolean",
	   "deleverage":"boolean",
	   "makerFee":"float",
	   "takerFee":"float",
	   "settlementFee":"float",
	   "insuranceFee":"float",
	   "fundingBaseSymbol":"symbol",
	   "fundingQuoteSymbol":"symbol",
	   "fundingPremiumSymbol":"symbol",
	   "fundingTimestamp":"timestamp",
	   "fundingInterval":"timespan",
	   "fundingRate":"float",
	   "indicativeFundingRate":"float",
	   "rebalanceTimestamp":"timestamp",
	   "rebalanceInterval":"timespan",
	   "openingTimestamp":"timestamp",
	   "closingTimestamp":"timestamp",
	   "sessionInterval":"timespan",
	   "prevClosePrice":"float",
	   "limitDownPrice":"float",
	   "limitUpPrice":"float",
	   "bankruptLimitDownPrice":"float",
	   "bankruptLimitUpPrice":"float",
	   "prevTotalVolume":"long",
	   "totalVolume":"long",
	   "volume":"long",
	   "volume24h":"long",
	   "prevTotalTurnover":"long",
	   "totalTurnover":"long",
	   "turnover":"long",
	   "turnover24h":"long",
	   "prevPrice24h":"float",
	   "vwap":"float",
	   "highPrice":"float",
	   "lowPrice":"float",
	   "lastPrice":"float",
	   "lastPriceProtected":"float",
	   "lastTickDirection":"symbol",
	   "lastChangePcnt":"float",
	   "bidPrice":"float",
	   "midPrice":"float",
	   "askPrice":"float",
	   "impactBidPrice":"float",
	   "impactMidPrice":"float",
	   "impactAskPrice":"float",
	   "hasLiquidity":"boolean",
	   "openInterest":"long",
	   "openValue":"long",
	   "fairMethod":"symbol",
	   "fairBasisRate":"float",
	   "fairBasis":"float",
	   "fairPrice":"float",
	   "markMethod":"symbol",
	   "markPrice":"float",
	   "indicativeTaxRate":"float",
	   "indicativeSettlePrice":"float",
	   "optionUnderlyingPrice":"float",
	   "settledPrice":"float",
	   "timestamp":"timestamp"
	},
	"foreignKeys":{
	   "inverseLeg":"instrument",
	   "sellLeg":"instrument",
	   "buyLeg":"instrument"
	},
	"attributes":{

	},
	"action":"partial",
	"data":[
	   {
		  "symbol":"XBTUSD",
		  "rootSymbol":"XBT",
		  "state":"Open",
		  "typ":"FFWCSX",
		  "listing":"2016-05-13T12:00:00.000Z",
		  "front":"2016-05-13T12:00:00.000Z",
		  "expiry":null,
		  "settle":null,
		  "relistInterval":null,
		  "inverseLeg":"",
		  "sellLeg":"",
		  "buyLeg":"",
		  "optionStrikePcnt":null,
		  "optionStrikeRound":null,
		  "optionStrikePrice":null,
		  "optionMultiplier":null,
		  "positionCurrency":"USD",
		  "underlying":"XBT",
		  "quoteCurrency":"USD",
		  "underlyingSymbol":"XBT=",
		  "reference":"BMEX",
		  "referenceSymbol":".BXBT",
		  "calcInterval":null,
		  "publishInterval":null,
		  "publishTime":null,
		  "maxOrderQty":10000000,
		  "maxPrice":1000000,
		  "lotSize":1,
		  "tickSize":0.5,
		  "multiplier":-100000000,
		  "settlCurrency":"XBt",
		  "underlyingToPositionMultiplier":null,
		  "underlyingToSettleMultiplier":-100000000,
		  "quoteToSettleMultiplier":null,
		  "isQuanto":false,
		  "isInverse":true,
		  "initMargin":0.01,
		  "maintMargin":0.005,
		  "riskLimit":20000000000,
		  "riskStep":10000000000,
		  "limit":null,
		  "capped":false,
		  "taxed":true,
		  "deleverage":true,
		  "makerFee":-0.00025,
		  "takerFee":0.00075,
		  "settlementFee":0,
		  "insuranceFee":0,
		  "fundingBaseSymbol":".XBTBON8H",
		  "fundingQuoteSymbol":".USDBON8H",
		  "fundingPremiumSymbol":".XBTUSDPI8H",
		  "fundingTimestamp":"2018-04-24T20:00:00.000Z",
		  "fundingInterval":"2000-01-01T08:00:00.000Z",
		  "fundingRate":-0.00007,
		  "indicativeFundingRate":0.000045,
		  "rebalanceTimestamp":null,
		  "rebalanceInterval":null,
		  "openingTimestamp":"2018-04-24T14:00:00.000Z",
		  "closingTimestamp":"2018-04-24T16:00:00.000Z",
		  "sessionInterval":"2000-01-01T02:00:00.000Z",
		  "prevClosePrice":9293.79,
		  "limitDownPrice":null,
		  "limitUpPrice":null,
		  "bankruptLimitDownPrice":null,
		  "bankruptLimitUpPrice":null,
		  "prevTotalVolume":403776486402,
		  "totalVolume":403983957030,
		  "volume":207470628,
		  "volume24h":3136547211,
		  "prevTotalTurnover":5349452421644829,
		  "totalTurnover":5351688081898598,
		  "turnover":2235660253769,
		  "turnover24h":34206869071222,
		  "prevPrice24h":8912.5,
		  "vwap":9170.1055,
		  "highPrice":9397,
		  "lowPrice":8783.5,
		  "lastPrice":9318,
		  "lastPriceProtected":9318,
		  "lastTickDirection":"PlusTick",
		  "lastChangePcnt":0.0455,
		  "bidPrice":9317.5,
		  "midPrice":9317.75,
		  "askPrice":9318,
		  "impactBidPrice":9317.9277,
		  "impactMidPrice":9318,
		  "impactAskPrice":9317.9277,
		  "hasLiquidity":false,
		  "openInterest":564581641,
		  "openValue":6056831844648,
		  "fairMethod":"FundingRate",
		  "fairBasisRate":-0.07665,
		  "fairBasis":-0.43,
		  "fairPrice":9321.36,
		  "markMethod":"FairPrice",
		  "markPrice":9321.36,
		  "indicativeTaxRate":0,
		  "indicativeSettlePrice":9321.79,
		  "optionUnderlyingPrice":null,
		  "settledPrice":null,
		  "timestamp":"2018-04-24T14:47:30.840Z"
	   }
	]
 }
*/
