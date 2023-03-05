// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: following.proto

// Specify package name to avoid name collision

package pb

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

// ユーザー情報を表すメッセージ型
type Following struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid               string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	FollowingingUserId int64  `protobuf:"varint,3,opt,name=followinging_user_id,json=followingingUserId,proto3" json:"followinging_user_id,omitempty"`
	FollowingedUserId  int64  `protobuf:"varint,4,opt,name=followinged_user_id,json=followingedUserId,proto3" json:"followinged_user_id,omitempty"`
	CreatedAt          string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt          string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Following) Reset() {
	*x = Following{}
	if protoimpl.UnsafeEnabled {
		mi := &file_following_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Following) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Following) ProtoMessage() {}

func (x *Following) ProtoReflect() protoreflect.Message {
	mi := &file_following_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Following.ProtoReflect.Descriptor instead.
func (*Following) Descriptor() ([]byte, []int) {
	return file_following_proto_rawDescGZIP(), []int{0}
}

func (x *Following) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Following) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Following) GetFollowingingUserId() int64 {
	if x != nil {
		return x.FollowingingUserId
	}
	return 0
}

func (x *Following) GetFollowingedUserId() int64 {
	if x != nil {
		return x.FollowingedUserId
	}
	return 0
}

func (x *Following) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Following) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type FollowingList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Following []*Following `protobuf:"bytes,1,rep,name=following,proto3" json:"following,omitempty"`
}

func (x *FollowingList) Reset() {
	*x = FollowingList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_following_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowingList) ProtoMessage() {}

func (x *FollowingList) ProtoReflect() protoreflect.Message {
	mi := &file_following_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowingList.ProtoReflect.Descriptor instead.
func (*FollowingList) Descriptor() ([]byte, []int) {
	return file_following_proto_rawDescGZIP(), []int{1}
}

func (x *FollowingList) GetFollowing() []*Following {
	if x != nil {
		return x.Following
	}
	return nil
}

// request
type FollowingIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FollowingIdRequest) Reset() {
	*x = FollowingIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_following_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowingIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowingIdRequest) ProtoMessage() {}

func (x *FollowingIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_following_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowingIdRequest.ProtoReflect.Descriptor instead.
func (*FollowingIdRequest) Descriptor() ([]byte, []int) {
	return file_following_proto_rawDescGZIP(), []int{2}
}

func (x *FollowingIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FollowingIdListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *FollowingIdListRequest) Reset() {
	*x = FollowingIdListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_following_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowingIdListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowingIdListRequest) ProtoMessage() {}

func (x *FollowingIdListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_following_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowingIdListRequest.ProtoReflect.Descriptor instead.
func (*FollowingIdListRequest) Descriptor() ([]byte, []int) {
	return file_following_proto_rawDescGZIP(), []int{3}
}

func (x *FollowingIdListRequest) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type FollowingUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *FollowingUserIdRequest) Reset() {
	*x = FollowingUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_following_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowingUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowingUserIdRequest) ProtoMessage() {}

func (x *FollowingUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_following_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowingUserIdRequest.ProtoReflect.Descriptor instead.
func (*FollowingUserIdRequest) Descriptor() ([]byte, []int) {
	return file_following_proto_rawDescGZIP(), []int{4}
}

func (x *FollowingUserIdRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// response
type FollowingBoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FollowingBoolResponse) Reset() {
	*x = FollowingBoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_following_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowingBoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowingBoolResponse) ProtoMessage() {}

func (x *FollowingBoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_following_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowingBoolResponse.ProtoReflect.Descriptor instead.
func (*FollowingBoolResponse) Descriptor() ([]byte, []int) {
	return file_following_proto_rawDescGZIP(), []int{5}
}

func (x *FollowingBoolResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

var File_following_proto protoreflect.FileDescriptor

var file_following_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x09, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x13, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x43, 0x0a, 0x0d, 0x46,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x09,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x22, 0x24, 0x0a, 0x12, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x31, 0x0a, 0x16, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x15, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x32, 0xde, 0x04, 0x0a, 0x10, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x14, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x1a, 0x14, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12,
	0x42, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x1a,
	0x20, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x2e,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x22, 0x00, 0x12, 0x56, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e,
	0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67,
	0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69,
	0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x69, 0x6e, 0x67, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_following_proto_rawDescOnce sync.Once
	file_following_proto_rawDescData = file_following_proto_rawDesc
)

func file_following_proto_rawDescGZIP() []byte {
	file_following_proto_rawDescOnce.Do(func() {
		file_following_proto_rawDescData = protoimpl.X.CompressGZIP(file_following_proto_rawDescData)
	})
	return file_following_proto_rawDescData
}

var file_following_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_following_proto_goTypes = []interface{}{
	(*Following)(nil),              // 0: following.Following
	(*FollowingList)(nil),          // 1: following.FollowingList
	(*FollowingIdRequest)(nil),     // 2: following.FollowingIdRequest
	(*FollowingIdListRequest)(nil), // 3: following.FollowingIdListRequest
	(*FollowingUserIdRequest)(nil), // 4: following.FollowingUserIdRequest
	(*FollowingBoolResponse)(nil),  // 5: following.FollowingBoolResponse
	(*emptypb.Empty)(nil),          // 6: google.protobuf.Empty
}
var file_following_proto_depIdxs = []int32{
	0, // 0: following.FollowingList.following:type_name -> following.Following
	0, // 1: following.FollowingService.Create:input_type -> following.Following
	0, // 2: following.FollowingService.Update:input_type -> following.Following
	2, // 3: following.FollowingService.Delete:input_type -> following.FollowingIdRequest
	2, // 4: following.FollowingService.GetById:input_type -> following.FollowingIdRequest
	4, // 5: following.FollowingService.GetFollowedListByUser:input_type -> following.FollowingUserIdRequest
	4, // 6: following.FollowingService.GetFollowingListByUser:input_type -> following.FollowingUserIdRequest
	3, // 7: following.FollowingService.GetListByIdList:input_type -> following.FollowingIdListRequest
	6, // 8: following.FollowingService.GetAll:input_type -> google.protobuf.Empty
	0, // 9: following.FollowingService.Create:output_type -> following.Following
	5, // 10: following.FollowingService.Update:output_type -> following.FollowingBoolResponse
	5, // 11: following.FollowingService.Delete:output_type -> following.FollowingBoolResponse
	0, // 12: following.FollowingService.GetById:output_type -> following.Following
	1, // 13: following.FollowingService.GetFollowedListByUser:output_type -> following.FollowingList
	1, // 14: following.FollowingService.GetFollowingListByUser:output_type -> following.FollowingList
	1, // 15: following.FollowingService.GetListByIdList:output_type -> following.FollowingList
	1, // 16: following.FollowingService.GetAll:output_type -> following.FollowingList
	9, // [9:17] is the sub-list for method output_type
	1, // [1:9] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_following_proto_init() }
func file_following_proto_init() {
	if File_following_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_following_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Following); i {
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
		file_following_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowingList); i {
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
		file_following_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowingIdRequest); i {
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
		file_following_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowingIdListRequest); i {
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
		file_following_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowingUserIdRequest); i {
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
		file_following_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowingBoolResponse); i {
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
			RawDescriptor: file_following_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_following_proto_goTypes,
		DependencyIndexes: file_following_proto_depIdxs,
		MessageInfos:      file_following_proto_msgTypes,
	}.Build()
	File_following_proto = out.File
	file_following_proto_rawDesc = nil
	file_following_proto_goTypes = nil
	file_following_proto_depIdxs = nil
}
