// Code generated by protoc-gen-go. DO NOT EDIT.
// source: enum.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Venue int32

const (
	Venue_DARKPOOL    Venue = 0
	Venue_COINBASEPRO Venue = 1
	Venue_BITMEX      Venue = 2
	Venue_BITFINEX    Venue = 3
	Venue_BINANCE     Venue = 4
	Venue_GEMINI      Venue = 5
)

var Venue_name = map[int32]string{
	0: "DARKPOOL",
	1: "COINBASEPRO",
	2: "BITMEX",
	3: "BITFINEX",
	4: "BINANCE",
	5: "GEMINI",
}
var Venue_value = map[string]int32{
	"DARKPOOL":    0,
	"COINBASEPRO": 1,
	"BITMEX":      2,
	"BITFINEX":    3,
	"BINANCE":     4,
	"GEMINI":      5,
}

func (x Venue) String() string {
	return proto.EnumName(Venue_name, int32(x))
}
func (Venue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{0}
}

type AccountMode int32

const (
	AccountMode_LIVE AccountMode = 0
	AccountMode_DEMO AccountMode = 1
)

var AccountMode_name = map[int32]string{
	0: "LIVE",
	1: "DEMO",
}
var AccountMode_value = map[string]int32{
	"LIVE": 0,
	"DEMO": 1,
}

func (x AccountMode) String() string {
	return proto.EnumName(AccountMode_name, int32(x))
}
func (AccountMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{1}
}

type AccountType int32

const (
	AccountType_NET               AccountType = 0
	AccountType_HEDGE             AccountType = 1
	AccountType_CROSS_VENUE_NET   AccountType = 2
	AccountType_CROSS_VENUE_HEDGE AccountType = 3
)

var AccountType_name = map[int32]string{
	0: "NET",
	1: "HEDGE",
	2: "CROSS_VENUE_NET",
	3: "CROSS_VENUE_HEDGE",
}
var AccountType_value = map[string]int32{
	"NET":               0,
	"HEDGE":             1,
	"CROSS_VENUE_NET":   2,
	"CROSS_VENUE_HEDGE": 3,
}

func (x AccountType) String() string {
	return proto.EnumName(AccountType_name, int32(x))
}
func (AccountType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{2}
}

type VenueType int32

const (
	VenueType_SPOT    VenueType = 0
	VenueType_FUTURES VenueType = 1
)

var VenueType_name = map[int32]string{
	0: "SPOT",
	1: "FUTURES",
}
var VenueType_value = map[string]int32{
	"SPOT":    0,
	"FUTURES": 1,
}

func (x VenueType) String() string {
	return proto.EnumName(VenueType_name, int32(x))
}
func (VenueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{3}
}

type Symbol int32

const (
	Symbol_BTC Symbol = 0
	Symbol_ETH Symbol = 1
	Symbol_USD Symbol = 2
	Symbol_LTC Symbol = 3
	Symbol_BCH Symbol = 4
	Symbol_ETC Symbol = 5
)

var Symbol_name = map[int32]string{
	0: "BTC",
	1: "ETH",
	2: "USD",
	3: "LTC",
	4: "BCH",
	5: "ETC",
}
var Symbol_value = map[string]int32{
	"BTC": 0,
	"ETH": 1,
	"USD": 2,
	"LTC": 3,
	"BCH": 4,
	"ETC": 5,
}

func (x Symbol) String() string {
	return proto.EnumName(Symbol_name, int32(x))
}
func (Symbol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{4}
}

type Product int32

const (
	Product_BTC_USD Product = 0
	Product_ETH_BTC Product = 1
	Product_ETH_USD Product = 2
	Product_LTC_USD Product = 3
	Product_LTC_BTC Product = 4
	Product_LTC_ETH Product = 5
	Product_BCH_USD Product = 6
	Product_BCH_BTC Product = 7
	Product_BCH_ETH Product = 8
	Product_ETC_USD Product = 9
	Product_ETC_BTC Product = 10
	Product_ETC_ETH Product = 11
)

var Product_name = map[int32]string{
	0:  "BTC_USD",
	1:  "ETH_BTC",
	2:  "ETH_USD",
	3:  "LTC_USD",
	4:  "LTC_BTC",
	5:  "LTC_ETH",
	6:  "BCH_USD",
	7:  "BCH_BTC",
	8:  "BCH_ETH",
	9:  "ETC_USD",
	10: "ETC_BTC",
	11: "ETC_ETH",
}
var Product_value = map[string]int32{
	"BTC_USD": 0,
	"ETH_BTC": 1,
	"ETH_USD": 2,
	"LTC_USD": 3,
	"LTC_BTC": 4,
	"LTC_ETH": 5,
	"BCH_USD": 6,
	"BCH_BTC": 7,
	"BCH_ETH": 8,
	"ETC_USD": 9,
	"ETC_BTC": 10,
	"ETC_ETH": 11,
}

func (x Product) String() string {
	return proto.EnumName(Product_name, int32(x))
}
func (Product) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{5}
}

type OrderType int32

const (
	OrderType_BUY             OrderType = 0
	OrderType_SELL            OrderType = 1
	OrderType_BUY_LIMIT       OrderType = 2
	OrderType_SELL_LIMIT      OrderType = 3
	OrderType_BUY_STOP        OrderType = 4
	OrderType_SELL_STOP       OrderType = 5
	OrderType_BUY_STOP_LIMIT  OrderType = 6
	OrderType_SELL_STOP_LIMIT OrderType = 7
	OrderType_CLOSING_BY      OrderType = 8
)

var OrderType_name = map[int32]string{
	0: "BUY",
	1: "SELL",
	2: "BUY_LIMIT",
	3: "SELL_LIMIT",
	4: "BUY_STOP",
	5: "SELL_STOP",
	6: "BUY_STOP_LIMIT",
	7: "SELL_STOP_LIMIT",
	8: "CLOSING_BY",
}
var OrderType_value = map[string]int32{
	"BUY":             0,
	"SELL":            1,
	"BUY_LIMIT":       2,
	"SELL_LIMIT":      3,
	"BUY_STOP":        4,
	"SELL_STOP":       5,
	"BUY_STOP_LIMIT":  6,
	"SELL_STOP_LIMIT": 7,
	"CLOSING_BY":      8,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}
func (OrderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{6}
}

type OrderState int32

