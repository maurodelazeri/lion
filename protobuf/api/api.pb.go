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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xd1, 0x6f, 0xd3, 0x30,
	0x10, 0xc6, 0x57, 0x4d, 0x80, 0xea, 0x8e, 0x41, 0x3d, 0xc6, 0xda, 0x00, 0x2f, 0x7b, 0xe2, 0x69,
	0xaa, 0x86, 0xc4, 0xd3, 0x24, 0xb4, 0x6a, 0xa2, 0x4c, 0x42, 0x10, 0x41, 0xc5, 0xbb, 0x97, 0x9c,
	0x86, 0xb5, 0xc8, 0x0e, 0xf6, 0x05, 0xf1, 0x2f, 0xf2, 0x5f, 0x4d, 0x71, 0xee, 0x1c, 0x37, 0x6d,
	0xdf, 0x7a, 0xdf, 0x77, 0xdf, 0xcf, 0xee, 0xe5, 0x2c, 0xc6, 0xaa, 0xd6, 0x17, 0xb5, 0xb3, 0x68,
	0xe5, 0xa1, 0xaa, 0x75, 0x26, 0x54, 0x83, 0xbf, 0x3b, 0x21, 0x3b, 0xfa, 0x0b, 0xa6, 0x01, 0x4f,
	0xd5, 0x0b, 0xf8, 0x07, 0x45, 0x83, 0xda, 0x1a, 0x12, 0xa6, 0x77, 0xaa, 0x78, 0x40, 0xf0, 0xa8,
	0xcd, 0x3d, 0x49, 0xcf, 0x1d, 0xd4, 0xd6, 0x21, 0x47, 0x8e, 0x55, 0x51, 0xd8, 0xc6, 0x70, 0x7d,
	0xf9, 0xff, 0x99, 0x38, 0xbc, 0xce, 0x6f, 0xe5, 0xa5, 0x18, 0x7f, 0x77, 0x25, 0xb8, 0x9f, 0x60,
	0x4a, 0x39, 0xbd, 0x68, 0xaf, 0x10, 0xea, 0x1f, 0xf0, 0xa7, 0x01, 0x8f, 0xd9, 0xcb, 0x54, 0xf2,
	0x4d, 0x85, 0xe7, 0x07, 0xf2, 0x4a, 0x1c, 0xad, 0x00, 0x73, 0xeb, 0x75, 0x7b, 0x05, 0x2f, 0x5f,
	0x85, 0x1e, 0xae, 0x39, 0xb9, 0xa9, 0xfa, 0x98, 0x5e, 0x8b, 0x93, 0x15, 0xe0, 0x17, 0xed, 0xd1,
	0xa2, 0x2e, 0x54, 0xb5, 0x76, 0xaa, 0x04, 0x2f, 0xdf, 0x86, 0xf6, 0x20, 0xbb, 0x5e, 0x66, 0xd8,
	0x9b, 0x3d, 0x6e, 0xc7, 0x5c, 0x8c, 0x88, 0xba, 0xec, 0xc7, 0xb0, 0xb6, 0x0f, 0x60, 0x88, 0x3a,
	0x94, 0x99, 0xfa, 0x6e, 0x8f, 0xeb, 0x6b, 0x6b, 0x3c, 0x9c, 0x1f, 0xc8, 0xcf, 0x62, 0x92, 0xb8,
	0xf2, 0x6c, 0xd8, 0xcf, 0xa0, 0xd9, 0xb6, 0xc1, 0x8c, 0xf7, 0xa3, 0xc5, 0x48, 0x7e, 0x14, 0xe3,
	0x15, 0xe0, 0xaf, 0xf0, 0x0d, 0xa5, 0x0c, 0xcd, 0x5d, 0xc1, 0x80, 0x93, 0x0d, 0x2d, 0x9e, 0xff,
	0x49, 0x88, 0xdc, 0x7a, 0x0e, 0xbe, 0x4e, 0x9a, 0x5a, 0x99, 0xc3, 0x67, 0x5b, 0x7a, 0x04, 0xdc,
	0x8a, 0x63, 0x3e, 0xf8, 0x06, 0x50, 0xe9, 0x4a, 0xce, 0xfb, 0xe6, 0x4e, 0x81, 0x92, 0x39, 0xd9,
	0x2e, 0x2b, 0xa2, 0xbe, 0x89, 0x53, 0x46, 0xe5, 0xce, 0x96, 0x4d, 0x81, 0x44, 0x9c, 0xf5, 0x31,
	0x32, 0x18, 0x38, 0xdf, 0xe1, 0x44, 0xde, 0x52, 0x4c, 0xda, 0xcb, 0x92, 0x41, 0xb3, 0xa5, 0x2a,
	0xfd, 0x77, 0xb3, 0x6d, 0x23, 0x32, 0x16, 0xe2, 0xc9, 0x57, 0x7b, 0xaf, 0x0d, 0x6d, 0x6e, 0xf8,
	0xcd, 0x39, 0x99, 0x4a, 0x31, 0x71, 0x25, 0x26, 0x2b, 0xc0, 0x6b, 0x7a, 0x0c, 0xb4, 0xba, 0x5c,
	0x72, 0xf4, 0x74, 0xa0, 0x26, 0x33, 0x98, 0xf6, 0xe9, 0xa5, 0xaa, 0x94, 0x29, 0x40, 0x66, 0x69,
	0x37, 0x89, 0x9b, 0x7b, 0x3b, 0xf4, 0x22, 0xef, 0xa6, 0x9b, 0x01, 0xf9, 0x34, 0x49, 0x3e, 0x37,
	0x1d, 0xc2, 0x7c, 0x87, 0xc3, 0x94, 0xbb, 0xa7, 0xe1, 0x49, 0x7f, 0x78, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0x54, 0x77, 0x93, 0x7d, 0x41, 0x04, 0x00, 0x00,
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
	// Get all Order for a user with filter - A server-to-client streaming RPC.
	// rpc GetOrders(OrderFilter) returns (stream Order) {}
	// Order request - A simple RPC
	OrderSend(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResult, error)
	GetPositions(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionsResult, error)
	GetHistoticalTrades(ctx context.Context, in *HistoricalTradesRequest, opts ...grpc.CallOption) (API_GetHistoticalTradesClient, error)
	GetBacktestingToken(ctx context.Context, in *BacktestingTokenRequest, opts ...grpc.CallOption) (*BacktestingTokenResponse, error)
	Backtesting(ctx context.Context, opts ...grpc.CallOption) (API_BacktestingClient, error)
	GetVenues(ctx context.Context, in *VenuesRequest, opts ...grpc.CallOption) (*VenuesResponse, error)
	PostVenues(ctx context.Context, in *VenuesPostRequest, opts ...grpc.CallOption) (*VenuesPostResponse, error)
	GetVenueDetail(ctx context.Context, in *VenueDetailedRequest, opts ...grpc.CallOption) (*VenueDetailedResponse, error)
	GetVenueProductDetail(ctx context.Context, in *VenueProductRequest, opts ...grpc.CallOption) (*VenueProductResponse, error)
	PostProduct(ctx context.Context, in *ProductPostRequest, opts ...grpc.CallOption) (*ProductPostResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	GetAccounts(ctx context.Context, in *AccountsRequest, opts ...grpc.CallOption) (*AccountsResponse, error)
	GetAccountBalance(ctx context.Context, in *AccountBalanceRequest, opts ...grpc.CallOption) (*AccountBalanceResponse, error)
	PostAccount(ctx context.Context, in *AccountsPostRequest, opts ...grpc.CallOption) (*AccountsPostResponse, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
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

func (c *aPIClient) GetVenues(ctx context.Context, in *VenuesRequest, opts ...grpc.CallOption) (*VenuesResponse, error) {
	out := new(VenuesResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetVenues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) PostVenues(ctx context.Context, in *VenuesPostRequest, opts ...grpc.CallOption) (*VenuesPostResponse, error) {
	out := new(VenuesPostResponse)
	err := c.cc.Invoke(ctx, "/api.API/PostVenues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetVenueDetail(ctx context.Context, in *VenueDetailedRequest, opts ...grpc.CallOption) (*VenueDetailedResponse, error) {
	out := new(VenueDetailedResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetVenueDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetVenueProductDetail(ctx context.Context, in *VenueProductRequest, opts ...grpc.CallOption) (*VenueProductResponse, error) {
	out := new(VenueProductResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetVenueProductDetail", in, out, opts...)
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

func (c *aPIClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/api.API/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetAccounts(ctx context.Context, in *AccountsRequest, opts ...grpc.CallOption) (*AccountsResponse, error) {
	out := new(AccountsResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) GetAccountBalance(ctx context.Context, in *AccountBalanceRequest, opts ...grpc.CallOption) (*AccountBalanceResponse, error) {
	out := new(AccountBalanceResponse)
	err := c.cc.Invoke(ctx, "/api.API/GetAccountBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) PostAccount(ctx context.Context, in *AccountsPostRequest, opts ...grpc.CallOption) (*AccountsPostResponse, error) {
	out := new(AccountsPostResponse)
	err := c.cc.Invoke(ctx, "/api.API/PostAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	// Get all Order for a user with filter - A server-to-client streaming RPC.
	// rpc GetOrders(OrderFilter) returns (stream Order) {}
	// Order request - A simple RPC
	OrderSend(context.Context, *OrderRequest) (*OrderResult, error)
	GetPositions(context.Context, *PositionRequest) (*PositionsResult, error)
	GetHistoticalTrades(*HistoricalTradesRequest, API_GetHistoticalTradesServer) error
	GetBacktestingToken(context.Context, *BacktestingTokenRequest) (*BacktestingTokenResponse, error)
	Backtesting(API_BacktestingServer) error
	GetVenues(context.Context, *VenuesRequest) (*VenuesResponse, error)
	PostVenues(context.Context, *VenuesPostRequest) (*VenuesPostResponse, error)
	GetVenueDetail(context.Context, *VenueDetailedRequest) (*VenueDetailedResponse, error)
	GetVenueProductDetail(context.Context, *VenueProductRequest) (*VenueProductResponse, error)
	PostProduct(context.Context, *ProductPostRequest) (*ProductPostResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	GetAccounts(context.Context, *AccountsRequest) (*AccountsResponse, error)
	GetAccountBalance(context.Context, *AccountBalanceRequest) (*AccountBalanceResponse, error)
	PostAccount(context.Context, *AccountsPostRequest) (*AccountsPostResponse, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
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

func _API_GetVenues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VenuesRequest)
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
		return srv.(APIServer).GetVenues(ctx, req.(*VenuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_PostVenues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VenuesPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).PostVenues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/PostVenues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).PostVenues(ctx, req.(*VenuesPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetVenueDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VenueDetailedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetVenueDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetVenueDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetVenueDetail(ctx, req.(*VenueDetailedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetVenueProductDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VenueProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetVenueProductDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetVenueProductDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetVenueProductDetail(ctx, req.(*VenueProductRequest))
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

func _API_GetAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountsRequest)
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
		return srv.(APIServer).GetAccounts(ctx, req.(*AccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_GetAccountBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetAccountBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/GetAccountBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetAccountBalance(ctx, req.(*AccountBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_PostAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountsPostRequest)
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
		return srv.(APIServer).PostAccount(ctx, req.(*AccountsPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
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
		{
			MethodName: "GetVenues",
			Handler:    _API_GetVenues_Handler,
		},
		{
			MethodName: "PostVenues",
			Handler:    _API_PostVenues_Handler,
		},
		{
			MethodName: "GetVenueDetail",
			Handler:    _API_GetVenueDetail_Handler,
		},
		{
			MethodName: "GetVenueProductDetail",
			Handler:    _API_GetVenueProductDetail_Handler,
		},
		{
			MethodName: "PostProduct",
			Handler:    _API_PostProduct_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _API_Login_Handler,
		},
		{
			MethodName: "GetAccounts",
			Handler:    _API_GetAccounts_Handler,
		},
		{
			MethodName: "GetAccountBalance",
			Handler:    _API_GetAccountBalance_Handler,
		},
		{
			MethodName: "PostAccount",
			Handler:    _API_PostAccount_Handler,
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
