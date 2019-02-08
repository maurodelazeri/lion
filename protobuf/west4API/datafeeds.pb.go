// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datafeeds.proto

package west4API

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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
	DatafeedId           int32              `protobuf:"varint,1,opt,name=datafeed_id,json=datafeedId,proto3" json:"datafeed_id,omitempty"`
	DatafeedDescription  string             `protobuf:"bytes,2,opt,name=datafeed_description,json=datafeedDescription,proto3" json:"datafeed_description,omitempty"`
	VenueId              int32              `protobuf:"varint,3,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	VenueDescription     string             `protobuf:"bytes,4,opt,name=venue_description,json=venueDescription,proto3" json:"venue_description,omitempty"`
	ContainerId          int32              `protobuf:"varint,5,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ServerId             string             `protobuf:"bytes,6,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ServerDescription    string             `protobuf:"bytes,7,opt,name=server_description,json=serverDescription,proto3" json:"server_description,omitempty"`
	StreamingServer      string             `protobuf:"bytes,8,opt,name=streaming_server,json=streamingServer,proto3" json:"streaming_server,omitempty"`
	EventServer          string             `protobuf:"bytes,9,opt,name=event_server,json=eventServer,proto3" json:"event_server,omitempty"`
	MemoryReservation    int64              `protobuf:"varint,10,opt,name=memory_reservation,json=memoryReservation,proto3" json:"memory_reservation,omitempty"`
	MemoryLimit          int64              `protobuf:"varint,11,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	Cpu                  int32              `protobuf:"varint,12,opt,name=cpu,proto3" json:"cpu,omitempty"`
	RestartPolicy        int32              `protobuf:"varint,13,opt,name=restart_policy,json=restartPolicy,proto3" json:"restart_policy,omitempty"`
	Products             []*DataFeedProduct `protobuf:"bytes,14,rep,name=products,proto3" json:"products,omitempty"`
	Env                  []*DataFeedEnv     `protobuf:"bytes,15,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
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

func (m *DataFeed) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *DataFeed) GetServerDescription() string {
	if m != nil {
		return m.ServerDescription
	}
	return ""
}

func (m *DataFeed) GetStreamingServer() string {
	if m != nil {
		return m.StreamingServer
	}
	return ""
}

func (m *DataFeed) GetEventServer() string {
	if m != nil {
		return m.EventServer
	}
	return ""
}

func (m *DataFeed) GetMemoryReservation() int64 {
	if m != nil {
		return m.MemoryReservation
	}
	return 0
}

func (m *DataFeed) GetMemoryLimit() int64 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *DataFeed) GetCpu() int32 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *DataFeed) GetRestartPolicy() int32 {
	if m != nil {
		return m.RestartPolicy
	}
	return 0
}

func (m *DataFeed) GetProducts() []*DataFeedProduct {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *DataFeed) GetEnv() []*DataFeedEnv {
	if m != nil {
		return m.Env
	}
	return nil
}

type DataFeedProduct struct {
	ProductId            int32    `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductDescription   string   `protobuf:"bytes,2,opt,name=product_description,json=productDescription,proto3" json:"product_description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataFeedProduct) Reset()         { *m = DataFeedProduct{} }
func (m *DataFeedProduct) String() string { return proto.CompactTextString(m) }
func (*DataFeedProduct) ProtoMessage()    {}
func (*DataFeedProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_6827cc0956531667, []int{1}
}

func (m *DataFeedProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataFeedProduct.Unmarshal(m, b)
}
func (m *DataFeedProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataFeedProduct.Marshal(b, m, deterministic)
}
func (m *DataFeedProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataFeedProduct.Merge(m, src)
}
func (m *DataFeedProduct) XXX_Size() int {
	return xxx_messageInfo_DataFeedProduct.Size(m)
}
func (m *DataFeedProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_DataFeedProduct.DiscardUnknown(m)
}

var xxx_messageInfo_DataFeedProduct proto.InternalMessageInfo

func (m *DataFeedProduct) GetProductId() int32 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *DataFeedProduct) GetProductDescription() string {
	if m != nil {
		return m.ProductDescription
	}
	return ""
}

type DataFeedEnv struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataFeedEnv) Reset()         { *m = DataFeedEnv{} }
func (m *DataFeedEnv) String() string { return proto.CompactTextString(m) }
func (*DataFeedEnv) ProtoMessage()    {}
func (*DataFeedEnv) Descriptor() ([]byte, []int) {
	return fileDescriptor_6827cc0956531667, []int{2}
}