const (
	OrderState_STARTED        OrderState = 0
	OrderState_PLACED         OrderState = 1
	OrderState_CANCELED       OrderState = 2
	OrderState_PARTIAL        OrderState = 3
	OrderState_FILLED         OrderState = 4
	OrderState_REJECTED       OrderState = 5
	OrderState_EXPIRED        OrderState = 6
	OrderState_REQUEST_ADD    OrderState = 7
	OrderState_REQUEST_MODIFY OrderState = 8
	OrderState_REQUEST_CANCEL OrderState = 9
)

var OrderState_name = map[int32]string{
	0: "STARTED",
	1: "PLACED",
	2: "CANCELED",
	3: "PARTIAL",
	4: "FILLED",
	5: "REJECTED",
	6: "EXPIRED",
	7: "REQUEST_ADD",
	8: "REQUEST_MODIFY",
	9: "REQUEST_CANCEL",
}
var OrderState_value = map[string]int32{
	"STARTED":        0,
	"PLACED":         1,
	"CANCELED":       2,
	"PARTIAL":        3,
	"FILLED":         4,
	"REJECTED":       5,
	"EXPIRED":        6,
	"REQUEST_ADD":    7,
	"REQUEST_MODIFY": 8,
	"REQUEST_CANCEL": 9,
}

func (x OrderState) String() string {
	return proto.EnumName(OrderState_name, int32(x))
}
func (OrderState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{7}
}

type OrderTypeFilling int32

const (
	OrderTypeFilling_FOK    OrderTypeFilling = 0
	OrderTypeFilling_IOC    OrderTypeFilling = 1
	OrderTypeFilling_RETURN OrderTypeFilling = 2
)

var OrderTypeFilling_name = map[int32]string{
	0: "FOK",
	1: "IOC",
	2: "RETURN",
}
var OrderTypeFilling_value = map[string]int32{
	"FOK":    0,
	"IOC":    1,
	"RETURN": 2,
}

func (x OrderTypeFilling) String() string {
	return proto.EnumName(OrderTypeFilling_name, int32(x))
}
func (OrderTypeFilling) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{8}
}

type OrderReason int32

const (
	OrderReason_ORDER_CLIENT   OrderReason = 0
	OrderReason_ORDER_MOBILE   OrderReason = 1
	OrderReason_ORDER_WEB      OrderReason = 2
	OrderReason_ORDER_STRATEGY OrderReason = 3
	OrderReason_ORDER_SL       OrderReason = 4
	OrderReason_ORDER_TP       OrderReason = 5
	OrderReason_ORDER_SO       OrderReason = 6
)

var OrderReason_name = map[int32]string{
	0: "ORDER_CLIENT",
	1: "ORDER_MOBILE",
	2: "ORDER_WEB",
	3: "ORDER_STRATEGY",
	4: "ORDER_SL",
	5: "ORDER_TP",
	6: "ORDER_SO",
}
var OrderReason_value = map[string]int32{
	"ORDER_CLIENT":   0,
	"ORDER_MOBILE":   1,
	"ORDER_WEB":      2,
	"ORDER_STRATEGY": 3,
	"ORDER_SL":       4,
	"ORDER_TP":       5,
	"ORDER_SO":       6,
}

func (x OrderReason) String() string {
	return proto.EnumName(OrderReason_name, int32(x))
}
func (OrderReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{9}
}

type OrderTypeTime int32

const (
	OrderTypeTime_GTC           OrderTypeTime = 0
	OrderTypeTime_DAY           OrderTypeTime = 1
	OrderTypeTime_SPECIFIED     OrderTypeTime = 2
	OrderTypeTime_SPECIFIED_DAY OrderTypeTime = 3
)

var OrderTypeTime_name = map[int32]string{
	0: "GTC",
	1: "DAY",
	2: "SPECIFIED",
	3: "SPECIFIED_DAY",
}
var OrderTypeTime_value = map[string]int32{
	"GTC":           0,
	"DAY":           1,
	"SPECIFIED":     2,
	"SPECIFIED_DAY": 3,
}

func (x OrderTypeTime) String() string {
	return proto.EnumName(OrderTypeTime_name, int32(x))
}
func (OrderTypeTime) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{10}
}

type DealType int32

const (
	DealType_DEAL_BUY      DealType = 0
	DealType_DEAL_SELL     DealType = 1
	DealType_BALANCE       DealType = 2
	DealType_CREDIT        DealType = 3
	DealType_CHARGE        DealType = 4
	DealType_CORRECTION    DealType = 5
	DealType_BONUS         DealType = 6
	DealType_COMMISSION    DealType = 7
	DealType_INTEREST      DealType = 8
	DealType_BUY_CANCELED  DealType = 9
	DealType_SELL_CANCELED DealType = 10
	DealType_DIVIDEND      DealType = 11
)

var DealType_name = map[int32]string{
	0:  "DEAL_BUY",
	1:  "DEAL_SELL",
	2:  "BALANCE",
	3:  "CREDIT",
	4:  "CHARGE",
	5:  "CORRECTION",
	6:  "BONUS",
	7:  "COMMISSION",
	8:  "INTEREST",
	9:  "BUY_CANCELED",
	10: "SELL_CANCELED",
	11: "DIVIDEND",
}
var DealType_value = map[string]int32{
	"DEAL_BUY":      0,
	"DEAL_SELL":     1,
	"BALANCE":       2,
	"CREDIT":        3,
	"CHARGE":        4,
	"CORRECTION":    5,
	"BONUS":         6,
	"COMMISSION":    7,
	"INTEREST":      8,
	"BUY_CANCELED":  9,
	"SELL_CANCELED": 10,
	"DIVIDEND":      11,
}

func (x DealType) String() string {
	return proto.EnumName(DealType_name, int32(x))
}
func (DealType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{11}
}

type DealEntry int32

const (
	DealEntry_IN     DealEntry = 0
	DealEntry_OUT    DealEntry = 1
	DealEntry_INOUT  DealEntry = 2
	DealEntry_OUT_BY DealEntry = 3
)

var DealEntry_name = map[int32]string{
	0: "IN",
	1: "OUT",
	2: "INOUT",
	3: "OUT_BY",
}
var DealEntry_value = map[string]int32{
	"IN":     0,
	"OUT":    1,
	"INOUT":  2,
	"OUT_BY": 3,
}

func (x DealEntry) String() string {
	return proto.EnumName(DealEntry_name, int32(x))
}
func (DealEntry) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{12}
}

type DealReason int32

