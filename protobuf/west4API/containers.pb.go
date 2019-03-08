// Code generated by protoc-gen-go. DO NOT EDIT.
// source: containers.proto

package west4API

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

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Names                string   `protobuf:"bytes,2,opt,name=names,proto3" json:"names,omitempty"`
	Image                string   `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	ImageId              string   `protobuf:"bytes,4,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	Created              string   `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Ports                string   `protobuf:"bytes,6,opt,name=ports,proto3" json:"ports,omitempty"`
	SizeRw               int64    `protobuf:"varint,7,opt,name=size_rw,json=sizeRw,proto3" json:"size_rw,omitempty"`
	SizeRootFs           int64    `protobuf:"varint,8,opt,name=size_root_fs,json=sizeRootFs,proto3" json:"size_root_fs,omitempty"`
	Labels               string   `protobuf:"bytes,9,opt,name=labels,proto3" json:"labels,omitempty"`
	State                string   `protobuf:"bytes,10,opt,name=state,proto3" json:"state,omitempty"`
	Status               string   `protobuf:"bytes,11,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_d029d11d13649966, []int{0}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetNames() string {
	if m != nil {
		return m.Names
	}
	return ""
}

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *Container) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Container) GetPorts() string {
	if m != nil {
		return m.Ports
	}
	return ""
}

func (m *Container) GetSizeRw() int64 {
	if m != nil {
		return m.SizeRw
	}
	return 0
}

func (m *Container) GetSizeRootFs() int64 {
	if m != nil {
		return m.SizeRootFs
	}
	return 0
}

func (m *Container) GetLabels() string {
	if m != nil {
		return m.Labels
	}
	return ""
}

func (m *Container) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Container) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ContainerCreate struct {
	ContainerId          int64    `protobuf:"varint,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ContainerIdentifier  string   `protobuf:"bytes,2,opt,name=container_identifier,json=containerIdentifier,proto3" json:"container_identifier,omitempty"`
	ContainerType        string   `protobuf:"bytes,3,opt,name=container_type,json=containerType,proto3" json:"container_type,omitempty"`
	ContainerDescription string   `protobuf:"bytes,4,opt,name=container_description,json=containerDescription,proto3" json:"container_description,omitempty"`
	Image                string   `protobuf:"bytes,5,opt,name=image,proto3" json:"image,omitempty"`
	Network              string   `protobuf:"bytes,6,opt,name=network,proto3" json:"network,omitempty"`
	Hostname             string   `protobuf:"bytes,7,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Env                  string   `protobuf:"bytes,8,opt,name=env,proto3" json:"env,omitempty"`
	RestartPolicy        string   `protobuf:"bytes,9,opt,name=restart_policy,json=restartPolicy,proto3" json:"restart_policy,omitempty"`
	MemoryReservation    int64    `protobuf:"varint,10,opt,name=memory_reservation,json=memoryReservation,proto3" json:"memory_reservation,omitempty"`
	MemoryLimit          int64    `protobuf:"varint,11,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	CpuLimit             int64    `protobuf:"varint,12,opt,name=cpu_limit,json=cpuLimit,proto3" json:"cpu_limit,omitempty"`
	StopTimeout          int64    `protobuf:"varint,13,opt,name=stop_timeout,json=stopTimeout,proto3" json:"stop_timeout,omitempty"`
	ServerId             int64    `protobuf:"varint,14,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	ContainerConfig      string   `protobuf:"bytes,15,opt,name=container_config,json=containerConfig,proto3" json:"container_config,omitempty"`
	LastUpdate           string   `protobuf:"bytes,16,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerCreate) Reset()         { *m = ContainerCreate{} }
func (m *ContainerCreate) String() string { return proto.CompactTextString(m) }
func (*ContainerCreate) ProtoMessage()    {}
func (*ContainerCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_d029d11d13649966, []int{1}
}

func (m *ContainerCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerCreate.Unmarshal(m, b)
}
func (m *ContainerCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerCreate.Marshal(b, m, deterministic)
}
func (m *ContainerCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerCreate.Merge(m, src)
}
func (m *ContainerCreate) XXX_Size() int {
	return xxx_messageInfo_ContainerCreate.Size(m)
}
func (m *ContainerCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerCreate.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerCreate proto.InternalMessageInfo

func (m *ContainerCreate) GetContainerId() int64 {
	if m != nil {
		return m.ContainerId
	}
	return 0
}

func (m *ContainerCreate) GetContainerIdentifier() string {
	if m != nil {
		return m.ContainerIdentifier
	}
	return ""
}

func (m *ContainerCreate) GetContainerType() string {
	if m != nil {
		return m.ContainerType
	}
	return ""
}

func (m *ContainerCreate) GetContainerDescription() string {
	if m != nil {
		return m.ContainerDescription
	}
	return ""
}

func (m *ContainerCreate) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ContainerCreate) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *ContainerCreate) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *ContainerCreate) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *ContainerCreate) GetRestartPolicy() string {
	if m != nil {
		return m.RestartPolicy
	}
	return ""
}

