// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: asp.proto

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

type Asp_Service int32

const (
	Asp_YOUTUBE   Asp_Service = 0
	Asp_TWITTER   Asp_Service = 1
	Asp_INSTAGRAM Asp_Service = 2
	Asp_FACEBOOK  Asp_Service = 3
	Asp_LINKEDIN  Asp_Service = 4
	Asp_TIKTOK    Asp_Service = 5
	Asp_PINTEREST Asp_Service = 6
	Asp_REDDIT    Asp_Service = 7
	Asp_SNAPCHAT  Asp_Service = 8
	Asp_TUMBLR    Asp_Service = 9
	Asp_TWITCH    Asp_Service = 10
	Asp_WEIBO     Asp_Service = 11
	Asp_WECHAT    Asp_Service = 12
	Asp_WHATSAPP  Asp_Service = 13
	Asp_LINE      Asp_Service = 14
	Asp_TELEGRAM  Asp_Service = 15
	Asp_VK        Asp_Service = 16
	Asp_YAHOO     Asp_Service = 17
	Asp_OTHER     Asp_Service = 18
)

// Enum value maps for Asp_Service.
var (
	Asp_Service_name = map[int32]string{
		0:  "YOUTUBE",
		1:  "TWITTER",
		2:  "INSTAGRAM",
		3:  "FACEBOOK",
		4:  "LINKEDIN",
		5:  "TIKTOK",
		6:  "PINTEREST",
		7:  "REDDIT",
		8:  "SNAPCHAT",
		9:  "TUMBLR",
		10: "TWITCH",
		11: "WEIBO",
		12: "WECHAT",
		13: "WHATSAPP",
		14: "LINE",
		15: "TELEGRAM",
		16: "VK",
		17: "YAHOO",
		18: "OTHER",
	}
	Asp_Service_value = map[string]int32{
		"YOUTUBE":   0,
		"TWITTER":   1,
		"INSTAGRAM": 2,
		"FACEBOOK":  3,
		"LINKEDIN":  4,
		"TIKTOK":    5,
		"PINTEREST": 6,
		"REDDIT":    7,
		"SNAPCHAT":  8,
		"TUMBLR":    9,
		"TWITCH":    10,
		"WEIBO":     11,
		"WECHAT":    12,
		"WHATSAPP":  13,
		"LINE":      14,
		"TELEGRAM":  15,
		"VK":        16,
		"YAHOO":     17,
		"OTHER":     18,
	}
)

func (x Asp_Service) Enum() *Asp_Service {
	p := new(Asp_Service)
	*p = x
	return p
}

func (x Asp_Service) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Asp_Service) Descriptor() protoreflect.EnumDescriptor {
	return file_asp_proto_enumTypes[0].Descriptor()
}

func (Asp_Service) Type() protoreflect.EnumType {
	return &file_asp_proto_enumTypes[0]
}

func (x Asp_Service) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Asp_Service.Descriptor instead.
func (Asp_Service) EnumDescriptor() ([]byte, []int) {
	return file_asp_proto_rawDescGZIP(), []int{0, 0}
}

type AspServiceRequest_AspService int32

const (
	AspServiceRequest_YOUTUBE   AspServiceRequest_AspService = 0
	AspServiceRequest_TWITTER   AspServiceRequest_AspService = 1
	AspServiceRequest_INSTAGRAM AspServiceRequest_AspService = 2
	AspServiceRequest_FACEBOOK  AspServiceRequest_AspService = 3
	AspServiceRequest_LINKEDIN  AspServiceRequest_AspService = 4
	AspServiceRequest_TIKTOK    AspServiceRequest_AspService = 5
	AspServiceRequest_PINTEREST AspServiceRequest_AspService = 6
	AspServiceRequest_REDDIT    AspServiceRequest_AspService = 7
	AspServiceRequest_SNAPCHAT  AspServiceRequest_AspService = 8
	AspServiceRequest_TUMBLR    AspServiceRequest_AspService = 9
	AspServiceRequest_TWITCH    AspServiceRequest_AspService = 10
	AspServiceRequest_WEIBO     AspServiceRequest_AspService = 11
	AspServiceRequest_WECHAT    AspServiceRequest_AspService = 12
	AspServiceRequest_WHATSAPP  AspServiceRequest_AspService = 13
	AspServiceRequest_LINE      AspServiceRequest_AspService = 14
	AspServiceRequest_TELEGRAM  AspServiceRequest_AspService = 15
	AspServiceRequest_VK        AspServiceRequest_AspService = 16
	AspServiceRequest_YAHOO     AspServiceRequest_AspService = 17
	AspServiceRequest_OTHER     AspServiceRequest_AspService = 18
)

