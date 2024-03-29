// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statistics.proto

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ------------------------------------------------------------------- //
type BuySellWeek struct {
	WeekDay              string   `protobuf:"bytes,1,opt,name=week_day,json=weekDay,proto3" json:"week_day,omitempty"`
	Sell                 int64    `protobuf:"varint,2,opt,name=sell,proto3" json:"sell,omitempty"`
	Buy                  int64    `protobuf:"varint,3,opt,name=buy,proto3" json:"buy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuySellWeek) Reset()         { *m = BuySellWeek{} }
func (m *BuySellWeek) String() string { return proto.CompactTextString(m) }
func (*BuySellWeek) ProtoMessage()    {}
func (*BuySellWeek) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe4ebcdede33dbb6, []int{0}
}

func (m *BuySellWeek) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuySellWeek.Unmarshal(m, b)
}
func (m *BuySellWeek) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuySellWeek.Marshal(b, m, deterministic)
}
func (m *BuySellWeek) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuySellWeek.Merge(m, src)
}
func (m *BuySellWeek) XXX_Size() int {
	return xxx_messageInfo_BuySellWeek.Size(m)
}
func (m *BuySellWeek) XXX_DiscardUnknown() {
	xxx_messageInfo_BuySellWeek.DiscardUnknown(m)
}

var xxx_messageInfo_BuySellWeek proto.InternalMessageInfo

func (m *BuySellWeek) GetWeekDay() string {
	if m != nil {
		return m.WeekDay
	}
	return ""
}

func (m *BuySellWeek) GetSell() int64 {
	if m != nil {
		return m.Sell
	}
	return 0
}

func (m *BuySellWeek) GetBuy() int64 {
	if m != nil {
		return m.Buy
	}
	return 0
}

type BuySellWeekRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Product              string   `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Venue                string   `protobuf:"bytes,3,opt,name=venue,proto3" json:"venue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuySellWeekRequest) Reset()         { *m = BuySellWeekRequest{} }
func (m *BuySellWeekRequest) String() string { return proto.CompactTextString(m) }
func (*BuySellWeekRequest) ProtoMessage()    {}
func (*BuySellWeekRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe4ebcdede33dbb6, []int{1}
}

func (m *BuySellWeekRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuySellWeekRequest.Unmarshal(m, b)
}
func (m *BuySellWeekRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuySellWeekRequest.Marshal(b, m, deterministic)
}
func (m *BuySellWeekRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuySellWeekRequest.Merge(m, src)
}
func (m *BuySellWeekRequest) XXX_Size() int {
	return xxx_messageInfo_BuySellWeekRequest.Size(m)
}
func (m *BuySellWeekRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuySellWeekRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuySellWeekRequest proto.InternalMessageInfo

func (m *BuySellWeekRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *BuySellWeekRequest) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *BuySellWeekRequest) GetVenue() string {
	if m != nil {
		return m.Venue
	}
	return ""
}

type BuySellWeekResponse struct {
	Retcode              Retcode        `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	BuySellWeek          []*BuySellWeek `protobuf:"bytes,2,rep,name=buy_sell_week,json=buySellWeek,proto3" json:"buy_sell_week,omitempty"`
	Comment              string         `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string         `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BuySellWeekResponse) Reset()         { *m = BuySellWeekResponse{} }
func (m *BuySellWeekResponse) String() string { return proto.CompactTextString(m) }
func (*BuySellWeekResponse) ProtoMessage()    {}
func (*BuySellWeekResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe4ebcdede33dbb6, []int{2}
}