func (m *ContainerCreate) GetMemoryReservation() int64 {
	if m != nil {
		return m.MemoryReservation
	}
	return 0
}

func (m *ContainerCreate) GetMemoryLimit() int64 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *ContainerCreate) GetCpuLimit() int64 {
	if m != nil {
		return m.CpuLimit
	}
	return 0
}

func (m *ContainerCreate) GetStopTimeout() int64 {
	if m != nil {
		return m.StopTimeout
	}
	return 0
}

func (m *ContainerCreate) GetServerId() int64 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *ContainerCreate) GetContainerConfig() string {
	if m != nil {
		return m.ContainerConfig
	}
	return ""
}

func (m *ContainerCreate) GetLastUpdate() string {
	if m != nil {
		return m.LastUpdate
	}
	return ""
}

// -------
type ContainerRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ContainerId          int64    `protobuf:"varint,2,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	ServerId             int64    `protobuf:"varint,3,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerRequest) Reset()         { *m = ContainerRequest{} }
func (m *ContainerRequest) String() string { return proto.CompactTextString(m) }
func (*ContainerRequest) ProtoMessage()    {}
func (*ContainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d029d11d13649966, []int{2}
}

func (m *ContainerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerRequest.Unmarshal(m, b)
}
func (m *ContainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerRequest.Marshal(b, m, deterministic)
}
func (m *ContainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerRequest.Merge(m, src)
}
func (m *ContainerRequest) XXX_Size() int {
	return xxx_messageInfo_ContainerRequest.Size(m)
}
func (m *ContainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerRequest proto.InternalMessageInfo

func (m *ContainerRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ContainerRequest) GetContainerId() int64 {
	if m != nil {
		return m.ContainerId
	}
	return 0
}

func (m *ContainerRequest) GetServerId() int64 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

type ContainerResponse struct {
	Retcode              Retcode      `protobuf:"varint,1,opt,name=retcode,proto3,enum=west4API.Retcode" json:"retcode,omitempty"`
	Container            []*Container `protobuf:"bytes,2,rep,name=container,proto3" json:"container,omitempty"`
	Comment              string       `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string       `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContainerResponse) Reset()         { *m = ContainerResponse{} }
func (m *ContainerResponse) String() string { return proto.CompactTextString(m) }
func (*ContainerResponse) ProtoMessage()    {}
func (*ContainerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d029d11d13649966, []int{3}
}

func (m *ContainerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerResponse.Unmarshal(m, b)
}
func (m *ContainerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerResponse.Marshal(b, m, deterministic)
}
func (m *ContainerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerResponse.Merge(m, src)
}
func (m *ContainerResponse) XXX_Size() int {
	return xxx_messageInfo_ContainerResponse.Size(m)
}
func (m *ContainerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerResponse proto.InternalMessageInfo

func (m *ContainerResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_DONE
}

func (m *ContainerResponse) GetContainer() []*Container {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *ContainerResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *ContainerResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type ContainerPostRequest struct {
	Token                string       `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ServerId             int64        `protobuf:"varint,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Container            []*Container `protobuf:"bytes,3,rep,name=container,proto3" json:"container,omitempty"`
	Action               Action       `protobuf:"varint,4,opt,name=action,proto3,enum=west4API.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContainerPostRequest) Reset()         { *m = ContainerPostRequest{} }
func (m *ContainerPostRequest) String() string { return proto.CompactTextString(m) }
func (*ContainerPostRequest) ProtoMessage()    {}
func (*ContainerPostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d029d11d13649966, []int{4}
}

func (m *ContainerPostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerPostRequest.Unmarshal(m, b)
}
func (m *ContainerPostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerPostRequest.Marshal(b, m, deterministic)
}
func (m *ContainerPostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerPostRequest.Merge(m, src)
}
func (m *ContainerPostRequest) XXX_Size() int {
	return xxx_messageInfo_ContainerPostRequest.Size(m)
}
func (m *ContainerPostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerPostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerPostRequest proto.InternalMessageInfo

func (m *ContainerPostRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ContainerPostRequest) GetServerId() int64 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *ContainerPostRequest) GetContainer() []*Container {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *ContainerPostRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_INSERT
}

type ContainerPostResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=west4API.Retcode" json:"retcode,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,3,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerPostResponse) Reset()         { *m = ContainerPostResponse{} }
func (m *ContainerPostResponse) String() string { return proto.CompactTextString(m) }
func (*ContainerPostResponse) ProtoMessage()    {}
func (*ContainerPostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d029d11d13649966, []int{5}
}

func (m *ContainerPostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerPostResponse.Unmarshal(m, b)
}
func (m *ContainerPostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerPostResponse.Marshal(b, m, deterministic)
}
func (m *ContainerPostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerPostResponse.Merge(m, src)
}
func (m *ContainerPostResponse) XXX_Size() int {
	return xxx_messageInfo_ContainerPostResponse.Size(m)
}
func (m *ContainerPostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerPostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerPostResponse proto.InternalMessageInfo

func (m *ContainerPostResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_DONE
}

func (m *ContainerPostResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *ContainerPostResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

// -------
type ContainerCreateRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ContainerCreate      int64    `protobuf:"varint,2,opt,name=container_create,json=containerCreate,proto3" json:"container_create,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerCreateRequest) Reset()         { *m = ContainerCreateRequest{} }
func (m *ContainerCreateRequest) String() string { return proto.CompactTextString(m) }
func (*ContainerCreateRequest) ProtoMessage()    {}
func (*ContainerCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d029d11d13649966, []int{6}
}

func (m *ContainerCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerCreateRequest.Unmarshal(m, b)
}
func (m *ContainerCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerCreateRequest.Marshal(b, m, deterministic)
}
func (m *ContainerCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerCreateRequest.Merge(m, src)
}
func (m *ContainerCreateRequest) XXX_Size() int {
	return xxx_messageInfo_ContainerCreateRequest.Size(m)
}
func (m *ContainerCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerCreateRequest proto.InternalMessageInfo

func (m *ContainerCreateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ContainerCreateRequest) GetContainerCreate() int64 {
	if m != nil {
		return m.ContainerCreate
	}
	return 0
}

type ContainerCreateResponse struct {
	Retcode              Retcode            `protobuf:"varint,1,opt,name=retcode,proto3,enum=west4API.Retcode" json:"retcode,omitempty"`
	ContainerCreate      []*ContainerCreate `protobuf:"bytes,2,rep,name=container_create,json=containerCreate,proto3" json:"container_create,omitempty"`
	Comment              string             `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string             `protobuf:"bytes,4,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ContainerCreateResponse) Reset()         { *m = ContainerCreateResponse{} }
func (m *ContainerCreateResponse) String() string { return proto.CompactTextString(m) }
func (*ContainerCreateResponse) ProtoMessage()    {}
func (*ContainerCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d029d11d13649966, []int{7}
}

func (m *ContainerCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerCreateResponse.Unmarshal(m, b)
}
func (m *ContainerCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerCreateResponse.Marshal(b, m, deterministic)
}
func (m *ContainerCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerCreateResponse.Merge(m, src)
}
func (m *ContainerCreateResponse) XXX_Size() int {
	return xxx_messageInfo_ContainerCreateResponse.Size(m)
}
func (m *ContainerCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerCreateResponse proto.InternalMessageInfo

func (m *ContainerCreateResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_DONE
}

func (m *ContainerCreateResponse) GetContainerCreate() []*ContainerCreate {
	if m != nil {
		return m.ContainerCreate
	}
	return nil
}

func (m *ContainerCreateResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *ContainerCreateResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

type ContainerCreatePostRequest struct {
	Token                string           `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ContainerCreate      *ContainerCreate `protobuf:"bytes,2,opt,name=container_create,json=containerCreate,proto3" json:"container_create,omitempty"`
	Action               Action           `protobuf:"varint,3,opt,name=action,proto3,enum=west4API.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ContainerCreatePostRequest) Reset()         { *m = ContainerCreatePostRequest{} }
func (m *ContainerCreatePostRequest) String() string { return proto.CompactTextString(m) }
func (*ContainerCreatePostRequest) ProtoMessage()    {}
func (*ContainerCreatePostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d029d11d13649966, []int{8}
}

func (m *ContainerCreatePostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerCreatePostRequest.Unmarshal(m, b)
}
func (m *ContainerCreatePostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerCreatePostRequest.Marshal(b, m, deterministic)
}
func (m *ContainerCreatePostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerCreatePostRequest.Merge(m, src)
}
func (m *ContainerCreatePostRequest) XXX_Size() int {
	return xxx_messageInfo_ContainerCreatePostRequest.Size(m)
}
func (m *ContainerCreatePostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerCreatePostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerCreatePostRequest proto.InternalMessageInfo

func (m *ContainerCreatePostRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ContainerCreatePostRequest) GetContainerCreate() *ContainerCreate {
	if m != nil {
		return m.ContainerCreate
	}
	return nil
}

func (m *ContainerCreatePostRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_INSERT
}

type ContainerCreatePostResponse struct {
	Retcode              Retcode  `protobuf:"varint,1,opt,name=retcode,proto3,enum=west4API.Retcode" json:"retcode,omitempty"`
	Comment              string   `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Elapsed              string   `protobuf:"bytes,3,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerCreatePostResponse) Reset()         { *m = ContainerCreatePostResponse{} }
func (m *ContainerCreatePostResponse) String() string { return proto.CompactTextString(m) }
func (*ContainerCreatePostResponse) ProtoMessage()    {}
func (*ContainerCreatePostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d029d11d13649966, []int{9}
}

func (m *ContainerCreatePostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerCreatePostResponse.Unmarshal(m, b)
}
func (m *ContainerCreatePostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerCreatePostResponse.Marshal(b, m, deterministic)
}
func (m *ContainerCreatePostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerCreatePostResponse.Merge(m, src)
}
func (m *ContainerCreatePostResponse) XXX_Size() int {
	return xxx_messageInfo_ContainerCreatePostResponse.Size(m)
}
func (m *ContainerCreatePostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerCreatePostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerCreatePostResponse proto.InternalMessageInfo

func (m *ContainerCreatePostResponse) GetRetcode() Retcode {
	if m != nil {
		return m.Retcode
	}
	return Retcode_DONE
}

func (m *ContainerCreatePostResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *ContainerCreatePostResponse) GetElapsed() string {
	if m != nil {
		return m.Elapsed
	}
	return ""
}

func init() {
	proto.RegisterType((*Container)(nil), "west4API.Container")
	proto.RegisterType((*ContainerCreate)(nil), "west4API.ContainerCreate")
	proto.RegisterType((*ContainerRequest)(nil), "west4API.ContainerRequest")
	proto.RegisterType((*ContainerResponse)(nil), "west4API.ContainerResponse")
	proto.RegisterType((*ContainerPostRequest)(nil), "west4API.ContainerPostRequest")
	proto.RegisterType((*ContainerPostResponse)(nil), "west4API.ContainerPostResponse")
	proto.RegisterType((*ContainerCreateRequest)(nil), "west4API.ContainerCreateRequest")
	proto.RegisterType((*ContainerCreateResponse)(nil), "west4API.ContainerCreateResponse")
	proto.RegisterType((*ContainerCreatePostRequest)(nil), "west4API.ContainerCreatePostRequest")
	proto.RegisterType((*ContainerCreatePostResponse)(nil), "west4API.ContainerCreatePostResponse")
}

func init() { proto.RegisterFile("containers.proto", fileDescriptor_d029d11d13649966) }

var fileDescriptor_d029d11d13649966 = []byte{
	// 741 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0xcb, 0x4e, 0xdb, 0x4c,
	0x14, 0x56, 0x6c, 0xc8, 0xe5, 0x24, 0x84, 0x30, 0xdc, 0x06, 0x58, 0xfc, 0x21, 0xd2, 0x2f, 0xe5,
	0xd7, 0xaf, 0x22, 0x01, 0x7d, 0x01, 0x04, 0xaa, 0x14, 0xa9, 0x0b, 0x64, 0xd1, 0x45, 0x57, 0x96,
	0xb1, 0x0f, 0x30, 0x22, 0xf6, 0xb8, 0x33, 0x13, 0xd2, 0x74, 0xd1, 0x97, 0xe8, 0x1b, 0x74, 0xd1,
	0xbe, 0x46, 0xa5, 0xbe, 0x58, 0x35, 0x97, 0xd8, 0x49, 0xb8, 0xb5, 0x2c, 0xba, 0xf3, 0xf7, 0x7d,
	0x67, 0xce, 0x39, 0xfe, 0xce, 0x19, 0x1b, 0x3a, 0x31, 0xcf, 0x54, 0xc4, 0x32, 0x14, 0xf2, 0x20,
	0x17, 0x5c, 0x71, 0x52, 0x1f, 0xa3, 0x54, 0xaf, 0x4f, 0xce, 0x07, 0xbb, 0x80, 0xd9, 0x28, 0xb5,
	0x6c, 0xef, 0x8b, 0x07, 0x8d, 0xd3, 0x69, 0x28, 0x69, 0x83, 0xc7, 0x12, 0x5a, 0xe9, 0x56, 0xfa,
	0x8d, 0xc0, 0x63, 0x09, 0xd9, 0x80, 0xe5, 0x2c, 0x4a, 0x51, 0x52, 0xcf, 0x50, 0x16, 0x68, 0x96,
	0xa5, 0xd1, 0x35, 0x52, 0xdf, 0xb2, 0x06, 0x90, 0x1d, 0xa8, 0x9b, 0x87, 0x90, 0x25, 0x74, 0xc9,
	0x08, 0x35, 0x83, 0x07, 0x09, 0xa1, 0x50, 0x8b, 0x05, 0x46, 0x0a, 0x13, 0xba, 0x6c, 0x15, 0x07,
	0x75, 0xaa, 0x9c, 0x0b, 0x25, 0x69, 0xd5, 0xa6, 0x32, 0x80, 0x6c, 0x43, 0x4d, 0xb2, 0x4f, 0x18,
	0x8a, 0x31, 0xad, 0x75, 0x2b, 0x7d, 0x3f, 0xa8, 0x6a, 0x18, 0x8c, 0x49, 0x17, 0x5a, 0x56, 0xe0,
	0x5c, 0x85, 0x57, 0x92, 0xd6, 0x8d, 0x0a, 0x46, 0xe5, 0x5c, 0xbd, 0x91, 0x64, 0x0b, 0xaa, 0xc3,
	0xe8, 0x12, 0x87, 0x92, 0x36, 0x4c, 0x46, 0x87, 0x74, 0x21, 0xa9, 0x22, 0x85, 0x14, 0x6c, 0x21,
	0x03, 0x74, 0xb4, 0x7e, 0x18, 0x49, 0xda, 0xb4, 0xd1, 0x16, 0xf5, 0x7e, 0x2e, 0xc1, 0x6a, 0xe1,
	0xca, 0xa9, 0xe9, 0x95, 0xec, 0x43, 0xab, 0xf0, 0x34, 0x74, 0x2e, 0xf9, 0x41, 0xb3, 0xe0, 0x06,
	0x09, 0x39, 0x84, 0x8d, 0xd9, 0x10, 0xcc, 0x14, 0xbb, 0x62, 0x28, 0x9c, 0x7b, 0xeb, 0x33, 0xa1,
	0x53, 0x89, 0xfc, 0x0b, 0xed, 0xf2, 0x88, 0x9a, 0xe4, 0x53, 0x53, 0x57, 0x0a, 0xf6, 0x62, 0x92,
	0x23, 0x39, 0x86, 0xcd, 0x32, 0x2c, 0x41, 0x19, 0x0b, 0x96, 0x2b, 0xc6, 0x33, 0xe7, 0x74, 0x59,
	0xf6, 0xac, 0xd4, 0xca, 0x39, 0x2d, 0xcf, 0xce, 0x89, 0x42, 0x2d, 0x43, 0x35, 0xe6, 0xe2, 0xd6,
	0x99, 0x3e, 0x85, 0x64, 0x17, 0xea, 0x37, 0x5c, 0x2a, 0x3d, 0x64, 0xe3, 0x7b, 0x23, 0x28, 0x30,
	0xe9, 0x80, 0x8f, 0xd9, 0x9d, 0x31, 0xbc, 0x11, 0xe8, 0x47, 0xdd, 0xb9, 0x40, 0xa9, 0x22, 0xa1,
	0xc2, 0x9c, 0x0f, 0x59, 0x3c, 0x71, 0x8e, 0xaf, 0x38, 0xf6, 0xdc, 0x90, 0xe4, 0x15, 0x90, 0x14,
	0x53, 0x2e, 0x26, 0xa1, 0x40, 0x89, 0xe2, 0x2e, 0x32, 0x6d, 0x83, 0x31, 0x6f, 0xcd, 0x2a, 0x41,
	0x29, 0x68, 0x97, 0x5d, 0xf8, 0x90, 0xa5, 0x4c, 0x99, 0xb9, 0xf8, 0x41, 0xd3, 0x72, 0x6f, 0x35,
	0x45, 0xf6, 0xa0, 0x11, 0xe7, 0x23, 0xa7, 0xb7, 0x8c, 0x5e, 0x8f, 0xf3, 0x91, 0x15, 0xf7, 0xa1,
	0x25, 0x15, 0xcf, 0x43, 0xc5, 0x52, 0xe4, 0x23, 0x45, 0x57, 0xec, 0x79, 0xcd, 0x5d, 0x58, 0x4a,
	0x9f, 0xd7, 0xf5, 0xec, 0x14, 0xdb, 0xf6, 0xbc, 0x25, 0x06, 0x09, 0xf9, 0x6f, 0xe6, 0xe6, 0x84,
	0x31, 0xcf, 0xae, 0xd8, 0x35, 0x5d, 0x35, 0xef, 0xb5, 0x5a, 0xf0, 0xa7, 0x86, 0x26, 0xff, 0x40,
	0x73, 0x18, 0x49, 0x15, 0x8e, 0xf2, 0x44, 0x2f, 0x56, 0xc7, 0x44, 0x81, 0xa6, 0xde, 0x19, 0xa6,
	0x77, 0x03, 0x9d, 0x62, 0x89, 0x02, 0xfc, 0x30, 0x42, 0xa9, 0xf4, 0x4c, 0x14, 0xbf, 0xc5, 0xcc,
	0x5d, 0x32, 0x0b, 0xee, 0xed, 0x96, 0x77, 0x7f, 0xb7, 0xe6, 0xba, 0xf6, 0xe7, 0xbb, 0xee, 0x7d,
	0xaf, 0xc0, 0xda, 0x4c, 0x29, 0x99, 0xf3, 0x4c, 0x22, 0xf9, 0x1f, 0x6a, 0x02, 0x55, 0xcc, 0x13,
	0x34, 0xd5, 0xda, 0x47, 0x6b, 0x07, 0xd3, 0x6f, 0xc0, 0x41, 0x60, 0x85, 0x60, 0x1a, 0x41, 0x0e,
	0xa1, 0x51, 0x94, 0xa3, 0x5e, 0xd7, 0xef, 0x37, 0x8f, 0xd6, 0xcb, 0xf0, 0x32, 0x79, 0x19, 0x65,
	0xae, 0x35, 0x4f, 0x53, 0xcc, 0x94, 0x5b, 0xda, 0x29, 0xd4, 0x0a, 0x0e, 0xa3, 0x5c, 0x62, 0xf1,
	0x29, 0x70, 0xb0, 0xf7, 0xad, 0x02, 0x1b, 0x45, 0xb2, 0x73, 0x2e, 0xd5, 0xd3, 0xc6, 0xcc, 0xbd,
	0xb5, 0xb7, 0x30, 0xab, 0xb9, 0x96, 0xfd, 0xdf, 0x6a, 0xb9, 0x0f, 0xd5, 0x28, 0x2e, 0x2e, 0x4e,
	0xfb, 0xa8, 0x53, 0xc6, 0x9f, 0x18, 0x3e, 0x70, 0x7a, 0xef, 0x23, 0x6c, 0x2e, 0xf4, 0xf9, 0x12,
	0x57, 0x67, 0x2c, 0xf2, 0x1e, 0xb5, 0xc8, 0x9f, 0xb7, 0xe8, 0x3d, 0x6c, 0x2d, 0x7c, 0x7b, 0x9e,
	0xf6, 0x68, 0x7e, 0x65, 0xcd, 0x01, 0x67, 0xd5, 0xcc, 0xca, 0x1a, 0xba, 0xf7, 0xa3, 0x02, 0xdb,
	0xf7, 0x72, 0xbf, 0xe4, 0xbd, 0xce, 0x1e, 0xac, 0xa9, 0x27, 0xb0, 0xf3, 0xc0, 0x04, 0x5c, 0xa5,
	0xc5, 0x76, 0x5e, 0xb4, 0x40, 0x5f, 0x2b, 0xb0, 0xbb, 0x90, 0xf8, 0xf9, 0x35, 0x3a, 0x7b, 0xc4,
	0xa2, 0x3f, 0x6b, 0xb7, 0x5c, 0x1e, 0xff, 0x99, 0xe5, 0xf9, 0x0c, 0x7b, 0x0f, 0xf6, 0xf8, 0x97,
	0x56, 0xe8, 0xb2, 0x6a, 0x7e, 0xee, 0xc7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xb6, 0xa2,
	0x95, 0x06, 0x08, 0x00, 0x00,
}
