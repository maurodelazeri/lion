// Code generated by protoc-gen-go. DO NOT EDIT.
// source: webapi.proto

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

// Venues
type VenuesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenuesRequest) Reset()         { *m = VenuesRequest{} }
func (m *VenuesRequest) String() string { return proto.CompactTextString(m) }
func (*VenuesRequest) ProtoMessage()    {}
func (*VenuesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f696248df6cc389, []int{0}
}

func (m *VenuesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenuesRequest.Unmarshal(m, b)
}
func (m *VenuesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenuesRequest.Marshal(b, m, deterministic)
}
func (m *VenuesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuesRequest.Merge(m, src)
}
func (m *VenuesRequest) XXX_Size() int {
	return xxx_messageInfo_VenuesRequest.Size(m)
}
func (m *VenuesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VenuesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VenuesRequest proto.InternalMessageInfo

type VenuesInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	VenueId              int32    `protobuf:"varint,3,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	Enabled              bool     `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenuesInfo) Reset()         { *m = VenuesInfo{} }
func (m *VenuesInfo) String() string { return proto.CompactTextString(m) }
func (*VenuesInfo) ProtoMessage()    {}
func (*VenuesInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f696248df6cc389, []int{1}
}

func (m *VenuesInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenuesInfo.Unmarshal(m, b)
}
func (m *VenuesInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenuesInfo.Marshal(b, m, deterministic)
}
func (m *VenuesInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuesInfo.Merge(m, src)
}
func (m *VenuesInfo) XXX_Size() int {
	return xxx_messageInfo_VenuesInfo.Size(m)
}
func (m *VenuesInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VenuesInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VenuesInfo proto.InternalMessageInfo

func (m *VenuesInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VenuesInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VenuesInfo) GetVenueId() int32 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *VenuesInfo) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type VenuesResponse struct {
	Retcode              Retcode       `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Venues               []*VenuesInfo `protobuf:"bytes,2,rep,name=venues,proto3" json:"venues,omitempty"`
	Comment              string        `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VenuesResponse) Reset()         { *m = VenuesResponse{} }
func (m *VenuesResponse) String() string { return proto.CompactTextString(m) }
func (*VenuesResponse) ProtoMessage()    {}
func (*VenuesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f696248df6cc389, []int{2}
}

