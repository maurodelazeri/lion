// Code generated by protoc-gen-go. DO NOT EDIT.
// source: west4API.proto

package west4API

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

// West4APIClient is the client API for West4API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type West4APIClient interface {
	GetDataFeeds(ctx context.Context, in *DataFeedRequest, opts ...grpc.CallOption) (*DataFeedResponse, error)
	PostDataFeed(ctx context.Context, in *DataFeedPostRequest, opts ...grpc.CallOption) (*DataFeedPostResponse, error)
	GetServers(ctx context.Context, in *ServerRequest, opts ...grpc.CallOption) (*ServerResponse, error)
	PostServer(ctx context.Context, in *ServerPostRequest, opts ...grpc.CallOption) (*ServerPostResponse, error)
	GetContainers(ctx context.Context, in *ContainerRequest, opts ...grpc.CallOption) (*ContainerResponse, error)
	PostContainer(ctx context.Context, in *ContainerPostRequest, opts ...grpc.CallOption) (*ContainerPostResponse, error)
	GetContainersCreate(ctx context.Context, in *ContainerCreateRequest, opts ...grpc.CallOption) (*ContainerCreateResponse, error)
	PostContainerCreate(ctx context.Context, in *ContainerCreatePostRequest, opts ...grpc.CallOption) (*ContainerCreatePostResponse, error)
}

type west4APIClient struct {
	cc *grpc.ClientConn
}

func NewWest4APIClient(cc *grpc.ClientConn) West4APIClient {
	return &west4APIClient{cc}
}