func (m *BuySellWeekResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuySellWeekResponse.Unmarshal(m, b)
}
func (m *BuySellWeekResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuySellWeekResponse.Marshal(b, m, deterministic)
}
func (m *BuySellWeekResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuySellWeekResponse.Merge(m, src)
}
func (m *BuySellWeekResponse) XXX_Size() int {
	return xxx_messageInfo_BuySellWeekResponse.Size(m)
}
func (m *BuySellWeekResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BuySellWeekResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BuySellWeekResponse proto.InternalMessageInfo

func (m *BuySellWeekResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_DONE
}

func (m *BuySellWeekResponse) GetBuySellWeek() []*BuySellWeek {
	if m != nil {
		return m.BuySellWeek
	}
	return nil
}

func (m *BuySellWeekResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *BuySellWeekResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type TradesDeadman struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductDescription   string   `protobuf:"bytes,2,opt,name=product_description,json=productDescription,proto3" json:"product_description,omitempty"`
	VenueId              int64    `protobuf:"varint,3,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	VenueDescription     string   `protobuf:"bytes,4,opt,name=venue_description,json=venueDescription,proto3" json:"venue_description,omitempty"`
	Price                float64  `protobuf:"fixed64,5,opt,name=price,proto3" json:"price,omitempty"`
	Volume               float64  `protobuf:"fixed64,6,opt,name=volume,proto3" json:"volume,omitempty"`
	Deadman              int64    `protobuf:"varint,7,opt,name=deadman,proto3" json:"deadman,omitempty"`
	Timestamp            string   `protobuf:"bytes,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TimeSince            string   `protobuf:"bytes,9,opt,name=time_since,json=timeSince,proto3" json:"time_since,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradesDeadman) Reset()         { *m = TradesDeadman{} }
func (m *TradesDeadman) String() string { return proto.CompactTextString(m) }
func (*TradesDeadman) ProtoMessage()    {}
func (*TradesDeadman) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe4ebcdede33dbb6, []int{3}
}

func (m *TradesDeadman) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradesDeadman.Unmarshal(m, b)
}
func (m *TradesDeadman) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradesDeadman.Marshal(b, m, deterministic)
}
func (m *TradesDeadman) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradesDeadman.Merge(m, src)
}
func (m *TradesDeadman) XXX_Size() int {
	return xxx_messageInfo_TradesDeadman.Size(m)
}
func (m *TradesDeadman) XXX_DiscardUnknown() {
	xxx_messageInfo_TradesDeadman.DiscardUnknown(m)
}

var xxx_messageInfo_TradesDeadman proto.InternalMessageInfo

func (m *TradesDeadman) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *TradesDeadman) GetProductDescription() string {
	if m != nil {
		return m.ProductDescription
	}
	return ""
}

func (m *TradesDeadman) GetVenueId() int64 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *TradesDeadman) GetVenueDescription() string {
	if m != nil {
		return m.VenueDescription
	}
	return ""
}

func (m *TradesDeadman) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *TradesDeadman) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *TradesDeadman) GetDeadman() int64 {
	if m != nil {
		return m.Deadman
	}
	return 0
}

func (m *TradesDeadman) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *TradesDeadman) GetTimeSince() string {
	if m != nil {
		return m.TimeSince
	}
	return ""
}

type TradesDeadmanRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradesDeadmanRequest) Reset()         { *m = TradesDeadmanRequest{} }
func (m *TradesDeadmanRequest) String() string { return proto.CompactTextString(m) }
func (*TradesDeadmanRequest) ProtoMessage()    {}
func (*TradesDeadmanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe4ebcdede33dbb6, []int{4}
}

