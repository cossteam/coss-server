// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: api/v1/group_join_request.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GroupJoinRequestService_GetGroupJoinRequestListByUserId_FullMethodName       = "/v1.GroupJoinRequestService/GetGroupJoinRequestListByUserId"
	GroupJoinRequestService_GetGroupJoinRequestByGroupIdAndUserId_FullMethodName = "/v1.GroupJoinRequestService/GetGroupJoinRequestByGroupIdAndUserId"
	GroupJoinRequestService_JoinGroup_FullMethodName                             = "/v1.GroupJoinRequestService/JoinGroup"
	GroupJoinRequestService_InviteJoinGroup_FullMethodName                       = "/v1.GroupJoinRequestService/InviteJoinGroup"
	GroupJoinRequestService_ManageGroupJoinRequestByID_FullMethodName            = "/v1.GroupJoinRequestService/ManageGroupJoinRequestByID"
	GroupJoinRequestService_GetGroupJoinRequestByID_FullMethodName               = "/v1.GroupJoinRequestService/GetGroupJoinRequestByID"
)

// GroupJoinRequestServiceClient is the client API for GroupJoinRequestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupJoinRequestServiceClient interface {
	// 获取用户加入群聊列表
	GetGroupJoinRequestListByUserId(ctx context.Context, in *GetGroupJoinRequestListRequest, opts ...grpc.CallOption) (*GroupJoinRequestListResponse, error)
	// 根据群聊id和用户id获取群聊加入请求
	GetGroupJoinRequestByGroupIdAndUserId(ctx context.Context, in *GetGroupJoinRequestByGroupIdAndUserIdRequest, opts ...grpc.CallOption) (*GetGroupJoinRequestByGroupIdAndUserIdResponse, error)
	// 申请加入群聊
	JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error)
	// 邀请用户加入群聊
	InviteJoinGroup(ctx context.Context, in *InviteJoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error)
	// 用户管理入群邀请
	ManageGroupJoinRequestByID(ctx context.Context, in *ManageGroupJoinRequestByIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 根据请求id获取申请
	GetGroupJoinRequestByID(ctx context.Context, in *GetGroupJoinRequestByIDRequest, opts ...grpc.CallOption) (*GetGroupJoinRequestByIDResponse, error)
}

type groupJoinRequestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupJoinRequestServiceClient(cc grpc.ClientConnInterface) GroupJoinRequestServiceClient {
	return &groupJoinRequestServiceClient{cc}
}

