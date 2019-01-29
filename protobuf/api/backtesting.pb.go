// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backtesting.proto

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

type Subscription struct {
	Subscribe            []*SubscriptionItems `protobuf:"bytes,1,rep,name=subscribe,proto3" json:"subscribe,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Subscription) Reset()         { *m = Subscription{} }
func (m *Subscription) String() string { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()    {}
func (*Subscription) Descriptor() ([]byte, []int) {
	return fileDescriptor_cad29974f829cd17, []int{0}
}

func (m *Subscription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subscription.Unmarshal(m, b)
}
func (m *Subscription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subscription.Marshal(b, m, deterministic)
}
func (m *Subscription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subscription.Merge(m, src)
}
func (m *Subscription) XXX_Size() int {
	return xxx_messageInfo_Subscription.Size(m)
}
func (m *Subscription) XXX_DiscardUnknown() {
	xxx_messageInfo_Subscription.DiscardUnknown(m)
}

var xxx_messageInfo_Subscription proto.InternalMessageInfo

func (m *Subscription) GetSubscribe() []*SubscriptionItems {
	if m != nil {
		return m.Subscribe
	}
	return nil
}

type SubscriptionItems struct {
	Venues               Venues   `protobuf:"varint,1,opt,name=Venues,proto3,enum=api.Venues" json:"Venues,omitempty"`
	Product              string   `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	Orderbook            bool     `protobuf:"varint,3,opt,name=orderbook,proto3" json:"orderbook,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscriptionItems) Reset()         { *m = SubscriptionItems{} }
func (m *SubscriptionItems) String() string { return proto.CompactTextString(m) }
func (*SubscriptionItems) ProtoMessage()    {}
func (*SubscriptionItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_cad29974f829cd17, []int{1}
}

func (m *SubscriptionItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriptionItems.Unmarshal(m, b)
}
func (m *SubscriptionItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriptionItems.Marshal(b, m, deterministic)
}
func (m *SubscriptionItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriptionItems.Merge(m, src)
}
func (m *SubscriptionItems) XXX_Size() int {
	return xxx_messageInfo_SubscriptionItems.Size(m)
}
func (m *SubscriptionItems) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriptionItems.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriptionItems proto.InternalMessageInfo

func (m *SubscriptionItems) GetVenues() Venues {
	if m != nil {
		return m.Venues
	}
	return Venues_DARKPOOL
}

func (m *SubscriptionItems) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *SubscriptionItems) GetOrderbook() bool {
	if m != nil {
		return m.Orderbook
	}
	return false
}

type BacktestingTokenRequest struct {
	Initialization       *ClientInitilization `protobuf:"bytes,1,opt,name=initialization,proto3" json:"initialization,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BacktestingTokenRequest) Reset()         { *m = BacktestingTokenRequest{} }
func (m *BacktestingTokenRequest) String() string { return proto.CompactTextString(m) }
func (*BacktestingTokenRequest) ProtoMessage()    {}
func (*BacktestingTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cad29974f829cd17, []int{2}
}

func (m *BacktestingTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BacktestingTokenRequest.Unmarshal(m, b)
}
func (m *BacktestingTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BacktestingTokenRequest.Marshal(b, m, deterministic)
}
func (m *BacktestingTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BacktestingTokenRequest.Merge(m, src)
}
func (m *BacktestingTokenRequest) XXX_Size() int {
	return xxx_messageInfo_BacktestingTokenRequest.Size(m)
}
func (m *BacktestingTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BacktestingTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BacktestingTokenRequest proto.InternalMessageInfo

func (m *BacktestingTokenRequest) GetInitialization() *ClientInitilization {
	if m != nil {
		return m.Initialization
	}
	return nil
}

type BacktestingTokenResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BacktestingTokenResponse) Reset()         { *m = BacktestingTokenResponse{} }
func (m *BacktestingTokenResponse) String() string { return proto.CompactTextString(m) }
func (*BacktestingTokenResponse) ProtoMessage()    {}
func (*BacktestingTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cad29974f829cd17, []int{3}
}

func (m *BacktestingTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BacktestingTokenResponse.Unmarshal(m, b)
}
func (m *BacktestingTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BacktestingTokenResponse.Marshal(b, m, deterministic)
}
func (m *BacktestingTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BacktestingTokenResponse.Merge(m, src)
}
func (m *BacktestingTokenResponse) XXX_Size() int {
	return xxx_messageInfo_BacktestingTokenResponse.Size(m)
}
func (m *BacktestingTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BacktestingTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BacktestingTokenResponse proto.InternalMessageInfo

func (m *BacktestingTokenResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *BacktestingTokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *BacktestingTokenResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type BacktestingRequest struct {
	Account              string          `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Subscription         *Subscription   `protobuf:"bytes,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
	StartDate            int64           `protobuf:"varint,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              int64           `protobuf:"varint,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	Code                 BacktestingCode `protobuf:"varint,5,opt,name=code,proto3,enum=api.BacktestingCode" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BacktestingRequest) Reset()         { *m = BacktestingRequest{} }
func (m *BacktestingRequest) String() string { return proto.CompactTextString(m) }
func (*BacktestingRequest) ProtoMessage()    {}
func (*BacktestingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cad29974f829cd17, []int{4}
}

func (m *BacktestingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BacktestingRequest.Unmarshal(m, b)
}
func (m *BacktestingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BacktestingRequest.Marshal(b, m, deterministic)
}
func (m *BacktestingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BacktestingRequest.Merge(m, src)
}
func (m *BacktestingRequest) XXX_Size() int {
	return xxx_messageInfo_BacktestingRequest.Size(m)
}
func (m *BacktestingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BacktestingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BacktestingRequest proto.InternalMessageInfo

func (m *BacktestingRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *BacktestingRequest) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (m *BacktestingRequest) GetStartDate() int64 {
	if m != nil {
		return m.StartDate
	}
	return 0
}

func (m *BacktestingRequest) GetEndDate() int64 {
	if m != nil {
		return m.EndDate
	}
	return 0
}

func (m *BacktestingRequest) GetCode() BacktestingCode {
	if m != nil {
		return m.Code
	}
	return BacktestingCode_START
}

type BacktestingResponse struct {
	Retcode              Retcode         `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Code                 BacktestingCode `protobuf:"varint,2,opt,name=code,proto3,enum=api.BacktestingCode" json:"code,omitempty"`
	Trade                *Trade          `protobuf:"bytes,3,opt,name=trade,proto3" json:"trade,omitempty"`
	Orderbook            *Orderbook      `protobuf:"bytes,4,opt,name=orderbook,proto3" json:"orderbook,omitempty"`
	Comment              string          `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BacktestingResponse) Reset()         { *m = BacktestingResponse{} }
func (m *BacktestingResponse) String() string { return proto.CompactTextString(m) }
func (*BacktestingResponse) ProtoMessage()    {}
func (*BacktestingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cad29974f829cd17, []int{5}
}

func (m *BacktestingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BacktestingResponse.Unmarshal(m, b)
}
func (m *BacktestingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BacktestingResponse.Marshal(b, m, deterministic)
}
func (m *BacktestingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BacktestingResponse.Merge(m, src)
}
func (m *BacktestingResponse) XXX_Size() int {
	return xxx_messageInfo_BacktestingResponse.Size(m)
}
func (m *BacktestingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BacktestingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BacktestingResponse proto.InternalMessageInfo

func (m *BacktestingResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *BacktestingResponse) GetCode() BacktestingCode {
	if m != nil {
		return m.Code
	}
	return BacktestingCode_START
}

func (m *BacktestingResponse) GetTrade() *Trade {
	if m != nil {
		return m.Trade
	}
	return nil
}

func (m *BacktestingResponse) GetOrderbook() *Orderbook {
	if m != nil {
		return m.Orderbook
	}
	return nil
}

func (m *BacktestingResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type ClientInitilization struct {
	AccountMode          AccountMode   `protobuf:"varint,1,opt,name=account_mode,json=accountMode,proto3,enum=api.AccountMode" json:"account_mode,omitempty"`
	Subscription         *Subscription `protobuf:"bytes,2,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Balances             []*Balance    `protobuf:"bytes,3,rep,name=balances,proto3" json:"balances,omitempty"`
	StartDate            int64         `protobuf:"varint,4,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate              int64         `protobuf:"varint,5,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	CandleGranularity    string        `protobuf:"bytes,6,opt,name=candle_granularity,json=candleGranularity,proto3" json:"candle_granularity,omitempty"`
	CandleGroupBy        CandleGroupBy `protobuf:"varint,7,opt,name=candle_group_by,json=candleGroupBy,proto3,enum=api.CandleGroupBy" json:"candle_group_by,omitempty"`
	Verbose              bool          `protobuf:"varint,8,opt,name=verbose,proto3" json:"verbose,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ClientInitilization) Reset()         { *m = ClientInitilization{} }
func (m *ClientInitilization) String() string { return proto.CompactTextString(m) }
func (*ClientInitilization) ProtoMessage()    {}
func (*ClientInitilization) Descriptor() ([]byte, []int) {
	return fileDescriptor_cad29974f829cd17, []int{6}
}

func (m *ClientInitilization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInitilization.Unmarshal(m, b)
}
func (m *ClientInitilization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInitilization.Marshal(b, m, deterministic)
}
func (m *ClientInitilization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInitilization.Merge(m, src)
}
func (m *ClientInitilization) XXX_Size() int {
	return xxx_messageInfo_ClientInitilization.Size(m)
}
func (m *ClientInitilization) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInitilization.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInitilization proto.InternalMessageInfo

func (m *ClientInitilization) GetAccountMode() AccountMode {
	if m != nil {
		return m.AccountMode
	}
	return AccountMode_DEMO
}

func (m *ClientInitilization) GetSubscription() *Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (m *ClientInitilization) GetBalances() []*Balance {
	if m != nil {
		return m.Balances
	}
	return nil
}

func (m *ClientInitilization) GetStartDate() int64 {
	if m != nil {
		return m.StartDate
	}
	return 0
}

func (m *ClientInitilization) GetEndDate() int64 {
	if m != nil {
		return m.EndDate
	}
	return 0
}

func (m *ClientInitilization) GetCandleGranularity() string {
	if m != nil {
		return m.CandleGranularity
	}
	return ""
}

func (m *ClientInitilization) GetCandleGroupBy() CandleGroupBy {
	if m != nil {
		return m.CandleGroupBy
	}
	return CandleGroupBy_TIME
}

func (m *ClientInitilization) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

func init() {
	proto.RegisterType((*Subscription)(nil), "api.Subscription")
	proto.RegisterType((*SubscriptionItems)(nil), "api.SubscriptionItems")
	proto.RegisterType((*BacktestingTokenRequest)(nil), "api.BacktestingTokenRequest")
	proto.RegisterType((*BacktestingTokenResponse)(nil), "api.BacktestingTokenResponse")
	proto.RegisterType((*BacktestingRequest)(nil), "api.BacktestingRequest")
	proto.RegisterType((*BacktestingResponse)(nil), "api.BacktestingResponse")
	proto.RegisterType((*ClientInitilization)(nil), "api.ClientInitilization")
}

func init() { proto.RegisterFile("backtesting.proto", fileDescriptor_cad29974f829cd17) }

var fileDescriptor_cad29974f829cd17 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x55, 0xf6, 0x3b, 0x93, 0x65, 0xe9, 0xba, 0x15, 0x98, 0x0a, 0xa4, 0x55, 0x90, 0x50, 0x0e,
	0xb0, 0x87, 0x2d, 0x5c, 0x38, 0x41, 0x5b, 0x09, 0xf5, 0x80, 0x90, 0x4c, 0xc5, 0x85, 0xc3, 0xca,
	0x49, 0xac, 0x2a, 0xda, 0x8d, 0x1d, 0x6c, 0x07, 0x69, 0xf9, 0x87, 0x1c, 0xf9, 0x1b, 0xfc, 0x0a,
	0x94, 0x71, 0xd2, 0x64, 0xbb, 0x08, 0x04, 0xc7, 0x79, 0xef, 0xc5, 0xcf, 0x6f, 0xc6, 0x13, 0x98,
	0xc7, 0x3c, 0xd9, 0x58, 0x61, 0x6c, 0x26, 0x6f, 0x96, 0x85, 0x56, 0x56, 0x91, 0x3e, 0x2f, 0xb2,
	0x53, 0x10, 0xb2, 0xcc, 0x1d, 0x70, 0x7a, 0x94, 0x73, 0xbd, 0x11, 0x36, 0xe5, 0x96, 0xd7, 0xc8,
	0x2c, 0xe6, 0x5b, 0x2e, 0x13, 0x61, 0x5c, 0x1d, 0x5e, 0xc2, 0xf4, 0x63, 0x19, 0x9b, 0x44, 0x67,
	0x85, 0xcd, 0x94, 0x24, 0x2f, 0xc1, 0x37, 0xae, 0x8e, 0x05, 0xf5, 0x16, 0xfd, 0x28, 0x58, 0x3d,
	0x58, 0xf2, 0x22, 0x5b, 0x76, 0x55, 0x57, 0x56, 0xe4, 0x86, 0xb5, 0xc2, 0xb0, 0x80, 0xf9, 0x01,
	0x4f, 0x9e, 0xc2, 0xe8, 0x93, 0x90, 0xa5, 0x30, 0xd4, 0x5b, 0x78, 0xd1, 0x6c, 0x15, 0xe0, 0x39,
	0x0e, 0x62, 0x35, 0x45, 0x28, 0x8c, 0x0b, 0xad, 0xd2, 0x32, 0xb1, 0xb4, 0xb7, 0xf0, 0x22, 0x9f,
	0x35, 0x25, 0x79, 0x0c, 0xbe, 0xd2, 0xa9, 0xd0, 0xb1, 0x52, 0x1b, 0xda, 0x5f, 0x78, 0xd1, 0x84,
	0xb5, 0x40, 0xf8, 0x19, 0x1e, 0x9e, 0xb7, 0xf9, 0xaf, 0xd5, 0x46, 0x48, 0x26, 0xbe, 0x94, 0xc2,
	0x58, 0xf2, 0x06, 0x66, 0x99, 0xcc, 0x6c, 0xc6, 0xb7, 0xd9, 0x37, 0x5e, 0x5d, 0x07, 0xfd, 0x83,
	0x15, 0x45, 0xff, 0x8b, 0x6d, 0x26, 0xa4, 0xbd, 0xaa, 0x04, 0x0d, 0xcf, 0xee, 0xe8, 0x43, 0x0d,
	0xf4, 0xf0, 0x70, 0x53, 0x28, 0x69, 0x04, 0x79, 0x06, 0x63, 0x2d, 0x6c, 0xa2, 0x52, 0x51, 0xc7,
	0x9a, 0xe2, 0xb1, 0xcc, 0x61, 0xac, 0x21, 0xc9, 0x09, 0x0c, 0x6d, 0xf5, 0x61, 0x1d, 0xcb, 0x15,
	0x55, 0xdc, 0x44, 0xe5, 0xb9, 0x90, 0x16, 0x23, 0xf9, 0xac, 0x29, 0xc3, 0xef, 0x1e, 0x90, 0x8e,
	0x69, 0x13, 0x86, 0xc2, 0x98, 0x27, 0x89, 0x2a, 0xa5, 0x45, 0x3b, 0x9f, 0x35, 0x25, 0x79, 0x05,
	0x53, 0xd3, 0xe9, 0x39, 0xfa, 0x04, 0xab, 0xf9, 0xc1, 0xb0, 0xd8, 0x9e, 0x8c, 0x3c, 0x01, 0x30,
	0x96, 0x6b, 0xbb, 0x4e, 0xb9, 0x15, 0x78, 0x89, 0x3e, 0xf3, 0x11, 0xb9, 0xe4, 0x56, 0x90, 0x47,
	0x30, 0x11, 0x32, 0x75, 0xe4, 0x00, 0xc9, 0xb1, 0x90, 0x29, 0x52, 0x11, 0x0c, 0x30, 0xf6, 0x10,
	0x63, 0x9f, 0xa0, 0x51, 0xe7, 0xc6, 0x17, 0x55, 0x7c, 0x54, 0x84, 0x3f, 0x3c, 0x38, 0xde, 0xcb,
	0xf2, 0x8f, 0xbd, 0x6b, 0x9c, 0x7a, 0x7f, 0x73, 0x22, 0x0b, 0x18, 0x5a, 0xcd, 0x53, 0x17, 0x24,
	0x58, 0x01, 0x4a, 0xaf, 0x2b, 0x84, 0x39, 0x82, 0x3c, 0xef, 0x3e, 0xa3, 0x01, 0xaa, 0x66, 0xa8,
	0xfa, 0xd0, 0xa0, 0x9d, 0x67, 0xd5, 0x9d, 0xcf, 0x70, 0x7f, 0x3e, 0x3f, 0x7b, 0x70, 0xfc, 0x9b,
	0xb7, 0x43, 0xce, 0x60, 0x5a, 0x4f, 0x64, 0x9d, 0xb7, 0xc1, 0x8e, 0xd0, 0xe2, 0xad, 0x23, 0xde,
	0x57, 0xf7, 0x0d, 0x78, 0x5b, 0xfc, 0xef, 0xec, 0x22, 0x98, 0x34, 0xeb, 0x4b, 0xfb, 0xb8, 0x9b,
	0xd3, 0xba, 0x37, 0x08, 0xb2, 0x5b, 0xf6, 0xce, 0x94, 0x07, 0x7f, 0x9a, 0xf2, 0x70, 0x7f, 0xca,
	0x2f, 0x80, 0x24, 0x5c, 0xa6, 0x5b, 0xb1, 0xbe, 0xd1, 0x5c, 0x96, 0x5b, 0xae, 0x33, 0xbb, 0xa3,
	0x23, 0x6c, 0xc6, 0xdc, 0x31, 0xef, 0x5a, 0x82, 0xbc, 0x86, 0xfb, 0xb7, 0x72, 0x55, 0x16, 0xeb,
	0x78, 0x47, 0xc7, 0xd8, 0x01, 0xe2, 0xb6, 0xad, 0xfe, 0x40, 0x95, 0xc5, 0xf9, 0x8e, 0xdd, 0x4b,
	0xba, 0x65, 0xd5, 0xec, 0xaf, 0x55, 0xdf, 0x8d, 0xa0, 0x13, 0xdc, 0xef, 0xa6, 0x8c, 0x47, 0xf8,
	0x73, 0x3a, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x35, 0x67, 0xdc, 0x0f, 0xe4, 0x04, 0x00, 0x00,
}
