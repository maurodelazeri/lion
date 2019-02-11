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
	DatafeedId           int64              `protobuf:"varint,1,opt,name=datafeed_id,json=datafeedId,proto3" json:"datafeed_id,omitempty"`
	DatafeedDescription  string             `protobuf:"bytes,2,opt,name=datafeed_description,json=datafeedDescription,proto3" json:"datafeed_description,omitempty"`
	VenueId              int64              `protobuf:"varint,3,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	VenueDescription     string             `protobuf:"bytes,4,opt,name=venue_description,json=venueDescription,proto3" json:"venue_description,omitempty"`
	ContainerId          int64              `protobuf:"varint,5,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ContainerDescription string             `protobuf:"bytes,6,opt,name=container_description,json=containerDescription,proto3" json:"container_description,omitempty"`
	ServerId             int64              `protobuf:"varint,7,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ServerDescription    string             `protobuf:"bytes,8,opt,name=server_description,json=serverDescription,proto3" json:"server_description,omitempty"`
	StreamingServer      string             `protobuf:"bytes,9,opt,name=streaming_server,json=streamingServer,proto3" json:"streaming_server,omitempty"`
	StreamingSecret      string             `protobuf:"bytes,10,opt,name=streaming_secret,json=streamingSecret,proto3" json:"streaming_secret,omitempty"`
	EventServer          string             `protobuf:"bytes,11,opt,name=event_server,json=eventServer,proto3" json:"event_server,omitempty"`
	Image                string             `protobuf:"bytes,12,opt,name=image,proto3" json:"image,omitempty"`
	MemoryReservation    float64            `protobuf:"fixed64,13,opt,name=memory_reservation,json=memoryReservation,proto3" json:"memory_reservation,omitempty"`
	MemoryLimit          float64            `protobuf:"fixed64,14,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	CpuLimit             float64            `protobuf:"fixed64,15,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`
	RestartPolicy        string             `protobuf:"bytes,16,opt,name=restart_policy,json=restartPolicy,proto3" json:"restart_policy,omitempty"`
	Products             []*DataFeedProduct `protobuf:"bytes,17,rep,name=products,proto3" json:"products,omitempty"`
	Env                  string             `protobuf:"bytes,18,opt,name=env,proto3" json:"env,omitempty"`
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

func (m *DataFeed) GetDatafeedId() int64 {
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

func (m *DataFeed) GetVenueId() int64 {
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

func (m *DataFeed) GetContainerId() int64 {
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

func (m *DataFeed) GetServerId() int64 {
	if m != nil {
		return m.ServerId
	}
	return 0
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

func (m *DataFeed) GetStreamingSecret() string {
	if m != nil {
		return m.StreamingSecret
	}
	return ""
}

func (m *DataFeed) GetEventServer() string {
	if m != nil {
		return m.EventServer
	}
	return ""
}

func (m *DataFeed) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *DataFeed) GetMemoryReservation() float64 {
	if m != nil {
		return m.MemoryReservation
	}
	return 0
}

func (m *DataFeed) GetMemoryLimit() float64 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *DataFeed) GetCpuLimit() float64 {
	if m != nil {
		return m.CpuLimit
	}
	return 0
}

func (m *DataFeed) GetRestartPolicy() string {
	if m != nil {
		return m.RestartPolicy
	}
	return ""
}

func (m *DataFeed) GetProducts() []*DataFeedProduct {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *DataFeed) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

type DataFeedProduct struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
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

func (m *DataFeedProduct) GetProductId() int64 {
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
	DatafeedId           int64    `protobuf:"varint,2,opt,name=datafeed_id,json=datafeedId,proto3" json:"datafeed_id,omitempty"`
	VenueId              int64    `protobuf:"varint,3,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
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

func (m *DataFeedRequest) GetDatafeedId() int64 {
	if m != nil {
		return m.DatafeedId
	}
	return 0
}

