// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

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

// Login
type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
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
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
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
	Elapsed              string   `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
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
func (dst *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(dst, src)
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

func (m *LoginResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "api.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "api.LoginResponse")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0x31, 0xae, 0xc2, 0x30,
	0x10, 0x05, 0x95, 0x9f, 0x0f, 0x81, 0x55, 0xa0, 0xb0, 0x28, 0xac, 0x54, 0x28, 0x05, 0xa2, 0x4a,
	0x01, 0x17, 0xe0, 0x00, 0x54, 0xbe, 0x81, 0x49, 0x56, 0x60, 0x11, 0x7b, 0x4d, 0xec, 0x88, 0x9a,
	0x9b, 0xa3, 0xd8, 0x0e, 0xe5, 0xbc, 0x91, 0x66, 0xb5, 0x00, 0x72, 0xf4, 0x8f, 0xc6, 0x0e, 0xe4,
	0x89, 0xe5, 0xd2, 0xaa, 0x0a, 0xd0, 0x8c, 0x3a, 0x0e, 0xf5, 0x05, 0xca, 0x2b, 0xdd, 0x95, 0x11,
	0xf8, 0x1a, 0xd1, 0x79, 0xb6, 0x83, 0x05, 0x6a, 0xa9, 0x7a, 0x9e, 0xed, 0xb3, 0xe3, 0x5a, 0x44,
	0x60, 0x15, 0xac, 0xac, 0x74, 0xee, 0x4d, 0x43, 0xc7, 0xff, 0x82, 0xf8, 0x71, 0xfd, 0xc9, 0x60,
	0x93, 0x12, 0xce, 0x92, 0x71, 0xc8, 0x0e, 0x50, 0x0c, 0xe8, 0x5b, 0xea, 0x30, 0x54, 0xb6, 0xa7,
	0xb2, 0x91, 0x56, 0x35, 0x22, 0x6e, 0x62, 0x96, 0xd3, 0x2d, 0x4f, 0x4f, 0x34, 0x29, 0x19, 0x81,
	0x71, 0x28, 0x5a, 0xd2, 0x1a, 0x8d, 0xe7, 0x79, 0xd8, 0x67, 0x9c, 0x0c, 0xf6, 0xd2, 0x3a, 0xec,
	0xf8, 0x7f, 0x34, 0x09, 0x6f, 0xcb, 0xf0, 0xcc, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x9f,
	0xea, 0xd9, 0xeb, 0x00, 0x00, 0x00,
}
