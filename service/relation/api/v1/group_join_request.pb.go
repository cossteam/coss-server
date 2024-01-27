// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: api/v1/group_join_request.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GroupRequestStatus int32

const (
	GroupRequestStatus_Pending    GroupRequestStatus = 0 //等待中
	GroupRequestStatus_Accepted   GroupRequestStatus = 1 //已通过
	GroupRequestStatus_Rejected   GroupRequestStatus = 2 //已拒绝
	GroupRequestStatus_Invitation GroupRequestStatus = 3 //邀请中
)

// Enum value maps for GroupRequestStatus.
var (
	GroupRequestStatus_name = map[int32]string{
		0: "Pending",
		1: "Accepted",
		2: "Rejected",
		3: "Invitation",
	}
	GroupRequestStatus_value = map[string]int32{
		"Pending":    0,
		"Accepted":   1,
		"Rejected":   2,
		"Invitation": 3,
	}
)

func (x GroupRequestStatus) Enum() *GroupRequestStatus {
	p := new(GroupRequestStatus)
	*p = x
	return p
}

func (x GroupRequestStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupRequestStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_group_join_request_proto_enumTypes[0].Descriptor()
}

func (GroupRequestStatus) Type() protoreflect.EnumType {
	return &file_api_v1_group_join_request_proto_enumTypes[0]
}

func (x GroupRequestStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupRequestStatus.Descriptor instead.
func (GroupRequestStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_group_join_request_proto_rawDescGZIP(), []int{0}
}

type GetGroupJoinRequestListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"string"
	UserId string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"string"`
}

func (x *GetGroupJoinRequestListRequest) Reset() {
	*x = GetGroupJoinRequestListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_join_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupJoinRequestListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupJoinRequestListRequest) ProtoMessage() {}

func (x *GetGroupJoinRequestListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_join_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupJoinRequestListRequest.ProtoReflect.Descriptor instead.
func (*GetGroupJoinRequestListRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_group_join_request_proto_rawDescGZIP(), []int{0}
}

func (x *GetGroupJoinRequestListRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GroupJoinRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"group_id"
	GroupId uint32 `protobuf:"varint,2,opt,name=GroupId,proto3" json:"group_id"`
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,3,opt,name=UserId,proto3" json:"user_id"`
	// @inject_tag: json:"inviter_id"
	InviterId string `protobuf:"bytes,4,opt,name=InviterId,proto3" json:"inviter_id"`
	// @inject_tag: json:"created_at"
	CreatedAt uint64 `protobuf:"varint,5,opt,name=CreatedAt,proto3" json:"created_at"`
	// @inject_tag: json:"status"
	Status GroupRequestStatus `protobuf:"varint,6,opt,name=Status,proto3,enum=v1.GroupRequestStatus" json:"status"`
	// @inject_tag: json:"remark"
	Remark string `protobuf:"bytes,7,opt,name=Remark,proto3" json:"remark"`
}

func (x *GroupJoinRequestResponse) Reset() {
	*x = GroupJoinRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_join_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupJoinRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupJoinRequestResponse) ProtoMessage() {}

func (x *GroupJoinRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_join_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupJoinRequestResponse.ProtoReflect.Descriptor instead.
func (*GroupJoinRequestResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_group_join_request_proto_rawDescGZIP(), []int{1}
}

func (x *GroupJoinRequestResponse) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GroupJoinRequestResponse) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupJoinRequestResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GroupJoinRequestResponse) GetInviterId() string {
	if x != nil {
		return x.InviterId
	}
	return ""
}

func (x *GroupJoinRequestResponse) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *GroupJoinRequestResponse) GetStatus() GroupRequestStatus {
	if x != nil {
		return x.Status
	}
	return GroupRequestStatus_Pending
}

func (x *GroupJoinRequestResponse) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type GroupJoinRequestListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"group_join_request_list"
	GroupJoinRequestResponses []*GroupJoinRequestResponse `protobuf:"bytes,1,rep,name=GroupJoinRequestResponses,proto3" json:"group_join_request_list"`
}

func (x *GroupJoinRequestListResponse) Reset() {
	*x = GroupJoinRequestListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_join_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupJoinRequestListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupJoinRequestListResponse) ProtoMessage() {}

