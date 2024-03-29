// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: gateway/gateway_struct.proto

package tm2_proto_gateway_go

import (
	srv_proto_company_go "github.com/chingkamhing/grpc-gateway/gen/srv-proto-company-go"
	srv_proto_user_go "github.com/chingkamhing/grpc-gateway/gen/srv-proto-user-go"
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

type UserDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    *srv_proto_user_go.UserInfo       `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Company *srv_proto_company_go.CompanyInfo `protobuf:"bytes,2,opt,name=company,proto3" json:"company,omitempty"`
}

func (x *UserDetail) Reset() {
	*x = UserDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_struct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetail) ProtoMessage() {}

func (x *UserDetail) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_struct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetail.ProtoReflect.Descriptor instead.
func (*UserDetail) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_struct_proto_rawDescGZIP(), []int{0}
}

func (x *UserDetail) GetUser() *srv_proto_user_go.UserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserDetail) GetCompany() *srv_proto_company_go.CompanyInfo {
	if x != nil {
		return x.Company
	}
	return nil
}

var File_gateway_gateway_struct_proto protoreflect.FileDescriptor

var file_gateway_gateway_struct_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0x16, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79,
	0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a,
	0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x22, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x2e, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x6e, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x42,
	0x54, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68,
	0x69, 0x6e, 0x67, 0x6b, 0x61, 0x6d, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x72, 0x76, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x67, 0x6f,
	0x3b, 0x74, 0x6d, 0x32, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x5f, 0x67, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_gateway_struct_proto_rawDescOnce sync.Once
	file_gateway_gateway_struct_proto_rawDescData = file_gateway_gateway_struct_proto_rawDesc
)

func file_gateway_gateway_struct_proto_rawDescGZIP() []byte {
	file_gateway_gateway_struct_proto_rawDescOnce.Do(func() {
		file_gateway_gateway_struct_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_gateway_struct_proto_rawDescData)
	})
	return file_gateway_gateway_struct_proto_rawDescData
}

var file_gateway_gateway_struct_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gateway_gateway_struct_proto_goTypes = []interface{}{
	(*UserDetail)(nil),                       // 0: gateway.UserDetail
	(*srv_proto_user_go.UserInfo)(nil),       // 1: user.UserInfo
	(*srv_proto_company_go.CompanyInfo)(nil), // 2: company.CompanyInfo
}
var file_gateway_gateway_struct_proto_depIdxs = []int32{
	1, // 0: gateway.UserDetail.user:type_name -> user.UserInfo
	2, // 1: gateway.UserDetail.company:type_name -> company.CompanyInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gateway_gateway_struct_proto_init() }
func file_gateway_gateway_struct_proto_init() {
	if File_gateway_gateway_struct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_gateway_struct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDetail); i {
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
			RawDescriptor: file_gateway_gateway_struct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gateway_gateway_struct_proto_goTypes,
		DependencyIndexes: file_gateway_gateway_struct_proto_depIdxs,
		MessageInfos:      file_gateway_gateway_struct_proto_msgTypes,
	}.Build()
	File_gateway_gateway_struct_proto = out.File
	file_gateway_gateway_struct_proto_rawDesc = nil
	file_gateway_gateway_struct_proto_goTypes = nil
	file_gateway_gateway_struct_proto_depIdxs = nil
}