const (
	DealReason_DEAL_CLIENT   DealReason = 0
	DealReason_DEAL_MOBILE   DealReason = 1
	DealReason_DEAL_WEB      DealReason = 2
	DealReason_DEAL_STRATEGY DealReason = 3
	DealReason_DEAL_SL       DealReason = 4
	DealReason_DEAL_TP       DealReason = 5
	DealReason_DEAL_SO       DealReason = 6
	DealReason_ROLLOVER      DealReason = 7
	DealReason_VMARGIN       DealReason = 8
	DealReason_SPLIT         DealReason = 9
	DealReason_AJUST         DealReason = 10
)

var DealReason_name = map[int32]string{
	0:  "DEAL_CLIENT",
	1:  "DEAL_MOBILE",
	2:  "DEAL_WEB",
	3:  "DEAL_STRATEGY",
	4:  "DEAL_SL",
	5:  "DEAL_TP",
	6:  "DEAL_SO",
	7:  "ROLLOVER",
	8:  "VMARGIN",
	9:  "SPLIT",
	10: "AJUST",
}
var DealReason_value = map[string]int32{
	"DEAL_CLIENT":   0,
	"DEAL_MOBILE":   1,
	"DEAL_WEB":      2,
	"DEAL_STRATEGY": 3,
	"DEAL_SL":       4,
	"DEAL_TP":       5,
	"DEAL_SO":       6,
	"ROLLOVER":      7,
	"VMARGIN":       8,
	"SPLIT":         9,
	"AJUST":         10,
}

func (x DealReason) String() string {
	return proto.EnumName(DealReason_name, int32(x))
}
func (DealReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{13}
}

type PositionReason int32

const (
	PositionReason_POSITION_CLIENT   PositionReason = 0
	PositionReason_POSITION_MOBILE   PositionReason = 1
	PositionReason_POSITION_WEB      PositionReason = 2
	PositionReason_POSITION_STRATEGY PositionReason = 3
)

var PositionReason_name = map[int32]string{
	0: "POSITION_CLIENT",
	1: "POSITION_MOBILE",
	2: "POSITION_WEB",
	3: "POSITION_STRATEGY",
}
var PositionReason_value = map[string]int32{
	"POSITION_CLIENT":   0,
	"POSITION_MOBILE":   1,
	"POSITION_WEB":      2,
	"POSITION_STRATEGY": 3,
}

func (x PositionReason) String() string {
	return proto.EnumName(PositionReason_name, int32(x))
}
func (PositionReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{14}
}

type PositionType int32

const (
	PositionType_POSITION_BUY  PositionType = 0
	PositionType_POSITION_SELL PositionType = 1
)

var PositionType_name = map[int32]string{
	0: "POSITION_BUY",
	1: "POSITION_SELL",
}
var PositionType_value = map[string]int32{
	"POSITION_BUY":  0,
	"POSITION_SELL": 1,
}

func (x PositionType) String() string {
	return proto.EnumName(PositionType_name, int32(x))
}
func (PositionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{15}
}

type TradeAction int32

const (
	TradeAction_DEAL     TradeAction = 0
	TradeAction_PENDING  TradeAction = 1
	TradeAction_SLTP     TradeAction = 2
	TradeAction_MODIFY   TradeAction = 3
	TradeAction_REMOVE   TradeAction = 4
	TradeAction_CLOSE_BY TradeAction = 5
)

var TradeAction_name = map[int32]string{
	0: "DEAL",
	1: "PENDING",
	2: "SLTP",
	3: "MODIFY",
	4: "REMOVE",
	5: "CLOSE_BY",
}
var TradeAction_value = map[string]int32{
	"DEAL":     0,
	"PENDING":  1,
	"SLTP":     2,
	"MODIFY":   3,
	"REMOVE":   4,
	"CLOSE_BY": 5,
}

func (x TradeAction) String() string {
	return proto.EnumName(TradeAction_name, int32(x))
}
func (TradeAction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{16}
}

type Retcode int32

const (
	Retcode_REJECTX                 Retcode = 0
	Retcode_REJECT                  Retcode = 2
	Retcode_CANCEL                  Retcode = 3
	Retcode_PLACED_TRADE            Retcode = 4
	Retcode_DONE                    Retcode = 5
	Retcode_DONE_PARTIAL            Retcode = 6
	Retcode_ERROR                   Retcode = 7
	Retcode_TIMEOUT                 Retcode = 8
	Retcode_INVALID                 Retcode = 9
	Retcode_INVALID_VOLUME          Retcode = 10
	Retcode_INVALID_PRICE           Retcode = 11
	Retcode_INVALID_STOPS           Retcode = 12
	Retcode_TRADE_DISABLED          Retcode = 13
	Retcode_MARKET_CLOSED           Retcode = 14
	Retcode_NO_MONEY                Retcode = 15
	Retcode_PRICE_CHANGED           Retcode = 16
	Retcode_PRICE_OFF               Retcode = 17
	Retcode_INVALID_EXPIRATION      Retcode = 18
	Retcode_ORDER_CHANGED           Retcode = 19
	Retcode_TOO_MANY_REQUESTS       Retcode = 21
	Retcode_NO_CHANGES              Retcode = 22
	Retcode_SERVER_DISABLES_AT      Retcode = 23
	Retcode_CLIENT_DISABLES_AT      Retcode = 24
	Retcode_LOCKED                  Retcode = 25
	Retcode_FROZEN                  Retcode = 26
	Retcode_INVALID_FILL            Retcode = 27
	Retcode_CONNECTION              Retcode = 28
	Retcode_ONLY_REAL               Retcode = 29
	Retcode_LIMIT_ORDERS            Retcode = 30
	Retcode_LIMIT_VOLUME            Retcode = 31
	Retcode_INVALID_ORDER           Retcode = 32
	Retcode_POSITION_CLOSED         Retcode = 33
	Retcode_CLOSE_ORDER_EXIST       Retcode = 34
	Retcode_LIMIT_POSITIONS         Retcode = 35
	Retcode_REJECT_CANCEL           Retcode = 36
	Retcode_LONG_ONLY               Retcode = 37
	Retcode_SHORT_ONLY              Retcode = 38
	Retcode_CLOSE_ONLY              Retcode = 39
	Retcode_REQUEST_WITH_NO_TOKEN   Retcode = 40
	Retcode_INVALID_TOKEN           Retcode = 41
	Retcode_INVALID_REQUEST_CONTEXT Retcode = 42
	Retcode_INVALID_REQUEST         Retcode = 43
)

