// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: protoc/fms.proto

package api

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

type InitFileUploadV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"` // в рамках тестового задания используем имя файла как уникальный идентификатор
	Size     uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *InitFileUploadV1Request) Reset() {
	*x = InitFileUploadV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_fms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitFileUploadV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitFileUploadV1Request) ProtoMessage() {}

func (x *InitFileUploadV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_fms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitFileUploadV1Request.ProtoReflect.Descriptor instead.
func (*InitFileUploadV1Request) Descriptor() ([]byte, []int) {
	return file_protoc_fms_proto_rawDescGZIP(), []int{0}
}

func (x *InitFileUploadV1Request) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *InitFileUploadV1Request) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type FilePart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// используем id, чтоб не делать неявную завязку на порядок частей в списке fileParts
	// также, по partId нужно сортировать части файла при разбиении и сборке файла из частей
	PartId  int32  `protobuf:"varint,1,opt,name=partId,proto3" json:"partId,omitempty"`
	Storage string `protobuf:"bytes,2,opt,name=storage,proto3" json:"storage,omitempty"` // ссылка на сервис B_n, куда сохранить эту часть
	Size    uint64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Path    string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"` // по какому пути на сервере-хранилище сохранить
}

func (x *FilePart) Reset() {
	*x = FilePart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_fms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilePart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilePart) ProtoMessage() {}

func (x *FilePart) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_fms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilePart.ProtoReflect.Descriptor instead.
func (*FilePart) Descriptor() ([]byte, []int) {
	return file_protoc_fms_proto_rawDescGZIP(), []int{1}
}

func (x *FilePart) GetPartId() int32 {
	if x != nil {
		return x.PartId
	}
	return 0
}

func (x *FilePart) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

func (x *FilePart) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FilePart) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type InitFileUploadV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileParts []*FilePart `protobuf:"bytes,1,rep,name=fileParts,proto3" json:"fileParts,omitempty"`
}

func (x *InitFileUploadV1Response) Reset() {
	*x = InitFileUploadV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_fms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitFileUploadV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitFileUploadV1Response) ProtoMessage() {}

func (x *InitFileUploadV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_fms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitFileUploadV1Response.ProtoReflect.Descriptor instead.
func (*InitFileUploadV1Response) Descriptor() ([]byte, []int) {
	return file_protoc_fms_proto_rawDescGZIP(), []int{2}
}

func (x *InitFileUploadV1Response) GetFileParts() []*FilePart {
	if x != nil {
		return x.FileParts
	}
	return nil
}

type ReportUploadProgressV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	PartId   int32  `protobuf:"varint,2,opt,name=partId,proto3" json:"partId,omitempty"`
}

func (x *ReportUploadProgressV1Request) Reset() {
	*x = ReportUploadProgressV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_fms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportUploadProgressV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportUploadProgressV1Request) ProtoMessage() {}

func (x *ReportUploadProgressV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_fms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportUploadProgressV1Request.ProtoReflect.Descriptor instead.
func (*ReportUploadProgressV1Request) Descriptor() ([]byte, []int) {
	return file_protoc_fms_proto_rawDescGZIP(), []int{3}
}

func (x *ReportUploadProgressV1Request) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ReportUploadProgressV1Request) GetPartId() int32 {
	if x != nil {
		return x.PartId
	}
	return 0
}

type GetFileDownloadInfoV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *GetFileDownloadInfoV1Request) Reset() {
	*x = GetFileDownloadInfoV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_fms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileDownloadInfoV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileDownloadInfoV1Request) ProtoMessage() {}

func (x *GetFileDownloadInfoV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_fms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileDownloadInfoV1Request.ProtoReflect.Descriptor instead.
func (*GetFileDownloadInfoV1Request) Descriptor() ([]byte, []int) {
	return file_protoc_fms_proto_rawDescGZIP(), []int{4}
}

func (x *GetFileDownloadInfoV1Request) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type GetFileDownloadInfoV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileParts []*FilePart `protobuf:"bytes,2,rep,name=fileParts,proto3" json:"fileParts,omitempty"`
}

