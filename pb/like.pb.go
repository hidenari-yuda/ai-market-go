// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: like.proto

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
type Like struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid      string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ItemId    int64  `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	UserId    int64  `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Like) Reset() {
	*x = Like{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Like) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Like) ProtoMessage() {}

func (x *Like) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Like.ProtoReflect.Descriptor instead.
func (*Like) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{0}
}

func (x *Like) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Like) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Like) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Like) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Like) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Like) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type LikeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Like []*Like `protobuf:"bytes,1,rep,name=like,proto3" json:"like,omitempty"`
}

func (x *LikeList) Reset() {
	*x = LikeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeList) ProtoMessage() {}

func (x *LikeList) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeList.ProtoReflect.Descriptor instead.
func (*LikeList) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{1}
}

func (x *LikeList) GetLike() []*Like {
	if x != nil {
		return x.Like
	}
	return nil
}

// request
type LikeIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LikeIdRequest) Reset() {
	*x = LikeIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeIdRequest) ProtoMessage() {}

func (x *LikeIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeIdRequest.ProtoReflect.Descriptor instead.
func (*LikeIdRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{2}
}

func (x *LikeIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type LikeIdListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *LikeIdListRequest) Reset() {
	*x = LikeIdListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeIdListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeIdListRequest) ProtoMessage() {}

func (x *LikeIdListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeIdListRequest.ProtoReflect.Descriptor instead.
func (*LikeIdListRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{3}
}

func (x *LikeIdListRequest) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type LikeUserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *LikeUserIdRequest) Reset() {
	*x = LikeUserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeUserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeUserIdRequest) ProtoMessage() {}

func (x *LikeUserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeUserIdRequest.ProtoReflect.Descriptor instead.
func (*LikeUserIdRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{4}
}

func (x *LikeUserIdRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type LikeItemIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId int64 `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *LikeItemIdRequest) Reset() {
	*x = LikeItemIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeItemIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeItemIdRequest) ProtoMessage() {}

func (x *LikeItemIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeItemIdRequest.ProtoReflect.Descriptor instead.
func (*LikeItemIdRequest) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{5}
}

func (x *LikeItemIdRequest) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

// response
type LikeBoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *LikeBoolResponse) Reset() {
	*x = LikeBoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeBoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeBoolResponse) ProtoMessage() {}

func (x *LikeBoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeBoolResponse.ProtoReflect.Descriptor instead.
func (*LikeBoolResponse) Descriptor() ([]byte, []int) {
	return file_like_proto_rawDescGZIP(), []int{6}
}

func (x *LikeBoolResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

var File_like_proto protoreflect.FileDescriptor

var file_like_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x69,
	0x6b, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9a, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69,
	0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2a, 0x0a, 0x08,
	0x4c, 0x69, 0x6b, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x6c, 0x69, 0x6b, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69,
	0x6b, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x4c, 0x69, 0x6b, 0x65,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x4c, 0x69, 0x6b,
	0x65, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c,
	0x0a, 0x11, 0x4c, 0x69, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x11,
	0x4c, 0x69, 0x6b, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x10, 0x4c, 0x69,
	0x6b, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x32, 0xb2, 0x03, 0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x0a,
	0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x1a, 0x0a, 0x2e, 0x6c, 0x69, 0x6b,
	0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x0a, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x1a, 0x16,
	0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x13, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c,
	0x69, 0x6b, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x13, 0x2e, 0x6c,
	0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x17, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6c, 0x69, 0x6b, 0x65,
	0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x2e, 0x6c,
	0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x6c, 0x69, 0x6b,
	0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x2e, 0x4c,
	0x69, 0x6b, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2e, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_like_proto_rawDescOnce sync.Once
	file_like_proto_rawDescData = file_like_proto_rawDesc
)

func file_like_proto_rawDescGZIP() []byte {
	file_like_proto_rawDescOnce.Do(func() {
		file_like_proto_rawDescData = protoimpl.X.CompressGZIP(file_like_proto_rawDescData)
	})
	return file_like_proto_rawDescData
}

var file_like_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_like_proto_goTypes = []interface{}{
	(*Like)(nil),              // 0: like.Like
	(*LikeList)(nil),          // 1: like.LikeList
	(*LikeIdRequest)(nil),     // 2: like.LikeIdRequest
	(*LikeIdListRequest)(nil), // 3: like.LikeIdListRequest
	(*LikeUserIdRequest)(nil), // 4: like.LikeUserIdRequest
	(*LikeItemIdRequest)(nil), // 5: like.LikeItemIdRequest
	(*LikeBoolResponse)(nil),  // 6: like.LikeBoolResponse
	(*emptypb.Empty)(nil),     // 7: google.protobuf.Empty
}
var file_like_proto_depIdxs = []int32{
	0, // 0: like.LikeList.like:type_name -> like.Like
	0, // 1: like.LikeService.Create:input_type -> like.Like
	0, // 2: like.LikeService.Update:input_type -> like.Like
	2, // 3: like.LikeService.Delete:input_type -> like.LikeIdRequest
	2, // 4: like.LikeService.GetById:input_type -> like.LikeIdRequest
	4, // 5: like.LikeService.GetListByUser:input_type -> like.LikeUserIdRequest
	5, // 6: like.LikeService.GetListByItem:input_type -> like.LikeItemIdRequest
	3, // 7: like.LikeService.GetListByIdList:input_type -> like.LikeIdListRequest
	7, // 8: like.LikeService.GetAll:input_type -> google.protobuf.Empty
	0, // 9: like.LikeService.Create:output_type -> like.Like
	6, // 10: like.LikeService.Update:output_type -> like.LikeBoolResponse
	6, // 11: like.LikeService.Delete:output_type -> like.LikeBoolResponse
	0, // 12: like.LikeService.GetById:output_type -> like.Like
	1, // 13: like.LikeService.GetListByUser:output_type -> like.LikeList
	1, // 14: like.LikeService.GetListByItem:output_type -> like.LikeList
	1, // 15: like.LikeService.GetListByIdList:output_type -> like.LikeList
	1, // 16: like.LikeService.GetAll:output_type -> like.LikeList
	9, // [9:17] is the sub-list for method output_type
	1, // [1:9] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_like_proto_init() }
func file_like_proto_init() {
	if File_like_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_like_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Like); i {
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
		file_like_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeList); i {
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
		file_like_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeIdRequest); i {
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
		file_like_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeIdListRequest); i {
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
		file_like_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeUserIdRequest); i {
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
		file_like_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeItemIdRequest); i {
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
		file_like_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeBoolResponse); i {
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
			RawDescriptor: file_like_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_like_proto_goTypes,
		DependencyIndexes: file_like_proto_depIdxs,
		MessageInfos:      file_like_proto_msgTypes,
	}.Build()
	File_like_proto = out.File
	file_like_proto_rawDesc = nil
	file_like_proto_goTypes = nil
	file_like_proto_depIdxs = nil
}
