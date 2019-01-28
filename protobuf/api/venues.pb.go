// Code generated by protoc-gen-go. DO NOT EDIT.
// source: venues.proto

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

type Venue struct {
	VenueId              int32    `protobuf:"varint,1,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	VenueDescription     string   `protobuf:"bytes,3,opt,name=venue_description,json=venueDescription,proto3" json:"venue_description,omitempty"`
	ApiKey               string   `protobuf:"bytes,4,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	ApiSecret            string   `protobuf:"bytes,5,opt,name=api_secret,json=apiSecret,proto3" json:"api_secret,omitempty"`
	Passphrase           string   `protobuf:"bytes,6,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	Enabled              bool     `protobuf:"varint,7,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Spot                 bool     `protobuf:"varint,8,opt,name=spot,proto3" json:"spot,omitempty"`
	Futures              bool     `protobuf:"varint,9,opt,name=futures,proto3" json:"futures,omitempty"`
	Options              bool     `protobuf:"varint,10,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Venue) Reset()         { *m = Venue{} }
func (m *Venue) String() string { return proto.CompactTextString(m) }
func (*Venue) ProtoMessage()    {}
func (*Venue) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c34a3f781dd794d, []int{0}
}
func (m *Venue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Venue.Unmarshal(m, b)
}
func (m *Venue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Venue.Marshal(b, m, deterministic)
}
func (dst *Venue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Venue.Merge(dst, src)
}
func (m *Venue) XXX_Size() int {
	return xxx_messageInfo_Venue.Size(m)
}
func (m *Venue) XXX_DiscardUnknown() {
	xxx_messageInfo_Venue.DiscardUnknown(m)
}

var xxx_messageInfo_Venue proto.InternalMessageInfo

func (m *Venue) GetVenueId() int32 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *Venue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Venue) GetVenueDescription() string {
	if m != nil {
		return m.VenueDescription
	}
	return ""
}

func (m *Venue) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *Venue) GetApiSecret() string {
	if m != nil {
		return m.ApiSecret
	}
	return ""
}

func (m *Venue) GetPassphrase() string {
	if m != nil {
		return m.Passphrase
	}
	return ""
}

func (m *Venue) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Venue) GetSpot() bool {
	if m != nil {
		return m.Spot
	}
	return false
}

func (m *Venue) GetFutures() bool {
	if m != nil {
		return m.Futures
	}
	return false
}

func (m *Venue) GetOptions() bool {
	if m != nil {
		return m.Options
	}
	return false
}

// Venue
type VenueRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	VenueId              int32    `protobuf:"varint,2,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenueRequest) Reset()         { *m = VenueRequest{} }
func (m *VenueRequest) String() string { return proto.CompactTextString(m) }
func (*VenueRequest) ProtoMessage()    {}
func (*VenueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c34a3f781dd794d, []int{1}
}
func (m *VenueRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenueRequest.Unmarshal(m, b)
}
func (m *VenueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenueRequest.Marshal(b, m, deterministic)
}
func (dst *VenueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueRequest.Merge(dst, src)
}
func (m *VenueRequest) XXX_Size() int {
	return xxx_messageInfo_VenueRequest.Size(m)
}
func (m *VenueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VenueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VenueRequest proto.InternalMessageInfo

func (m *VenueRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *VenueRequest) GetVenueId() int32 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *VenueRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type VenueResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Venues               []*Venue `protobuf:"bytes,2,rep,name=venues,proto3" json:"venues,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenueResponse) Reset()         { *m = VenueResponse{} }
func (m *VenueResponse) String() string { return proto.CompactTextString(m) }
func (*VenueResponse) ProtoMessage()    {}
func (*VenueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c34a3f781dd794d, []int{2}
}
func (m *VenueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenueResponse.Unmarshal(m, b)
}
func (m *VenueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenueResponse.Marshal(b, m, deterministic)
}
func (dst *VenueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueResponse.Merge(dst, src)
}
func (m *VenueResponse) XXX_Size() int {
	return xxx_messageInfo_VenueResponse.Size(m)
}
func (m *VenueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VenueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VenueResponse proto.InternalMessageInfo

func (m *VenueResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *VenueResponse) GetVenues() []*Venue {
	if m != nil {
		return m.Venues
	}
	return nil
}

func (m *VenueResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *VenueResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type VenuePostRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Venue                *Venue   `protobuf:"bytes,2,opt,name=venue,proto3" json:"venue,omitempty"`
	Action               Action   `protobuf:"varint,3,opt,name=action,proto3,enum=api.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenuePostRequest) Reset()         { *m = VenuePostRequest{} }
func (m *VenuePostRequest) String() string { return proto.CompactTextString(m) }
func (*VenuePostRequest) ProtoMessage()    {}
func (*VenuePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c34a3f781dd794d, []int{3}
}
func (m *VenuePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenuePostRequest.Unmarshal(m, b)
}
func (m *VenuePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenuePostRequest.Marshal(b, m, deterministic)
}
func (dst *VenuePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuePostRequest.Merge(dst, src)
}
func (m *VenuePostRequest) XXX_Size() int {
	return xxx_messageInfo_VenuePostRequest.Size(m)
}
func (m *VenuePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VenuePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VenuePostRequest proto.InternalMessageInfo

func (m *VenuePostRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *VenuePostRequest) GetVenue() *Venue {
	if m != nil {
		return m.Venue
	}
	return nil
}

func (m *VenuePostRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_INSERT
}

type VenuePostResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,3,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenuePostResponse) Reset()         { *m = VenuePostResponse{} }
func (m *VenuePostResponse) String() string { return proto.CompactTextString(m) }
func (*VenuePostResponse) ProtoMessage()    {}
func (*VenuePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c34a3f781dd794d, []int{4}
}
func (m *VenuePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenuePostResponse.Unmarshal(m, b)
}
func (m *VenuePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenuePostResponse.Marshal(b, m, deterministic)
}
func (dst *VenuePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuePostResponse.Merge(dst, src)
}
func (m *VenuePostResponse) XXX_Size() int {
	return xxx_messageInfo_VenuePostResponse.Size(m)
}
func (m *VenuePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VenuePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VenuePostResponse proto.InternalMessageInfo

func (m *VenuePostResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *VenuePostResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *VenuePostResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

// Venue Products
type VenueProduct struct {
	VenueId              int32      `protobuf:"varint,1,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	VenueDescription     string     `protobuf:"bytes,2,opt,name=venue_description,json=venueDescription,proto3" json:"venue_description,omitempty"`
	Total                int32      `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Enabled              int32      `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Disabled             int32      `protobuf:"varint,5,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Products             []*Product `protobuf:"bytes,6,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *VenueProduct) Reset()         { *m = VenueProduct{} }
func (m *VenueProduct) String() string { return proto.CompactTextString(m) }
func (*VenueProduct) ProtoMessage()    {}
func (*VenueProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c34a3f781dd794d, []int{5}
}
func (m *VenueProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VenueProduct.Unmarshal(m, b)
}
func (m *VenueProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VenueProduct.Marshal(b, m, deterministic)
}
func (dst *VenueProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueProduct.Merge(dst, src)
}
func (m *VenueProduct) XXX_Size() int {
	return xxx_messageInfo_VenueProduct.Size(m)
}
func (m *VenueProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_VenueProduct.DiscardUnknown(m)
}

var xxx_messageInfo_VenueProduct proto.InternalMessageInfo

func (m *VenueProduct) GetVenueId() int32 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *VenueProduct) GetVenueDescription() string {
	if m != nil {
		return m.VenueDescription
	}
	return ""
}

func (m *VenueProduct) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *VenueProduct) GetEnabled() int32 {
	if m != nil {
		return m.Enabled
	}
	return 0
}

func (m *VenueProduct) GetDisabled() int32 {
	if m != nil {
		return m.Disabled
	}
	return 0
}

func (m *VenueProduct) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

type VenueProductRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	VenueId              int32    `protobuf:"varint,2,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VenueProductRequest) Reset()         { *m = VenueProductRequest{} }
func (m *VenueProductRequest) String() string { return proto.CompactTextString(m) }
func (*VenueProductRequest) ProtoMessage()    {}
func (*VenueProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c34a3f781dd794d, []int{6}
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

func (m *VenueProductRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *VenueProductRequest) GetVenueId() int32 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *VenueProductRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type VenueProductResponse struct {
	Retcode              Retcode         `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	VenueProducts        []*VenueProduct `protobuf:"bytes,2,rep,name=venue_products,json=venueProducts,proto3" json:"venue_products,omitempty"`
	Comment              string          `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string          `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *VenueProductResponse) Reset()         { *m = VenueProductResponse{} }
func (m *VenueProductResponse) String() string { return proto.CompactTextString(m) }
func (*VenueProductResponse) ProtoMessage()    {}
func (*VenueProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c34a3f781dd794d, []int{7}
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

func (m *VenueProductResponse) GetVenueProducts() []*VenueProduct {
	if m != nil {
		return m.VenueProducts
	}
	return nil
}

func (m *VenueProductResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *VenueProductResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

func init() {
	proto.RegisterType((*Venue)(nil), "api.Venue")
	proto.RegisterType((*VenueRequest)(nil), "api.VenueRequest")
	proto.RegisterType((*VenueResponse)(nil), "api.VenueResponse")
	proto.RegisterType((*VenuePostRequest)(nil), "api.VenuePostRequest")
	proto.RegisterType((*VenuePostResponse)(nil), "api.VenuePostResponse")
	proto.RegisterType((*VenueProduct)(nil), "api.VenueProduct")
	proto.RegisterType((*VenueProductRequest)(nil), "api.VenueProductRequest")
	proto.RegisterType((*VenueProductResponse)(nil), "api.VenueProductResponse")
}

func init() { proto.RegisterFile("venues.proto", fileDescriptor_8c34a3f781dd794d) }

var fileDescriptor_8c34a3f781dd794d = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0x46, 0xb2, 0xf5, 0xe3, 0x89, 0x63, 0xe2, 0x6d, 0xa0, 0x5b, 0x43, 0x8b, 0x50, 0xa1, 0x08,
	0x0a, 0x3e, 0xb8, 0x97, 0x5e, 0x0b, 0xbd, 0x94, 0x5e, 0xc2, 0x06, 0x7a, 0xe8, 0x25, 0x6c, 0xa4,
	0x29, 0x15, 0xb1, 0xb5, 0x5b, 0xed, 0x2a, 0x90, 0xa7, 0xe8, 0x13, 0xf4, 0x0d, 0xfa, 0x28, 0x7d,
	0xa8, 0xa2, 0xd9, 0x95, 0x2d, 0x43, 0x12, 0x48, 0xe9, 0x4d, 0xdf, 0x7c, 0xa3, 0xf9, 0xf9, 0xe6,
	0x93, 0x60, 0x7e, 0x8b, 0x4d, 0x87, 0x66, 0xad, 0x5b, 0x65, 0x15, 0x9b, 0x48, 0x5d, 0xaf, 0x00,
	0x9b, 0x6e, 0xe7, 0x02, 0xab, 0x85, 0x6e, 0x55, 0xd5, 0x95, 0xd6, 0x27, 0xe4, 0xbf, 0x42, 0x88,
	0xbe, 0xf4, 0x6f, 0xb0, 0x17, 0x90, 0xd2, 0xab, 0x57, 0x75, 0xc5, 0x83, 0x2c, 0x28, 0x22, 0x91,
	0x10, 0xfe, 0x54, 0x31, 0x06, 0xd3, 0x46, 0xee, 0x90, 0x87, 0x59, 0x50, 0xcc, 0x04, 0x3d, 0xb3,
	0xb7, 0xb0, 0x74, 0xe9, 0x15, 0x9a, 0xb2, 0xad, 0xb5, 0xad, 0x55, 0xc3, 0x27, 0x94, 0x70, 0x46,
	0xc4, 0xc7, 0x43, 0x9c, 0x3d, 0x87, 0x44, 0xea, 0xfa, 0xea, 0x06, 0xef, 0xf8, 0x94, 0x52, 0x62,
	0xa9, 0xeb, 0xcf, 0x78, 0xc7, 0x5e, 0x02, 0xf4, 0x84, 0xc1, 0xb2, 0x45, 0xcb, 0x23, 0xe2, 0x66,
	0x52, 0xd7, 0x97, 0x14, 0x60, 0xaf, 0x00, 0xb4, 0x34, 0x46, 0x7f, 0x6f, 0xa5, 0x41, 0x1e, 0x13,
	0x3d, 0x8a, 0x30, 0x0e, 0x09, 0x36, 0xf2, 0x7a, 0x8b, 0x15, 0x4f, 0xb2, 0xa0, 0x48, 0xc5, 0x00,
	0xfb, 0x91, 0x8d, 0x56, 0x96, 0xa7, 0x14, 0xa6, 0xe7, 0x3e, 0xfb, 0x5b, 0x67, 0xbb, 0x16, 0x0d,
	0x9f, 0xb9, 0x6c, 0x0f, 0x7b, 0x46, 0xd1, 0xa4, 0x86, 0x83, 0x63, 0x3c, 0xcc, 0x2f, 0x61, 0x4e,
	0xf2, 0x08, 0xfc, 0xd1, 0xa1, 0xb1, 0xec, 0x1c, 0x22, 0xab, 0x6e, 0xb0, 0x21, 0x89, 0x66, 0xc2,
	0x81, 0x23, 0xed, 0xc2, 0xfb, 0xb5, 0x9b, 0x1c, 0xb4, 0xcb, 0x7f, 0x06, 0x70, 0xea, 0xab, 0x1a,
	0xad, 0x1a, 0x83, 0xec, 0x0d, 0x24, 0x2d, 0xda, 0x52, 0x55, 0x48, 0x85, 0x17, 0x9b, 0xf9, 0x5a,
	0xea, 0x7a, 0x2d, 0x5c, 0x4c, 0x0c, 0x24, 0xcb, 0x21, 0x76, 0xf7, 0xe5, 0x61, 0x36, 0x29, 0x4e,
	0x36, 0x40, 0x69, 0xae, 0x96, 0x67, 0xfa, 0x65, 0x4a, 0xb5, 0xdb, 0x61, 0x63, 0x7d, 0xd3, 0x01,
	0x92, 0x5c, 0x5b, 0xa9, 0x0d, 0x56, 0xfe, 0x0c, 0x03, 0xcc, 0x15, 0x9c, 0x51, 0x91, 0x0b, 0x65,
	0xec, 0xe3, 0xab, 0x66, 0x10, 0x51, 0x1f, 0xda, 0xf3, 0x78, 0x00, 0x47, 0xb0, 0xd7, 0x10, 0xcb,
	0x72, 0x6f, 0x87, 0xc5, 0xe6, 0x84, 0x52, 0x3e, 0x50, 0x48, 0x78, 0x2a, 0x57, 0xb0, 0x1c, 0x35,
	0x7c, 0xa2, 0x0a, 0xa3, 0x0d, 0xc3, 0x07, 0x37, 0x9c, 0x1c, 0x6f, 0xf8, 0x27, 0xf0, 0x97, 0xbc,
	0x70, 0x1f, 0xc0, 0x63, 0x7e, 0xbf, 0xd7, 0xdb, 0xe1, 0x03, 0xde, 0x26, 0x99, 0xac, 0xdc, 0x52,
	0xc3, 0x48, 0x38, 0x30, 0x76, 0xe6, 0xd4, 0x15, 0x1f, 0x9c, 0xb9, 0x82, 0xb4, 0xaa, 0x8d, 0xa3,
	0x22, 0xa2, 0xf6, 0x98, 0x15, 0x90, 0x0e, 0xdf, 0x27, 0x8f, 0xe9, 0xc0, 0x4e, 0x01, 0x3f, 0xb3,
	0xd8, 0xb3, 0xf9, 0x57, 0x78, 0x36, 0xde, 0xe6, 0xbf, 0xda, 0xf3, 0x77, 0x00, 0xe7, 0xc7, 0xc5,
	0x9f, 0x78, 0x9f, 0xf7, 0xb0, 0x70, 0xfd, 0xf6, 0xcb, 0x38, 0xb7, 0x2e, 0x0f, 0x66, 0x19, 0x4a,
	0x9f, 0xde, 0x8e, 0xd0, 0x3f, 0x79, 0xf7, 0x3a, 0xa6, 0x3f, 0xd9, 0xbb, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x2a, 0xe8, 0x27, 0x86, 0xfa, 0x04, 0x00, 0x00,
}
