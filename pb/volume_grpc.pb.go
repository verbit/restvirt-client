// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: minivirt/volume.proto

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
	VolumeService_GetVolume_FullMethodName             = "/VolumeService/GetVolume"
	VolumeService_ListVolumes_FullMethodName           = "/VolumeService/ListVolumes"
	VolumeService_CreateVolume_FullMethodName          = "/VolumeService/CreateVolume"
	VolumeService_UpdateVolume_FullMethodName          = "/VolumeService/UpdateVolume"
	VolumeService_DeleteVolume_FullMethodName          = "/VolumeService/DeleteVolume"
	VolumeService_ListVolumeAttachments_FullMethodName = "/VolumeService/ListVolumeAttachments"
	VolumeService_GetVolumeAttachment_FullMethodName   = "/VolumeService/GetVolumeAttachment"
	VolumeService_AttachVolume_FullMethodName          = "/VolumeService/AttachVolume"
	VolumeService_DetachVolume_FullMethodName          = "/VolumeService/DetachVolume"
)

// VolumeServiceClient is the client API for VolumeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VolumeServiceClient interface {
	GetVolume(ctx context.Context, in *GetVolumeRequest, opts ...grpc.CallOption) (*Volume, error)
	ListVolumes(ctx context.Context, in *ListVolumesRequest, opts ...grpc.CallOption) (*ListVolumesResponse, error)
	CreateVolume(ctx context.Context, in *CreateVolumeRequest, opts ...grpc.CallOption) (*Volume, error)
	UpdateVolume(ctx context.Context, in *UpdateVolumeRequest, opts ...grpc.CallOption) (*Volume, error)
	DeleteVolume(ctx context.Context, in *DeleteVolumeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListVolumeAttachments(ctx context.Context, in *ListVolumeAttachmentsRequest, opts ...grpc.CallOption) (*ListVolumeAttachmentsResponse, error)
	GetVolumeAttachment(ctx context.Context, in *VolumeAttachmentIdentifier, opts ...grpc.CallOption) (*VolumeAttachment, error)
	AttachVolume(ctx context.Context, in *VolumeAttachmentIdentifier, opts ...grpc.CallOption) (*VolumeAttachment, error)
	DetachVolume(ctx context.Context, in *VolumeAttachmentIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type volumeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVolumeServiceClient(cc grpc.ClientConnInterface) VolumeServiceClient {
	return &volumeServiceClient{cc}
}

func (c *volumeServiceClient) GetVolume(ctx context.Context, in *GetVolumeRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := c.cc.Invoke(ctx, VolumeService_GetVolume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) ListVolumes(ctx context.Context, in *ListVolumesRequest, opts ...grpc.CallOption) (*ListVolumesResponse, error) {
	out := new(ListVolumesResponse)
	err := c.cc.Invoke(ctx, VolumeService_ListVolumes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) CreateVolume(ctx context.Context, in *CreateVolumeRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := c.cc.Invoke(ctx, VolumeService_CreateVolume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) UpdateVolume(ctx context.Context, in *UpdateVolumeRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := c.cc.Invoke(ctx, VolumeService_UpdateVolume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) DeleteVolume(ctx context.Context, in *DeleteVolumeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, VolumeService_DeleteVolume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) ListVolumeAttachments(ctx context.Context, in *ListVolumeAttachmentsRequest, opts ...grpc.CallOption) (*ListVolumeAttachmentsResponse, error) {
	out := new(ListVolumeAttachmentsResponse)
	err := c.cc.Invoke(ctx, VolumeService_ListVolumeAttachments_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) GetVolumeAttachment(ctx context.Context, in *VolumeAttachmentIdentifier, opts ...grpc.CallOption) (*VolumeAttachment, error) {
	out := new(VolumeAttachment)
	err := c.cc.Invoke(ctx, VolumeService_GetVolumeAttachment_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) AttachVolume(ctx context.Context, in *VolumeAttachmentIdentifier, opts ...grpc.CallOption) (*VolumeAttachment, error) {
	out := new(VolumeAttachment)
	err := c.cc.Invoke(ctx, VolumeService_AttachVolume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeServiceClient) DetachVolume(ctx context.Context, in *VolumeAttachmentIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, VolumeService_DetachVolume_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VolumeServiceServer is the server API for VolumeService service.
// All implementations must embed UnimplementedVolumeServiceServer
// for forward compatibility
type VolumeServiceServer interface {
	GetVolume(context.Context, *GetVolumeRequest) (*Volume, error)
	ListVolumes(context.Context, *ListVolumesRequest) (*ListVolumesResponse, error)
	CreateVolume(context.Context, *CreateVolumeRequest) (*Volume, error)
	UpdateVolume(context.Context, *UpdateVolumeRequest) (*Volume, error)
	DeleteVolume(context.Context, *DeleteVolumeRequest) (*emptypb.Empty, error)
	ListVolumeAttachments(context.Context, *ListVolumeAttachmentsRequest) (*ListVolumeAttachmentsResponse, error)
	GetVolumeAttachment(context.Context, *VolumeAttachmentIdentifier) (*VolumeAttachment, error)
	AttachVolume(context.Context, *VolumeAttachmentIdentifier) (*VolumeAttachment, error)
	DetachVolume(context.Context, *VolumeAttachmentIdentifier) (*emptypb.Empty, error)
	mustEmbedUnimplementedVolumeServiceServer()
}

// UnimplementedVolumeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVolumeServiceServer struct {
}

func (UnimplementedVolumeServiceServer) GetVolume(context.Context, *GetVolumeRequest) (*Volume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVolume not implemented")
}
func (UnimplementedVolumeServiceServer) ListVolumes(context.Context, *ListVolumesRequest) (*ListVolumesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVolumes not implemented")
}
func (UnimplementedVolumeServiceServer) CreateVolume(context.Context, *CreateVolumeRequest) (*Volume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVolume not implemented")
}
func (UnimplementedVolumeServiceServer) UpdateVolume(context.Context, *UpdateVolumeRequest) (*Volume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVolume not implemented")
}
func (UnimplementedVolumeServiceServer) DeleteVolume(context.Context, *DeleteVolumeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVolume not implemented")
}
func (UnimplementedVolumeServiceServer) ListVolumeAttachments(context.Context, *ListVolumeAttachmentsRequest) (*ListVolumeAttachmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVolumeAttachments not implemented")
}
func (UnimplementedVolumeServiceServer) GetVolumeAttachment(context.Context, *VolumeAttachmentIdentifier) (*VolumeAttachment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVolumeAttachment not implemented")
}
func (UnimplementedVolumeServiceServer) AttachVolume(context.Context, *VolumeAttachmentIdentifier) (*VolumeAttachment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachVolume not implemented")
}
func (UnimplementedVolumeServiceServer) DetachVolume(context.Context, *VolumeAttachmentIdentifier) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachVolume not implemented")
}
func (UnimplementedVolumeServiceServer) mustEmbedUnimplementedVolumeServiceServer() {}

// UnsafeVolumeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VolumeServiceServer will
// result in compilation errors.
type UnsafeVolumeServiceServer interface {
	mustEmbedUnimplementedVolumeServiceServer()
}

func RegisterVolumeServiceServer(s grpc.ServiceRegistrar, srv VolumeServiceServer) {
	s.RegisterService(&VolumeService_ServiceDesc, srv)
}

func _VolumeService_GetVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).GetVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolumeService_GetVolume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).GetVolume(ctx, req.(*GetVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_ListVolumes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVolumesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).ListVolumes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolumeService_ListVolumes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).ListVolumes(ctx, req.(*ListVolumesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_CreateVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).CreateVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolumeService_CreateVolume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).CreateVolume(ctx, req.(*CreateVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_UpdateVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).UpdateVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolumeService_UpdateVolume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).UpdateVolume(ctx, req.(*UpdateVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_DeleteVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).DeleteVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolumeService_DeleteVolume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).DeleteVolume(ctx, req.(*DeleteVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_ListVolumeAttachments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVolumeAttachmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).ListVolumeAttachments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolumeService_ListVolumeAttachments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).ListVolumeAttachments(ctx, req.(*ListVolumeAttachmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_GetVolumeAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeAttachmentIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).GetVolumeAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolumeService_GetVolumeAttachment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).GetVolumeAttachment(ctx, req.(*VolumeAttachmentIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_AttachVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeAttachmentIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).AttachVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolumeService_AttachVolume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).AttachVolume(ctx, req.(*VolumeAttachmentIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _VolumeService_DetachVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeAttachmentIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumeServiceServer).DetachVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VolumeService_DetachVolume_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumeServiceServer).DetachVolume(ctx, req.(*VolumeAttachmentIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

// VolumeService_ServiceDesc is the grpc.ServiceDesc for VolumeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VolumeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VolumeService",
	HandlerType: (*VolumeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVolume",
			Handler:    _VolumeService_GetVolume_Handler,
		},
		{
			MethodName: "ListVolumes",
			Handler:    _VolumeService_ListVolumes_Handler,
		},
		{
			MethodName: "CreateVolume",
			Handler:    _VolumeService_CreateVolume_Handler,
		},
		{
			MethodName: "UpdateVolume",
			Handler:    _VolumeService_UpdateVolume_Handler,
		},
		{
			MethodName: "DeleteVolume",
			Handler:    _VolumeService_DeleteVolume_Handler,
		},
		{
			MethodName: "ListVolumeAttachments",
			Handler:    _VolumeService_ListVolumeAttachments_Handler,
		},
		{
			MethodName: "GetVolumeAttachment",
			Handler:    _VolumeService_GetVolumeAttachment_Handler,
		},
		{
			MethodName: "AttachVolume",
			Handler:    _VolumeService_AttachVolume_Handler,
		},
		{
			MethodName: "DetachVolume",
			Handler:    _VolumeService_DetachVolume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "minivirt/volume.proto",
}
