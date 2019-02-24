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
	ContainerEvent       string             `protobuf:"bytes,6,opt,name=container_event,json=containerEvent,proto3" json:"container_event,omitempty"`
	ContainerIdentifier  string             `protobuf:"bytes,7,opt,name=container_identifier,json=containerIdentifier,proto3" json:"container_identifier,omitempty"`
	ServerId             int64              `protobuf:"varint,8,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ServerDescription    string             `protobuf:"bytes,9,opt,name=server_description,json=serverDescription,proto3" json:"server_description,omitempty"`
	StreamingServer      string             `protobuf:"bytes,10,opt,name=streaming_server,json=streamingServer,proto3" json:"streaming_server,omitempty"`
	StreamingSecret      string             `protobuf:"bytes,11,opt,name=streaming_secret,json=streamingSecret,proto3" json:"streaming_secret,omitempty"`
	EventServer          string             `protobuf:"bytes,12,opt,name=event_server,json=eventServer,proto3" json:"event_server,omitempty"`
	Image                string             `protobuf:"bytes,13,opt,name=image,proto3" json:"image,omitempty"`
	MemoryReservation    float64            `protobuf:"fixed64,14,opt,name=memory_reservation,json=memoryReservation,proto3" json:"memory_reservation,omitempty"`
	MemoryLimit          float64            `protobuf:"fixed64,15,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	CpuLimit             float64            `protobuf:"fixed64,16,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`
	RestartPolicy        string             `protobuf:"bytes,17,opt,name=restart_policy,json=restartPolicy,proto3" json:"restart_policy,omitempty"`
	Products             []*DataFeedProduct `protobuf:"bytes,18,rep,name=products,proto3" json:"products,omitempty"`
	Env                  string             `protobuf:"bytes,19,opt,name=env,proto3" json:"env,omitempty"`
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

func (m *DataFeed) GetContainerEvent() string {
	if m != nil {
		return m.ContainerEvent
	}
	return ""
}

func (m *DataFeed) GetContainerIdentifier() string {
	if m != nil {
		return m.ContainerIdentifier
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
	return Retcode_REQUEST_WITH_NO_TOKEN
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
	ContainerIdentifier  string   `protobuf:"bytes,2,opt,name=container_identifier,json=containerIdentifier,proto3" json:"container_identifier,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
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
	return Retcode_REQUEST_WITH_NO_TOKEN
}

