// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: protoc/admin.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeleteCorruptedFilesV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCorruptedFilesV1Request) Reset() {
	*x = DeleteCorruptedFilesV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCorruptedFilesV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCorruptedFilesV1Request) ProtoMessage() {}

func (x *DeleteCorruptedFilesV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCorruptedFilesV1Request.ProtoReflect.Descriptor instead.
func (*DeleteCorruptedFilesV1Request) Descriptor() ([]byte, []int) {
	return file_protoc_admin_proto_rawDescGZIP(), []int{0}
}

type DeleteCorruptedFilesV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCorruptedFilesV1Response) Reset() {
	*x = DeleteCorruptedFilesV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCorruptedFilesV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCorruptedFilesV1Response) ProtoMessage() {}

func (x *DeleteCorruptedFilesV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCorruptedFilesV1Response.ProtoReflect.Descriptor instead.
func (*DeleteCorruptedFilesV1Response) Descriptor() ([]byte, []int) {
	return file_protoc_admin_proto_rawDescGZIP(), []int{1}
}

var File_protoc_admin_proto protoreflect.FileDescriptor

var file_protoc_admin_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x64, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x20, 0x0a, 0x1e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x91, 0x01, 0x0a, 0x10, 0x46, 0x69, 0x6c,
	0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7d, 0x0a,
	0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x56, 0x31, 0x12, 0x30, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x64, 0x69, 0x73, 0x74,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x72, 0x72, 0x75, 0x70, 0x74, 0x65, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoc_admin_proto_rawDescOnce sync.Once
	file_protoc_admin_proto_rawDescData = file_protoc_admin_proto_rawDesc
)

func file_protoc_admin_proto_rawDescGZIP() []byte {
	file_protoc_admin_proto_rawDescOnce.Do(func() {
		file_protoc_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoc_admin_proto_rawDescData)
	})
	return file_protoc_admin_proto_rawDescData
}

var file_protoc_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protoc_admin_proto_goTypes = []interface{}{
	(*DeleteCorruptedFilesV1Request)(nil),  // 0: dist_file_storage.DeleteCorruptedFilesV1Request
	(*DeleteCorruptedFilesV1Response)(nil), // 1: dist_file_storage.DeleteCorruptedFilesV1Response
}
var file_protoc_admin_proto_depIdxs = []int32{
	0, // 0: dist_file_storage.FileAdminService.DeleteCorruptedFilesV1:input_type -> dist_file_storage.DeleteCorruptedFilesV1Request
	1, // 1: dist_file_storage.FileAdminService.DeleteCorruptedFilesV1:output_type -> dist_file_storage.DeleteCorruptedFilesV1Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protoc_admin_proto_init() }
func file_protoc_admin_proto_init() {
	if File_protoc_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoc_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCorruptedFilesV1Request); i {
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
		file_protoc_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCorruptedFilesV1Response); i {
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
			RawDescriptor: file_protoc_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protoc_admin_proto_goTypes,
		DependencyIndexes: file_protoc_admin_proto_depIdxs,
		MessageInfos:      file_protoc_admin_proto_msgTypes,
	}.Build()
	File_protoc_admin_proto = out.File
	file_protoc_admin_proto_rawDesc = nil
	file_protoc_admin_proto_goTypes = nil
	file_protoc_admin_proto_depIdxs = nil
}
