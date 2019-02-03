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
	VenueId              int32    `protobuf:"varint,2,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	ContainerId          int32    `protobuf:"varint,3,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	DatafeedDescription  string   `protobuf:"bytes,4,opt,name=datafeed_description,json=datafeedDescription,proto3" json:"datafeed_description,omitempty"`
	Enabled              bool     `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
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

func (m *DataFeed) GetVenueId() int32 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *DataFeed) GetContainerId() int32 {
	if m != nil {
		return m.ContainerId
	}
	return 0
}

func (m *DataFeed) GetDatafeedDescription() string {
	if m != nil {
		return m.DatafeedDescription
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
	ContainerId          int32    `protobuf:"varint,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
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

func (m *DataFeedRequest) GetContainerId() int32 {
	if m != nil {
		return m.ContainerId
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
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x49, 0xeb, 0xb6, 0xee, 0x4c, 0xdc, 0xcc, 0x76, 0x51, 0xbd, 0xb1, 0xf6, 0xaa, 0x20,
	0x14, 0x9c, 0xbe, 0xc0, 0x60, 0x08, 0xf5, 0x6a, 0xe4, 0x05, 0x24, 0x6b, 0x8e, 0x50, 0xdc, 0x92,
	0xda, 0x64, 0xee, 0x09, 0x7c, 0x12, 0xdf, 0xc1, 0xe7, 0x93, 0x65, 0x4b, 0x6b, 0x19, 0x2a, 0x78,
	0xf9, 0x9f, 0xff, 0xf4, 0xe7, 0x3b, 0x7f, 0x03, 0x43, 0xc1, 0x0d, 0x7f, 0x46, 0x14, 0x3a, 0x2d,
	0x2b, 0x65, 0x14, 0x0d, 0xb6, 0xa8, 0xcd, 0xfd, 0x6c, 0x91, 0x5d, 0x02, 0xca, 0xcd, 0x7a, 0x3f,
	0x8d, 0x3f, 0x09, 0x04, 0x73, 0x6e, 0xf8, 0x03, 0xa2, 0xa0, 0x57, 0x30, 0x70, 0x5f, 0x3d, 0x15,
	0x22, 0x24, 0x11, 0x49, 0x3a, 0x0c, 0xdc, 0x28, 0x13, 0xf4, 0x02, 0x82, 0x37, 0x94, 0x1b, 0xdc,
	0xb9, 0x9e, 0x75, 0x7b, 0x56, 0x67, 0x82, 0x5e, 0xc3, 0x69, 0xae, 0xa4, 0xe1, 0x85, 0xc4, 0x6a,
	0x67, 0xfb, 0xd6, 0x1e, 0xd4, 0xb3, 0x4c, 0xd0, 0x5b, 0x98, 0xd4, 0xf1, 0x02, 0x75, 0x5e, 0x15,
	0xa5, 0x29, 0x94, 0x0c, 0x4f, 0x22, 0x92, 0xf4, 0xd9, 0xd8, 0x79, 0xf3, 0xc6, 0xa2, 0x21, 0xf4,
	0x50, 0xf2, 0xe5, 0x0a, 0x45, 0xd8, 0x89, 0x48, 0x12, 0x30, 0x27, 0xe3, 0x47, 0x18, 0x3a, 0x6e,
	0x86, 0xaf, 0x1b, 0xd4, 0x86, 0x4e, 0xa0, 0x63, 0xd4, 0x0b, 0x4a, 0x0b, 0xde, 0x67, 0x7b, 0x71,
	0x04, 0xe6, 0x1d, 0x81, 0xc5, 0x1f, 0x04, 0x46, 0x4d, 0x98, 0x2e, 0x95, 0xd4, 0x48, 0x6f, 0xa0,
	0x57, 0xa1, 0xc9, 0x95, 0x40, 0x9b, 0x77, 0x36, 0x3d, 0x4f, 0x5d, 0x83, 0x29, 0xdb, 0x1b, 0xcc,
	0x6d, 0xd0, 0x14, 0x02, 0x87, 0x1f, 0x7a, 0x91, 0x9f, 0x0c, 0xa6, 0xb4, 0xd9, 0xae, 0xa3, 0xeb,
	0x9d, 0xdd, 0x5d, 0xb9, 0x5a, 0xaf, 0x51, 0x1a, 0x5b, 0x54, 0x9f, 0x39, 0x69, 0x2f, 0x5e, 0xf1,
	0x52, 0xa3, 0x38, 0xf4, 0xe2, 0x64, 0xfc, 0x4e, 0x60, 0xec, 0xa2, 0x16, 0x4a, 0x9b, 0xdf, 0xcf,
	0x6e, 0x13, 0x91, 0x3f, 0x89, 0x12, 0xe8, 0xf2, 0xdc, 0xfe, 0x0e, 0xdf, 0x5e, 0x3b, 0x6a, 0xb6,
	0x67, 0x76, 0xce, 0x0e, 0x7e, 0xbc, 0x85, 0x49, 0x1b, 0xe3, 0x3f, 0x85, 0x7d, 0x2b, 0xc0, 0xfb,
	0xb1, 0x00, 0xbf, 0x55, 0xc0, 0xb2, 0x6b, 0x9f, 0xec, 0xdd, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x80, 0x89, 0xc4, 0xae, 0xdb, 0x02, 0x00, 0x00,
}
