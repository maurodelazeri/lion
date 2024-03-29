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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Balance struct {
	BalanceId            int64    `protobuf:"varint,1,opt,name=balance_id,json=balanceId,proto3" json:"balance_id,omitempty"`
	CurrencyId           int64    `protobuf:"varint,2,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
	CurrencyDescription  string   `protobuf:"bytes,3,opt,name=currency_description,json=currencyDescription,proto3" json:"currency_description,omitempty"`
	UserId               int64    `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Balance              float64  `protobuf:"fixed64,5,opt,name=balance,proto3" json:"balance,omitempty"`
	Available            float64  `protobuf:"fixed64,6,opt,name=available,proto3" json:"available,omitempty"`
	Hold                 float64  `protobuf:"fixed64,7,opt,name=hold,proto3" json:"hold,omitempty"`
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

func (m *Balance) GetBalanceId() int64 {
	if m != nil {
		return m.BalanceId
	}
	return 0
}

func (m *Balance) GetCurrencyId() int64 {
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

func (m *Balance) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Balance) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
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
	BalanceMode          string   `protobuf:"bytes,2,opt,name=balance_mode,json=balanceMode,proto3" json:"balance_mode,omitempty"`
	BalanceId            int64    `protobuf:"varint,3,opt,name=balance_id,json=balanceId,proto3" json:"balance_id,omitempty"`
	CurrencyId           int64    `protobuf:"varint,4,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
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

func (m *BalanceRequest) GetBalanceMode() string {
	if m != nil {
		return m.BalanceMode
	}
	return ""
}

func (m *BalanceRequest) GetBalanceId() int64 {
	if m != nil {
		return m.BalanceId
	}
	return 0
}

func (m *BalanceRequest) GetCurrencyId() int64 {
	if m != nil {
		return m.CurrencyId
	}
	return 0
}

type BalanceResponse struct {
	Retcode              Retcode    `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Balances             []*Balance `protobuf:"bytes,2,rep,name=balances,proto3" json:"balances,omitempty"`
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
	return Retcode_DONE
}

