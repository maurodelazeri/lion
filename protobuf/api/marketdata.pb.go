// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketdata.proto

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

// Market Data
type Item struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Volume               float64  `protobuf:"fixed64,2,opt,name=volume,proto3" json:"volume,omitempty"`
	Price                float64  `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6b85761666e9c8c, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Item) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

type Orderbook struct {
	Product              string    `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Venue                Venues    `protobuf:"varint,2,opt,name=venue,proto3,enum=api.Venues" json:"venue,omitempty"`
	Timestamp            int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Levels               int32     `protobuf:"varint,4,opt,name=levels,proto3" json:"levels,omitempty"`
	Asks                 []*Item   `protobuf:"bytes,5,rep,name=asks,proto3" json:"asks,omitempty"`
	Bids                 []*Item   `protobuf:"bytes,6,rep,name=bids,proto3" json:"bids,omitempty"`
	VenueType            VenueType `protobuf:"varint,7,opt,name=venue_type,json=venueType,proto3,enum=api.VenueType" json:"venue_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Orderbook) Reset()         { *m = Orderbook{} }
func (m *Orderbook) String() string { return proto.CompactTextString(m) }
func (*Orderbook) ProtoMessage()    {}
func (*Orderbook) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6b85761666e9c8c, []int{1}
}

func (m *Orderbook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Orderbook.Unmarshal(m, b)
}
func (m *Orderbook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Orderbook.Marshal(b, m, deterministic)
}
func (m *Orderbook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Orderbook.Merge(m, src)
}
func (m *Orderbook) XXX_Size() int {
	return xxx_messageInfo_Orderbook.Size(m)
}
func (m *Orderbook) XXX_DiscardUnknown() {
	xxx_messageInfo_Orderbook.DiscardUnknown(m)
}

var xxx_messageInfo_Orderbook proto.InternalMessageInfo

func (m *Orderbook) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *Orderbook) GetVenue() Venues {
	if m != nil {
		return m.Venue
	}
	return Venues_DARKPOOL
}

func (m *Orderbook) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Orderbook) GetLevels() int32 {
	if m != nil {
		return m.Levels
	}
	return 0
}

func (m *Orderbook) GetAsks() []*Item {
	if m != nil {
		return m.Asks
	}
	return nil
}

func (m *Orderbook) GetBids() []*Item {
	if m != nil {
		return m.Bids
	}
	return nil
}

func (m *Orderbook) GetVenueType() VenueType {
	if m != nil {
		return m.VenueType
	}
	return VenueType_SPOT
}

type Trade struct {
	Product              string    `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Venue                Venues    `protobuf:"varint,2,opt,name=venue,proto3,enum=api.Venues" json:"venue,omitempty"`
	Timestamp            int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Price                float64   `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Volume               float64   `protobuf:"fixed64,5,opt,name=volume,proto3" json:"volume,omitempty"`
	OrderSide            Side      `protobuf:"varint,6,opt,name=order_side,json=orderSide,proto3,enum=api.Side" json:"order_side,omitempty"`
	VenueType            VenueType `protobuf:"varint,7,opt,name=venue_type,json=venueType,proto3,enum=api.VenueType" json:"venue_type,omitempty"`
	Asks                 []*Item   `protobuf:"bytes,8,rep,name=asks,proto3" json:"asks,omitempty"`
	Bids                 []*Item   `protobuf:"bytes,9,rep,name=bids,proto3" json:"bids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Trade) Reset()         { *m = Trade{} }
func (m *Trade) String() string { return proto.CompactTextString(m) }
func (*Trade) ProtoMessage()    {}
func (*Trade) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6b85761666e9c8c, []int{2}
}

func (m *Trade) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trade.Unmarshal(m, b)
}
func (m *Trade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trade.Marshal(b, m, deterministic)
}
func (m *Trade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trade.Merge(m, src)
}
func (m *Trade) XXX_Size() int {
	return xxx_messageInfo_Trade.Size(m)
}
func (m *Trade) XXX_DiscardUnknown() {
	xxx_messageInfo_Trade.DiscardUnknown(m)
}

var xxx_messageInfo_Trade proto.InternalMessageInfo

func (m *Trade) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *Trade) GetVenue() Venues {
	if m != nil {
		return m.Venue
	}
	return Venues_DARKPOOL
}

func (m *Trade) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Trade) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Trade) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Trade) GetOrderSide() Side {
	if m != nil {
		return m.OrderSide
	}
	return Side_BUY
}

