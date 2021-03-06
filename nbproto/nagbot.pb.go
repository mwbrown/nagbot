// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nbproto/nagbot.proto

/*
Package nbproto is a generated protocol buffer package.

It is generated from these files:
	nbproto/nagbot.proto

It has these top-level messages:
	OwnerInfo
	TaskDefinition
	TaskSchedule
	TaskInstance
	LoginRequest
	LoginResponse
	LogoutRequest
	LogoutResponse
	CheckLoginRequest
	CheckLoginResponse
	AddTaskDefRequest
	AddTaskDefResponse
	DelTaskDefRequest
	DelTaskDefResponse
*/
package nbproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type OwnerInfo_OwnerType int32

const (
	OwnerInfo_UNKNOWN OwnerInfo_OwnerType = 0
	OwnerInfo_USER    OwnerInfo_OwnerType = 1
	OwnerInfo_GROUP   OwnerInfo_OwnerType = 2
)

var OwnerInfo_OwnerType_name = map[int32]string{
	0: "UNKNOWN",
	1: "USER",
	2: "GROUP",
}
var OwnerInfo_OwnerType_value = map[string]int32{
	"UNKNOWN": 0,
	"USER":    1,
	"GROUP":   2,
}

func (x OwnerInfo_OwnerType) String() string {
	return proto.EnumName(OwnerInfo_OwnerType_name, int32(x))
}
func (OwnerInfo_OwnerType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type TaskSchedule_ScheduleInfo_ScheduleType int32

const (
	TaskSchedule_ScheduleInfo_UNKNOWN       TaskSchedule_ScheduleInfo_ScheduleType = 0
	TaskSchedule_ScheduleInfo_ONESHOT       TaskSchedule_ScheduleInfo_ScheduleType = 1
	TaskSchedule_ScheduleInfo_INTERVAL      TaskSchedule_ScheduleInfo_ScheduleType = 2
	TaskSchedule_ScheduleInfo_WEEKLY        TaskSchedule_ScheduleInfo_ScheduleType = 3
	TaskSchedule_ScheduleInfo_MONTH_DAY     TaskSchedule_ScheduleInfo_ScheduleType = 4
	TaskSchedule_ScheduleInfo_MONTH_WEEKDAY TaskSchedule_ScheduleInfo_ScheduleType = 5
	TaskSchedule_ScheduleInfo_ANNUAL        TaskSchedule_ScheduleInfo_ScheduleType = 6
)

var TaskSchedule_ScheduleInfo_ScheduleType_name = map[int32]string{
	0: "UNKNOWN",
	1: "ONESHOT",
	2: "INTERVAL",
	3: "WEEKLY",
	4: "MONTH_DAY",
	5: "MONTH_WEEKDAY",
	6: "ANNUAL",
}
var TaskSchedule_ScheduleInfo_ScheduleType_value = map[string]int32{
	"UNKNOWN":       0,
	"ONESHOT":       1,
	"INTERVAL":      2,
	"WEEKLY":        3,
	"MONTH_DAY":     4,
	"MONTH_WEEKDAY": 5,
	"ANNUAL":        6,
}

func (x TaskSchedule_ScheduleInfo_ScheduleType) String() string {
	return proto.EnumName(TaskSchedule_ScheduleInfo_ScheduleType_name, int32(x))
}
func (TaskSchedule_ScheduleInfo_ScheduleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0, 0}
}

// TODO: decide if this is the right approach to use separate ID spaces
// for groups and users, or allow them to overlap.
type OwnerInfo struct {
	Id   uint32              `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type OwnerInfo_OwnerType `protobuf:"varint,2,opt,name=type,enum=nbproto.OwnerInfo_OwnerType" json:"type,omitempty"`
}

func (m *OwnerInfo) Reset()                    { *m = OwnerInfo{} }
func (m *OwnerInfo) String() string            { return proto.CompactTextString(m) }
func (*OwnerInfo) ProtoMessage()               {}
func (*OwnerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OwnerInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OwnerInfo) GetType() OwnerInfo_OwnerType {
	if m != nil {
		return m.Type
	}
	return OwnerInfo_UNKNOWN
}

// Defines an individual Todo task, of which multiple instances
// can exist.
type TaskDefinition struct {
	Id    uint32     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Desc  string     `protobuf:"bytes,2,opt,name=desc" json:"desc,omitempty"`
	Owner *OwnerInfo `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
}

