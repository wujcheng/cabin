// Code generated by protoc-gen-go. DO NOT EDIT.
// source: main.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	main.proto

It has these top-level messages:
	UpdateBrokerRequest
	ListBrokerRequest
	Broker
	ListBrokerResponse
	PublishRequest
	RetResponse
	User
	AddUserRequest
	DelUserRequest
	GetUserRequest
	AddMembersRequest
	DelMembersRequest
	IsMemberRequest
	ListMembersRequest
	ListMembersResponse
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/mwitkow/go-proto-validators"

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

type UpdateBrokerRequest struct {
	External string `protobuf:"bytes,1,opt,name=external" json:"external,omitempty"`
	Internal string `protobuf:"bytes,2,opt,name=internal" json:"internal,omitempty"`
	Conn     uint64 `protobuf:"varint,3,opt,name=conn" json:"conn,omitempty"`
	MaxConn  int64  `protobuf:"zigzag64,4,opt,name=max_conn,json=maxConn" json:"max_conn,omitempty"`
}

func (m *UpdateBrokerRequest) Reset()                    { *m = UpdateBrokerRequest{} }
func (m *UpdateBrokerRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateBrokerRequest) ProtoMessage()               {}
func (*UpdateBrokerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpdateBrokerRequest) GetExternal() string {
	if m != nil {
		return m.External
	}
	return ""
}

func (m *UpdateBrokerRequest) GetInternal() string {
	if m != nil {
		return m.Internal
	}
	return ""
}

func (m *UpdateBrokerRequest) GetConn() uint64 {
	if m != nil {
		return m.Conn
	}
	return 0
}

func (m *UpdateBrokerRequest) GetMaxConn() int64 {
	if m != nil {
		return m.MaxConn
	}
	return 0
}

type ListBrokerRequest struct {
}

func (m *ListBrokerRequest) Reset()                    { *m = ListBrokerRequest{} }
func (m *ListBrokerRequest) String() string            { return proto.CompactTextString(m) }
func (*ListBrokerRequest) ProtoMessage()               {}
func (*ListBrokerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Broker struct {
	External   string `protobuf:"bytes,1,opt,name=external" json:"external,omitempty"`
	Internal   string `protobuf:"bytes,2,opt,name=internal" json:"internal,omitempty"`
	Conn       uint64 `protobuf:"varint,3,opt,name=conn" json:"conn,omitempty"`
	MaxConn    int64  `protobuf:"zigzag64,4,opt,name=max_conn,json=maxConn" json:"max_conn,omitempty"`
	UpdateTime int64  `protobuf:"varint,5,opt,name=update_time,json=updateTime" json:"update_time,omitempty"`
}

func (m *Broker) Reset()                    { *m = Broker{} }
func (m *Broker) String() string            { return proto.CompactTextString(m) }
func (*Broker) ProtoMessage()               {}
func (*Broker) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Broker) GetExternal() string {
	if m != nil {
		return m.External
	}
	return ""
}

func (m *Broker) GetInternal() string {
	if m != nil {
		return m.Internal
	}
	return ""
}

func (m *Broker) GetConn() uint64 {
	if m != nil {
		return m.Conn
	}
	return 0
}

func (m *Broker) GetMaxConn() int64 {
	if m != nil {
		return m.MaxConn
	}
	return 0
}

func (m *Broker) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

type ListBrokerResponse struct {
	Brokers []*Broker `protobuf:"bytes,1,rep,name=brokers" json:"brokers,omitempty"`
}

func (m *ListBrokerResponse) Reset()                    { *m = ListBrokerResponse{} }
func (m *ListBrokerResponse) String() string            { return proto.CompactTextString(m) }
func (*ListBrokerResponse) ProtoMessage()               {}
func (*ListBrokerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ListBrokerResponse) GetBrokers() []*Broker {
	if m != nil {
		return m.Brokers
	}
	return nil
}