func (c *west4APIClient) GetDataFeeds(ctx context.Context, in *DataFeedRequest, opts ...grpc.CallOption) (*DataFeedResponse, error) {
	out := new(DataFeedResponse)
	err := c.cc.Invoke(ctx, "/west4API.west4API/GetDataFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *west4APIClient) PostDataFeed(ctx context.Context, in *DataFeedPostRequest, opts ...grpc.CallOption) (*DataFeedPostResponse, error) {
	out := new(DataFeedPostResponse)
	err := c.cc.Invoke(ctx, "/west4API.west4API/PostDataFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *west4APIClient) GetServers(ctx context.Context, in *ServerRequest, opts ...grpc.CallOption) (*ServerResponse, error) {
	out := new(ServerResponse)
	err := c.cc.Invoke(ctx, "/west4API.west4API/GetServers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *west4APIClient) PostServer(ctx context.Context, in *ServerPostRequest, opts ...grpc.CallOption) (*ServerPostResponse, error) {
	out := new(ServerPostResponse)
	err := c.cc.Invoke(ctx, "/west4API.west4API/PostServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *west4APIClient) GetContainers(ctx context.Context, in *ContainerRequest, opts ...grpc.CallOption) (*ContainerResponse, error) {
	out := new(ContainerResponse)
	err := c.cc.Invoke(ctx, "/west4API.west4API/GetContainers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *west4APIClient) PostContainer(ctx context.Context, in *ContainerPostRequest, opts ...grpc.CallOption) (*ContainerPostResponse, error) {
	out := new(ContainerPostResponse)
	err := c.cc.Invoke(ctx, "/west4API.west4API/PostContainer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *west4APIClient) GetContainersCreate(ctx context.Context, in *ContainerCreateRequest, opts ...grpc.CallOption) (*ContainerCreateResponse, error) {
	out := new(ContainerCreateResponse)
	err := c.cc.Invoke(ctx, "/west4API.west4API/GetContainersCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *west4APIClient) PostContainerCreate(ctx context.Context, in *ContainerCreatePostRequest, opts ...grpc.CallOption) (*ContainerCreatePostResponse, error) {
	out := new(ContainerCreatePostResponse)
	err := c.cc.Invoke(ctx, "/west4API.west4API/PostContainerCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// West4APIServer is the server API for West4API service.
type West4APIServer interface {
	GetDataFeeds(context.Context, *DataFeedRequest) (*DataFeedResponse, error)
	PostDataFeed(context.Context, *DataFeedPostRequest) (*DataFeedPostResponse, error)
	GetServers(context.Context, *ServerRequest) (*ServerResponse, error)
	PostServer(context.Context, *ServerPostRequest) (*ServerPostResponse, error)
	GetContainers(context.Context, *ContainerRequest) (*ContainerResponse, error)
	PostContainer(context.Context, *ContainerPostRequest) (*ContainerPostResponse, error)
	GetContainersCreate(context.Context, *ContainerCreateRequest) (*ContainerCreateResponse, error)
	PostContainerCreate(context.Context, *ContainerCreatePostRequest) (*ContainerCreatePostResponse, error)
}

func RegisterWest4APIServer(s *grpc.Server, srv West4APIServer) {
	s.RegisterService(&_West4API_serviceDesc, srv)
}

func _West4API_GetDataFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(West4APIServer).GetDataFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/west4API.west4API/GetDataFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(West4APIServer).GetDataFeeds(ctx, req.(*DataFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _West4API_PostDataFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataFeedPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(West4APIServer).PostDataFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/west4API.west4API/PostDataFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(West4APIServer).PostDataFeed(ctx, req.(*DataFeedPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _West4API_GetServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(West4APIServer).GetServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/west4API.west4API/GetServers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(West4APIServer).GetServers(ctx, req.(*ServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _West4API_PostServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(West4APIServer).PostServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/west4API.west4API/PostServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(West4APIServer).PostServer(ctx, req.(*ServerPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _West4API_GetContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(West4APIServer).GetContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/west4API.west4API/GetContainers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(West4APIServer).GetContainers(ctx, req.(*ContainerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _West4API_PostContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(West4APIServer).PostContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/west4API.west4API/PostContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(West4APIServer).PostContainer(ctx, req.(*ContainerPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _West4API_GetContainersCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(West4APIServer).GetContainersCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/west4API.west4API/GetContainersCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(West4APIServer).GetContainersCreate(ctx, req.(*ContainerCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _West4API_PostContainerCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerCreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(West4APIServer).PostContainerCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/west4API.west4API/PostContainerCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(West4APIServer).PostContainerCreate(ctx, req.(*ContainerCreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _West4API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "west4API.west4API",
	HandlerType: (*West4APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDataFeeds",
			Handler:    _West4API_GetDataFeeds_Handler,
		},
		{
			MethodName: "PostDataFeed",
			Handler:    _West4API_PostDataFeed_Handler,
		},
		{
			MethodName: "GetServers",
			Handler:    _West4API_GetServers_Handler,
		},
		{
			MethodName: "PostServer",
			Handler:    _West4API_PostServer_Handler,
		},
		{
			MethodName: "GetContainers",
			Handler:    _West4API_GetContainers_Handler,
		},
		{
			MethodName: "PostContainer",
			Handler:    _West4API_PostContainer_Handler,
		},
		{
			MethodName: "GetContainersCreate",
			Handler:    _West4API_GetContainersCreate_Handler,
		},
		{
			MethodName: "PostContainerCreate",
			Handler:    _West4API_PostContainerCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "west4API.proto",
}

func init() { proto.RegisterFile("west4API.proto", fileDescriptor_81ea3877c7e6a19f) }

var fileDescriptor_81ea3877c7e6a19f = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x59, 0x40, 0xe8, 0xd4, 0x00, 0xba, 0x0e, 0x40, 0x0a, 0xe5, 0x8f, 0x60, 0xed, 0x00,
	0xbc, 0x40, 0x55, 0x44, 0x54, 0x16, 0xaa, 0xb2, 0xb2, 0x18, 0x72, 0x48, 0x2c, 0x71, 0xb1, 0x0f,
	0x78, 0x13, 0x9e, 0x17, 0x19, 0x73, 0x87, 0x4b, 0x1c, 0x3a, 0xe6, 0x7e, 0xdf, 0xfd, 0xfc, 0xe9,
	0x14, 0xd8, 0xfa, 0x20, 0xcf, 0x57, 0xe3, 0xd9, 0x74, 0xb4, 0x70, 0x96, 0x2d, 0x6e, 0xca, 0x77,
	0xb9, 0xf3, 0x64, 0x1b, 0x36, 0x2f, 0x0d, 0x39, 0x1f, 0x59, 0xb9, 0x5d, 0x1b, 0x36, 0xcf, 0x44,
	0xb5, 0x0c, 0x0a, 0x4f, 0xee, 0x5d, 0xf9, 0xc5, 0xe7, 0x3a, 0xe8, 0x3a, 0x56, 0xd0, 0xab, 0x88,
	0xaf, 0x0d, 0x9b, 0x9b, 0xb0, 0x81, 0xfb, 0x23, 0x7d, 0x49, 0x86, 0x73, 0x7a, 0x7d, 0x23, 0xcf,
	0x65, 0x99, 0x43, 0x7e, 0x61, 0x1b, 0x4f, 0xa7, 0x6b, 0x78, 0x07, 0xbd, 0x99, 0xf5, 0x6a, 0xc2,
	0xc3, 0x76, 0x3a, 0x70, 0x91, 0x0d, 0xbb, 0xb0, 0x0a, 0xc7, 0x00, 0x15, 0xf1, 0x7d, 0xac, 0x8e,
	0xbb, 0xbf, 0xf9, 0x38, 0x12, 0xd1, 0x5e, 0x1b, 0xa8, 0x62, 0x0a, 0x10, 0xa4, 0x71, 0x8e, 0x83,
	0xbf, 0xc9, 0xb4, 0xcf, 0x41, 0x1e, 0xaa, 0xea, 0x16, 0x8a, 0x8a, 0x78, 0xa2, 0xb7, 0xc6, 0xe4,
	0x1a, 0x3a, 0x15, 0xd9, 0x20, 0xcb, 0xd4, 0x35, 0x87, 0x22, 0xd8, 0x15, 0xe1, 0x30, 0x93, 0x4f,
	0xcb, 0x1d, 0x75, 0x72, 0x75, 0x3e, 0x40, 0x7f, 0xa9, 0xdf, 0xc4, 0x91, 0x61, 0xc2, 0xe3, 0xcc,
	0x66, 0x44, 0xe2, 0x3e, 0xf9, 0x27, 0xa1, 0xf6, 0x1a, 0xfa, 0x4b, 0x8d, 0x7f, 0xec, 0x67, 0x9d,
	0xbb, 0x69, 0xfb, 0xf3, 0x15, 0x29, 0x79, 0xe5, 0x71, 0xe3, 0xfb, 0xff, 0xbc, 0xfc, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x2f, 0x07, 0xd7, 0x7c, 0xed, 0x02, 0x00, 0x00,
}
