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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

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
	return Retcode_DONE
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
	return Retcode_DONE
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
	AccountMode          SystemMode    `protobuf:"varint,1,opt,name=account_mode,json=accountMode,proto3,enum=api.SystemMode" json:"account_mode,omitempty"`
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

func (m *ClientInitilization) GetAccountMode() SystemMode {
	if m != nil {
		return m.AccountMode
	}
	return SystemMode_DEMO
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
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0xa4, 0x49, 0x26, 0x21, 0x25, 0xdb, 0x0a, 0x96, 0x0a, 0xa4, 0xc8, 0x48, 0xc8,
	0x07, 0xe8, 0xc1, 0xc0, 0x85, 0x13, 0x6a, 0x2b, 0xa1, 0x1e, 0x10, 0xd2, 0x52, 0x71, 0xe1, 0x10,
	0xad, 0xed, 0x51, 0x65, 0x25, 0xde, 0x35, 0xbb, 0x6b, 0xa4, 0xf0, 0x87, 0x1c, 0xf9, 0x0c, 0xfe,
	0x02, 0x79, 0xd6, 0x6e, 0x9c, 0x06, 0x81, 0xe0, 0x38, 0xef, 0x3d, 0xef, 0xec, 0x9b, 0xb7, 0x63,
	0x98, 0x27, 0x32, 0x5d, 0x39, 0xb4, 0x2e, 0x57, 0x37, 0x67, 0xa5, 0xd1, 0x4e, 0xb3, 0x9e, 0x2c,
	0xf3, 0x53, 0x40, 0x55, 0x15, 0x1e, 0x38, 0xbd, 0x5f, 0x48, 0xb3, 0x42, 0x97, 0x49, 0x27, 0x1b,
	0x64, 0x96, 0xc8, 0xb5, 0x54, 0x29, 0x5a, 0x5f, 0x87, 0x97, 0x30, 0xfd, 0x58, 0x25, 0x36, 0x35,
	0x79, 0xe9, 0x72, 0xad, 0xd8, 0x2b, 0x18, 0x5b, 0x5f, 0x27, 0xc8, 0x83, 0x45, 0x2f, 0x9a, 0xc4,
	0x0f, 0xce, 0x64, 0x99, 0x9f, 0x75, 0x55, 0x57, 0x0e, 0x0b, 0x2b, 0xb6, 0xc2, 0xb0, 0x84, 0xf9,
	0x1e, 0xcf, 0x9e, 0xc2, 0xe1, 0x27, 0x54, 0x15, 0x5a, 0x1e, 0x2c, 0x82, 0x68, 0x16, 0x4f, 0xe8,
	0x1c, 0x0f, 0x89, 0x86, 0x62, 0x1c, 0x86, 0xa5, 0xd1, 0x59, 0x95, 0x3a, 0x7e, 0xb0, 0x08, 0xa2,
	0xb1, 0x68, 0x4b, 0xf6, 0x18, 0xc6, 0xda, 0x64, 0x68, 0x12, 0xad, 0x57, 0xbc, 0xb7, 0x08, 0xa2,
	0x91, 0xd8, 0x02, 0xe1, 0x67, 0x78, 0x78, 0xbe, 0xf5, 0x7f, 0xad, 0x57, 0xa8, 0x04, 0x7e, 0xa9,
	0xd0, 0x3a, 0xf6, 0x16, 0x66, 0xb9, 0xca, 0x5d, 0x2e, 0xd7, 0xf9, 0x37, 0x59, 0x5f, 0x87, 0xfa,
	0x4f, 0x62, 0x4e, 0xfd, 0x2f, 0xd6, 0x39, 0x2a, 0x77, 0x55, 0x0b, 0x5a, 0x5e, 0xdc, 0xd1, 0x87,
	0x06, 0xf8, 0xfe, 0xe1, 0xb6, 0xd4, 0xca, 0x22, 0x7b, 0x06, 0x43, 0x83, 0x2e, 0xd5, 0x19, 0x36,
	0xb6, 0xa6, 0x74, 0xac, 0xf0, 0x98, 0x68, 0x49, 0x76, 0x02, 0x03, 0x57, 0x7f, 0xd8, 0xd8, 0xf2,
	0x45, 0x6d, 0x37, 0xd5, 0x45, 0x81, 0xca, 0x91, 0xa5, 0xb1, 0x68, 0xcb, 0xf0, 0x7b, 0x00, 0xac,
	0xd3, 0xb4, 0x35, 0xc3, 0x61, 0x28, 0xd3, 0x54, 0x57, 0xca, 0x51, 0xbb, 0xb1, 0x68, 0x4b, 0xf6,
	0x1a, 0xa6, 0xb6, 0x33, 0x73, 0xea, 0x33, 0x89, 0xe7, 0x7b, 0x61, 0x89, 0x1d, 0x19, 0x7b, 0x02,
	0x60, 0x9d, 0x34, 0x6e, 0x99, 0x49, 0x87, 0x74, 0x89, 0x9e, 0x18, 0x13, 0x72, 0x29, 0x1d, 0xb2,
	0x47, 0x30, 0x42, 0x95, 0x79, 0xb2, 0x4f, 0xe4, 0x10, 0x55, 0x46, 0x54, 0x04, 0x7d, 0xb2, 0x3d,
	0x20, 0xdb, 0x27, 0xd4, 0xa8, 0x73, 0xe3, 0x8b, 0xda, 0x3e, 0x29, 0xc2, 0x1f, 0x01, 0x1c, 0xef,
	0x78, 0xf9, 0xc7, 0xd9, 0xb5, 0x9d, 0x0e, 0xfe, 0xd6, 0x89, 0x2d, 0x60, 0xe0, 0x8c, 0xcc, 0xbc,
	0x91, 0x49, 0x0c, 0x24, 0xbd, 0xae, 0x11, 0xe1, 0x09, 0xf6, 0xbc, 0xfb, 0x8c, 0xfa, 0xa4, 0x9a,
	0x91, 0xea, 0x43, 0x8b, 0x76, 0x9e, 0x55, 0x37, 0x9f, 0xc1, 0x6e, 0x3e, 0x3f, 0x0f, 0xe0, 0xf8,
	0x37, 0x6f, 0x87, 0xc5, 0x30, 0x6d, 0x12, 0x59, 0x16, 0x5b, 0x63, 0x47, 0x3e, 0x86, 0x8d, 0x75,
	0x58, 0xbc, 0xaf, 0xaf, 0x3b, 0x69, 0x44, 0x75, 0xf1, 0xbf, 0xd1, 0x45, 0x30, 0x6a, 0xb7, 0x97,
	0xf7, 0x68, 0x35, 0xa7, 0xcd, 0x68, 0x08, 0x14, 0xb7, 0xec, 0x9d, 0x90, 0xfb, 0x7f, 0x0a, 0x79,
	0xb0, 0x1b, 0xf2, 0x0b, 0x60, 0xa9, 0x54, 0xd9, 0x1a, 0x97, 0x37, 0x46, 0xaa, 0x6a, 0x2d, 0x4d,
	0xee, 0x36, 0xfc, 0x90, 0x66, 0x31, 0xf7, 0xcc, 0xbb, 0x2d, 0xc1, 0xde, 0xc0, 0xd1, 0xad, 0x5c,
	0x57, 0xe5, 0x32, 0xd9, 0xf0, 0x21, 0x0d, 0x80, 0xf9, 0x65, 0x6b, 0x3e, 0xd0, 0x55, 0x79, 0xbe,
	0x11, 0xf7, 0xd2, 0x6e, 0x59, 0xcf, 0xfa, 0x6b, 0x3d, 0x76, 0x8b, 0x7c, 0x44, 0xeb, 0xdd, 0x96,
	0xc9, 0x21, 0xfd, 0x9b, 0x5e, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x3a, 0x10, 0x28, 0xe3,
	0x04, 0x00, 0x00,
}
