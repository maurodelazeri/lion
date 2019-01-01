// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

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

// Events
type Event struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Event                string   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	User                 string   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Account              string   `protobuf:"bytes,4,opt,name=account,proto3" json:"account,omitempty"`
	Container            string   `protobuf:"bytes,5,opt,name=container,proto3" json:"container,omitempty"`
	Strategy             string   `protobuf:"bytes,6,opt,name=strategy,proto3" json:"strategy,omitempty"`
	Timestamp            int64    `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Command              string   `protobuf:"bytes,8,opt,name=command,proto3" json:"command,omitempty"`
	Message              string   `protobuf:"bytes,9,opt,name=message,proto3" json:"message,omitempty"`
	Error                bool     `protobuf:"varint,10,opt,name=error,proto3" json:"error,omitempty"`
	Payload              string   `protobuf:"bytes,11,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *Event) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Event) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Event) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *Event) GetStrategy() string {
	if m != nil {
		return m.Strategy
	}
	return ""
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Event) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Event) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

func (m *Event) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "api.Event")
}

func init() { proto.RegisterFile("events.proto", fileDescriptor_8f22242cb04491f9) }

var fileDescriptor_8f22242cb04491f9 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x95, 0xa4, 0x69, 0x93, 0x01, 0xb1, 0xb0, 0x58, 0x8c, 0x10, 0x8b, 0x88, 0x55, 0x56,
	0x6c, 0x38, 0x03, 0x17, 0xc8, 0x0d, 0x86, 0x64, 0x54, 0x59, 0xc2, 0x3f, 0xb2, 0xa7, 0x48, 0x3d,
	0x09, 0xd7, 0x45, 0x1e, 0x87, 0x76, 0xe7, 0xef, 0x7b, 0xcf, 0x7a, 0xb2, 0xe1, 0x91, 0x7f, 0xd8,
	0x4b, 0x7e, 0x8f, 0x29, 0x48, 0x30, 0x1d, 0x45, 0xfb, 0xf6, 0xdb, 0x42, 0xff, 0x59, 0xac, 0x79,
	0x82, 0xd6, 0x6e, 0xd8, 0x4c, 0xcd, 0x3c, 0x2e, 0xad, 0xdd, 0xcc, 0x33, 0xf4, 0x5a, 0xc7, 0x56,
	0x55, 0x05, 0x63, 0xe0, 0x70, 0xc9, 0x9c, 0xb0, 0x53, 0xa9, 0x67, 0x83, 0x70, 0xa2, 0x75, 0x0d,
	0x17, 0x2f, 0x78, 0x50, 0xfd, 0x8f, 0xe6, 0x15, 0xc6, 0x35, 0x78, 0x21, 0xeb, 0x39, 0x61, 0xaf,
	0xd9, 0x5d, 0x98, 0x17, 0x18, 0xb2, 0x24, 0x12, 0x3e, 0x5f, 0xf1, 0xa8, 0xe1, 0x8d, 0xcb, 0x4d,
	0xb1, 0x8e, 0xb3, 0x90, 0x8b, 0x78, 0x9a, 0x9a, 0xb9, 0x5b, 0xee, 0xa2, 0x2c, 0xae, 0xc1, 0x39,
	0xf2, 0x1b, 0x0e, 0x75, 0x71, 0xc7, 0x92, 0x38, 0xce, 0x99, 0xce, 0x8c, 0x63, 0x4d, 0x76, 0xd4,
	0xf7, 0xa4, 0x14, 0x12, 0xc2, 0xd4, 0xcc, 0xc3, 0x52, 0xa1, 0xf4, 0x23, 0x5d, 0xbf, 0x03, 0x6d,
	0xf8, 0x50, 0xfb, 0x3b, 0x7e, 0x1d, 0xf5, 0x97, 0x3e, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc9,
	0x5b, 0x39, 0x7e, 0x35, 0x01, 0x00, 0x00,
}