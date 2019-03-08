// Code generated by protoc-gen-go. DO NOT EDIT.
// source: heraldsquareAPI.proto

package heraldsquareAPI

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

func init() { proto.RegisterFile("heraldsquareAPI.proto", fileDescriptor_951f66b9e303173e) }

var fileDescriptor_951f66b9e303173e = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x48, 0x2d, 0x4a,
	0xcc, 0x49, 0x29, 0x2e, 0x2c, 0x4d, 0x2c, 0x4a, 0x75, 0x0c, 0xf0, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x47, 0x13, 0x96, 0xe2, 0x49, 0x2d, 0x4b, 0xcd, 0x2b, 0x29, 0x86, 0x48, 0x1b,
	0xc5, 0x73, 0xa1, 0x2b, 0x10, 0xf2, 0xe1, 0xe2, 0x74, 0x4f, 0x2d, 0x71, 0x05, 0xab, 0x12, 0x92,
	0xd5, 0x43, 0x37, 0x16, 0x2c, 0x11, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x25, 0x87, 0x4b,
	0xba, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x89, 0x21, 0x89, 0x0d, 0x6c, 0x8f, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x5a, 0x0e, 0x28, 0x7b, 0x9f, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeraldsquareAPIClient is the client API for HeraldsquareAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeraldsquareAPIClient interface {
	GetEvents(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
}

type heraldsquareAPIClient struct {
	cc *grpc.ClientConn
}

func NewHeraldsquareAPIClient(cc *grpc.ClientConn) HeraldsquareAPIClient {
	return &heraldsquareAPIClient{cc}
}

func (c *heraldsquareAPIClient) GetEvents(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/heraldsquareAPI.heraldsquareAPI/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeraldsquareAPIServer is the server API for HeraldsquareAPI service.
type HeraldsquareAPIServer interface {
	GetEvents(context.Context, *EventRequest) (*EventResponse, error)
}

func RegisterHeraldsquareAPIServer(s *grpc.Server, srv HeraldsquareAPIServer) {
	s.RegisterService(&_HeraldsquareAPI_serviceDesc, srv)
}

func _HeraldsquareAPI_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeraldsquareAPIServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/heraldsquareAPI.heraldsquareAPI/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeraldsquareAPIServer).GetEvents(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HeraldsquareAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "heraldsquareAPI.heraldsquareAPI",
	HandlerType: (*HeraldsquareAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvents",
			Handler:    _HeraldsquareAPI_GetEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heraldsquareAPI.proto",
}
