// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events.proto

package heraldsquareAPI

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

type Event struct {
	EventId              int64    `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	System               System   `protobuf:"varint,2,opt,name=system,proto3,enum=heraldsquareAPI.System" json:"system,omitempty"`
	UserId               string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Timestamp            string   `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Function             string   `protobuf:"bytes,6,opt,name=function,proto3" json:"function,omitempty"`
	Payload              string   `protobuf:"bytes,7,opt,name=payload,proto3" json:"payload,omitempty"`
	Error                bool     `protobuf:"varint,8,opt,name=error,proto3" json:"error,omitempty"`
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

func (m *Event) GetEventId() int64 {
	if m != nil {
		return m.EventId
	}
	return 0
}

func (m *Event) GetSystem() System {
	if m != nil {
		return m.System
	}
	return System_RESERVED
}

func (m *Event) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Event) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Event) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Event) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *Event) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *Event) GetError() bool {
	if m != nil {
		return m.Error
	}
	return false
}

type EventRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	EventId              int64    `protobuf:"varint,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{1}
}

func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (m *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(m, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

func (m *EventRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *EventRequest) GetEventId() int64 {
	if m != nil {
		return m.EventId
	}
	return 0
}

type EventResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=heraldsquareAPI.Retcode" json:"retcode,omitempty"`
	Event                []*Event `protobuf:"bytes,2,rep,name=event,proto3" json:"event,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f22242cb04491f9, []int{2}
}

func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (m *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(m, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

func (m *EventResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECT
}

func (m *EventResponse) GetEvent() []*Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *EventResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *EventResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "heraldsquareAPI.Event")
	proto.RegisterType((*EventRequest)(nil), "heraldsquareAPI.EventRequest")
	proto.RegisterType((*EventResponse)(nil), "heraldsquareAPI.EventResponse")
}

func init() { proto.RegisterFile("events.proto", fileDescriptor_8f22242cb04491f9) }

var fileDescriptor_8f22242cb04491f9 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4b, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xe5, 0xf6, 0xe6, 0xd1, 0x73, 0x4b, 0x91, 0x2c, 0x44, 0x4d, 0xc5, 0x20, 0xea, 0x28,
	0x03, 0x54, 0xa4, 0xb0, 0x00, 0xc4, 0x80, 0x41, 0x67, 0xc8, 0x2c, 0x00, 0x99, 0xfa, 0x00, 0x15,
	0x71, 0x9c, 0xda, 0x0e, 0x52, 0x17, 0xc4, 0xfe, 0x58, 0x02, 0xb2, 0x9d, 0x80, 0x4a, 0x87, 0x5f,
	0xfe, 0x47, 0x74, 0x7e, 0xc3, 0x14, 0x3f, 0xb0, 0x71, 0x76, 0xd5, 0x1a, 0xed, 0x34, 0x3d, 0x7d,
	0x43, 0x23, 0x6a, 0x69, 0x77, 0x9d, 0x30, 0x78, 0xf7, 0xb0, 0x5e, 0x00, 0x36, 0x9d, 0x8a, 0xe2,
	0xf2, 0x8b, 0x40, 0x72, 0xef, 0xdd, 0xf4, 0x02, 0xf2, 0x10, 0x7b, 0xda, 0x4a, 0x46, 0x0a, 0x52,
	0x8e, 0x79, 0x16, 0x78, 0x2d, 0xe9, 0x35, 0xa4, 0x76, 0x6f, 0x1d, 0x2a, 0x36, 0x2a, 0x48, 0x39,
	0xab, 0xe6, 0xab, 0x3f, 0x95, 0xab, 0xc7, 0x20, 0xf3, 0xde, 0x46, 0xe7, 0x90, 0x75, 0x16, 0x8d,
	0xaf, 0x1a, 0x17, 0xa4, 0x9c, 0xf0, 0xd4, 0xe3, 0x5a, 0xd2, 0x4b, 0x98, 0xb8, 0xad, 0x42, 0xeb,
	0x84, 0x6a, 0xd9, 0xbf, 0x20, 0xfd, 0x7e, 0xa0, 0x0c, 0x32, 0x85, 0xd6, 0x8a, 0x57, 0x64, 0x49,
	0xd0, 0x06, 0xa4, 0x0b, 0xc8, 0x5f, 0xba, 0x66, 0xe3, 0xb6, 0xba, 0x61, 0x69, 0x90, 0x7e, 0xd8,
	0xa7, 0x5a, 0xb1, 0xaf, 0xb5, 0x90, 0x2c, 0x8b, 0xa9, 0x1e, 0xe9, 0x19, 0x24, 0x68, 0x8c, 0x36,
	0x2c, 0x2f, 0x48, 0x99, 0xf3, 0x08, 0xcb, 0x5b, 0x98, 0x86, 0x8b, 0x39, 0xee, 0x3a, 0xb4, 0xce,
	0xbb, 0x9c, 0x7e, 0xc7, 0x26, 0x5c, 0x3d, 0xe1, 0x11, 0x0e, 0xe6, 0x18, 0x1d, 0xcc, 0xb1, 0xfc,
	0x24, 0x70, 0xd2, 0x37, 0xd8, 0x56, 0x37, 0x16, 0x69, 0x05, 0x99, 0x41, 0xb7, 0xd1, 0x12, 0x43,
	0xc9, 0xac, 0x62, 0x47, 0x0b, 0xf1, 0xa8, 0xf3, 0xc1, 0x48, 0xaf, 0x20, 0x09, 0x85, 0x6c, 0x54,
	0x8c, 0xcb, 0xff, 0xd5, 0xf9, 0x51, 0x22, 0xfe, 0x22, 0x9a, 0xfc, 0x91, 0x1b, 0xad, 0x94, 0xf7,
	0xc7, 0x45, 0x07, 0xf4, 0x0a, 0xd6, 0xa2, 0xb5, 0x28, 0xfb, 0x41, 0x07, 0x7c, 0x4e, 0xc3, 0x13,
	0xdf, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0x30, 0x4a, 0x2b, 0x1f, 0x0f, 0x02, 0x00, 0x00,
}