func (m *DataFeedPostResponse) GetContainerIdentifier() string {
	if m != nil {
		return m.ContainerIdentifier
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
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x95, 0x93, 0xb6, 0x71, 0xc6, 0x6d, 0x7e, 0x36, 0xbd, 0x70, 0xbf, 0x4f, 0x88, 0x60, 0x09,
	0x11, 0x54, 0x11, 0x44, 0x01, 0x71, 0x5d, 0xa9, 0x45, 0x8a, 0xc4, 0x45, 0x65, 0x1e, 0x20, 0x5a,
	0xbc, 0xd3, 0xca, 0x22, 0x5e, 0x9b, 0xf5, 0x3a, 0xa8, 0x0f, 0xc0, 0x93, 0x70, 0xc7, 0x4b, 0xf1,
	0x2a, 0x68, 0x67, 0xbd, 0xb6, 0xdb, 0xa8, 0x20, 0xf5, 0xce, 0x73, 0xce, 0xd9, 0x33, 0x3f, 0xbb,
	0x63, 0x18, 0x0b, 0xae, 0xf9, 0x35, 0xa2, 0x28, 0x97, 0x85, 0xca, 0x75, 0xce, 0xfc, 0xef, 0x58,
	0xea, 0x77, 0xe7, 0x57, 0xab, 0xff, 0x00, 0x65, 0x95, 0x59, 0x34, 0xfa, 0xbd, 0x0f, 0xfe, 0x05,
	0xd7, 0xfc, 0x23, 0xa2, 0x60, 0x4f, 0x21, 0x70, 0xa7, 0xd6, 0xa9, 0x08, 0xbd, 0xb9, 0xb7, 0xe8,
	0xc7, 0xe0, 0xa0, 0x95, 0x60, 0x6f, 0xe0, 0xb8, 0x11, 0x08, 0x2c, 0x13, 0x95, 0x16, 0x3a, 0xcd,
	0x65, 0xd8, 0x9b, 0x7b, 0x8b, 0x61, 0x3c, 0x73, 0xdc, 0x45, 0x4b, 0xb1, 0x13, 0xf0, 0xb7, 0x28,
	0x2b, 0x34, 0x86, 0x7d, 0x32, 0x1c, 0x50, 0xbc, 0x12, 0xec, 0x14, 0xa6, 0x96, 0xea, 0x5a, 0xed,
	0x91, 0xd5, 0x84, 0x88, 0xae, 0xcf, 0x33, 0x38, 0x4c, 0x72, 0xa9, 0x79, 0x2a, 0x51, 0x19, 0xaf,
	0x7d, 0xf2, 0x0a, 0x1a, 0x6c, 0x25, 0xd8, 0x0b, 0x18, 0xb7, 0x12, 0xdc, 0xa2, 0xd4, 0xe1, 0x01,
	0xb9, 0x8d, 0x1a, 0xf8, 0xd2, 0xa0, 0xa6, 0x8d, 0xae, 0x17, 0x4a, 0x9d, 0x5e, 0xa7, 0xa8, 0xc2,
	0x81, 0x6d, 0xa3, 0xe3, 0xe9, 0x28, 0xf6, 0x3f, 0x0c, 0x4b, 0x54, 0x5b, 0x9b, 0xdb, 0xa7, 0xdc,
	0xbe, 0x05, 0x56, 0x82, 0xbd, 0x02, 0x56, 0x93, 0xdd, 0x4e, 0x86, 0xe4, 0x36, 0xb5, 0x4c, 0xb7,
	0x95, 0x97, 0x30, 0x29, 0xb5, 0x42, 0x9e, 0xa5, 0xf2, 0x66, 0x6d, 0xe9, 0x10, 0x48, 0x3c, 0x6e,
	0xf0, 0xcf, 0x04, 0xdf, 0x97, 0x26, 0x0a, 0x75, 0x18, 0xec, 0x48, 0x0d, 0x6c, 0x06, 0x44, 0x3d,
	0x3b, 0xc7, 0x43, 0x92, 0x05, 0x84, 0xd5, 0x6e, 0xc7, 0xb0, 0x9f, 0x66, 0xfc, 0x06, 0xc3, 0x23,
	0xe2, 0x6c, 0x60, 0xaa, 0xcf, 0x30, 0xcb, 0xd5, 0xed, 0x5a, 0xa1, 0x39, 0xcb, 0xa9, 0xfa, 0xd1,
	0xdc, 0x5b, 0x78, 0xf1, 0xd4, 0x32, 0x71, 0x4b, 0x98, 0x3c, 0xb5, 0x7c, 0x93, 0x66, 0xa9, 0x0e,
	0xc7, 0x24, 0x0c, 0x2c, 0xf6, 0xc9, 0x40, 0x66, 0x58, 0x49, 0x51, 0xd5, 0xfc, 0x84, 0x78, 0x3f,
	0x29, 0x2a, 0x4b, 0x3e, 0x87, 0x91, 0xc2, 0x52, 0x73, 0xa5, 0xd7, 0x45, 0xbe, 0x49, 0x93, 0xdb,
	0x70, 0x4a, 0xd5, 0x1c, 0xd5, 0xe8, 0x15, 0x81, 0xec, 0x3d, 0xf8, 0x85, 0xca, 0x45, 0x95, 0xe8,
	0x32, 0x64, 0xf3, 0xfe, 0x22, 0x38, 0x3b, 0x59, 0xba, 0x17, 0xbc, 0x74, 0x2f, 0xf6, 0xca, 0x2a,
	0xe2, 0x46, 0xca, 0x26, 0xd0, 0x47, 0xb9, 0x0d, 0x67, 0x64, 0x69, 0x3e, 0x23, 0x0e, 0xe3, 0x7b,
	0x72, 0xf6, 0x04, 0xa0, 0x3e, 0xd0, 0x3e, 0xf3, 0x61, 0x8d, 0xac, 0x04, 0x7b, 0x0d, 0x33, 0x47,
	0xef, 0x3e, 0x72, 0x56, 0x53, 0x9d, 0x0b, 0x8d, 0x3e, 0x40, 0xe0, 0x52, 0x5c, 0xca, 0x2d, 0x63,
	0xb0, 0x27, 0x79, 0x86, 0x64, 0x3c, 0x8c, 0xe9, 0xdb, 0x8c, 0x7e, 0xcb, 0x37, 0x15, 0xd6, 0x2e,
	0x36, 0x88, 0x92, 0xb6, 0xb6, 0x18, 0xbf, 0x55, 0x58, 0x6a, 0x23, 0xd4, 0xf9, 0x57, 0x94, 0xf5,
	0x69, 0x1b, 0xdc, 0xdf, 0xcc, 0xde, 0xce, 0x66, 0x3e, 0xbc, 0x66, 0xd1, 0x4f, 0x0f, 0x26, 0x6d,
	0x96, 0xb2, 0xc8, 0x65, 0x89, 0xec, 0x14, 0x06, 0x0a, 0x75, 0x92, 0x0b, 0x5b, 0xe6, 0xe8, 0x6c,
	0xda, 0x4e, 0x37, 0xb6, 0x44, 0xec, 0x14, 0x6c, 0x09, 0xbe, 0x4b, 0x15, 0xf6, 0xe8, 0x2e, 0xd8,
	0xee, 0x5d, 0xc4, 0x8d, 0x86, 0x85, 0x30, 0x48, 0xf2, 0x2c, 0x33, 0x0b, 0xd8, 0xa7, 0x2e, 0x5c,
	0x68, 0x18, 0xdc, 0xf0, 0xa2, 0x44, 0x51, 0x2f, 0xba, 0x0b, 0xa3, 0x1f, 0x1e, 0xcc, 0x9a, 0x7b,
	0xca, 0x4b, 0xfd, 0xf7, 0x79, 0xdc, 0xad, 0xc8, 0xfb, 0x67, 0x45, 0x0b, 0x38, 0xe0, 0x09, 0xdd,
	0x62, 0x9f, 0xba, 0x9d, 0xb4, 0xea, 0x73, 0xc2, 0xe3, 0x9a, 0x8f, 0x7e, 0x79, 0x70, 0x7c, 0xb7,
	0x8e, 0xc7, 0x4c, 0xec, 0xa1, 0x3f, 0x4c, 0xef, 0xe1, 0x3f, 0xcc, 0x23, 0x86, 0xf6, 0xe5, 0x80,
	0x7e, 0xe2, 0x6f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x98, 0x8d, 0xdd, 0xed, 0x05, 0x00,
	0x00,
}