func (m *BalanceResponse) GetBalances() []*Balance {
	if m != nil {
		return m.Balances
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

type AccountFunds struct {
	AccountFundId        int64    `protobuf:"varint,1,opt,name=account_fund_id,json=accountFundId,proto3" json:"account_fund_id,omitempty"`
	CurrencyId           int64    `protobuf:"varint,2,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
	CurrencyDescription  string   `protobuf:"bytes,3,opt,name=currency_description,json=currencyDescription,proto3" json:"currency_description,omitempty"`
	AccountId            int64    `protobuf:"varint,4,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AccountDescription   string   `protobuf:"bytes,5,opt,name=account_description,json=accountDescription,proto3" json:"account_description,omitempty"`
	FundMode             string   `protobuf:"bytes,6,opt,name=fund_mode,json=fundMode,proto3" json:"fund_mode,omitempty"`
	UserId               int64    `protobuf:"varint,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VenueId              int64    `protobuf:"varint,8,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	VenueDescription     string   `protobuf:"bytes,9,opt,name=venue_description,json=venueDescription,proto3" json:"venue_description,omitempty"`
	Balance              float64  `protobuf:"fixed64,10,opt,name=balance,proto3" json:"balance,omitempty"`
	Available            float64  `protobuf:"fixed64,11,opt,name=available,proto3" json:"available,omitempty"`
	Hold                 float64  `protobuf:"fixed64,12,opt,name=hold,proto3" json:"hold,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountFunds) Reset()         { *m = AccountFunds{} }
func (m *AccountFunds) String() string { return proto.CompactTextString(m) }
func (*AccountFunds) ProtoMessage()    {}
func (*AccountFunds) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{3}
}

func (m *AccountFunds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountFunds.Unmarshal(m, b)
}
func (m *AccountFunds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountFunds.Marshal(b, m, deterministic)
}
func (m *AccountFunds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountFunds.Merge(m, src)
}
func (m *AccountFunds) XXX_Size() int {
	return xxx_messageInfo_AccountFunds.Size(m)
}
func (m *AccountFunds) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountFunds.DiscardUnknown(m)
}

var xxx_messageInfo_AccountFunds proto.InternalMessageInfo

func (m *AccountFunds) GetAccountFundId() int64 {
	if m != nil {
		return m.AccountFundId
	}
	return 0
}

func (m *AccountFunds) GetCurrencyId() int64 {
	if m != nil {
		return m.CurrencyId
	}
	return 0
}

func (m *AccountFunds) GetCurrencyDescription() string {
	if m != nil {
		return m.CurrencyDescription
	}
	return ""
}

func (m *AccountFunds) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *AccountFunds) GetAccountDescription() string {
	if m != nil {
		return m.AccountDescription
	}
	return ""
}

func (m *AccountFunds) GetFundMode() string {
	if m != nil {
		return m.FundMode
	}
	return ""
}

func (m *AccountFunds) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AccountFunds) GetVenueId() int64 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *AccountFunds) GetVenueDescription() string {
	if m != nil {
		return m.VenueDescription
	}
	return ""
}

func (m *AccountFunds) GetBalance() float64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *AccountFunds) GetAvailable() float64 {
	if m != nil {
		return m.Available
	}
	return 0
}

func (m *AccountFunds) GetHold() float64 {
	if m != nil {
		return m.Hold
	}
	return 0
}

type Fund struct {
	AccountId            int64    `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	VenueId              int64    `protobuf:"varint,2,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	CurrencyId           int64    `protobuf:"varint,3,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
	Amount               float64  `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Fund) Reset()         { *m = Fund{} }
func (m *Fund) String() string { return proto.CompactTextString(m) }
func (*Fund) ProtoMessage()    {}
func (*Fund) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{4}
}

func (m *Fund) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fund.Unmarshal(m, b)
}
func (m *Fund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fund.Marshal(b, m, deterministic)
}
func (m *Fund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fund.Merge(m, src)
}
func (m *Fund) XXX_Size() int {
	return xxx_messageInfo_Fund.Size(m)
}
func (m *Fund) XXX_DiscardUnknown() {
	xxx_messageInfo_Fund.DiscardUnknown(m)
}

var xxx_messageInfo_Fund proto.InternalMessageInfo

func (m *Fund) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *Fund) GetVenueId() int64 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *Fund) GetCurrencyId() int64 {
	if m != nil {
		return m.CurrencyId
	}
	return 0
}

func (m *Fund) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type WithdrawalFund struct {
	AccountFundId        int64    `protobuf:"varint,1,opt,name=account_fund_id,json=accountFundId,proto3" json:"account_fund_id,omitempty"`
	AccountId            int64    `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Amount               float64  `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	WithdrawAll          bool     `protobuf:"varint,4,opt,name=withdraw_all,json=withdrawAll,proto3" json:"withdraw_all,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithdrawalFund) Reset()         { *m = WithdrawalFund{} }
func (m *WithdrawalFund) String() string { return proto.CompactTextString(m) }
func (*WithdrawalFund) ProtoMessage()    {}
func (*WithdrawalFund) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{5}
}

func (m *WithdrawalFund) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawalFund.Unmarshal(m, b)
}
func (m *WithdrawalFund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawalFund.Marshal(b, m, deterministic)
}
func (m *WithdrawalFund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawalFund.Merge(m, src)
}
func (m *WithdrawalFund) XXX_Size() int {
	return xxx_messageInfo_WithdrawalFund.Size(m)
}
func (m *WithdrawalFund) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawalFund.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawalFund proto.InternalMessageInfo

func (m *WithdrawalFund) GetAccountFundId() int64 {
	if m != nil {
		return m.AccountFundId
	}
	return 0
}

func (m *WithdrawalFund) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *WithdrawalFund) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *WithdrawalFund) GetWithdrawAll() bool {
	if m != nil {
		return m.WithdrawAll
	}
	return false
}