var Retcode_name = map[int32]string{
	0:  "REJECTX",
	2:  "REJECT",
	3:  "CANCEL",
	4:  "PLACED_TRADE",
	5:  "DONE",
	6:  "DONE_PARTIAL",
	7:  "ERROR",
	8:  "TIMEOUT",
	9:  "INVALID",
	10: "INVALID_VOLUME",
	11: "INVALID_PRICE",
	12: "INVALID_STOPS",
	13: "TRADE_DISABLED",
	14: "MARKET_CLOSED",
	15: "NO_MONEY",
	16: "PRICE_CHANGED",
	17: "PRICE_OFF",
	18: "INVALID_EXPIRATION",
	19: "ORDER_CHANGED",
	21: "TOO_MANY_REQUESTS",
	22: "NO_CHANGES",
	23: "SERVER_DISABLES_AT",
	24: "CLIENT_DISABLES_AT",
	25: "LOCKED",
	26: "FROZEN",
	27: "INVALID_FILL",
	28: "CONNECTION",
	29: "ONLY_REAL",
	30: "LIMIT_ORDERS",
	31: "LIMIT_VOLUME",
	32: "INVALID_ORDER",
	33: "POSITION_CLOSED",
	34: "CLOSE_ORDER_EXIST",
	35: "LIMIT_POSITIONS",
	36: "REJECT_CANCEL",
	37: "LONG_ONLY",
	38: "SHORT_ONLY",
	39: "CLOSE_ONLY",
	40: "REQUEST_WITH_NO_TOKEN",
	41: "INVALID_TOKEN",
	42: "INVALID_REQUEST_CONTEXT",
	43: "INVALID_REQUEST",
}
var Retcode_value = map[string]int32{
	"REJECTX":                 0,
	"REJECT":                  2,
	"CANCEL":                  3,
	"PLACED_TRADE":            4,
	"DONE":                    5,
	"DONE_PARTIAL":            6,
	"ERROR":                   7,
	"TIMEOUT":                 8,
	"INVALID":                 9,
	"INVALID_VOLUME":          10,
	"INVALID_PRICE":           11,
	"INVALID_STOPS":           12,
	"TRADE_DISABLED":          13,
	"MARKET_CLOSED":           14,
	"NO_MONEY":                15,
	"PRICE_CHANGED":           16,
	"PRICE_OFF":               17,
	"INVALID_EXPIRATION":      18,
	"ORDER_CHANGED":           19,
	"TOO_MANY_REQUESTS":       21,
	"NO_CHANGES":              22,
	"SERVER_DISABLES_AT":      23,
	"CLIENT_DISABLES_AT":      24,
	"LOCKED":                  25,
	"FROZEN":                  26,
	"INVALID_FILL":            27,
	"CONNECTION":              28,
	"ONLY_REAL":               29,
	"LIMIT_ORDERS":            30,
	"LIMIT_VOLUME":            31,
	"INVALID_ORDER":           32,
	"POSITION_CLOSED":         33,
	"CLOSE_ORDER_EXIST":       34,
	"LIMIT_POSITIONS":         35,
	"REJECT_CANCEL":           36,
	"LONG_ONLY":               37,
	"SHORT_ONLY":              38,
	"CLOSE_ONLY":              39,
	"REQUEST_WITH_NO_TOKEN":   40,
	"INVALID_TOKEN":           41,
	"INVALID_REQUEST_CONTEXT": 42,
	"INVALID_REQUEST":         43,
}

func (x Retcode) String() string {
	return proto.EnumName(Retcode_name, int32(x))
}
func (Retcode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{17}
}

type Granularity int32

const (
	Granularity_GRANULARITY_0M  Granularity = 0
	Granularity_GRANULARITY_1M  Granularity = 60
	Granularity_GRANULARITY_2M  Granularity = 120
	Granularity_GRANULARITY_3M  Granularity = 180
	Granularity_GRANULARITY_4M  Granularity = 240
	Granularity_GRANULARITY_5M  Granularity = 300
	Granularity_GRANULARITY_6M  Granularity = 360
	Granularity_GRANULARITY_7M  Granularity = 420
	Granularity_GRANULARITY_8M  Granularity = 480
	Granularity_GRANULARITY_9M  Granularity = 540
	Granularity_GRANULARITY_10M Granularity = 600
	Granularity_GRANULARITY_15M Granularity = 900
	Granularity_GRANULARITY_20M Granularity = 1200
	Granularity_GRANULARITY_30M Granularity = 1800
	Granularity_GRANULARITY_40M Granularity = 2400
	Granularity_GRANULARITY_50M Granularity = 3000
	Granularity_GRANULARITY_1H  Granularity = 3600
	Granularity_GRANULARITY_2H  Granularity = 7200
	Granularity_GRANULARITY_3H  Granularity = 10800
	Granularity_GRANULARITY_4H  Granularity = 14400
	Granularity_GRANULARITY_5H  Granularity = 18000
	Granularity_GRANULARITY_6H  Granularity = 21600
	Granularity_GRANULARITY_7H  Granularity = 25200
	Granularity_GRANULARITY_8H  Granularity = 28800
	Granularity_GRANULARITY_9H  Granularity = 32400
	Granularity_GRANULARITY_10H Granularity = 36000
	Granularity_GRANULARITY_11H Granularity = 39600
	Granularity_GRANULARITY_12H Granularity = 43200
	Granularity_GRANULARITY_1D  Granularity = 86400
)

var Granularity_name = map[int32]string{
	0:     "GRANULARITY_0M",
	60:    "GRANULARITY_1M",
	120:   "GRANULARITY_2M",
	180:   "GRANULARITY_3M",
	240:   "GRANULARITY_4M",
	300:   "GRANULARITY_5M",
	360:   "GRANULARITY_6M",
	420:   "GRANULARITY_7M",
	480:   "GRANULARITY_8M",
	540:   "GRANULARITY_9M",
	600:   "GRANULARITY_10M",
	900:   "GRANULARITY_15M",
	1200:  "GRANULARITY_20M",
	1800:  "GRANULARITY_30M",
	2400:  "GRANULARITY_40M",
	3000:  "GRANULARITY_50M",
	3600:  "GRANULARITY_1H",
	7200:  "GRANULARITY_2H",
	10800: "GRANULARITY_3H",
	14400: "GRANULARITY_4H",
	18000: "GRANULARITY_5H",
	21600: "GRANULARITY_6H",
	25200: "GRANULARITY_7H",
	28800: "GRANULARITY_8H",
	32400: "GRANULARITY_9H",
	36000: "GRANULARITY_10H",
	39600: "GRANULARITY_11H",
	43200: "GRANULARITY_12H",
	86400: "GRANULARITY_1D",
}
var Granularity_value = map[string]int32{
	"GRANULARITY_0M":  0,
	"GRANULARITY_1M":  60,
	"GRANULARITY_2M":  120,
	"GRANULARITY_3M":  180,
	"GRANULARITY_4M":  240,
	"GRANULARITY_5M":  300,
	"GRANULARITY_6M":  360,
	"GRANULARITY_7M":  420,
	"GRANULARITY_8M":  480,
	"GRANULARITY_9M":  540,
	"GRANULARITY_10M": 600,
	"GRANULARITY_15M": 900,
	"GRANULARITY_20M": 1200,
	"GRANULARITY_30M": 1800,
	"GRANULARITY_40M": 2400,
	"GRANULARITY_50M": 3000,
	"GRANULARITY_1H":  3600,
	"GRANULARITY_2H":  7200,
	"GRANULARITY_3H":  10800,
	"GRANULARITY_4H":  14400,
	"GRANULARITY_5H":  18000,
	"GRANULARITY_6H":  21600,
	"GRANULARITY_7H":  25200,
	"GRANULARITY_8H":  28800,
	"GRANULARITY_9H":  32400,
	"GRANULARITY_10H": 36000,
	"GRANULARITY_11H": 39600,
	"GRANULARITY_12H": 43200,
	"GRANULARITY_1D":  86400,
}