func (x *GetFileDownloadInfoV1Response) Reset() {
	*x = GetFileDownloadInfoV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_fms_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileDownloadInfoV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileDownloadInfoV1Response) ProtoMessage() {}

func (x *GetFileDownloadInfoV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_fms_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileDownloadInfoV1Response.ProtoReflect.Descriptor instead.
func (*GetFileDownloadInfoV1Response) Descriptor() ([]byte, []int) {
	return file_protoc_fms_proto_rawDescGZIP(), []int{5}
}

func (x *GetFileDownloadInfoV1Response) GetFileParts() []*FilePart {
	if x != nil {
		return x.FileParts
	}
	return nil
}

type CancelUploadV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *CancelUploadV1Request) Reset() {
	*x = CancelUploadV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_fms_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelUploadV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelUploadV1Request) ProtoMessage() {}

func (x *CancelUploadV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_fms_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelUploadV1Request.ProtoReflect.Descriptor instead.
func (*CancelUploadV1Request) Descriptor() ([]byte, []int) {
	return file_protoc_fms_proto_rawDescGZIP(), []int{6}
}

func (x *CancelUploadV1Request) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type AddStorageV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr                string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	SpaceAvailableBytes uint64 `protobuf:"varint,2,opt,name=space_available_bytes,json=spaceAvailableBytes,proto3" json:"space_available_bytes,omitempty"`
}

func (x *AddStorageV1Request) Reset() {
	*x = AddStorageV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_fms_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStorageV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStorageV1Request) ProtoMessage() {}

func (x *AddStorageV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_fms_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStorageV1Request.ProtoReflect.Descriptor instead.
func (*AddStorageV1Request) Descriptor() ([]byte, []int) {
	return file_protoc_fms_proto_rawDescGZIP(), []int{7}
}

func (x *AddStorageV1Request) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *AddStorageV1Request) GetSpaceAvailableBytes() uint64 {
	if x != nil {
		return x.SpaceAvailableBytes
	}
	return 0
}

type AddStorageV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddStorageV1Response) Reset() {
	*x = AddStorageV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protoc_fms_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddStorageV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddStorageV1Response) ProtoMessage() {}

func (x *AddStorageV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_protoc_fms_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddStorageV1Response.ProtoReflect.Descriptor instead.
func (*AddStorageV1Response) Descriptor() ([]byte, []int) {
	return file_protoc_fms_proto_rawDescGZIP(), []int{8}
}

func (x *AddStorageV1Response) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_protoc_fms_proto protoreflect.FileDescriptor

var file_protoc_fms_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x66, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x64, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x49, 0x0a, 0x17, 0x49, 0x6e, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x64, 0x0a,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x72, 0x74, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x22, 0x55, 0x0a, 0x18, 0x49, 0x6e, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x39, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x52,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x73, 0x22, 0x53, 0x0a, 0x1d, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x61, 0x72, 0x74, 0x49, 0x64, 0x22,
	0x3a, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5a, 0x0a, 0x1d, 0x47,
	0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x52, 0x09, 0x66, 0x69,
	0x6c, 0x65, 0x50, 0x61, 0x72, 0x74, 0x73, 0x22, 0x33, 0x0a, 0x15, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x13,
	0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x73, 0x70, 0x61, 0x63, 0x65, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x14, 0x41,
	0x64, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x32, 0xb8, 0x03, 0x0a, 0x15, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a,
	0x10, 0x49, 0x6e, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56,
	0x31, 0x12, 0x2a, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x64, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x16, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x56, 0x31, 0x12, 0x30, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x52,
	0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x56, 0x31,
	0x12, 0x28, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x7a, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x56, 0x31, 0x12, 0x2f, 0x2e, 0x64, 0x69,
	0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x64,
	0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x7d,
	0x0a, 0x1a, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0c,
	0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x31, 0x12, 0x26, 0x2e, 0x64,
	0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protoc_fms_proto_rawDescOnce sync.Once
	file_protoc_fms_proto_rawDescData = file_protoc_fms_proto_rawDesc
)

func file_protoc_fms_proto_rawDescGZIP() []byte {
	file_protoc_fms_proto_rawDescOnce.Do(func() {
		file_protoc_fms_proto_rawDescData = protoimpl.X.CompressGZIP(file_protoc_fms_proto_rawDescData)
	})
	return file_protoc_fms_proto_rawDescData
}