type Transfer struct {
	AccountFundId        int64    `protobuf:"varint,1,opt,name=account_fund_id,json=accountFundId,proto3" json:"account_fund_id,omitempty"`
	AccountIdDestination int64    `protobuf:"varint,2,opt,name=account_id_destination,json=accountIdDestination,proto3" json:"account_id_destination,omitempty"`
	VenueIdDestination   int64    `protobuf:"varint,3,opt,name=venue_id_destination,json=venueIdDestination,proto3" json:"venue_id_destination,omitempty"`
	Amount               float64  `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transfer) Reset()         { *m = Transfer{} }
func (m *Transfer) String() string { return proto.CompactTextString(m) }
func (*Transfer) ProtoMessage()    {}
func (*Transfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{6}
}

func (m *Transfer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transfer.Unmarshal(m, b)
}
func (m *Transfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transfer.Marshal(b, m, deterministic)
}
func (m *Transfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transfer.Merge(m, src)
}
func (m *Transfer) XXX_Size() int {
	return xxx_messageInfo_Transfer.Size(m)
}
func (m *Transfer) XXX_DiscardUnknown() {
	xxx_messageInfo_Transfer.DiscardUnknown(m)
}

var xxx_messageInfo_Transfer proto.InternalMessageInfo

func (m *Transfer) GetAccountFundId() int64 {
	if m != nil {
		return m.AccountFundId
	}
	return 0
}

func (m *Transfer) GetAccountIdDestination() int64 {
	if m != nil {
		return m.AccountIdDestination
	}
	return 0
}

func (m *Transfer) GetVenueIdDestination() int64 {
	if m != nil {
		return m.VenueIdDestination
	}
	return 0
}

func (m *Transfer) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type AccountFundsRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	AccountFundId        int64    `protobuf:"varint,2,opt,name=account_fund_id,json=accountFundId,proto3" json:"account_fund_id,omitempty"`
	AccountId            int64    `protobuf:"varint,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	VenueId              int64    `protobuf:"varint,4,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	CurrencyId           int64    `protobuf:"varint,5,opt,name=currency_id,json=currencyId,proto3" json:"currency_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountFundsRequest) Reset()         { *m = AccountFundsRequest{} }
func (m *AccountFundsRequest) String() string { return proto.CompactTextString(m) }
func (*AccountFundsRequest) ProtoMessage()    {}
func (*AccountFundsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{7}
}

func (m *AccountFundsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountFundsRequest.Unmarshal(m, b)
}
func (m *AccountFundsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountFundsRequest.Marshal(b, m, deterministic)
}
func (m *AccountFundsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountFundsRequest.Merge(m, src)
}
func (m *AccountFundsRequest) XXX_Size() int {
	return xxx_messageInfo_AccountFundsRequest.Size(m)
}
func (m *AccountFundsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountFundsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountFundsRequest proto.InternalMessageInfo

func (m *AccountFundsRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AccountFundsRequest) GetAccountFundId() int64 {
	if m != nil {
		return m.AccountFundId
	}
	return 0
}

func (m *AccountFundsRequest) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *AccountFundsRequest) GetVenueId() int64 {
	if m != nil {
		return m.VenueId
	}
	return 0
}

func (m *AccountFundsRequest) GetCurrencyId() int64 {
	if m != nil {
		return m.CurrencyId
	}
	return 0
}

type AccountFundsResponse struct {
	Retcode              Retcode         `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	AccountFunds         []*AccountFunds `protobuf:"bytes,2,rep,name=account_funds,json=accountFunds,proto3" json:"account_funds,omitempty"`
	Comment              string          `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string          `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AccountFundsResponse) Reset()         { *m = AccountFundsResponse{} }
func (m *AccountFundsResponse) String() string { return proto.CompactTextString(m) }
func (*AccountFundsResponse) ProtoMessage()    {}
func (*AccountFundsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{8}
}