func (m *DataFeedEnv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataFeedEnv.Unmarshal(m, b)
}
func (m *DataFeedEnv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataFeedEnv.Marshal(b, m, deterministic)
}
func (m *DataFeedEnv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataFeedEnv.Merge(m, src)
}
func (m *DataFeedEnv) XXX_Size() int {
	return xxx_messageInfo_DataFeedEnv.Size(m)
}
func (m *DataFeedEnv) XXX_DiscardUnknown() {
	xxx_messageInfo_DataFeedEnv.DiscardUnknown(m)
}

var xxx_messageInfo_DataFeedEnv proto.InternalMessageInfo

func (m *DataFeedEnv) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataFeedEnv) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
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
	return fileDescriptor_6827cc0956531667, []int{3}
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
	return fileDescriptor_6827cc0956531667, []int{4}
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
	return fileDescriptor_6827cc0956531667, []int{5}
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
	return fileDescriptor_6827cc0956531667, []int{6}
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
	proto.RegisterType((*DataFeedProduct)(nil), "west4API.DataFeedProduct")
	proto.RegisterType((*DataFeedEnv)(nil), "west4API.DataFeedEnv")
	proto.RegisterType((*DataFeedRequest)(nil), "west4API.DataFeedRequest")
	proto.RegisterType((*DataFeedResponse)(nil), "west4API.DataFeedResponse")
	proto.RegisterType((*DataFeedPostRequest)(nil), "west4API.DataFeedPostRequest")
	proto.RegisterType((*DataFeedPostResponse)(nil), "west4API.DataFeedPostResponse")
}

func init() { proto.RegisterFile("datafeeds.proto", fileDescriptor_6827cc0956531667) }