// Enum value maps for AspServiceRequest_AspService.
var (
	AspServiceRequest_AspService_name = map[int32]string{
		0:  "YOUTUBE",
		1:  "TWITTER",
		2:  "INSTAGRAM",
		3:  "FACEBOOK",
		4:  "LINKEDIN",
		5:  "TIKTOK",
		6:  "PINTEREST",
		7:  "REDDIT",
		8:  "SNAPCHAT",
		9:  "TUMBLR",
		10: "TWITCH",
		11: "WEIBO",
		12: "WECHAT",
		13: "WHATSAPP",
		14: "LINE",
		15: "TELEGRAM",
		16: "VK",
		17: "YAHOO",
		18: "OTHER",
	}
	AspServiceRequest_AspService_value = map[string]int32{
		"YOUTUBE":   0,
		"TWITTER":   1,
		"INSTAGRAM": 2,
		"FACEBOOK":  3,
		"LINKEDIN":  4,
		"TIKTOK":    5,
		"PINTEREST": 6,
		"REDDIT":    7,
		"SNAPCHAT":  8,
		"TUMBLR":    9,
		"TWITCH":    10,
		"WEIBO":     11,
		"WECHAT":    12,
		"WHATSAPP":  13,
		"LINE":      14,
		"TELEGRAM":  15,
		"VK":        16,
		"YAHOO":     17,
		"OTHER":     18,
	}
)

func (x AspServiceRequest_AspService) Enum() *AspServiceRequest_AspService {
	p := new(AspServiceRequest_AspService)
	*p = x
	return p
}

func (x AspServiceRequest_AspService) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AspServiceRequest_AspService) Descriptor() protoreflect.EnumDescriptor {
	return file_asp_proto_enumTypes[1].Descriptor()
}

func (AspServiceRequest_AspService) Type() protoreflect.EnumType {
	return &file_asp_proto_enumTypes[1]
}

func (x AspServiceRequest_AspService) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AspServiceRequest_AspService.Descriptor instead.
func (AspServiceRequest_AspService) EnumDescriptor() ([]byte, []int) {
	return file_asp_proto_rawDescGZIP(), []int{4, 0}
}

