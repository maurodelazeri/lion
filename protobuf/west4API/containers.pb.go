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
	Containers           []*Container `protobuf:"bytes,2,rep,name=containers,proto3" json:"containers,omitempty"`
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

func (m *ContainerResponse) GetContainers() []*Container {
	if m != nil {
		return m.Containers
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
	Containers           []*Container `protobuf:"bytes,3,rep,name=containers,proto3" json:"containers,omitempty"`
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

func (m *ContainerPostRequest) GetContainers() []*Container {
	if m != nil {
		return m.Containers
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
	ContainersCreate     []*ContainerCreate `protobuf:"bytes,2,rep,name=containers_create,json=containersCreate,proto3" json:"containers_create,omitempty"`
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

func (m *ContainerCreateResponse) GetContainersCreate() []*ContainerCreate {
	if m != nil {
		return m.ContainersCreate
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
	// 749 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x55, 0x4b, 0x6e, 0xdb, 0x48,
	0x10, 0x85, 0x48, 0x5b, 0x9f, 0x92, 0x2c, 0x4b, 0xed, 0x5f, 0xdb, 0x5e, 0x8c, 0x2c, 0x60, 0x00,
	0x0d, 0x06, 0x63, 0x60, 0xec, 0x5c, 0xc0, 0xb0, 0x61, 0x40, 0x40, 0x16, 0x06, 0xe1, 0x2c, 0xb2,
	0x22, 0x68, 0xb2, 0x6c, 0x37, 0x2c, 0xb2, 0x19, 0x76, 0xcb, 0x8a, 0xb2, 0xc8, 0x25, 0x72, 0x83,
	0xac, 0x92, 0x6b, 0x24, 0x17, 0x0b, 0xfa, 0xc3, 0x8f, 0xe4, 0x5f, 0xe0, 0x45, 0x76, 0x7c, 0xef,
	0x55, 0x57, 0x15, 0x5f, 0x15, 0x9b, 0xd0, 0x0b, 0x79, 0x22, 0x03, 0x96, 0x60, 0x26, 0x0e, 0xd3,
	0x8c, 0x4b, 0x4e, 0x9a, 0x33, 0x14, 0xf2, 0xcd, 0xc9, 0xc5, 0x78, 0x0f, 0x30, 0x99, 0xc6, 0x86,
	0x1d, 0x7e, 0x71, 0xa0, 0x75, 0x9a, 0x87, 0x92, 0x2e, 0x38, 0x2c, 0xa2, 0xb5, 0x41, 0x6d, 0xd4,
	0xf2, 0x1c, 0x16, 0x91, 0x4d, 0x58, 0x4d, 0x82, 0x18, 0x05, 0x75, 0x34, 0x65, 0x80, 0x62, 0x59,
	0x1c, 0xdc, 0x20, 0x75, 0x0d, 0xab, 0x01, 0xd9, 0x85, 0xa6, 0x7e, 0xf0, 0x59, 0x44, 0x57, 0xb4,
	0xd0, 0xd0, 0x78, 0x1c, 0x11, 0x0a, 0x8d, 0x30, 0xc3, 0x40, 0x62, 0x44, 0x57, 0x8d, 0x62, 0xa1,
	0x4a, 0x95, 0xf2, 0x4c, 0x0a, 0x5a, 0x37, 0xa9, 0x34, 0x20, 0x3b, 0xd0, 0x10, 0xec, 0x13, 0xfa,
	0xd9, 0x8c, 0x36, 0x06, 0xb5, 0x91, 0xeb, 0xd5, 0x15, 0xf4, 0x66, 0x64, 0x00, 0x1d, 0x23, 0x70,
	0x2e, 0xfd, 0x6b, 0x41, 0x9b, 0x5a, 0x05, 0xad, 0x72, 0x2e, 0xcf, 0x05, 0xd9, 0x86, 0xfa, 0x24,
	0xb8, 0xc2, 0x89, 0xa0, 0x2d, 0x9d, 0xd1, 0x22, 0x55, 0x48, 0xc8, 0x40, 0x22, 0x05, 0x53, 0x48,
	0x03, 0x15, 0xad, 0x1e, 0xa6, 0x82, 0xb6, 0x4d, 0xb4, 0x41, 0xc3, 0x9f, 0x2b, 0xb0, 0x5e, 0xb8,
	0x72, 0xaa, 0x7b, 0x25, 0x07, 0xd0, 0x29, 0x3c, 0xf5, 0xad, 0x4b, 0xae, 0xd7, 0x2e, 0xb8, 0x71,
	0x44, 0xfe, 0x87, 0xcd, 0x6a, 0x08, 0x26, 0x92, 0x5d, 0x33, 0xcc, 0xac, 0x7b, 0x1b, 0x95, 0xd0,
	0x5c, 0x22, 0x7f, 0x43, 0xb7, 0x3c, 0x22, 0xe7, 0x69, 0x6e, 0xea, 0x5a, 0xc1, 0x5e, 0xce, 0x53,
	0x24, 0xc7, 0xb0, 0x55, 0x86, 0x45, 0x28, 0xc2, 0x8c, 0xa5, 0x92, 0xf1, 0xc4, 0x3a, 0x5d, 0x96,
	0x3d, 0x2b, 0xb5, 0x72, 0x4e, 0xab, 0xd5, 0x39, 0x51, 0x68, 0x24, 0x28, 0x67, 0x3c, 0xbb, 0xb3,
	0xa6, 0xe7, 0x90, 0xec, 0x41, 0xf3, 0x96, 0x0b, 0xa9, 0x86, 0xac, 0x7d, 0x6f, 0x79, 0x05, 0x26,
	0x3d, 0x70, 0x31, 0xb9, 0xd7, 0x86, 0xb7, 0x3c, 0xf5, 0xa8, 0x3a, 0xcf, 0x50, 0xc8, 0x20, 0x93,
	0x7e, 0xca, 0x27, 0x2c, 0x9c, 0x5b, 0xc7, 0xd7, 0x2c, 0x7b, 0xa1, 0x49, 0xf2, 0x1f, 0x90, 0x18,
	0x63, 0x9e, 0xcd, 0xfd, 0x0c, 0x05, 0x66, 0xf7, 0x81, 0x6e, 0x1b, 0xb4, 0x79, 0x7d, 0xa3, 0x78,
	0xa5, 0xa0, 0x5c, 0xb6, 0xe1, 0x13, 0x16, 0x33, 0xa9, 0xe7, 0xe2, 0x7a, 0x6d, 0xc3, 0xbd, 0x55,
	0x14, 0xd9, 0x87, 0x56, 0x98, 0x4e, 0xad, 0xde, 0xd1, 0x7a, 0x33, 0x4c, 0xa7, 0x46, 0x3c, 0x80,
	0x8e, 0x90, 0x3c, 0xf5, 0x25, 0x8b, 0x91, 0x4f, 0x25, 0x5d, 0x33, 0xe7, 0x15, 0x77, 0x69, 0x28,
	0x75, 0x5e, 0xd5, 0x33, 0x53, 0xec, 0x9a, 0xf3, 0x86, 0x18, 0x47, 0xe4, 0x9f, 0xca, 0x97, 0xe3,
	0x87, 0x3c, 0xb9, 0x66, 0x37, 0x74, 0x5d, 0xbf, 0xd7, 0x7a, 0xc1, 0x9f, 0x6a, 0x9a, 0xfc, 0x05,
	0xed, 0x49, 0x20, 0xa4, 0x3f, 0x4d, 0x23, 0xb5, 0x58, 0x3d, 0x1d, 0x05, 0x8a, 0x7a, 0xa7, 0x99,
	0xe1, 0x2d, 0xf4, 0x8a, 0x25, 0xf2, 0xf0, 0xc3, 0x14, 0x85, 0x54, 0x33, 0x91, 0xfc, 0x0e, 0x13,
	0xfb, 0x91, 0x19, 0xf0, 0x60, 0xb7, 0x9c, 0x87, 0xbb, 0xb5, 0xd0, 0xb5, 0xbb, 0xd8, 0xf5, 0xf0,
	0x7b, 0x0d, 0xfa, 0x95, 0x52, 0x22, 0xe5, 0x89, 0x40, 0xf2, 0x2f, 0x34, 0x32, 0x94, 0x21, 0x8f,
	0x50, 0x57, 0xeb, 0x1e, 0xf5, 0x0f, 0xf3, 0x3b, 0xe0, 0xd0, 0x33, 0x82, 0x97, 0x47, 0x90, 0x63,
	0x80, 0xf2, 0xca, 0xa0, 0xce, 0xc0, 0x1d, 0xb5, 0x8f, 0x36, 0xca, 0xf8, 0x32, 0x7b, 0x25, 0x4c,
	0x7f, 0xd8, 0x3c, 0x8e, 0x31, 0x91, 0x76, 0x6d, 0x73, 0xa8, 0x14, 0x9c, 0x04, 0xa9, 0xc0, 0xe2,
	0x32, 0xb0, 0x70, 0xf8, 0xad, 0x06, 0x9b, 0x45, 0xb6, 0x0b, 0x2e, 0xe4, 0xf3, 0xd6, 0x2c, 0xbc,
	0xb7, 0xb3, 0x34, 0xad, 0xc5, 0xa6, 0xdd, 0xdf, 0x6b, 0x7a, 0x04, 0xf5, 0x20, 0x2c, 0x3e, 0x9e,
	0xee, 0x51, 0xaf, 0x3c, 0x70, 0xa2, 0x79, 0xcf, 0xea, 0xc3, 0x8f, 0xb0, 0xb5, 0xd4, 0xe9, 0x6b,
	0x9c, 0xad, 0x98, 0xe4, 0x3c, 0x69, 0x92, 0xbb, 0x68, 0xd2, 0x7b, 0xd8, 0x5e, 0xba, 0x7f, 0x9e,
	0x77, 0x69, 0x71, 0x6d, 0xf5, 0x01, 0x6b, 0x56, 0x65, 0x6d, 0x35, 0x3d, 0xfc, 0x51, 0x83, 0x9d,
	0x07, 0xb9, 0x5f, 0xf3, 0x5e, 0xe7, 0xd0, 0x2f, 0x5d, 0x2d, 0x8b, 0xaa, 0x19, 0xec, 0x3e, 0x32,
	0x03, 0x5b, 0xaa, 0xf2, 0x63, 0xb2, 0x17, 0xeb, 0x6b, 0x96, 0xe8, 0x6b, 0x0d, 0xf6, 0x96, 0x32,
	0xbf, 0xbc, 0x4a, 0x67, 0x4f, 0x98, 0xf4, 0x6c, 0xbf, 0xcb, 0xfe, 0x55, 0xd6, 0xc7, 0x7d, 0x61,
	0x7d, 0x3e, 0xc3, 0xfe, 0xa3, 0x3d, 0xfe, 0xa1, 0x25, 0xba, 0xaa, 0xeb, 0x5f, 0xfc, 0xf1, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x99, 0xd3, 0xb6, 0x41, 0x0c, 0x08, 0x00, 0x00,
}
