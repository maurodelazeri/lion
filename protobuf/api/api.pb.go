// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdb, 0x6e, 0x13, 0x31,
	0x10, 0xa5, 0xaa, 0x00, 0xc5, 0xbd, 0xa5, 0xee, 0x3d, 0xd0, 0x72, 0x79, 0xe2, 0xa9, 0x8a, 0x0a,
	0x4f, 0x08, 0x21, 0x35, 0x5c, 0x02, 0x2a, 0x15, 0x11, 0x09, 0x94, 0x47, 0x9c, 0xdd, 0x51, 0xb1,
	0x92, 0xac, 0x17, 0xdb, 0x8b, 0xc8, 0xbf, 0xf2, 0x31, 0x68, 0xbd, 0xe3, 0x59, 0xaf, 0x93, 0xf0,
	0x94, 0xcc, 0x39, 0x33, 0xe7, 0x8c, 0xed, 0x59, 0x9b, 0xb5, 0x44, 0x2e, 0xcf, 0x73, 0xad, 0xac,
	0xe2, 0xeb, 0x22, 0x97, 0x1d, 0x26, 0x0a, 0xfb, 0xb3, 0x02, 0x3a, 0x9b, 0xbf, 0x21, 0x2b, 0xc0,
	0x60, 0xb4, 0x03, 0x7f, 0x20, 0x29, 0xac, 0x54, 0x19, 0x02, 0xbb, 0x63, 0x91, 0x4c, 0x2c, 0x18,
	0x2b, 0xb3, 0x5b, 0x84, 0xb6, 0x34, 0xe4, 0x4a, 0x5b, 0x5f, 0xb2, 0x2d, 0x92, 0x44, 0x15, 0x19,
	0xc5, 0xed, 0xa4, 0xd0, 0x1a, 0xb2, 0x44, 0x92, 0xe8, 0x76, 0xae, 0x55, 0x5a, 0x24, 0x75, 0xc5,
	0x58, 0x4c, 0x45, 0x96, 0x10, 0xdf, 0x9e, 0x09, 0x3d, 0x01, 0x9b, 0x0a, 0x2b, 0x7c, 0x53, 0x15,
	0xe2, 0x79, 0x63, 0x85, 0x95, 0xc6, 0xca, 0x04, 0x2b, 0x2e, 0xfe, 0x6e, 0xb1, 0xf5, 0xcb, 0xc1,
	0x47, 0xde, 0x65, 0x77, 0x3f, 0xa9, 0x5b, 0x99, 0xf1, 0xdd, 0xf3, 0x72, 0x89, 0xee, 0xff, 0x17,
	0xf8, 0x55, 0x80, 0xb1, 0x1d, 0x1e, 0x42, 0x26, 0x57, 0x99, 0x81, 0xa7, 0x77, 0xf8, 0x6b, 0xb6,
	0xd5, 0x07, 0xfb, 0x86, 0x5a, 0xe4, 0xfb, 0x2e, 0x0d, 0x81, 0xb9, 0x2f, 0x3e, 0x88, 0x50, 0xaa,
	0x7f, 0xc7, 0x36, 0x07, 0xca, 0x78, 0x81, 0x39, 0x3f, 0x6e, 0x24, 0x96, 0x94, 0x97, 0x38, 0x59,
	0xc2, 0x90, 0xcc, 0x0b, 0xd6, 0xea, 0x83, 0xfd, 0xe6, 0xb6, 0x1e, 0x9b, 0x77, 0x41, 0xb3, 0x79,
	0x84, 0xa8, 0xea, 0x15, 0x6b, 0x95, 0x3a, 0x0e, 0xe6, 0x07, 0x75, 0x4a, 0x68, 0x7b, 0x18, 0xc3,
	0x54, 0xfd, 0x92, 0x6d, 0xf4, 0xc1, 0x0e, 0xf0, 0x2c, 0xf8, 0x9e, 0x4b, 0xc4, 0xd0, 0x57, 0xef,
	0x37, 0x41, 0xaa, 0xed, 0xb1, 0x8d, 0x52, 0x0d, 0x09, 0x7e, 0x14, 0xa6, 0x85, 0xee, 0xc7, 0x8b,
	0x44, 0xe4, 0x7f, 0x89, 0xd3, 0x82, 0xfe, 0x18, 0x36, 0xfd, 0x09, 0x8c, 0xfd, 0x91, 0x40, 0x7f,
	0x8c, 0x16, 0xfd, 0x1b, 0x44, 0xe4, 0xdf, 0xc3, 0xd9, 0x43, 0x7f, 0x0c, 0x9b, 0xfe, 0x04, 0xc6,
	0xfe, 0x48, 0xa0, 0x3f, 0x46, 0x8b, 0xfe, 0x0d, 0x82, 0x34, 0x46, 0x6c, 0x2f, 0x58, 0x3f, 0xf5,
	0xf1, 0x30, 0x6c, 0x99, 0x60, 0x2f, 0x78, 0xba, 0x82, 0x25, 0xd5, 0xef, 0x6c, 0x2f, 0xd8, 0x19,
	0x9f, 0xc1, 0xcf, 0xc2, 0xba, 0x25, 0x8d, 0x3e, 0x5a, 0xc9, 0x47, 0x33, 0x3a, 0xd2, 0x22, 0xa5,
	0x19, 0x75, 0x41, 0x73, 0x46, 0x11, 0xa2, 0xaa, 0x1b, 0xb6, 0xdf, 0x07, 0x7b, 0xed, 0xbe, 0xdf,
	0x2b, 0x99, 0xa5, 0xc3, 0x62, 0x36, 0x13, 0x7a, 0xce, 0xab, 0x85, 0x2c, 0xe0, 0x5e, 0xec, 0x6c,
	0x15, 0x4d, 0xc2, 0x3f, 0xd8, 0x11, 0x09, 0xe3, 0x80, 0x79, 0xed, 0xc7, 0x41, 0x71, 0x93, 0xf2,
	0xf2, 0x4f, 0xfe, 0x93, 0x41, 0x0e, 0x57, 0xac, 0x4d, 0x0b, 0x7e, 0x0b, 0x22, 0x9d, 0x89, 0x8c,
	0x9f, 0xd4, 0x8b, 0xf4, 0x98, 0xd7, 0xec, 0x2c, 0xa3, 0x48, 0xec, 0xab, 0xdb, 0x87, 0xcf, 0x3a,
	0x05, 0x3d, 0x56, 0x6a, 0x42, 0x82, 0xd5, 0x71, 0x13, 0x1e, 0x69, 0x9e, 0xae, 0x60, 0x49, 0xf6,
	0x9c, 0xdd, 0x77, 0x43, 0xac, 0x0d, 0xdf, 0xc1, 0x59, 0xd3, 0xbe, 0xb8, 0x5d, 0x03, 0x94, 0x7f,
	0xcd, 0x0e, 0xeb, 0x8f, 0xfe, 0x06, 0x60, 0x32, 0x04, 0x98, 0x96, 0xbf, 0x7e, 0x86, 0x8b, 0xf9,
	0x10, 0xa6, 0x0e, 0x89, 0x66, 0x38, 0x24, 0x48, 0xee, 0x82, 0xb5, 0x5c, 0x73, 0x43, 0xc8, 0x52,
	0x9c, 0x09, 0x17, 0x37, 0x5b, 0x40, 0xc8, 0x14, 0x53, 0xeb, 0x6e, 0xad, 0xcd, 0xb2, 0x05, 0x65,
	0x64, 0xf9, 0xae, 0xf8, 0x1b, 0xd7, 0xc7, 0xd1, 0xcd, 0xe3, 0xb3, 0xa8, 0xba, 0xfa, 0x6a, 0x3e,
	0x48, 0x63, 0x95, 0x95, 0x89, 0x98, 0xe2, 0x3c, 0x56, 0xdb, 0xe8, 0x60, 0x5d, 0xc3, 0x5e, 0xec,
	0xc1, 0x0a, 0xb6, 0xd2, 0xec, 0xae, 0xa1, 0x6a, 0xaf, 0x7e, 0xdb, 0x46, 0x6a, 0x02, 0xfe, 0x70,
	0x62, 0xb8, 0x79, 0x38, 0x8b, 0x2c, 0xed, 0xce, 0x7b, 0xb6, 0x11, 0xb0, 0x74, 0x4b, 0x10, 0x12,
	0xdf, 0x12, 0x01, 0xe1, 0x35, 0x9e, 0xad, 0x75, 0xd7, 0xc6, 0xf7, 0xdc, 0x2b, 0xf7, 0xfc, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0xe7, 0x52, 0xf3, 0xb8, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetCurrencies(ctx context.Context, in *CurrencyRequest, opts ...grpc.CallOption) (*CurrencyResponse, error)
	PostCurrency(ctx context.Context, in *CurrencyPostRequest, opts ...grpc.CallOption) (*CurrencyPostResponse, error)
	GetVenues(ctx context.Context, in *VenueRequest, opts ...grpc.CallOption) (*VenueResponse, error)
	PostVenue(ctx context.Context, in *VenuePostRequest, opts ...grpc.CallOption) (*VenuePostResponse, error)
	GetProducts(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	PostProduct(ctx context.Context, in *ProductPostRequest, opts ...grpc.CallOption) (*ProductPostResponse, error)
	GetAccounts(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	PostAccount(ctx context.Context, in *AccountPostRequest, opts ...grpc.CallOption) (*AccountPostResponse, error)
	GetBalances(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	PostBalance(ctx context.Context, in *BalancePostRequest, opts ...grpc.CallOption) (*BalancePostResponse, error)
	GetAccountsBalances(ctx context.Context, in *AccountsBalancesRequest, opts ...grpc.CallOption) (*AccountsBalancesResponse, error)
	PostAccountsBalance(ctx context.Context, in *AccountBalancePostRequest, opts ...grpc.CallOption) (*AccountBalancePostResponse, error)
	GetTrades(ctx context.Context, in *TradeRequest, opts ...grpc.CallOption) (*TradeResponse, error)
	GetMarketKindSummary(ctx context.Context, in *MarketKindSummaryRequest, opts ...grpc.CallOption) (*MarketKindSummaryResponse, error)
	GetMarketProductSummary(ctx context.Context, in *MarketProductSummaryRequest, opts ...grpc.CallOption) (*MarketProductSummaryResponse, error)
	GetTradesDeadman(ctx context.Context, in *TradesDeadmanRequest, opts ...grpc.CallOption) (*TradesDeadmanResponse, error)
	GetOrderbooksDeadman(ctx context.Context, in *OrderbookDeadmanRequest, opts ...grpc.CallOption) (*OrderbookDeadmanResponse, error)
	GetBars(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*BarResponse, error)
	GetProductWeekSeelWeek(ctx context.Context, in *BuySellWeekRequest, opts ...grpc.CallOption) (*BuySellWeekResponse, error)
	// Get all Order for a user with filter - A server-to-client streaming RPC.
	// rpc GetOrders(OrderFilter) returns (stream Order) {}
	// Order request - A simple RPC
	OrderSend(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResult, error)
	GetPositions(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionsResult, error)
	GetHistoticalTrades(ctx context.Context, in *HistoricalTradesRequest, opts ...grpc.CallOption) (API_GetHistoticalTradesClient, error)
	GetBacktestingToken(ctx context.Context, in *BacktestingTokenRequest, opts ...grpc.CallOption) (*BacktestingTokenResponse, error)
	Backtesting(ctx context.Context, opts ...grpc.CallOption) (API_BacktestingClient, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/api.API/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetCurrencies(ctx context.Context, in *CurrencyRequest, opts ...grpc.CallOption) (*CurrencyResponse, error) {
	out := new(CurrencyResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetCurrencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) PostCurrency(ctx context.Context, in *CurrencyPostRequest, opts ...grpc.CallOption) (*CurrencyPostResponse, error) {
	out := new(CurrencyPostResponse)
	err := c.cc.Invoke(ctx, "/api.API/PostCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetVenues(ctx context.Context, in *VenueRequest, opts ...grpc.CallOption) (*VenueResponse, error) {
	out := new(VenueResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetVenues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) PostVenue(ctx context.Context, in *VenuePostRequest, opts ...grpc.CallOption) (*VenuePostResponse, error) {
	out := new(VenuePostResponse)
	err := c.cc.Invoke(ctx, "/api.API/PostVenue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetProducts(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) PostProduct(ctx context.Context, in *ProductPostRequest, opts ...grpc.CallOption) (*ProductPostResponse, error) {
	out := new(ProductPostResponse)
	err := c.cc.Invoke(ctx, "/api.API/PostProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetAccounts(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) PostAccount(ctx context.Context, in *AccountPostRequest, opts ...grpc.CallOption) (*AccountPostResponse, error) {
	out := new(AccountPostResponse)
	err := c.cc.Invoke(ctx, "/api.API/PostAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetBalances(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) PostBalance(ctx context.Context, in *BalancePostRequest, opts ...grpc.CallOption) (*BalancePostResponse, error) {
	out := new(BalancePostResponse)
	err := c.cc.Invoke(ctx, "/api.API/PostBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetAccountsBalances(ctx context.Context, in *AccountsBalancesRequest, opts ...grpc.CallOption) (*AccountsBalancesResponse, error) {
	out := new(AccountsBalancesResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetAccountsBalances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) PostAccountsBalance(ctx context.Context, in *AccountBalancePostRequest, opts ...grpc.CallOption) (*AccountBalancePostResponse, error) {
	out := new(AccountBalancePostResponse)
	err := c.cc.Invoke(ctx, "/api.API/PostAccountsBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetTrades(ctx context.Context, in *TradeRequest, opts ...grpc.CallOption) (*TradeResponse, error) {
	out := new(TradeResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetTrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetMarketKindSummary(ctx context.Context, in *MarketKindSummaryRequest, opts ...grpc.CallOption) (*MarketKindSummaryResponse, error) {
	out := new(MarketKindSummaryResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetMarketKindSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetMarketProductSummary(ctx context.Context, in *MarketProductSummaryRequest, opts ...grpc.CallOption) (*MarketProductSummaryResponse, error) {
	out := new(MarketProductSummaryResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetMarketProductSummary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetTradesDeadman(ctx context.Context, in *TradesDeadmanRequest, opts ...grpc.CallOption) (*TradesDeadmanResponse, error) {
	out := new(TradesDeadmanResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetTradesDeadman", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetOrderbooksDeadman(ctx context.Context, in *OrderbookDeadmanRequest, opts ...grpc.CallOption) (*OrderbookDeadmanResponse, error) {
	out := new(OrderbookDeadmanResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetOrderbooksDeadman", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetBars(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*BarResponse, error) {
	out := new(BarResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetBars", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetProductWeekSeelWeek(ctx context.Context, in *BuySellWeekRequest, opts ...grpc.CallOption) (*BuySellWeekResponse, error) {
	out := new(BuySellWeekResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetProductWeekSeelWeek", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) OrderSend(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResult, error) {
	out := new(OrderResult)
	err := c.cc.Invoke(ctx, "/api.API/OrderSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetPositions(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionsResult, error) {
	out := new(PositionsResult)
	err := c.cc.Invoke(ctx, "/api.API/GetPositions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetHistoticalTrades(ctx context.Context, in *HistoricalTradesRequest, opts ...grpc.CallOption) (API_GetHistoticalTradesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[0], "/api.API/GetHistoticalTrades", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIGetHistoticalTradesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_GetHistoticalTradesClient interface {
	Recv() (*HistoricalTradesResult, error)
	grpc.ClientStream
}

type aPIGetHistoticalTradesClient struct {
	grpc.ClientStream
}

func (x *aPIGetHistoticalTradesClient) Recv() (*HistoricalTradesResult, error) {
	m := new(HistoricalTradesResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) GetBacktestingToken(ctx context.Context, in *BacktestingTokenRequest, opts ...grpc.CallOption) (*BacktestingTokenResponse, error) {
	out := new(BacktestingTokenResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetBacktestingToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Backtesting(ctx context.Context, opts ...grpc.CallOption) (API_BacktestingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[1], "/api.API/Backtesting", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIBacktestingClient{stream}
	return x, nil
}

type API_BacktestingClient interface {
	Send(*BacktestingRequest) error
	Recv() (*BacktestingResponse, error)
	grpc.ClientStream
}

type aPIBacktestingClient struct {
	grpc.ClientStream
}

func (x *aPIBacktestingClient) Send(m *BacktestingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aPIBacktestingClient) Recv() (*BacktestingResponse, error) {
	m := new(BacktestingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetCurrencies(context.Context, *CurrencyRequest) (*CurrencyResponse, error)
	PostCurrency(context.Context, *CurrencyPostRequest) (*CurrencyPostResponse, error)
	GetVenues(context.Context, *VenueRequest) (*VenueResponse, error)
	PostVenue(context.Context, *VenuePostRequest) (*VenuePostResponse, error)
	GetProducts(context.Context, *ProductRequest) (*ProductResponse, error)
	PostProduct(context.Context, *ProductPostRequest) (*ProductPostResponse, error)
	GetAccounts(context.Context, *AccountRequest) (*AccountResponse, error)
	PostAccount(context.Context, *AccountPostRequest) (*AccountPostResponse, error)
	GetBalances(context.Context, *BalanceRequest) (*BalanceResponse, error)
	PostBalance(context.Context, *BalancePostRequest) (*BalancePostResponse, error)
	GetAccountsBalances(context.Context, *AccountsBalancesRequest) (*AccountsBalancesResponse, error)
	PostAccountsBalance(context.Context, *AccountBalancePostRequest) (*AccountBalancePostResponse, error)
	GetTrades(context.Context, *TradeRequest) (*TradeResponse, error)
	GetMarketKindSummary(context.Context, *MarketKindSummaryRequest) (*MarketKindSummaryResponse, error)
	GetMarketProductSummary(context.Context, *MarketProductSummaryRequest) (*MarketProductSummaryResponse, error)
	GetTradesDeadman(context.Context, *TradesDeadmanRequest) (*TradesDeadmanResponse, error)
	GetOrderbooksDeadman(context.Context, *OrderbookDeadmanRequest) (*OrderbookDeadmanResponse, error)
	GetBars(context.Context, *BarRequest) (*BarResponse, error)
	GetProductWeekSeelWeek(context.Context, *BuySellWeekRequest) (*BuySellWeekResponse, error)
	// Get all Order for a user with filter - A server-to-client streaming RPC.
	// rpc GetOrders(OrderFilter) returns (stream Order) {}
	// Order request - A simple RPC
	OrderSend(context.Context, *OrderRequest) (*OrderResult, error)
	GetPositions(context.Context, *PositionRequest) (*PositionsResult, error)
	GetHistoticalTrades(*HistoricalTradesRequest, API_GetHistoticalTradesServer) error
	GetBacktestingToken(context.Context, *BacktestingTokenRequest) (*BacktestingTokenResponse, error)
	Backtesting(API_BacktestingServer) error
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetCurrencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetCurrencies(ctx, req.(*CurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_PostCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrencyPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).PostCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/PostCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).PostCurrency(ctx, req.(*CurrencyPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetVenues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VenueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetVenues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetVenues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetVenues(ctx, req.(*VenueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_PostVenue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VenuePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).PostVenue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/PostVenue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).PostVenue(ctx, req.(*VenuePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetProducts(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_PostProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).PostProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/PostProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).PostProduct(ctx, req.(*ProductPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetAccounts(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_PostAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).PostAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/PostAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).PostAccount(ctx, req.(*AccountPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetBalances(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_PostBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalancePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).PostBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/PostBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).PostBalance(ctx, req.(*BalancePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetAccountsBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountsBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetAccountsBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetAccountsBalances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetAccountsBalances(ctx, req.(*AccountsBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_PostAccountsBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountBalancePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).PostAccountsBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/PostAccountsBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).PostAccountsBalance(ctx, req.(*AccountBalancePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetTrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetTrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetTrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetTrades(ctx, req.(*TradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetMarketKindSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketKindSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetMarketKindSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetMarketKindSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetMarketKindSummary(ctx, req.(*MarketKindSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetMarketProductSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarketProductSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetMarketProductSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetMarketProductSummary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetMarketProductSummary(ctx, req.(*MarketProductSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetTradesDeadman_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TradesDeadmanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetTradesDeadman(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetTradesDeadman",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetTradesDeadman(ctx, req.(*TradesDeadmanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetOrderbooksDeadman_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderbookDeadmanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetOrderbooksDeadman(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetOrderbooksDeadman",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetOrderbooksDeadman(ctx, req.(*OrderbookDeadmanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetBars_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetBars(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetBars",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetBars(ctx, req.(*BarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetProductWeekSeelWeek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuySellWeekRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetProductWeekSeelWeek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetProductWeekSeelWeek",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetProductWeekSeelWeek(ctx, req.(*BuySellWeekRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_OrderSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).OrderSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/OrderSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).OrderSend(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetPositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetPositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetPositions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetPositions(ctx, req.(*PositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetHistoticalTrades_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HistoricalTradesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).GetHistoticalTrades(m, &aPIGetHistoticalTradesServer{stream})
}

type API_GetHistoticalTradesServer interface {
	Send(*HistoricalTradesResult) error
	grpc.ServerStream
}

type aPIGetHistoticalTradesServer struct {
	grpc.ServerStream
}

func (x *aPIGetHistoticalTradesServer) Send(m *HistoricalTradesResult) error {
	return x.ServerStream.SendMsg(m)
}

func _API_GetBacktestingToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BacktestingTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetBacktestingToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetBacktestingToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetBacktestingToken(ctx, req.(*BacktestingTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Backtesting_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(APIServer).Backtesting(&aPIBacktestingServer{stream})
}

type API_BacktestingServer interface {
	Send(*BacktestingResponse) error
	Recv() (*BacktestingRequest, error)
	grpc.ServerStream
}

type aPIBacktestingServer struct {
	grpc.ServerStream
}

func (x *aPIBacktestingServer) Send(m *BacktestingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aPIBacktestingServer) Recv() (*BacktestingRequest, error) {
	m := new(BacktestingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _API_Login_Handler,
		},
		{
			MethodName: "GetCurrencies",
			Handler:    _API_GetCurrencies_Handler,
		},
		{
			MethodName: "PostCurrency",
			Handler:    _API_PostCurrency_Handler,
		},
		{
			MethodName: "GetVenues",
			Handler:    _API_GetVenues_Handler,
		},
		{
			MethodName: "PostVenue",
			Handler:    _API_PostVenue_Handler,
		},
		{
			MethodName: "GetProducts",
			Handler:    _API_GetProducts_Handler,
		},
		{
			MethodName: "PostProduct",
			Handler:    _API_PostProduct_Handler,
		},
		{
			MethodName: "GetAccounts",
			Handler:    _API_GetAccounts_Handler,
		},
		{
			MethodName: "PostAccount",
			Handler:    _API_PostAccount_Handler,
		},
		{
			MethodName: "GetBalances",
			Handler:    _API_GetBalances_Handler,
		},
		{
			MethodName: "PostBalance",
			Handler:    _API_PostBalance_Handler,
		},
		{
			MethodName: "GetAccountsBalances",
			Handler:    _API_GetAccountsBalances_Handler,
		},
		{
			MethodName: "PostAccountsBalance",
			Handler:    _API_PostAccountsBalance_Handler,
		},
		{
			MethodName: "GetTrades",
			Handler:    _API_GetTrades_Handler,
		},
		{
			MethodName: "GetMarketKindSummary",
			Handler:    _API_GetMarketKindSummary_Handler,
		},
		{
			MethodName: "GetMarketProductSummary",
			Handler:    _API_GetMarketProductSummary_Handler,
		},
		{
			MethodName: "GetTradesDeadman",
			Handler:    _API_GetTradesDeadman_Handler,
		},
		{
			MethodName: "GetOrderbooksDeadman",
			Handler:    _API_GetOrderbooksDeadman_Handler,
		},
		{
			MethodName: "GetBars",
			Handler:    _API_GetBars_Handler,
		},
		{
			MethodName: "GetProductWeekSeelWeek",
			Handler:    _API_GetProductWeekSeelWeek_Handler,
		},
		{
			MethodName: "OrderSend",
			Handler:    _API_OrderSend_Handler,
		},
		{
			MethodName: "GetPositions",
			Handler:    _API_GetPositions_Handler,
		},
		{
			MethodName: "GetBacktestingToken",
			Handler:    _API_GetBacktestingToken_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetHistoticalTrades",
			Handler:       _API_GetHistoticalTrades_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Backtesting",
			Handler:       _API_Backtesting_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}