// ユーザー情報を表すメッセージ型
type Asp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid      string        `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ItemId    int64         `protobuf:"varint,3,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	UserId    int64         `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Service   []Asp_Service `protobuf:"varint,6,rep,packed,name=service,proto3,enum=asp.Asp_Service" json:"service,omitempty"`
	CreatedAt string        `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string        `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsDeleted bool          `protobuf:"varint,9,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
}

func (x *Asp) Reset() {
	*x = Asp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asp) ProtoMessage() {}

func (x *Asp) ProtoReflect() protoreflect.Message {
	mi := &file_asp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asp.ProtoReflect.Descriptor instead.
func (*Asp) Descriptor() ([]byte, []int) {
	return file_asp_proto_rawDescGZIP(), []int{0}
}

func (x *Asp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Asp) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Asp) GetItemId() int64 {
	if x != nil {
		return x.ItemId
	}
	return 0
}

func (x *Asp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Asp) GetService() []Asp_Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *Asp) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Asp) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Asp) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

type AspList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Asp []*Asp `protobuf:"bytes,1,rep,name=Asp,proto3" json:"Asp,omitempty"`
}

func (x *AspList) Reset() {
	*x = AspList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AspList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AspList) ProtoMessage() {}

func (x *AspList) ProtoReflect() protoreflect.Message {
	mi := &file_asp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AspList.ProtoReflect.Descriptor instead.
func (*AspList) Descriptor() ([]byte, []int) {
	return file_asp_proto_rawDescGZIP(), []int{1}
}

func (x *AspList) GetAsp() []*Asp {
	if x != nil {
		return x.Asp
	}
	return nil
}

// request
type AspIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AspIdRequest) Reset() {
	*x = AspIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AspIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AspIdRequest) ProtoMessage() {}

func (x *AspIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AspIdRequest.ProtoReflect.Descriptor instead.
func (*AspIdRequest) Descriptor() ([]byte, []int) {
	return file_asp_proto_rawDescGZIP(), []int{2}
}

func (x *AspIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AspIdListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *AspIdListRequest) Reset() {
	*x = AspIdListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AspIdListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AspIdListRequest) ProtoMessage() {}

func (x *AspIdListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AspIdListRequest.ProtoReflect.Descriptor instead.
func (*AspIdListRequest) Descriptor() ([]byte, []int) {
	return file_asp_proto_rawDescGZIP(), []int{3}
}

func (x *AspIdListRequest) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

// ユーザーの情報リクエストに使用するメッセージ型
type AspServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AspType AspServiceRequest_AspService `protobuf:"varint,1,opt,name=Asp_type,json=AspType,proto3,enum=asp.AspServiceRequest_AspService" json:"Asp_type,omitempty"`
}

func (x *AspServiceRequest) Reset() {
	*x = AspServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AspServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AspServiceRequest) ProtoMessage() {}

func (x *AspServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_asp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AspServiceRequest.ProtoReflect.Descriptor instead.
func (*AspServiceRequest) Descriptor() ([]byte, []int) {
	return file_asp_proto_rawDescGZIP(), []int{4}
}

func (x *AspServiceRequest) GetAspType() AspServiceRequest_AspService {
	if x != nil {
		return x.AspType
	}
	return AspServiceRequest_YOUTUBE
}

// response
type AspBoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *AspBoolResponse) Reset() {
	*x = AspBoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_asp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AspBoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AspBoolResponse) ProtoMessage() {}

func (x *AspBoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_asp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AspBoolResponse.ProtoReflect.Descriptor instead.
func (*AspBoolResponse) Descriptor() ([]byte, []int) {
	return file_asp_proto_rawDescGZIP(), []int{5}
}

func (x *AspBoolResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

var File_asp_proto protoreflect.FileDescriptor

var file_asp_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x73, 0x70,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x03,
	0x0a, 0x03, 0x41, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61,
	0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x22, 0xf6, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x59, 0x4f, 0x55, 0x54, 0x55, 0x42, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x54, 0x57, 0x49, 0x54, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e,
	0x53, 0x54, 0x41, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x41, 0x43,
	0x45, 0x42, 0x4f, 0x4f, 0x4b, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x49, 0x4e, 0x4b, 0x45,
	0x44, 0x49, 0x4e, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x49, 0x4b, 0x54, 0x4f, 0x4b, 0x10,
	0x05, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x45, 0x53, 0x54, 0x10, 0x06,
	0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x44, 0x44, 0x49, 0x54, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08,
	0x53, 0x4e, 0x41, 0x50, 0x43, 0x48, 0x41, 0x54, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x55,
	0x4d, 0x42, 0x4c, 0x52, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x57, 0x49, 0x54, 0x43, 0x48,
	0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x45, 0x49, 0x42, 0x4f, 0x10, 0x0b, 0x12, 0x0a, 0x0a,
	0x06, 0x57, 0x45, 0x43, 0x48, 0x41, 0x54, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x48, 0x41,
	0x54, 0x53, 0x41, 0x50, 0x50, 0x10, 0x0d, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4e, 0x45, 0x10,
	0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x45, 0x4c, 0x45, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x0f, 0x12,
	0x06, 0x0a, 0x02, 0x56, 0x4b, 0x10, 0x10, 0x12, 0x09, 0x0a, 0x05, 0x59, 0x41, 0x48, 0x4f, 0x4f,
	0x10, 0x11, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x12, 0x22, 0x25, 0x0a,
	0x07, 0x41, 0x73, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x41, 0x73, 0x70, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x61, 0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x52,
	0x03, 0x41, 0x73, 0x70, 0x22, 0x1e, 0x0a, 0x0c, 0x41, 0x73, 0x70, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x41, 0x73, 0x70, 0x49, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0xcd, 0x02, 0x0a, 0x11, 0x41, 0x73, 0x70,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c,
	0x0a, 0x08, 0x41, 0x73, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x21, 0x2e, 0x61, 0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x73, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x07, 0x41, 0x73, 0x70, 0x54, 0x79, 0x70, 0x65, 0x22, 0xf9, 0x01, 0x0a,
	0x0a, 0x41, 0x73, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x59,
	0x4f, 0x55, 0x54, 0x55, 0x42, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x57, 0x49, 0x54,
	0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x47, 0x52,
	0x41, 0x4d, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x41, 0x43, 0x45, 0x42, 0x4f, 0x4f, 0x4b,
	0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x49, 0x4e, 0x4b, 0x45, 0x44, 0x49, 0x4e, 0x10, 0x04,
	0x12, 0x0a, 0x0a, 0x06, 0x54, 0x49, 0x4b, 0x54, 0x4f, 0x4b, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09,
	0x50, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x45, 0x53, 0x54, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x52,
	0x45, 0x44, 0x44, 0x49, 0x54, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x4e, 0x41, 0x50, 0x43,
	0x48, 0x41, 0x54, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x55, 0x4d, 0x42, 0x4c, 0x52, 0x10,
	0x09, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x57, 0x49, 0x54, 0x43, 0x48, 0x10, 0x0a, 0x12, 0x09, 0x0a,
	0x05, 0x57, 0x45, 0x49, 0x42, 0x4f, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x45, 0x43, 0x48,
	0x41, 0x54, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x57, 0x48, 0x41, 0x54, 0x53, 0x41, 0x50, 0x50,
	0x10, 0x0d, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08,
	0x54, 0x45, 0x4c, 0x45, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x0f, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x4b,
	0x10, 0x10, 0x12, 0x09, 0x0a, 0x05, 0x59, 0x41, 0x48, 0x4f, 0x4f, 0x10, 0x11, 0x12, 0x09, 0x0a,
	0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x12, 0x22, 0x27, 0x0a, 0x0f, 0x41, 0x73, 0x70, 0x42,
	0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x32, 0xc9, 0x03, 0x0a, 0x0a, 0x41, 0x73, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x1e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x61, 0x73, 0x70,
	0x2e, 0x41, 0x73, 0x70, 0x1a, 0x08, 0x2e, 0x61, 0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x2a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x61, 0x73, 0x70,
	0x2e, 0x41, 0x73, 0x70, 0x1a, 0x14, 0x2e, 0x61, 0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x42, 0x6f,
	0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x61, 0x73, 0x70, 0x2e, 0x41, 0x73, 0x70,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x73, 0x70, 0x2e,
	0x41, 0x73, 0x70, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x28, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x61,
	0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x08, 0x2e, 0x61, 0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x61,
	0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x61, 0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12,
	0x32, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x74, 0x65, 0x6d,
	0x12, 0x11, 0x2e, 0x61, 0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x61, 0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x73, 0x70, 0x2e, 0x41, 0x73,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x61, 0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x49, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x61, 0x73, 0x70, 0x2e,
	0x41, 0x73, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x28, 0x01, 0x12, 0x30, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e,
	0x61, 0x73, 0x70, 0x2e, 0x41, 0x73, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x11, 0x5a,
	0x0f, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_asp_proto_rawDescOnce sync.Once
	file_asp_proto_rawDescData = file_asp_proto_rawDesc
)

func file_asp_proto_rawDescGZIP() []byte {
	file_asp_proto_rawDescOnce.Do(func() {
		file_asp_proto_rawDescData = protoimpl.X.CompressGZIP(file_asp_proto_rawDescData)
	})
	return file_asp_proto_rawDescData
}

var file_asp_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_asp_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_asp_proto_goTypes = []interface{}{
	(Asp_Service)(0),                  // 0: asp.Asp.Service
	(AspServiceRequest_AspService)(0), // 1: asp.AspServiceRequest.AspService
	(*Asp)(nil),                       // 2: asp.Asp
	(*AspList)(nil),                   // 3: asp.AspList
	(*AspIdRequest)(nil),              // 4: asp.AspIdRequest
	(*AspIdListRequest)(nil),          // 5: asp.AspIdListRequest
	(*AspServiceRequest)(nil),         // 6: asp.AspServiceRequest
	(*AspBoolResponse)(nil),           // 7: asp.AspBoolResponse
	(*emptypb.Empty)(nil),             // 8: google.protobuf.Empty
}
var file_asp_proto_depIdxs = []int32{
	0,  // 0: asp.Asp.service:type_name -> asp.Asp.Service
	2,  // 1: asp.AspList.Asp:type_name -> asp.Asp
	1,  // 2: asp.AspServiceRequest.Asp_type:type_name -> asp.AspServiceRequest.AspService
	2,  // 3: asp.AspService.Create:input_type -> asp.Asp
	2,  // 4: asp.AspService.Update:input_type -> asp.Asp
	4,  // 5: asp.AspService.Delete:input_type -> asp.AspIdRequest
	4,  // 6: asp.AspService.GetById:input_type -> asp.AspIdRequest
	4,  // 7: asp.AspService.GetListByUser:input_type -> asp.AspIdRequest
	4,  // 8: asp.AspService.GetListByItem:input_type -> asp.AspIdRequest
	6,  // 9: asp.AspService.GetListByService:input_type -> asp.AspServiceRequest
	5,  // 10: asp.AspService.GetListByIdList:input_type -> asp.AspIdListRequest
	8,  // 11: asp.AspService.GetAll:input_type -> google.protobuf.Empty
	2,  // 12: asp.AspService.Create:output_type -> asp.Asp
	7,  // 13: asp.AspService.Update:output_type -> asp.AspBoolResponse
	7,  // 14: asp.AspService.Delete:output_type -> asp.AspBoolResponse
	2,  // 15: asp.AspService.GetById:output_type -> asp.Asp
	3,  // 16: asp.AspService.GetListByUser:output_type -> asp.AspList
	3,  // 17: asp.AspService.GetListByItem:output_type -> asp.AspList
	3,  // 18: asp.AspService.GetListByService:output_type -> asp.AspList
	3,  // 19: asp.AspService.GetListByIdList:output_type -> asp.AspList
	3,  // 20: asp.AspService.GetAll:output_type -> asp.AspList
	12, // [12:21] is the sub-list for method output_type
	3,  // [3:12] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_asp_proto_init() }
func file_asp_proto_init() {
	if File_asp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_asp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Asp); i {
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
		file_asp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AspList); i {
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
		file_asp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AspIdRequest); i {
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
		file_asp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AspIdListRequest); i {
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
		file_asp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AspServiceRequest); i {
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
		file_asp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AspBoolResponse); i {
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
			RawDescriptor: file_asp_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_asp_proto_goTypes,
		DependencyIndexes: file_asp_proto_depIdxs,
		EnumInfos:         file_asp_proto_enumTypes,
		MessageInfos:      file_asp_proto_msgTypes,
	}.Build()
	File_asp_proto = out.File
	file_asp_proto_rawDesc = nil
	file_asp_proto_goTypes = nil
	file_asp_proto_depIdxs = nil
}