func (m *Trade) GetVenueType() VenueType {
	if m != nil {
		return m.VenueType
	}
	return VenueType_SPOT
}

func (m *Trade) GetAsks() []*Item {
	if m != nil {
		return m.Asks
	}
	return nil
}

func (m *Trade) GetBids() []*Item {
	if m != nil {
		return m.Bids
	}
	return nil
}

type Ticker struct {
	Product              string    `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Venue                Venues    `protobuf:"varint,2,opt,name=venue,proto3,enum=api.Venues" json:"venue,omitempty"`
	Timestamp            int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Price                float64   `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	OrderSide            Side      `protobuf:"varint,5,opt,name=order_side,json=orderSide,proto3,enum=api.Side" json:"order_side,omitempty"`
	BestBid              float64   `protobuf:"fixed64,6,opt,name=best_bid,json=bestBid,proto3" json:"best_bid,omitempty"`
	BestAsk              float64   `protobuf:"fixed64,7,opt,name=best_ask,json=bestAsk,proto3" json:"best_ask,omitempty"`
	VenueType            VenueType `protobuf:"varint,8,opt,name=venue_type,json=venueType,proto3,enum=api.VenueType" json:"venue_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Ticker) Reset()         { *m = Ticker{} }
func (m *Ticker) String() string { return proto.CompactTextString(m) }
func (*Ticker) ProtoMessage()    {}
func (*Ticker) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6b85761666e9c8c, []int{3}
}

func (m *Ticker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ticker.Unmarshal(m, b)
}
func (m *Ticker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ticker.Marshal(b, m, deterministic)
}
func (m *Ticker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ticker.Merge(m, src)
}
func (m *Ticker) XXX_Size() int {
	return xxx_messageInfo_Ticker.Size(m)
}
func (m *Ticker) XXX_DiscardUnknown() {
	xxx_messageInfo_Ticker.DiscardUnknown(m)
}

var xxx_messageInfo_Ticker proto.InternalMessageInfo

func (m *Ticker) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *Ticker) GetVenue() Venues {
	if m != nil {
		return m.Venue
	}
	return Venues_DARKPOOL
}

func (m *Ticker) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Ticker) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Ticker) GetOrderSide() Side {
	if m != nil {
		return m.OrderSide
	}
	return Side_BUY
}

func (m *Ticker) GetBestBid() float64 {
	if m != nil {
		return m.BestBid
	}
	return 0
}

func (m *Ticker) GetBestAsk() float64 {
	if m != nil {
		return m.BestAsk
	}
	return 0
}

func (m *Ticker) GetVenueType() VenueType {
	if m != nil {
		return m.VenueType
	}
	return VenueType_SPOT
}

type Candle struct {
	Venue                Venues   `protobuf:"varint,1,opt,name=venue,proto3,enum=api.Venues" json:"venue,omitempty"`
	Product              string   `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Granularity          int32    `protobuf:"varint,3,opt,name=granularity,proto3" json:"granularity,omitempty"`
	Point                int64    `protobuf:"varint,4,opt,name=point,proto3" json:"point,omitempty"`
	Open                 float64  `protobuf:"fixed64,5,opt,name=open,proto3" json:"open,omitempty"`
	Close                float64  `protobuf:"fixed64,6,opt,name=close,proto3" json:"close,omitempty"`
	High                 float64  `protobuf:"fixed64,7,opt,name=high,proto3" json:"high,omitempty"`
	Low                  float64  `protobuf:"fixed64,8,opt,name=low,proto3" json:"low,omitempty"`
	Volume               float64  `protobuf:"fixed64,9,opt,name=volume,proto3" json:"volume,omitempty"`
	Total                float64  `protobuf:"fixed64,10,opt,name=total,proto3" json:"total,omitempty"`
	TotalTrades          int32    `protobuf:"varint,11,opt,name=total_trades,json=totalTrades,proto3" json:"total_trades,omitempty"`
	BuyTotal             int32    `protobuf:"varint,12,opt,name=buy_total,json=buyTotal,proto3" json:"buy_total,omitempty"`
	SellTotal            int32    `protobuf:"varint,13,opt,name=sell_total,json=sellTotal,proto3" json:"sell_total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Candle) Reset()         { *m = Candle{} }
func (m *Candle) String() string { return proto.CompactTextString(m) }
func (*Candle) ProtoMessage()    {}
func (*Candle) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6b85761666e9c8c, []int{4}
}

func (m *Candle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Candle.Unmarshal(m, b)
}
func (m *Candle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Candle.Marshal(b, m, deterministic)
}
func (m *Candle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Candle.Merge(m, src)
}
func (m *Candle) XXX_Size() int {
	return xxx_messageInfo_Candle.Size(m)
}
func (m *Candle) XXX_DiscardUnknown() {
	xxx_messageInfo_Candle.DiscardUnknown(m)
}

var xxx_messageInfo_Candle proto.InternalMessageInfo

func (m *Candle) GetVenue() Venues {
	if m != nil {
		return m.Venue
	}
	return Venues_DARKPOOL
}

func (m *Candle) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *Candle) GetGranularity() int32 {
	if m != nil {
		return m.Granularity
	}
	return 0
}

func (m *Candle) GetPoint() int64 {
	if m != nil {
		return m.Point
	}
	return 0
}

func (m *Candle) GetOpen() float64 {
	if m != nil {
		return m.Open
	}
	return 0
}

func (m *Candle) GetClose() float64 {
	if m != nil {
		return m.Close
	}
	return 0
}

func (m *Candle) GetHigh() float64 {
	if m != nil {
		return m.High
	}
	return 0
}

func (m *Candle) GetLow() float64 {
	if m != nil {
		return m.Low
	}
	return 0
}

func (m *Candle) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Candle) GetTotal() float64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Candle) GetTotalTrades() int32 {
	if m != nil {
		return m.TotalTrades
	}
	return 0
}

func (m *Candle) GetBuyTotal() int32 {
	if m != nil {
		return m.BuyTotal
	}
	return 0
}

func (m *Candle) GetSellTotal() int32 {
	if m != nil {
		return m.SellTotal
	}
	return 0
}

func init() {
	proto.RegisterType((*Item)(nil), "api.Item")
	proto.RegisterType((*Orderbook)(nil), "api.Orderbook")
	proto.RegisterType((*Trade)(nil), "api.Trade")
	proto.RegisterType((*Ticker)(nil), "api.Ticker")
	proto.RegisterType((*Candle)(nil), "api.Candle")
}

func init() { proto.RegisterFile("marketdata.proto", fileDescriptor_e6b85761666e9c8c) }

var fileDescriptor_e6b85761666e9c8c = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0x95, 0xa4, 0xc9, 0x34, 0xa7, 0x43, 0x35, 0xb2, 0x10, 0x32, 0x97, 0x91, 0x3a, 0x5d,
	0x75, 0x43, 0x17, 0xc3, 0x13, 0x70, 0xd9, 0xb0, 0x42, 0x32, 0x15, 0xdb, 0xca, 0xa9, 0x8f, 0x66,
	0xac, 0x5c, 0x1c, 0xc5, 0x4e, 0x51, 0xdf, 0x80, 0x37, 0xe1, 0x9d, 0x78, 0x05, 0x5e, 0x02, 0xf9,
	0x38, 0xd3, 0xcb, 0x00, 0x83, 0x58, 0xc0, 0xee, 0x9c, 0xff, 0x4f, 0xec, 0xf3, 0x7f, 0x71, 0x0c,
	0x17, 0xb5, 0xec, 0x4a, 0x74, 0x4a, 0x3a, 0xb9, 0x6c, 0x3b, 0xe3, 0x0c, 0x4b, 0x64, 0xab, 0x9f,
	0x01, 0x36, 0x7d, 0x1d, 0x84, 0xf9, 0x3b, 0x18, 0xbd, 0x77, 0x58, 0xb3, 0x29, 0xc4, 0x5a, 0xf1,
	0x68, 0x16, 0x2d, 0x12, 0x11, 0x6b, 0xc5, 0x9e, 0x40, 0xb6, 0x35, 0x55, 0x5f, 0x23, 0x8f, 0x67,
	0xd1, 0x22, 0x12, 0x43, 0xc7, 0x1e, 0x43, 0xda, 0x76, 0x7a, 0x83, 0x3c, 0x21, 0x39, 0x34, 0xf3,
	0xef, 0x11, 0xe4, 0x1f, 0x3a, 0x85, 0x5d, 0x61, 0x4c, 0xc9, 0x38, 0x9c, 0xb5, 0x9d, 0x51, 0xfd,
	0xc6, 0xd1, 0x82, 0xb9, 0xb8, 0x6b, 0xd9, 0x15, 0xa4, 0x5b, 0x6c, 0xfa, 0xb0, 0xe8, 0xf4, 0x7a,
	0xb2, 0x94, 0xad, 0x5e, 0x7e, 0xf2, 0x8a, 0x15, 0xc1, 0x61, 0x2f, 0x20, 0x77, 0xba, 0x46, 0xeb,
	0x64, 0xdd, 0xd2, 0x26, 0x89, 0x38, 0x08, 0x7e, 0xac, 0x0a, 0xb7, 0x58, 0x59, 0x3e, 0x9a, 0x45,
	0x8b, 0x54, 0x0c, 0x1d, 0xbb, 0x84, 0x91, 0xb4, 0xa5, 0xe5, 0xe9, 0x2c, 0x59, 0x4c, 0xae, 0x73,
	0x5a, 0xd7, 0xe7, 0x12, 0x24, 0x7b, 0xbb, 0xd0, 0xca, 0xf2, 0xec, 0x27, 0xdb, 0xcb, 0xec, 0x25,
	0x00, 0x6d, 0xbe, 0x76, 0xbb, 0x16, 0xf9, 0x19, 0xcd, 0x36, 0x3d, 0xcc, 0xb6, 0xda, 0xb5, 0x28,
	0xf2, 0xed, 0x5d, 0x39, 0xff, 0x1a, 0x43, 0xba, 0xea, 0xa4, 0xc2, 0x7f, 0x99, 0x74, 0x0f, 0x7a,
	0x74, 0x04, 0xfa, 0xe8, 0xb3, 0xa4, 0x27, 0x9f, 0x65, 0x01, 0x60, 0x3c, 0xff, 0xb5, 0xd5, 0x0a,
	0x79, 0x46, 0x7b, 0x86, 0x98, 0x1f, 0xb5, 0x42, 0x91, 0x93, 0xe9, 0xcb, 0xbf, 0xcc, 0xba, 0x07,
	0x3b, 0x7e, 0x18, 0x6c, 0xfe, 0x4b, 0xb0, 0xf3, 0x2f, 0x31, 0x64, 0x2b, 0xbd, 0x29, 0xb1, 0xfb,
	0xff, 0xa8, 0x4e, 0x91, 0xa4, 0x0f, 0x20, 0x79, 0x0a, 0xe3, 0x02, 0xad, 0x5b, 0x17, 0x5a, 0x11,
	0xba, 0x48, 0x9c, 0xf9, 0xfe, 0x8d, 0x56, 0x7b, 0x4b, 0xda, 0x92, 0x58, 0x0d, 0xd6, 0x6b, 0x5b,
	0xde, 0x03, 0x39, 0xfe, 0xd3, 0xa1, 0xf9, 0x16, 0x43, 0xf6, 0x56, 0x36, 0xaa, 0xc2, 0x43, 0xe0,
	0xe8, 0xb7, 0x81, 0x8f, 0x68, 0xc5, 0xa7, 0xb4, 0x66, 0x30, 0xb9, 0xe9, 0x64, 0xd3, 0x57, 0xb2,
	0xd3, 0x6e, 0x47, 0x30, 0x52, 0x71, 0x2c, 0x11, 0x0e, 0xa3, 0x1b, 0x47, 0x38, 0x12, 0x11, 0x1a,
	0xc6, 0x60, 0x64, 0x5a, 0x6c, 0x86, 0x73, 0x43, 0xb5, 0x7f, 0x72, 0x53, 0x19, 0x8b, 0x43, 0xea,
	0xd0, 0xf8, 0x27, 0x6f, 0xf5, 0xcd, 0xed, 0x90, 0x97, 0x6a, 0x76, 0x01, 0x49, 0x65, 0x3e, 0x53,
	0xca, 0x48, 0xf8, 0xf2, 0xe8, 0x24, 0xe6, 0xf7, 0x2f, 0x08, 0x67, 0x9c, 0xac, 0x38, 0x84, 0x35,
	0xa9, 0x61, 0x57, 0x70, 0x4e, 0xc5, 0xda, 0xf9, 0xff, 0xc6, 0xf2, 0x49, 0x18, 0x9b, 0x34, 0xfa,
	0x95, 0x2c, 0x7b, 0x0e, 0x79, 0xd1, 0xef, 0xd6, 0xe1, 0xe5, 0x73, 0xf2, 0xc7, 0x45, 0xbf, 0x5b,
	0xd1, 0xfb, 0x97, 0x00, 0x16, 0xab, 0x6a, 0x70, 0x1f, 0x91, 0x9b, 0x7b, 0x85, 0xec, 0x22, 0xa3,
	0xcb, 0xec, 0xd5, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0xe9, 0x40, 0xe0, 0xf1, 0x04, 0x00,
	0x00,
}