var file_protoc_fms_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protoc_fms_proto_goTypes = []interface{}{
	(*InitFileUploadV1Request)(nil),       // 0: dist_file_storage.InitFileUploadV1Request
	(*FilePart)(nil),                      // 1: dist_file_storage.FilePart
	(*InitFileUploadV1Response)(nil),      // 2: dist_file_storage.InitFileUploadV1Response
	(*ReportUploadProgressV1Request)(nil), // 3: dist_file_storage.ReportUploadProgressV1Request
	(*GetFileDownloadInfoV1Request)(nil),  // 4: dist_file_storage.GetFileDownloadInfoV1Request
	(*GetFileDownloadInfoV1Response)(nil), // 5: dist_file_storage.GetFileDownloadInfoV1Response
	(*CancelUploadV1Request)(nil),         // 6: dist_file_storage.CancelUploadV1Request
	(*AddStorageV1Request)(nil),           // 7: dist_file_storage.AddStorageV1Request
	(*AddStorageV1Response)(nil),          // 8: dist_file_storage.AddStorageV1Response
	(*emptypb.Empty)(nil),                 // 9: google.protobuf.Empty
}
var file_protoc_fms_proto_depIdxs = []int32{
	1, // 0: dist_file_storage.InitFileUploadV1Response.fileParts:type_name -> dist_file_storage.FilePart
	1, // 1: dist_file_storage.GetFileDownloadInfoV1Response.fileParts:type_name -> dist_file_storage.FilePart
	0, // 2: dist_file_storage.FileManagementService.InitFileUploadV1:input_type -> dist_file_storage.InitFileUploadV1Request
	3, // 3: dist_file_storage.FileManagementService.ReportUploadProgressV1:input_type -> dist_file_storage.ReportUploadProgressV1Request
	6, // 4: dist_file_storage.FileManagementService.CancelUploadV1:input_type -> dist_file_storage.CancelUploadV1Request
	4, // 5: dist_file_storage.FileManagementService.GetFileDownloadInfoV1:input_type -> dist_file_storage.GetFileDownloadInfoV1Request
	7, // 6: dist_file_storage.FileManagementAdminService.AddStorageV1:input_type -> dist_file_storage.AddStorageV1Request
	2, // 7: dist_file_storage.FileManagementService.InitFileUploadV1:output_type -> dist_file_storage.InitFileUploadV1Response
	9, // 8: dist_file_storage.FileManagementService.ReportUploadProgressV1:output_type -> google.protobuf.Empty
	9, // 9: dist_file_storage.FileManagementService.CancelUploadV1:output_type -> google.protobuf.Empty
	5, // 10: dist_file_storage.FileManagementService.GetFileDownloadInfoV1:output_type -> dist_file_storage.GetFileDownloadInfoV1Response
	8, // 11: dist_file_storage.FileManagementAdminService.AddStorageV1:output_type -> dist_file_storage.AddStorageV1Response
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protoc_fms_proto_init() }
func file_protoc_fms_proto_init() {
	if File_protoc_fms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protoc_fms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitFileUploadV1Request); i {
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
		file_protoc_fms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilePart); i {
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
		file_protoc_fms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitFileUploadV1Response); i {
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
		file_protoc_fms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportUploadProgressV1Request); i {
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
		file_protoc_fms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileDownloadInfoV1Request); i {
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
		file_protoc_fms_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFileDownloadInfoV1Response); i {
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
		file_protoc_fms_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelUploadV1Request); i {
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
		file_protoc_fms_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStorageV1Request); i {
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
		file_protoc_fms_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddStorageV1Response); i {
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
			RawDescriptor: file_protoc_fms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_protoc_fms_proto_goTypes,
		DependencyIndexes: file_protoc_fms_proto_depIdxs,
		MessageInfos:      file_protoc_fms_proto_msgTypes,
	}.Build()
	File_protoc_fms_proto = out.File
	file_protoc_fms_proto_rawDesc = nil
	file_protoc_fms_proto_goTypes = nil
	file_protoc_fms_proto_depIdxs = nil
}