func (x *GroupJoinRequestListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_join_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupJoinRequestListResponse.ProtoReflect.Descriptor instead.
func (*GroupJoinRequestListResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_group_join_request_proto_rawDescGZIP(), []int{2}
}

func (x *GroupJoinRequestListResponse) GetGroupJoinRequestResponses() []*GroupJoinRequestResponse {
	if x != nil {
		return x.GroupJoinRequestResponses
	}
	return nil
}

type GetGroupJoinRequestByGroupIdAndUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"group_id"
	GroupId uint32 `protobuf:"varint,1,opt,name=GroupId,proto3" json:"group_id"`
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"user_id"`
}

func (x *GetGroupJoinRequestByGroupIdAndUserIdRequest) Reset() {
	*x = GetGroupJoinRequestByGroupIdAndUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_join_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupJoinRequestByGroupIdAndUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupJoinRequestByGroupIdAndUserIdRequest) ProtoMessage() {}

func (x *GetGroupJoinRequestByGroupIdAndUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_join_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupJoinRequestByGroupIdAndUserIdRequest.ProtoReflect.Descriptor instead.
func (*GetGroupJoinRequestByGroupIdAndUserIdRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_group_join_request_proto_rawDescGZIP(), []int{3}
}

func (x *GetGroupJoinRequestByGroupIdAndUserIdRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GetGroupJoinRequestByGroupIdAndUserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetGroupJoinRequestByGroupIdAndUserIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"group_id"
	GroupId uint32 `protobuf:"varint,2,opt,name=GroupId,proto3" json:"group_id"`
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,3,opt,name=UserId,proto3" json:"user_id"`
	// @inject_tag: json:"inviter_id"
	InviterId string `protobuf:"bytes,4,opt,name=InviterId,proto3" json:"inviter_id"`
	// @inject_tag: json:"created_at"
	CreatedAt uint64 `protobuf:"varint,5,opt,name=CreatedAt,proto3" json:"created_at"`
	// @inject_tag: json:"status"
	Status GroupRequestStatus `protobuf:"varint,6,opt,name=Status,proto3,enum=v1.GroupRequestStatus" json:"status"`
	// @inject_tag: json:"remark"
	Remark string `protobuf:"bytes,7,opt,name=Remark,proto3" json:"remark"`
}

func (x *GetGroupJoinRequestByGroupIdAndUserIdResponse) Reset() {
	*x = GetGroupJoinRequestByGroupIdAndUserIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_join_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupJoinRequestByGroupIdAndUserIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupJoinRequestByGroupIdAndUserIdResponse) ProtoMessage() {}

func (x *GetGroupJoinRequestByGroupIdAndUserIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_join_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupJoinRequestByGroupIdAndUserIdResponse.ProtoReflect.Descriptor instead.
func (*GetGroupJoinRequestByGroupIdAndUserIdResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_group_join_request_proto_rawDescGZIP(), []int{4}
}

func (x *GetGroupJoinRequestByGroupIdAndUserIdResponse) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GetGroupJoinRequestByGroupIdAndUserIdResponse) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GetGroupJoinRequestByGroupIdAndUserIdResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetGroupJoinRequestByGroupIdAndUserIdResponse) GetInviterId() string {
	if x != nil {
		return x.InviterId
	}
	return ""
}

func (x *GetGroupJoinRequestByGroupIdAndUserIdResponse) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *GetGroupJoinRequestByGroupIdAndUserIdResponse) GetStatus() GroupRequestStatus {
	if x != nil {
		return x.Status
	}
	return GroupRequestStatus_Pending
}

func (x *GetGroupJoinRequestByGroupIdAndUserIdResponse) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type JoinGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"group_id"
	GroupId uint32 `protobuf:"varint,1,opt,name=GroupId,proto3" json:"group_id"` // 群聊id
	// @inject_tag: json:"user_id"
	UserId string `protobuf:"bytes,2,opt,name=UserId,proto3" json:"user_id"` // 用户id
	// @inject_tag: json:"msg"
	Msg string `protobuf:"bytes,3,opt,name=Msg,proto3" json:"msg"` // 申请信息
}

func (x *JoinGroupRequest) Reset() {
	*x = JoinGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_join_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGroupRequest) ProtoMessage() {}

