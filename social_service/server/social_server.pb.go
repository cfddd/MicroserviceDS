// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.0
// source: social_server.proto

package server

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FollowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags:json:"user_id" form:"user_id" uri:"user_id"
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	// @gotags:json:"to_user_id" form:"to_user_id" uri:"to_user_id"
	ToUserId int64 `protobuf:"varint,2,opt,name=ToUserId,proto3" json:"ToUserId,omitempty"`
	// @gotags:json:"action_type" form:"action_type" uri:"action_type"
	ActionType int32 `protobuf:"varint,3,opt,name=ActionType,proto3" json:"ActionType,omitempty"`
}

func (x *FollowRequest) Reset() {
	*x = FollowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRequest) ProtoMessage() {}

func (x *FollowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_social_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRequest.ProtoReflect.Descriptor instead.
func (*FollowRequest) Descriptor() ([]byte, []int) {
	return file_social_server_proto_rawDescGZIP(), []int{0}
}

func (x *FollowRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FollowRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *FollowRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type FollowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags:json:"status_code" form:"status_code" uri:"status_code"
	StatusCode int32 `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	// @gotags:json:"status_msg" form:"status_msg" uri:"status_msg"
	StatusMsg string `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
}

func (x *FollowResponse) Reset() {
	*x = FollowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowResponse) ProtoMessage() {}

func (x *FollowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_social_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowResponse.ProtoReflect.Descriptor instead.
func (*FollowResponse) Descriptor() ([]byte, []int) {
	return file_social_server_proto_rawDescGZIP(), []int{1}
}

func (x *FollowResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FollowResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type FollowListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags:json:"user_id" form:"user_id" uri:"user_id"
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *FollowListRequest) Reset() {
	*x = FollowListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowListRequest) ProtoMessage() {}

func (x *FollowListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_social_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowListRequest.ProtoReflect.Descriptor instead.
func (*FollowListRequest) Descriptor() ([]byte, []int) {
	return file_social_server_proto_rawDescGZIP(), []int{2}
}

func (x *FollowListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FollowListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags:json:"status_code" form:"status_code" uri:"status_code"
	StatusCode int32 `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	// @gotags:json:"status_msg" form:"status_msg" uri:"status_msg"
	StatusMsg string `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
	// @gotags:json:"user_list" form:"user_list" uri:"user_list"
	UserId []int64 `protobuf:"varint,3,rep,packed,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *FollowListResponse) Reset() {
	*x = FollowListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowListResponse) ProtoMessage() {}

func (x *FollowListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_social_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowListResponse.ProtoReflect.Descriptor instead.
func (*FollowListResponse) Descriptor() ([]byte, []int) {
	return file_social_server_proto_rawDescGZIP(), []int{3}
}

func (x *FollowListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FollowListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *FollowListResponse) GetUserId() []int64 {
	if x != nil {
		return x.UserId
	}
	return nil
}

type FollowInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags:json:"user_id" form:"user_id" uri:"user_id"
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	// @gotags:json:"to_user_id" form:"to_user_id" uri:"to_user_id"
	ToUserId []int64 `protobuf:"varint,2,rep,packed,name=ToUserId,proto3" json:"ToUserId,omitempty"`
}

func (x *FollowInfoRequest) Reset() {
	*x = FollowInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowInfoRequest) ProtoMessage() {}

func (x *FollowInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_social_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowInfoRequest.ProtoReflect.Descriptor instead.
func (*FollowInfoRequest) Descriptor() ([]byte, []int) {
	return file_social_server_proto_rawDescGZIP(), []int{4}
}

func (x *FollowInfoRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FollowInfoRequest) GetToUserId() []int64 {
	if x != nil {
		return x.ToUserId
	}
	return nil
}

type FollowInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags:json:"is_follow" form:"is_follow" uri:"is_follow"
	IsFollow bool `protobuf:"varint,1,opt,name=IsFollow,proto3" json:"IsFollow,omitempty"`
	// @gotags:json:"follow_count" form:"follow_count" uri:"follow_count"
	FollowCount int64 `protobuf:"varint,2,opt,name=FollowCount,proto3" json:"FollowCount,omitempty"`
	// @gotags:json:"follower_count" form:"follower_count" uri:"follower_count"
	FollowerCount int64 `protobuf:"varint,3,opt,name=FollowerCount,proto3" json:"FollowerCount,omitempty"`
	// @gotags:json:"to_user_id" form:"to_user_id" uri:"to_user_id"
	ToUserId int64 `protobuf:"varint,4,opt,name=ToUserId,proto3" json:"ToUserId,omitempty"`
}

func (x *FollowInfo) Reset() {
	*x = FollowInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowInfo) ProtoMessage() {}

func (x *FollowInfo) ProtoReflect() protoreflect.Message {
	mi := &file_social_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowInfo.ProtoReflect.Descriptor instead.
func (*FollowInfo) Descriptor() ([]byte, []int) {
	return file_social_server_proto_rawDescGZIP(), []int{5}
}

func (x *FollowInfo) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *FollowInfo) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *FollowInfo) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *FollowInfo) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

type FollowInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags:json:"status_code" form:"status_code" uri:"status_code"
	StatusCode int32 `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	// @gotags:json:"status_msg" form:"status_msg" uri:"status_msg"
	StatusMsg string `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
	// @gotags:json:"follow_info" form:"follow_info" uri:"follow_info"
	FollowInfo []*FollowInfo `protobuf:"bytes,3,rep,name=FollowInfo,proto3" json:"FollowInfo,omitempty"`
}

func (x *FollowInfoResponse) Reset() {
	*x = FollowInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowInfoResponse) ProtoMessage() {}

func (x *FollowInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_social_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowInfoResponse.ProtoReflect.Descriptor instead.
func (*FollowInfoResponse) Descriptor() ([]byte, []int) {
	return file_social_server_proto_rawDescGZIP(), []int{6}
}

func (x *FollowInfoResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FollowInfoResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *FollowInfoResponse) GetFollowInfo() []*FollowInfo {
	if x != nil {
		return x.FollowInfo
	}
	return nil
}

type PostMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags:json:"user_id" form:"user_id" uri:"user_id"
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	// @gotags:json:"to_user_id" form:"to_user_id" uri:"to_user_id"
	ToUserId int64 `protobuf:"varint,2,opt,name=ToUserId,proto3" json:"ToUserId,omitempty"`
	// @gotags:json:"action_type" form:"action_type" uri:"action_type"
	ActionType int32 `protobuf:"varint,3,opt,name=ActionType,proto3" json:"ActionType,omitempty"`
	// @gotags:json:"content" form:"content" uri:"content"
	Content string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *PostMessageRequest) Reset() {
	*x = PostMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostMessageRequest) ProtoMessage() {}

func (x *PostMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_social_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostMessageRequest.ProtoReflect.Descriptor instead.
func (*PostMessageRequest) Descriptor() ([]byte, []int) {
	return file_social_server_proto_rawDescGZIP(), []int{7}
}

func (x *PostMessageRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PostMessageRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *PostMessageRequest) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

func (x *PostMessageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type PostMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags:json:"status_code" form:"status_code" uri:"status_code"
	StatusCode int32 `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	// @gotags:json:"status_msg" form:"status_msg" uri:"status_msg"
	StatusMsg string `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
}

func (x *PostMessageResponse) Reset() {
	*x = PostMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostMessageResponse) ProtoMessage() {}

func (x *PostMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_social_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostMessageResponse.ProtoReflect.Descriptor instead.
func (*PostMessageResponse) Descriptor() ([]byte, []int) {
	return file_social_server_proto_rawDescGZIP(), []int{8}
}

func (x *PostMessageResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *PostMessageResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

type GetMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags:json:"user_id" form:"user_id" uri:"user_id"
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	// @gotags:json:"to_user_id" form:"to_user_id" uri:"to_user_id"
	ToUserId int64 `protobuf:"varint,2,opt,name=ToUserId,proto3" json:"ToUserId,omitempty"`
	// @gotags:json:"pre_msg_time" form:"pre_msg_time" uri:"pre_msg_time"
	PreMsgTime int64 `protobuf:"varint,3,opt,name=preMsgTime,proto3" json:"preMsgTime,omitempty"`
}

func (x *GetMessageRequest) Reset() {
	*x = GetMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequest) ProtoMessage() {}

func (x *GetMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_social_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequest.ProtoReflect.Descriptor instead.
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return file_social_server_proto_rawDescGZIP(), []int{9}
}

func (x *GetMessageRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetMessageRequest) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *GetMessageRequest) GetPreMsgTime() int64 {
	if x != nil {
		return x.PreMsgTime
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags:json:"id" form:"id" uri:"id"
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// @gotags:json:"user_id" form:"user_id" uri:"user_id"
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	// @gotags:json:"to_user_id" form:"to_user_id" uri:"to_user_id"
	ToUserId int64 `protobuf:"varint,3,opt,name=ToUserId,proto3" json:"ToUserId,omitempty"`
	// @gotags:json:"content" form:"content" uri:"content"
	Content string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	// @gotags:json:"created_at" form:"created_at" uri:"created_at"
	CreatedAt int64 `protobuf:"varint,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_server_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_social_server_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_social_server_proto_rawDescGZIP(), []int{10}
}

func (x *Message) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Message) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type GetMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags:json:"status_code" form:"status_code" uri:"status_code"
	StatusCode int32 `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	// @gotags:json:"status_msg" form:"status_msg" uri:"status_msg"
	StatusMsg string `protobuf:"bytes,2,opt,name=StatusMsg,proto3" json:"StatusMsg,omitempty"`
	// @gotags:json:"message_list" form:"message_list" uri:"message_list"
	Message []*Message `protobuf:"bytes,3,rep,name=message,proto3" json:"message,omitempty"`
}

func (x *GetMessageResponse) Reset() {
	*x = GetMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_social_server_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageResponse) ProtoMessage() {}

func (x *GetMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_social_server_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageResponse.ProtoReflect.Descriptor instead.
func (*GetMessageResponse) Descriptor() ([]byte, []int) {
	return file_social_server_proto_rawDescGZIP(), []int{11}
}

func (x *GetMessageResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetMessageResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *GetMessageResponse) GetMessage() []*Message {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_social_server_proto protoreflect.FileDescriptor

var file_social_server_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x63, 0x0a, 0x0d, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4e,
	0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x2b,
	0x0a, 0x11, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x6a, 0x0a, 0x12, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x11, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x8c, 0x01, 0x0a, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x49, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x20, 0x0a, 0x0b, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x82, 0x01, 0x0a, 0x12, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x12, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x53, 0x0a, 0x13, 0x50, 0x6f, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x67,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x4d, 0x73,
	0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x65,
	0x4d, 0x73, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x54,
	0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0x79, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4d, 0x73, 0x67, 0x12, 0x25, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc5, 0x03, 0x0a, 0x0d, 0x53,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0c,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x12, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_social_server_proto_rawDescOnce sync.Once
	file_social_server_proto_rawDescData = file_social_server_proto_rawDesc
)

func file_social_server_proto_rawDescGZIP() []byte {
	file_social_server_proto_rawDescOnce.Do(func() {
		file_social_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_social_server_proto_rawDescData)
	})
	return file_social_server_proto_rawDescData
}

var file_social_server_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_social_server_proto_goTypes = []interface{}{
	(*FollowRequest)(nil),       // 0: pb.FollowRequest
	(*FollowResponse)(nil),      // 1: pb.FollowResponse
	(*FollowListRequest)(nil),   // 2: pb.FollowListRequest
	(*FollowListResponse)(nil),  // 3: pb.FollowListResponse
	(*FollowInfoRequest)(nil),   // 4: pb.FollowInfoRequest
	(*FollowInfo)(nil),          // 5: pb.FollowInfo
	(*FollowInfoResponse)(nil),  // 6: pb.FollowInfoResponse
	(*PostMessageRequest)(nil),  // 7: pb.PostMessageRequest
	(*PostMessageResponse)(nil), // 8: pb.PostMessageResponse
	(*GetMessageRequest)(nil),   // 9: pb.GetMessageRequest
	(*Message)(nil),             // 10: pb.Message
	(*GetMessageResponse)(nil),  // 11: pb.GetMessageResponse
}
var file_social_server_proto_depIdxs = []int32{
	5,  // 0: pb.FollowInfoResponse.FollowInfo:type_name -> pb.FollowInfo
	10, // 1: pb.GetMessageResponse.message:type_name -> pb.Message
	0,  // 2: pb.SocialService.FollowAction:input_type -> pb.FollowRequest
	2,  // 3: pb.SocialService.GetFollowList:input_type -> pb.FollowListRequest
	2,  // 4: pb.SocialService.GetFollowerList:input_type -> pb.FollowListRequest
	2,  // 5: pb.SocialService.GetFriendList:input_type -> pb.FollowListRequest
	4,  // 6: pb.SocialService.GetFollowInfo:input_type -> pb.FollowInfoRequest
	7,  // 7: pb.SocialService.PostMessage:input_type -> pb.PostMessageRequest
	9,  // 8: pb.SocialService.GetMessage:input_type -> pb.GetMessageRequest
	1,  // 9: pb.SocialService.FollowAction:output_type -> pb.FollowResponse
	3,  // 10: pb.SocialService.GetFollowList:output_type -> pb.FollowListResponse
	3,  // 11: pb.SocialService.GetFollowerList:output_type -> pb.FollowListResponse
	3,  // 12: pb.SocialService.GetFriendList:output_type -> pb.FollowListResponse
	6,  // 13: pb.SocialService.GetFollowInfo:output_type -> pb.FollowInfoResponse
	8,  // 14: pb.SocialService.PostMessage:output_type -> pb.PostMessageResponse
	11, // 15: pb.SocialService.GetMessage:output_type -> pb.GetMessageResponse
	9,  // [9:16] is the sub-list for method output_type
	2,  // [2:9] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_social_server_proto_init() }
func file_social_server_proto_init() {
	if File_social_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_social_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRequest); i {
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
		file_social_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowResponse); i {
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
		file_social_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowListRequest); i {
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
		file_social_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowListResponse); i {
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
		file_social_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowInfoRequest); i {
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
		file_social_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowInfo); i {
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
		file_social_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowInfoResponse); i {
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
		file_social_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostMessageRequest); i {
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
		file_social_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostMessageResponse); i {
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
		file_social_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequest); i {
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
		file_social_server_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_social_server_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageResponse); i {
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
			RawDescriptor: file_social_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_social_server_proto_goTypes,
		DependencyIndexes: file_social_server_proto_depIdxs,
		MessageInfos:      file_social_server_proto_msgTypes,
	}.Build()
	File_social_server_proto = out.File
	file_social_server_proto_rawDesc = nil
	file_social_server_proto_goTypes = nil
	file_social_server_proto_depIdxs = nil
}
