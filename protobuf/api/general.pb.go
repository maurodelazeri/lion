// Code generated by protoc-gen-go. DO NOT EDIT.
// source: general.proto

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

type ProductRequest struct {
	Venue                Venue    `protobuf:"varint,1,opt,name=venue,proto3,enum=api.Venue" json:"venue,omitempty"`
	Product              Product  `protobuf:"varint,2,opt,name=product,proto3,enum=api.Product" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductRequest) Reset()         { *m = ProductRequest{} }
func (m *ProductRequest) String() string { return proto.CompactTextString(m) }
func (*ProductRequest) ProtoMessage()    {}
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a913b1a5d8940539, []int{0}
}

func (m *ProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductRequest.Unmarshal(m, b)
}
func (m *ProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductRequest.Marshal(b, m, deterministic)
}
func (m *ProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductRequest.Merge(m, src)
}
func (m *ProductRequest) XXX_Size() int {
	return xxx_messageInfo_ProductRequest.Size(m)
}
func (m *ProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProductRequest proto.InternalMessageInfo

func (m *ProductRequest) GetVenue() Venue {
	if m != nil {
		return m.Venue
	}
	return Venue_DARKPOOL
}

func (m *ProductRequest) GetProduct() Product {
	if m != nil {
		return m.Product
	}
	return Product_BTC_USD
}

type ProductResult struct {
	Retcode              Retcode       `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	VenueProducts        *VenueProduct `protobuf:"bytes,2,opt,name=venue_products,json=venueProducts,proto3" json:"venue_products,omitempty"`
	Comment              string        `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ProductResult) Reset()         { *m = ProductResult{} }
func (m *ProductResult) String() string { return proto.CompactTextString(m) }
func (*ProductResult) ProtoMessage()    {}
func (*ProductResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_a913b1a5d8940539, []int{1}
}

func (m *ProductResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductResult.Unmarshal(m, b)
}
func (m *ProductResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductResult.Marshal(b, m, deterministic)
}
func (m *ProductResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductResult.Merge(m, src)
}
func (m *ProductResult) XXX_Size() int {
	return xxx_messageInfo_ProductResult.Size(m)
}
func (m *ProductResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductResult.DiscardUnknown(m)
}

var xxx_messageInfo_ProductResult proto.InternalMessageInfo

func (m *ProductResult) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *ProductResult) GetVenueProducts() *VenueProduct {
	if m != nil {
		return m.VenueProducts
	}
	return nil
}

func (m *ProductResult) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type VenueProduct struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	IndividualConnection bool     `protobuf:"varint,2,opt,name=individual_connection,json=individualConnection,proto3" json:"individual_connection,omitempty"`
	MinimumOrdersSize    float64  `protobuf:"fixed64,3,opt,name=minimum_orders_size,json=minimumOrdersSize,proto3" json:"minimum_orders_size,omitempty"`
	StepSize             float64  `protobuf:"fixed64,4,opt,name=step_size,json=stepSize,proto3" json:"step_size,omitempty"`
	MakerFee             float64  `protobuf:"fixed64,5,opt,name=maker_fee,json=makerFee,proto3" json:"maker_fee,omitempty"`
	TakerFee             float64  `protobuf:"fixed64,6,opt,name=taker_fee,json=takerFee,proto3" json:"taker_fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenueProduct) Reset()         { *m = VenueProduct{} }
func (m *VenueProduct) String() string { return proto.CompactTextString(m) }
func (*VenueProduct) ProtoMessage()    {}
func (*VenueProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_a913b1a5d8940539, []int{2}
}

func (m *VenueProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenueProduct.Unmarshal(m, b)
}
func (m *VenueProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenueProduct.Marshal(b, m, deterministic)
}
func (m *VenueProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueProduct.Merge(m, src)
}
func (m *VenueProduct) XXX_Size() int {
	return xxx_messageInfo_VenueProduct.Size(m)
}
func (m *VenueProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_VenueProduct.DiscardUnknown(m)
}

var xxx_messageInfo_VenueProduct proto.InternalMessageInfo

func (m *VenueProduct) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *VenueProduct) GetIndividualConnection() bool {
	if m != nil {
		return m.IndividualConnection
	}
	return false
}

func (m *VenueProduct) GetMinimumOrdersSize() float64 {
	if m != nil {
		return m.MinimumOrdersSize
	}
	return 0
}

func (m *VenueProduct) GetStepSize() float64 {
	if m != nil {
		return m.StepSize
	}
	return 0
}

