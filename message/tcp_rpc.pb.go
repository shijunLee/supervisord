// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tcp_rpc.proto

package message

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

// message id 1
type HeartbeatAck struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatAck) Reset()         { *m = HeartbeatAck{} }
func (m *HeartbeatAck) String() string { return proto.CompactTextString(m) }
func (*HeartbeatAck) ProtoMessage()    {}
func (*HeartbeatAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{0}
}

func (m *HeartbeatAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatAck.Unmarshal(m, b)
}
func (m *HeartbeatAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatAck.Marshal(b, m, deterministic)
}
func (m *HeartbeatAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatAck.Merge(m, src)
}
func (m *HeartbeatAck) XXX_Size() int {
	return xxx_messageInfo_HeartbeatAck.Size(m)
}
func (m *HeartbeatAck) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatAck.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatAck proto.InternalMessageInfo

func (m *HeartbeatAck) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// message id 2
type HeartbeatReq struct {
	Id                   int32    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatReq) Reset()         { *m = HeartbeatReq{} }
func (m *HeartbeatReq) String() string { return proto.CompactTextString(m) }
func (*HeartbeatReq) ProtoMessage()    {}
func (*HeartbeatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{1}
}

func (m *HeartbeatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatReq.Unmarshal(m, b)
}
func (m *HeartbeatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatReq.Marshal(b, m, deterministic)
}
func (m *HeartbeatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatReq.Merge(m, src)
}
func (m *HeartbeatReq) XXX_Size() int {
	return xxx_messageInfo_HeartbeatReq.Size(m)
}
func (m *HeartbeatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatReq.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatReq proto.InternalMessageInfo

func (m *HeartbeatReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type InstanceLoginReq struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	InstanceID           string   `protobuf:"bytes,2,opt,name=instanceID,proto3" json:"instanceID,omitempty"`
	GroupID              string   `protobuf:"bytes,3,opt,name=groupID,proto3" json:"groupID,omitempty"`
	AccessToken          string   `protobuf:"bytes,4,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstanceLoginReq) Reset()         { *m = InstanceLoginReq{} }
func (m *InstanceLoginReq) String() string { return proto.CompactTextString(m) }
func (*InstanceLoginReq) ProtoMessage()    {}
func (*InstanceLoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{2}
}

func (m *InstanceLoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceLoginReq.Unmarshal(m, b)
}
func (m *InstanceLoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceLoginReq.Marshal(b, m, deterministic)
}
func (m *InstanceLoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceLoginReq.Merge(m, src)
}
func (m *InstanceLoginReq) XXX_Size() int {
	return xxx_messageInfo_InstanceLoginReq.Size(m)
}
func (m *InstanceLoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceLoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceLoginReq proto.InternalMessageInfo

func (m *InstanceLoginReq) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *InstanceLoginReq) GetInstanceID() string {
	if m != nil {
		return m.InstanceID
	}
	return ""
}

func (m *InstanceLoginReq) GetGroupID() string {
	if m != nil {
		return m.GroupID
	}
	return ""
}

func (m *InstanceLoginReq) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

// message login id 4
type InstanceLoginAck struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstanceLoginAck) Reset()         { *m = InstanceLoginAck{} }
func (m *InstanceLoginAck) String() string { return proto.CompactTextString(m) }
func (*InstanceLoginAck) ProtoMessage()    {}
func (*InstanceLoginAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{3}
}

func (m *InstanceLoginAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstanceLoginAck.Unmarshal(m, b)
}
func (m *InstanceLoginAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstanceLoginAck.Marshal(b, m, deterministic)
}
func (m *InstanceLoginAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstanceLoginAck.Merge(m, src)
}
func (m *InstanceLoginAck) XXX_Size() int {
	return xxx_messageInfo_InstanceLoginAck.Size(m)
}
func (m *InstanceLoginAck) XXX_DiscardUnknown() {
	xxx_messageInfo_InstanceLoginAck.DiscardUnknown(m)
}

var xxx_messageInfo_InstanceLoginAck proto.InternalMessageInfo

func (m *InstanceLoginAck) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

// message start app 5
type StartAppReq struct {
	ApplicationName      string   `protobuf:"bytes,1,opt,name=applicationName,proto3" json:"applicationName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartAppReq) Reset()         { *m = StartAppReq{} }
func (m *StartAppReq) String() string { return proto.CompactTextString(m) }
func (*StartAppReq) ProtoMessage()    {}
func (*StartAppReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{4}
}

func (m *StartAppReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartAppReq.Unmarshal(m, b)
}
func (m *StartAppReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartAppReq.Marshal(b, m, deterministic)
}
func (m *StartAppReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartAppReq.Merge(m, src)
}
func (m *StartAppReq) XXX_Size() int {
	return xxx_messageInfo_StartAppReq.Size(m)
}
func (m *StartAppReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StartAppReq.DiscardUnknown(m)
}

var xxx_messageInfo_StartAppReq proto.InternalMessageInfo

func (m *StartAppReq) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

// message start app 6
type StartAppAck struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartAppAck) Reset()         { *m = StartAppAck{} }
func (m *StartAppAck) String() string { return proto.CompactTextString(m) }
func (*StartAppAck) ProtoMessage()    {}
func (*StartAppAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{5}
}

func (m *StartAppAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartAppAck.Unmarshal(m, b)
}
func (m *StartAppAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartAppAck.Marshal(b, m, deterministic)
}
func (m *StartAppAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartAppAck.Merge(m, src)
}
func (m *StartAppAck) XXX_Size() int {
	return xxx_messageInfo_StartAppAck.Size(m)
}
func (m *StartAppAck) XXX_DiscardUnknown() {
	xxx_messageInfo_StartAppAck.DiscardUnknown(m)
}

var xxx_messageInfo_StartAppAck proto.InternalMessageInfo

func (m *StartAppAck) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *StartAppAck) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// message start app 7
type StopAppReq struct {
	ApplicationName      string   `protobuf:"bytes,1,opt,name=applicationName,proto3" json:"applicationName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopAppReq) Reset()         { *m = StopAppReq{} }
func (m *StopAppReq) String() string { return proto.CompactTextString(m) }
func (*StopAppReq) ProtoMessage()    {}
func (*StopAppReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{6}
}

func (m *StopAppReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopAppReq.Unmarshal(m, b)
}
func (m *StopAppReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopAppReq.Marshal(b, m, deterministic)
}
func (m *StopAppReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopAppReq.Merge(m, src)
}
func (m *StopAppReq) XXX_Size() int {
	return xxx_messageInfo_StopAppReq.Size(m)
}
func (m *StopAppReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StopAppReq.DiscardUnknown(m)
}

var xxx_messageInfo_StopAppReq proto.InternalMessageInfo

func (m *StopAppReq) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

// message stop app 8
type StopAppAck struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopAppAck) Reset()         { *m = StopAppAck{} }
func (m *StopAppAck) String() string { return proto.CompactTextString(m) }
func (*StopAppAck) ProtoMessage()    {}
func (*StopAppAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{7}
}

func (m *StopAppAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopAppAck.Unmarshal(m, b)
}
func (m *StopAppAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopAppAck.Marshal(b, m, deterministic)
}
func (m *StopAppAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopAppAck.Merge(m, src)
}
func (m *StopAppAck) XXX_Size() int {
	return xxx_messageInfo_StopAppAck.Size(m)
}
func (m *StopAppAck) XXX_DiscardUnknown() {
	xxx_messageInfo_StopAppAck.DiscardUnknown(m)
}

var xxx_messageInfo_StopAppAck proto.InternalMessageInfo

func (m *StopAppAck) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *StopAppAck) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//message get log request 9
type GetLogReq struct {
	ApplicationName      string   `protobuf:"bytes,1,opt,name=applicationName,proto3" json:"applicationName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLogReq) Reset()         { *m = GetLogReq{} }
func (m *GetLogReq) String() string { return proto.CompactTextString(m) }
func (*GetLogReq) ProtoMessage()    {}
func (*GetLogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{8}
}

func (m *GetLogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLogReq.Unmarshal(m, b)
}
func (m *GetLogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLogReq.Marshal(b, m, deterministic)
}
func (m *GetLogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLogReq.Merge(m, src)
}
func (m *GetLogReq) XXX_Size() int {
	return xxx_messageInfo_GetLogReq.Size(m)
}
func (m *GetLogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLogReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetLogReq proto.InternalMessageInfo

func (m *GetLogReq) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

// message get log result 10
type GetLogAck struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLogAck) Reset()         { *m = GetLogAck{} }
func (m *GetLogAck) String() string { return proto.CompactTextString(m) }
func (*GetLogAck) ProtoMessage()    {}
func (*GetLogAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{9}
}

func (m *GetLogAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLogAck.Unmarshal(m, b)
}
func (m *GetLogAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLogAck.Marshal(b, m, deterministic)
}
func (m *GetLogAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLogAck.Merge(m, src)
}
func (m *GetLogAck) XXX_Size() int {
	return xxx_messageInfo_GetLogAck.Size(m)
}
func (m *GetLogAck) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLogAck.DiscardUnknown(m)
}