type PublishRequest struct {
	Topic    string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	Payload  string `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	Qos      int32  `protobuf:"zigzag32,3,opt,name=qos" json:"qos,omitempty"`
	Retained bool   `protobuf:"varint,4,opt,name=retained" json:"retained,omitempty"`
}

func (m *PublishRequest) Reset()                    { *m = PublishRequest{} }
func (m *PublishRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()               {}
func (*PublishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PublishRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *PublishRequest) GetQos() int32 {
	if m != nil {
		return m.Qos
	}
	return 0
}

func (m *PublishRequest) GetRetained() bool {
	if m != nil {
		return m.Retained
	}
	return false
}

type RetResponse struct {
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (m *RetResponse) Reset()                    { *m = RetResponse{} }
func (m *RetResponse) String() string            { return proto.CompactTextString(m) }
func (*RetResponse) ProtoMessage()               {}
func (*RetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RetResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type User struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Ip       string `protobuf:"bytes,4,opt,name=ip" json:"ip,omitempty"`
	Broker   string `protobuf:"bytes,5,opt,name=broker" json:"broker,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *User) GetBroker() string {
	if m != nil {
		return m.Broker
	}
	return ""
}

type AddUserRequest struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Ip       string `protobuf:"bytes,4,opt,name=ip" json:"ip,omitempty"`
}

