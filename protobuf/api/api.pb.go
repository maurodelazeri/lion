// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
	GetTrades(ctx context.Context, in *TradeRequest, opts ...grpc.CallOption) (*TradeResponse, error)
	GetMarketKindSummary(ctx context.Context, in *MarketKindSummaryRequest, opts ...grpc.CallOption) (*MarketKindSummaryResponse, error)
	GetMarketProductSummary(ctx context.Context, in *MarketProductSummaryRequest, opts ...grpc.CallOption) (*MarketProductSummaryResponse, error)
	GetBars(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*BarResponse, error)
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

func (c *aPIClient) GetBars(ctx context.Context, in *BarRequest, opts ...grpc.CallOption) (*BarResponse, error) {
	out := new(BarResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetBars", in, out, opts...)
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
	GetTrades(context.Context, *TradeRequest) (*TradeResponse, error)
	GetMarketKindSummary(context.Context, *MarketKindSummaryRequest) (*MarketKindSummaryResponse, error)
	GetMarketProductSummary(context.Context, *MarketProductSummaryRequest) (*MarketProductSummaryResponse, error)
	GetBars(context.Context, *BarRequest) (*BarResponse, error)
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
			MethodName: "GetBars",
			Handler:    _API_GetBars_Handler,
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

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 529 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x25, 0xaa, 0x00, 0xc5, 0x49, 0xdb, 0xd4, 0x4d, 0x69, 0x09, 0x14, 0x01, 0x27, 0x4e, 0x51,
	0x54, 0x38, 0x21, 0x84, 0xd4, 0x20, 0x08, 0x08, 0x10, 0x11, 0x8d, 0xe0, 0x8a, 0xb3, 0x3b, 0x2a,
	0x56, 0x12, 0x7b, 0xb1, 0xbd, 0x88, 0x7c, 0x1a, 0x7f, 0x87, 0xd6, 0x19, 0xcf, 0xae, 0x1d, 0xd2,
	0x5b, 0xe6, 0xbd, 0x79, 0xef, 0x39, 0xe3, 0x59, 0xb3, 0xb6, 0x28, 0xe4, 0xb0, 0x30, 0xda, 0x69,
	0xbe, 0x27, 0x0a, 0x39, 0x60, 0xa2, 0x74, 0x3f, 0x37, 0xc0, 0xa0, 0xfb, 0x1b, 0x54, 0x09, 0x16,
	0xab, 0x43, 0xf8, 0x03, 0x59, 0xe9, 0xa4, 0x56, 0x08, 0x1c, 0xcd, 0x45, 0xb6, 0x70, 0x60, 0x9d,
	0x54, 0xd7, 0x08, 0xed, 0x1b, 0x28, 0xb4, 0x71, 0x41, 0x72, 0x20, 0xb2, 0x4c, 0x97, 0x8a, 0xea,
	0x5e, 0x56, 0x1a, 0x03, 0x2a, 0x93, 0x64, 0x7a, 0x50, 0x18, 0x9d, 0x97, 0x59, 0xad, 0x98, 0x8b,
	0xa5, 0x50, 0x19, 0xf1, 0xbd, 0x95, 0x30, 0x0b, 0x70, 0xb9, 0x70, 0x22, 0x1c, 0x6a, 0x83, 0x6c,
	0xaa, 0x8b, 0xbf, 0x8c, 0xed, 0x5d, 0x4e, 0x3f, 0xf0, 0x11, 0xbb, 0xfd, 0x49, 0x5f, 0x4b, 0xc5,
	0x8f, 0x86, 0xd5, 0x1f, 0xf2, 0xbf, 0xbf, 0xc2, 0xaf, 0x12, 0xac, 0x1b, 0xf0, 0x26, 0x64, 0x0b,
	0xad, 0x2c, 0x3c, 0xbd, 0xc5, 0x5f, 0xb3, 0xfd, 0x09, 0xb8, 0x37, 0x74, 0x20, 0xde, 0xf7, 0x6d,
	0x08, 0xac, 0x83, 0xf8, 0x24, 0x41, 0x49, 0xff, 0x96, 0x75, 0xa7, 0xda, 0x06, 0x83, 0x35, 0x3f,
	0x8b, 0x1a, 0x2b, 0x2a, 0x58, 0xdc, 0xff, 0x0f, 0x43, 0x36, 0x2f, 0x58, 0x7b, 0x02, 0xee, 0x9b,
	0x1f, 0x34, 0x1e, 0xde, 0x17, 0xf1, 0xe1, 0x11, 0x22, 0xd5, 0x2b, 0xd6, 0xae, 0x7c, 0x3c, 0xcc,
	0x4f, 0xea, 0x96, 0x66, 0xec, 0xbd, 0x14, 0x26, 0xf5, 0x4b, 0xd6, 0x99, 0x80, 0x9b, 0xe2, 0xe4,
	0xf9, 0xb1, 0x6f, 0xc4, 0x32, 0xa8, 0xfb, 0x31, 0x48, 0xda, 0x31, 0xeb, 0x54, 0x6e, 0x48, 0xf0,
	0xd3, 0x66, 0x5b, 0x33, 0xfd, 0x6c, 0x9b, 0x48, 0xf2, 0x2f, 0x71, 0x37, 0x30, 0x1f, 0xcb, 0x38,
	0x9f, 0xc0, 0x34, 0x1f, 0x09, 0xcc, 0xc7, 0x6a, 0x3b, 0x3f, 0x22, 0x92, 0xfc, 0x31, 0x6e, 0x1a,
	0xe6, 0x63, 0x19, 0xe7, 0x13, 0x98, 0xe6, 0x23, 0x81, 0xf9, 0x58, 0x6d, 0xe7, 0x47, 0x44, 0x72,
	0xe7, 0x33, 0x23, 0x72, 0xba, 0x73, 0x5f, 0xc4, 0x77, 0x8e, 0x10, 0xa9, 0xbe, 0xb3, 0xfe, 0x04,
	0xdc, 0x67, 0xbf, 0xfd, 0x1f, 0xa5, 0xca, 0xaf, 0xca, 0xd5, 0x4a, 0x98, 0x35, 0x3f, 0xf7, 0xdd,
	0x5b, 0x78, 0x30, 0x7b, 0xb4, 0x8b, 0x26, 0xe3, 0x1f, 0xec, 0x94, 0x8c, 0xf1, 0xc2, 0x82, 0xf7,
	0xe3, 0x86, 0x38, 0xa6, 0x82, 0xfd, 0x93, 0x1b, 0x3a, 0x28, 0x61, 0xc8, 0xee, 0xfa, 0x81, 0x1b,
	0xcb, 0x0f, 0x71, 0x2e, 0x26, 0x18, 0xf4, 0x6a, 0x80, 0xfa, 0x2f, 0x58, 0xfb, 0x8b, 0xc9, 0xc1,
	0x5c, 0x81, 0xca, 0x71, 0x40, 0xbe, 0x8e, 0x35, 0x08, 0xd9, 0x72, 0xe9, 0xfc, 0x27, 0xd1, 0xad,
	0x96, 0x5a, 0x5b, 0x59, 0x3d, 0x51, 0xe1, 0x73, 0x0e, 0x75, 0xb2, 0xd6, 0xa1, 0x8b, 0xd4, 0x33,
	0x76, 0x3c, 0x01, 0xf7, 0x5e, 0x5a, 0xa7, 0x9d, 0xcc, 0xc4, 0x12, 0x2f, 0xe7, 0xa1, 0x6f, 0xf7,
	0xb0, 0xa9, 0xe1, 0x60, 0xf6, 0x60, 0x07, 0xbb, 0xf1, 0x1c, 0xb5, 0xd0, 0x75, 0x5c, 0x3f, 0x93,
	0x33, 0xbd, 0x00, 0x85, 0xae, 0x29, 0x1c, 0x5c, 0xcf, 0x77, 0xb0, 0x34, 0x9d, 0x77, 0xac, 0xd3,
	0x60, 0x69, 0x05, 0x09, 0x49, 0x57, 0xb0, 0x41, 0x04, 0x8f, 0x67, 0xad, 0x51, 0x6b, 0x7e, 0xc7,
	0x3f, 0xa1, 0xcf, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x63, 0x39, 0x79, 0x3f, 0x03, 0x06, 0x00,
	0x00,
}
