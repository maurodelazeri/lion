// Code generated by protoc-gen-go. DO NOT EDIT.
// source: venues.proto

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

type Venue struct {
	VenueId              int64    `protobuf:"varint,1,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	VenueDescription     string   `protobuf:"bytes,3,opt,name=venue_description,json=venueDescription,proto3" json:"venue_description,omitempty"`
	ApiKey               string   `protobuf:"bytes,4,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	ApiSecret            string   `protobuf:"bytes,5,opt,name=api_secret,json=apiSecret,proto3" json:"api_secret,omitempty"`
	Passphrase           string   `protobuf:"bytes,6,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
	Enabled              bool     `protobuf:"varint,7,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Spot                 bool     `protobuf:"varint,8,opt,name=spot,proto3" json:"spot,omitempty"`
	Futures              bool     `protobuf:"varint,9,opt,name=futures,proto3" json:"futures,omitempty"`
	Options              bool     `protobuf:"varint,10,opt,name=options,proto3" json:"options,omitempty"`
	Swaps                bool     `protobuf:"varint,11,opt,name=swaps,proto3" json:"swaps,omitempty"`
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
func (m *Venue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Venue.Merge(m, src)
}
func (m *Venue) XXX_Size() int {
	return xxx_messageInfo_Venue.Size(m)
}
func (m *Venue) XXX_DiscardUnknown() {
	xxx_messageInfo_Venue.DiscardUnknown(m)
}

var xxx_messageInfo_Venue proto.InternalMessageInfo

func (m *Venue) GetVenueId() int64 {
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

func (m *Venue) GetSwaps() bool {
	if m != nil {
		return m.Swaps
	}
	return false
}

type VenueRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	VenueId              int64    `protobuf:"varint,2,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
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
func (m *VenueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueRequest.Merge(m, src)
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

func (m *VenueRequest) GetVenueId() int64 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

type VenueResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Venue                []*Venue `protobuf:"bytes,2,rep,name=venue,proto3" json:"venue,omitempty"`
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
func (m *VenueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenueResponse.Merge(m, src)
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
	return Retcode_DONE
}

func (m *VenueResponse) GetVenue() []*Venue {
	if m != nil {
		return m.Venue
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
func (m *VenuePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuePostRequest.Merge(m, src)
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
func (m *VenuePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VenuePostResponse.Merge(m, src)
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
	return Retcode_DONE
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

func init() {
	proto.RegisterType((*Venue)(nil), "api.Venue")
	proto.RegisterType((*VenueRequest)(nil), "api.VenueRequest")
	proto.RegisterType((*VenueResponse)(nil), "api.VenueResponse")
	proto.RegisterType((*VenuePostRequest)(nil), "api.VenuePostRequest")
	proto.RegisterType((*VenuePostResponse)(nil), "api.VenuePostResponse")
}

func init() { proto.RegisterFile("venues.proto", fileDescriptor_8c34a3f781dd794d) }

var fileDescriptor_8c34a3f781dd794d = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xdd, 0x8a, 0xd4, 0x30,
	0x14, 0xa6, 0xed, 0xce, 0x74, 0x7a, 0x66, 0x1c, 0x76, 0x83, 0x60, 0x5c, 0x50, 0x4a, 0x05, 0x29,
	0x08, 0x73, 0x31, 0x3e, 0x80, 0x08, 0xde, 0x88, 0x37, 0x12, 0xc1, 0xdb, 0x25, 0xdb, 0x1e, 0x31,
	0xec, 0x36, 0x89, 0x3d, 0xa9, 0x32, 0x0f, 0xe1, 0xd3, 0xf8, 0x82, 0xd2, 0x93, 0x56, 0x3b, 0xa0,
	0x82, 0x77, 0xf9, 0x7e, 0x72, 0x92, 0xef, 0x4b, 0x60, 0xf7, 0x15, 0xed, 0x80, 0x74, 0xf0, 0xbd,
	0x0b, 0x4e, 0x64, 0xda, 0x9b, 0x6b, 0x40, 0x3b, 0x74, 0x91, 0xb8, 0xde, 0xfb, 0xde, 0xb5, 0x43,
	0x13, 0x26, 0x43, 0xf5, 0x23, 0x85, 0xd5, 0xc7, 0x71, 0x87, 0x78, 0x0c, 0x1b, 0xde, 0x7a, 0x63,
	0x5a, 0x99, 0x94, 0x49, 0x9d, 0xa9, 0x9c, 0xf1, 0xdb, 0x56, 0x08, 0xb8, 0xb0, 0xba, 0x43, 0x99,
	0x96, 0x49, 0x5d, 0x28, 0x5e, 0x8b, 0x17, 0x70, 0x15, 0xed, 0x2d, 0x52, 0xd3, 0x1b, 0x1f, 0x8c,
	0xb3, 0x32, 0x63, 0xc3, 0x25, 0x0b, 0x6f, 0x7e, 0xf3, 0xe2, 0x11, 0xe4, 0xda, 0x9b, 0x9b, 0x3b,
	0x3c, 0xc9, 0x0b, 0xb6, 0xac, 0xb5, 0x37, 0xef, 0xf0, 0x24, 0x9e, 0x00, 0x8c, 0x02, 0x61, 0xd3,
	0x63, 0x90, 0x2b, 0xd6, 0x0a, 0xed, 0xcd, 0x07, 0x26, 0xc4, 0x53, 0x00, 0xaf, 0x89, 0xfc, 0xe7,
	0x5e, 0x13, 0xca, 0x35, 0xcb, 0x0b, 0x46, 0x48, 0xc8, 0xd1, 0xea, 0xdb, 0x7b, 0x6c, 0x65, 0x5e,
	0x26, 0xf5, 0x46, 0xcd, 0x70, 0xbc, 0x32, 0x79, 0x17, 0xe4, 0x86, 0x69, 0x5e, 0x8f, 0xee, 0x4f,
	0x43, 0x18, 0x7a, 0x24, 0x59, 0x44, 0xf7, 0x04, 0x47, 0xc5, 0xf1, 0x4d, 0x49, 0x42, 0x54, 0x26,
	0x28, 0x1e, 0xc2, 0x8a, 0xbe, 0x69, 0x4f, 0x72, 0xcb, 0x7c, 0x04, 0xd5, 0x2b, 0xd8, 0x71, 0x69,
	0x0a, 0xbf, 0x0c, 0x48, 0x61, 0x74, 0x05, 0x77, 0x87, 0x96, 0x8b, 0x2b, 0x54, 0x04, 0x67, 0x8d,
	0xa6, 0x67, 0x8d, 0x56, 0xdf, 0x13, 0x78, 0x30, 0x4d, 0x20, 0xef, 0x2c, 0xa1, 0x78, 0x0e, 0x79,
	0x8f, 0xa1, 0x71, 0x2d, 0xf2, 0x90, 0xfd, 0x71, 0x77, 0xd0, 0xde, 0x1c, 0x54, 0xe4, 0xd4, 0x2c,
	0x8a, 0x12, 0x56, 0x3c, 0x44, 0xa6, 0x65, 0x56, 0x6f, 0x8f, 0xc0, 0xae, 0x38, 0x2a, 0x0a, 0x63,
	0x98, 0xc6, 0x75, 0x1d, 0xda, 0x30, 0xbd, 0xc7, 0x0c, 0xb9, 0xae, 0x7b, 0xed, 0x09, 0xdb, 0xe9,
	0x19, 0x66, 0x58, 0x39, 0xb8, 0xe4, 0x19, 0xef, 0x1d, 0x85, 0x7f, 0x87, 0x5a, 0x9c, 0x9f, 0xfc,
	0xf9, 0xfc, 0x67, 0xb0, 0xd6, 0xcd, 0xaf, 0xef, 0xb0, 0x3f, 0x6e, 0xd9, 0xf2, 0x9a, 0x29, 0x35,
	0x49, 0x95, 0x83, 0xab, 0xc5, 0x81, 0xff, 0xd9, 0xc1, 0x22, 0x61, 0xfa, 0xd7, 0x84, 0xd9, 0x59,
	0xc2, 0xdb, 0x35, 0xff, 0xf7, 0x97, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x64, 0x55, 0x94, 0x00,
	0x20, 0x03, 0x00, 0x00,
}
