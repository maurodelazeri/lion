// Code generated by protoc-gen-go. DO NOT EDIT.
// source: products.proto

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

type Product struct {
	ProductId              int64    `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	VenueId                int64    `protobuf:"varint,2,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	Kind                   string   `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	AssetClass             string   `protobuf:"bytes,4,opt,name=asset_class,json=assetClass,proto3" json:"asset_class,omitempty"`
	BaseCurrency           int64    `protobuf:"varint,5,opt,name=base_currency,json=baseCurrency,proto3" json:"base_currency,omitempty"`
	QuoteCurrency          int64    `protobuf:"varint,6,opt,name=quote_currency,json=quoteCurrency,proto3" json:"quote_currency,omitempty"`
	SystemSymbolIdentifier string   `protobuf:"bytes,7,opt,name=system_symbol_identifier,json=systemSymbolIdentifier,proto3" json:"system_symbol_identifier,omitempty"`
	IndividualConnection   bool     `protobuf:"varint,8,opt,name=individual_connection,json=individualConnection,proto3" json:"individual_connection,omitempty"`
	StreamingSave          bool     `protobuf:"varint,9,opt,name=streaming_save,json=streamingSave,proto3" json:"streaming_save,omitempty"`
	VenueSymbolIdentifier  string   `protobuf:"bytes,10,opt,name=venue_symbol_identifier,json=venueSymbolIdentifier,proto3" json:"venue_symbol_identifier,omitempty"`
	MinimumOrdersSize      float64  `protobuf:"fixed64,11,opt,name=minimum_orders_size,json=minimumOrdersSize,proto3" json:"minimum_orders_size,omitempty"`
	StepSize               float64  `protobuf:"fixed64,12,opt,name=step_size,json=stepSize,proto3" json:"step_size,omitempty"`
	PricePrecision         int64    `protobuf:"varint,13,opt,name=price_precision,json=pricePrecision,proto3" json:"price_precision,omitempty"`
	MakerFee               float64  `protobuf:"fixed64,14,opt,name=maker_fee,json=makerFee,proto3" json:"maker_fee,omitempty"`
	TakerFee               float64  `protobuf:"fixed64,15,opt,name=taker_fee,json=takerFee,proto3" json:"taker_fee,omitempty"`
	Enabled                bool     `protobuf:"varint,16,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Settlement             string   `protobuf:"bytes,17,opt,name=settlement,proto3" json:"settlement,omitempty"`
	Expiration             string   `protobuf:"bytes,18,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6e54f42122eb82, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *Product) GetVenueId() int64 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *Product) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Product) GetAssetClass() string {
	if m != nil {
		return m.AssetClass
	}
	return ""
}

func (m *Product) GetBaseCurrency() int64 {
	if m != nil {
		return m.BaseCurrency
	}
	return 0
}

func (m *Product) GetQuoteCurrency() int64 {
	if m != nil {
		return m.QuoteCurrency
	}
	return 0
}

func (m *Product) GetSystemSymbolIdentifier() string {
	if m != nil {
		return m.SystemSymbolIdentifier
	}
	return ""
}

func (m *Product) GetIndividualConnection() bool {
	if m != nil {
		return m.IndividualConnection
	}
	return false
}

func (m *Product) GetStreamingSave() bool {
	if m != nil {
		return m.StreamingSave
	}
	return false
}

func (m *Product) GetVenueSymbolIdentifier() string {
	if m != nil {
		return m.VenueSymbolIdentifier
	}
	return ""
}

func (m *Product) GetMinimumOrdersSize() float64 {
	if m != nil {
		return m.MinimumOrdersSize
	}
	return 0
}

func (m *Product) GetStepSize() float64 {
	if m != nil {
		return m.StepSize
	}
	return 0
}

func (m *Product) GetPricePrecision() int64 {
	if m != nil {
		return m.PricePrecision
	}
	return 0
}