func (m *DataFeedRequest) GetVenueId() int64 {
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
	ContainerId          string   `protobuf:"bytes,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ContainerDescription string   `protobuf:"bytes,3,opt,name=container_description,json=containerDescription,proto3" json:"container_description,omitempty"`
	Comment              string   `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,5,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
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

func (m *DataFeedPostResponse) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *DataFeedPostResponse) GetContainerDescription() string {
	if m != nil {
		return m.ContainerDescription
	}
	return ""
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
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xe1, 0x6e, 0xd3, 0x3c,
	0x14, 0x55, 0xd6, 0x75, 0x4d, 0x6f, 0xb6, 0xb5, 0xf5, 0xfa, 0x49, 0xd9, 0x87, 0x10, 0x25, 0x12,
	0x52, 0xd1, 0x44, 0x11, 0x1b, 0x88, 0xdf, 0x93, 0x06, 0x52, 0x25, 0x7e, 0x4c, 0xe1, 0x01, 0x2a,
	0x13, 0x5f, 0xa6, 0x88, 0xc6, 0x09, 0x89, 0x53, 0xb4, 0x07, 0xe0, 0x49, 0x78, 0x09, 0xde, 0x82,
	0x57, 0x42, 0xbe, 0xb6, 0x93, 0xac, 0xd5, 0x06, 0xe2, 0x5f, 0xee, 0x39, 0xc7, 0xc7, 0x27, 0xf7,
	0xda, 0x86, 0x91, 0xe0, 0x8a, 0x7f, 0x46, 0x14, 0xd5, 0xa2, 0x28, 0x73, 0x95, 0x33, 0xff, 0x1b,
	0x56, 0xea, 0xf5, 0xe5, 0xf5, 0xf2, 0x7f, 0x40, 0x59, 0x67, 0x06, 0x8d, 0x7e, 0xf6, 0xc1, 0xbf,
	0xe2, 0x8a, 0xbf, 0x47, 0x14, 0xec, 0x09, 0x04, 0x6e, 0xd5, 0x2a, 0x15, 0xa1, 0x37, 0xf3, 0xe6,
	0xbd, 0x18, 0x1c, 0xb4, 0x14, 0xec, 0x15, 0x4c, 0x1b, 0x81, 0xc0, 0x2a, 0x29, 0xd3, 0x42, 0xa5,
	0xb9, 0x0c, 0xf7, 0x66, 0xde, 0x7c, 0x18, 0x9f, 0x38, 0xee, 0xaa, 0xa5, 0xd8, 0x29, 0xf8, 0x1b,
	0x94, 0x35, 0x6a, 0xc3, 0x1e, 0x19, 0x0e, 0xa8, 0x5e, 0x0a, 0x76, 0x06, 0x13, 0x43, 0x75, 0xad,
	0xf6, 0xc9, 0x6a, 0x4c, 0x44, 0xd7, 0xe7, 0x29, 0x1c, 0x26, 0xb9, 0x54, 0x3c, 0x95, 0x58, 0x6a,
	0xaf, 0x3e, 0x79, 0x05, 0x0d, 0xb6, 0x14, 0xec, 0x02, 0xfe, 0x6b, 0x25, 0x5d, 0xcf, 0x03, 0xf2,
	0x9c, 0x36, 0x64, 0xd7, 0xf7, 0x11, 0x0c, 0x2b, 0x2c, 0x37, 0xc6, 0x74, 0x40, 0xa6, 0xbe, 0x01,
	0x96, 0x82, 0xbd, 0x00, 0x66, 0xc9, 0xae, 0x9d, 0x4f, 0x76, 0x13, 0xc3, 0x74, 0xbd, 0x9e, 0xc3,
	0xb8, 0x52, 0x25, 0xf2, 0x2c, 0x95, 0x37, 0x2b, 0x43, 0x87, 0x43, 0x12, 0x8f, 0x1a, 0xfc, 0x23,
	0xc1, 0xdb, 0xd2, 0xa4, 0x44, 0x15, 0xc2, 0x8e, 0x54, 0xc3, 0xfa, 0xcf, 0x71, 0x83, 0x52, 0x39,
	0xc7, 0x80, 0x64, 0x01, 0x61, 0xd6, 0x6d, 0x0a, 0xfd, 0x34, 0xe3, 0x37, 0x18, 0x1e, 0x12, 0x67,
	0x0a, 0x9d, 0x3e, 0xc3, 0x2c, 0x2f, 0x6f, 0x57, 0x25, 0xea, 0xb5, 0x9c, 0xd2, 0x1f, 0xcd, 0xbc,
	0xb9, 0x17, 0x4f, 0x0c, 0x13, 0xb7, 0x84, 0xde, 0xc7, 0xca, 0xd7, 0x69, 0x96, 0xaa, 0xf0, 0x98,
	0x84, 0x81, 0xc1, 0x3e, 0x68, 0x48, 0x37, 0x2b, 0x29, 0x6a, 0xcb, 0x8f, 0x88, 0xf7, 0x93, 0xa2,
	0x36, 0xe4, 0x33, 0x38, 0x2e, 0xb1, 0x52, 0xbc, 0x54, 0xab, 0x22, 0x5f, 0xa7, 0xc9, 0x6d, 0x38,
	0xa6, 0x34, 0x47, 0x16, 0xbd, 0x26, 0x90, 0xbd, 0x01, 0xbf, 0x28, 0x73, 0x51, 0x27, 0xaa, 0x0a,
	0x27, 0xb3, 0xde, 0x3c, 0x38, 0x3f, 0x5d, 0xb8, 0xa3, 0xb9, 0x70, 0x47, 0xf1, 0xda, 0x28, 0xe2,
	0x46, 0xca, 0xc6, 0xd0, 0x43, 0xb9, 0x09, 0x19, 0x59, 0xea, 0xcf, 0x88, 0xc3, 0x68, 0x4b, 0xce,
	0x1e, 0x03, 0xd8, 0x05, 0xed, 0xf9, 0x1d, 0x5a, 0x64, 0x29, 0xd8, 0x4b, 0x38, 0x71, 0xf4, 0xee,
	0xe9, 0x65, 0x96, 0xea, 0x0c, 0x34, 0x7a, 0x0b, 0x81, 0xdb, 0xe2, 0x9d, 0xdc, 0x30, 0x06, 0xfb,
	0x92, 0x67, 0x48, 0xc6, 0xc3, 0x98, 0xbe, 0x75, 0xeb, 0x37, 0x7c, 0x5d, 0xa3, 0x75, 0x31, 0x45,
	0x94, 0xb4, 0xd9, 0x62, 0xfc, 0x5a, 0x63, 0xa5, 0xb4, 0x50, 0xe5, 0x5f, 0x50, 0xda, 0xd5, 0xa6,
	0xd8, 0xbe, 0x72, 0x7b, 0x3b, 0x57, 0xee, 0xfe, 0xfb, 0x13, 0xfd, 0xf0, 0x60, 0xdc, 0xee, 0x52,
	0x15, 0xb9, 0xac, 0x90, 0x9d, 0xc1, 0xa0, 0x44, 0x95, 0xe4, 0xc2, 0xc4, 0x3c, 0x3e, 0x9f, 0xb4,
	0xdd, 0x8d, 0x0d, 0x11, 0x3b, 0x05, 0x5b, 0x80, 0xef, 0xb6, 0x0a, 0xf7, 0x68, 0x16, 0x6c, 0x77,
	0x16, 0x71, 0xa3, 0x61, 0x21, 0x0c, 0x92, 0x3c, 0xcb, 0x50, 0x2a, 0xca, 0x32, 0x8c, 0x5d, 0xa9,
	0x19, 0x5c, 0xf3, 0xa2, 0x42, 0x61, 0x6f, 0xb0, 0x2b, 0xa3, 0xef, 0x1e, 0x9c, 0x34, 0x73, 0xca,
	0x2b, 0xf5, 0x70, 0x3f, 0xee, 0x26, 0xf2, 0xfe, 0x98, 0x68, 0x0e, 0x07, 0x3c, 0xa1, 0x29, 0xf6,
	0xe8, 0x6f, 0xc7, 0xad, 0xfa, 0x92, 0xf0, 0xd8, 0xf2, 0xd1, 0x2f, 0x0f, 0xa6, 0x77, 0x73, 0xfc,
	0x4b, 0xc7, 0xb6, 0x9f, 0x21, 0x33, 0xf5, 0xbf, 0x7b, 0x86, 0x7a, 0x0f, 0x3c, 0x43, 0x9d, 0xce,
	0xee, 0xdf, 0xdb, 0xd9, 0xfe, 0x9d, 0xce, 0x7e, 0x3a, 0xa0, 0x27, 0xfc, 0xe2, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc6, 0xde, 0x2c, 0x99, 0xeb, 0x05, 0x00, 0x00,
}