var xxx_messageInfo_GetLogAck proto.InternalMessageInfo

func (m *GetLogAck) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GetLogAck) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type StartProgramesReq struct {
	ApplicationNames     []string `protobuf:"bytes,1,rep,name=applicationNames,proto3" json:"applicationNames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartProgramesReq) Reset()         { *m = StartProgramesReq{} }
func (m *StartProgramesReq) String() string { return proto.CompactTextString(m) }
func (*StartProgramesReq) ProtoMessage()    {}
func (*StartProgramesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{10}
}

func (m *StartProgramesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartProgramesReq.Unmarshal(m, b)
}
func (m *StartProgramesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartProgramesReq.Marshal(b, m, deterministic)
}
func (m *StartProgramesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartProgramesReq.Merge(m, src)
}
func (m *StartProgramesReq) XXX_Size() int {
	return xxx_messageInfo_StartProgramesReq.Size(m)
}
func (m *StartProgramesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StartProgramesReq.DiscardUnknown(m)
}

var xxx_messageInfo_StartProgramesReq proto.InternalMessageInfo

func (m *StartProgramesReq) GetApplicationNames() []string {
	if m != nil {
		return m.ApplicationNames
	}
	return nil
}

// message start startPrograms 12
type StartProgramsAck struct {
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartProgramsAck) Reset()         { *m = StartProgramsAck{} }
func (m *StartProgramsAck) String() string { return proto.CompactTextString(m) }
func (*StartProgramsAck) ProtoMessage()    {}
func (*StartProgramsAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{11}
}

func (m *StartProgramsAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartProgramsAck.Unmarshal(m, b)
}
func (m *StartProgramsAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartProgramsAck.Marshal(b, m, deterministic)
}
func (m *StartProgramsAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartProgramsAck.Merge(m, src)
}
func (m *StartProgramsAck) XXX_Size() int {
	return xxx_messageInfo_StartProgramsAck.Size(m)
}
func (m *StartProgramsAck) XXX_DiscardUnknown() {
	xxx_messageInfo_StartProgramsAck.DiscardUnknown(m)
}

var xxx_messageInfo_StartProgramsAck proto.InternalMessageInfo

func (m *StartProgramsAck) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StartProgramsAck) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

// message start startPrograms 13
type StopProgramesReq struct {
	ApplicationNames     []string `protobuf:"bytes,1,rep,name=applicationNames,proto3" json:"applicationNames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopProgramesReq) Reset()         { *m = StopProgramesReq{} }
func (m *StopProgramesReq) String() string { return proto.CompactTextString(m) }
func (*StopProgramesReq) ProtoMessage()    {}
func (*StopProgramesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{12}
}

func (m *StopProgramesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopProgramesReq.Unmarshal(m, b)
}
func (m *StopProgramesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopProgramesReq.Marshal(b, m, deterministic)
}
func (m *StopProgramesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopProgramesReq.Merge(m, src)
}
func (m *StopProgramesReq) XXX_Size() int {
	return xxx_messageInfo_StopProgramesReq.Size(m)
}
func (m *StopProgramesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StopProgramesReq.DiscardUnknown(m)
}

var xxx_messageInfo_StopProgramesReq proto.InternalMessageInfo

func (m *StopProgramesReq) GetApplicationNames() []string {
	if m != nil {
		return m.ApplicationNames
	}
	return nil
}

// message start startPrograms 14
type StopProgramsAck struct {
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StopProgramsAck) Reset()         { *m = StopProgramsAck{} }
func (m *StopProgramsAck) String() string { return proto.CompactTextString(m) }
func (*StopProgramsAck) ProtoMessage()    {}
func (*StopProgramsAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3dfdba9700c00cfa, []int{13}
}

func (m *StopProgramsAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StopProgramsAck.Unmarshal(m, b)
}
func (m *StopProgramsAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StopProgramsAck.Marshal(b, m, deterministic)
}
func (m *StopProgramsAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StopProgramsAck.Merge(m, src)
}
func (m *StopProgramsAck) XXX_Size() int {
	return xxx_messageInfo_StopProgramsAck.Size(m)
}
func (m *StopProgramsAck) XXX_DiscardUnknown() {
	xxx_messageInfo_StopProgramsAck.DiscardUnknown(m)
}

var xxx_messageInfo_StopProgramsAck proto.InternalMessageInfo

func (m *StopProgramsAck) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *StopProgramsAck) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*HeartbeatAck)(nil), "message.HeartbeatAck")
	proto.RegisterType((*HeartbeatReq)(nil), "message.HeartbeatReq")
	proto.RegisterType((*InstanceLoginReq)(nil), "message.InstanceLoginReq")
	proto.RegisterType((*InstanceLoginAck)(nil), "message.InstanceLoginAck")
	proto.RegisterType((*StartAppReq)(nil), "message.StartAppReq")
	proto.RegisterType((*StartAppAck)(nil), "message.StartAppAck")
	proto.RegisterType((*StopAppReq)(nil), "message.StopAppReq")
	proto.RegisterType((*StopAppAck)(nil), "message.StopAppAck")
	proto.RegisterType((*GetLogReq)(nil), "message.GetLogReq")
	proto.RegisterType((*GetLogAck)(nil), "message.GetLogAck")
	proto.RegisterType((*StartProgramesReq)(nil), "message.startProgramesReq")
	proto.RegisterType((*StartProgramsAck)(nil), "message.startProgramsAck")
	proto.RegisterType((*StopProgramesReq)(nil), "message.stopProgramesReq")
	proto.RegisterType((*StopProgramsAck)(nil), "message.stopProgramsAck")
}

