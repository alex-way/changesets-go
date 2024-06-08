// https://github.com/alex-way/changesets/blob/main/pkg/plugin/plugin.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.26.1
// source: plugin.proto

package plugin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	VersionGetterSetterService_Request_FullMethodName = "/plugin.VersionGetterSetterService/Request"
)

// VersionGetterSetterServiceClient is the client API for VersionGetterSetterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VersionGetterSetterServiceClient interface {
	Request(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*Response, error)
}

type versionGetterSetterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVersionGetterSetterServiceClient(cc grpc.ClientConnInterface) VersionGetterSetterServiceClient {
	return &versionGetterSetterServiceClient{cc}
}

func (c *versionGetterSetterServiceClient) Request(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, VersionGetterSetterService_Request_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VersionGetterSetterServiceServer is the server API for VersionGetterSetterService service.
// All implementations must embed UnimplementedVersionGetterSetterServiceServer
// for forward compatibility
type VersionGetterSetterServiceServer interface {
	Request(context.Context, *RequestMessage) (*Response, error)
	mustEmbedUnimplementedVersionGetterSetterServiceServer()
}

// UnimplementedVersionGetterSetterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVersionGetterSetterServiceServer struct {
}

func (UnimplementedVersionGetterSetterServiceServer) Request(context.Context, *RequestMessage) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (UnimplementedVersionGetterSetterServiceServer) mustEmbedUnimplementedVersionGetterSetterServiceServer() {
}

// UnsafeVersionGetterSetterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VersionGetterSetterServiceServer will
// result in compilation errors.
type UnsafeVersionGetterSetterServiceServer interface {
	mustEmbedUnimplementedVersionGetterSetterServiceServer()
}

func RegisterVersionGetterSetterServiceServer(s grpc.ServiceRegistrar, srv VersionGetterSetterServiceServer) {
	s.RegisterService(&VersionGetterSetterService_ServiceDesc, srv)
}

func _VersionGetterSetterService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionGetterSetterServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VersionGetterSetterService_Request_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionGetterSetterServiceServer).Request(ctx, req.(*RequestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// VersionGetterSetterService_ServiceDesc is the grpc.ServiceDesc for VersionGetterSetterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VersionGetterSetterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plugin.VersionGetterSetterService",
	HandlerType: (*VersionGetterSetterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _VersionGetterSetterService_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin.proto",
}