func (x Granularity) String() string {
	return proto.EnumName(Granularity_name, int32(x))
}
func (Granularity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{18}
}

type CandleGroupBy int32

const (
	CandleGroupBy_TIME   CandleGroupBy = 0
	CandleGroupBy_VOLUME CandleGroupBy = 2
	CandleGroupBy_TRADE  CandleGroupBy = 3
)

var CandleGroupBy_name = map[int32]string{
	0: "TIME",
	2: "VOLUME",
	3: "TRADE",
}
var CandleGroupBy_value = map[string]int32{
	"TIME":   0,
	"VOLUME": 2,
	"TRADE":  3,
}

func (x CandleGroupBy) String() string {
	return proto.EnumName(CandleGroupBy_name, int32(x))
}
func (CandleGroupBy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_enum_3eccc383bf0b1dd4, []int{19}
}

func init() {
	proto.RegisterEnum("api.Venue", Venue_name, Venue_value)
	proto.RegisterEnum("api.AccountMode", AccountMode_name, AccountMode_value)
	proto.RegisterEnum("api.AccountType", AccountType_name, AccountType_value)
	proto.RegisterEnum("api.VenueType", VenueType_name, VenueType_value)
	proto.RegisterEnum("api.Symbol", Symbol_name, Symbol_value)
	proto.RegisterEnum("api.Product", Product_name, Product_value)
	proto.RegisterEnum("api.OrderType", OrderType_name, OrderType_value)
	proto.RegisterEnum("api.OrderState", OrderState_name, OrderState_value)
	proto.RegisterEnum("api.OrderTypeFilling", OrderTypeFilling_name, OrderTypeFilling_value)
	proto.RegisterEnum("api.OrderReason", OrderReason_name, OrderReason_value)
	proto.RegisterEnum("api.OrderTypeTime", OrderTypeTime_name, OrderTypeTime_value)
	proto.RegisterEnum("api.DealType", DealType_name, DealType_value)
	proto.RegisterEnum("api.DealEntry", DealEntry_name, DealEntry_value)
	proto.RegisterEnum("api.DealReason", DealReason_name, DealReason_value)
	proto.RegisterEnum("api.PositionReason", PositionReason_name, PositionReason_value)
	proto.RegisterEnum("api.PositionType", PositionType_name, PositionType_value)
	proto.RegisterEnum("api.TradeAction", TradeAction_name, TradeAction_value)
	proto.RegisterEnum("api.Retcode", Retcode_name, Retcode_value)
	proto.RegisterEnum("api.Granularity", Granularity_name, Granularity_value)
	proto.RegisterEnum("api.CandleGroupBy", CandleGroupBy_name, CandleGroupBy_value)
}

func init() { proto.RegisterFile("enum.proto", fileDescriptor_enum_3eccc383bf0b1dd4) }