func (m *TradesDeadmanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradesDeadmanRequest.Unmarshal(m, b)
}
func (m *TradesDeadmanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradesDeadmanRequest.Marshal(b, m, deterministic)
}
func (m *TradesDeadmanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradesDeadmanRequest.Merge(m, src)
}
func (m *TradesDeadmanRequest) XXX_Size() int {
	return xxx_messageInfo_TradesDeadmanRequest.Size(m)
}
func (m *TradesDeadmanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TradesDeadmanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TradesDeadmanRequest proto.InternalMessageInfo

func (m *TradesDeadmanRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type TradesDeadmanResponse struct {
	Retcode              Retcode          `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	TradesDeadman        []*TradesDeadman `protobuf:"bytes,2,rep,name=trades_deadman,json=tradesDeadman,proto3" json:"trades_deadman,omitempty"`
	Comment              string           `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string           `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TradesDeadmanResponse) Reset()         { *m = TradesDeadmanResponse{} }
func (m *TradesDeadmanResponse) String() string { return proto.CompactTextString(m) }
func (*TradesDeadmanResponse) ProtoMessage()    {}
func (*TradesDeadmanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe4ebcdede33dbb6, []int{5}
}

func (m *TradesDeadmanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradesDeadmanResponse.Unmarshal(m, b)
}
func (m *TradesDeadmanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradesDeadmanResponse.Marshal(b, m, deterministic)
}
func (m *TradesDeadmanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradesDeadmanResponse.Merge(m, src)
}
func (m *TradesDeadmanResponse) XXX_Size() int {
	return xxx_messageInfo_TradesDeadmanResponse.Size(m)
}
func (m *TradesDeadmanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TradesDeadmanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TradesDeadmanResponse proto.InternalMessageInfo

func (m *TradesDeadmanResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_DONE
}

func (m *TradesDeadmanResponse) GetTradesDeadman() []*TradesDeadman {
	if m != nil {
		return m.TradesDeadman
	}
	return nil
}

func (m *TradesDeadmanResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *TradesDeadmanResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type OrderbookDeadman struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	ProductDescription   string   `protobuf:"bytes,2,opt,name=product_description,json=productDescription,proto3" json:"product_description,omitempty"`
	VenueId              int64    `protobuf:"varint,3,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	VenueDescription     string   `protobuf:"bytes,4,opt,name=venue_description,json=venueDescription,proto3" json:"venue_description,omitempty"`
	Deadman              int64    `protobuf:"varint,5,opt,name=deadman,proto3" json:"deadman,omitempty"`
	Timestamp            string   `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TimeSince            string   `protobuf:"bytes,7,opt,name=time_since,json=timeSince,proto3" json:"time_since,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderbookDeadman) Reset()         { *m = OrderbookDeadman{} }
func (m *OrderbookDeadman) String() string { return proto.CompactTextString(m) }
func (*OrderbookDeadman) ProtoMessage()    {}
func (*OrderbookDeadman) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe4ebcdede33dbb6, []int{6}
}

func (m *OrderbookDeadman) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderbookDeadman.Unmarshal(m, b)
}
func (m *OrderbookDeadman) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderbookDeadman.Marshal(b, m, deterministic)
}
func (m *OrderbookDeadman) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderbookDeadman.Merge(m, src)
}
func (m *OrderbookDeadman) XXX_Size() int {
	return xxx_messageInfo_OrderbookDeadman.Size(m)
}
func (m *OrderbookDeadman) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderbookDeadman.DiscardUnknown(m)
}

var xxx_messageInfo_OrderbookDeadman proto.InternalMessageInfo

func (m *OrderbookDeadman) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *OrderbookDeadman) GetProductDescription() string {
	if m != nil {
		return m.ProductDescription
	}
	return ""
}

func (m *OrderbookDeadman) GetVenueId() int64 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *OrderbookDeadman) GetVenueDescription() string {
	if m != nil {
		return m.VenueDescription
	}
	return ""
}

func (m *OrderbookDeadman) GetDeadman() int64 {
	if m != nil {
		return m.Deadman
	}
	return 0
}

func (m *OrderbookDeadman) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *OrderbookDeadman) GetTimeSince() string {
	if m != nil {
		return m.TimeSince
	}
	return ""
}

type OrderbookDeadmanRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderbookDeadmanRequest) Reset()         { *m = OrderbookDeadmanRequest{} }
func (m *OrderbookDeadmanRequest) String() string { return proto.CompactTextString(m) }
func (*OrderbookDeadmanRequest) ProtoMessage()    {}
func (*OrderbookDeadmanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe4ebcdede33dbb6, []int{7}
}

