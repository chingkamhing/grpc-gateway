// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: company/company_rr.proto

package tm2_proto_company_go

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

type CreateCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *CompanyInfo `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CreateCompanyRequest) Reset() {
	*x = CreateCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_company_rr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCompanyRequest) ProtoMessage() {}

func (x *CreateCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_company_rr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCompanyRequest.ProtoReflect.Descriptor instead.
func (*CreateCompanyRequest) Descriptor() ([]byte, []int) {
	return file_company_company_rr_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCompanyRequest) GetValue() *CompanyInfo {
	if x != nil {
		return x.Value
	}
	return nil
}

type CreateCompanyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *CompanyInfo `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CreateCompanyReply) Reset() {
	*x = CreateCompanyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_company_rr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCompanyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCompanyReply) ProtoMessage() {}

func (x *CreateCompanyReply) ProtoReflect() protoreflect.Message {
	mi := &file_company_company_rr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCompanyReply.ProtoReflect.Descriptor instead.
func (*CreateCompanyReply) Descriptor() ([]byte, []int) {
	return file_company_company_rr_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCompanyReply) GetValue() *CompanyInfo {
	if x != nil {
		return x.Value
	}
	return nil
}

type ListCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListCompanyRequest) Reset() {
	*x = ListCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_company_rr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCompanyRequest) ProtoMessage() {}

func (x *ListCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_company_rr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCompanyRequest.ProtoReflect.Descriptor instead.
func (*ListCompanyRequest) Descriptor() ([]byte, []int) {
	return file_company_company_rr_proto_rawDescGZIP(), []int{2}
}

func (x *ListCompanyRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListCompanyRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListCompanyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*CompanyInfo `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	Offset int32          `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32          `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Count  int32          `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListCompanyReply) Reset() {
	*x = ListCompanyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_company_rr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCompanyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCompanyReply) ProtoMessage() {}

func (x *ListCompanyReply) ProtoReflect() protoreflect.Message {
	mi := &file_company_company_rr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCompanyReply.ProtoReflect.Descriptor instead.
func (*ListCompanyReply) Descriptor() ([]byte, []int) {
	return file_company_company_rr_proto_rawDescGZIP(), []int{3}
}

func (x *ListCompanyReply) GetValues() []*CompanyInfo {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *ListCompanyReply) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListCompanyReply) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListCompanyReply) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCompanyRequest) Reset() {
	*x = GetCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_company_rr_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyRequest) ProtoMessage() {}