func (m *VenueProduct) GetMakerFee() float64 {
	if m != nil {
		return m.MakerFee
	}
	return 0
}

func (m *VenueProduct) GetTakerFee() float64 {
	if m != nil {
		return m.TakerFee
	}
	return 0
}

func init() {
	proto.RegisterType((*ProductRequest)(nil), "api.ProductRequest")
	proto.RegisterType((*ProductResult)(nil), "api.ProductResult")
	proto.RegisterType((*VenueProduct)(nil), "api.VenueProduct")
}

func init() { proto.RegisterFile("general.proto", fileDescriptor_a913b1a5d8940539) }

var fileDescriptor_a913b1a5d8940539 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x15, 0x4a, 0xff, 0x2e, 0x6d, 0xa5, 0x1a, 0x90, 0x22, 0x58, 0xaa, 0x0e, 0xa8, 0x53,
	0x86, 0x76, 0x61, 0x47, 0x62, 0x05, 0x19, 0x89, 0x81, 0x25, 0x4a, 0xe3, 0x0b, 0xb2, 0x88, 0xed,
	0xe0, 0x9f, 0x0e, 0x7d, 0x05, 0x5e, 0x93, 0x07, 0x41, 0xb9, 0xae, 0xdb, 0x8e, 0xf7, 0x7c, 0xe7,
	0x1c, 0x1f, 0xc9, 0x30, 0xfd, 0x42, 0x8d, 0xb6, 0x6a, 0x8a, 0xd6, 0x1a, 0x6f, 0x58, 0xaf, 0x6a,
	0xe5, 0x1d, 0xa0, 0x0e, 0x2a, 0x0a, 0xcb, 0x0f, 0x98, 0xbd, 0x5a, 0x23, 0x42, 0xed, 0x39, 0xfe,
	0x04, 0x74, 0x9e, 0x2d, 0xa0, 0xbf, 0x43, 0x1d, 0x30, 0xcf, 0x16, 0xd9, 0x6a, 0xb6, 0x86, 0xa2,
	0x6a, 0x65, 0xf1, 0xde, 0x29, 0x3c, 0x02, 0xf6, 0x00, 0xc3, 0x36, 0x66, 0xf2, 0x0b, 0xf2, 0x4c,
	0xc8, 0x93, 0x7a, 0x12, 0x5c, 0xfe, 0x66, 0x30, 0x3d, 0x96, 0xbb, 0xd0, 0xf8, 0x2e, 0x69, 0xd1,
	0xd7, 0x46, 0xa4, 0xf6, 0x98, 0xe4, 0x51, 0xe3, 0x09, 0xb2, 0x47, 0x98, 0xd1, 0x53, 0xe5, 0xa1,
	0xca, 0xd1, 0x43, 0x57, 0xeb, 0xf9, 0x69, 0x4c, 0x2a, 0x9e, 0xee, 0xce, 0x2e, 0xc7, 0x72, 0x18,
	0xd6, 0x46, 0x29, 0xd4, 0x3e, 0xef, 0x2d, 0xb2, 0xd5, 0x98, 0xa7, 0x73, 0xf9, 0x97, 0xc1, 0xe4,
	0x3c, 0xd9, 0x59, 0x51, 0x57, 0xdb, 0x06, 0x05, 0x8d, 0x19, 0xf1, 0x74, 0xb2, 0x0d, 0xdc, 0x4a,
	0x2d, 0xe4, 0x4e, 0x8a, 0x50, 0x35, 0x65, 0x6d, 0xb4, 0xc6, 0xda, 0x4b, 0xa3, 0x69, 0xc5, 0x88,
	0xdf, 0x9c, 0xe0, 0xd3, 0x91, 0xb1, 0x02, 0xae, 0x95, 0xd4, 0x52, 0x05, 0x55, 0x1a, 0x2b, 0xd0,
	0xba, 0xd2, 0xc9, 0x3d, 0xd2, 0x8a, 0x8c, 0xcf, 0x0f, 0xe8, 0x85, 0xc8, 0x9b, 0xdc, 0x23, 0xbb,
	0x87, 0xb1, 0xf3, 0xd8, 0x46, 0xd7, 0x25, 0xb9, 0x46, 0x9d, 0x90, 0xa0, 0xaa, 0xbe, 0xd1, 0x96,
	0x9f, 0x88, 0x79, 0x3f, 0x42, 0x12, 0x9e, 0x91, 0xa0, 0x3f, 0xc2, 0x41, 0x84, 0xfe, 0x00, 0xb7,
	0x03, 0xfa, 0xd7, 0xcd, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x90, 0xef, 0xab, 0xf9, 0x01,
	0x00, 0x00,
}
