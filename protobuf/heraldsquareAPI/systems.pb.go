// Code generated by protoc-gen-go. DO NOT EDIT.
// source: systems.proto

package heraldsquareAPI

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

type SystemSummer struct {
	Action               Action   `protobuf:"varint,1,opt,name=action,proto3,enum=heraldsquareAPI.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SystemSummer) Reset()         { *m = SystemSummer{} }
func (m *SystemSummer) String() string { return proto.CompactTextString(m) }
func (*SystemSummer) ProtoMessage()    {}
func (*SystemSummer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec938d4cda008df6, []int{0}
}
func (m *SystemSummer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemSummer.Unmarshal(m, b)
}
func (m *SystemSummer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemSummer.Marshal(b, m, deterministic)
}
func (dst *SystemSummer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemSummer.Merge(dst, src)
}
func (m *SystemSummer) XXX_Size() int {
	return xxx_messageInfo_SystemSummer.Size(m)
}
func (m *SystemSummer) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemSummer.DiscardUnknown(m)
}

var xxx_messageInfo_SystemSummer proto.InternalMessageInfo

func (m *SystemSummer) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_INSERT
}

func init() {
	proto.RegisterType((*SystemSummer)(nil), "heraldsquareAPI.SystemSummer")
}

func init() { proto.RegisterFile("systems.proto", fileDescriptor_ec938d4cda008df6) }

var fileDescriptor_ec938d4cda008df6 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xae, 0x2c, 0x2e,
	0x49, 0xcd, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcf, 0x48, 0x2d, 0x4a, 0xcc,
	0x49, 0x29, 0x2e, 0x2c, 0x4d, 0x2c, 0x4a, 0x75, 0x0c, 0xf0, 0x94, 0xe2, 0x4a, 0xcd, 0x2b, 0xcd,
	0x85, 0x48, 0x2a, 0xd9, 0x73, 0xf1, 0x04, 0x83, 0x55, 0x07, 0x97, 0xe6, 0xe6, 0xa6, 0x16, 0x09,
	0xe9, 0x73, 0xb1, 0x25, 0x26, 0x97, 0x64, 0xe6, 0xe7, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19,
	0x89, 0xeb, 0xa1, 0xe9, 0xd6, 0x73, 0x04, 0x4b, 0x07, 0x41, 0x95, 0x25, 0xb1, 0x81, 0xcd, 0x31,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xa0, 0x3a, 0xe2, 0x75, 0x00, 0x00, 0x00,
}