func (x *GetCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_company_rr_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyRequest.ProtoReflect.Descriptor instead.
func (*GetCompanyRequest) Descriptor() ([]byte, []int) {
	return file_company_company_rr_proto_rawDescGZIP(), []int{4}
}

func (x *GetCompanyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCompanyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *CompanyInfo `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetCompanyReply) Reset() {
	*x = GetCompanyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_company_rr_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCompanyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCompanyReply) ProtoMessage() {}

func (x *GetCompanyReply) ProtoReflect() protoreflect.Message {
	mi := &file_company_company_rr_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCompanyReply.ProtoReflect.Descriptor instead.
func (*GetCompanyReply) Descriptor() ([]byte, []int) {
	return file_company_company_rr_proto_rawDescGZIP(), []int{5}
}

func (x *GetCompanyReply) GetValue() *CompanyInfo {
	if x != nil {
		return x.Value
	}
	return nil
}

type UpdateCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *CompanyInfo `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateCompanyRequest) Reset() {
	*x = UpdateCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_company_rr_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCompanyRequest) ProtoMessage() {}

func (x *UpdateCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_company_rr_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCompanyRequest.ProtoReflect.Descriptor instead.
func (*UpdateCompanyRequest) Descriptor() ([]byte, []int) {
	return file_company_company_rr_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCompanyRequest) GetValue() *CompanyInfo {
	if x != nil {
		return x.Value
	}
	return nil
}

type UpdateCompanyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *CompanyInfo `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UpdateCompanyReply) Reset() {
	*x = UpdateCompanyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_company_rr_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCompanyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCompanyReply) ProtoMessage() {}

func (x *UpdateCompanyReply) ProtoReflect() protoreflect.Message {
	mi := &file_company_company_rr_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCompanyReply.ProtoReflect.Descriptor instead.
func (*UpdateCompanyReply) Descriptor() ([]byte, []int) {
	return file_company_company_rr_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCompanyReply) GetValue() *CompanyInfo {
	if x != nil {
		return x.Value
	}
	return nil
}

type DeleteCompanyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCompanyRequest) Reset() {
	*x = DeleteCompanyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_company_rr_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCompanyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCompanyRequest) ProtoMessage() {}

func (x *DeleteCompanyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_company_company_rr_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCompanyRequest.ProtoReflect.Descriptor instead.
func (*DeleteCompanyRequest) Descriptor() ([]byte, []int) {
	return file_company_company_rr_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCompanyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteCompanyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCompanyReply) Reset() {
	*x = DeleteCompanyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_company_company_rr_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCompanyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCompanyReply) ProtoMessage() {}

func (x *DeleteCompanyReply) ProtoReflect() protoreflect.Message {
	mi := &file_company_company_rr_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCompanyReply.ProtoReflect.Descriptor instead.
func (*DeleteCompanyReply) Descriptor() ([]byte, []int) {
	return file_company_company_rr_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCompanyReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_company_company_rr_proto protoreflect.FileDescriptor

var file_company_company_rr_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x5f, 0x72, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x42, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x40, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x42, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x84, 0x01, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x2c, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x42, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x40, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x26, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x54, 0x5a, 0x52, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x6b, 0x61,
	0x6d, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x72, 0x76, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2d, 0x67, 0x6f, 0x3b, 0x74, 0x6d, 0x32, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x5f, 0x67, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_company_company_rr_proto_rawDescOnce sync.Once
	file_company_company_rr_proto_rawDescData = file_company_company_rr_proto_rawDesc
)

func file_company_company_rr_proto_rawDescGZIP() []byte {
	file_company_company_rr_proto_rawDescOnce.Do(func() {
		file_company_company_rr_proto_rawDescData = protoimpl.X.CompressGZIP(file_company_company_rr_proto_rawDescData)
	})
	return file_company_company_rr_proto_rawDescData
}

var file_company_company_rr_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_company_company_rr_proto_goTypes = []interface{}{
	(*CreateCompanyRequest)(nil), // 0: company.CreateCompanyRequest
	(*CreateCompanyReply)(nil),   // 1: company.CreateCompanyReply
	(*ListCompanyRequest)(nil),   // 2: company.ListCompanyRequest
	(*ListCompanyReply)(nil),     // 3: company.ListCompanyReply
	(*GetCompanyRequest)(nil),    // 4: company.GetCompanyRequest
	(*GetCompanyReply)(nil),      // 5: company.GetCompanyReply
	(*UpdateCompanyRequest)(nil), // 6: company.UpdateCompanyRequest
	(*UpdateCompanyReply)(nil),   // 7: company.UpdateCompanyReply
	(*DeleteCompanyRequest)(nil), // 8: company.DeleteCompanyRequest
	(*DeleteCompanyReply)(nil),   // 9: company.DeleteCompanyReply
	(*CompanyInfo)(nil),          // 10: company.CompanyInfo
}
var file_company_company_rr_proto_depIdxs = []int32{
	10, // 0: company.CreateCompanyRequest.value:type_name -> company.CompanyInfo
	10, // 1: company.CreateCompanyReply.value:type_name -> company.CompanyInfo
	10, // 2: company.ListCompanyReply.values:type_name -> company.CompanyInfo
	10, // 3: company.GetCompanyReply.value:type_name -> company.CompanyInfo
	10, // 4: company.UpdateCompanyRequest.value:type_name -> company.CompanyInfo
	10, // 5: company.UpdateCompanyReply.value:type_name -> company.CompanyInfo
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_company_company_rr_proto_init() }
func file_company_company_rr_proto_init() {
	if File_company_company_rr_proto != nil {
		return
	}
	file_company_company_struct_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_company_company_rr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCompanyRequest); i {
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
		file_company_company_rr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCompanyReply); i {
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
		file_company_company_rr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCompanyRequest); i {
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
		file_company_company_rr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCompanyReply); i {
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
		file_company_company_rr_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompanyRequest); i {
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
		file_company_company_rr_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCompanyReply); i {
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
		file_company_company_rr_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCompanyRequest); i {
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
		file_company_company_rr_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCompanyReply); i {
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
		file_company_company_rr_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCompanyRequest); i {
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
		file_company_company_rr_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCompanyReply); i {
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
			RawDescriptor: file_company_company_rr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_company_company_rr_proto_goTypes,
		DependencyIndexes: file_company_company_rr_proto_depIdxs,
		MessageInfos:      file_company_company_rr_proto_msgTypes,
	}.Build()
	File_company_company_rr_proto = out.File
	file_company_company_rr_proto_rawDesc = nil
	file_company_company_rr_proto_goTypes = nil
	file_company_company_rr_proto_depIdxs = nil
}
