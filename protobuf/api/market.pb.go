// Code generated by protoc-gen-go. DO NOT EDIT.
// source: market.proto

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

type MarketKindSummary struct {
	MarketDataId          int64    `protobuf:"varint,1,opt,name=market_data_id,json=marketDataId,proto3" json:"market_data_id,omitempty"`
	Product               string   `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	VenuesTotal           int64    `protobuf:"varint,3,opt,name=venues_total,json=venuesTotal,proto3" json:"venues_total,omitempty"`
	ProductsEnabledTotal  int64    `protobuf:"varint,4,opt,name=products_enabled_total,json=productsEnabledTotal,proto3" json:"products_enabled_total,omitempty"`
	ProductsDisabledTotal int64    `protobuf:"varint,5,opt,name=products_disabled_total,json=productsDisabledTotal,proto3" json:"products_disabled_total,omitempty"`
	StreamingSaveEnabled  int64    `protobuf:"varint,6,opt,name=streaming_save_enabled,json=streamingSaveEnabled,proto3" json:"streaming_save_enabled,omitempty"`
	StreamingSaveDisabled int64    `protobuf:"varint,7,opt,name=streaming_save_disabled,json=streamingSaveDisabled,proto3" json:"streaming_save_disabled,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *MarketKindSummary) Reset()         { *m = MarketKindSummary{} }
func (m *MarketKindSummary) String() string { return proto.CompactTextString(m) }
func (*MarketKindSummary) ProtoMessage()    {}
func (*MarketKindSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{0}
}
func (m *MarketKindSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketKindSummary.Unmarshal(m, b)
}
func (m *MarketKindSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketKindSummary.Marshal(b, m, deterministic)
}
func (dst *MarketKindSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketKindSummary.Merge(dst, src)
}
func (m *MarketKindSummary) XXX_Size() int {
	return xxx_messageInfo_MarketKindSummary.Size(m)
}
func (m *MarketKindSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketKindSummary.DiscardUnknown(m)
}

var xxx_messageInfo_MarketKindSummary proto.InternalMessageInfo

func (m *MarketKindSummary) GetMarketDataId() int64 {
	if m != nil {
		return m.MarketDataId
	}
	return 0
}

func (m *MarketKindSummary) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *MarketKindSummary) GetVenuesTotal() int64 {
	if m != nil {
		return m.VenuesTotal
	}
	return 0
}

func (m *MarketKindSummary) GetProductsEnabledTotal() int64 {
	if m != nil {
		return m.ProductsEnabledTotal
	}
	return 0
}

func (m *MarketKindSummary) GetProductsDisabledTotal() int64 {
	if m != nil {
		return m.ProductsDisabledTotal
	}
	return 0
}

func (m *MarketKindSummary) GetStreamingSaveEnabled() int64 {
	if m != nil {
		return m.StreamingSaveEnabled
	}
	return 0
}

func (m *MarketKindSummary) GetStreamingSaveDisabled() int64 {
	if m != nil {
		return m.StreamingSaveDisabled
	}
	return 0
}

type MarketKindSummaryRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Kind                 string   `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarketKindSummaryRequest) Reset()         { *m = MarketKindSummaryRequest{} }
func (m *MarketKindSummaryRequest) String() string { return proto.CompactTextString(m) }
func (*MarketKindSummaryRequest) ProtoMessage()    {}
func (*MarketKindSummaryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{1}
}
func (m *MarketKindSummaryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketKindSummaryRequest.Unmarshal(m, b)
}
func (m *MarketKindSummaryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketKindSummaryRequest.Marshal(b, m, deterministic)
}
func (dst *MarketKindSummaryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketKindSummaryRequest.Merge(dst, src)
}
func (m *MarketKindSummaryRequest) XXX_Size() int {
	return xxx_messageInfo_MarketKindSummaryRequest.Size(m)
}
func (m *MarketKindSummaryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketKindSummaryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MarketKindSummaryRequest proto.InternalMessageInfo

func (m *MarketKindSummaryRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *MarketKindSummaryRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

type MarketKindSummaryResponse struct {
	Retcode              Retcode              `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	MarketKindSummary    []*MarketKindSummary `protobuf:"bytes,2,rep,name=market_kind_summary,json=marketKindSummary,proto3" json:"market_kind_summary,omitempty"`
	Comment              string               `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string               `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MarketKindSummaryResponse) Reset()         { *m = MarketKindSummaryResponse{} }
