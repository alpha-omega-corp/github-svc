// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: proto/github/github.proto

package github

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

// GithubServiceClient is the client API for GithubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GithubServiceClient interface {
	GetSecrets(ctx context.Context, in *GetSecretsRequest, opts ...grpc.CallOption) (*GetSecretsResponse, error)
	CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*CreateSecretResponse, error)
	DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*DeleteSecretResponse, error)
}

type githubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGithubServiceClient(cc grpc.ClientConnInterface) GithubServiceClient {
	return &githubServiceClient{cc}
}

func (c *githubServiceClient) GetSecrets(ctx context.Context, in *GetSecretsRequest, opts ...grpc.CallOption) (*GetSecretsResponse, error) {
	out := new(GetSecretsResponse)
	err := c.cc.Invoke(ctx, "/docker.GithubService/GetSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubServiceClient) CreateSecret(ctx context.Context, in *CreateSecretRequest, opts ...grpc.CallOption) (*CreateSecretResponse, error) {
	out := new(CreateSecretResponse)
	err := c.cc.Invoke(ctx, "/docker.GithubService/CreateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *githubServiceClient) DeleteSecret(ctx context.Context, in *DeleteSecretRequest, opts ...grpc.CallOption) (*DeleteSecretResponse, error) {
	out := new(DeleteSecretResponse)
	err := c.cc.Invoke(ctx, "/docker.GithubService/DeleteSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GithubServiceServer is the server API for GithubService service.
// All implementations must embed UnimplementedGithubServiceServer
// for forward compatibility
type GithubServiceServer interface {
	GetSecrets(context.Context, *GetSecretsRequest) (*GetSecretsResponse, error)
	CreateSecret(context.Context, *CreateSecretRequest) (*CreateSecretResponse, error)
	DeleteSecret(context.Context, *DeleteSecretRequest) (*DeleteSecretResponse, error)
	mustEmbedUnimplementedGithubServiceServer()
}

// UnimplementedGithubServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGithubServiceServer struct {
}

func (UnimplementedGithubServiceServer) GetSecrets(context.Context, *GetSecretsRequest) (*GetSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecrets not implemented")
}
func (UnimplementedGithubServiceServer) CreateSecret(context.Context, *CreateSecretRequest) (*CreateSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}
func (UnimplementedGithubServiceServer) DeleteSecret(context.Context, *DeleteSecretRequest) (*DeleteSecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}
func (UnimplementedGithubServiceServer) mustEmbedUnimplementedGithubServiceServer() {}

// UnsafeGithubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GithubServiceServer will
// result in compilation errors.
type UnsafeGithubServiceServer interface {
	mustEmbedUnimplementedGithubServiceServer()
}

func RegisterGithubServiceServer(s grpc.ServiceRegistrar, srv GithubServiceServer) {
	s.RegisterService(&GithubService_ServiceDesc, srv)
}

func _GithubService_GetSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubServiceServer).GetSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.GithubService/GetSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubServiceServer).GetSecrets(ctx, req.(*GetSecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubService_CreateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubServiceServer).CreateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.GithubService/CreateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubServiceServer).CreateSecret(ctx, req.(*CreateSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GithubService_DeleteSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubServiceServer).DeleteSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/docker.GithubService/DeleteSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubServiceServer).DeleteSecret(ctx, req.(*DeleteSecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GithubService_ServiceDesc is the grpc.ServiceDesc for GithubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GithubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "docker.GithubService",
	HandlerType: (*GithubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSecrets",
			Handler:    _GithubService_GetSecrets_Handler,
		},
		{
			MethodName: "CreateSecret",
			Handler:    _GithubService_CreateSecret_Handler,
		},
		{
			MethodName: "DeleteSecret",
			Handler:    _GithubService_DeleteSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/github/github.proto",
}