func (m *Product) GetMakerFee() float64 {
	if m != nil {
		return m.MakerFee
	}
	return 0
}

func (m *Product) GetTakerFee() float64 {
	if m != nil {
		return m.TakerFee
	}
	return 0
}

func (m *Product) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Product) GetSettlement() string {
	if m != nil {
		return m.Settlement
	}
	return ""
}

func (m *Product) GetExpiration() string {
	if m != nil {
		return m.Expiration
	}
	return ""
}

type ProductRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ProductId            int64    `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	VenueId              int64    `protobuf:"varint,3,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductRequest) Reset()         { *m = ProductRequest{} }
func (m *ProductRequest) String() string { return proto.CompactTextString(m) }
func (*ProductRequest) ProtoMessage()    {}
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6e54f42122eb82, []int{1}
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

func (m *ProductRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ProductRequest) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *ProductRequest) GetVenueId() int64 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

type ProductResponse struct {
	Retcode              Retcode    `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Product              []*Product `protobuf:"bytes,2,rep,name=product,proto3" json:"product,omitempty"`
	Comment              string     `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string     `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ProductResponse) Reset()         { *m = ProductResponse{} }
func (m *ProductResponse) String() string { return proto.CompactTextString(m) }
func (*ProductResponse) ProtoMessage()    {}
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6e54f42122eb82, []int{2}
}

func (m *ProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductResponse.Unmarshal(m, b)
}
func (m *ProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductResponse.Marshal(b, m, deterministic)
}
func (m *ProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductResponse.Merge(m, src)
}
func (m *ProductResponse) XXX_Size() int {
	return xxx_messageInfo_ProductResponse.Size(m)
}
func (m *ProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductResponse proto.InternalMessageInfo

func (m *ProductResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *ProductResponse) GetProduct() []*Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *ProductResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *ProductResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type ProductPostRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Product              *Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Action               Action   `protobuf:"varint,3,opt,name=action,proto3,enum=api.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductPostRequest) Reset()         { *m = ProductPostRequest{} }
func (m *ProductPostRequest) String() string { return proto.CompactTextString(m) }
func (*ProductPostRequest) ProtoMessage()    {}
func (*ProductPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6e54f42122eb82, []int{3}
}

func (m *ProductPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductPostRequest.Unmarshal(m, b)
}
func (m *ProductPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductPostRequest.Marshal(b, m, deterministic)
}
func (m *ProductPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductPostRequest.Merge(m, src)
}
func (m *ProductPostRequest) XXX_Size() int {
	return xxx_messageInfo_ProductPostRequest.Size(m)
}
func (m *ProductPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProductPostRequest proto.InternalMessageInfo

func (m *ProductPostRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ProductPostRequest) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
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
	Elapsed              string   `protobuf:"bytes,3,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductPostResponse) Reset()         { *m = ProductPostResponse{} }
func (m *ProductPostResponse) String() string { return proto.CompactTextString(m) }
func (*ProductPostResponse) ProtoMessage()    {}
func (*ProductPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6e54f42122eb82, []int{4}
}

func (m *ProductPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductPostResponse.Unmarshal(m, b)
}
func (m *ProductPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductPostResponse.Marshal(b, m, deterministic)
}
func (m *ProductPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductPostResponse.Merge(m, src)
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

func (m *ProductPostResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

func init() {
	proto.RegisterType((*Product)(nil), "api.Product")
	proto.RegisterType((*ProductRequest)(nil), "api.ProductRequest")
	proto.RegisterType((*ProductResponse)(nil), "api.ProductResponse")
	proto.RegisterType((*ProductPostRequest)(nil), "api.ProductPostRequest")
	proto.RegisterType((*ProductPostResponse)(nil), "api.ProductPostResponse")
}

func init() { proto.RegisterFile("products.proto", fileDescriptor_8c6e54f42122eb82) }

var fileDescriptor_8c6e54f42122eb82 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5f, 0x6f, 0xd3, 0x3c,
	0x14, 0xc6, 0x95, 0x66, 0x5b, 0xdb, 0xd3, 0x35, 0x7b, 0xe7, 0x6d, 0x2f, 0x06, 0x04, 0x54, 0x9d,
	0x80, 0x5e, 0xf5, 0x82, 0x49, 0x88, 0x5b, 0x34, 0x09, 0x69, 0x57, 0x4c, 0xd9, 0x07, 0x08, 0x6e,
	0x7c, 0x86, 0xac, 0x25, 0x76, 0x66, 0x3b, 0x85, 0xed, 0x7b, 0x70, 0xc7, 0x87, 0x45, 0x39, 0x4e,
	0xb2, 0x95, 0x69, 0x20, 0xee, 0x72, 0x9e, 0xdf, 0xa3, 0xf3, 0x4f, 0xc7, 0x81, 0xa4, 0xb2, 0x46,
	0xd6, 0xb9, 0x77, 0xcb, 0xca, 0x1a, 0x6f, 0x58, 0x2c, 0x2a, 0xf5, 0x0c, 0x50, 0xd7, 0x65, 0x10,
	0xe6, 0x3f, 0xb7, 0x61, 0x78, 0x1e, 0x3c, 0xec, 0x05, 0x40, 0x6b, 0xcf, 0x94, 0xe4, 0xd1, 0x2c,
	0x5a, 0xc4, 0xe9, 0xb8, 0x55, 0xce, 0x24, 0x7b, 0x0a, 0xa3, 0x35, 0xea, 0x1a, 0x1b, 0x38, 0x20,
	0x38, 0xa4, 0xf8, 0x4c, 0x32, 0x06, 0x5b, 0x57, 0x4a, 0x4b, 0x1e, 0xcf, 0xa2, 0xc5, 0x38, 0xa5,
	0x6f, 0xf6, 0x0a, 0x26, 0xc2, 0x39, 0xf4, 0x59, 0x5e, 0x08, 0xe7, 0xf8, 0x16, 0x21, 0x20, 0xe9,
	0xb4, 0x51, 0xd8, 0x31, 0x4c, 0x57, 0xc2, 0x61, 0x96, 0xd7, 0xd6, 0xa2, 0xce, 0x6f, 0xf8, 0x36,
	0x25, 0xdd, 0x6d, 0xc4, 0xd3, 0x56, 0x63, 0xaf, 0x21, 0xb9, 0xae, 0x8d, 0xbf, 0xe7, 0xda, 0x21,
	0xd7, 0x94, 0xd4, 0xde, 0xf6, 0x01, 0xb8, 0xbb, 0x71, 0x1e, 0xcb, 0xcc, 0xdd, 0x94, 0x2b, 0x53,
	0x64, 0x4a, 0xa2, 0xf6, 0xea, 0x52, 0xa1, 0xe5, 0x43, 0xaa, 0xfc, 0x7f, 0xe0, 0x17, 0x84, 0xcf,
	0x7a, 0xca, 0x4e, 0xe0, 0x48, 0x69, 0xa9, 0xd6, 0x4a, 0xd6, 0xa2, 0xc8, 0x72, 0xa3, 0x35, 0xe6,
	0x5e, 0x19, 0xcd, 0x47, 0xb3, 0x68, 0x31, 0x4a, 0x0f, 0xef, 0xe0, 0x69, 0xcf, 0x9a, 0xae, 0x9c,
	0xb7, 0x28, 0x4a, 0xa5, 0xbf, 0x66, 0x4e, 0xac, 0x91, 0x8f, 0xc9, 0x3d, 0xed, 0xd5, 0x0b, 0xb1,
	0x46, 0xf6, 0x1e, 0x9e, 0x84, 0x8d, 0x3d, 0x6c, 0x0a, 0xa8, 0xa9, 0x23, 0xc2, 0x0f, 0x7a, 0x5a,
	0xc2, 0x41, 0xa9, 0xb4, 0x2a, 0xeb, 0x32, 0x33, 0x56, 0xa2, 0x75, 0x99, 0x53, 0xb7, 0xc8, 0x27,
	0xb3, 0x68, 0x11, 0xa5, 0xfb, 0x2d, 0xfa, 0x4c, 0xe4, 0x42, 0xdd, 0x22, 0x7b, 0x0e, 0x63, 0xe7,
	0xb1, 0x0a, 0xae, 0x5d, 0x72, 0x8d, 0x1a, 0x81, 0xe0, 0x5b, 0xd8, 0xab, 0xac, 0xca, 0x31, 0xab,
	0x2c, 0xe6, 0xca, 0x35, 0xa3, 0x4d, 0x69, 0x85, 0x09, 0xc9, 0xe7, 0x9d, 0xda, 0x64, 0x29, 0xc5,
	0x15, 0xda, 0xec, 0x12, 0x91, 0x27, 0x21, 0x0b, 0x09, 0x9f, 0x90, 0x4a, 0xf8, 0x1e, 0xee, 0x05,
	0xe8, 0x3b, 0xc8, 0x61, 0x88, 0x5a, 0xac, 0x0a, 0x94, 0xfc, 0x3f, 0xda, 0x43, 0x17, 0xb2, 0x97,
	0x00, 0x0e, 0xbd, 0x2f, 0xb0, 0x44, 0xed, 0xf9, 0x7e, 0xb8, 0x81, 0x3b, 0xa5, 0xe1, 0xf8, 0xbd,
	0x52, 0x56, 0xd0, 0xca, 0x59, 0xe0, 0x77, 0xca, 0xfc, 0x0b, 0x24, 0xed, 0x75, 0xa6, 0x78, 0x5d,
	0xa3, 0xf3, 0xec, 0x10, 0xb6, 0xbd, 0xb9, 0x42, 0x4d, 0xf7, 0x39, 0x4e, 0x43, 0xf0, 0xdb, 0xe9,
	0x0e, 0xfe, 0x74, 0xba, 0xf1, 0xc6, 0xe9, 0xce, 0x7f, 0x44, 0xb0, 0xd7, 0x97, 0x70, 0x95, 0xd1,
	0x0e, 0xd9, 0x1b, 0x18, 0x5a, 0xf4, 0xb9, 0x91, 0x48, 0x55, 0x92, 0x77, 0xbb, 0x4b, 0x51, 0xa9,
	0x65, 0x1a, 0xb4, 0xb4, 0x83, 0x8d, 0xaf, 0xad, 0xc1, 0x07, 0xb3, 0x78, 0x31, 0x69, 0x7d, 0x5d,
	0xba, 0x0e, 0x36, 0xfb, 0xc9, 0x4d, 0x49, 0x2b, 0x08, 0x2f, 0xa4, 0x0b, 0x69, 0x73, 0x85, 0xa8,
	0x1c, 0xca, 0xf6, 0x81, 0x74, 0xe1, 0xfc, 0x1b, 0xb0, 0x36, 0xcf, 0xb9, 0x71, 0x7f, 0x99, 0x7e,
	0xa3, 0x8f, 0xe8, 0xf1, 0x3e, 0x8e, 0x61, 0x47, 0x84, 0xe3, 0x8e, 0x69, 0xac, 0x09, 0xd9, 0x3e,
	0x92, 0x94, 0xb6, 0x68, 0x7e, 0x0d, 0x07, 0x1b, 0x85, 0xff, 0x71, 0x27, 0xf7, 0x66, 0x1d, 0x3c,
	0x3a, 0x6b, 0xbc, 0x31, 0xeb, 0x6a, 0x87, 0xfe, 0x45, 0x27, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xe0, 0xb0, 0x76, 0x4c, 0xae, 0x04, 0x00, 0x00,
}
