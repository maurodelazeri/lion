// Code generated by protoc-gen-go. DO NOT EDIT.
// source: currencies.proto

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

type Currency struct {
	CurrencyId           int32    `protobuf:"varint,1,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CurrencyDescription  string   `protobuf:"bytes,3,opt,name=currency_description,json=currencyDescription,proto3" json:"currency_description,omitempty"`
	MinConfirmation      int32    `protobuf:"varint,4,opt,name=min_confirmation,json=minConfirmation,proto3" json:"min_confirmation,omitempty"`
	Enabled              bool     `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	TxFee                float64  `protobuf:"fixed64,6,opt,name=tx_fee,json=txFee,proto3" json:"tx_fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Currency) Reset()         { *m = Currency{} }
func (m *Currency) String() string { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()    {}
func (*Currency) Descriptor() ([]byte, []int) {
	return fileDescriptor_1988b70e90d5a630, []int{0}
}

func (m *Currency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Currency.Unmarshal(m, b)
}
func (m *Currency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Currency.Marshal(b, m, deterministic)
}
func (m *Currency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Currency.Merge(m, src)
}
func (m *Currency) XXX_Size() int {
	return xxx_messageInfo_Currency.Size(m)
}
func (m *Currency) XXX_DiscardUnknown() {
	xxx_messageInfo_Currency.DiscardUnknown(m)
}

var xxx_messageInfo_Currency proto.InternalMessageInfo

func (m *Currency) GetCurrencyId() int32 {
	if m != nil {
		return m.CurrencyId
	}
	return 0
}

func (m *Currency) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Currency) GetCurrencyDescription() string {
	if m != nil {
		return m.CurrencyDescription
	}
	return ""
}

func (m *Currency) GetMinConfirmation() int32 {
	if m != nil {
		return m.MinConfirmation
	}
	return 0
}

func (m *Currency) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Currency) GetTxFee() float64 {
	if m != nil {
		return m.TxFee
	}
	return 0
}

type CurrencyRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	CurrencyId           int32    `protobuf:"varint,2,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrencyRequest) Reset()         { *m = CurrencyRequest{} }
func (m *CurrencyRequest) String() string { return proto.CompactTextString(m) }
func (*CurrencyRequest) ProtoMessage()    {}
func (*CurrencyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1988b70e90d5a630, []int{1}
}

func (m *CurrencyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrencyRequest.Unmarshal(m, b)
}
func (m *CurrencyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrencyRequest.Marshal(b, m, deterministic)
}
func (m *CurrencyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyRequest.Merge(m, src)
}
func (m *CurrencyRequest) XXX_Size() int {
	return xxx_messageInfo_CurrencyRequest.Size(m)
}
func (m *CurrencyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyRequest proto.InternalMessageInfo

func (m *CurrencyRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CurrencyRequest) GetCurrencyId() int32 {
	if m != nil {
		return m.CurrencyId
	}
	return 0
}

func (m *CurrencyRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CurrencyResponse struct {
	Retcode              Retcode     `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Currencies           []*Currency `protobuf:"bytes,2,rep,name=currencies,proto3" json:"currencies,omitempty"`
	Comment              string      `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string      `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CurrencyResponse) Reset()         { *m = CurrencyResponse{} }
func (m *CurrencyResponse) String() string { return proto.CompactTextString(m) }
func (*CurrencyResponse) ProtoMessage()    {}
func (*CurrencyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1988b70e90d5a630, []int{2}
}

func (m *CurrencyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrencyResponse.Unmarshal(m, b)
}
func (m *CurrencyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrencyResponse.Marshal(b, m, deterministic)
}
func (m *CurrencyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyResponse.Merge(m, src)
}
func (m *CurrencyResponse) XXX_Size() int {
	return xxx_messageInfo_CurrencyResponse.Size(m)
}
func (m *CurrencyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyResponse proto.InternalMessageInfo

func (m *CurrencyResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *CurrencyResponse) GetCurrencies() []*Currency {
	if m != nil {
		return m.Currencies
	}
	return nil
}

func (m *CurrencyResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *CurrencyResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type CurrencyPostRequest struct {
	Token                string    `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Currency             *Currency `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Action               Action    `protobuf:"varint,3,opt,name=action,proto3,enum=api.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CurrencyPostRequest) Reset()         { *m = CurrencyPostRequest{} }
func (m *CurrencyPostRequest) String() string { return proto.CompactTextString(m) }
func (*CurrencyPostRequest) ProtoMessage()    {}
func (*CurrencyPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1988b70e90d5a630, []int{3}
}

func (m *CurrencyPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrencyPostRequest.Unmarshal(m, b)
}
func (m *CurrencyPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrencyPostRequest.Marshal(b, m, deterministic)
}
func (m *CurrencyPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyPostRequest.Merge(m, src)
}
func (m *CurrencyPostRequest) XXX_Size() int {
	return xxx_messageInfo_CurrencyPostRequest.Size(m)
}
func (m *CurrencyPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyPostRequest proto.InternalMessageInfo

func (m *CurrencyPostRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *CurrencyPostRequest) GetCurrency() *Currency {
	if m != nil {
		return m.Currency
	}
	return nil
}

func (m *CurrencyPostRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_INSERT
}

type CurrencyPostResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,3,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrencyPostResponse) Reset()         { *m = CurrencyPostResponse{} }
func (m *CurrencyPostResponse) String() string { return proto.CompactTextString(m) }
func (*CurrencyPostResponse) ProtoMessage()    {}
func (*CurrencyPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1988b70e90d5a630, []int{4}
}

func (m *CurrencyPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrencyPostResponse.Unmarshal(m, b)
}
func (m *CurrencyPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrencyPostResponse.Marshal(b, m, deterministic)
}
func (m *CurrencyPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyPostResponse.Merge(m, src)
}
func (m *CurrencyPostResponse) XXX_Size() int {
	return xxx_messageInfo_CurrencyPostResponse.Size(m)
}
func (m *CurrencyPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyPostResponse proto.InternalMessageInfo

func (m *CurrencyPostResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *CurrencyPostResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *CurrencyPostResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

func init() {
	proto.RegisterType((*Currency)(nil), "api.Currency")
	proto.RegisterType((*CurrencyRequest)(nil), "api.CurrencyRequest")
	proto.RegisterType((*CurrencyResponse)(nil), "api.CurrencyResponse")
	proto.RegisterType((*CurrencyPostRequest)(nil), "api.CurrencyPostRequest")
	proto.RegisterType((*CurrencyPostResponse)(nil), "api.CurrencyPostResponse")
}

func init() { proto.RegisterFile("currencies.proto", fileDescriptor_1988b70e90d5a630) }

var fileDescriptor_1988b70e90d5a630 = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xd1, 0x4a, 0xfb, 0x30,
	0x14, 0xc6, 0x49, 0xbb, 0x76, 0xdb, 0xd9, 0xff, 0xbf, 0x8d, 0x6c, 0x42, 0xf1, 0xc6, 0x52, 0x41,
	0xba, 0x0b, 0x07, 0xce, 0x27, 0x90, 0x89, 0xe0, 0x9d, 0xe4, 0x5a, 0x18, 0x5d, 0x7b, 0x06, 0xc1,
	0x35, 0xa9, 0x4d, 0x06, 0x1b, 0x3e, 0x8b, 0xaf, 0xe4, 0x33, 0x89, 0xe9, 0xd2, 0x6e, 0xca, 0x04,
	0xef, 0x7a, 0xbe, 0xef, 0xcb, 0x39, 0xf9, 0x9d, 0x14, 0x86, 0xe9, 0xa6, 0x2c, 0x51, 0xa4, 0x1c,
	0xd5, 0xb4, 0x28, 0xa5, 0x96, 0xd4, 0x4d, 0x0a, 0x7e, 0x0e, 0x28, 0x36, 0x79, 0x25, 0x44, 0x1f,
	0x04, 0x3a, 0xf3, 0x2a, 0xb5, 0xa3, 0x17, 0xd0, 0xdb, 0x9f, 0xd8, 0x2d, 0x78, 0x16, 0x90, 0x90,
	0xc4, 0x1e, 0x03, 0x2b, 0x3d, 0x66, 0x94, 0x42, 0x4b, 0x24, 0x39, 0x06, 0x4e, 0x48, 0xe2, 0x2e,
	0x33, 0xdf, 0xf4, 0x06, 0xc6, 0xf5, 0xa1, 0x0c, 0x55, 0x5a, 0xf2, 0x42, 0x73, 0x29, 0x02, 0xd7,
	0x64, 0x46, 0xd6, 0xbb, 0x6f, 0x2c, 0x3a, 0x81, 0x61, 0xce, 0xc5, 0x22, 0x95, 0x62, 0xc5, 0xcb,
	0x3c, 0x31, 0xf1, 0x96, 0x19, 0x36, 0xc8, 0xb9, 0x98, 0x1f, 0xc8, 0x34, 0x80, 0x36, 0x8a, 0x64,
	0xb9, 0xc6, 0x2c, 0xf0, 0x42, 0x12, 0x77, 0x98, 0x2d, 0xe9, 0x19, 0xf8, 0x7a, 0xbb, 0x58, 0x21,
	0x06, 0x7e, 0x48, 0x62, 0xc2, 0x3c, 0xbd, 0x7d, 0x40, 0x8c, 0x9e, 0x61, 0x60, 0x79, 0x18, 0xbe,
	0x6e, 0x50, 0x69, 0x3a, 0x06, 0x4f, 0xcb, 0x17, 0x14, 0x06, 0xa8, 0xcb, 0xaa, 0xe2, 0x3b, 0xac,
	0x73, 0x12, 0xd6, 0x6d, 0x60, 0xa3, 0x77, 0x02, 0xc3, 0xa6, 0xbd, 0x2a, 0xa4, 0x50, 0x48, 0xaf,
	0xa0, 0x5d, 0xa2, 0x4e, 0x65, 0x86, 0x66, 0x42, 0x7f, 0xf6, 0x6f, 0x9a, 0x14, 0x7c, 0xca, 0x2a,
	0x8d, 0x59, 0x93, 0x5e, 0x03, 0x34, 0x0f, 0x12, 0x38, 0xa1, 0x1b, 0xf7, 0x66, 0xff, 0x4d, 0xb4,
	0x6e, 0x79, 0x10, 0xf8, 0x42, 0x4f, 0x65, 0x9e, 0xa3, 0xd0, 0xfb, 0x2b, 0xd8, 0xd2, 0x2c, 0x65,
	0x9d, 0x14, 0x0a, 0x33, 0xb3, 0xb6, 0x2e, 0xb3, 0x65, 0xf4, 0x06, 0x23, 0xdb, 0xeb, 0x49, 0x2a,
	0xfd, 0xfb, 0x06, 0x26, 0xd0, 0xb1, 0xb8, 0x06, 0xff, 0xc7, 0x6d, 0x6a, 0x9b, 0x5e, 0x82, 0x9f,
	0xa4, 0xf5, 0xb3, 0xf6, 0x67, 0x3d, 0x13, 0xbc, 0x33, 0x12, 0xdb, 0x5b, 0x51, 0x09, 0xe3, 0xe3,
	0xe1, 0x7f, 0xdc, 0xcf, 0x01, 0xb0, 0x73, 0x12, 0xd8, 0x3d, 0x02, 0x5e, 0xfa, 0xe6, 0x37, 0xbe,
	0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xab, 0x71, 0xef, 0xeb, 0x02, 0x00, 0x00,
}