func (c *groupJoinRequestServiceClient) GetGroupJoinRequestListByUserId(ctx context.Context, in *GetGroupJoinRequestListRequest, opts ...grpc.CallOption) (*GroupJoinRequestListResponse, error) {
	out := new(GroupJoinRequestListResponse)
	err := c.cc.Invoke(ctx, GroupJoinRequestService_GetGroupJoinRequestListByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupJoinRequestServiceClient) GetGroupJoinRequestByGroupIdAndUserId(ctx context.Context, in *GetGroupJoinRequestByGroupIdAndUserIdRequest, opts ...grpc.CallOption) (*GetGroupJoinRequestByGroupIdAndUserIdResponse, error) {
	out := new(GetGroupJoinRequestByGroupIdAndUserIdResponse)
	err := c.cc.Invoke(ctx, GroupJoinRequestService_GetGroupJoinRequestByGroupIdAndUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupJoinRequestServiceClient) JoinGroup(ctx context.Context, in *JoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error) {
	out := new(JoinGroupResponse)
	err := c.cc.Invoke(ctx, GroupJoinRequestService_JoinGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupJoinRequestServiceClient) InviteJoinGroup(ctx context.Context, in *InviteJoinGroupRequest, opts ...grpc.CallOption) (*JoinGroupResponse, error) {
	out := new(JoinGroupResponse)
	err := c.cc.Invoke(ctx, GroupJoinRequestService_InviteJoinGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupJoinRequestServiceClient) ManageGroupJoinRequestByID(ctx context.Context, in *ManageGroupJoinRequestByIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, GroupJoinRequestService_ManageGroupJoinRequestByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupJoinRequestServiceClient) GetGroupJoinRequestByID(ctx context.Context, in *GetGroupJoinRequestByIDRequest, opts ...grpc.CallOption) (*GetGroupJoinRequestByIDResponse, error) {
	out := new(GetGroupJoinRequestByIDResponse)
	err := c.cc.Invoke(ctx, GroupJoinRequestService_GetGroupJoinRequestByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupJoinRequestServiceServer is the server API for GroupJoinRequestService service.
// All implementations must embed UnimplementedGroupJoinRequestServiceServer
// for forward compatibility
type GroupJoinRequestServiceServer interface {
	// 获取用户加入群聊列表
	GetGroupJoinRequestListByUserId(context.Context, *GetGroupJoinRequestListRequest) (*GroupJoinRequestListResponse, error)
	// 根据群聊id和用户id获取群聊加入请求
	GetGroupJoinRequestByGroupIdAndUserId(context.Context, *GetGroupJoinRequestByGroupIdAndUserIdRequest) (*GetGroupJoinRequestByGroupIdAndUserIdResponse, error)
	// 申请加入群聊
	JoinGroup(context.Context, *JoinGroupRequest) (*JoinGroupResponse, error)
	// 邀请用户加入群聊
	InviteJoinGroup(context.Context, *InviteJoinGroupRequest) (*JoinGroupResponse, error)
	// 用户管理入群邀请
	ManageGroupJoinRequestByID(context.Context, *ManageGroupJoinRequestByIDRequest) (*emptypb.Empty, error)
	// 根据请求id获取申请
	GetGroupJoinRequestByID(context.Context, *GetGroupJoinRequestByIDRequest) (*GetGroupJoinRequestByIDResponse, error)
	mustEmbedUnimplementedGroupJoinRequestServiceServer()
}

// UnimplementedGroupJoinRequestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupJoinRequestServiceServer struct {
}

func (UnimplementedGroupJoinRequestServiceServer) GetGroupJoinRequestListByUserId(context.Context, *GetGroupJoinRequestListRequest) (*GroupJoinRequestListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupJoinRequestListByUserId not implemented")
}
func (UnimplementedGroupJoinRequestServiceServer) GetGroupJoinRequestByGroupIdAndUserId(context.Context, *GetGroupJoinRequestByGroupIdAndUserIdRequest) (*GetGroupJoinRequestByGroupIdAndUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupJoinRequestByGroupIdAndUserId not implemented")
}
func (UnimplementedGroupJoinRequestServiceServer) JoinGroup(context.Context, *JoinGroupRequest) (*JoinGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGroup not implemented")
}
func (UnimplementedGroupJoinRequestServiceServer) InviteJoinGroup(context.Context, *InviteJoinGroupRequest) (*JoinGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteJoinGroup not implemented")
}
func (UnimplementedGroupJoinRequestServiceServer) ManageGroupJoinRequestByID(context.Context, *ManageGroupJoinRequestByIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManageGroupJoinRequestByID not implemented")
}
func (UnimplementedGroupJoinRequestServiceServer) GetGroupJoinRequestByID(context.Context, *GetGroupJoinRequestByIDRequest) (*GetGroupJoinRequestByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupJoinRequestByID not implemented")
}
func (UnimplementedGroupJoinRequestServiceServer) mustEmbedUnimplementedGroupJoinRequestServiceServer() {
}

// UnsafeGroupJoinRequestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupJoinRequestServiceServer will
// result in compilation errors.
type UnsafeGroupJoinRequestServiceServer interface {
	mustEmbedUnimplementedGroupJoinRequestServiceServer()
}

func RegisterGroupJoinRequestServiceServer(s grpc.ServiceRegistrar, srv GroupJoinRequestServiceServer) {
	s.RegisterService(&GroupJoinRequestService_ServiceDesc, srv)
}

func _GroupJoinRequestService_GetGroupJoinRequestListByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupJoinRequestListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupJoinRequestServiceServer).GetGroupJoinRequestListByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupJoinRequestService_GetGroupJoinRequestListByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupJoinRequestServiceServer).GetGroupJoinRequestListByUserId(ctx, req.(*GetGroupJoinRequestListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupJoinRequestService_GetGroupJoinRequestByGroupIdAndUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupJoinRequestByGroupIdAndUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupJoinRequestServiceServer).GetGroupJoinRequestByGroupIdAndUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupJoinRequestService_GetGroupJoinRequestByGroupIdAndUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupJoinRequestServiceServer).GetGroupJoinRequestByGroupIdAndUserId(ctx, req.(*GetGroupJoinRequestByGroupIdAndUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupJoinRequestService_JoinGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupJoinRequestServiceServer).JoinGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupJoinRequestService_JoinGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupJoinRequestServiceServer).JoinGroup(ctx, req.(*JoinGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupJoinRequestService_InviteJoinGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteJoinGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupJoinRequestServiceServer).InviteJoinGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupJoinRequestService_InviteJoinGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupJoinRequestServiceServer).InviteJoinGroup(ctx, req.(*InviteJoinGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupJoinRequestService_ManageGroupJoinRequestByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManageGroupJoinRequestByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupJoinRequestServiceServer).ManageGroupJoinRequestByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupJoinRequestService_ManageGroupJoinRequestByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupJoinRequestServiceServer).ManageGroupJoinRequestByID(ctx, req.(*ManageGroupJoinRequestByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupJoinRequestService_GetGroupJoinRequestByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupJoinRequestByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupJoinRequestServiceServer).GetGroupJoinRequestByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupJoinRequestService_GetGroupJoinRequestByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupJoinRequestServiceServer).GetGroupJoinRequestByID(ctx, req.(*GetGroupJoinRequestByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupJoinRequestService_ServiceDesc is the grpc.ServiceDesc for GroupJoinRequestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupJoinRequestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.GroupJoinRequestService",
	HandlerType: (*GroupJoinRequestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroupJoinRequestListByUserId",
			Handler:    _GroupJoinRequestService_GetGroupJoinRequestListByUserId_Handler,
		},
		{
			MethodName: "GetGroupJoinRequestByGroupIdAndUserId",
			Handler:    _GroupJoinRequestService_GetGroupJoinRequestByGroupIdAndUserId_Handler,
		},
		{
			MethodName: "JoinGroup",
			Handler:    _GroupJoinRequestService_JoinGroup_Handler,
		},
		{
			MethodName: "InviteJoinGroup",
			Handler:    _GroupJoinRequestService_InviteJoinGroup_Handler,
		},
		{
			MethodName: "ManageGroupJoinRequestByID",
			Handler:    _GroupJoinRequestService_ManageGroupJoinRequestByID_Handler,
		},
		{
			MethodName: "GetGroupJoinRequestByID",
			Handler:    _GroupJoinRequestService_GetGroupJoinRequestByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/group_join_request.proto",
}
