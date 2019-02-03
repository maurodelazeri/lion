// Code generated by protoc-gen-go. DO NOT EDIT.
// source: balances.proto

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

type Balance struct {
	BalanceId            int32    `protobuf:"varint,1,opt,name=balance_id,json=balanceId,proto3" json:"balance_id,omitempty"`
	CurrencyId           int32    `protobuf:"varint,2,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
	CurrencyDescription  string   `protobuf:"bytes,3,opt,name=currency_description,json=currencyDescription,proto3" json:"currency_description,omitempty"`
	AccountId            int32    `protobuf:"varint,4,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AccountDescription   string   `protobuf:"bytes,5,opt,name=account_description,json=accountDescription,proto3" json:"account_description,omitempty"`
	UserId               int32    `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VenueId              int32    `protobuf:"varint,7,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	VenueDescription     string   `protobuf:"bytes,8,opt,name=venue_description,json=venueDescription,proto3" json:"venue_description,omitempty"`
	Available            float64  `protobuf:"fixed64,9,opt,name=available,proto3" json:"available,omitempty"`
	Hold                 float64  `protobuf:"fixed64,10,opt,name=hold,proto3" json:"hold,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Balance) Reset()         { *m = Balance{} }
func (m *Balance) String() string { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()    {}
func (*Balance) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{0}
}

func (m *Balance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Balance.Unmarshal(m, b)
}
func (m *Balance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Balance.Marshal(b, m, deterministic)
}
func (m *Balance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balance.Merge(m, src)
}
func (m *Balance) XXX_Size() int {
	return xxx_messageInfo_Balance.Size(m)
}
func (m *Balance) XXX_DiscardUnknown() {
	xxx_messageInfo_Balance.DiscardUnknown(m)
}

var xxx_messageInfo_Balance proto.InternalMessageInfo

func (m *Balance) GetBalanceId() int32 {
	if m != nil {
		return m.BalanceId
	}
	return 0
}

func (m *Balance) GetCurrencyId() int32 {
	if m != nil {
		return m.CurrencyId
	}
	return 0
}

func (m *Balance) GetCurrencyDescription() string {
	if m != nil {
		return m.CurrencyDescription
	}
	return ""
}

func (m *Balance) GetAccountId() int32 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *Balance) GetAccountDescription() string {
	if m != nil {
		return m.AccountDescription
	}
	return ""
}

func (m *Balance) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Balance) GetVenueId() int32 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *Balance) GetVenueDescription() string {
	if m != nil {
		return m.VenueDescription
	}
	return ""
}

func (m *Balance) GetAvailable() float64 {
	if m != nil {
		return m.Available
	}
	return 0
}

func (m *Balance) GetHold() float64 {
	if m != nil {
		return m.Hold
	}
	return 0
}

type BalanceRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	BalanceId            int32    `protobuf:"varint,2,opt,name=balance_id,json=balanceId,proto3" json:"balance_id,omitempty"`
	AccountId            int32    `protobuf:"varint,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceRequest) Reset()         { *m = BalanceRequest{} }
func (m *BalanceRequest) String() string { return proto.CompactTextString(m) }
func (*BalanceRequest) ProtoMessage()    {}
func (*BalanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{1}
}

func (m *BalanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceRequest.Unmarshal(m, b)
}
func (m *BalanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceRequest.Marshal(b, m, deterministic)
}
func (m *BalanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceRequest.Merge(m, src)
}
func (m *BalanceRequest) XXX_Size() int {
	return xxx_messageInfo_BalanceRequest.Size(m)
}
func (m *BalanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceRequest proto.InternalMessageInfo

func (m *BalanceRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *BalanceRequest) GetBalanceId() int32 {
	if m != nil {
		return m.BalanceId
	}
	return 0
}

func (m *BalanceRequest) GetAccountId() int32 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

type BalanceResponse struct {
	Retcode              Retcode    `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Balance              []*Balance `protobuf:"bytes,2,rep,name=balance,proto3" json:"balance,omitempty"`
	Comment              string     `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string     `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BalanceResponse) Reset()         { *m = BalanceResponse{} }
func (m *BalanceResponse) String() string { return proto.CompactTextString(m) }
func (*BalanceResponse) ProtoMessage()    {}
func (*BalanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{2}
}

func (m *BalanceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceResponse.Unmarshal(m, b)
}
func (m *BalanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceResponse.Marshal(b, m, deterministic)
}
func (m *BalanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceResponse.Merge(m, src)
}
func (m *BalanceResponse) XXX_Size() int {
	return xxx_messageInfo_BalanceResponse.Size(m)
}
func (m *BalanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceResponse proto.InternalMessageInfo

func (m *BalanceResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *BalanceResponse) GetBalance() []*Balance {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *BalanceResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *BalanceResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type BalancePostRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Balance              *Balance `protobuf:"bytes,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Action               Action   `protobuf:"varint,3,opt,name=action,proto3,enum=api.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalancePostRequest) Reset()         { *m = BalancePostRequest{} }
func (m *BalancePostRequest) String() string { return proto.CompactTextString(m) }
func (*BalancePostRequest) ProtoMessage()    {}
func (*BalancePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{3}
}

func (m *BalancePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalancePostRequest.Unmarshal(m, b)
}
func (m *BalancePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalancePostRequest.Marshal(b, m, deterministic)
}
func (m *BalancePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalancePostRequest.Merge(m, src)
}
func (m *BalancePostRequest) XXX_Size() int {
	return xxx_messageInfo_BalancePostRequest.Size(m)
}
func (m *BalancePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BalancePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BalancePostRequest proto.InternalMessageInfo

func (m *BalancePostRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *BalancePostRequest) GetBalance() *Balance {
	if m != nil {
		return m.Balance
	}
	return nil
}

func (m *BalancePostRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_INSERT
}

type BalancePostResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,3,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalancePostResponse) Reset()         { *m = BalancePostResponse{} }
func (m *BalancePostResponse) String() string { return proto.CompactTextString(m) }
func (*BalancePostResponse) ProtoMessage()    {}
func (*BalancePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{4}
}

func (m *BalancePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalancePostResponse.Unmarshal(m, b)
}
func (m *BalancePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalancePostResponse.Marshal(b, m, deterministic)
}
func (m *BalancePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalancePostResponse.Merge(m, src)
}
func (m *BalancePostResponse) XXX_Size() int {
	return xxx_messageInfo_BalancePostResponse.Size(m)
}
func (m *BalancePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BalancePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BalancePostResponse proto.InternalMessageInfo

func (m *BalancePostResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *BalancePostResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *BalancePostResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

func init() {
	proto.RegisterType((*Balance)(nil), "api.Balance")
	proto.RegisterType((*BalanceRequest)(nil), "api.BalanceRequest")
	proto.RegisterType((*BalanceResponse)(nil), "api.BalanceResponse")
	proto.RegisterType((*BalancePostRequest)(nil), "api.BalancePostRequest")
	proto.RegisterType((*BalancePostResponse)(nil), "api.BalancePostResponse")
}

func init() { proto.RegisterFile("balances.proto", fileDescriptor_9d5f5974e1c89e1a) }

var fileDescriptor_9d5f5974e1c89e1a = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xe1, 0xea, 0xd3, 0x30,
	0x14, 0xc5, 0x69, 0xfb, 0x5f, 0xfb, 0xef, 0x9d, 0x54, 0xcd, 0xfe, 0x60, 0x15, 0xc5, 0x52, 0x41,
	0x0a, 0xc2, 0xc4, 0xf9, 0x04, 0x8a, 0x5f, 0xfa, 0x4d, 0xf2, 0x02, 0x92, 0x25, 0x17, 0x2c, 0x76,
	0x49, 0xd7, 0xa4, 0x13, 0x1f, 0xc4, 0xe7, 0xf3, 0x55, 0xa4, 0x69, 0xb2, 0xb5, 0x83, 0x09, 0x7e,
	0xdb, 0x3d, 0xbf, 0xcb, 0x39, 0xbd, 0xa7, 0x2b, 0x64, 0x7b, 0xd6, 0x32, 0xc9, 0x51, 0x6f, 0xbb,
	0x5e, 0x19, 0x45, 0x22, 0xd6, 0x35, 0x2f, 0x00, 0xe5, 0x70, 0x98, 0x84, 0xf2, 0x4f, 0x08, 0xc9,
	0xe7, 0x69, 0x87, 0xbc, 0x02, 0x70, 0xeb, 0xdf, 0x1a, 0x91, 0x07, 0x45, 0x50, 0xad, 0x68, 0xea,
	0x94, 0x5a, 0x90, 0xd7, 0xb0, 0xe6, 0x43, 0xdf, 0xa3, 0xe4, 0xbf, 0x46, 0x1e, 0x5a, 0x0e, 0x5e,
	0xaa, 0x05, 0xf9, 0x00, 0x0f, 0xe7, 0x05, 0x81, 0x9a, 0xf7, 0x4d, 0x67, 0x1a, 0x25, 0xf3, 0xa8,
	0x08, 0xaa, 0x94, 0x6e, 0x3c, 0xfb, 0x72, 0x41, 0x63, 0x24, 0xe3, 0x5c, 0x0d, 0xd2, 0x8c, 0x96,
	0x77, 0x53, 0xa4, 0x53, 0x6a, 0x41, 0xde, 0xc3, 0xc6, 0xe3, 0xb9, 0xe1, 0xca, 0x1a, 0x12, 0x87,
	0xe6, 0x7e, 0xcf, 0x20, 0x19, 0x34, 0xf6, 0xa3, 0x59, 0x6c, 0xcd, 0xe2, 0x71, 0xac, 0x05, 0x79,
	0x0e, 0xf7, 0x27, 0x94, 0x83, 0xbd, 0x2c, 0xb1, 0x24, 0xb1, 0x73, 0x2d, 0xc8, 0x3b, 0x78, 0x3a,
	0xa1, 0x79, 0xc4, 0xbd, 0x8d, 0x78, 0x62, 0xc1, 0x3c, 0xe0, 0x25, 0xa4, 0xec, 0xc4, 0x9a, 0x96,
	0xed, 0x5b, 0xcc, 0xd3, 0x22, 0xa8, 0x02, 0x7a, 0x11, 0x08, 0x81, 0xbb, 0xef, 0xaa, 0x15, 0x39,
	0x58, 0x60, 0x7f, 0x97, 0x02, 0x32, 0x57, 0x30, 0xc5, 0xe3, 0x80, 0xda, 0x90, 0x07, 0x58, 0x19,
	0xf5, 0x03, 0xa5, 0xad, 0x38, 0xa5, 0xd3, 0x70, 0xd5, 0x7e, 0x78, 0xdd, 0xfe, 0xb2, 0xa9, 0xe8,
	0xaa, 0xa9, 0xf2, 0x77, 0x00, 0x8f, 0xcf, 0x31, 0xba, 0x53, 0x52, 0x23, 0x79, 0x0b, 0x49, 0x8f,
	0x86, 0x2b, 0x81, 0x36, 0x29, 0xdb, 0x3d, 0xda, 0xb2, 0xae, 0xd9, 0xd2, 0x49, 0xa3, 0x1e, 0x8e,
	0x7b, 0x2e, 0x27, 0x0f, 0x8b, 0xa8, 0x5a, 0xbb, 0x3d, 0x6f, 0xe7, 0x21, 0xc9, 0x21, 0xe1, 0xea,
	0x70, 0x40, 0x69, 0xdc, 0x2b, 0xf5, 0xe3, 0x48, 0xb0, 0x65, 0x9d, 0xc6, 0xe9, 0x1d, 0xa6, 0xd4,
	0x8f, 0xe5, 0x4f, 0x20, 0xce, 0xe7, 0xab, 0xd2, 0xe6, 0xdf, 0x0d, 0x2c, 0x9e, 0x23, 0xb8, 0xfd,
	0x1c, 0x6f, 0x20, 0x66, 0xfc, 0xfc, 0xcf, 0xca, 0x76, 0x6b, 0xbb, 0xf6, 0xc9, 0x4a, 0xd4, 0xa1,
	0xf2, 0x08, 0x9b, 0x45, 0xf0, 0x7f, 0x76, 0x32, 0xbb, 0x35, 0xbc, 0x79, 0x6b, 0xb4, 0xb8, 0x75,
	0x1f, 0xdb, 0x4f, 0xea, 0xe3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0xae, 0xfa, 0x1b, 0x75,
	0x03, 0x00, 0x00,
}