func (m *AddUserRequest) Reset()                    { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string            { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()               {}
func (*AddUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AddUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AddUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AddUserRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type DelUserRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *DelUserRequest) Reset()                    { *m = DelUserRequest{} }
func (m *DelUserRequest) String() string            { return proto.CompactTextString(m) }
func (*DelUserRequest) ProtoMessage()               {}
func (*DelUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DelUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GetUserRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *GetUserRequest) Reset()                    { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()               {}
func (*GetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type AddMembersRequest struct {
	GroupId string   `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	Members []string `protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
}

func (m *AddMembersRequest) Reset()                    { *m = AddMembersRequest{} }
func (m *AddMembersRequest) String() string            { return proto.CompactTextString(m) }
func (*AddMembersRequest) ProtoMessage()               {}
func (*AddMembersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AddMembersRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *AddMembersRequest) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type DelMembersRequest struct {
	GroupId string   `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	Members []string `protobuf:"bytes,2,rep,name=members" json:"members,omitempty"`
}

func (m *DelMembersRequest) Reset()                    { *m = DelMembersRequest{} }
func (m *DelMembersRequest) String() string            { return proto.CompactTextString(m) }
func (*DelMembersRequest) ProtoMessage()               {}
func (*DelMembersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DelMembersRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *DelMembersRequest) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

type IsMemberRequest struct {
	GroupId  string `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
	MemberId string `protobuf:"bytes,2,opt,name=member_id,json=memberId" json:"member_id,omitempty"`
}

func (m *IsMemberRequest) Reset()                    { *m = IsMemberRequest{} }
func (m *IsMemberRequest) String() string            { return proto.CompactTextString(m) }
func (*IsMemberRequest) ProtoMessage()               {}
func (*IsMemberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *IsMemberRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *IsMemberRequest) GetMemberId() string {
	if m != nil {
		return m.MemberId
	}
	return ""
}

type ListMembersRequest struct {
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId" json:"group_id,omitempty"`
}

func (m *ListMembersRequest) Reset()                    { *m = ListMembersRequest{} }
func (m *ListMembersRequest) String() string            { return proto.CompactTextString(m) }
func (*ListMembersRequest) ProtoMessage()               {}
func (*ListMembersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ListMembersRequest) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

type ListMembersResponse struct {
	Members []string `protobuf:"bytes,1,rep,name=members" json:"members,omitempty"`
}

func (m *ListMembersResponse) Reset()                    { *m = ListMembersResponse{} }
func (m *ListMembersResponse) String() string            { return proto.CompactTextString(m) }
func (*ListMembersResponse) ProtoMessage()               {}
func (*ListMembersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ListMembersResponse) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateBrokerRequest)(nil), "main.UpdateBrokerRequest")
	proto.RegisterType((*ListBrokerRequest)(nil), "main.ListBrokerRequest")
	proto.RegisterType((*Broker)(nil), "main.Broker")
	proto.RegisterType((*ListBrokerResponse)(nil), "main.ListBrokerResponse")
	proto.RegisterType((*PublishRequest)(nil), "main.PublishRequest")
	proto.RegisterType((*RetResponse)(nil), "main.RetResponse")
	proto.RegisterType((*User)(nil), "main.User")
	proto.RegisterType((*AddUserRequest)(nil), "main.AddUserRequest")
	proto.RegisterType((*DelUserRequest)(nil), "main.DelUserRequest")
	proto.RegisterType((*GetUserRequest)(nil), "main.GetUserRequest")
	proto.RegisterType((*AddMembersRequest)(nil), "main.AddMembersRequest")
	proto.RegisterType((*DelMembersRequest)(nil), "main.DelMembersRequest")
	proto.RegisterType((*IsMemberRequest)(nil), "main.IsMemberRequest")
	proto.RegisterType((*ListMembersRequest)(nil), "main.ListMembersRequest")
	proto.RegisterType((*ListMembersResponse)(nil), "main.ListMembersResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BrokerSvc service

type BrokerSvcClient interface {
	Update(ctx context.Context, in *UpdateBrokerRequest, opts ...grpc.CallOption) (*Broker, error)
	List(ctx context.Context, in *ListBrokerRequest, opts ...grpc.CallOption) (*ListBrokerResponse, error)
}

type brokerSvcClient struct {
	cc *grpc.ClientConn
}

func NewBrokerSvcClient(cc *grpc.ClientConn) BrokerSvcClient {
	return &brokerSvcClient{cc}
}

func (c *brokerSvcClient) Update(ctx context.Context, in *UpdateBrokerRequest, opts ...grpc.CallOption) (*Broker, error) {
	out := new(Broker)
	err := grpc.Invoke(ctx, "/main.BrokerSvc/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brokerSvcClient) List(ctx context.Context, in *ListBrokerRequest, opts ...grpc.CallOption) (*ListBrokerResponse, error) {
	out := new(ListBrokerResponse)
	err := grpc.Invoke(ctx, "/main.BrokerSvc/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BrokerSvc service

type BrokerSvcServer interface {
	Update(context.Context, *UpdateBrokerRequest) (*Broker, error)
	List(context.Context, *ListBrokerRequest) (*ListBrokerResponse, error)
}

func RegisterBrokerSvcServer(s *grpc.Server, srv BrokerSvcServer) {
	s.RegisterService(&_BrokerSvc_serviceDesc, srv)
}

func _BrokerSvc_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBrokerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerSvcServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.BrokerSvc/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerSvcServer).Update(ctx, req.(*UpdateBrokerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BrokerSvc_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBrokerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerSvcServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.BrokerSvc/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerSvcServer).List(ctx, req.(*ListBrokerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BrokerSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.BrokerSvc",
	HandlerType: (*BrokerSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _BrokerSvc_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _BrokerSvc_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}

// Client API for IMSvc service

type IMSvcClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*RetResponse, error)
}

type iMSvcClient struct {
	cc *grpc.ClientConn
}

func NewIMSvcClient(cc *grpc.ClientConn) IMSvcClient {
	return &iMSvcClient{cc}
}

func (c *iMSvcClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*RetResponse, error) {
	out := new(RetResponse)
	err := grpc.Invoke(ctx, "/main.IMSvc/Publish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IMSvc service

type IMSvcServer interface {
	Publish(context.Context, *PublishRequest) (*RetResponse, error)
}

func RegisterIMSvcServer(s *grpc.Server, srv IMSvcServer) {
	s.RegisterService(&_IMSvc_serviceDesc, srv)
}

func _IMSvc_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMSvcServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.IMSvc/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMSvcServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IMSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.IMSvc",
	HandlerType: (*IMSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _IMSvc_Publish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}

// Client API for UserSvc service

type UserSvcClient interface {
	Add(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*User, error)
	Del(ctx context.Context, in *DelUserRequest, opts ...grpc.CallOption) (*User, error)
	Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
}

type userSvcClient struct {
	cc *grpc.ClientConn
}

func NewUserSvcClient(cc *grpc.ClientConn) UserSvcClient {
	return &userSvcClient{cc}
}

func (c *userSvcClient) Add(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/main.UserSvc/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) Del(ctx context.Context, in *DelUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/main.UserSvc/Del", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/main.UserSvc/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserSvc service

type UserSvcServer interface {
	Add(context.Context, *AddUserRequest) (*User, error)
	Del(context.Context, *DelUserRequest) (*User, error)
	Get(context.Context, *GetUserRequest) (*User, error)
}

func RegisterUserSvcServer(s *grpc.Server, srv UserSvcServer) {
	s.RegisterService(&_UserSvc_serviceDesc, srv)
}

func _UserSvc_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.UserSvc/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).Add(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_Del_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).Del(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.UserSvc/Del",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).Del(ctx, req.(*DelUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.UserSvc/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).Get(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.UserSvc",
	HandlerType: (*UserSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _UserSvc_Add_Handler,
		},
		{
			MethodName: "Del",
			Handler:    _UserSvc_Del_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserSvc_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}

// Client API for GroupSvc service

type GroupSvcClient interface {
	AddMembers(ctx context.Context, in *AddMembersRequest, opts ...grpc.CallOption) (*RetResponse, error)
	DelMembers(ctx context.Context, in *DelMembersRequest, opts ...grpc.CallOption) (*RetResponse, error)
	IsMember(ctx context.Context, in *IsMemberRequest, opts ...grpc.CallOption) (*RetResponse, error)
	ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error)
}

type groupSvcClient struct {
	cc *grpc.ClientConn
}

func NewGroupSvcClient(cc *grpc.ClientConn) GroupSvcClient {
	return &groupSvcClient{cc}
}

func (c *groupSvcClient) AddMembers(ctx context.Context, in *AddMembersRequest, opts ...grpc.CallOption) (*RetResponse, error) {
	out := new(RetResponse)
	err := grpc.Invoke(ctx, "/main.GroupSvc/AddMembers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupSvcClient) DelMembers(ctx context.Context, in *DelMembersRequest, opts ...grpc.CallOption) (*RetResponse, error) {
	out := new(RetResponse)
	err := grpc.Invoke(ctx, "/main.GroupSvc/DelMembers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupSvcClient) IsMember(ctx context.Context, in *IsMemberRequest, opts ...grpc.CallOption) (*RetResponse, error) {
	out := new(RetResponse)
	err := grpc.Invoke(ctx, "/main.GroupSvc/IsMember", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupSvcClient) ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error) {
	out := new(ListMembersResponse)
	err := grpc.Invoke(ctx, "/main.GroupSvc/ListMembers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GroupSvc service

type GroupSvcServer interface {
	AddMembers(context.Context, *AddMembersRequest) (*RetResponse, error)
	DelMembers(context.Context, *DelMembersRequest) (*RetResponse, error)
	IsMember(context.Context, *IsMemberRequest) (*RetResponse, error)
	ListMembers(context.Context, *ListMembersRequest) (*ListMembersResponse, error)
}

func RegisterGroupSvcServer(s *grpc.Server, srv GroupSvcServer) {
	s.RegisterService(&_GroupSvc_serviceDesc, srv)
}

func _GroupSvc_AddMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupSvcServer).AddMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.GroupSvc/AddMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupSvcServer).AddMembers(ctx, req.(*AddMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupSvc_DelMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupSvcServer).DelMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.GroupSvc/DelMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupSvcServer).DelMembers(ctx, req.(*DelMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupSvc_IsMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupSvcServer).IsMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.GroupSvc/IsMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupSvcServer).IsMember(ctx, req.(*IsMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupSvc_ListMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupSvcServer).ListMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/main.GroupSvc/ListMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupSvcServer).ListMembers(ctx, req.(*ListMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GroupSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "main.GroupSvc",
	HandlerType: (*GroupSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMembers",
			Handler:    _GroupSvc_AddMembers_Handler,
		},
		{
			MethodName: "DelMembers",
			Handler:    _GroupSvc_DelMembers_Handler,
		},
		{
			MethodName: "IsMember",
			Handler:    _GroupSvc_IsMember_Handler,
		},
		{
			MethodName: "ListMembers",
			Handler:    _GroupSvc_ListMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}

func init() { proto.RegisterFile("main.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 870 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x96, 0xdf, 0x6e, 0x1b, 0x45,
	0x14, 0xc6, 0x35, 0x6b, 0xc7, 0x7f, 0x8e, 0x8b, 0x5b, 0x4f, 0xd2, 0x74, 0xb3, 0x0a, 0xc4, 0x4c,
	0x25, 0x64, 0x55, 0x34, 0x2b, 0x82, 0x04, 0x52, 0xc5, 0x45, 0x13, 0x22, 0x45, 0x41, 0x54, 0xa0,
	0x6d, 0x23, 0x35, 0x57, 0x61, 0x9d, 0x19, 0x39, 0xa3, 0xec, 0xee, 0x6c, 0x77, 0xc6, 0x49, 0xb8,
	0xe1, 0x82, 0x6b, 0x2e, 0x40, 0x3c, 0x0a, 0x0f, 0xc2, 0x05, 0x0f, 0x50, 0xa9, 0xe2, 0x31, 0x90,
	0x40, 0x33, 0xb3, 0x63, 0x7b, 0xd7, 0x96, 0x4a, 0x10, 0xf5, 0xd5, 0xee, 0x99, 0x33, 0xbf, 0xf9,
	0xce, 0xd9, 0xf3, 0x8d, 0x0c, 0x90, 0xc6, 0x3c, 0xdb, 0xcd, 0x0b, 0xa1, 0x04, 0x6e, 0xea, 0xe7,
	0x60, 0x7b, 0x22, 0xc4, 0x24, 0x61, 0x61, 0x9c, 0xf3, 0x30, 0xce, 0x32, 0xa1, 0x62, 0xc5, 0x45,
	0x26, 0x6d, 0x4e, 0xf0, 0xd9, 0x84, 0xab, 0x8b, 0xe9, 0x78, 0xf7, 0x5c, 0xa4, 0x61, 0x7a, 0xcd,
	0xd5, 0xa5, 0xb8, 0x0e, 0x27, 0xe2, 0xb1, 0x59, 0x7c, 0x7c, 0x15, 0x27, 0x9c, 0xc6, 0x4a, 0x14,
	0x32, 0x9c, 0x3d, 0xda, 0x7d, 0xe4, 0x27, 0x04, 0xeb, 0x27, 0x39, 0x8d, 0x15, 0x3b, 0x28, 0xc4,
	0x25, 0x2b, 0x22, 0xf6, 0x6a, 0xca, 0xa4, 0xc2, 0x04, 0x3a, 0xec, 0x46, 0xb1, 0x22, 0x8b, 0x13,
	0x1f, 0x0d, 0xd1, 0xa8, 0x7b, 0xd0, 0x7a, 0xf3, 0x7a, 0xc7, 0x7b, 0x89, 0xa2, 0x59, 0x5c, 0xe7,
	0xf0, 0xac, 0xcc, 0xf1, 0xaa, 0x39, 0x2e, 0x8e, 0x31, 0x34, 0xcf, 0x45, 0x96, 0xf9, 0x8d, 0x21,
	0x1a, 0x35, 0x23, 0xf3, 0x8c, 0xb7, 0xa0, 0x93, 0xc6, 0x37, 0x67, 0x26, 0xde, 0x1c, 0xa2, 0x11,
	0x8e, 0xda, 0x69, 0x7c, 0xf3, 0xa5, 0xc8, 0x32, 0xb2, 0x0e, 0x83, 0xaf, 0xb9, 0x54, 0x15, 0x2d,
	0xe4, 0x67, 0x04, 0x2d, 0x1b, 0xc1, 0x41, 0x5d, 0xd6, 0x82, 0x9c, 0xa0, 0x2e, 0xe7, 0x3f, 0xcb,
	0xc0, 0x3b, 0xd0, 0x9b, 0x9a, 0xa6, 0x9c, 0x29, 0x9e, 0x32, 0x7f, 0x6d, 0x88, 0x46, 0x8d, 0x08,
	0x6c, 0xe8, 0x05, 0x4f, 0x19, 0xf9, 0x02, 0xf0, 0xa2, 0x4e, 0x99, 0x8b, 0x4c, 0x32, 0xfc, 0x11,
	0xb4, 0xc7, 0x26, 0x22, 0x7d, 0x34, 0x6c, 0x8c, 0x7a, 0x7b, 0x77, 0x76, 0xcd, 0x67, 0x2c, 0xd3,
	0xdc, 0x22, 0xf9, 0x05, 0x41, 0xff, 0xdb, 0xe9, 0x38, 0xe1, 0xf2, 0xc2, 0xf5, 0x7b, 0x1b, 0xd6,
	0x94, 0xc8, 0xf9, 0x79, 0xad, 0xd9, 0x36, 0x88, 0x87, 0xd0, 0xce, 0xe3, 0xef, 0x13, 0x11, 0xd3,
	0x5a, 0xa3, 0x5d, 0x18, 0x3f, 0x84, 0xc6, 0x2b, 0x21, 0x4d, 0x7d, 0x83, 0x83, 0xc1, 0x9b, 0xd7,
	0x3b, 0xef, 0xdd, 0xfb, 0xdb, 0xfd, 0x90, 0xdf, 0x88, 0xf4, 0xaa, 0xee, 0x50, 0xc1, 0x54, 0xcc,
	0x33, 0x46, 0x4d, 0xc5, 0x9d, 0x68, 0xf6, 0x4e, 0xde, 0x87, 0x5e, 0xc4, 0xd4, 0xac, 0x94, 0x3e,
	0x78, 0xe2, 0xd2, 0x88, 0xe9, 0x44, 0x9e, 0xb8, 0x24, 0x3f, 0x40, 0xf3, 0x44, 0xb2, 0x02, 0x6f,
	0x82, 0xc7, 0x69, 0x4d, 0xa4, 0xc7, 0xa9, 0x46, 0x4f, 0xa5, 0xee, 0x75, 0xca, 0x5c, 0xf3, 0xdd,
	0xbb, 0x5e, 0xcb, 0x63, 0x29, 0xaf, 0x45, 0x41, 0x8d, 0xc0, 0x6e, 0x34, 0x7b, 0xd7, 0xe7, 0xf0,
	0xdc, 0x88, 0xe9, 0x46, 0x1e, 0xcf, 0xf1, 0x26, 0xb4, 0x6c, 0x97, 0x4c, 0xd3, 0xbb, 0x51, 0xf9,
	0x46, 0x2e, 0xa0, 0xbf, 0x4f, 0xa9, 0x96, 0xe0, 0x3a, 0xd6, 0x9f, 0x2b, 0xf9, 0x3f, 0x15, 0x90,
	0x8f, 0xa1, 0x7f, 0xc8, 0x92, 0xc5, 0x93, 0x16, 0xc9, 0xa8, 0x4a, 0xd6, 0xd9, 0x47, 0x4c, 0xfd,
	0xdb, 0xec, 0x97, 0x30, 0xd8, 0xa7, 0xf4, 0x19, 0x4b, 0xc7, 0xac, 0x90, 0x6e, 0xc3, 0x87, 0xd0,
	0x99, 0x14, 0x62, 0x9a, 0x9f, 0x2d, 0x35, 0xb6, 0x6d, 0xe2, 0xc7, 0x54, 0x7f, 0xff, 0xd4, 0x6e,
	0xf2, 0xbd, 0x61, 0xc3, 0x65, 0x7c, 0x87, 0x22, 0x17, 0xd6, 0xe4, 0x43, 0x96, 0xbc, 0x0b, 0xf2,
	0x29, 0xdc, 0x3d, 0x96, 0x16, 0x7c, 0x0b, 0xee, 0x43, 0xe8, 0x5a, 0x80, 0xce, 0xa9, 0x5d, 0x0e,
	0x76, 0xe1, 0x98, 0x92, 0xcf, 0xad, 0x8b, 0x6e, 0xad, 0x9a, 0x84, 0xb0, 0x5e, 0xd9, 0x58, 0x0e,
	0xad, 0x3f, 0x2f, 0x46, 0xfb, 0xaf, 0x3b, 0x2b, 0x62, 0xef, 0x37, 0x04, 0x5d, 0xeb, 0xc2, 0xe7,
	0x57, 0xe7, 0xf8, 0x1b, 0x68, 0xd9, 0x3b, 0x0f, 0x6f, 0x59, 0x83, 0xae, 0xb8, 0x01, 0x83, 0x8a,
	0x77, 0xc9, 0xf6, 0x8f, 0x7f, 0xfc, 0xf9, 0xab, 0xb7, 0x49, 0x06, 0xe1, 0xd5, 0x27, 0xa1, 0x9d,
	0xc9, 0xd0, 0xde, 0x07, 0x4f, 0xd0, 0x23, 0x7c, 0x02, 0x4d, 0xad, 0x07, 0x3f, 0xb0, 0x7b, 0x96,
	0xae, 0xb0, 0xc0, 0x5f, 0x5e, 0xb0, 0x9a, 0x49, 0x60, 0xc0, 0x1b, 0xe4, 0xee, 0x02, 0x38, 0xe1,
	0x52, 0x3d, 0x41, 0x8f, 0xf6, 0x9e, 0xc3, 0xda, 0xf1, 0x33, 0x2d, 0xf8, 0x2b, 0x68, 0x97, 0xf7,
	0x05, 0xde, 0xb0, 0xa4, 0xea, 0xf5, 0x11, 0x0c, 0x6c, 0x74, 0xc1, 0xc1, 0x64, 0xd3, 0x80, 0xef,
	0x91, 0x9e, 0x06, 0xe7, 0x36, 0x5d, 0x43, 0x7f, 0x47, 0xd0, 0xd6, 0xf3, 0xaa, 0xb9, 0x4f, 0xa1,
	0xb1, 0x4f, 0xa9, 0x63, 0x56, 0x0d, 0x16, 0x40, 0xd9, 0x1b, 0xc9, 0x0a, 0xf2, 0xc0, 0xc0, 0x06,
	0xe4, 0x8e, 0x86, 0xe9, 0x71, 0x0e, 0x63, 0x4a, 0x75, 0xe5, 0x4f, 0xa1, 0x71, 0xc8, 0x12, 0x47,
	0xa8, 0x1a, 0xe7, 0x2d, 0x04, 0xca, 0x92, 0x92, 0x70, 0xc4, 0x94, 0x23, 0x54, 0xcd, 0xf4, 0x16,
	0xc2, 0x84, 0x99, 0x36, 0xfd, 0xe5, 0x41, 0xe7, 0x48, 0x4f, 0x86, 0x2e, 0xe9, 0x14, 0x60, 0x6e,
	0x31, 0xf7, 0x41, 0x96, 0x4c, 0xb7, 0xaa, 0x61, 0x3b, 0x86, 0xbf, 0x45, 0x36, 0x34, 0xdf, 0x4c,
	0x5a, 0x68, 0xc7, 0xc7, 0xd5, 0x7a, 0x0a, 0x30, 0xf7, 0x98, 0x43, 0x2f, 0xb9, 0xee, 0x76, 0xe8,
	0xb2, 0x09, 0x2f, 0xa0, 0xe3, 0x4c, 0x86, 0xef, 0xdb, 0xfd, 0x35, 0xd3, 0xad, 0xc2, 0x7e, 0x60,
	0xb0, 0x3e, 0x59, 0x9f, 0x63, 0xb9, 0x2c, 0xc9, 0x9a, 0x4a, 0xa1, 0xb7, 0x60, 0x13, 0xbc, 0x30,
	0x84, 0x35, 0xc9, 0x5b, 0x2b, 0x56, 0xca, 0x33, 0x86, 0xe6, 0x8c, 0x80, 0xdc, 0x5f, 0x92, 0x5e,
	0x4e, 0xe9, 0xb8, 0x65, 0xfe, 0x49, 0x7c, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x6c,
	0x2d, 0xe8, 0xb3, 0x08, 0x00, 0x00,
}