var fileDescriptor_enum_3eccc383bf0b1dd4 = []byte{
	// 1576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x57, 0xbd, 0x6f, 0x23, 0xc7,
	0x15, 0xd7, 0xf2, 0x9b, 0x8f, 0xfa, 0x78, 0xda, 0xb3, 0x7c, 0x71, 0xec, 0x7c, 0xd8, 0xf9, 0xde,
	0x00, 0x81, 0x4e, 0xf2, 0xd9, 0x67, 0x20, 0x08, 0xb0, 0xdc, 0x1d, 0x72, 0xc7, 0xda, 0x99, 0x61,
	0x66, 0x87, 0xb4, 0x98, 0x66, 0x21, 0x4b, 0x44, 0x20, 0x40, 0x47, 0x0a, 0x34, 0x05, 0x44, 0x4d,
	0xe0, 0x22, 0x45, 0x8a, 0x14, 0x06, 0x52, 0x05, 0x48, 0x71, 0x40, 0x12, 0x24, 0x45, 0x8a, 0x2b,
	0x12, 0xc0, 0xa5, 0x8b, 0xfc, 0x01, 0x49, 0x97, 0xf2, 0xd2, 0xa5, 0x74, 0x97, 0xea, 0x80, 0xe0,
	0xcd, 0xec, 0x52, 0xba, 0x65, 0x37, 0xef, 0x37, 0x6f, 0xde, 0xfc, 0xe6, 0xf7, 0x3e, 0x96, 0x04,
	0x98, 0xcd, 0x6f, 0x9e, 0xfe, 0xe8, 0x7a, 0xb9, 0x58, 0x2d, 0xfc, 0xfa, 0xd9, 0xf5, 0x65, 0x30,
	0x85, 0xe6, 0x64, 0x36, 0xbf, 0x99, 0xf9, 0xdb, 0xd0, 0x89, 0x43, 0x7d, 0x32, 0x52, 0x2a, 0xc5,
	0x2d, 0x7f, 0x0f, 0x7a, 0x91, 0xe2, 0xb2, 0x1f, 0x66, 0x6c, 0xa4, 0x15, 0x7a, 0x3e, 0x40, 0xab,
	0xcf, 0x8d, 0x60, 0xa7, 0x58, 0x23, 0xd7, 0x3e, 0x37, 0x03, 0x2e, 0xd9, 0x29, 0xd6, 0xfd, 0x1e,
	0xb4, 0xfb, 0x5c, 0x86, 0x32, 0x62, 0xd8, 0x20, 0xb7, 0x21, 0x13, 0x5c, 0x72, 0x6c, 0x06, 0x6f,
	0x43, 0x2f, 0x3c, 0x3f, 0x5f, 0xdc, 0xcc, 0x57, 0x62, 0x71, 0x31, 0xf3, 0x3b, 0xd0, 0x48, 0xf9,
	0x84, 0xe1, 0x16, 0xad, 0x62, 0x26, 0x14, 0x7a, 0x81, 0x58, 0xbb, 0x98, 0xdb, 0xeb, 0x99, 0xdf,
	0x86, 0xba, 0x64, 0x06, 0xb7, 0xfc, 0x2e, 0x34, 0x13, 0x16, 0x0f, 0x19, 0x7a, 0xfe, 0x03, 0xd8,
	0x8b, 0xb4, 0xca, 0xb2, 0x7c, 0xc2, 0xe4, 0x98, 0xe5, 0xb4, 0x5f, 0xf3, 0x0f, 0x60, 0xff, 0x3e,
	0xe8, 0x7c, 0xeb, 0xc1, 0x3b, 0xd0, 0xb5, 0x8f, 0xb1, 0xc1, 0x3a, 0xd0, 0xc8, 0x46, 0x8a, 0xa2,
	0xf5, 0xa0, 0x3d, 0x18, 0x9b, 0xb1, 0x66, 0x19, 0x7a, 0xc1, 0x4f, 0xa0, 0x95, 0xdd, 0x3e, 0xfd,
	0x78, 0x71, 0x45, 0xb7, 0xf5, 0x4d, 0x84, 0x5b, 0xb4, 0x60, 0x26, 0x41, 0x8f, 0x16, 0xe3, 0x2c,
	0xc6, 0x1a, 0x2d, 0x52, 0x13, 0x61, 0xdd, 0xfa, 0x44, 0x09, 0x36, 0x9c, 0x4f, 0x84, 0xcd, 0xe0,
	0x4f, 0x1e, 0xb4, 0x47, 0xcb, 0xc5, 0xc5, 0xcd, 0xf9, 0xca, 0x3e, 0xdd, 0x44, 0x39, 0x9d, 0xb1,
	0xb7, 0x30, 0x93, 0xe4, 0x14, 0xd2, 0x2b, 0x0d, 0x17, 0xad, 0x07, 0xed, 0xb4, 0x70, 0xab, 0x97,
	0x06, 0xb9, 0x35, 0x4a, 0x83, 0x6e, 0x6f, 0xda, 0x68, 0x91, 0x3b, 0xd3, 0x2a, 0x0d, 0x72, 0x6b,
	0x97, 0x06, 0xb9, 0x75, 0x5c, 0x68, 0x17, 0xad, 0x5b, 0x1a, 0xe4, 0x06, 0xa5, 0x41, 0x6e, 0xbd,
	0xe0, 0xb7, 0x1e, 0x74, 0xd5, 0xf2, 0x62, 0xb6, 0x2c, 0x95, 0xed, 0x8f, 0xa7, 0x4e, 0xfb, 0x8c,
	0xa5, 0x29, 0x7a, 0xfe, 0x0e, 0x74, 0xfb, 0xe3, 0x69, 0x9e, 0x72, 0xc1, 0x49, 0xd2, 0x5d, 0x00,
	0xda, 0x28, 0xec, 0xba, 0x4d, 0xf2, 0x78, 0x9a, 0x67, 0x46, 0x8d, 0xb0, 0x41, 0xce, 0x76, 0xd7,
	0x9a, 0x4d, 0xdf, 0x87, 0xdd, 0x72, 0xb3, 0x38, 0xd0, 0xa2, 0x44, 0xad, 0x5d, 0x0a, 0xb0, 0x4d,
	0x51, 0xa3, 0x54, 0x65, 0x5c, 0x0e, 0xf3, 0xfe, 0x14, 0x3b, 0xc1, 0x33, 0x0f, 0xc0, 0xb2, 0xca,
	0x56, 0x67, 0xab, 0x19, 0x31, 0xce, 0x4c, 0xa8, 0x0d, 0x23, 0x01, 0x01, 0x5a, 0xa3, 0x34, 0x8c,
	0x58, 0x8c, 0x1e, 0xdd, 0x1e, 0x51, 0x49, 0xa5, 0xac, 0x10, 0x70, 0x14, 0x6a, 0xc3, 0xc3, 0x14,
	0xeb, 0xe4, 0x36, 0xe0, 0x29, 0x6d, 0x34, 0xc8, 0x4d, 0xb3, 0x0f, 0x59, 0x44, 0x01, 0xac, 0x80,
	0xec, 0x74, 0xc4, 0x35, 0x23, 0x01, 0xf7, 0xa0, 0xa7, 0xd9, 0x4f, 0xc7, 0x2c, 0x33, 0x79, 0x18,
	0xc7, 0xd8, 0x26, 0xce, 0x25, 0x20, 0x54, 0xcc, 0x07, 0x53, 0xec, 0xdc, 0xc7, 0xdc, 0x75, 0xd8,
	0x0d, 0x0e, 0x01, 0xd7, 0xba, 0x0d, 0x2e, 0xaf, 0xae, 0x2e, 0xe7, 0x3f, 0x27, 0xf9, 0x06, 0xea,
	0xc4, 0x95, 0x0a, 0x57, 0x91, 0xeb, 0x07, 0xcd, 0xcc, 0x58, 0x4b, 0xac, 0x05, 0xbf, 0x84, 0x9e,
	0x3d, 0xa1, 0x67, 0x67, 0x9f, 0x2c, 0xe6, 0x3e, 0xc2, 0xb6, 0xd2, 0x31, 0xd3, 0x79, 0x94, 0x72,
	0x26, 0xa9, 0x00, 0xd7, 0x88, 0x50, 0x7d, 0x9e, 0x32, 0x27, 0xbe, 0x43, 0x3e, 0x62, 0x7d, 0xac,
	0x11, 0x0f, 0x67, 0x66, 0x46, 0x87, 0x86, 0x0d, 0xa7, 0x2e, 0x01, 0x05, 0x96, 0xba, 0x97, 0x3a,
	0xcb, 0x90, 0xfe, 0x77, 0x7b, 0x0a, 0x5b, 0x41, 0x04, 0x3b, 0x6b, 0xc6, 0xe6, 0xf2, 0xa9, 0xcd,
	0xf6, 0xb0, 0xac, 0xec, 0x38, 0x9c, 0xba, 0xfb, 0xb2, 0x11, 0x8b, 0xf8, 0x80, 0x5b, 0x41, 0xf7,
	0x61, 0x67, 0x6d, 0xe6, 0xe4, 0x51, 0x0f, 0x3e, 0xf7, 0xa0, 0x13, 0xcf, 0xce, 0xae, 0x6c, 0xb9,
	0xd0, 0x30, 0x60, 0x61, 0x9a, 0xbb, 0x9a, 0xd9, 0x81, 0xae, 0xb5, 0x8a, 0xc2, 0xa1, 0x6a, 0x0c,
	0x53, 0xdb, 0xf0, 0x35, 0xd2, 0x21, 0xd2, 0x2c, 0xb6, 0x25, 0x43, 0xeb, 0x24, 0xd4, 0x43, 0x1a,
	0x04, 0x94, 0x78, 0xa5, 0x35, 0x8b, 0x0c, 0x57, 0x12, 0x9b, 0xd4, 0xd1, 0x7d, 0x25, 0xc7, 0x19,
	0xb6, 0xdc, 0x96, 0x10, 0x3c, 0xcb, 0x68, 0xab, 0x4d, 0x97, 0x71, 0x69, 0x98, 0x66, 0x99, 0xc1,
	0x0e, 0x69, 0x45, 0xa5, 0xb5, 0xce, 0x7e, 0xd7, 0x92, 0xa5, 0xc2, 0x5a, 0x43, 0x60, 0xf9, 0xf1,
	0x09, 0x8f, 0x99, 0x8c, 0xb1, 0x17, 0x1c, 0x43, 0x97, 0x98, 0xb3, 0xf9, 0x6a, 0x79, 0xeb, 0xb7,
	0xa0, 0xc6, 0xa5, 0x7b, 0xba, 0x1a, 0x1b, 0xf4, 0xe8, 0x66, 0x2e, 0x69, 0x69, 0xc9, 0xaa, 0xb1,
	0xa1, 0x4a, 0xac, 0x07, 0x7f, 0xf6, 0x00, 0xe8, 0x54, 0x91, 0xb4, 0x3d, 0xe8, 0xd9, 0x37, 0xae,
	0x73, 0x56, 0x02, 0xeb, 0x94, 0x95, 0x9a, 0xb8, 0x8c, 0xed, 0xc3, 0x8e, 0xd3, 0xe4, 0x2e, 0x61,
	0x3d, 0x68, 0x3b, 0x28, 0x75, 0x9d, 0x6d, 0x0d, 0x9b, 0xae, 0xf5, 0x8e, 0xc2, 0x96, 0xad, 0x59,
	0x95, 0xa6, 0x6a, 0xc2, 0xb4, 0x6b, 0xed, 0x89, 0x08, 0xf5, 0x90, 0x4b, 0xec, 0x10, 0xd5, 0x6c,
	0x94, 0x72, 0x83, 0x5d, 0x5a, 0x86, 0x1f, 0x8e, 0x33, 0x83, 0x10, 0x9c, 0xc3, 0xee, 0x68, 0xf1,
	0xc9, 0xe5, 0xea, 0x72, 0x31, 0x2f, 0xc8, 0x3e, 0x80, 0xbd, 0x91, 0xca, 0x38, 0x49, 0x7b, 0x47,
	0xf8, 0x3e, 0xb8, 0x26, 0x8d, 0xb0, 0xbd, 0x06, 0x1d, 0xf1, 0x03, 0xd8, 0x5f, 0x23, 0x77, 0xe4,
	0x83, 0x63, 0xd8, 0x2e, 0x2f, 0xb1, 0x15, 0x70, 0xff, 0xa0, 0xab, 0x82, 0x7d, 0xd8, 0xb9, 0x3b,
	0x68, 0x2b, 0x21, 0x30, 0xd0, 0x33, 0xcb, 0xb3, 0x8b, 0x59, 0x78, 0x4e, 0xe7, 0xdc, 0x5c, 0x0f,
	0x53, 0x37, 0x0b, 0x47, 0x4c, 0xc6, 0x5c, 0x0e, 0xd1, 0xb3, 0x23, 0x27, 0x35, 0x23, 0xa7, 0x7f,
	0xd1, 0x7a, 0x75, 0xd7, 0x40, 0x42, 0x4d, 0x98, 0x2b, 0x6e, 0x9a, 0x12, 0x8c, 0x32, 0xd3, 0x0c,
	0xfe, 0xd7, 0x84, 0xb6, 0x9e, 0xad, 0xce, 0xe9, 0xa3, 0xd1, 0x83, 0xb6, 0x6b, 0xf0, 0x53, 0x37,
	0x20, 0x9c, 0x51, 0xd4, 0x9d, 0xeb, 0xd8, 0xba, 0xe5, 0x6a, 0x07, 0x47, 0x6e, 0x74, 0x18, 0x53,
	0x40, 0x62, 0xa2, 0x24, 0xc3, 0x26, 0xed, 0xd1, 0x2a, 0x2f, 0xe7, 0x47, 0x8b, 0x94, 0x65, 0x5a,
	0xab, 0x42, 0x7c, 0xc3, 0x05, 0xa3, 0xe2, 0xb0, 0x73, 0x95, 0xcb, 0x49, 0x98, 0x72, 0xaa, 0x39,
	0x1f, 0x76, 0x0b, 0x23, 0x9f, 0xa8, 0x74, 0x2c, 0x18, 0x02, 0x09, 0x50, 0x62, 0x23, 0xcd, 0x23,
	0x86, 0xbd, 0xfb, 0x10, 0x8d, 0xbd, 0x0c, 0xb7, 0xe9, 0xa4, 0x65, 0x91, 0xc7, 0x3c, 0x0b, 0xfb,
	0x54, 0xae, 0x3b, 0xe4, 0x26, 0x42, 0x7d, 0xc2, 0x4c, 0x6e, 0x9f, 0x19, 0xe3, 0x2e, 0x3d, 0x59,
	0xaa, 0x5c, 0x28, 0xc9, 0xa6, 0xb8, 0x67, 0xb5, 0xa5, 0x90, 0x79, 0x94, 0x84, 0x72, 0xc8, 0x62,
	0x44, 0x6a, 0x3a, 0x07, 0xa9, 0xc1, 0x00, 0xf7, 0xfd, 0xd7, 0xc1, 0x2f, 0x6f, 0xb2, 0x33, 0x2e,
	0xb4, 0x7d, 0xe5, 0xd3, 0xc9, 0x62, 0xd8, 0x14, 0x27, 0x1f, 0x50, 0x86, 0x8d, 0x52, 0xb9, 0x08,
	0xe5, 0x34, 0x2f, 0xa6, 0x5b, 0x86, 0x07, 0xd4, 0x76, 0x52, 0x15, 0x6e, 0x19, 0xbe, 0x4e, 0x11,
	0x33, 0xa6, 0x27, 0x4c, 0x97, 0x4c, 0xb3, 0x3c, 0x34, 0xf8, 0x90, 0x70, 0x57, 0x53, 0xaf, 0xe0,
	0x5f, 0x21, 0xc5, 0x53, 0x15, 0x9d, 0xb0, 0x18, 0xdf, 0xb0, 0x33, 0x58, 0xab, 0x9f, 0x31, 0x89,
	0x5f, 0x25, 0x85, 0x4b, 0x66, 0x34, 0x97, 0xf1, 0x4d, 0xd7, 0xe0, 0x52, 0x16, 0xbd, 0xff, 0x96,
	0x1d, 0x76, 0x32, 0x25, 0x32, 0x61, 0x8a, 0x5f, 0xa3, 0x03, 0xf6, 0xf3, 0x90, 0x5b, 0xe2, 0x19,
	0x7e, 0xfd, 0x0e, 0x29, 0xb4, 0xfe, 0xc6, 0x7d, 0x61, 0xad, 0x17, 0x7e, 0xb3, 0x52, 0xf4, 0x56,
	0xc6, 0xb7, 0xed, 0x0f, 0x01, 0x5b, 0x39, 0x4e, 0x04, 0x76, 0xca, 0x33, 0x83, 0xef, 0x90, 0xaf,
	0x0b, 0x58, 0x9e, 0xc8, 0xf0, 0x5b, 0x14, 0xd3, 0x95, 0x4f, 0x39, 0xeb, 0xbf, 0x4d, 0xcc, 0x52,
	0x25, 0x87, 0x39, 0xd1, 0xc3, 0xef, 0xd8, 0x6f, 0x60, 0xa2, 0xb4, 0x71, 0xf6, 0x77, 0xcb, 0xaf,
	0x17, 0x73, 0xf6, 0xf7, 0xfc, 0x37, 0xe0, 0xa0, 0xfc, 0x5c, 0x7c, 0xc4, 0x4d, 0x92, 0x4b, 0x95,
	0x1b, 0x75, 0xc2, 0x24, 0x7e, 0xff, 0x3e, 0x61, 0x07, 0xfd, 0xc0, 0x7f, 0x13, 0x1e, 0x96, 0xd0,
	0xfa, 0x23, 0xa3, 0xa4, 0x61, 0xa7, 0x06, 0x03, 0x62, 0x58, 0xd9, 0xc4, 0x1f, 0x06, 0x2f, 0x1b,
	0xd0, 0x1b, 0x2e, 0xcf, 0xe6, 0x37, 0x57, 0x67, 0xcb, 0xcb, 0xd5, 0x2d, 0xd5, 0xd2, 0x50, 0x87,
	0x72, 0x9c, 0x86, 0x9a, 0x9b, 0x69, 0x7e, 0x28, 0x70, 0xab, 0x8a, 0x3d, 0x12, 0xf8, 0xe3, 0x2a,
	0x76, 0x24, 0xf0, 0x17, 0xfe, 0x83, 0x57, 0xb1, 0x63, 0x81, 0x7f, 0xf3, 0xaa, 0xe0, 0xbb, 0x02,
	0xbf, 0xdc, 0x00, 0x1f, 0x0b, 0xfc, 0x6b, 0xad, 0x0a, 0xbe, 0x27, 0xf0, 0xbf, 0x1b, 0xe0, 0xfb,
	0x02, 0xff, 0x58, 0xaf, 0x82, 0x4f, 0x04, 0xbe, 0xd8, 0x00, 0x3f, 0x10, 0xf8, 0xfb, 0x86, 0xff,
	0x1a, 0xec, 0xbd, 0x42, 0xfd, 0x50, 0xe0, 0xbf, 0x37, 0xd1, 0xc7, 0x02, 0x7f, 0xd5, 0xae, 0xa2,
	0x47, 0x87, 0x02, 0x9f, 0x77, 0xab, 0xe8, 0xf1, 0xa1, 0xc0, 0x5f, 0xef, 0x56, 0xd1, 0x77, 0x0f,
	0x05, 0xbe, 0xf0, 0xab, 0xe8, 0xe3, 0x43, 0x81, 0x9f, 0x3f, 0xac, 0x12, 0x7b, 0x94, 0xe0, 0x67,
	0x6f, 0x55, 0xc1, 0xa3, 0x04, 0x9f, 0x3d, 0xd9, 0x10, 0x30, 0xc1, 0xe7, 0x66, 0x43, 0xc0, 0x04,
	0xbf, 0xb8, 0xf6, 0x5f, 0xab, 0x08, 0x98, 0xe0, 0x3f, 0x7f, 0xe3, 0x55, 0xd1, 0xf7, 0x12, 0x7c,
	0xf1, 0x97, 0x0d, 0xf4, 0xfd, 0x04, 0xbf, 0xfc, 0xc7, 0x06, 0xfa, 0x24, 0xc1, 0x4f, 0xff, 0xb3,
	0x81, 0x7e, 0x90, 0xe0, 0x67, 0x2f, 0x3d, 0xff, 0xa0, 0xaa, 0x62, 0x82, 0xcf, 0x7e, 0x57, 0xdb,
	0x80, 0x1f, 0x25, 0xf8, 0xfc, 0xef, 0x9b, 0xf0, 0x51, 0x82, 0x5f, 0xfc, 0xab, 0x56, 0x0d, 0xfd,
	0x28, 0xc6, 0x4f, 0xff, 0xd0, 0x0c, 0x0e, 0x61, 0x27, 0x3a, 0x9b, 0x5f, 0x5c, 0xcd, 0x86, 0xcb,
	0xc5, 0xcd, 0x75, 0xff, 0x96, 0xe6, 0x28, 0x0d, 0x48, 0x37, 0x7b, 0x8b, 0xe6, 0xac, 0xd1, 0x04,
	0x75, 0x83, 0xb6, 0xfe, 0x71, 0xcb, 0xfe, 0x95, 0x38, 0xfe, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xa3, 0x91, 0x04, 0x7d, 0x58, 0x0c, 0x00, 0x00,
}
