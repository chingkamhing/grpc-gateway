// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: gateway/gateway_rr.proto

package tm2_proto_gateway_go

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *UserDetail `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_rr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_rr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_rr_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetValue() *UserDetail {
	if x != nil {
		return x.Value
	}
	return nil
}

type CreateUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *UserDetail `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Error string      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *CreateUserReply) Reset() {
	*x = CreateUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_rr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReply) ProtoMessage() {}

func (x *CreateUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_rr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserReply.ProtoReflect.Descriptor instead.
func (*CreateUserReply) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_rr_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserReply) GetValue() *UserDetail {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *CreateUserReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ListUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListUserRequest) Reset() {
	*x = ListUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_rr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRequest) ProtoMessage() {}

func (x *ListUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_rr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRequest.ProtoReflect.Descriptor instead.
func (*ListUserRequest) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_rr_proto_rawDescGZIP(), []int{2}
}

func (x *ListUserRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListUserRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*UserDetail `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	Offset int32         `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32         `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Count  int32         `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Error  string        `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ListUserReply) Reset() {
	*x = ListUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_rr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserReply) ProtoMessage() {}

func (x *ListUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_rr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserReply.ProtoReflect.Descriptor instead.
func (*ListUserReply) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_rr_proto_rawDescGZIP(), []int{3}
}

func (x *ListUserReply) GetValues() []*UserDetail {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *ListUserReply) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListUserReply) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListUserReply) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListUserReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_rr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_rr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_rr_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *UserDetail `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetUserReply) Reset() {
	*x = GetUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_rr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserReply) ProtoMessage() {}

func (x *GetUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_rr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserReply.ProtoReflect.Descriptor instead.
func (*GetUserReply) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_rr_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserReply) GetValue() *UserDetail {
	if x != nil {
		return x.Value
	}
	return nil
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *UserDetail `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_rr_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_rr_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_rr_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateUserRequest) GetValue() *UserDetail {
	if x != nil {
		return x.Value
	}
	return nil
}

type UpdateUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *UserDetail `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateUserReply) Reset() {
	*x = UpdateUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_rr_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserReply) ProtoMessage() {}

func (x *UpdateUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_rr_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserReply.ProtoReflect.Descriptor instead.
func (*UpdateUserReply) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_rr_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateUserReply) GetValue() *UserDetail {
	if x != nil {
		return x.Value
	}
	return nil
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_rr_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_rr_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_rr_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteUserRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteUserReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *DeleteUserReply) Reset() {
	*x = DeleteUserReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_rr_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserReply) ProtoMessage() {}

func (x *DeleteUserReply) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_rr_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserReply.ProtoReflect.Descriptor instead.
func (*DeleteUserReply) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_rr_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteUserReply) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_gateway_gateway_rr_proto protoreflect.FileDescriptor

var file_gateway_gateway_rr_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x72, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x1a, 0x1c, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x52, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3f, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x96, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x39, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3e, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3c, 0x0a, 0x0f,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x29, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x27, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x6b, 0x61, 0x6d, 0x68,
	0x69, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x74, 0x6d, 0x32, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x67, 0x6f, 0x3b, 0x74, 0x6d, 0x32, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x67, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_gateway_rr_proto_rawDescOnce sync.Once
	file_gateway_gateway_rr_proto_rawDescData = file_gateway_gateway_rr_proto_rawDesc
)

func file_gateway_gateway_rr_proto_rawDescGZIP() []byte {
	file_gateway_gateway_rr_proto_rawDescOnce.Do(func() {
		file_gateway_gateway_rr_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_gateway_rr_proto_rawDescData)
	})
	return file_gateway_gateway_rr_proto_rawDescData
}

var file_gateway_gateway_rr_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_gateway_gateway_rr_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil), // 0: gateway.CreateUserRequest
	(*CreateUserReply)(nil),   // 1: gateway.CreateUserReply
	(*ListUserRequest)(nil),   // 2: gateway.ListUserRequest
	(*ListUserReply)(nil),     // 3: gateway.ListUserReply
	(*GetUserRequest)(nil),    // 4: gateway.GetUserRequest
	(*GetUserReply)(nil),      // 5: gateway.GetUserReply
	(*UpdateUserRequest)(nil), // 6: gateway.UpdateUserRequest
	(*UpdateUserReply)(nil),   // 7: gateway.UpdateUserReply
	(*DeleteUserRequest)(nil), // 8: gateway.DeleteUserRequest
	(*DeleteUserReply)(nil),   // 9: gateway.DeleteUserReply
	(*UserDetail)(nil),        // 10: gateway.UserDetail
}
var file_gateway_gateway_rr_proto_depIdxs = []int32{
	10, // 0: gateway.CreateUserRequest.value:type_name -> gateway.UserDetail
	10, // 1: gateway.CreateUserReply.value:type_name -> gateway.UserDetail
	10, // 2: gateway.ListUserReply.values:type_name -> gateway.UserDetail
	10, // 3: gateway.GetUserReply.value:type_name -> gateway.UserDetail
	10, // 4: gateway.UpdateUserRequest.value:type_name -> gateway.UserDetail
	10, // 5: gateway.UpdateUserReply.value:type_name -> gateway.UserDetail
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_gateway_gateway_rr_proto_init() }
func file_gateway_gateway_rr_proto_init() {
	if File_gateway_gateway_rr_proto != nil {
		return
	}
	file_gateway_gateway_struct_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gateway_gateway_rr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_gateway_gateway_rr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserReply); i {
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
		file_gateway_gateway_rr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserRequest); i {
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
		file_gateway_gateway_rr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUserReply); i {
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
		file_gateway_gateway_rr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_gateway_gateway_rr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserReply); i {
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
		file_gateway_gateway_rr_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRequest); i {
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
		file_gateway_gateway_rr_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserReply); i {
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
		file_gateway_gateway_rr_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
		file_gateway_gateway_rr_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserReply); i {
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
			RawDescriptor: file_gateway_gateway_rr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_gateway_rr_proto_goTypes,
		DependencyIndexes: file_gateway_gateway_rr_proto_depIdxs,
		MessageInfos:      file_gateway_gateway_rr_proto_msgTypes,
	}.Build()
	File_gateway_gateway_rr_proto = out.File
	file_gateway_gateway_rr_proto_rawDesc = nil
	file_gateway_gateway_rr_proto_goTypes = nil
	file_gateway_gateway_rr_proto_depIdxs = nil
}
