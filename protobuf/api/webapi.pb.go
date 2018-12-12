// Code generated by protoc-gen-go. DO NOT EDIT.
// source: webapi.proto

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
func (dst *VenuesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuesRequest.Merge(dst, src)
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
	Venue                Venue    `protobuf:"varint,3,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
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
func (dst *VenuesInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuesInfo.Merge(dst, src)
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

func (m *VenuesInfo) GetVenue() Venue {
	if m != nil {
		return m.Venue
	}
	return Venue_DARKPOOL
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
func (dst *VenuesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuesResponse.Merge(dst, src)
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
	Venue                Venue    `protobuf:"varint,1,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
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
func (dst *VenuesPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuesPostRequest.Merge(dst, src)
}
func (m *VenuesPostRequest) XXX_Size() int {
	return xxx_messageInfo_VenuesPostRequest.Size(m)
}
func (m *VenuesPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VenuesPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VenuesPostRequest proto.InternalMessageInfo

func (m *VenuesPostRequest) GetVenue() Venue {
	if m != nil {
		return m.Venue
	}
	return Venue_DARKPOOL
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
func (dst *VenuesPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuesPostResponse.Merge(dst, src)
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
func (dst *VenueDetailedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueDetailedRequest.Merge(dst, src)
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
	Venue                Venue                   `protobuf:"varint,4,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
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
func (dst *VenueDetailedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueDetailedResponse.Merge(dst, src)
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

func (m *VenueDetailedResponse) GetVenue() Venue {
	if m != nil {
		return m.Venue
	}
	return Venue_DARKPOOL
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
	Venue                Venue    `protobuf:"varint,2,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
	Product              Product  `protobuf:"varint,3,opt,name=product,proto3,enum=api.Product" json:"product,omitempty"`
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
func (dst *VenueDetailedProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueDetailedProduct.Merge(dst, src)
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

func (m *VenueDetailedProduct) GetVenue() Venue {
	if m != nil {
		return m.Venue
	}
	return Venue_DARKPOOL
}

func (m *VenueDetailedProduct) GetProduct() Product {
	if m != nil {
		return m.Product
	}
	return Product_BTC_USD
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
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
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
func (dst *VenueProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueProductRequest.Merge(dst, src)
}
func (m *VenueProductRequest) XXX_Size() int {
	return xxx_messageInfo_VenueProductRequest.Size(m)
}
func (m *VenueProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VenueProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VenueProductRequest proto.InternalMessageInfo

func (m *VenueProductRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type VenueProductResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Venue                Venue    `protobuf:"varint,3,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
	Product              Product  `protobuf:"varint,4,opt,name=product,proto3,enum=api.Product" json:"product,omitempty"`
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
func (dst *VenueProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueProductResponse.Merge(dst, src)
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

func (m *VenueProductResponse) GetVenue() Venue {
	if m != nil {
		return m.Venue
	}
	return Venue_DARKPOOL
}

func (m *VenueProductResponse) GetProduct() Product {
	if m != nil {
		return m.Product
	}
	return Product_BTC_USD
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

type ProductPostRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Venue                Venue    `protobuf:"varint,2,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
	Product              Product  `protobuf:"varint,3,opt,name=product,proto3,enum=api.Product" json:"product,omitempty"`
	Enabled              bool     `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	IndividualConnection bool     `protobuf:"varint,5,opt,name=individual_connection,json=individualConnection,proto3" json:"individual_connection,omitempty"`
	VenueName            string   `protobuf:"bytes,6,opt,name=venue_name,json=venueName,proto3" json:"venue_name,omitempty"`
	ApiName              string   `protobuf:"bytes,7,opt,name=api_name,json=apiName,proto3" json:"api_name,omitempty"`
	MinimumOrdersSize    float64  `protobuf:"fixed64,8,opt,name=minimum_orders_size,json=minimumOrdersSize,proto3" json:"minimum_orders_size,omitempty"`
	StepSize             float64  `protobuf:"fixed64,9,opt,name=step_size,json=stepSize,proto3" json:"step_size,omitempty"`
	MakerFee             float64  `protobuf:"fixed64,10,opt,name=maker_fee,json=makerFee,proto3" json:"maker_fee,omitempty"`
	TakerFee             float64  `protobuf:"fixed64,11,opt,name=taker_fee,json=takerFee,proto3" json:"taker_fee,omitempty"`
	Action               Action   `protobuf:"varint,12,opt,name=action,proto3,enum=api.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductPostRequest) Reset()         { *m = ProductPostRequest{} }
func (m *ProductPostRequest) String() string { return proto.CompactTextString(m) }
func (*ProductPostRequest) ProtoMessage()    {}
func (*ProductPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f696248df6cc389, []int{10}
}
func (m *ProductPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductPostRequest.Unmarshal(m, b)
}
func (m *ProductPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductPostRequest.Marshal(b, m, deterministic)
}
func (dst *ProductPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductPostRequest.Merge(dst, src)
}
func (m *ProductPostRequest) XXX_Size() int {
	return xxx_messageInfo_ProductPostRequest.Size(m)
}
func (m *ProductPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProductPostRequest proto.InternalMessageInfo

func (m *ProductPostRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProductPostRequest) GetVenue() Venue {
	if m != nil {
		return m.Venue
	}
	return Venue_DARKPOOL
}

func (m *ProductPostRequest) GetProduct() Product {
	if m != nil {
		return m.Product
	}
	return Product_BTC_USD
}

func (m *ProductPostRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *ProductPostRequest) GetIndividualConnection() bool {
	if m != nil {
		return m.IndividualConnection
	}
	return false
}

func (m *ProductPostRequest) GetVenueName() string {
	if m != nil {
		return m.VenueName
	}
	return ""
}

func (m *ProductPostRequest) GetApiName() string {
	if m != nil {
		return m.ApiName
	}
	return ""
}

func (m *ProductPostRequest) GetMinimumOrdersSize() float64 {
	if m != nil {
		return m.MinimumOrdersSize
	}
	return 0
}

func (m *ProductPostRequest) GetStepSize() float64 {
	if m != nil {
		return m.StepSize
	}
	return 0
}

func (m *ProductPostRequest) GetMakerFee() float64 {
	if m != nil {
		return m.MakerFee
	}
	return 0
}

func (m *ProductPostRequest) GetTakerFee() float64 {
	if m != nil {
		return m.TakerFee
	}
	return 0
}

func (m *ProductPostRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_INSERT
}

type ProductPostResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductPostResponse) Reset()         { *m = ProductPostResponse{} }
func (m *ProductPostResponse) String() string { return proto.CompactTextString(m) }
func (*ProductPostResponse) ProtoMessage()    {}
func (*ProductPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f696248df6cc389, []int{11}
}
func (m *ProductPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductPostResponse.Unmarshal(m, b)
}
func (m *ProductPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductPostResponse.Marshal(b, m, deterministic)
}
func (dst *ProductPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductPostResponse.Merge(dst, src)
}
func (m *ProductPostResponse) XXX_Size() int {
	return xxx_messageInfo_ProductPostResponse.Size(m)
}
func (m *ProductPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductPostResponse proto.InternalMessageInfo

func (m *ProductPostResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *ProductPostResponse) GetComment() string {
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
	proto.RegisterType((*ProductPostRequest)(nil), "api.ProductPostRequest")
	proto.RegisterType((*ProductPostResponse)(nil), "api.ProductPostResponse")
}

func init() { proto.RegisterFile("webapi.proto", fileDescriptor_9f696248df6cc389) }

var fileDescriptor_9f696248df6cc389 = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x96, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0x65, 0xe7, 0xaf, 0x27, 0x69, 0xaa, 0x6e, 0x5b, 0xfd, 0xdc, 0x1f, 0x02, 0x45, 0x46,
	0x94, 0x9c, 0x7a, 0x68, 0x2f, 0x1c, 0xb8, 0x20, 0x10, 0x12, 0x42, 0x82, 0x6a, 0x2b, 0x95, 0xa3,
	0xb5, 0x8d, 0xa7, 0x62, 0xd5, 0x78, 0x6d, 0xbc, 0x4e, 0x51, 0xcb, 0x85, 0x77, 0xe1, 0xc0, 0x7b,
	0xf0, 0x02, 0xbc, 0x12, 0xf2, 0xec, 0x3a, 0xb1, 0x53, 0x1a, 0x5c, 0x81, 0x38, 0x71, 0x8b, 0xe7,
	0xbb, 0xbb, 0x33, 0xfb, 0xfd, 0xcc, 0x38, 0x86, 0xe1, 0x47, 0x3c, 0x13, 0xa9, 0x3c, 0x48, 0xb3,
	0x24, 0x4f, 0x58, 0x4b, 0xa4, 0xf2, 0x7f, 0x40, 0x35, 0x8f, 0x4d, 0x20, 0xd8, 0x84, 0x8d, 0x53,
	0x54, 0x73, 0xd4, 0x1c, 0x3f, 0xcc, 0x51, 0xe7, 0xc1, 0x0c, 0xc0, 0x04, 0x5e, 0xa9, 0xf3, 0x84,
	0x8d, 0xc0, 0x95, 0x91, 0xef, 0x8c, 0x9d, 0x89, 0xc7, 0x5d, 0x19, 0x31, 0x06, 0x6d, 0x25, 0x62,
	0xf4, 0x5d, 0x8a, 0xd0, 0x6f, 0x36, 0x86, 0xce, 0x65, 0xb1, 0xc3, 0x6f, 0x8d, 0x9d, 0xc9, 0xe8,
	0x10, 0x0e, 0x8a, 0x74, 0x74, 0x06, 0x37, 0x02, 0xf3, 0xa1, 0x87, 0x4a, 0x9c, 0xcd, 0x30, 0xf2,
	0xdb, 0x63, 0x67, 0xd2, 0xe7, 0xe5, 0x63, 0xf0, 0x09, 0x46, 0x65, 0x7a, 0x9d, 0x26, 0x4a, 0x23,
	0xdb, 0x87, 0x5e, 0x86, 0xf9, 0x34, 0x89, 0x90, 0xd2, 0x8e, 0x0e, 0x87, 0x74, 0x1e, 0x37, 0x31,
	0x5e, 0x8a, 0xec, 0x31, 0x74, 0xe9, 0x70, 0xed, 0xbb, 0xe3, 0xd6, 0x64, 0x70, 0xb8, 0xb9, 0x4c,
	0x4b, 0xa5, 0x73, 0x2b, 0x17, 0xc9, 0xa7, 0x49, 0x1c, 0xa3, 0xca, 0xa9, 0x40, 0x8f, 0x97, 0x8f,
	0xc1, 0x57, 0x07, 0xb6, 0xcc, 0x86, 0xe3, 0x44, 0xe7, 0xd6, 0x80, 0xe5, 0x75, 0x9c, 0x06, 0xd7,
	0x71, 0x6b, 0xd7, 0x61, 0xff, 0x41, 0x4f, 0xa4, 0x32, 0xbc, 0xc0, 0x2b, 0x9b, 0xab, 0x2b, 0x52,
	0xf9, 0x1a, 0xaf, 0xd8, 0x7d, 0x80, 0x42, 0xd0, 0x38, 0xcd, 0x30, 0x27, 0x13, 0x3c, 0xee, 0x89,
	0x54, 0x9e, 0x50, 0x80, 0x3d, 0x00, 0x48, 0x85, 0xd6, 0xe9, 0xfb, 0x4c, 0x68, 0xf4, 0x3b, 0x24,
	0x57, 0x22, 0xc1, 0x29, 0xb0, 0x6a, 0xa1, 0x77, 0xb4, 0xaa, 0xe2, 0x80, 0x5b, 0x77, 0xe0, 0x09,
	0xec, 0xd0, 0xb9, 0x2f, 0x30, 0x17, 0x72, 0x86, 0x51, 0x63, 0x0f, 0x82, 0xef, 0x2e, 0xec, 0xae,
	0x6c, 0xbd, 0x63, 0x55, 0xa6, 0xb5, 0xdc, 0x1b, 0xad, 0xd5, 0xfa, 0x59, 0x6b, 0xb5, 0x1b, 0xb0,
	0xe8, 0xdc, 0xca, 0xa2, 0xbb, 0x86, 0x45, 0x6f, 0x3d, 0x8b, 0xfe, 0x2a, 0x0b, 0xf6, 0x14, 0x86,
	0x69, 0x96, 0x44, 0xf3, 0x69, 0x1e, 0x4a, 0x75, 0x9e, 0xf8, 0x1e, 0xb5, 0xdf, 0xde, 0xb2, 0xb4,
	0xd2, 0x91, 0x63, 0xb3, 0x8a, 0x0f, 0xec, 0x72, 0x1a, 0xa8, 0x0a, 0x0b, 0xa8, 0xb3, 0xf8, 0xdc,
	0x5a, 0x81, 0x61, 0xf7, 0xdf, 0x98, 0xc1, 0x85, 0x29, 0xee, 0x6d, 0xa6, 0xec, 0x43, 0xcf, 0xe6,
	0xb4, 0x33, 0x69, 0x10, 0x94, 0x05, 0x95, 0xe2, 0xed, 0x73, 0xc9, 0x8e, 0x60, 0x57, 0xaa, 0x48,
	0x5e, 0xca, 0x68, 0x2e, 0x66, 0xe1, 0x34, 0x51, 0x0a, 0xa7, 0xb9, 0x4c, 0x94, 0x35, 0x79, 0x67,
	0x29, 0x3e, 0x5f, 0x68, 0x85, 0xb1, 0x94, 0x3f, 0x24, 0x8e, 0xc6, 0x74, 0x8f, 0x22, 0x6f, 0x0a,
	0x98, 0x7b, 0xd0, 0x2f, 0x7c, 0x27, 0xd1, 0xb8, 0x5e, 0x00, 0x22, 0xe9, 0x00, 0xb6, 0x63, 0xa9,
	0x64, 0x3c, 0x8f, 0xc3, 0x24, 0x8b, 0x30, 0xd3, 0xa1, 0x96, 0xd7, 0xc6, 0x7c, 0x87, 0x6f, 0x59,
	0xe9, 0x2d, 0x29, 0x27, 0xf2, 0x1a, 0xd9, 0x3d, 0xf0, 0x74, 0x8e, 0xa9, 0x59, 0xe5, 0xd1, 0xaa,
	0x7e, 0x11, 0x28, 0xc5, 0x58, 0x5c, 0x60, 0x16, 0x9e, 0x23, 0x92, 0xc9, 0x0e, 0xef, 0x53, 0xe0,
	0x25, 0x92, 0x98, 0x2f, 0xc4, 0x81, 0x11, 0x73, 0x2b, 0x06, 0x8f, 0x60, 0x9b, 0x7c, 0x2c, 0x8d,
	0xb2, 0xd3, 0xb0, 0x02, 0x20, 0xf8, 0x56, 0x92, 0x5a, 0xac, 0xfb, 0xcd, 0xd6, 0xff, 0xf5, 0x1b,
	0xb4, 0x42, 0xb4, 0xdd, 0x90, 0x68, 0xa7, 0x21, 0xd1, 0x6e, 0x63, 0xa2, 0xbd, 0x75, 0x44, 0xfb,
	0x8d, 0x88, 0x7a, 0x8d, 0x88, 0xc2, 0x3a, 0xa2, 0x83, 0x75, 0x44, 0x87, 0x75, 0xa2, 0xd5, 0x71,
	0xdb, 0xa8, 0x8f, 0xdb, 0x97, 0x16, 0x30, 0x6b, 0x5f, 0xf5, 0xed, 0xff, 0x6f, 0xd8, 0xfe, 0xd2,
	0xb0, 0xb1, 0x87, 0xd0, 0x15, 0xe6, 0x9a, 0x43, 0xb2, 0x6d, 0x40, 0xb6, 0x3d, 0xa3, 0x10, 0xb7,
	0x52, 0xf0, 0x0e, 0xb6, 0x6b, 0x90, 0xfe, 0xd4, 0x3f, 0xdf, 0x59, 0x97, 0x3e, 0x7f, 0x8e, 0x7e,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xa8, 0x5e, 0xa7, 0x1f, 0x09, 0x00, 0x00,
}