func (m *TaskDefinition) Reset()                    { *m = TaskDefinition{} }
func (m *TaskDefinition) String() string            { return proto.CompactTextString(m) }
func (*TaskDefinition) ProtoMessage()               {}
func (*TaskDefinition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TaskDefinition) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TaskDefinition) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *TaskDefinition) GetOwner() *OwnerInfo {
	if m != nil {
		return m.Owner
	}
	return nil
}

type TaskSchedule struct {
	Id       uint32                     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	TaskId   uint32                     `protobuf:"varint,2,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	Owner    *OwnerInfo                 `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	Schedule *TaskSchedule_ScheduleInfo `protobuf:"bytes,4,opt,name=schedule" json:"schedule,omitempty"`
	NextDue  uint64                     `protobuf:"varint,5,opt,name=next_due,json=nextDue" json:"next_due,omitempty"`
	IsActive bool                       `protobuf:"varint,6,opt,name=is_active,json=isActive" json:"is_active,omitempty"`
}

func (m *TaskSchedule) Reset()                    { *m = TaskSchedule{} }
func (m *TaskSchedule) String() string            { return proto.CompactTextString(m) }
func (*TaskSchedule) ProtoMessage()               {}
func (*TaskSchedule) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TaskSchedule) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TaskSchedule) GetTaskId() uint32 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *TaskSchedule) GetOwner() *OwnerInfo {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *TaskSchedule) GetSchedule() *TaskSchedule_ScheduleInfo {
	if m != nil {
		return m.Schedule
	}
	return nil
}

func (m *TaskSchedule) GetNextDue() uint64 {
	if m != nil {
		return m.NextDue
	}
	return 0
}

func (m *TaskSchedule) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

type TaskSchedule_ScheduleInfo struct {
	Type TaskSchedule_ScheduleInfo_ScheduleType `protobuf:"varint,1,opt,name=type,enum=nbproto.TaskSchedule_ScheduleInfo_ScheduleType" json:"type,omitempty"`
	// This controls whether a more advanced match will occur if the date
	// in question does not exist in a given month. For instance, Feb 29th
	// is not guaranteed to happen, so if exact_only is false, normally the
	// event would be generated once Mar 1st occurs. If exact_only is true,
	// the event would be skipped every year except leap years.
	ExactOnly bool `protobuf:"varint,2,opt,name=exact_only,json=exactOnly" json:"exact_only,omitempty"`
	// Various optional members to control timing information.
	Time    uint32 `protobuf:"varint,3,opt,name=time" json:"time,omitempty"`
	Weekday uint32 `protobuf:"varint,4,opt,name=weekday" json:"weekday,omitempty"`
}

func (m *TaskSchedule_ScheduleInfo) Reset()                    { *m = TaskSchedule_ScheduleInfo{} }
func (m *TaskSchedule_ScheduleInfo) String() string            { return proto.CompactTextString(m) }
func (*TaskSchedule_ScheduleInfo) ProtoMessage()               {}
func (*TaskSchedule_ScheduleInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *TaskSchedule_ScheduleInfo) GetType() TaskSchedule_ScheduleInfo_ScheduleType {
	if m != nil {
		return m.Type
	}
	return TaskSchedule_ScheduleInfo_UNKNOWN
}

func (m *TaskSchedule_ScheduleInfo) GetExactOnly() bool {
	if m != nil {
		return m.ExactOnly
	}
	return false
}

func (m *TaskSchedule_ScheduleInfo) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *TaskSchedule_ScheduleInfo) GetWeekday() uint32 {
	if m != nil {
		return m.Weekday
	}
	return 0
}

// Instance of a task, tied to an owner (group or user)
type TaskInstance struct {
	Id     uint32     `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	TaskId uint32     `protobuf:"varint,2,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	Owner  *OwnerInfo `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
}

