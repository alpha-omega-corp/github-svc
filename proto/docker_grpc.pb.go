// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: proto/docker.proto

package proto

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

// DockerServiceClient is the client API for DockerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DockerServiceClient interface {
	CreatePackage(ctx context.Context, in *CreatePackageRequest, opts ...grpc.CallOption) (*CreatePackageResponse, error)
	GetPackages(ctx context.Context, in *GetPackagesRequest, opts ...grpc.CallOption) (*GetPackagesResponse, error)
	DeletePackage(ctx context.Context, in *DeletePackageRequest, opts ...grpc.CallOption) (*DeletePackageResponse, error)
	GetContainers(ctx context.Context, in *GetContainersRequest, opts ...grpc.CallOption) (*GetContainersResponse, error)
	GetContainerLogs(ctx context.Context, in *GetContainerLogsRequest, opts ...grpc.CallOption) (*GetContainerLogsResponse, error)
}

type dockerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDockerServiceClient(cc grpc.ClientConnInterface) DockerServiceClient {
	return &dockerServiceClient{cc}
}

func (c *dockerServiceClient) CreatePackage(ctx context.Context, in *CreatePackageRequest, opts ...grpc.CallOption) (*CreatePackageResponse, error) {
	out := new(CreatePackageResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/CreatePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) GetPackages(ctx context.Context, in *GetPackagesRequest, opts ...grpc.CallOption) (*GetPackagesResponse, error) {
	out := new(GetPackagesResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/GetPackages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) DeletePackage(ctx context.Context, in *DeletePackageRequest, opts ...grpc.CallOption) (*DeletePackageResponse, error) {
	out := new(DeletePackageResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/DeletePackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) GetContainers(ctx context.Context, in *GetContainersRequest, opts ...grpc.CallOption) (*GetContainersResponse, error) {
	out := new(GetContainersResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/GetContainers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dockerServiceClient) GetContainerLogs(ctx context.Context, in *GetContainerLogsRequest, opts ...grpc.CallOption) (*GetContainerLogsResponse, error) {
	out := new(GetContainerLogsResponse)
	err := c.cc.Invoke(ctx, "/docker.DockerService/GetContainerLogs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DockerServiceServer is the server API for DockerService service.
// All implementations must embed UnimplementedDockerServiceServer
// for forward compatibility
type DockerServiceServer interface {
	CreatePackage(context.Context, *CreatePackageRequest) (*CreatePackageResponse, error)
	GetPackages(context.Context, *GetPackagesRequest) (*GetPackagesResponse, error)
	DeletePackage(context.Context, *DeletePackageRequest) (*DeletePackageResponse, error)
	GetContainers(context.Context, *GetContainersRequest) (*GetContainersResponse, error)
	GetContainerLogs(context.Context, *GetContainerLogsRequest) (*GetContainerLogsResponse, error)
	mustEmbedUnimplementedDockerServiceServer()
}

// UnimplementedDockerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDockerServiceServer struct {
}

func (UnimplementedDockerServiceServer) CreatePackage(context.Context, *CreatePackageRequest) (*CreatePackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePackage not implemented")
}
func (UnimplementedDockerServiceServer) GetPackages(context.Context, *GetPackagesRequest) (*GetPackagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPackages not implemented")
}
func (UnimplementedDockerServiceServer) DeletePackage(context.Context, *DeletePackageRequest) (*DeletePackageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePackage not implemented")
}
func (UnimplementedDockerServiceServer) GetContainers(context.Context, *GetContainersRequest) (*GetContainersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainers not implemented")
}
func (UnimplementedDockerServiceServer) GetContainerLogs(context.Context, *GetContainerLogsRequest) (*GetContainerLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContainerLogs not implemented")
}
func (UnimplementedDockerServiceServer) mustEmbedUnimplementedDockerServiceServer() {}

// UnsafeDockerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DockerServiceServer will
// result in compilation errors.
type UnsafeDockerServiceServer interface {
	mustEmbedUnimplementedDockerServiceServer()
}

func RegisterDockerServiceServer(s grpc.ServiceRegistrar, srv DockerServiceServer) {
	s.RegisterService(&DockerService_ServiceDesc, srv)
}

func _DockerService_CreatePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).CreatePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/CreatePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).CreatePackage(ctx, req.(*CreatePackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_GetPackages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPackagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).GetPackages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/GetPackages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).GetPackages(ctx, req.(*GetPackagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_DeletePackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).DeletePackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/DeletePackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).DeletePackage(ctx, req.(*DeletePackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_GetContainers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).GetContainers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/GetContainers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).GetContainers(ctx, req.(*GetContainersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DockerService_GetContainerLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContainerLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DockerServiceServer).GetContainerLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.DockerService/GetContainerLogs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DockerServiceServer).GetContainerLogs(ctx, req.(*GetContainerLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DockerService_ServiceDesc is the grpc.ServiceDesc for DockerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DockerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "docker.DockerService",
	HandlerType: (*DockerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePackage",
			Handler:    _DockerService_CreatePackage_Handler,
		},
		{
			MethodName: "GetPackages",
			Handler:    _DockerService_GetPackages_Handler,
		},
		{
			MethodName: "DeletePackage",
			Handler:    _DockerService_DeletePackage_Handler,
		},
		{
			MethodName: "GetContainers",
			Handler:    _DockerService_GetContainers_Handler,
		},
		{
			MethodName: "GetContainerLogs",
			Handler:    _DockerService_GetContainerLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/docker.proto",
}
