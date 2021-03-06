// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/user.proto

package pb

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

type User struct {
	UserID               int64    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_45dc2bac12b48cd0, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Users struct {
	U                    []*User  `protobuf:"bytes,1,rep,name=u,proto3" json:"u,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_45dc2bac12b48cd0, []int{1}
}
func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (dst *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(dst, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetU() []*User {
	if m != nil {
		return m.U
	}
	return nil
}

type Result struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_45dc2bac12b48cd0, []int{2}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Result) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetAllUserRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllUserRequest) Reset()         { *m = GetAllUserRequest{} }
func (m *GetAllUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllUserRequest) ProtoMessage()    {}
func (*GetAllUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_45dc2bac12b48cd0, []int{3}
}
func (m *GetAllUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllUserRequest.Unmarshal(m, b)
}
func (m *GetAllUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllUserRequest.Marshal(b, m, deterministic)
}
func (dst *GetAllUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllUserRequest.Merge(dst, src)
}
func (m *GetAllUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllUserRequest.Size(m)
}
func (m *GetAllUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllUserRequest proto.InternalMessageInfo

type AddUserRequest struct {
	UserID               int64    `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserRequest) Reset()         { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()    {}
func (*AddUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_45dc2bac12b48cd0, []int{4}
}
func (m *AddUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserRequest.Unmarshal(m, b)
}
func (m *AddUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserRequest.Marshal(b, m, deterministic)
}
func (dst *AddUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserRequest.Merge(dst, src)
}
func (m *AddUserRequest) XXX_Size() int {
	return xxx_messageInfo_AddUserRequest.Size(m)
}
func (m *AddUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserRequest proto.InternalMessageInfo

func (m *AddUserRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *AddUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*Users)(nil), "pb.Users")
	proto.RegisterType((*Result)(nil), "pb.Result")
	proto.RegisterType((*GetAllUserRequest)(nil), "pb.GetAllUserRequest")
	proto.RegisterType((*AddUserRequest)(nil), "pb.AddUserRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	GetAllUser(ctx context.Context, in *GetAllUserRequest, opts ...grpc.CallOption) (*Users, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*Result, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetAllUser(ctx context.Context, in *GetAllUserRequest, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/pb.UserService/GetAllUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/pb.UserService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	GetAllUser(context.Context, *GetAllUserRequest) (*Users, error)
	AddUser(context.Context, *AddUserRequest) (*Result, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetAllUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAllUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/GetAllUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAllUser(ctx, req.(*GetAllUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllUser",
			Handler:    _UserService_GetAllUser_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _UserService_AddUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/user.proto",
}

func init() { proto.RegisterFile("pb/user.proto", fileDescriptor_user_45dc2bac12b48cd0) }

var fileDescriptor_user_45dc2bac12b48cd0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xbb, 0xad, 0xa6, 0xe9, 0x14, 0x05, 0x47, 0x2c, 0x21, 0x17, 0xc3, 0x9e, 0x02, 0x42,
	0x94, 0x7a, 0x13, 0x2f, 0x85, 0x82, 0x78, 0x5d, 0xf1, 0x03, 0x34, 0xe9, 0x20, 0xc2, 0xd6, 0xae,
	0x3b, 0xbb, 0x7e, 0x7e, 0x99, 0x98, 0xf5, 0x0f, 0xde, 0x7a, 0xdb, 0xf7, 0x86, 0xfd, 0xcd, 0x7b,
	0x03, 0x27, 0xae, 0xbd, 0x8e, 0x4c, 0xbe, 0x71, 0x7e, 0x1f, 0xf6, 0x38, 0x76, 0xad, 0xbe, 0x83,
	0xa3, 0x67, 0x26, 0x8f, 0x0b, 0xc8, 0x64, 0xf2, 0xb8, 0x2e, 0x54, 0xa5, 0xea, 0x89, 0x19, 0x14,
	0x96, 0x90, 0xcb, 0xeb, 0x6d, 0xb3, 0xa3, 0x62, 0x5c, 0xa9, 0x7a, 0x66, 0xbe, 0xb5, 0xbe, 0x84,
	0x63, 0xf9, 0xcb, 0xb8, 0x00, 0x15, 0x0b, 0x55, 0x4d, 0xea, 0xf9, 0x32, 0x6f, 0x5c, 0xdb, 0x88,
	0x6b, 0x54, 0xd4, 0xf7, 0x90, 0x19, 0xe2, 0x68, 0x03, 0x16, 0x30, 0xe5, 0xd8, 0x75, 0xc4, 0xdc,
	0xf3, 0x73, 0x93, 0xa4, 0x4c, 0x76, 0xc4, 0xbc, 0x79, 0x49, 0xfc, 0x24, 0xf5, 0x39, 0x9c, 0x3d,
	0x50, 0x58, 0x59, 0xdb, 0xe3, 0xe8, 0x3d, 0x12, 0x07, 0xbd, 0x86, 0xd3, 0xd5, 0x76, 0xfb, 0xcb,
	0x39, 0x24, 0xf9, 0xd2, 0xc2, 0x5c, 0x10, 0x4f, 0xe4, 0x3f, 0x5e, 0x3b, 0xc2, 0x1b, 0x80, 0x9f,
	0x4d, 0x78, 0x21, 0x15, 0xfe, 0x6d, 0x2e, 0x67, 0xa9, 0x19, 0xeb, 0x11, 0x5e, 0xc1, 0x74, 0x88,
	0x81, 0x28, 0xfe, 0xdf, 0x4c, 0x25, 0x88, 0xf7, 0x55, 0x5d, 0x8f, 0xda, 0xac, 0x3f, 0xf7, 0xed,
	0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0x25, 0xe0, 0x19, 0x7f, 0x01, 0x00, 0x00,
}
