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
	Product              Product   `protobuf:"varint,1,opt,name=product,proto3,enum=api.Product" json:"product,omitempty"`
	Venue                Venue     `protobuf:"varint,2,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
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

func (m *Orderbook) GetProduct() Product {
	if m != nil {
		return m.Product
	}
	return Product_BTC_USD
}

func (m *Orderbook) GetVenue() Venue {
	if m != nil {
		return m.Venue
	}
	return Venue_DARKPOOL
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
	Product              Product   `protobuf:"varint,1,opt,name=product,proto3,enum=api.Product" json:"product,omitempty"`
	Venue                Venue     `protobuf:"varint,2,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
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

func (m *Trade) GetProduct() Product {
	if m != nil {
		return m.Product
	}
	return Product_BTC_USD
}

func (m *Trade) GetVenue() Venue {
	if m != nil {
		return m.Venue
	}
	return Venue_DARKPOOL
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
	Product              Product   `protobuf:"varint,1,opt,name=product,proto3,enum=api.Product" json:"product,omitempty"`
	Venue                Venue     `protobuf:"varint,2,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
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

func (m *Ticker) GetProduct() Product {
	if m != nil {
		return m.Product
	}
	return Product_BTC_USD
}

func (m *Ticker) GetVenue() Venue {
	if m != nil {
		return m.Venue
	}
	return Venue_DARKPOOL
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
	Venue                Venue    `protobuf:"varint,1,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
	Product              Product  `protobuf:"varint,2,opt,name=product,proto3,enum=api.Product" json:"product,omitempty"`
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

func (m *Candle) GetVenue() Venue {
	if m != nil {
		return m.Venue
	}
	return Venue_DARKPOOL
}

func (m *Candle) GetProduct() Product {
	if m != nil {
		return m.Product
	}
	return Product_BTC_USD
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
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0x92, 0x26, 0x6d, 0xa6, 0xa5, 0x5a, 0x59, 0x08, 0x85, 0x9f, 0x95, 0x4a, 0x0f, 0xa8,
	0x17, 0x7a, 0x28, 0x4f, 0xc0, 0xcf, 0x85, 0x13, 0xc8, 0x54, 0x5c, 0x2b, 0xa7, 0x1e, 0xed, 0x5a,
	0xf9, 0x71, 0x14, 0x3b, 0x45, 0x7d, 0x10, 0xde, 0x84, 0x77, 0xe2, 0x15, 0x38, 0x22, 0x8f, 0xb3,
	0x6d, 0x77, 0x81, 0x22, 0x2e, 0xec, 0x6d, 0xe6, 0xfb, 0x6c, 0xcf, 0x7c, 0xdf, 0xd8, 0x86, 0x8b,
	0x4a, 0xb4, 0x05, 0x5a, 0x29, 0xac, 0x58, 0x36, 0xad, 0xb6, 0x9a, 0x45, 0xa2, 0x51, 0x4f, 0x00,
	0xeb, 0xae, 0xf2, 0xc0, 0xfc, 0x1d, 0x0c, 0xde, 0x5b, 0xac, 0xd8, 0x14, 0x42, 0x25, 0xb3, 0x60,
	0x16, 0x2c, 0x22, 0x1e, 0x2a, 0xc9, 0x1e, 0x41, 0xb2, 0xd3, 0x65, 0x57, 0x61, 0x16, 0xce, 0x82,
	0x45, 0xc0, 0xfb, 0x8c, 0x3d, 0x84, 0xb8, 0x69, 0xd5, 0x16, 0xb3, 0x88, 0x60, 0x9f, 0xcc, 0x7f,
	0x04, 0x90, 0x7e, 0x68, 0x25, 0xb6, 0xb9, 0xd6, 0x05, 0x7b, 0x01, 0xc3, 0xa6, 0xd5, 0xb2, 0xdb,
	0x5a, 0x3a, 0x70, 0xba, 0x9a, 0x2c, 0x45, 0xa3, 0x96, 0x1f, 0x3d, 0xc6, 0x6f, 0x48, 0x36, 0x83,
	0x78, 0x87, 0x75, 0xe7, 0x4b, 0x4c, 0x57, 0x40, 0xab, 0x3e, 0x3b, 0x84, 0x7b, 0x82, 0x3d, 0x83,
	0xd4, 0xaa, 0x0a, 0x8d, 0x15, 0x55, 0x43, 0x15, 0x23, 0x7e, 0x04, 0x5c, 0x8f, 0x25, 0xee, 0xb0,
	0x34, 0xd9, 0x60, 0x16, 0x2c, 0x62, 0xde, 0x67, 0xec, 0x12, 0x06, 0xc2, 0x14, 0x26, 0x8b, 0x67,
	0xd1, 0x62, 0xbc, 0x4a, 0xe9, 0x58, 0x27, 0x92, 0x13, 0xec, 0xe8, 0x5c, 0x49, 0x93, 0x25, 0xbf,
	0xd0, 0x0e, 0x66, 0x2f, 0x01, 0xa8, 0xf8, 0xc6, 0xee, 0x1b, 0xcc, 0x86, 0xd4, 0xda, 0xf4, 0xd8,
	0xda, 0x7a, 0xdf, 0x20, 0x4f, 0x77, 0x37, 0xe1, 0xfc, 0x5b, 0x08, 0xf1, 0xba, 0x15, 0x12, 0xff,
	0x9b, 0xec, 0xc3, 0x08, 0x06, 0x27, 0x23, 0x38, 0x19, 0x58, 0x7c, 0x6b, 0x60, 0x0b, 0x00, 0xed,
	0x26, 0xb3, 0x31, 0x4a, 0x62, 0x96, 0x50, 0x49, 0xaf, 0xf9, 0x93, 0x92, 0xc8, 0x53, 0x22, 0x5d,
	0xf8, 0x8f, 0xc2, 0x0f, 0x2e, 0x8f, 0xce, 0xbb, 0x9c, 0xfe, 0xd6, 0xe5, 0xf9, 0xd7, 0x10, 0x92,
	0xb5, 0xda, 0x16, 0xd8, 0xde, 0xb3, 0x6f, 0xb7, 0xfd, 0x89, 0xcf, 0xf8, 0xf3, 0x18, 0x46, 0x39,
	0x1a, 0xbb, 0xc9, 0x95, 0x24, 0x1f, 0x03, 0x3e, 0x74, 0xf9, 0x1b, 0x25, 0x0f, 0x94, 0x30, 0x05,
	0x19, 0xd7, 0x53, 0xaf, 0x4d, 0x71, 0xc7, 0xd5, 0xd1, 0xdf, 0xae, 0xd3, 0xf7, 0x10, 0x92, 0xb7,
	0xa2, 0x96, 0x25, 0x1e, 0xf5, 0x06, 0x7f, 0xd2, 0x7b, 0xe2, 0x5c, 0x78, 0xde, 0xb9, 0xf1, 0x55,
	0x2b, 0xea, 0xae, 0x14, 0xad, 0xb2, 0x7b, 0x72, 0x26, 0xe6, 0xa7, 0x10, 0x79, 0xa3, 0x55, 0x6d,
	0xc9, 0x9b, 0x88, 0xfb, 0x84, 0x31, 0x18, 0xe8, 0x06, 0xeb, 0xfe, 0x46, 0x51, 0xec, 0x56, 0x6e,
	0x4b, 0x6d, 0xb0, 0xb7, 0xc0, 0x27, 0x6e, 0xe5, 0xb5, 0xba, 0xba, 0xee, 0xc5, 0x53, 0xcc, 0x2e,
	0x20, 0x2a, 0xf5, 0x17, 0x92, 0x1c, 0x70, 0x17, 0x9e, 0xdc, 0xd1, 0xf4, 0xee, 0xa7, 0x62, 0xb5,
	0x15, 0x65, 0x06, 0xfe, 0x4c, 0x4a, 0xd8, 0x73, 0x98, 0x50, 0xb0, 0xb1, 0xee, 0x79, 0x99, 0x6c,
	0xec, 0xdb, 0x26, 0x8c, 0x5e, 0x9c, 0x61, 0x4f, 0x21, 0xcd, 0xbb, 0xfd, 0xc6, 0x6f, 0x9e, 0x10,
	0x3f, 0xca, 0xbb, 0xfd, 0x9a, 0xf6, 0x5f, 0x02, 0x18, 0x2c, 0xcb, 0x9e, 0x7d, 0x40, 0x6c, 0xea,
	0x10, 0xa2, 0xf3, 0x84, 0x3e, 0xc0, 0x57, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x69, 0x01, 0xc7,
	0xb3, 0x25, 0x05, 0x00, 0x00,
}
