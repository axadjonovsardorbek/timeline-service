// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: shared_memories.proto

package genproto

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
	SharedMemoriesService_Create_FullMethodName  = "/time_capsule.SharedMemoriesService/Create"
	SharedMemoriesService_GetById_FullMethodName = "/time_capsule.SharedMemoriesService/GetById"
	SharedMemoriesService_GetAll_FullMethodName  = "/time_capsule.SharedMemoriesService/GetAll"
	SharedMemoriesService_Update_FullMethodName  = "/time_capsule.SharedMemoriesService/Update"
	SharedMemoriesService_Delete_FullMethodName  = "/time_capsule.SharedMemoriesService/Delete"
)

// SharedMemoriesServiceClient is the client API for SharedMemoriesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SharedMemoriesServiceClient interface {
	Create(ctx context.Context, in *SharedMemoriesCreateReq, opts ...grpc.CallOption) (*Void, error)
	GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*SharedMemoriesGetByIdRes, error)
	GetAll(ctx context.Context, in *SharedMemoriesGetAllReq, opts ...grpc.CallOption) (*SharedMemoriesGetAllRes, error)
	Update(ctx context.Context, in *SharedMemoriesUpdateReq, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
}

type sharedMemoriesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSharedMemoriesServiceClient(cc grpc.ClientConnInterface) SharedMemoriesServiceClient {
	return &sharedMemoriesServiceClient{cc}
}

func (c *sharedMemoriesServiceClient) Create(ctx context.Context, in *SharedMemoriesCreateReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, SharedMemoriesService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedMemoriesServiceClient) GetById(ctx context.Context, in *ById, opts ...grpc.CallOption) (*SharedMemoriesGetByIdRes, error) {
	out := new(SharedMemoriesGetByIdRes)
	err := c.cc.Invoke(ctx, SharedMemoriesService_GetById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedMemoriesServiceClient) GetAll(ctx context.Context, in *SharedMemoriesGetAllReq, opts ...grpc.CallOption) (*SharedMemoriesGetAllRes, error) {
	out := new(SharedMemoriesGetAllRes)
	err := c.cc.Invoke(ctx, SharedMemoriesService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedMemoriesServiceClient) Update(ctx context.Context, in *SharedMemoriesUpdateReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, SharedMemoriesService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharedMemoriesServiceClient) Delete(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, SharedMemoriesService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharedMemoriesServiceServer is the server API for SharedMemoriesService service.
// All implementations must embed UnimplementedSharedMemoriesServiceServer
// for forward compatibility
type SharedMemoriesServiceServer interface {
	Create(context.Context, *SharedMemoriesCreateReq) (*Void, error)
	GetById(context.Context, *ById) (*SharedMemoriesGetByIdRes, error)
	GetAll(context.Context, *SharedMemoriesGetAllReq) (*SharedMemoriesGetAllRes, error)
	Update(context.Context, *SharedMemoriesUpdateReq) (*Void, error)
	Delete(context.Context, *ById) (*Void, error)
	mustEmbedUnimplementedSharedMemoriesServiceServer()
}

// UnimplementedSharedMemoriesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSharedMemoriesServiceServer struct {
}

func (UnimplementedSharedMemoriesServiceServer) Create(context.Context, *SharedMemoriesCreateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSharedMemoriesServiceServer) GetById(context.Context, *ById) (*SharedMemoriesGetByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedSharedMemoriesServiceServer) GetAll(context.Context, *SharedMemoriesGetAllReq) (*SharedMemoriesGetAllRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedSharedMemoriesServiceServer) Update(context.Context, *SharedMemoriesUpdateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSharedMemoriesServiceServer) Delete(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSharedMemoriesServiceServer) mustEmbedUnimplementedSharedMemoriesServiceServer() {}

// UnsafeSharedMemoriesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SharedMemoriesServiceServer will
// result in compilation errors.
type UnsafeSharedMemoriesServiceServer interface {
	mustEmbedUnimplementedSharedMemoriesServiceServer()
}

func RegisterSharedMemoriesServiceServer(s grpc.ServiceRegistrar, srv SharedMemoriesServiceServer) {
	s.RegisterService(&SharedMemoriesService_ServiceDesc, srv)
}

func _SharedMemoriesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SharedMemoriesCreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedMemoriesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SharedMemoriesService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedMemoriesServiceServer).Create(ctx, req.(*SharedMemoriesCreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedMemoriesService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedMemoriesServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SharedMemoriesService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedMemoriesServiceServer).GetById(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedMemoriesService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SharedMemoriesGetAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedMemoriesServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SharedMemoriesService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedMemoriesServiceServer).GetAll(ctx, req.(*SharedMemoriesGetAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedMemoriesService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SharedMemoriesUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedMemoriesServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SharedMemoriesService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedMemoriesServiceServer).Update(ctx, req.(*SharedMemoriesUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SharedMemoriesService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharedMemoriesServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SharedMemoriesService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharedMemoriesServiceServer).Delete(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

// SharedMemoriesService_ServiceDesc is the grpc.ServiceDesc for SharedMemoriesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SharedMemoriesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "time_capsule.SharedMemoriesService",
	HandlerType: (*SharedMemoriesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SharedMemoriesService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _SharedMemoriesService_GetById_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _SharedMemoriesService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SharedMemoriesService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SharedMemoriesService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shared_memories.proto",
}