func (x *JoinGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_join_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGroupRequest.ProtoReflect.Descriptor instead.
func (*JoinGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_group_join_request_proto_rawDescGZIP(), []int{5}
}

func (x *JoinGroupRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *JoinGroupRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *JoinGroupRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type JoinGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *JoinGroupResponse) Reset() {
	*x = JoinGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_join_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinGroupResponse) ProtoMessage() {}

func (x *JoinGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_join_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinGroupResponse.ProtoReflect.Descriptor instead.
func (*JoinGroupResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_group_join_request_proto_rawDescGZIP(), []int{6}
}

type InviteJoinGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"group_id"
	GroupId uint32 `protobuf:"varint,1,opt,name=GroupId,proto3" json:"group_id"`
	// @inject_tag: json:"inviter_id"
	InviterId string `protobuf:"bytes,2,opt,name=InviterId,proto3" json:"inviter_id"`
	// @inject_tag: json:"member"
	Member []string `protobuf:"bytes,3,rep,name=Member,proto3" json:"member"`
}

func (x *InviteJoinGroupRequest) Reset() {
	*x = InviteJoinGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_join_request_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteJoinGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteJoinGroupRequest) ProtoMessage() {}

func (x *InviteJoinGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_join_request_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteJoinGroupRequest.ProtoReflect.Descriptor instead.
func (*InviteJoinGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_group_join_request_proto_rawDescGZIP(), []int{7}
}

func (x *InviteJoinGroupRequest) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *InviteJoinGroupRequest) GetInviterId() string {
	if x != nil {
		return x.InviterId
	}
	return ""
}

func (x *InviteJoinGroupRequest) GetMember() []string {
	if x != nil {
		return x.Member
	}
	return nil
}

type ManageGroupJoinRequestByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	// @inject_tag: json:"status"
	Status GroupRequestStatus `protobuf:"varint,2,opt,name=Status,proto3,enum=v1.GroupRequestStatus" json:"status"`
}

func (x *ManageGroupJoinRequestByIDRequest) Reset() {
	*x = ManageGroupJoinRequestByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_group_join_request_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManageGroupJoinRequestByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManageGroupJoinRequestByIDRequest) ProtoMessage() {}

func (x *ManageGroupJoinRequestByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_group_join_request_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManageGroupJoinRequestByIDRequest.ProtoReflect.Descriptor instead.
func (*ManageGroupJoinRequestByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_group_join_request_proto_rawDescGZIP(), []int{8}
}

func (x *ManageGroupJoinRequestByIDRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ManageGroupJoinRequestByIDRequest) GetStatus() GroupRequestStatus {
	if x != nil {
		return x.Status
	}
	return GroupRequestStatus_Pending
}

var File_api_v1_group_join_request_proto protoreflect.FileDescriptor

var file_api_v1_group_join_request_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6a,
	0x6f, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x38, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xe0, 0x01, 0x0a,
	0x18, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x49,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22,
	0x7a, 0x0a, 0x1c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5a, 0x0a, 0x19, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x19, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x60, 0x0a, 0x2c, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xf5, 0x01,
	0x0a, 0x2d, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x41, 0x6e,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2e, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x56, 0x0a, 0x10, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4d,
	0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x13, 0x0a,
	0x11, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x68, 0x0a, 0x16, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x63, 0x0a, 0x21,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2a, 0x4d, 0x0a, 0x12, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x03,
	0x32, 0xf8, 0x03, 0x0a, 0x17, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x1f,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x22, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8e, 0x01, 0x0a, 0x25, 0x47, 0x65, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42,
	0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x30, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a,
	0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x09, 0x4a, 0x6f, 0x69, 0x6e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x76, 0x31,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x4a, 0x6f,
	0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x1a,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x25, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x73, 0x69, 0x6d,
	0x2f, 0x63, 0x6f, 0x73, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_group_join_request_proto_rawDescOnce sync.Once
	file_api_v1_group_join_request_proto_rawDescData = file_api_v1_group_join_request_proto_rawDesc
)

func file_api_v1_group_join_request_proto_rawDescGZIP() []byte {
	file_api_v1_group_join_request_proto_rawDescOnce.Do(func() {
		file_api_v1_group_join_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_group_join_request_proto_rawDescData)
	})
	return file_api_v1_group_join_request_proto_rawDescData
}