func (m *AccountFundsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountFundsResponse.Unmarshal(m, b)
}
func (m *AccountFundsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountFundsResponse.Marshal(b, m, deterministic)
}
func (m *AccountFundsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountFundsResponse.Merge(m, src)
}
func (m *AccountFundsResponse) XXX_Size() int {
	return xxx_messageInfo_AccountFundsResponse.Size(m)
}
func (m *AccountFundsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountFundsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountFundsResponse proto.InternalMessageInfo

func (m *AccountFundsResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_DONE
}

func (m *AccountFundsResponse) GetAccountFunds() []*AccountFunds {
	if m != nil {
		return m.AccountFunds
	}
	return nil
}

func (m *AccountFundsResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *AccountFundsResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type AccountFundPostRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Fund                 *Fund    `protobuf:"bytes,2,opt,name=fund,proto3" json:"fund,omitempty"`
	Action               Action   `protobuf:"varint,3,opt,name=action,proto3,enum=api.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountFundPostRequest) Reset()         { *m = AccountFundPostRequest{} }
func (m *AccountFundPostRequest) String() string { return proto.CompactTextString(m) }
func (*AccountFundPostRequest) ProtoMessage()    {}
func (*AccountFundPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{9}
}

func (m *AccountFundPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountFundPostRequest.Unmarshal(m, b)
}
func (m *AccountFundPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountFundPostRequest.Marshal(b, m, deterministic)
}
func (m *AccountFundPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountFundPostRequest.Merge(m, src)
}
func (m *AccountFundPostRequest) XXX_Size() int {
	return xxx_messageInfo_AccountFundPostRequest.Size(m)
}
func (m *AccountFundPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountFundPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountFundPostRequest proto.InternalMessageInfo

func (m *AccountFundPostRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AccountFundPostRequest) GetFund() *Fund {
	if m != nil {
		return m.Fund
	}
	return nil
}

func (m *AccountFundPostRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_INSERT
}

type AccountFundPostResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,3,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountFundPostResponse) Reset()         { *m = AccountFundPostResponse{} }
func (m *AccountFundPostResponse) String() string { return proto.CompactTextString(m) }
func (*AccountFundPostResponse) ProtoMessage()    {}
func (*AccountFundPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{10}
}

func (m *AccountFundPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountFundPostResponse.Unmarshal(m, b)
}
func (m *AccountFundPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountFundPostResponse.Marshal(b, m, deterministic)
}
func (m *AccountFundPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountFundPostResponse.Merge(m, src)
}
func (m *AccountFundPostResponse) XXX_Size() int {
	return xxx_messageInfo_AccountFundPostResponse.Size(m)
}
func (m *AccountFundPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountFundPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountFundPostResponse proto.InternalMessageInfo

func (m *AccountFundPostResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_DONE
}

func (m *AccountFundPostResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *AccountFundPostResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type AccountTransferPostRequest struct {
	Token                string    `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Transfer             *Transfer `protobuf:"bytes,2,opt,name=transfer,proto3" json:"transfer,omitempty"`
	Action               Action    `protobuf:"varint,3,opt,name=action,proto3,enum=api.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AccountTransferPostRequest) Reset()         { *m = AccountTransferPostRequest{} }
func (m *AccountTransferPostRequest) String() string { return proto.CompactTextString(m) }
func (*AccountTransferPostRequest) ProtoMessage()    {}
func (*AccountTransferPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{11}
}

func (m *AccountTransferPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountTransferPostRequest.Unmarshal(m, b)
}
func (m *AccountTransferPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountTransferPostRequest.Marshal(b, m, deterministic)
}
func (m *AccountTransferPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountTransferPostRequest.Merge(m, src)
}
func (m *AccountTransferPostRequest) XXX_Size() int {
	return xxx_messageInfo_AccountTransferPostRequest.Size(m)
}
func (m *AccountTransferPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountTransferPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountTransferPostRequest proto.InternalMessageInfo

func (m *AccountTransferPostRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AccountTransferPostRequest) GetTransfer() *Transfer {
	if m != nil {
		return m.Transfer
	}
	return nil
}

func (m *AccountTransferPostRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_INSERT
}

type AccountTransferPostResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,3,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountTransferPostResponse) Reset()         { *m = AccountTransferPostResponse{} }
func (m *AccountTransferPostResponse) String() string { return proto.CompactTextString(m) }
func (*AccountTransferPostResponse) ProtoMessage()    {}
func (*AccountTransferPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{12}
}

func (m *AccountTransferPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountTransferPostResponse.Unmarshal(m, b)
}
func (m *AccountTransferPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountTransferPostResponse.Marshal(b, m, deterministic)
}
func (m *AccountTransferPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountTransferPostResponse.Merge(m, src)
}
func (m *AccountTransferPostResponse) XXX_Size() int {
	return xxx_messageInfo_AccountTransferPostResponse.Size(m)
}
func (m *AccountTransferPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountTransferPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountTransferPostResponse proto.InternalMessageInfo

func (m *AccountTransferPostResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_DONE
}

func (m *AccountTransferPostResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *AccountTransferPostResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type AccountWithdrawalPostRequest struct {
	Token                string          `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	WithdrawalFund       *WithdrawalFund `protobuf:"bytes,2,opt,name=withdrawal_fund,json=withdrawalFund,proto3" json:"withdrawal_fund,omitempty"`
	Action               Action          `protobuf:"varint,3,opt,name=action,proto3,enum=api.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AccountWithdrawalPostRequest) Reset()         { *m = AccountWithdrawalPostRequest{} }
func (m *AccountWithdrawalPostRequest) String() string { return proto.CompactTextString(m) }
func (*AccountWithdrawalPostRequest) ProtoMessage()    {}
func (*AccountWithdrawalPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{13}
}

func (m *AccountWithdrawalPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountWithdrawalPostRequest.Unmarshal(m, b)
}
func (m *AccountWithdrawalPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountWithdrawalPostRequest.Marshal(b, m, deterministic)
}
func (m *AccountWithdrawalPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountWithdrawalPostRequest.Merge(m, src)
}
func (m *AccountWithdrawalPostRequest) XXX_Size() int {
	return xxx_messageInfo_AccountWithdrawalPostRequest.Size(m)
}
func (m *AccountWithdrawalPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountWithdrawalPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountWithdrawalPostRequest proto.InternalMessageInfo

func (m *AccountWithdrawalPostRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AccountWithdrawalPostRequest) GetWithdrawalFund() *WithdrawalFund {
	if m != nil {
		return m.WithdrawalFund
	}
	return nil
}

func (m *AccountWithdrawalPostRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_INSERT
}

type AccountWithdrawalPostResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,3,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountWithdrawalPostResponse) Reset()         { *m = AccountWithdrawalPostResponse{} }
func (m *AccountWithdrawalPostResponse) String() string { return proto.CompactTextString(m) }
func (*AccountWithdrawalPostResponse) ProtoMessage()    {}
func (*AccountWithdrawalPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d5f5974e1c89e1a, []int{14}
}

func (m *AccountWithdrawalPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountWithdrawalPostResponse.Unmarshal(m, b)
}
func (m *AccountWithdrawalPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountWithdrawalPostResponse.Marshal(b, m, deterministic)
}
func (m *AccountWithdrawalPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountWithdrawalPostResponse.Merge(m, src)
}
func (m *AccountWithdrawalPostResponse) XXX_Size() int {
	return xxx_messageInfo_AccountWithdrawalPostResponse.Size(m)
}
func (m *AccountWithdrawalPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountWithdrawalPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountWithdrawalPostResponse proto.InternalMessageInfo

func (m *AccountWithdrawalPostResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_DONE
}

func (m *AccountWithdrawalPostResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *AccountWithdrawalPostResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

func init() {
	proto.RegisterType((*Balance)(nil), "api.Balance")
	proto.RegisterType((*BalanceRequest)(nil), "api.BalanceRequest")
	proto.RegisterType((*BalanceResponse)(nil), "api.BalanceResponse")
	proto.RegisterType((*AccountFunds)(nil), "api.AccountFunds")
	proto.RegisterType((*Fund)(nil), "api.Fund")
	proto.RegisterType((*WithdrawalFund)(nil), "api.WithdrawalFund")
	proto.RegisterType((*Transfer)(nil), "api.Transfer")
	proto.RegisterType((*AccountFundsRequest)(nil), "api.AccountFundsRequest")
	proto.RegisterType((*AccountFundsResponse)(nil), "api.AccountFundsResponse")
	proto.RegisterType((*AccountFundPostRequest)(nil), "api.AccountFundPostRequest")
	proto.RegisterType((*AccountFundPostResponse)(nil), "api.AccountFundPostResponse")
	proto.RegisterType((*AccountTransferPostRequest)(nil), "api.AccountTransferPostRequest")
	proto.RegisterType((*AccountTransferPostResponse)(nil), "api.AccountTransferPostResponse")
	proto.RegisterType((*AccountWithdrawalPostRequest)(nil), "api.AccountWithdrawalPostRequest")
	proto.RegisterType((*AccountWithdrawalPostResponse)(nil), "api.AccountWithdrawalPostResponse")
}

func init() { proto.RegisterFile("balances.proto", fileDescriptor_9d5f5974e1c89e1a) }

var fileDescriptor_9d5f5974e1c89e1a = []byte{
	// 764 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0xc6, 0xf9, 0xf3, 0x24, 0x4d, 0xe9, 0x26, 0x6a, 0x4d, 0xdb, 0x88, 0x60, 0xa4, 0x2a,
	0x08, 0xa9, 0x40, 0x41, 0x9c, 0xb8, 0x14, 0x55, 0x48, 0x39, 0x20, 0xa1, 0x15, 0x12, 0xc7, 0x68,
	0x6b, 0x6f, 0x55, 0x0b, 0xc7, 0x0e, 0xfe, 0x69, 0x54, 0x71, 0xe3, 0xc6, 0x95, 0x4b, 0x5f, 0x80,
	0x13, 0x07, 0x0e, 0xbc, 0x11, 0x6f, 0x82, 0x76, 0xbd, 0x6b, 0xaf, 0xd3, 0xe6, 0xa7, 0x48, 0xe5,
	0x96, 0x99, 0x6f, 0x3c, 0xf3, 0xcd, 0x37, 0xb3, 0xa3, 0x40, 0xe7, 0x94, 0xfa, 0x34, 0x70, 0x58,
	0x7c, 0x38, 0x8d, 0xc2, 0x24, 0xc4, 0x06, 0x9d, 0x7a, 0xbb, 0xc0, 0x82, 0x74, 0x92, 0x39, 0xec,
	0x3f, 0x08, 0x1a, 0x6f, 0xb2, 0x18, 0xdc, 0x07, 0x90, 0xe1, 0x63, 0xcf, 0xb5, 0xd0, 0x00, 0x0d,
	0x0d, 0x62, 0x4a, 0xcf, 0xc8, 0xc5, 0x0f, 0xa0, 0xe5, 0xa4, 0x51, 0xc4, 0x02, 0xe7, 0x92, 0xe3,
	0x15, 0x81, 0x83, 0x72, 0x8d, 0x5c, 0xfc, 0x1c, 0x7a, 0x79, 0x80, 0xcb, 0x62, 0x27, 0xf2, 0xa6,
	0x89, 0x17, 0x06, 0x96, 0x31, 0x40, 0x43, 0x93, 0x74, 0x15, 0x76, 0x52, 0x40, 0x78, 0x07, 0x1a,
	0x69, 0xcc, 0x22, 0x9e, 0xaf, 0x2a, 0xf2, 0xd5, 0xb9, 0x39, 0x72, 0xb1, 0x05, 0x0d, 0x59, 0xd9,
	0xaa, 0x0d, 0xd0, 0x10, 0x11, 0x65, 0xe2, 0x7d, 0x30, 0xe9, 0x05, 0xf5, 0x7c, 0x7a, 0xea, 0x33,
	0xab, 0x2e, 0xb0, 0xc2, 0x81, 0x31, 0x54, 0xcf, 0x43, 0xdf, 0xb5, 0x1a, 0x02, 0x10, 0xbf, 0xed,
	0x6f, 0x08, 0x3a, 0xb2, 0x47, 0xc2, 0x3e, 0xa7, 0x2c, 0x4e, 0x70, 0x0f, 0x6a, 0x49, 0xf8, 0x89,
	0x05, 0xa2, 0x4b, 0x93, 0x64, 0x06, 0x7e, 0x08, 0x6d, 0x25, 0xc0, 0x24, 0x74, 0x99, 0x68, 0xd1,
	0x24, 0x2d, 0xe9, 0x7b, 0x17, 0xba, 0xf3, 0x1a, 0x19, 0x2b, 0x34, 0xaa, 0xce, 0x6b, 0x64, 0x5f,
	0x21, 0xd8, 0xcc, 0xb9, 0xc4, 0xd3, 0x30, 0x88, 0x19, 0x3e, 0x80, 0x46, 0xc4, 0x12, 0x87, 0x57,
	0xe4, 0x74, 0x3a, 0x47, 0xed, 0x43, 0x3a, 0xf5, 0x0e, 0x49, 0xe6, 0x23, 0x0a, 0xc4, 0x43, 0x68,
	0xaa, 0x71, 0x5a, 0x95, 0x81, 0x31, 0x6c, 0xc9, 0x40, 0x95, 0x2f, 0x47, 0xb9, 0x7a, 0x4e, 0x38,
	0x99, 0xb0, 0x20, 0x91, 0xe2, 0x2b, 0x93, 0x23, 0xcc, 0xa7, 0xd3, 0x98, 0x65, 0xe4, 0x4c, 0xa2,
	0x4c, 0xfb, 0x87, 0x01, 0xed, 0x63, 0xc7, 0x09, 0xd3, 0x20, 0x79, 0x9b, 0x06, 0x6e, 0x8c, 0x0f,
	0x60, 0x93, 0x66, 0xf6, 0xf8, 0x2c, 0x0d, 0xdc, 0x62, 0x27, 0x36, 0x68, 0x11, 0x76, 0x47, 0x7b,
	0xd1, 0x07, 0x50, 0xb5, 0x73, 0x19, 0x4d, 0xe9, 0x19, 0xb9, 0xf8, 0x29, 0x74, 0x15, 0xac, 0x27,
	0xac, 0x89, 0x84, 0x58, 0x42, 0x7a, 0xbe, 0x3d, 0x30, 0x45, 0x0f, 0x62, 0xac, 0x75, 0x11, 0xd6,
	0xe4, 0x0e, 0x31, 0x53, 0x6d, 0x09, 0x1b, 0xa5, 0x25, 0xbc, 0x0f, 0xcd, 0x0b, 0x16, 0xa4, 0x62,
	0xd4, 0x4d, 0x81, 0x34, 0x84, 0x3d, 0x72, 0xf1, 0x13, 0xd8, 0xca, 0x20, 0xbd, 0xbe, 0x29, 0x12,
	0xdf, 0x13, 0x80, 0x5e, 0x5d, 0x5b, 0x66, 0x58, 0xb2, 0xcc, 0xad, 0x45, 0xcb, 0xdc, 0xd6, 0x96,
	0xf9, 0x12, 0xaa, 0x5c, 0xf7, 0x39, 0x85, 0xd0, 0xbc, 0x42, 0x3a, 0xf5, 0x4a, 0x99, 0xfa, 0xdc,
	0xbc, 0x8c, 0x6b, 0xf3, 0xda, 0x86, 0x3a, 0x9d, 0xf0, 0x3c, 0x42, 0x78, 0x44, 0xa4, 0x65, 0x7f,
	0x47, 0xd0, 0xf9, 0xe8, 0x25, 0xe7, 0x6e, 0x44, 0x67, 0xd4, 0x17, 0x2c, 0xd6, 0xdd, 0x91, 0x32,
	0xdb, 0xca, 0x3c, 0xdb, 0xa2, 0xa2, 0xa1, 0x57, 0xe4, 0x0f, 0x72, 0x26, 0x0b, 0x8e, 0xa9, 0xef,
	0x0b, 0x3e, 0x4d, 0xd2, 0x52, 0xbe, 0x63, 0xdf, 0xb7, 0x7f, 0x23, 0x68, 0x7e, 0x88, 0x68, 0x10,
	0x9f, 0xb1, 0x68, 0x6d, 0x3a, 0x2f, 0x61, 0xbb, 0xa0, 0xc3, 0x47, 0x98, 0x78, 0x01, 0x15, 0x23,
	0xcc, 0xa8, 0xf5, 0x72, 0x6a, 0x27, 0x05, 0x86, 0x9f, 0x41, 0x4f, 0x69, 0x5a, 0xfa, 0x26, 0x53,
	0x10, 0x4b, 0x7d, 0xf5, 0x2f, 0x16, 0x29, 0xf9, 0x0b, 0x41, 0x57, 0x7f, 0x6b, 0xcb, 0xcf, 0xd2,
	0x0d, 0x5d, 0x55, 0x56, 0x8b, 0x6c, 0x2c, 0x5b, 0x89, 0xea, 0xd2, 0x95, 0xa8, 0x5d, 0x3b, 0x5b,
	0x3f, 0x11, 0xf4, 0xca, 0x84, 0x6f, 0x79, 0xbb, 0x5e, 0xc1, 0x86, 0xde, 0x83, 0x3a, 0x60, 0x5b,
	0x22, 0xba, 0x94, 0xb9, 0x4d, 0xf5, 0x23, 0xf4, 0x2f, 0x97, 0x2c, 0x82, 0x6d, 0x2d, 0xe3, 0xfb,
	0x30, 0x4e, 0x96, 0xeb, 0xdb, 0x87, 0x2a, 0xe7, 0x24, 0x44, 0x6d, 0x1d, 0x99, 0x82, 0x12, 0xff,
	0x92, 0x08, 0x37, 0x7e, 0x04, 0x75, 0xea, 0xe4, 0x83, 0xee, 0x1c, 0xb5, 0x24, 0x67, 0xee, 0x22,
	0x12, 0xb2, 0x53, 0xd8, 0xb9, 0x56, 0xf3, 0x96, 0x12, 0x69, 0xad, 0x56, 0x16, 0xb6, 0x6a, 0x94,
	0x5b, 0xfd, 0x8a, 0x60, 0x57, 0xd6, 0x55, 0x8f, 0x60, 0x75, 0xbf, 0x8f, 0xa1, 0x99, 0xc8, 0x60,
	0xd9, 0xf3, 0x86, 0x60, 0xa4, 0x32, 0x90, 0x1c, 0x5e, 0xaf, 0xf7, 0x4b, 0xd8, 0xbb, 0x91, 0xc3,
	0x7f, 0xe8, 0xff, 0x0a, 0xc1, 0xbe, 0xac, 0x5d, 0x5c, 0xa6, 0xd5, 0x0a, 0xbc, 0x86, 0xcd, 0x59,
	0x1e, 0x3e, 0xd6, 0x86, 0xdf, 0x15, 0xd4, 0xca, 0x47, 0x8e, 0x74, 0x66, 0xe5, 0xa3, 0xb7, 0x96,
	0x28, 0x5f, 0xa0, 0xbf, 0x80, 0xd8, 0xdd, 0xcb, 0x72, 0x5a, 0x17, 0x7f, 0xee, 0x5e, 0xfc, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xda, 0x0a, 0xf6, 0x17, 0xff, 0x09, 0x00, 0x00,
}
