// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: minivirt/port_forwarding.proto

package pb

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
	PortForwardingService_GetPortForwarding_FullMethodName    = "/PortForwardingService/GetPortForwarding"
	PortForwardingService_ListPortForwardings_FullMethodName  = "/PortForwardingService/ListPortForwardings"
	PortForwardingService_PutPortForwarding_FullMethodName    = "/PortForwardingService/PutPortForwarding"
	PortForwardingService_DeletePortForwarding_FullMethodName = "/PortForwardingService/DeletePortForwarding"
)

// PortForwardingServiceClient is the client API for PortForwardingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortForwardingServiceClient interface {
	GetPortForwarding(ctx context.Context, in *PortForwardingIdentifier, opts ...grpc.CallOption) (*PortForwarding, error)
	ListPortForwardings(ctx context.Context, in *ListPortForwardingsRequest, opts ...grpc.CallOption) (*ListPortForwardingsResponse, error)
	PutPortForwarding(ctx context.Context, in *PutPortForwardingRequest, opts ...grpc.CallOption) (*PortForwarding, error)
	DeletePortForwarding(ctx context.Context, in *PortForwardingIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type portForwardingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPortForwardingServiceClient(cc grpc.ClientConnInterface) PortForwardingServiceClient {
	return &portForwardingServiceClient{cc}
}

func (c *portForwardingServiceClient) GetPortForwarding(ctx context.Context, in *PortForwardingIdentifier, opts ...grpc.CallOption) (*PortForwarding, error) {
	out := new(PortForwarding)
	err := c.cc.Invoke(ctx, PortForwardingService_GetPortForwarding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portForwardingServiceClient) ListPortForwardings(ctx context.Context, in *ListPortForwardingsRequest, opts ...grpc.CallOption) (*ListPortForwardingsResponse, error) {
	out := new(ListPortForwardingsResponse)
	err := c.cc.Invoke(ctx, PortForwardingService_ListPortForwardings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portForwardingServiceClient) PutPortForwarding(ctx context.Context, in *PutPortForwardingRequest, opts ...grpc.CallOption) (*PortForwarding, error) {
	out := new(PortForwarding)
	err := c.cc.Invoke(ctx, PortForwardingService_PutPortForwarding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portForwardingServiceClient) DeletePortForwarding(ctx context.Context, in *PortForwardingIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PortForwardingService_DeletePortForwarding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortForwardingServiceServer is the server API for PortForwardingService service.
// All implementations must embed UnimplementedPortForwardingServiceServer
// for forward compatibility
type PortForwardingServiceServer interface {
	GetPortForwarding(context.Context, *PortForwardingIdentifier) (*PortForwarding, error)
	ListPortForwardings(context.Context, *ListPortForwardingsRequest) (*ListPortForwardingsResponse, error)
	PutPortForwarding(context.Context, *PutPortForwardingRequest) (*PortForwarding, error)
	DeletePortForwarding(context.Context, *PortForwardingIdentifier) (*emptypb.Empty, error)
	mustEmbedUnimplementedPortForwardingServiceServer()
}

// UnimplementedPortForwardingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPortForwardingServiceServer struct {
}

func (UnimplementedPortForwardingServiceServer) GetPortForwarding(context.Context, *PortForwardingIdentifier) (*PortForwarding, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPortForwarding not implemented")
}
func (UnimplementedPortForwardingServiceServer) ListPortForwardings(context.Context, *ListPortForwardingsRequest) (*ListPortForwardingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPortForwardings not implemented")
}
func (UnimplementedPortForwardingServiceServer) PutPortForwarding(context.Context, *PutPortForwardingRequest) (*PortForwarding, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPortForwarding not implemented")
}
func (UnimplementedPortForwardingServiceServer) DeletePortForwarding(context.Context, *PortForwardingIdentifier) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePortForwarding not implemented")
}
func (UnimplementedPortForwardingServiceServer) mustEmbedUnimplementedPortForwardingServiceServer() {}

// UnsafePortForwardingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortForwardingServiceServer will
// result in compilation errors.
type UnsafePortForwardingServiceServer interface {
	mustEmbedUnimplementedPortForwardingServiceServer()
}

func RegisterPortForwardingServiceServer(s grpc.ServiceRegistrar, srv PortForwardingServiceServer) {
	s.RegisterService(&PortForwardingService_ServiceDesc, srv)
}

func _PortForwardingService_GetPortForwarding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortForwardingIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortForwardingServiceServer).GetPortForwarding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortForwardingService_GetPortForwarding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortForwardingServiceServer).GetPortForwarding(ctx, req.(*PortForwardingIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortForwardingService_ListPortForwardings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPortForwardingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortForwardingServiceServer).ListPortForwardings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortForwardingService_ListPortForwardings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortForwardingServiceServer).ListPortForwardings(ctx, req.(*ListPortForwardingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortForwardingService_PutPortForwarding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutPortForwardingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortForwardingServiceServer).PutPortForwarding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortForwardingService_PutPortForwarding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortForwardingServiceServer).PutPortForwarding(ctx, req.(*PutPortForwardingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortForwardingService_DeletePortForwarding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortForwardingIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortForwardingServiceServer).DeletePortForwarding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortForwardingService_DeletePortForwarding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortForwardingServiceServer).DeletePortForwarding(ctx, req.(*PortForwardingIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

// PortForwardingService_ServiceDesc is the grpc.ServiceDesc for PortForwardingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortForwardingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PortForwardingService",
	HandlerType: (*PortForwardingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPortForwarding",
			Handler:    _PortForwardingService_GetPortForwarding_Handler,
		},
		{
			MethodName: "ListPortForwardings",
			Handler:    _PortForwardingService_ListPortForwardings_Handler,
		},
		{
			MethodName: "PutPortForwarding",
			Handler:    _PortForwardingService_PutPortForwarding_Handler,
		},
		{
			MethodName: "DeletePortForwarding",
			Handler:    _PortForwardingService_DeletePortForwarding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "minivirt/port_forwarding.proto",
}