var file_api_v1_group_join_request_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v1_group_join_request_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_v1_group_join_request_proto_goTypes = []interface{}{
	(GroupRequestStatus)(0),                               // 0: v1.GroupRequestStatus
	(*GetGroupJoinRequestListRequest)(nil),                // 1: v1.GetGroupJoinRequestListRequest
	(*GroupJoinRequestResponse)(nil),                      // 2: v1.GroupJoinRequestResponse
	(*GroupJoinRequestListResponse)(nil),                  // 3: v1.GroupJoinRequestListResponse
	(*GetGroupJoinRequestByGroupIdAndUserIdRequest)(nil),  // 4: v1.GetGroupJoinRequestByGroupIdAndUserIdRequest
	(*GetGroupJoinRequestByGroupIdAndUserIdResponse)(nil), // 5: v1.GetGroupJoinRequestByGroupIdAndUserIdResponse
	(*JoinGroupRequest)(nil),                              // 6: v1.JoinGroupRequest
	(*JoinGroupResponse)(nil),                             // 7: v1.JoinGroupResponse
	(*InviteJoinGroupRequest)(nil),                        // 8: v1.InviteJoinGroupRequest
	(*ManageGroupJoinRequestByIDRequest)(nil),             // 9: v1.ManageGroupJoinRequestByIDRequest
	(*emptypb.Empty)(nil),                                 // 10: google.protobuf.Empty
}
var file_api_v1_group_join_request_proto_depIdxs = []int32{
	0,  // 0: v1.GroupJoinRequestResponse.Status:type_name -> v1.GroupRequestStatus
	2,  // 1: v1.GroupJoinRequestListResponse.GroupJoinRequestResponses:type_name -> v1.GroupJoinRequestResponse
	0,  // 2: v1.GetGroupJoinRequestByGroupIdAndUserIdResponse.Status:type_name -> v1.GroupRequestStatus
	0,  // 3: v1.ManageGroupJoinRequestByIDRequest.Status:type_name -> v1.GroupRequestStatus
	1,  // 4: v1.GroupJoinRequestService.GetGroupJoinRequestListByUserId:input_type -> v1.GetGroupJoinRequestListRequest
	4,  // 5: v1.GroupJoinRequestService.GetGroupJoinRequestByGroupIdAndUserId:input_type -> v1.GetGroupJoinRequestByGroupIdAndUserIdRequest
	6,  // 6: v1.GroupJoinRequestService.JoinGroup:input_type -> v1.JoinGroupRequest
	8,  // 7: v1.GroupJoinRequestService.InviteJoinGroup:input_type -> v1.InviteJoinGroupRequest
	9,  // 8: v1.GroupJoinRequestService.ManageGroupJoinRequestByID:input_type -> v1.ManageGroupJoinRequestByIDRequest
	3,  // 9: v1.GroupJoinRequestService.GetGroupJoinRequestListByUserId:output_type -> v1.GroupJoinRequestListResponse
	5,  // 10: v1.GroupJoinRequestService.GetGroupJoinRequestByGroupIdAndUserId:output_type -> v1.GetGroupJoinRequestByGroupIdAndUserIdResponse
	7,  // 11: v1.GroupJoinRequestService.JoinGroup:output_type -> v1.JoinGroupResponse
	7,  // 12: v1.GroupJoinRequestService.InviteJoinGroup:output_type -> v1.JoinGroupResponse
	10, // 13: v1.GroupJoinRequestService.ManageGroupJoinRequestByID:output_type -> google.protobuf.Empty
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1_group_join_request_proto_init() }
func file_api_v1_group_join_request_proto_init() {
	if File_api_v1_group_join_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_group_join_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupJoinRequestListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_group_join_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupJoinRequestResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_group_join_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupJoinRequestListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_group_join_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupJoinRequestByGroupIdAndUserIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_group_join_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupJoinRequestByGroupIdAndUserIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_group_join_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_group_join_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinGroupResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_group_join_request_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteJoinGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_v1_group_join_request_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManageGroupJoinRequestByIDRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_v1_group_join_request_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_group_join_request_proto_goTypes,
		DependencyIndexes: file_api_v1_group_join_request_proto_depIdxs,
		EnumInfos:         file_api_v1_group_join_request_proto_enumTypes,
		MessageInfos:      file_api_v1_group_join_request_proto_msgTypes,
	}.Build()
	File_api_v1_group_join_request_proto = out.File
	file_api_v1_group_join_request_proto_rawDesc = nil
	file_api_v1_group_join_request_proto_goTypes = nil
	file_api_v1_group_join_request_proto_depIdxs = nil
}