func init() { proto.RegisterFile("tcp_rpc.proto", fileDescriptor_3dfdba9700c00cfa) }

var fileDescriptor_3dfdba9700c00cfa = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6a, 0xc2, 0x40,
	0x10, 0x25, 0x5a, 0xb5, 0x19, 0x6d, 0xb5, 0xa1, 0x94, 0x9c, 0x44, 0xf6, 0x24, 0x1e, 0x7a, 0x29,
	0x6d, 0x6f, 0x8a, 0x54, 0x68, 0x03, 0xb6, 0x94, 0xd8, 0x7b, 0x59, 0xd7, 0x21, 0x04, 0x35, 0xbb,
	0xdd, 0x1d, 0xff, 0xa0, 0x1f, 0x5e, 0x36, 0x46, 0x9b, 0xa4, 0x50, 0x30, 0xc7, 0x79, 0x6f, 0xe7,
	0xbd, 0x37, 0x0f, 0x16, 0x2e, 0x48, 0xa8, 0x4f, 0xad, 0xc4, 0xad, 0xd2, 0x92, 0xa4, 0xd7, 0xda,
	0xa2, 0x31, 0x3c, 0x42, 0xd6, 0x87, 0xce, 0x0b, 0x72, 0x4d, 0x4b, 0xe4, 0x34, 0x15, 0x6b, 0xef,
	0x12, 0x6a, 0xf1, 0xca, 0x77, 0x06, 0xce, 0xb0, 0x11, 0xd6, 0xe2, 0x55, 0x81, 0x0f, 0xf1, 0x2b,
	0xe3, 0x6b, 0x47, 0xfe, 0xdb, 0x81, 0x5e, 0x90, 0x18, 0xe2, 0x89, 0xc0, 0xb9, 0x8c, 0xe2, 0xc4,
	0x3e, 0xba, 0x86, 0x06, 0x57, 0x2a, 0xd8, 0xeb, 0xb8, 0xe1, 0x7e, 0xf0, 0xfa, 0x00, 0x71, 0xf6,
	0x32, 0x98, 0xa5, 0x12, 0x6e, 0x98, 0x43, 0x3c, 0x1f, 0x5a, 0x91, 0x96, 0x3b, 0x15, 0xcc, 0xfc,
	0x7a, 0x4a, 0x1e, 0x46, 0x6f, 0x00, 0x6d, 0x2e, 0x04, 0x1a, 0xf3, 0x21, 0xd7, 0x98, 0xf8, 0x67,
	0x29, 0x9b, 0x87, 0xd8, 0xa8, 0x94, 0xc2, 0x9e, 0x72, 0x03, 0x4d, 0x8d, 0x66, 0xb7, 0xa1, 0x34,
	0xc6, 0x79, 0x98, 0x4d, 0xec, 0x11, 0xda, 0x0b, 0xe2, 0x9a, 0xa6, 0x4a, 0xd9, 0xb0, 0x43, 0xe8,
	0x72, 0xa5, 0x36, 0xb1, 0xe0, 0x14, 0xcb, 0xe4, 0x8d, 0x6f, 0x31, 0x8b, 0x5d, 0x86, 0xd9, 0xe4,
	0x77, 0xf1, 0x1f, 0x7d, 0x7b, 0x47, 0xd6, 0x6e, 0x76, 0xe4, 0xb1, 0xec, 0x07, 0x80, 0x05, 0x49,
	0x75, 0xb2, 0xf1, 0xf8, 0xb8, 0x57, 0xcd, 0xf7, 0x1e, 0xdc, 0x67, 0xa4, 0xb9, 0x8c, 0x4e, 0xb3,
	0x0d, 0x0e, 0x6b, 0xd6, 0x35, 0xa7, 0xee, 0x14, 0xd4, 0x3d, 0x06, 0x1d, 0xd4, 0x5a, 0xea, 0xd7,
	0x82, 0x79, 0x01, 0x63, 0x13, 0xb8, 0x32, 0xb6, 0xba, 0x77, 0x2d, 0x23, 0xcd, 0xb7, 0x68, 0x6c,
	0x92, 0x11, 0xf4, 0x4a, 0x96, 0xc6, 0x77, 0x06, 0xf5, 0xa1, 0x1b, 0xfe, 0xc1, 0xd9, 0x0c, 0x7a,
	0x79, 0x01, 0x53, 0x8a, 0x54, 0x3c, 0xd8, 0x56, 0x64, 0x88, 0xd3, 0xce, 0x1c, 0x2a, 0xda, 0x4f,
	0x6c, 0x6c, 0x55, 0xa4, 0xaa, 0x9c, 0xe2, 0x09, 0xba, 0xb9, 0xfd, 0x6a, 0x21, 0x96, 0xcd, 0xf4,
	0x0b, 0xde, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x11, 0xe2, 0xa0, 0x36, 0x93, 0x03, 0x00, 0x00,
}