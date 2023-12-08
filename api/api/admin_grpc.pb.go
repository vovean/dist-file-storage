// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: protoc/admin.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FileAdminService_DeleteCorruptedFilesV1_FullMethodName = "/dist_file_storage.FileAdminService/DeleteCorruptedFilesV1"
)

// FileAdminServiceClient is the client API for FileAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileAdminServiceClient interface {
	DeleteCorruptedFilesV1(ctx context.Context, in *DeleteCorruptedFilesV1Request, opts ...grpc.CallOption) (*DeleteCorruptedFilesV1Response, error)
}

type fileAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileAdminServiceClient(cc grpc.ClientConnInterface) FileAdminServiceClient {
	return &fileAdminServiceClient{cc}
}

func (c *fileAdminServiceClient) DeleteCorruptedFilesV1(ctx context.Context, in *DeleteCorruptedFilesV1Request, opts ...grpc.CallOption) (*DeleteCorruptedFilesV1Response, error) {
	out := new(DeleteCorruptedFilesV1Response)
	err := c.cc.Invoke(ctx, FileAdminService_DeleteCorruptedFilesV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileAdminServiceServer is the server API for FileAdminService service.
// All implementations should embed UnimplementedFileAdminServiceServer
// for forward compatibility
type FileAdminServiceServer interface {
	DeleteCorruptedFilesV1(context.Context, *DeleteCorruptedFilesV1Request) (*DeleteCorruptedFilesV1Response, error)
}

// UnimplementedFileAdminServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFileAdminServiceServer struct {
}

func (UnimplementedFileAdminServiceServer) DeleteCorruptedFilesV1(context.Context, *DeleteCorruptedFilesV1Request) (*DeleteCorruptedFilesV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCorruptedFilesV1 not implemented")
}

// UnsafeFileAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileAdminServiceServer will
// result in compilation errors.
type UnsafeFileAdminServiceServer interface {
	mustEmbedUnimplementedFileAdminServiceServer()
}

func RegisterFileAdminServiceServer(s grpc.ServiceRegistrar, srv FileAdminServiceServer) {
	s.RegisterService(&FileAdminService_ServiceDesc, srv)
}

func _FileAdminService_DeleteCorruptedFilesV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCorruptedFilesV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileAdminServiceServer).DeleteCorruptedFilesV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FileAdminService_DeleteCorruptedFilesV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileAdminServiceServer).DeleteCorruptedFilesV1(ctx, req.(*DeleteCorruptedFilesV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// FileAdminService_ServiceDesc is the grpc.ServiceDesc for FileAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dist_file_storage.FileAdminService",
	HandlerType: (*FileAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteCorruptedFilesV1",
			Handler:    _FileAdminService_DeleteCorruptedFilesV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protoc/admin.proto",
}