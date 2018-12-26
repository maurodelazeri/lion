// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Login
type LoginRequest struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "api.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "api.LoginResponse")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8e, 0x31, 0xae, 0xc2, 0x30,
	0x0c, 0x40, 0xd5, 0xdf, 0x0f, 0x05, 0xab, 0x30, 0x58, 0x0c, 0x51, 0x27, 0xd4, 0x01, 0x31, 0x75,
	0x80, 0x9d, 0x13, 0x30, 0xe5, 0x06, 0xa1, 0xb5, 0x4a, 0x85, 0x1a, 0x87, 0xc4, 0x11, 0xd7, 0x47,
	0x4a, 0x5b, 0x36, 0xbf, 0x67, 0xf9, 0xc9, 0x00, 0x26, 0xca, 0xb3, 0x71, 0x9e, 0x85, 0x31, 0x37,
	0x6e, 0xa8, 0x80, 0x6c, 0x1c, 0x27, 0x51, 0xdf, 0xa0, 0xbc, 0x73, 0x3f, 0x58, 0x4d, 0xef, 0x48,
	0x41, 0x10, 0xe1, 0x3f, 0x06, 0xf2, 0x2a, 0x3b, 0x66, 0xe7, 0xad, 0x4e, 0x33, 0x56, 0xb0, 0x71,
	0x26, 0x84, 0x0f, 0xfb, 0x4e, 0xfd, 0x25, 0xff, 0xe3, 0xba, 0x87, 0xdd, 0x7c, 0x1f, 0x1c, 0xdb,
	0x40, 0x78, 0x82, 0xc2, 0x93, 0xb4, 0xdc, 0x51, 0x6a, 0xec, 0x2f, 0x65, 0x63, 0xdc, 0xd0, 0xe8,
	0xc9, 0xe9, 0x65, 0x89, 0x07, 0x58, 0x09, 0xbf, 0xc8, 0xce, 0xc5, 0x09, 0x50, 0x41, 0xd1, 0xf2,
	0x38, 0x92, 0x15, 0x95, 0x27, 0xbf, 0xe0, 0x63, 0x9d, 0xfe, 0xbd, 0x7e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xac, 0xd2, 0x9e, 0x97, 0xce, 0x00, 0x00, 0x00,
}