func (m *TaskInstance) Reset()                    { *m = TaskInstance{} }
func (m *TaskInstance) String() string            { return proto.CompactTextString(m) }
func (*TaskInstance) ProtoMessage()               {}
func (*TaskInstance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TaskInstance) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *TaskInstance) GetTaskId() uint32 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *TaskInstance) GetOwner() *OwnerInfo {
	if m != nil {
		return m.Owner
	}
	return nil
}

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutRequest struct {
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type LogoutResponse struct {
}

func (m *LogoutResponse) Reset()                    { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string            { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()               {}
func (*LogoutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// This message is mainly for debugging purposes.
type CheckLoginRequest struct {
}

func (m *CheckLoginRequest) Reset()                    { *m = CheckLoginRequest{} }
func (m *CheckLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckLoginRequest) ProtoMessage()               {}
func (*CheckLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type CheckLoginResponse struct {
}

func (m *CheckLoginResponse) Reset()                    { *m = CheckLoginResponse{} }
func (m *CheckLoginResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckLoginResponse) ProtoMessage()               {}
func (*CheckLoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type AddTaskDefRequest struct {
}

func (m *AddTaskDefRequest) Reset()                    { *m = AddTaskDefRequest{} }
func (m *AddTaskDefRequest) String() string            { return proto.CompactTextString(m) }
func (*AddTaskDefRequest) ProtoMessage()               {}
func (*AddTaskDefRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type AddTaskDefResponse struct {
}

func (m *AddTaskDefResponse) Reset()                    { *m = AddTaskDefResponse{} }
func (m *AddTaskDefResponse) String() string            { return proto.CompactTextString(m) }
func (*AddTaskDefResponse) ProtoMessage()               {}
func (*AddTaskDefResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type DelTaskDefRequest struct {
}

func (m *DelTaskDefRequest) Reset()                    { *m = DelTaskDefRequest{} }
func (m *DelTaskDefRequest) String() string            { return proto.CompactTextString(m) }
func (*DelTaskDefRequest) ProtoMessage()               {}
func (*DelTaskDefRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type DelTaskDefResponse struct {
}

func (m *DelTaskDefResponse) Reset()                    { *m = DelTaskDefResponse{} }
func (m *DelTaskDefResponse) String() string            { return proto.CompactTextString(m) }
func (*DelTaskDefResponse) ProtoMessage()               {}
func (*DelTaskDefResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*OwnerInfo)(nil), "nbproto.OwnerInfo")
	proto.RegisterType((*TaskDefinition)(nil), "nbproto.TaskDefinition")
	proto.RegisterType((*TaskSchedule)(nil), "nbproto.TaskSchedule")
	proto.RegisterType((*TaskSchedule_ScheduleInfo)(nil), "nbproto.TaskSchedule.ScheduleInfo")
	proto.RegisterType((*TaskInstance)(nil), "nbproto.TaskInstance")
	proto.RegisterType((*LoginRequest)(nil), "nbproto.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "nbproto.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "nbproto.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "nbproto.LogoutResponse")
	proto.RegisterType((*CheckLoginRequest)(nil), "nbproto.CheckLoginRequest")
	proto.RegisterType((*CheckLoginResponse)(nil), "nbproto.CheckLoginResponse")
	proto.RegisterType((*AddTaskDefRequest)(nil), "nbproto.AddTaskDefRequest")
	proto.RegisterType((*AddTaskDefResponse)(nil), "nbproto.AddTaskDefResponse")
	proto.RegisterType((*DelTaskDefRequest)(nil), "nbproto.DelTaskDefRequest")
	proto.RegisterType((*DelTaskDefResponse)(nil), "nbproto.DelTaskDefResponse")
	proto.RegisterEnum("nbproto.OwnerInfo_OwnerType", OwnerInfo_OwnerType_name, OwnerInfo_OwnerType_value)
	proto.RegisterEnum("nbproto.TaskSchedule_ScheduleInfo_ScheduleType", TaskSchedule_ScheduleInfo_ScheduleType_name, TaskSchedule_ScheduleInfo_ScheduleType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Nagbot service

type NagbotClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	AddTaskDef(ctx context.Context, in *AddTaskDefRequest, opts ...grpc.CallOption) (*AddTaskDefResponse, error)
	DelTaskDef(ctx context.Context, in *DelTaskDefRequest, opts ...grpc.CallOption) (*DelTaskDefResponse, error)
	CheckLogin(ctx context.Context, in *CheckLoginRequest, opts ...grpc.CallOption) (*CheckLoginResponse, error)
}

type nagbotClient struct {
	cc *grpc.ClientConn
}

func NewNagbotClient(cc *grpc.ClientConn) NagbotClient {
	return &nagbotClient{cc}
}

func (c *nagbotClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/nbproto.Nagbot/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nagbotClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := grpc.Invoke(ctx, "/nbproto.Nagbot/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nagbotClient) AddTaskDef(ctx context.Context, in *AddTaskDefRequest, opts ...grpc.CallOption) (*AddTaskDefResponse, error) {
	out := new(AddTaskDefResponse)
	err := grpc.Invoke(ctx, "/nbproto.Nagbot/AddTaskDef", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nagbotClient) DelTaskDef(ctx context.Context, in *DelTaskDefRequest, opts ...grpc.CallOption) (*DelTaskDefResponse, error) {
	out := new(DelTaskDefResponse)
	err := grpc.Invoke(ctx, "/nbproto.Nagbot/DelTaskDef", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nagbotClient) CheckLogin(ctx context.Context, in *CheckLoginRequest, opts ...grpc.CallOption) (*CheckLoginResponse, error) {
	out := new(CheckLoginResponse)
	err := grpc.Invoke(ctx, "/nbproto.Nagbot/CheckLogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Nagbot service

type NagbotServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	AddTaskDef(context.Context, *AddTaskDefRequest) (*AddTaskDefResponse, error)
	DelTaskDef(context.Context, *DelTaskDefRequest) (*DelTaskDefResponse, error)
	CheckLogin(context.Context, *CheckLoginRequest) (*CheckLoginResponse, error)
}

func RegisterNagbotServer(s *grpc.Server, srv NagbotServer) {
	s.RegisterService(&_Nagbot_serviceDesc, srv)
}

func _Nagbot_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NagbotServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbproto.Nagbot/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NagbotServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nagbot_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NagbotServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbproto.Nagbot/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NagbotServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nagbot_AddTaskDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTaskDefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NagbotServer).AddTaskDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbproto.Nagbot/AddTaskDef",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NagbotServer).AddTaskDef(ctx, req.(*AddTaskDefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nagbot_DelTaskDef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelTaskDefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NagbotServer).DelTaskDef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbproto.Nagbot/DelTaskDef",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NagbotServer).DelTaskDef(ctx, req.(*DelTaskDefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nagbot_CheckLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NagbotServer).CheckLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nbproto.Nagbot/CheckLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NagbotServer).CheckLogin(ctx, req.(*CheckLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Nagbot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nbproto.Nagbot",
	HandlerType: (*NagbotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Nagbot_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Nagbot_Logout_Handler,
		},
		{
			MethodName: "AddTaskDef",
			Handler:    _Nagbot_AddTaskDef_Handler,
		},
		{
			MethodName: "DelTaskDef",
			Handler:    _Nagbot_DelTaskDef_Handler,
		},
		{
			MethodName: "CheckLogin",
			Handler:    _Nagbot_CheckLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nbproto/nagbot.proto",
}

func init() { proto.RegisterFile("nbproto/nagbot.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 649 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xad, 0xdd, 0xc4, 0xb1, 0x6f, 0xe3, 0x7c, 0xee, 0xfd, 0x0a, 0x35, 0x29, 0x48, 0x91, 0x25,
	0xa4, 0x6c, 0x48, 0x51, 0xd9, 0x20, 0x21, 0x21, 0x45, 0x4d, 0x68, 0xa3, 0x06, 0x1b, 0x4d, 0x13,
	0xaa, 0x6e, 0x88, 0xdc, 0x78, 0xda, 0x5a, 0x49, 0xc7, 0x21, 0x9e, 0xd0, 0x46, 0xe2, 0x29, 0x78,
	0x3c, 0x1e, 0x82, 0x67, 0x40, 0x1e, 0xbb, 0xf6, 0x80, 0x2b, 0xc1, 0x86, 0xdd, 0x3d, 0x67, 0xce,
	0x3d, 0xf7, 0xcf, 0x09, 0xec, 0xb0, 0x8b, 0xc5, 0x32, 0xe2, 0xd1, 0x3e, 0xf3, 0xaf, 0x2e, 0x22,
	0xde, 0x11, 0x00, 0x6b, 0x19, 0xeb, 0x7c, 0x05, 0xc3, 0xbb, 0x65, 0x74, 0x39, 0x60, 0x97, 0x11,
	0x36, 0x40, 0x0d, 0x03, 0x5b, 0x69, 0x29, 0x6d, 0x93, 0xa8, 0x61, 0x80, 0x2f, 0xa1, 0xc2, 0xd7,
	0x0b, 0x6a, 0xab, 0x2d, 0xa5, 0xdd, 0x38, 0x78, 0xda, 0xc9, 0x92, 0x3a, 0x79, 0x46, 0x1a, 0x8d,
	0xd6, 0x0b, 0x4a, 0x84, 0xd2, 0x79, 0x91, 0xd9, 0x25, 0x14, 0x6e, 0x41, 0x6d, 0xec, 0x9e, 0xb8,
	0xde, 0x99, 0x6b, 0x6d, 0xa0, 0x0e, 0x95, 0xf1, 0x69, 0x9f, 0x58, 0x0a, 0x1a, 0x50, 0x3d, 0x22,
	0xde, 0xf8, 0x83, 0xa5, 0x3a, 0x9f, 0xa0, 0x31, 0xf2, 0xe3, 0x59, 0x8f, 0x5e, 0x86, 0x2c, 0xe4,
	0x61, 0xc4, 0x4a, 0x2d, 0x20, 0x54, 0x02, 0x1a, 0x4f, 0x45, 0x0b, 0x06, 0x11, 0x31, 0xb6, 0xa1,
	0x1a, 0x25, 0x45, 0xec, 0xcd, 0x96, 0xd2, 0xde, 0x3a, 0xc0, 0x72, 0x5f, 0x24, 0x15, 0x38, 0x3f,
	0x36, 0xa1, 0x9e, 0x14, 0x38, 0x9d, 0x5e, 0xd3, 0x60, 0x35, 0xa7, 0x25, 0xfb, 0x5d, 0xa8, 0x71,
	0x3f, 0x9e, 0x4d, 0xc2, 0x40, 0x54, 0x30, 0x89, 0x96, 0xc0, 0x41, 0xf0, 0xf7, 0x35, 0xf0, 0x2d,
	0xe8, 0x71, 0x66, 0x6f, 0x57, 0x84, 0xd8, 0xc9, 0xc5, 0x72, 0xed, 0xce, 0x7d, 0x20, 0x92, 0xf3,
	0x1c, 0x7c, 0x02, 0x3a, 0xa3, 0x77, 0x7c, 0x12, 0xac, 0xa8, 0x5d, 0x6d, 0x29, 0xed, 0x0a, 0xa9,
	0x25, 0xb8, 0xb7, 0xa2, 0xb8, 0x07, 0x46, 0x18, 0x4f, 0xfc, 0x29, 0x0f, 0xbf, 0x50, 0x5b, 0x6b,
	0x29, 0x6d, 0x9d, 0xe8, 0x61, 0xdc, 0x15, 0xb8, 0xf9, 0x4d, 0x85, 0xba, 0x6c, 0x89, 0x87, 0xd9,
	0xb5, 0x14, 0x71, 0xad, 0xfd, 0x3f, 0x37, 0x91, 0x83, 0xe2, 0x80, 0xf8, 0x0c, 0x80, 0xde, 0xf9,
	0x53, 0x3e, 0x89, 0xd8, 0x7c, 0x2d, 0x76, 0xa2, 0x13, 0x43, 0x30, 0x1e, 0x9b, 0xaf, 0x93, 0x73,
	0xf0, 0xf0, 0x86, 0x8a, 0xad, 0x98, 0x44, 0xc4, 0x68, 0x43, 0xed, 0x96, 0xd2, 0x59, 0xe0, 0xaf,
	0xc5, 0xfc, 0x26, 0xb9, 0x87, 0xce, 0xa2, 0xe8, 0xb0, 0xfc, 0x41, 0x6c, 0x41, 0xcd, 0x73, 0xfb,
	0xa7, 0xc7, 0xde, 0xc8, 0x52, 0xb0, 0x0e, 0xfa, 0xc0, 0x1d, 0xf5, 0xc9, 0xc7, 0xee, 0xd0, 0x52,
	0x11, 0x40, 0x3b, 0xeb, 0xf7, 0x4f, 0x86, 0xe7, 0xd6, 0x26, 0x9a, 0x60, 0xbc, 0xf7, 0xdc, 0xd1,
	0xf1, 0xa4, 0xd7, 0x3d, 0xb7, 0x2a, 0xb8, 0x0d, 0x66, 0x0a, 0x13, 0x41, 0x42, 0x55, 0x13, 0x75,
	0xd7, 0x75, 0xc7, 0xdd, 0xa1, 0xa5, 0x39, 0x7e, 0x7a, 0xef, 0x01, 0x8b, 0xb9, 0xcf, 0xa6, 0xff,
	0xe2, 0xde, 0xce, 0x3b, 0xa8, 0x0f, 0xa3, 0xab, 0x90, 0x11, 0xfa, 0x79, 0x45, 0x63, 0x8e, 0x4d,
	0xd0, 0x57, 0x31, 0x5d, 0x32, 0xff, 0x26, 0x5d, 0xbd, 0x41, 0x72, 0x9c, 0xbc, 0x2d, 0xfc, 0x38,
	0xbe, 0x8d, 0x96, 0x41, 0xf6, 0x05, 0xe7, 0xd8, 0x79, 0x0e, 0x66, 0xe6, 0x13, 0x2f, 0x22, 0x16,
	0x53, 0xdc, 0x81, 0x2a, 0x8f, 0x66, 0x94, 0x65, 0x2e, 0x29, 0x70, 0xfe, 0x13, 0xb2, 0x68, 0xc5,
	0xb3, 0x7a, 0x8e, 0x05, 0x8d, 0x7b, 0x22, 0x4d, 0x74, 0xfe, 0x87, 0xed, 0xc3, 0x6b, 0x3a, 0x9d,
	0xc9, 0x6d, 0x39, 0x3b, 0x80, 0x32, 0x59, 0x48, 0xbb, 0x41, 0x90, 0xfd, 0xe6, 0x24, 0xa9, 0x4c,
	0x16, 0xd2, 0x1e, 0x9d, 0x97, 0xa5, 0x32, 0x99, 0x4a, 0x0f, 0xbe, 0xab, 0xa0, 0xb9, 0xe2, 0xef,
	0x05, 0x5f, 0x43, 0x55, 0x54, 0xc4, 0x47, 0xf9, 0x06, 0xe5, 0xb6, 0x9a, 0x8f, 0x7f, 0xa7, 0xb3,
	0x6a, 0x1b, 0xf8, 0x06, 0xb4, 0x74, 0x2e, 0xfc, 0x45, 0x53, 0x4c, 0xde, 0xdc, 0x2d, 0xf1, 0x79,
	0xf2, 0x11, 0x40, 0x31, 0x02, 0x36, 0x73, 0x61, 0x69, 0xd8, 0xe6, 0xde, 0x83, 0x6f, 0xb2, 0x51,
	0x31, 0xa0, 0x64, 0x54, 0x5a, 0x85, 0x64, 0x54, 0xde, 0x48, 0x6a, 0x54, 0xec, 0x5f, 0x32, 0x2a,
	0x5d, 0x4a, 0x32, 0x7a, 0xe0, 0x60, 0x1b, 0x17, 0x9a, 0x78, 0x7b, 0xf5, 0x33, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0xc0, 0x1d, 0xd0, 0xc9, 0x05, 0x00, 0x00,
}