func (m *VenuesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenuesResponse.Unmarshal(m, b)
}
func (m *VenuesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenuesResponse.Marshal(b, m, deterministic)
}
func (m *VenuesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuesResponse.Merge(m, src)
}
func (m *VenuesResponse) XXX_Size() int {
	return xxx_messageInfo_VenuesResponse.Size(m)
}
func (m *VenuesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VenuesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VenuesResponse proto.InternalMessageInfo

func (m *VenuesResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *VenuesResponse) GetVenues() []*VenuesInfo {
	if m != nil {
		return m.Venues
	}
	return nil
}

func (m *VenuesResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type VenuesPostRequest struct {
	VenueId              int32    `protobuf:"varint,1,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	Enabled              bool     `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ApiKey               string   `protobuf:"bytes,3,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	ApiSecret            string   `protobuf:"bytes,4,opt,name=api_secret,json=apiSecret,proto3" json:"api_secret,omitempty"`
	Passphrase           string   `protobuf:"bytes,5,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenuesPostRequest) Reset()         { *m = VenuesPostRequest{} }
func (m *VenuesPostRequest) String() string { return proto.CompactTextString(m) }
func (*VenuesPostRequest) ProtoMessage()    {}
func (*VenuesPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f696248df6cc389, []int{3}
}

func (m *VenuesPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenuesPostRequest.Unmarshal(m, b)
}
func (m *VenuesPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenuesPostRequest.Marshal(b, m, deterministic)
}
func (m *VenuesPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuesPostRequest.Merge(m, src)
}
func (m *VenuesPostRequest) XXX_Size() int {
	return xxx_messageInfo_VenuesPostRequest.Size(m)
}
func (m *VenuesPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VenuesPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VenuesPostRequest proto.InternalMessageInfo

func (m *VenuesPostRequest) GetVenueId() int32 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *VenuesPostRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *VenuesPostRequest) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *VenuesPostRequest) GetApiSecret() string {
	if m != nil {
		return m.ApiSecret
	}
	return ""
}

func (m *VenuesPostRequest) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

type VenuesPostResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenuesPostResponse) Reset()         { *m = VenuesPostResponse{} }
func (m *VenuesPostResponse) String() string { return proto.CompactTextString(m) }
func (*VenuesPostResponse) ProtoMessage()    {}
func (*VenuesPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f696248df6cc389, []int{4}
}

func (m *VenuesPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenuesPostResponse.Unmarshal(m, b)
}
func (m *VenuesPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenuesPostResponse.Marshal(b, m, deterministic)
}
func (m *VenuesPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuesPostResponse.Merge(m, src)
}
func (m *VenuesPostResponse) XXX_Size() int {
	return xxx_messageInfo_VenuesPostResponse.Size(m)
}
func (m *VenuesPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VenuesPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VenuesPostResponse proto.InternalMessageInfo

func (m *VenuesPostResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *VenuesPostResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

// Venues Detailed
type VenueDetailedRequest struct {
	Venue                Venue    `protobuf:"varint,1,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenueDetailedRequest) Reset()         { *m = VenueDetailedRequest{} }
func (m *VenueDetailedRequest) String() string { return proto.CompactTextString(m) }
func (*VenueDetailedRequest) ProtoMessage()    {}
func (*VenueDetailedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f696248df6cc389, []int{5}
}

func (m *VenueDetailedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenueDetailedRequest.Unmarshal(m, b)
}
func (m *VenueDetailedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenueDetailedRequest.Marshal(b, m, deterministic)
}
func (m *VenueDetailedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueDetailedRequest.Merge(m, src)
}
func (m *VenueDetailedRequest) XXX_Size() int {
	return xxx_messageInfo_VenueDetailedRequest.Size(m)
}
func (m *VenueDetailedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VenueDetailedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VenueDetailedRequest proto.InternalMessageInfo

func (m *VenueDetailedRequest) GetVenue() Venue {
	if m != nil {
		return m.Venue
	}
	return Venue_DARKPOOL
}

type VenueDetailedResponse struct {
	Retcode              Retcode                 `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Id                   string                  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	VenueId              int32                   `protobuf:"varint,4,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	Enabled              bool                    `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	ApiKey               string                  `protobuf:"bytes,6,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	ApiSecret            string                  `protobuf:"bytes,7,opt,name=api_secret,json=apiSecret,proto3" json:"api_secret,omitempty"`
	Passphrase           string                  `protobuf:"bytes,8,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	ProductInfo          []*VenueDetailedProduct `protobuf:"bytes,9,rep,name=product_info,json=productInfo,proto3" json:"product_info,omitempty"`
	Comment              string                  `protobuf:"bytes,10,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *VenueDetailedResponse) Reset()         { *m = VenueDetailedResponse{} }
func (m *VenueDetailedResponse) String() string { return proto.CompactTextString(m) }
func (*VenueDetailedResponse) ProtoMessage()    {}
func (*VenueDetailedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f696248df6cc389, []int{6}
}

func (m *VenueDetailedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenueDetailedResponse.Unmarshal(m, b)
}
func (m *VenueDetailedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenueDetailedResponse.Marshal(b, m, deterministic)
}
func (m *VenueDetailedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueDetailedResponse.Merge(m, src)
}
func (m *VenueDetailedResponse) XXX_Size() int {
	return xxx_messageInfo_VenueDetailedResponse.Size(m)
}
func (m *VenueDetailedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VenueDetailedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VenueDetailedResponse proto.InternalMessageInfo

func (m *VenueDetailedResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *VenueDetailedResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VenueDetailedResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VenueDetailedResponse) GetVenueId() int32 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *VenueDetailedResponse) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *VenueDetailedResponse) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *VenueDetailedResponse) GetApiSecret() string {
	if m != nil {
		return m.ApiSecret
	}
	return ""
}

func (m *VenueDetailedResponse) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *VenueDetailedResponse) GetProductInfo() []*VenueDetailedProduct {
	if m != nil {
		return m.ProductInfo
	}
	return nil
}

func (m *VenueDetailedResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type VenueDetailedProduct struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VenueId              string   `protobuf:"bytes,2,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	Product              int32    `protobuf:"varint,3,opt,name=product,proto3" json:"product,omitempty"`
	Enabled              bool     `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	IndividualConnection bool     `protobuf:"varint,5,opt,name=individual_connection,json=individualConnection,proto3" json:"individual_connection,omitempty"`
	VenueName            string   `protobuf:"bytes,6,opt,name=venue_name,json=venueName,proto3" json:"venue_name,omitempty"`
	ApiName              string   `protobuf:"bytes,7,opt,name=api_name,json=apiName,proto3" json:"api_name,omitempty"`
	MinimumOrdersSize    float64  `protobuf:"fixed64,8,opt,name=minimum_orders_size,json=minimumOrdersSize,proto3" json:"minimum_orders_size,omitempty"`
	StepSize             float64  `protobuf:"fixed64,9,opt,name=step_size,json=stepSize,proto3" json:"step_size,omitempty"`
	MakerFee             float64  `protobuf:"fixed64,10,opt,name=maker_fee,json=makerFee,proto3" json:"maker_fee,omitempty"`
	TakerFee             float64  `protobuf:"fixed64,11,opt,name=taker_fee,json=takerFee,proto3" json:"taker_fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenueDetailedProduct) Reset()         { *m = VenueDetailedProduct{} }
func (m *VenueDetailedProduct) String() string { return proto.CompactTextString(m) }
func (*VenueDetailedProduct) ProtoMessage()    {}
func (*VenueDetailedProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f696248df6cc389, []int{7}
}

func (m *VenueDetailedProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenueDetailedProduct.Unmarshal(m, b)
}
func (m *VenueDetailedProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenueDetailedProduct.Marshal(b, m, deterministic)
}
func (m *VenueDetailedProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueDetailedProduct.Merge(m, src)
}
func (m *VenueDetailedProduct) XXX_Size() int {
	return xxx_messageInfo_VenueDetailedProduct.Size(m)
}
func (m *VenueDetailedProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_VenueDetailedProduct.DiscardUnknown(m)
}

var xxx_messageInfo_VenueDetailedProduct proto.InternalMessageInfo

func (m *VenueDetailedProduct) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VenueDetailedProduct) GetVenueId() string {
	if m != nil {
		return m.VenueId
	}
	return ""
}

func (m *VenueDetailedProduct) GetProduct() int32 {
	if m != nil {
		return m.Product
	}
	return 0
}

func (m *VenueDetailedProduct) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *VenueDetailedProduct) GetIndividualConnection() bool {
	if m != nil {
		return m.IndividualConnection
	}
	return false
}

func (m *VenueDetailedProduct) GetVenueName() string {
	if m != nil {
		return m.VenueName
	}
	return ""
}

func (m *VenueDetailedProduct) GetApiName() string {
	if m != nil {
		return m.ApiName
	}
	return ""
}

func (m *VenueDetailedProduct) GetMinimumOrdersSize() float64 {
	if m != nil {
		return m.MinimumOrdersSize
	}
	return 0
}

func (m *VenueDetailedProduct) GetStepSize() float64 {
	if m != nil {
		return m.StepSize
	}
	return 0
}

func (m *VenueDetailedProduct) GetMakerFee() float64 {
	if m != nil {
		return m.MakerFee
	}
	return 0
}

func (m *VenueDetailedProduct) GetTakerFee() float64 {
	if m != nil {
		return m.TakerFee
	}
	return 0
}

// Venue Product
type VenueProductRequest struct {
	Venue                Venue    `protobuf:"varint,1,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
	Product              Product  `protobuf:"varint,2,opt,name=product,proto3,enum=api.Product" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenueProductRequest) Reset()         { *m = VenueProductRequest{} }
func (m *VenueProductRequest) String() string { return proto.CompactTextString(m) }
func (*VenueProductRequest) ProtoMessage()    {}
func (*VenueProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f696248df6cc389, []int{8}
}

func (m *VenueProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenueProductRequest.Unmarshal(m, b)
}
func (m *VenueProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenueProductRequest.Marshal(b, m, deterministic)
}
func (m *VenueProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueProductRequest.Merge(m, src)
}
func (m *VenueProductRequest) XXX_Size() int {
	return xxx_messageInfo_VenueProductRequest.Size(m)
}
func (m *VenueProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VenueProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VenueProductRequest proto.InternalMessageInfo

func (m *VenueProductRequest) GetVenue() Venue {
	if m != nil {
		return m.Venue
	}
	return Venue_DARKPOOL
}

func (m *VenueProductRequest) GetProduct() Product {
	if m != nil {
		return m.Product
	}
	return Product_BTC_USD
}

type VenueProductResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	VenueId              string   `protobuf:"bytes,3,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	Product              int32    `protobuf:"varint,4,opt,name=product,proto3" json:"product,omitempty"`
	Enabled              bool     `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	IndividualConnection bool     `protobuf:"varint,6,opt,name=individual_connection,json=individualConnection,proto3" json:"individual_connection,omitempty"`
	VenueName            string   `protobuf:"bytes,7,opt,name=venue_name,json=venueName,proto3" json:"venue_name,omitempty"`
	ApiName              string   `protobuf:"bytes,8,opt,name=api_name,json=apiName,proto3" json:"api_name,omitempty"`
	MinimumOrdersSize    float64  `protobuf:"fixed64,9,opt,name=minimum_orders_size,json=minimumOrdersSize,proto3" json:"minimum_orders_size,omitempty"`
	StepSize             float64  `protobuf:"fixed64,10,opt,name=step_size,json=stepSize,proto3" json:"step_size,omitempty"`
	MakerFee             float64  `protobuf:"fixed64,11,opt,name=maker_fee,json=makerFee,proto3" json:"maker_fee,omitempty"`
	TakerFee             float64  `protobuf:"fixed64,12,opt,name=taker_fee,json=takerFee,proto3" json:"taker_fee,omitempty"`
	Comment              string   `protobuf:"bytes,13,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenueProductResponse) Reset()         { *m = VenueProductResponse{} }
func (m *VenueProductResponse) String() string { return proto.CompactTextString(m) }
func (*VenueProductResponse) ProtoMessage()    {}
func (*VenueProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f696248df6cc389, []int{9}
}

func (m *VenueProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenueProductResponse.Unmarshal(m, b)
}
func (m *VenueProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenueProductResponse.Marshal(b, m, deterministic)
}
func (m *VenueProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueProductResponse.Merge(m, src)
}
func (m *VenueProductResponse) XXX_Size() int {
	return xxx_messageInfo_VenueProductResponse.Size(m)
}
func (m *VenueProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VenueProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VenueProductResponse proto.InternalMessageInfo

func (m *VenueProductResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *VenueProductResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VenueProductResponse) GetVenueId() string {
	if m != nil {
		return m.VenueId
	}
	return ""
}

func (m *VenueProductResponse) GetProduct() int32 {
	if m != nil {
		return m.Product
	}
	return 0
}

func (m *VenueProductResponse) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *VenueProductResponse) GetIndividualConnection() bool {
	if m != nil {
		return m.IndividualConnection
	}
	return false
}

func (m *VenueProductResponse) GetVenueName() string {
	if m != nil {
		return m.VenueName
	}
	return ""
}

func (m *VenueProductResponse) GetApiName() string {
	if m != nil {
		return m.ApiName
	}
	return ""
}

func (m *VenueProductResponse) GetMinimumOrdersSize() float64 {
	if m != nil {
		return m.MinimumOrdersSize
	}
	return 0
}

func (m *VenueProductResponse) GetStepSize() float64 {
	if m != nil {
		return m.StepSize
	}
	return 0
}

func (m *VenueProductResponse) GetMakerFee() float64 {
	if m != nil {
		return m.MakerFee
	}
	return 0
}

func (m *VenueProductResponse) GetTakerFee() float64 {
	if m != nil {
		return m.TakerFee
	}
	return 0
}

func (m *VenueProductResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func init() {
	proto.RegisterType((*VenuesRequest)(nil), "api.VenuesRequest")
	proto.RegisterType((*VenuesInfo)(nil), "api.VenuesInfo")
	proto.RegisterType((*VenuesResponse)(nil), "api.VenuesResponse")
	proto.RegisterType((*VenuesPostRequest)(nil), "api.VenuesPostRequest")
	proto.RegisterType((*VenuesPostResponse)(nil), "api.VenuesPostResponse")
	proto.RegisterType((*VenueDetailedRequest)(nil), "api.VenueDetailedRequest")
	proto.RegisterType((*VenueDetailedResponse)(nil), "api.VenueDetailedResponse")
	proto.RegisterType((*VenueDetailedProduct)(nil), "api.VenueDetailedProduct")
	proto.RegisterType((*VenueProductRequest)(nil), "api.VenueProductRequest")
	proto.RegisterType((*VenueProductResponse)(nil), "api.VenueProductResponse")
}

func init() { proto.RegisterFile("webapi.proto", fileDescriptor_9f696248df6cc389) }

var fileDescriptor_9f696248df6cc389 = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x65, 0xe7, 0xd3, 0x93, 0xb4, 0x55, 0xb7, 0xad, 0x70, 0x41, 0xa0, 0xc8, 0x07, 0xc8,
	0xa9, 0x87, 0xf6, 0xc2, 0x81, 0x1b, 0x08, 0xa9, 0x42, 0x82, 0x6a, 0x2b, 0xf5, 0x6a, 0xb9, 0xf1,
	0x54, 0xac, 0x5a, 0xef, 0x1a, 0xef, 0xa6, 0xa8, 0xe5, 0x65, 0x78, 0x01, 0xde, 0x83, 0x77, 0xe1,
	0x25, 0x90, 0x67, 0x77, 0x1b, 0x27, 0xc5, 0x51, 0x2a, 0x6e, 0xd9, 0xf9, 0xef, 0xc7, 0xcc, 0xff,
	0x37, 0x13, 0xc3, 0xf8, 0x3b, 0x5e, 0x66, 0xa5, 0x38, 0x2a, 0x2b, 0x65, 0x14, 0xeb, 0x64, 0xa5,
	0x78, 0x0e, 0x28, 0xe7, 0x85, 0x0d, 0x24, 0x3b, 0xb0, 0x75, 0x81, 0x72, 0x8e, 0x9a, 0xe3, 0xb7,
	0x39, 0x6a, 0x93, 0x20, 0x80, 0x0d, 0x9c, 0xca, 0x2b, 0xc5, 0xb6, 0x21, 0x14, 0x79, 0x1c, 0x4c,
	0x82, 0x69, 0xc4, 0x43, 0x91, 0x33, 0x06, 0x5d, 0x99, 0x15, 0x18, 0x87, 0x14, 0xa1, 0xdf, 0xec,
	0x10, 0x86, 0xb7, 0xf5, 0x89, 0x54, 0xe4, 0x71, 0x67, 0x12, 0x4c, 0x7b, 0x7c, 0x40, 0xeb, 0xd3,
	0x9c, 0xc5, 0x30, 0x40, 0x99, 0x5d, 0xde, 0x60, 0x1e, 0x77, 0x27, 0xc1, 0x74, 0xc8, 0xfd, 0x32,
	0xf9, 0x01, 0xdb, 0xfe, 0x5d, 0x5d, 0x2a, 0xa9, 0x91, 0xbd, 0x86, 0x41, 0x85, 0x66, 0xa6, 0x72,
	0xa4, 0xf7, 0xb6, 0x8f, 0xc7, 0x47, 0x75, 0xde, 0xdc, 0xc6, 0xb8, 0x17, 0xd9, 0x1b, 0xe8, 0xd3,
	0xf5, 0x3a, 0x0e, 0x27, 0x9d, 0xe9, 0xe8, 0x78, 0x87, 0xb6, 0x2d, 0x72, 0xe6, 0x4e, 0xae, 0x1f,
	0x9f, 0xa9, 0xa2, 0x40, 0x69, 0x28, 0xad, 0x88, 0xfb, 0x65, 0xf2, 0x33, 0x80, 0x5d, 0x7b, 0xe0,
	0x4c, 0x69, 0xe3, 0x2a, 0x5f, 0xaa, 0x23, 0x68, 0xad, 0x23, 0x5c, 0xaa, 0x83, 0x3d, 0x83, 0x41,
	0x56, 0x8a, 0xf4, 0x1a, 0xef, 0xdc, 0x23, 0xfd, 0xac, 0x14, 0x9f, 0xf0, 0x8e, 0xbd, 0x04, 0xa8,
	0x05, 0x8d, 0xb3, 0x0a, 0x0d, 0x55, 0x1f, 0xf1, 0x28, 0x2b, 0xc5, 0x39, 0x05, 0xd8, 0x2b, 0x80,
	0x32, 0xd3, 0xba, 0xfc, 0x5a, 0x65, 0x1a, 0xe3, 0x1e, 0xc9, 0x8d, 0x48, 0x72, 0x01, 0xac, 0x99,
	0xe1, 0x13, 0x3d, 0x6a, 0x94, 0x1e, 0x2e, 0x97, 0xfe, 0x16, 0xf6, 0xe9, 0xde, 0x0f, 0x68, 0x32,
	0x71, 0x83, 0xb9, 0x2f, 0x7e, 0x02, 0x3d, 0x2a, 0xd6, 0xdd, 0x0b, 0x0b, 0x53, 0xb9, 0x15, 0x92,
	0xdf, 0x21, 0x1c, 0xac, 0x1c, 0x7d, 0x62, 0x56, 0xb6, 0x99, 0xc2, 0x47, 0xcd, 0xd4, 0x69, 0x69,
	0xa6, 0x6e, 0x2b, 0x84, 0x5e, 0x2b, 0x84, 0xfe, 0x1a, 0x08, 0x83, 0xf5, 0x10, 0x86, 0xab, 0x10,
	0xd8, 0x3b, 0x18, 0x97, 0x95, 0xca, 0xe7, 0x33, 0x93, 0x0a, 0x79, 0xa5, 0xe2, 0x88, 0x1a, 0xee,
	0x70, 0xe1, 0x8d, 0xb7, 0xe2, 0xcc, 0xee, 0xe2, 0x23, 0xb7, 0x9d, 0x66, 0xa7, 0x01, 0x01, 0x96,
	0x21, 0xfc, 0x09, 0x57, 0x28, 0xb8, 0xf3, 0x8f, 0xc6, 0xad, 0xe9, 0x86, 0x03, 0xd9, 0x70, 0xc3,
	0x3d, 0xe6, 0x87, 0xce, 0x2d, 0xdb, 0x87, 0x8e, 0x9d, 0xc0, 0x81, 0x90, 0xb9, 0xb8, 0x15, 0xf9,
	0x3c, 0xbb, 0x49, 0x67, 0x4a, 0x4a, 0x9c, 0x19, 0xa1, 0xa4, 0xf3, 0x73, 0x7f, 0x21, 0xbe, 0x7f,
	0xd0, 0x6a, 0x0f, 0x6d, 0x0e, 0xc4, 0xca, 0xfa, 0x1b, 0x51, 0xe4, 0xb3, 0x03, 0x56, 0x5b, 0x4c,
	0xa2, 0x35, 0xb8, 0x66, 0x41, 0xd2, 0x11, 0xec, 0x15, 0x42, 0x8a, 0x62, 0x5e, 0xa4, 0xaa, 0xca,
	0xb1, 0xd2, 0xa9, 0x16, 0xf7, 0xd6, 0xe7, 0x80, 0xef, 0x3a, 0xe9, 0x0b, 0x29, 0xe7, 0xe2, 0x1e,
	0xd9, 0x0b, 0x88, 0xb4, 0xc1, 0xd2, 0xee, 0x8a, 0x68, 0xd7, 0xb0, 0x0e, 0x78, 0xb1, 0xc8, 0xae,
	0xb1, 0x4a, 0xaf, 0x10, 0xc9, 0xcf, 0x80, 0x0f, 0x29, 0xf0, 0x11, 0x49, 0x34, 0x0f, 0xe2, 0xc8,
	0x8a, 0xc6, 0x89, 0x49, 0x0a, 0x7b, 0x64, 0xb6, 0x87, 0xb4, 0x69, 0xc7, 0xd7, 0x7d, 0xed, 0x2d,
	0x0e, 0x1b, 0x7d, 0xed, 0xef, 0xf1, 0x62, 0xf2, 0xab, 0xe3, 0x70, 0x3e, 0xbc, 0xf0, 0x9f, 0x83,
	0xb1, 0xfa, 0x8f, 0xfa, 0x6f, 0xec, 0xdd, 0x56, 0xec, 0xbd, 0x0d, 0xb1, 0xf7, 0x37, 0xc6, 0x3e,
	0x58, 0x87, 0x7d, 0xb8, 0x11, 0xf6, 0x68, 0x23, 0xec, 0xb0, 0x0e, 0xfb, 0x68, 0x1d, 0xf6, 0xf1,
	0x32, 0xf6, 0xe6, 0xf8, 0x6d, 0x2d, 0x8d, 0xdf, 0x65, 0x9f, 0x3e, 0x7d, 0x27, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0xf7, 0x86, 0xf3, 0xed, 0x1b, 0x07, 0x00, 0x00,
}
