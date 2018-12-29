// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accounts.proto

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

type Account struct {
	Id                   string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId               string      `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AccountMode          AccountMode `protobuf:"varint,3,opt,name=account_mode,json=accountMode,proto3,enum=api.AccountMode" json:"account_mode,omitempty"`
	AccountType          AccountType `protobuf:"varint,4,opt,name=account_type,json=accountType,proto3,enum=api.AccountType" json:"account_type,omitempty"`
	Enabled              bool        `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *Account) GetAccountMode() AccountMode {
	if m != nil {
		return m.AccountMode
	}
	return AccountMode_DEMO
}

func (m *Account) GetAccountType() AccountType {
	if m != nil {
		return m.AccountType
	}
	return AccountType_NET
}

func (m *Account) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type AccountsRequest struct {
	Token                string      `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	AccountId            string      `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	AccountMode          AccountMode `protobuf:"varint,3,opt,name=account_mode,json=accountMode,proto3,enum=api.AccountMode" json:"account_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AccountsRequest) Reset()         { *m = AccountsRequest{} }
func (m *AccountsRequest) String() string { return proto.CompactTextString(m) }
func (*AccountsRequest) ProtoMessage()    {}
func (*AccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{1}
}

func (m *AccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountsRequest.Unmarshal(m, b)
}
func (m *AccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountsRequest.Marshal(b, m, deterministic)
}
func (m *AccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountsRequest.Merge(m, src)
}
func (m *AccountsRequest) XXX_Size() int {
	return xxx_messageInfo_AccountsRequest.Size(m)
}
func (m *AccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountsRequest proto.InternalMessageInfo

func (m *AccountsRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AccountsRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *AccountsRequest) GetAccountMode() AccountMode {
	if m != nil {
		return m.AccountMode
	}
	return AccountMode_DEMO
}

type AccountsResponse struct {
	Retcode              Retcode    `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Accounts             []*Account `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	Comment              string     `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string     `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AccountsResponse) Reset()         { *m = AccountsResponse{} }
func (m *AccountsResponse) String() string { return proto.CompactTextString(m) }
func (*AccountsResponse) ProtoMessage()    {}
func (*AccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{2}
}

func (m *AccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountsResponse.Unmarshal(m, b)
}
func (m *AccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountsResponse.Marshal(b, m, deterministic)
}
func (m *AccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountsResponse.Merge(m, src)
}
func (m *AccountsResponse) XXX_Size() int {
	return xxx_messageInfo_AccountsResponse.Size(m)
}
func (m *AccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountsResponse proto.InternalMessageInfo

func (m *AccountsResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *AccountsResponse) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *AccountsResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *AccountsResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type AccountsPostRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Action               Action   `protobuf:"varint,2,opt,name=action,proto3,enum=api.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountsPostRequest) Reset()         { *m = AccountsPostRequest{} }
func (m *AccountsPostRequest) String() string { return proto.CompactTextString(m) }
func (*AccountsPostRequest) ProtoMessage()    {}
func (*AccountsPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{3}
}

func (m *AccountsPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountsPostRequest.Unmarshal(m, b)
}
func (m *AccountsPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountsPostRequest.Marshal(b, m, deterministic)
}
func (m *AccountsPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountsPostRequest.Merge(m, src)
}
func (m *AccountsPostRequest) XXX_Size() int {
	return xxx_messageInfo_AccountsPostRequest.Size(m)
}
func (m *AccountsPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountsPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AccountsPostRequest proto.InternalMessageInfo

func (m *AccountsPostRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AccountsPostRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_INSERT
}

type AccountsPostResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=api.Retcode" json:"retcode,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,3,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountsPostResponse) Reset()         { *m = AccountsPostResponse{} }
func (m *AccountsPostResponse) String() string { return proto.CompactTextString(m) }
func (*AccountsPostResponse) ProtoMessage()    {}
func (*AccountsPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1e7723af4c007b7, []int{4}
}

func (m *AccountsPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountsPostResponse.Unmarshal(m, b)
}
func (m *AccountsPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountsPostResponse.Marshal(b, m, deterministic)
}
func (m *AccountsPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountsPostResponse.Merge(m, src)
}
func (m *AccountsPostResponse) XXX_Size() int {
	return xxx_messageInfo_AccountsPostResponse.Size(m)
}
func (m *AccountsPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountsPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AccountsPostResponse proto.InternalMessageInfo

func (m *AccountsPostResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_REJECTX
}

func (m *AccountsPostResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *AccountsPostResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

func init() {
	proto.RegisterType((*Account)(nil), "api.Account")
	proto.RegisterType((*AccountsRequest)(nil), "api.AccountsRequest")
	proto.RegisterType((*AccountsResponse)(nil), "api.AccountsResponse")
	proto.RegisterType((*AccountsPostRequest)(nil), "api.AccountsPostRequest")
	proto.RegisterType((*AccountsPostResponse)(nil), "api.AccountsPostResponse")
}

func init() { proto.RegisterFile("accounts.proto", fileDescriptor_e1e7723af4c007b7) }

var fileDescriptor_e1e7723af4c007b7 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x6a, 0xf3, 0x30,
	0x10, 0xc5, 0xb1, 0xfd, 0x25, 0x8e, 0x27, 0xc1, 0x5f, 0x50, 0x03, 0x35, 0x85, 0x82, 0x71, 0xa1,
	0x78, 0x95, 0x45, 0x72, 0x82, 0x2e, 0xbb, 0x28, 0x04, 0xd1, 0x7d, 0x50, 0xac, 0x59, 0x98, 0xc6,
	0x92, 0x6a, 0xc9, 0x8b, 0xd0, 0x8b, 0xf4, 0x24, 0x3d, 0x5f, 0xb1, 0x2c, 0xe5, 0x0f, 0xa5, 0x85,
	0x76, 0xf9, 0xde, 0x68, 0xde, 0xcc, 0x6f, 0x6c, 0x48, 0x59, 0x55, 0xc9, 0x4e, 0x18, 0xbd, 0x54,
	0xad, 0x34, 0x92, 0x44, 0x4c, 0xd5, 0x37, 0x80, 0xa2, 0x6b, 0x06, 0xa3, 0xf8, 0x08, 0x20, 0x7e,
	0x18, 0xde, 0x90, 0x14, 0xc2, 0x9a, 0x67, 0x41, 0x1e, 0x94, 0x09, 0x0d, 0x6b, 0x4e, 0xae, 0x21,
	0xee, 0x34, 0xb6, 0xdb, 0x9a, 0x67, 0xa1, 0x35, 0xc7, 0xbd, 0x7c, 0xe4, 0x64, 0x0d, 0x33, 0x97,
	0xbb, 0x6d, 0x24, 0xc7, 0x2c, 0xca, 0x83, 0x32, 0x5d, 0xcd, 0x97, 0x4c, 0xd5, 0x4b, 0x17, 0xf6,
	0x24, 0x39, 0xd2, 0x29, 0x3b, 0x89, 0xf3, 0x26, 0x73, 0x50, 0x98, 0xfd, 0xfb, 0xda, 0xf4, 0x7c,
	0x50, 0xa7, 0xa6, 0x5e, 0x90, 0x0c, 0x62, 0x14, 0x6c, 0xb7, 0x47, 0x9e, 0x8d, 0xf2, 0xa0, 0x9c,
	0x50, 0x2f, 0x8b, 0x37, 0xf8, 0xef, 0xba, 0x34, 0xc5, 0xd7, 0x0e, 0xb5, 0x21, 0x0b, 0x18, 0x19,
	0xf9, 0x82, 0xc2, 0x21, 0x0c, 0x82, 0xdc, 0x02, 0xf8, 0xb9, 0x47, 0x90, 0xc4, 0x39, 0x7f, 0x64,
	0x29, 0xde, 0x03, 0x98, 0x9f, 0xa6, 0x6b, 0x25, 0x85, 0x46, 0x72, 0x0f, 0x71, 0x8b, 0xa6, 0xea,
	0x43, 0x02, 0x1b, 0x32, 0xb3, 0x21, 0x74, 0xf0, 0xa8, 0x2f, 0x92, 0x12, 0x26, 0xfe, 0xab, 0x64,
	0x61, 0x1e, 0x95, 0x53, 0xf7, 0xd0, 0x05, 0xd2, 0x63, 0xb5, 0xa7, 0xaf, 0x64, 0xd3, 0xa0, 0x30,
	0x76, 0xad, 0x84, 0x7a, 0x69, 0xef, 0xb2, 0x67, 0x4a, 0x23, 0xb7, 0x77, 0x4c, 0xa8, 0x97, 0xc5,
	0x06, 0xae, 0xfc, 0x66, 0x1b, 0xa9, 0xcd, 0xcf, 0xb7, 0xb9, 0x83, 0x31, 0xab, 0x4c, 0x2d, 0x85,
	0xbd, 0x4b, 0xba, 0x9a, 0xba, 0x45, 0x7a, 0x8b, 0xba, 0x52, 0xd1, 0xc2, 0xe2, 0x32, 0xf1, 0x97,
	0xbc, 0x67, 0x14, 0xe1, 0xb7, 0x14, 0xd1, 0x05, 0xc5, 0x6e, 0x6c, 0xff, 0xce, 0xf5, 0x67, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc5, 0xc6, 0x67, 0x29, 0xc0, 0x02, 0x00, 0x00,
}