func (m *MarketKindSummaryResponse) String() string { return proto.CompactTextString(m) }
func (*MarketKindSummaryResponse) ProtoMessage()    {}
func (*MarketKindSummaryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{2}
}
func (m *MarketKindSummaryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketKindSummaryResponse.Unmarshal(m, b)
}
func (m *MarketKindSummaryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketKindSummaryResponse.Marshal(b, m, deterministic)
}
func (dst *MarketKindSummaryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketKindSummaryResponse.Merge(dst, src)
}
func (m *MarketKindSummaryResponse) XXX_Size() int {
	return xxx_messageInfo_MarketKindSummaryResponse.Size(m)
}
func (m *MarketKindSummaryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketKindSummaryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MarketKindSummaryResponse proto.InternalMessageInfo

func (m *MarketKindSummaryResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *MarketKindSummaryResponse) GetMarketKindSummary() []*MarketKindSummary {
	if m != nil {
		return m.MarketKindSummary
	}
	return nil
}

func (m *MarketKindSummaryResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *MarketKindSummaryResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type MarketProductSummary struct {
	MarketDataId         int64    `protobuf:"varint,1,opt,name=market_data_id,json=marketDataId,proto3" json:"market_data_id,omitempty"`
	ProductId            int64    `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductDescription   string   `protobuf:"bytes,3,opt,name=product_description,json=productDescription,proto3" json:"product_description,omitempty"`
	VenueId              int64    `protobuf:"varint,4,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	VenueDescription     string   `protobuf:"bytes,5,opt,name=venue_description,json=venueDescription,proto3" json:"venue_description,omitempty"`
	Price                float64  `protobuf:"fixed64,6,opt,name=price,proto3" json:"price,omitempty"`
	Volume               float64  `protobuf:"fixed64,7,opt,name=volume,proto3" json:"volume,omitempty"`
	OrderSide            string   `protobuf:"bytes,8,opt,name=order_side,json=orderSide,proto3" json:"order_side,omitempty"`
	SystemTimestamp      string   `protobuf:"bytes,9,opt,name=system_timestamp,json=systemTimestamp,proto3" json:"system_timestamp,omitempty"`
	VenueTimestamp       string   `protobuf:"bytes,10,opt,name=venue_timestamp,json=venueTimestamp,proto3" json:"venue_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarketProductSummary) Reset()         { *m = MarketProductSummary{} }
func (m *MarketProductSummary) String() string { return proto.CompactTextString(m) }
func (*MarketProductSummary) ProtoMessage()    {}
func (*MarketProductSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{3}
}
func (m *MarketProductSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketProductSummary.Unmarshal(m, b)
}
func (m *MarketProductSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketProductSummary.Marshal(b, m, deterministic)
}
func (dst *MarketProductSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketProductSummary.Merge(dst, src)
}
func (m *MarketProductSummary) XXX_Size() int {
	return xxx_messageInfo_MarketProductSummary.Size(m)
}
func (m *MarketProductSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketProductSummary.DiscardUnknown(m)
}

var xxx_messageInfo_MarketProductSummary proto.InternalMessageInfo

func (m *MarketProductSummary) GetMarketDataId() int64 {
	if m != nil {
		return m.MarketDataId
	}
	return 0
}

func (m *MarketProductSummary) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *MarketProductSummary) GetProductDescription() string {
	if m != nil {
		return m.ProductDescription
	}
	return ""
}

func (m *MarketProductSummary) GetVenueId() int64 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *MarketProductSummary) GetVenueDescription() string {
	if m != nil {
		return m.VenueDescription
	}
	return ""
}

func (m *MarketProductSummary) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *MarketProductSummary) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *MarketProductSummary) GetOrderSide() string {
	if m != nil {
		return m.OrderSide
	}
	return ""
}

func (m *MarketProductSummary) GetSystemTimestamp() string {
	if m != nil {
		return m.SystemTimestamp
	}
	return ""
}

func (m *MarketProductSummary) GetVenueTimestamp() string {
	if m != nil {
		return m.VenueTimestamp
	}
	return ""
}

type MarketProductSummaryRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Kind                 string   `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	ProductDescription   string   `protobuf:"bytes,3,opt,name=product_description,json=productDescription,proto3" json:"product_description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarketProductSummaryRequest) Reset()         { *m = MarketProductSummaryRequest{} }
func (m *MarketProductSummaryRequest) String() string { return proto.CompactTextString(m) }
func (*MarketProductSummaryRequest) ProtoMessage()    {}
func (*MarketProductSummaryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{4}
}
func (m *MarketProductSummaryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketProductSummaryRequest.Unmarshal(m, b)
}
func (m *MarketProductSummaryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketProductSummaryRequest.Marshal(b, m, deterministic)
}
func (dst *MarketProductSummaryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketProductSummaryRequest.Merge(dst, src)
}
func (m *MarketProductSummaryRequest) XXX_Size() int {
	return xxx_messageInfo_MarketProductSummaryRequest.Size(m)
}
func (m *MarketProductSummaryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketProductSummaryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MarketProductSummaryRequest proto.InternalMessageInfo

func (m *MarketProductSummaryRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *MarketProductSummaryRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *MarketProductSummaryRequest) GetProductDescription() string {
	if m != nil {
		return m.ProductDescription
	}
	return ""
}

type MarketProductSummaryResponse struct {
	Retcode              Retcode                 `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	MarketProductSummary []*MarketProductSummary `protobuf:"bytes,2,rep,name=market_product_summary,json=marketProductSummary,proto3" json:"market_product_summary,omitempty"`
	Comment              string                  `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string                  `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MarketProductSummaryResponse) Reset()         { *m = MarketProductSummaryResponse{} }
func (m *MarketProductSummaryResponse) String() string { return proto.CompactTextString(m) }
func (*MarketProductSummaryResponse) ProtoMessage()    {}
func (*MarketProductSummaryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{5}
}
func (m *MarketProductSummaryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketProductSummaryResponse.Unmarshal(m, b)
}
func (m *MarketProductSummaryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketProductSummaryResponse.Marshal(b, m, deterministic)
}
func (dst *MarketProductSummaryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketProductSummaryResponse.Merge(dst, src)
}
func (m *MarketProductSummaryResponse) XXX_Size() int {
	return xxx_messageInfo_MarketProductSummaryResponse.Size(m)
}
func (m *MarketProductSummaryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketProductSummaryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MarketProductSummaryResponse proto.InternalMessageInfo

func (m *MarketProductSummaryResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *MarketProductSummaryResponse) GetMarketProductSummary() []*MarketProductSummary {
	if m != nil {
		return m.MarketProductSummary
	}
	return nil
}

func (m *MarketProductSummaryResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *MarketProductSummaryResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type Bar struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Open                 float64  `protobuf:"fixed64,2,opt,name=open,proto3" json:"open,omitempty"`
	High                 float64  `protobuf:"fixed64,3,opt,name=high,proto3" json:"high,omitempty"`
	Low                  float64  `protobuf:"fixed64,4,opt,name=low,proto3" json:"low,omitempty"`
	Close                float64  `protobuf:"fixed64,5,opt,name=close,proto3" json:"close,omitempty"`
	Volume               float64  `protobuf:"fixed64,6,opt,name=volume,proto3" json:"volume,omitempty"`
	Trades               int64    `protobuf:"varint,7,opt,name=trades,proto3" json:"trades,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bar) Reset()         { *m = Bar{} }
func (m *Bar) String() string { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()    {}
func (*Bar) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{6}
}
func (m *Bar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bar.Unmarshal(m, b)
}
func (m *Bar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bar.Marshal(b, m, deterministic)
}
func (dst *Bar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bar.Merge(dst, src)
}
func (m *Bar) XXX_Size() int {
	return xxx_messageInfo_Bar.Size(m)
}
func (m *Bar) XXX_DiscardUnknown() {
	xxx_messageInfo_Bar.DiscardUnknown(m)
}

var xxx_messageInfo_Bar proto.InternalMessageInfo

func (m *Bar) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Bar) GetOpen() float64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *Bar) GetHigh() float64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *Bar) GetLow() float64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *Bar) GetClose() float64 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *Bar) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Bar) GetTrades() int64 {
	if m != nil {
		return m.Trades
	}
	return 0
}

type BarRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Grouping             string   `protobuf:"bytes,2,opt,name=grouping,proto3" json:"grouping,omitempty"`
	Product              string   `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
	Venue                string   `protobuf:"bytes,4,opt,name=venue,proto3" json:"venue,omitempty"`
	BeginDate            int64    `protobuf:"varint,5,opt,name=begin_date,json=beginDate,proto3" json:"begin_date,omitempty"`
	EndDate              int64    `protobuf:"varint,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BarRequest) Reset()         { *m = BarRequest{} }
func (m *BarRequest) String() string { return proto.CompactTextString(m) }
func (*BarRequest) ProtoMessage()    {}
func (*BarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{7}
}
func (m *BarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarRequest.Unmarshal(m, b)
}
func (m *BarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarRequest.Marshal(b, m, deterministic)
}
func (dst *BarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarRequest.Merge(dst, src)
}
func (m *BarRequest) XXX_Size() int {
	return xxx_messageInfo_BarRequest.Size(m)
}
func (m *BarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BarRequest proto.InternalMessageInfo

func (m *BarRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *BarRequest) GetGrouping() string {
	if m != nil {
		return m.Grouping
	}
	return ""
}

func (m *BarRequest) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *BarRequest) GetVenue() string {
	if m != nil {
		return m.Venue
	}
	return ""
}

func (m *BarRequest) GetBeginDate() int64 {
	if m != nil {
		return m.BeginDate
	}
	return 0
}

func (m *BarRequest) GetEndDate() int64 {
	if m != nil {
		return m.EndDate
	}
	return 0
}

type BarResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Bars                 []*Bar   `protobuf:"bytes,2,rep,name=bars,proto3" json:"bars,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BarResponse) Reset()         { *m = BarResponse{} }
func (m *BarResponse) String() string { return proto.CompactTextString(m) }
func (*BarResponse) ProtoMessage()    {}
func (*BarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{8}
}
func (m *BarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarResponse.Unmarshal(m, b)
}
func (m *BarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarResponse.Marshal(b, m, deterministic)
}
func (dst *BarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarResponse.Merge(dst, src)
}
func (m *BarResponse) XXX_Size() int {
	return xxx_messageInfo_BarResponse.Size(m)
}
func (m *BarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BarResponse proto.InternalMessageInfo

func (m *BarResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *BarResponse) GetBars() []*Bar {
	if m != nil {
		return m.Bars
	}
	return nil
}

func (m *BarResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *BarResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

func init() {
	proto.RegisterType((*MarketKindSummary)(nil), "api.MarketKindSummary")
	proto.RegisterType((*MarketKindSummaryRequest)(nil), "api.MarketKindSummaryRequest")
	proto.RegisterType((*MarketKindSummaryResponse)(nil), "api.MarketKindSummaryResponse")
	proto.RegisterType((*MarketProductSummary)(nil), "api.MarketProductSummary")
	proto.RegisterType((*MarketProductSummaryRequest)(nil), "api.MarketProductSummaryRequest")
	proto.RegisterType((*MarketProductSummaryResponse)(nil), "api.MarketProductSummaryResponse")
	proto.RegisterType((*Bar)(nil), "api.Bar")
	proto.RegisterType((*BarRequest)(nil), "api.BarRequest")
	proto.RegisterType((*BarResponse)(nil), "api.BarResponse")
}

func init() { proto.RegisterFile("market.proto", fileDescriptor_3f90997f23a2c3f8) }

var fileDescriptor_3f90997f23a2c3f8 = []byte{
	// 718 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x5d, 0x6a, 0x1b, 0x49,
	0x10, 0x66, 0x24, 0xeb, 0x67, 0xca, 0xc6, 0x3f, 0x6d, 0xad, 0x76, 0xec, 0xf5, 0x82, 0x77, 0x58,
	0x76, 0x1d, 0x02, 0x0e, 0x38, 0x21, 0x07, 0x30, 0x4a, 0xc0, 0x84, 0x90, 0x30, 0xf6, 0xfb, 0xd0,
	0x52, 0x17, 0x72, 0x63, 0xcd, 0xf4, 0xa4, 0xbb, 0xa5, 0xc4, 0x27, 0xc8, 0x19, 0x72, 0x83, 0xdc,
	0x22, 0x07, 0xc8, 0x5b, 0xee, 0x92, 0xf7, 0xd0, 0xd5, 0x3d, 0xb2, 0x64, 0x8b, 0x80, 0xfc, 0xd6,
	0xf5, 0x7d, 0x55, 0x5f, 0xb5, 0xaa, 0xbe, 0x69, 0xc1, 0x56, 0xc1, 0xf5, 0x0d, 0xda, 0xd3, 0x4a,
	0x2b, 0xab, 0x58, 0x93, 0x57, 0xf2, 0x10, 0xb0, 0x9c, 0x16, 0x1e, 0x48, 0x7f, 0x34, 0x60, 0xef,
	0x2d, 0x65, 0xbc, 0x91, 0xa5, 0xb8, 0x9c, 0x16, 0x05, 0xd7, 0xb7, 0xec, 0x5f, 0xd8, 0xf6, 0x65,
	0xb9, 0xe0, 0x96, 0xe7, 0x52, 0x24, 0xd1, 0x71, 0x74, 0xd2, 0xcc, 0x82, 0xd8, 0x80, 0x5b, 0x7e,
	0x21, 0x58, 0x02, 0x9d, 0x4a, 0x2b, 0x31, 0x1d, 0xd9, 0xa4, 0x71, 0x1c, 0x9d, 0xc4, 0x59, 0x1d,
	0xb2, 0x7f, 0x60, 0x6b, 0x86, 0xe5, 0x14, 0x4d, 0x6e, 0x95, 0xe5, 0x93, 0xa4, 0x49, 0xd5, 0x9b,
	0x1e, 0xbb, 0x72, 0x10, 0x7b, 0x01, 0xfd, 0x90, 0x6d, 0x72, 0x2c, 0xf9, 0x70, 0x82, 0x22, 0x24,
	0x6f, 0x50, 0x72, 0xaf, 0x66, 0x5f, 0x79, 0xd2, 0x57, 0xbd, 0x84, 0x3f, 0xe7, 0x55, 0x42, 0x9a,
	0xc5, 0xb2, 0x16, 0x95, 0xfd, 0x51, 0xd3, 0x83, 0xc0, 0xce, 0xbb, 0x19, 0xab, 0x91, 0x17, 0xb2,
	0x1c, 0xe7, 0x86, 0xcf, 0xb0, 0xee, 0x99, 0xb4, 0x7d, 0xb7, 0x39, 0x7b, 0xc9, 0x67, 0x18, 0x5a,
	0xba, 0x6e, 0xf7, 0xaa, 0xea, 0x9e, 0x49, 0xc7, 0x77, 0x5b, 0x2a, 0xab, 0x5b, 0xa6, 0x03, 0x48,
	0x1e, 0xcc, 0x34, 0xc3, 0x0f, 0x53, 0x34, 0x96, 0xf5, 0xa0, 0x65, 0xd5, 0x0d, 0x96, 0x34, 0xd1,
	0x38, 0xf3, 0x01, 0x63, 0xb0, 0x71, 0x23, 0x4b, 0x11, 0xe6, 0x48, 0xe7, 0xf4, 0x5b, 0x04, 0x07,
	0x2b, 0x64, 0x4c, 0xa5, 0x4a, 0x83, 0xec, 0x3f, 0xe8, 0x68, 0xb4, 0x23, 0x25, 0x90, 0x94, 0xb6,
	0xcf, 0xb6, 0x4e, 0x79, 0x25, 0x4f, 0x33, 0x8f, 0x65, 0x35, 0xc9, 0x5e, 0xc3, 0x7e, 0x58, 0xa5,
	0x13, 0xcd, 0x8d, 0x97, 0x49, 0x1a, 0xc7, 0xcd, 0x93, 0xcd, 0xb3, 0x3e, 0xd5, 0x3c, 0x6c, 0xb2,
	0x57, 0x3c, 0xb0, 0x44, 0x02, 0x9d, 0x91, 0x2a, 0x0a, 0x2c, 0x2d, 0x6d, 0x33, 0xce, 0xea, 0xd0,
	0x31, 0x38, 0xe1, 0x95, 0x41, 0x41, 0xab, 0x8b, 0xb3, 0x3a, 0x4c, 0x7f, 0x36, 0xa0, 0xe7, 0xc5,
	0xdf, 0xfb, 0xad, 0xac, 0xe7, 0xaf, 0xbf, 0x01, 0xc2, 0x36, 0x5d, 0x46, 0x83, 0x32, 0xe2, 0x80,
	0x5c, 0x08, 0xf6, 0x0c, 0xf6, 0x6b, 0x5a, 0xa0, 0x19, 0x69, 0x59, 0x59, 0xa9, 0xca, 0x70, 0x3b,
	0x16, 0xa8, 0xc1, 0x1d, 0xc3, 0x0e, 0xa0, 0x4b, 0x0e, 0x74, 0x6a, 0xde, 0x64, 0x1d, 0x8a, 0x2f,
	0x04, 0x7b, 0x0a, 0x7b, 0x9e, 0x5a, 0x54, 0x6a, 0x91, 0xd2, 0x2e, 0x11, 0x8b, 0x3a, 0x3d, 0x68,
	0x55, 0x5a, 0x8e, 0x90, 0xbc, 0x13, 0x65, 0x3e, 0x60, 0x7d, 0x68, 0xcf, 0xd4, 0x64, 0x5a, 0x20,
	0x79, 0x23, 0xca, 0x42, 0xe4, 0x7e, 0x85, 0xd2, 0x02, 0x75, 0x6e, 0xa4, 0xc0, 0xa4, 0x4b, 0x9a,
	0x31, 0x21, 0x97, 0x52, 0x20, 0x7b, 0x02, 0xbb, 0xe6, 0xd6, 0x58, 0x2c, 0x72, 0x2b, 0x0b, 0x34,
	0x96, 0x17, 0x55, 0x12, 0x53, 0xd2, 0x8e, 0xc7, 0xaf, 0x6a, 0x98, 0xfd, 0x0f, 0x3b, 0xfe, 0x92,
	0x77, 0x99, 0x40, 0x99, 0xdb, 0x04, 0xcf, 0x13, 0xd3, 0x4f, 0xf0, 0xd7, 0xaa, 0xb1, 0xaf, 0x6d,
	0xc1, 0xb5, 0x47, 0x9c, 0x7e, 0x8f, 0xe0, 0x68, 0x75, 0xeb, 0x35, 0x6d, 0xfb, 0x0e, 0xfa, 0xc1,
	0x21, 0xf5, 0x05, 0x96, 0x9d, 0x7b, 0xb0, 0xe0, 0xdc, 0x7b, 0xad, 0x7a, 0xc5, 0x2a, 0xcb, 0x3d,
	0xc6, 0xbf, 0x5f, 0x22, 0x68, 0x9e, 0x73, 0xed, 0x46, 0xe3, 0x46, 0x1e, 0x4c, 0x4a, 0x67, 0x87,
	0xa9, 0x0a, 0x4b, 0x1a, 0x57, 0x94, 0xd1, 0xd9, 0x61, 0xd7, 0x72, 0x7c, 0x4d, 0x0d, 0xa2, 0x8c,
	0xce, 0x6c, 0x17, 0x9a, 0x13, 0xf5, 0x91, 0x94, 0xa3, 0xcc, 0x1d, 0xdd, 0xf8, 0x47, 0x13, 0x65,
	0x90, 0xfc, 0x15, 0x65, 0x3e, 0x58, 0xb0, 0x4f, 0x7b, 0xc9, 0x3e, 0x7d, 0x68, 0x5b, 0xcd, 0x05,
	0x9a, 0xf0, 0xe4, 0x84, 0x28, 0xfd, 0x1a, 0x01, 0x9c, 0x73, 0xfd, 0xfb, 0x9d, 0x1e, 0x42, 0x77,
	0xac, 0xd5, 0xb4, 0x92, 0xe5, 0x38, 0xec, 0x75, 0x1e, 0x2f, 0xbe, 0xde, 0xcd, 0xe5, 0xd7, 0xbb,
	0x07, 0x2d, 0x32, 0x54, 0x18, 0x87, 0x0f, 0x9c, 0x8f, 0x87, 0x38, 0x96, 0xa5, 0xfb, 0x64, 0x31,
	0xbc, 0xb6, 0x31, 0x21, 0x03, 0x6e, 0xd1, 0x7d, 0x5c, 0x58, 0x0a, 0x4f, 0xfa, 0x37, 0xb5, 0x83,
	0xa5, 0x70, 0x54, 0xfa, 0x39, 0x82, 0x4d, 0xba, 0xea, 0x9a, 0x1e, 0x38, 0x82, 0x8d, 0x21, 0xd7,
	0x26, 0x6c, 0xbc, 0x4b, 0x49, 0x4e, 0x87, 0xd0, 0xc7, 0x2c, 0x74, 0xd8, 0xa6, 0x3f, 0xbd, 0xe7,
	0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xed, 0x96, 0x19, 0x16, 0x15, 0x07, 0x00, 0x00,
}