func (m *OrderbookDeadmanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderbookDeadmanRequest.Unmarshal(m, b)
}
func (m *OrderbookDeadmanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderbookDeadmanRequest.Marshal(b, m, deterministic)
}
func (m *OrderbookDeadmanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderbookDeadmanRequest.Merge(m, src)
}
func (m *OrderbookDeadmanRequest) XXX_Size() int {
	return xxx_messageInfo_OrderbookDeadmanRequest.Size(m)
}
func (m *OrderbookDeadmanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderbookDeadmanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderbookDeadmanRequest proto.InternalMessageInfo

func (m *OrderbookDeadmanRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type OrderbookDeadmanResponse struct {
	Retcode              Retcode             `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	OrderbookDeadman     []*OrderbookDeadman `protobuf:"bytes,2,rep,name=orderbook_deadman,json=orderbookDeadman,proto3" json:"orderbook_deadman,omitempty"`
	Comment              string              `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string              `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *OrderbookDeadmanResponse) Reset()         { *m = OrderbookDeadmanResponse{} }
func (m *OrderbookDeadmanResponse) String() string { return proto.CompactTextString(m) }
func (*OrderbookDeadmanResponse) ProtoMessage()    {}
func (*OrderbookDeadmanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe4ebcdede33dbb6, []int{8}
}

func (m *OrderbookDeadmanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderbookDeadmanResponse.Unmarshal(m, b)
}
func (m *OrderbookDeadmanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderbookDeadmanResponse.Marshal(b, m, deterministic)
}
func (m *OrderbookDeadmanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderbookDeadmanResponse.Merge(m, src)
}
func (m *OrderbookDeadmanResponse) XXX_Size() int {
	return xxx_messageInfo_OrderbookDeadmanResponse.Size(m)
}
func (m *OrderbookDeadmanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderbookDeadmanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderbookDeadmanResponse proto.InternalMessageInfo

func (m *OrderbookDeadmanResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_DONE
}

func (m *OrderbookDeadmanResponse) GetOrderbookDeadman() []*OrderbookDeadman {
	if m != nil {
		return m.OrderbookDeadman
	}
	return nil
}

func (m *OrderbookDeadmanResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *OrderbookDeadmanResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

func init() {
	proto.RegisterType((*BuySellWeek)(nil), "api.BuySellWeek")
	proto.RegisterType((*BuySellWeekRequest)(nil), "api.BuySellWeekRequest")
	proto.RegisterType((*BuySellWeekResponse)(nil), "api.BuySellWeekResponse")
	proto.RegisterType((*TradesDeadman)(nil), "api.TradesDeadman")
	proto.RegisterType((*TradesDeadmanRequest)(nil), "api.TradesDeadmanRequest")
	proto.RegisterType((*TradesDeadmanResponse)(nil), "api.TradesDeadmanResponse")
	proto.RegisterType((*OrderbookDeadman)(nil), "api.OrderbookDeadman")
	proto.RegisterType((*OrderbookDeadmanRequest)(nil), "api.OrderbookDeadmanRequest")
	proto.RegisterType((*OrderbookDeadmanResponse)(nil), "api.OrderbookDeadmanResponse")
}

func init() { proto.RegisterFile("statistics.proto", fileDescriptor_fe4ebcdede33dbb6) }

var fileDescriptor_fe4ebcdede33dbb6 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0x41, 0x6b, 0xdb, 0x4c,
	0x10, 0x45, 0x56, 0x6c, 0x45, 0xe3, 0xcf, 0x41, 0xd9, 0x38, 0x5f, 0xb7, 0xa5, 0x05, 0xa3, 0x43,
	0x31, 0xb4, 0x38, 0x90, 0xf6, 0xd2, 0x6b, 0xf0, 0x25, 0x97, 0x16, 0x94, 0x42, 0xa1, 0x17, 0xb1,
	0xd6, 0xce, 0x61, 0xb1, 0xa4, 0x55, 0xb5, 0xab, 0x14, 0x1f, 0xfb, 0x6b, 0x7a, 0x29, 0xfd, 0x01,
	0xfd, 0x75, 0x65, 0x77, 0x25, 0x2c, 0x2b, 0x60, 0x48, 0x4e, 0xbd, 0xe9, 0xbd, 0xb7, 0xfb, 0x76,
	0x66, 0xde, 0x20, 0x88, 0x94, 0x66, 0x5a, 0x28, 0x2d, 0x32, 0xb5, 0xaa, 0x6a, 0xa9, 0x25, 0xf1,
	0x59, 0x25, 0x5e, 0x00, 0x96, 0x4d, 0xe1, 0x88, 0xf8, 0x23, 0x4c, 0x6f, 0x9a, 0xdd, 0x1d, 0xe6,
	0xf9, 0x17, 0xc4, 0x2d, 0x79, 0x0e, 0xa7, 0xdf, 0x11, 0xb7, 0x29, 0x67, 0x3b, 0xea, 0x2d, 0xbc,
	0x65, 0x98, 0x04, 0x06, 0xaf, 0xd9, 0x8e, 0x10, 0x38, 0x51, 0x98, 0xe7, 0x74, 0xb4, 0xf0, 0x96,
	0x7e, 0x62, 0xbf, 0x49, 0x04, 0xfe, 0xa6, 0xd9, 0x51, 0xdf, 0x52, 0xe6, 0x33, 0xfe, 0x0a, 0xa4,
	0xe7, 0x97, 0xe0, 0xb7, 0x06, 0x95, 0x26, 0x73, 0x18, 0x6b, 0xb9, 0xc5, 0xb2, 0xf5, 0x74, 0x80,
	0x50, 0x08, 0xaa, 0x5a, 0xf2, 0x26, 0xd3, 0xd6, 0x34, 0x4c, 0x3a, 0x68, 0xce, 0xdf, 0x63, 0xd9,
	0xa0, 0x75, 0x0e, 0x13, 0x07, 0xe2, 0x9f, 0x1e, 0x5c, 0x1c, 0x98, 0xab, 0x4a, 0x96, 0x0a, 0xc9,
	0x6b, 0x08, 0x6a, 0xd4, 0x99, 0xe4, 0x68, 0xfd, 0xcf, 0xae, 0xff, 0x5b, 0xb1, 0x4a, 0xac, 0x12,
	0xc7, 0x25, 0x9d, 0x48, 0xde, 0xc3, 0x6c, 0xd3, 0xec, 0x52, 0x53, 0x79, 0x6a, 0xba, 0xa2, 0xa3,
	0x85, 0xbf, 0x9c, 0x5e, 0x47, 0xf6, 0x74, 0xdf, 0x78, 0xba, 0xe9, 0x8d, 0x84, 0x42, 0x90, 0xc9,
	0xa2, 0xc0, 0x52, 0xb7, 0xd5, 0x74, 0xd0, 0x28, 0x98, 0xb3, 0x4a, 0x21, 0xa7, 0x27, 0x4e, 0x69,
	0x61, 0xfc, 0x6b, 0x04, 0xb3, 0xcf, 0x35, 0xe3, 0xa8, 0xd6, 0xc8, 0x78, 0xc1, 0x4a, 0xf2, 0x0a,
	0xa0, 0x6d, 0x2e, 0x15, 0xdc, 0x96, 0xe9, 0x27, 0x61, 0xcb, 0xdc, 0x72, 0x72, 0x05, 0x17, 0x9d,
	0xcc, 0x51, 0x65, 0xb5, 0xa8, 0xb4, 0x90, 0x65, 0x3b, 0x16, 0xd2, 0x4a, 0xeb, 0xbd, 0x62, 0x82,
	0xb2, 0x43, 0x31, 0x6e, 0x6e, 0xfc, 0x81, 0xc5, 0xb7, 0x9c, 0xbc, 0x81, 0x73, 0x27, 0xf5, 0x9d,
	0x5c, 0x81, 0x91, 0x15, 0xfa, 0x3e, 0x73, 0x18, 0x57, 0xb5, 0xc8, 0x90, 0x8e, 0x17, 0xde, 0xd2,
	0x4b, 0x1c, 0x20, 0xff, 0xc3, 0xe4, 0x5e, 0xe6, 0x4d, 0x81, 0x74, 0x62, 0xe9, 0x16, 0x99, 0x8e,
	0xb9, 0x6b, 0x88, 0x06, 0xee, 0xd1, 0x16, 0x92, 0x97, 0x10, 0x6a, 0x51, 0xa0, 0xd2, 0xac, 0xa8,
	0xe8, 0xa9, 0x7d, 0x6c, 0x4f, 0x98, 0xee, 0x0d, 0x48, 0x95, 0x28, 0x33, 0xa4, 0xe1, 0x5e, 0xbe,
	0x33, 0x44, 0xfc, 0x16, 0xe6, 0x07, 0xd3, 0x3a, 0xba, 0x36, 0xf1, 0x6f, 0x0f, 0x2e, 0x07, 0xc7,
	0x1f, 0xb9, 0x08, 0x1f, 0xe0, 0x4c, 0x5b, 0x83, 0xb4, 0xeb, 0xc6, 0x6d, 0x02, 0xb1, 0xc7, 0x0f,
	0xbd, 0x67, 0xfa, 0x20, 0xc7, 0xa7, 0x6c, 0xc3, 0x8f, 0x11, 0x44, 0x9f, 0x6a, 0x8e, 0xf5, 0x46,
	0xca, 0xed, 0xbf, 0xbe, 0x10, 0xbd, 0x88, 0xc7, 0x47, 0x22, 0x9e, 0x1c, 0x8f, 0x38, 0x18, 0x46,
	0x7c, 0x05, 0xcf, 0x86, 0x23, 0x38, 0x9e, 0xf2, 0x1f, 0x0f, 0xe8, 0xc3, 0x1b, 0x8f, 0x0c, 0xfa,
	0x06, 0xce, 0x65, 0xe7, 0x31, 0xc8, 0xfa, 0xd2, 0xde, 0x78, 0xf0, 0x42, 0x24, 0x87, 0x41, 0x3d,
	0x21, 0xf1, 0xcd, 0xc4, 0xfe, 0x5c, 0xdf, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x31, 0x09, 0x35,
	0xb4, 0x81, 0x05, 0x00, 0x00,
}