var fileDescriptor_6827cc0956531667 = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x95, 0xe3, 0xb6, 0x71, 0xc6, 0x6d, 0x93, 0x6c, 0xfb, 0x49, 0xee, 0x87, 0x10, 0xc1, 0x12,
	0x22, 0xa8, 0x6a, 0x10, 0x05, 0xc4, 0x75, 0xa5, 0x82, 0x14, 0x89, 0x8b, 0x6a, 0x79, 0x80, 0x68,
	0xf1, 0x0e, 0xc8, 0x22, 0xde, 0x35, 0xde, 0x8d, 0xab, 0x3e, 0x00, 0x4f, 0xc2, 0xfb, 0xf0, 0x4c,
	0xc8, 0xb3, 0xfe, 0x4b, 0xa3, 0x82, 0xc4, 0x9d, 0xe7, 0x9c, 0xb3, 0x67, 0x26, 0x67, 0x67, 0x03,
	0x63, 0x29, 0xac, 0xf8, 0x82, 0x28, 0xcd, 0x22, 0x2f, 0xb4, 0xd5, 0x2c, 0xb8, 0x45, 0x63, 0xdf,
	0x5c, 0xdd, 0x2c, 0xff, 0x07, 0x54, 0x9b, 0xcc, 0xa1, 0xf1, 0xaf, 0x3d, 0x08, 0xae, 0x85, 0x15,
	0x1f, 0x10, 0x25, 0x7b, 0x02, 0x61, 0x73, 0x6a, 0x95, 0xca, 0xc8, 0x9b, 0x79, 0xf3, 0x7d, 0x0e,
	0x0d, 0xb4, 0x94, 0xec, 0x15, 0x9c, 0xb6, 0x02, 0x89, 0x26, 0x29, 0xd2, 0xdc, 0xa6, 0x5a, 0x45,
	0x83, 0x99, 0x37, 0x1f, 0xf1, 0x93, 0x86, 0xbb, 0xee, 0x28, 0x76, 0x06, 0x41, 0x89, 0x6a, 0x83,
	0x95, 0xa1, 0x4f, 0x86, 0x43, 0xaa, 0x97, 0x92, 0x9d, 0xc3, 0xd4, 0x51, 0x7d, 0xab, 0x3d, 0xb2,
	0x9a, 0x10, 0xd1, 0xf7, 0x79, 0x0a, 0x87, 0x89, 0x56, 0x56, 0xa4, 0x0a, 0x8b, 0xca, 0x6b, 0x9f,
	0xbc, 0xc2, 0x16, 0x5b, 0x4a, 0xf6, 0x08, 0x46, 0x06, 0x8b, 0xd2, 0xf1, 0x07, 0xe4, 0x13, 0x38,
	0x60, 0x29, 0xd9, 0x05, 0xb0, 0x9a, 0xec, 0x77, 0x1b, 0x92, 0x6a, 0xea, 0x98, 0x7e, 0xbb, 0x17,
	0x30, 0x31, 0xb6, 0x40, 0x91, 0xa5, 0xea, 0xeb, 0xca, 0xd1, 0x51, 0x40, 0xe2, 0x71, 0x8b, 0x7f,
	0x22, 0xb8, 0x9a, 0x0c, 0x4b, 0x54, 0xb6, 0x91, 0x8d, 0x48, 0x16, 0x12, 0x56, 0x4b, 0x2e, 0x80,
	0x65, 0x98, 0xe9, 0xe2, 0x6e, 0x55, 0x60, 0xa5, 0x12, 0xd4, 0x1c, 0x66, 0xde, 0xdc, 0xe7, 0x53,
	0xc7, 0xf0, 0x8e, 0xa8, 0x1c, 0x6b, 0xf9, 0x3a, 0xcd, 0x52, 0x1b, 0x85, 0x24, 0x0c, 0x1d, 0xf6,
	0xb1, 0x82, 0xd8, 0x04, 0xfc, 0x24, 0xdf, 0x44, 0x87, 0x94, 0x42, 0xf5, 0xc9, 0x9e, 0xc1, 0x71,
	0x81, 0xc6, 0x8a, 0xc2, 0xae, 0x72, 0xbd, 0x4e, 0x93, 0xbb, 0xe8, 0x88, 0xc8, 0xa3, 0x1a, 0xbd,
	0x21, 0x90, 0xbd, 0x85, 0x20, 0x2f, 0xb4, 0xdc, 0x24, 0xd6, 0x44, 0xc7, 0x33, 0x7f, 0x1e, 0x5e,
	0x9e, 0x2d, 0x9a, 0xcd, 0x58, 0x34, 0x9b, 0x70, 0xe3, 0x14, 0xbc, 0x95, 0xb2, 0xe7, 0xe0, 0xa3,
	0x2a, 0xa3, 0x31, 0x9d, 0xf8, 0x6f, 0xf7, 0xc4, 0x7b, 0x55, 0xf2, 0x4a, 0x11, 0x0b, 0x18, 0xdf,
	0x73, 0x61, 0x8f, 0x01, 0x6a, 0x9f, 0x6e, 0xab, 0x46, 0x35, 0xb2, 0x94, 0xec, 0x25, 0x9c, 0x34,
	0xf4, 0xee, 0x4e, 0xb1, 0x9a, 0xea, 0xdd, 0x4d, 0xfc, 0x0e, 0xc2, 0x5e, 0x5b, 0xc6, 0x60, 0x4f,
	0x89, 0x0c, 0xc9, 0x78, 0xc4, 0xe9, 0x9b, 0x9d, 0xc2, 0x7e, 0x29, 0xd6, 0x1b, 0xac, 0x5d, 0x5c,
	0x11, 0x27, 0xdd, 0x6c, 0x1c, 0xbf, 0x6f, 0xd0, 0xd8, 0x4a, 0x68, 0xf5, 0x37, 0x54, 0xf5, 0x69,
	0x57, 0xdc, 0x7f, 0x08, 0x83, 0x9d, 0x87, 0xf0, 0xf0, 0x56, 0xc7, 0x3f, 0x3d, 0x98, 0x74, 0x5d,
	0x4c, 0xae, 0x95, 0x41, 0x76, 0x0e, 0xc3, 0x02, 0x6d, 0xa2, 0xa5, 0x1b, 0xf3, 0xf8, 0x72, 0xda,
	0x45, 0xc8, 0x1d, 0xc1, 0x1b, 0x05, 0x5b, 0x40, 0xd0, 0xb4, 0x8a, 0x06, 0x14, 0x38, 0xdb, 0x0d,
	0x9c, 0xb7, 0x1a, 0x16, 0xc1, 0x30, 0xd1, 0x59, 0x86, 0xca, 0xd2, 0x2c, 0x23, 0xde, 0x94, 0x15,
	0x83, 0x6b, 0x91, 0x1b, 0x94, 0xf5, 0xbb, 0x6a, 0xca, 0xf8, 0x87, 0x07, 0x27, 0xed, 0x3d, 0x69,
	0x63, 0xff, 0x9c, 0xc7, 0xf6, 0x44, 0xde, 0x5f, 0x27, 0x9a, 0xc3, 0x81, 0x48, 0xe8, 0x16, 0x7d,
	0xfa, 0xb5, 0x93, 0x4e, 0x7d, 0x45, 0x38, 0xaf, 0xf9, 0xf8, 0x16, 0x4e, 0xb7, 0xc7, 0xf8, 0x97,
	0xc0, 0x7a, 0x01, 0x0c, 0x1e, 0x0c, 0xc0, 0xdf, 0x0a, 0xe0, 0xf3, 0x01, 0xfd, 0xff, 0xbd, 0xfe,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xb7, 0x6a, 0x17, 0x28, 0x05, 0x00, 0x00,
}
