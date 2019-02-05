// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datafeeds.proto

package west4API

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

type DataFeed struct {
	DatafeedId           int32    `protobuf:"varint,1,opt,name=datafeed_id,json=datafeedId,proto3" json:"datafeed_id,omitempty"`
	DatafeedDescription  string   `protobuf:"bytes,2,opt,name=datafeed_description,json=datafeedDescription,proto3" json:"datafeed_description,omitempty"`
	VenueId              int32    `protobuf:"varint,3,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	VenueDescription     string   `protobuf:"bytes,4,opt,name=venue_description,json=venueDescription,proto3" json:"venue_description,omitempty"`
	ContainerId          int32    `protobuf:"varint,5,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ContainerDescription string   `protobuf:"bytes,6,opt,name=container_description,json=containerDescription,proto3" json:"container_description,omitempty"`
	Enabled              bool     `protobuf:"varint,7,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataFeed) Reset()         { *m = DataFeed{} }
func (m *DataFeed) String() string { return proto.CompactTextString(m) }
func (*DataFeed) ProtoMessage()    {}
func (*DataFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_6827cc0956531667, []int{0}
}

func (m *DataFeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataFeed.Unmarshal(m, b)
}
func (m *DataFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataFeed.Marshal(b, m, deterministic)
}
func (m *DataFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataFeed.Merge(m, src)
}
func (m *DataFeed) XXX_Size() int {
	return xxx_messageInfo_DataFeed.Size(m)
}
func (m *DataFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_DataFeed.DiscardUnknown(m)
}

var xxx_messageInfo_DataFeed proto.InternalMessageInfo

func (m *DataFeed) GetDatafeedId() int32 {
	if m != nil {
		return m.DatafeedId
	}
	return 0
}

func (m *DataFeed) GetDatafeedDescription() string {
	if m != nil {
		return m.DatafeedDescription
	}
	return ""
}

func (m *DataFeed) GetVenueId() int32 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *DataFeed) GetVenueDescription() string {
	if m != nil {
		return m.VenueDescription
	}
	return ""
}

func (m *DataFeed) GetContainerId() int32 {
	if m != nil {
		return m.ContainerId
	}
	return 0
}

func (m *DataFeed) GetContainerDescription() string {
	if m != nil {
		return m.ContainerDescription
	}
	return ""
}

func (m *DataFeed) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type DataFeedRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	DatafeedId           int32    `protobuf:"varint,2,opt,name=datafeed_id,json=datafeedId,proto3" json:"datafeed_id,omitempty"`
	VenueId              int32    `protobuf:"varint,3,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataFeedRequest) Reset()         { *m = DataFeedRequest{} }
func (m *DataFeedRequest) String() string { return proto.CompactTextString(m) }
func (*DataFeedRequest) ProtoMessage()    {}
func (*DataFeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6827cc0956531667, []int{1}
}

func (m *DataFeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataFeedRequest.Unmarshal(m, b)
}
func (m *DataFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataFeedRequest.Marshal(b, m, deterministic)
}
func (m *DataFeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataFeedRequest.Merge(m, src)
}
func (m *DataFeedRequest) XXX_Size() int {
	return xxx_messageInfo_DataFeedRequest.Size(m)
}
func (m *DataFeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataFeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataFeedRequest proto.InternalMessageInfo

func (m *DataFeedRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DataFeedRequest) GetDatafeedId() int32 {
	if m != nil {
		return m.DatafeedId
	}
	return 0
}

func (m *DataFeedRequest) GetVenueId() int32 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

type DataFeedResponse struct {
	Retcode              Retcode     `protobuf:"varint,1,opt,name=retcode,proto3,enum=west4API.Retcode" json:"retcode,omitempty"`
	Datafeed             []*DataFeed `protobuf:"bytes,2,rep,name=datafeed,proto3" json:"datafeed,omitempty"`
	Comment              string      `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string      `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DataFeedResponse) Reset()         { *m = DataFeedResponse{} }
func (m *DataFeedResponse) String() string { return proto.CompactTextString(m) }
func (*DataFeedResponse) ProtoMessage()    {}
func (*DataFeedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6827cc0956531667, []int{2}
}

func (m *DataFeedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataFeedResponse.Unmarshal(m, b)
}
func (m *DataFeedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataFeedResponse.Marshal(b, m, deterministic)
}
func (m *DataFeedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataFeedResponse.Merge(m, src)
}
func (m *DataFeedResponse) XXX_Size() int {
	return xxx_messageInfo_DataFeedResponse.Size(m)
}
func (m *DataFeedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataFeedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataFeedResponse proto.InternalMessageInfo

func (m *DataFeedResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECT
}

func (m *DataFeedResponse) GetDatafeed() []*DataFeed {
	if m != nil {
		return m.Datafeed
	}
	return nil
}

func (m *DataFeedResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *DataFeedResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type DataFeedPostRequest struct {
	Token                string    `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Datafeed             *DataFeed `protobuf:"bytes,2,opt,name=datafeed,proto3" json:"datafeed,omitempty"`
	Action               Action    `protobuf:"varint,3,opt,name=action,proto3,enum=west4API.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DataFeedPostRequest) Reset()         { *m = DataFeedPostRequest{} }
func (m *DataFeedPostRequest) String() string { return proto.CompactTextString(m) }
func (*DataFeedPostRequest) ProtoMessage()    {}
func (*DataFeedPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6827cc0956531667, []int{3}
}

func (m *DataFeedPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataFeedPostRequest.Unmarshal(m, b)
}
func (m *DataFeedPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataFeedPostRequest.Marshal(b, m, deterministic)
}
func (m *DataFeedPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataFeedPostRequest.Merge(m, src)
}
func (m *DataFeedPostRequest) XXX_Size() int {
	return xxx_messageInfo_DataFeedPostRequest.Size(m)
}
func (m *DataFeedPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataFeedPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataFeedPostRequest proto.InternalMessageInfo

func (m *DataFeedPostRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DataFeedPostRequest) GetDatafeed() *DataFeed {
	if m != nil {
		return m.Datafeed
	}
	return nil
}

func (m *DataFeedPostRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_INSERT
}

type DataFeedPostResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=west4API.Retcode" json:"retcode,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,3,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataFeedPostResponse) Reset()         { *m = DataFeedPostResponse{} }
func (m *DataFeedPostResponse) String() string { return proto.CompactTextString(m) }
func (*DataFeedPostResponse) ProtoMessage()    {}
func (*DataFeedPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6827cc0956531667, []int{4}
}

func (m *DataFeedPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataFeedPostResponse.Unmarshal(m, b)
}
func (m *DataFeedPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataFeedPostResponse.Marshal(b, m, deterministic)
}
func (m *DataFeedPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataFeedPostResponse.Merge(m, src)
}
func (m *DataFeedPostResponse) XXX_Size() int {
	return xxx_messageInfo_DataFeedPostResponse.Size(m)
}
func (m *DataFeedPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataFeedPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataFeedPostResponse proto.InternalMessageInfo

func (m *DataFeedPostResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECT
}

func (m *DataFeedPostResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *DataFeedPostResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

func init() {
	proto.RegisterType((*DataFeed)(nil), "west4API.DataFeed")
	proto.RegisterType((*DataFeedRequest)(nil), "west4API.DataFeedRequest")
	proto.RegisterType((*DataFeedResponse)(nil), "west4API.DataFeedResponse")
	proto.RegisterType((*DataFeedPostRequest)(nil), "west4API.DataFeedPostRequest")
	proto.RegisterType((*DataFeedPostResponse)(nil), "west4API.DataFeedPostResponse")
}

func init() { proto.RegisterFile("datafeeds.proto", fileDescriptor_6827cc0956531667) }

var fileDescriptor_6827cc0956531667 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xd1, 0x8a, 0xda, 0x40,
	0x14, 0x86, 0x49, 0x52, 0x4d, 0x3c, 0x96, 0xaa, 0x63, 0x0a, 0x69, 0x6f, 0x9a, 0xe6, 0x2a, 0x20,
	0x04, 0xaa, 0x7d, 0x01, 0x41, 0x0a, 0xb9, 0x93, 0x79, 0x81, 0x32, 0x66, 0x4e, 0x21, 0x54, 0x67,
	0xb2, 0x99, 0x71, 0x7d, 0x82, 0x7d, 0x87, 0xbd, 0xdf, 0x17, 0x5d, 0x9c, 0x38, 0x49, 0xdc, 0xc5,
	0x5d, 0xd8, 0xcb, 0x73, 0xfe, 0x3f, 0xdf, 0x39, 0xf3, 0x1f, 0x02, 0x13, 0xce, 0x34, 0xfb, 0x87,
	0xc8, 0x55, 0x56, 0xd5, 0x52, 0x4b, 0x12, 0x9c, 0x50, 0xe9, 0xdf, 0xeb, 0x6d, 0xfe, 0x1d, 0x50,
	0x1c, 0x0f, 0x4d, 0x37, 0x79, 0x74, 0x21, 0xd8, 0x30, 0xcd, 0xfe, 0x20, 0x72, 0xf2, 0x03, 0xc6,
	0xf6, 0xab, 0xbf, 0x25, 0x8f, 0x9c, 0xd8, 0x49, 0x07, 0x14, 0x6c, 0x2b, 0xe7, 0xe4, 0x17, 0x84,
	0xad, 0x81, 0xa3, 0x2a, 0xea, 0xb2, 0xd2, 0xa5, 0x14, 0x91, 0x1b, 0x3b, 0xe9, 0x88, 0xce, 0xad,
	0xb6, 0xe9, 0x24, 0xf2, 0x0d, 0x82, 0x7b, 0x14, 0x47, 0x3c, 0x03, 0x3d, 0x03, 0xf4, 0x4d, 0x9d,
	0x73, 0xb2, 0x80, 0x59, 0x23, 0xf5, 0x51, 0x9f, 0x0c, 0x6a, 0x6a, 0x84, 0x3e, 0xe7, 0x27, 0x7c,
	0x2e, 0xa4, 0xd0, 0xac, 0x14, 0x58, 0x9f, 0x59, 0x03, 0xc3, 0x1a, 0xb7, 0xbd, 0x9c, 0x93, 0x15,
	0x7c, 0xed, 0x2c, 0x7d, 0xe6, 0xd0, 0x30, 0xc3, 0x56, 0xec, 0x73, 0x23, 0xf0, 0x51, 0xb0, 0xdd,
	0x1e, 0x79, 0xe4, 0xc7, 0x4e, 0x1a, 0x50, 0x5b, 0x26, 0x05, 0x4c, 0x6c, 0x32, 0x14, 0xef, 0x8e,
	0xa8, 0x34, 0x09, 0x61, 0xa0, 0xe5, 0x7f, 0x14, 0x26, 0x9a, 0x11, 0x6d, 0x8a, 0x97, 0xb1, 0xb9,
	0xaf, 0x62, 0xbb, 0x9d, 0x41, 0xf2, 0xe4, 0xc0, 0xb4, 0x9b, 0xa2, 0x2a, 0x29, 0x14, 0x92, 0x05,
	0xf8, 0x35, 0xea, 0x42, 0x72, 0x34, 0x83, 0xbe, 0x2c, 0x67, 0x99, 0x3d, 0x5e, 0x46, 0x1b, 0x81,
	0x5a, 0x07, 0xc9, 0x20, 0xb0, 0xa3, 0x22, 0x37, 0xf6, 0xd2, 0xf1, 0x92, 0x74, 0xee, 0x16, 0xdd,
	0x7a, 0xce, 0x0f, 0x2e, 0xe4, 0xe1, 0x80, 0x42, 0x9b, 0x5d, 0x46, 0xd4, 0x96, 0x26, 0x8a, 0x3d,
	0xab, 0x14, 0xf2, 0xcb, 0x15, 0x6c, 0x99, 0x3c, 0x38, 0x30, 0xb7, 0xa8, 0xad, 0x54, 0xfa, 0xed,
	0x3c, 0xae, 0x37, 0x72, 0xde, 0xdd, 0x28, 0x85, 0x21, 0x2b, 0xcc, 0xa1, 0x3c, 0xf3, 0xda, 0x69,
	0xe7, 0x5e, 0x9b, 0x3e, 0xbd, 0xe8, 0xc9, 0x09, 0xc2, 0xeb, 0x35, 0x3e, 0x12, 0x58, 0x2f, 0x00,
	0xf7, 0x66, 0x00, 0xde, 0x55, 0x00, 0xbb, 0xa1, 0xf9, 0x5b, 0x56, 0xcf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x8a, 0x6f, 0xd5, 0xe3, 0x56, 0x03, 0x00, 0x00,
}
