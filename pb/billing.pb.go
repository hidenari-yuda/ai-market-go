// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: billing.proto

// Specify package name to avoid name collision

package pb

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

type Billing_Service int32

const (
	Billing_NONE     Billing_Service = 0
	Billing_NORMAL   Billing_Service = 1
	Billing_ADMIN    Billing_Service = 2
	Billing_GUEST    Billing_Service = 3
	Billing_DISABLED Billing_Service = 4
)

// Enum value maps for Billing_Service.
var (
	Billing_Service_name = map[int32]string{
		0: "NONE",
		1: "NORMAL",
		2: "ADMIN",
		3: "GUEST",
		4: "DISABLED",
	}
	Billing_Service_value = map[string]int32{
		"NONE":     0,
		"NORMAL":   1,
		"ADMIN":    2,
		"GUEST":    3,
		"DISABLED": 4,
	}
)

func (x Billing_Service) Enum() *Billing_Service {
	p := new(Billing_Service)
	*p = x
	return p
}

func (x Billing_Service) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Billing_Service) Descriptor() protoreflect.EnumDescriptor {
	return file_billing_proto_enumTypes[0].Descriptor()
}

func (Billing_Service) Type() protoreflect.EnumType {
	return &file_billing_proto_enumTypes[0]
}

func (x Billing_Service) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Billing_Service.Descriptor instead.
func (Billing_Service) EnumDescriptor() ([]byte, []int) {
	return file_billing_proto_rawDescGZIP(), []int{0, 0}
}

// ユーザー情報を表すメッセージ型
type Billing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid      string          `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserId    int64           `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title     string          `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Number    string          `protobuf:"bytes,5,opt,name=number,proto3" json:"number,omitempty"`
	Password  string          `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Service   Billing_Service `protobuf:"varint,7,opt,name=service,proto3,enum=billing.Billing_Service" json:"service,omitempty"`
	CreatedAt string          `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string          `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsDeleted bool            `protobuf:"varint,10,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
}

func (x *Billing) Reset() {
	*x = Billing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Billing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Billing) ProtoMessage() {}

func (x *Billing) ProtoReflect() protoreflect.Message {
	mi := &file_billing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Billing.ProtoReflect.Descriptor instead.
func (*Billing) Descriptor() ([]byte, []int) {
	return file_billing_proto_rawDescGZIP(), []int{0}
}

func (x *Billing) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Billing) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Billing) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Billing) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Billing) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Billing) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Billing) GetService() Billing_Service {
	if x != nil {
		return x.Service
	}
	return Billing_NONE
}

func (x *Billing) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Billing) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Billing) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

type BillingList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Billing []*Billing `protobuf:"bytes,1,rep,name=billing,proto3" json:"billing,omitempty"`
}

func (x *BillingList) Reset() {
	*x = BillingList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingList) ProtoMessage() {}

func (x *BillingList) ProtoReflect() protoreflect.Message {
	mi := &file_billing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingList.ProtoReflect.Descriptor instead.
func (*BillingList) Descriptor() ([]byte, []int) {
	return file_billing_proto_rawDescGZIP(), []int{1}
}

func (x *BillingList) GetBilling() []*Billing {
	if x != nil {
		return x.Billing
	}
	return nil
}

// request
type BillingIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BillingIdRequest) Reset() {
	*x = BillingIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingIdRequest) ProtoMessage() {}

func (x *BillingIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_billing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingIdRequest.ProtoReflect.Descriptor instead.
func (*BillingIdRequest) Descriptor() ([]byte, []int) {
	return file_billing_proto_rawDescGZIP(), []int{2}
}

func (x *BillingIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BillingIdListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *BillingIdListRequest) Reset() {
	*x = BillingIdListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingIdListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingIdListRequest) ProtoMessage() {}

func (x *BillingIdListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_billing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingIdListRequest.ProtoReflect.Descriptor instead.
func (*BillingIdListRequest) Descriptor() ([]byte, []int) {
	return file_billing_proto_rawDescGZIP(), []int{3}
}

func (x *BillingIdListRequest) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

// res
type BillingBoolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *BillingBoolResponse) Reset() {
	*x = BillingBoolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_billing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BillingBoolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingBoolResponse) ProtoMessage() {}

func (x *BillingBoolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_billing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingBoolResponse.ProtoReflect.Descriptor instead.
func (*BillingBoolResponse) Descriptor() ([]byte, []int) {
	return file_billing_proto_rawDescGZIP(), []int{4}
}

func (x *BillingBoolResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

var File_billing_proto protoreflect.FileDescriptor

var file_billing_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0xe6, 0x02, 0x0a, 0x07, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x32, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x43, 0x0a, 0x07,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x55, 0x45, 0x53,
	0x54, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10,
	0x04, 0x22, 0x39, 0x0a, 0x0b, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2a, 0x0a, 0x07, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x52, 0x07, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x22, 0x0a, 0x10,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x26, 0x0a, 0x14, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x13, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xf8, 0x01, 0x0a, 0x0e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x10, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x10, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x10, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x1c, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x19, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x67, 0x6f,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_billing_proto_rawDescOnce sync.Once
	file_billing_proto_rawDescData = file_billing_proto_rawDesc
)

func file_billing_proto_rawDescGZIP() []byte {
	file_billing_proto_rawDescOnce.Do(func() {
		file_billing_proto_rawDescData = protoimpl.X.CompressGZIP(file_billing_proto_rawDescData)
	})
	return file_billing_proto_rawDescData
}

var file_billing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_billing_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_billing_proto_goTypes = []interface{}{
	(Billing_Service)(0),         // 0: billing.Billing.Service
	(*Billing)(nil),              // 1: billing.Billing
	(*BillingList)(nil),          // 2: billing.BillingList
	(*BillingIdRequest)(nil),     // 3: billing.BillingIdRequest
	(*BillingIdListRequest)(nil), // 4: billing.BillingIdListRequest
	(*BillingBoolResponse)(nil),  // 5: billing.BillingBoolResponse
}
var file_billing_proto_depIdxs = []int32{
	0, // 0: billing.Billing.service:type_name -> billing.Billing.Service
	1, // 1: billing.BillingList.billing:type_name -> billing.Billing
	1, // 2: billing.BillingService.Create:input_type -> billing.Billing
	1, // 3: billing.BillingService.Update:input_type -> billing.Billing
	3, // 4: billing.BillingService.GetById:input_type -> billing.BillingIdRequest
	3, // 5: billing.BillingService.GetByUserId:input_type -> billing.BillingIdRequest
	1, // 6: billing.BillingService.Create:output_type -> billing.Billing
	5, // 7: billing.BillingService.Update:output_type -> billing.BillingBoolResponse
	1, // 8: billing.BillingService.GetById:output_type -> billing.Billing
	2, // 9: billing.BillingService.GetByUserId:output_type -> billing.BillingList
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_billing_proto_init() }
func file_billing_proto_init() {
	if File_billing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_billing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Billing); i {
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
		file_billing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingList); i {
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
		file_billing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingIdRequest); i {
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
		file_billing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingIdListRequest); i {
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
		file_billing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BillingBoolResponse); i {
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
			RawDescriptor: file_billing_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_billing_proto_goTypes,
		DependencyIndexes: file_billing_proto_depIdxs,
		EnumInfos:         file_billing_proto_enumTypes,
		MessageInfos:      file_billing_proto_msgTypes,
	}.Build()
	File_billing_proto = out.File
	file_billing_proto_rawDesc = nil
	file_billing_proto_goTypes = nil
	file_billing_proto_depIdxs = nil
}
