// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/v1/canary/canarysvc.proto

package canary

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CanaryService_Health_FullMethodName = "/canary.CanaryService/Health"
	CanaryService_Ping_FullMethodName   = "/canary.CanaryService/Ping"
)

// CanaryServiceClient is the client API for CanaryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CanaryServiceClient interface {
	// Sends a health response
	Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthReply, error)
	// Replies Pong
	Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingReply, error)
}

type canaryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCanaryServiceClient(cc grpc.ClientConnInterface) CanaryServiceClient {
	return &canaryServiceClient{cc}
}

func (c *canaryServiceClient) Health(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*HealthReply, error) {
	out := new(HealthReply)
	err := c.cc.Invoke(ctx, CanaryService_Health_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *canaryServiceClient) Ping(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := c.cc.Invoke(ctx, CanaryService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CanaryServiceServer is the server API for CanaryService service.
// All implementations must embed UnimplementedCanaryServiceServer
// for forward compatibility
type CanaryServiceServer interface {
	// Sends a health response
	Health(context.Context, *emptypb.Empty) (*HealthReply, error)
	// Replies Pong
	Ping(context.Context, *emptypb.Empty) (*PingReply, error)
	mustEmbedUnimplementedCanaryServiceServer()
}

// UnimplementedCanaryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCanaryServiceServer struct {
}

func (UnimplementedCanaryServiceServer) Health(context.Context, *emptypb.Empty) (*HealthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedCanaryServiceServer) Ping(context.Context, *emptypb.Empty) (*PingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCanaryServiceServer) mustEmbedUnimplementedCanaryServiceServer() {}

// UnsafeCanaryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CanaryServiceServer will
// result in compilation errors.
type UnsafeCanaryServiceServer interface {
	mustEmbedUnimplementedCanaryServiceServer()
}

func RegisterCanaryServiceServer(s grpc.ServiceRegistrar, srv CanaryServiceServer) {
	s.RegisterService(&CanaryService_ServiceDesc, srv)
}

func _CanaryService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CanaryServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CanaryService_Health_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CanaryServiceServer).Health(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CanaryService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CanaryServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CanaryService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CanaryServiceServer).Ping(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CanaryService_ServiceDesc is the grpc.ServiceDesc for CanaryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CanaryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "canary.CanaryService",
	HandlerType: (*CanaryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _CanaryService_Health_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _CanaryService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/canary/canarysvc.proto",
}