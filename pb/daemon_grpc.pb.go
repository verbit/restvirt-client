// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// DaemonServiceClient is the client API for DaemonService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaemonServiceClient interface {
	GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*Network, error)
	ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error)
	CreateNetwork(ctx context.Context, in *CreateNetworkRequest, opts ...grpc.CallOption) (*Network, error)
	DeleteNetwork(ctx context.Context, in *DeleteNetworkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StartDomain(ctx context.Context, in *StartDomainRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StopDomain(ctx context.Context, in *StopDomainRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*Domain, error)
	ListDomains(ctx context.Context, in *ListDomainsRequest, opts ...grpc.CallOption) (*ListDomainsResponse, error)
	CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*Domain, error)
	DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DownloadImage(ctx context.Context, in *DownloadImageRequest, opts ...grpc.CallOption) (DaemonService_DownloadImageClient, error)
	GetVolume(ctx context.Context, in *GetVolumeRequest, opts ...grpc.CallOption) (*Volume, error)
	ListVolumes(ctx context.Context, in *ListVolumesRequest, opts ...grpc.CallOption) (*ListVolumesResponse, error)
	CreateVolume(ctx context.Context, in *CreateVolumeRequest, opts ...grpc.CallOption) (*Volume, error)
	UpdateVolume(ctx context.Context, in *UpdateVolumeRequest, opts ...grpc.CallOption) (*Volume, error)
	DeleteVolume(ctx context.Context, in *DeleteVolumeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListVolumeAttachments(ctx context.Context, in *ListVolumeAttachmentsRequest, opts ...grpc.CallOption) (*ListVolumeAttachmentsResponse, error)
	GetVolumeAttachment(ctx context.Context, in *VolumeAttachmentIdentifier, opts ...grpc.CallOption) (*VolumeAttachment, error)
	AttachVolume(ctx context.Context, in *VolumeAttachmentIdentifier, opts ...grpc.CallOption) (*VolumeAttachment, error)
	DetachVolume(ctx context.Context, in *VolumeAttachmentIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetPortForwarding(ctx context.Context, in *PortForwardingIdentifier, opts ...grpc.CallOption) (*PortForwarding, error)
	ListPortForwardings(ctx context.Context, in *ListPortForwardingsRequest, opts ...grpc.CallOption) (*ListPortForwardingsResponse, error)
	PutPortForwarding(ctx context.Context, in *PutPortForwardingRequest, opts ...grpc.CallOption) (*PortForwarding, error)
	DeletePortForwarding(ctx context.Context, in *PortForwardingIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SyncRoutes(ctx context.Context, in *SyncRoutesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type daemonServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDaemonServiceClient(cc grpc.ClientConnInterface) DaemonServiceClient {
	return &daemonServiceClient{cc}
}

func (c *daemonServiceClient) GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*Network, error) {
	out := new(Network)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/GetNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) ListNetworks(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error) {
	out := new(ListNetworksResponse)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/ListNetworks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) CreateNetwork(ctx context.Context, in *CreateNetworkRequest, opts ...grpc.CallOption) (*Network, error) {
	out := new(Network)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/CreateNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) DeleteNetwork(ctx context.Context, in *DeleteNetworkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/DeleteNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) StartDomain(ctx context.Context, in *StartDomainRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/StartDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) StopDomain(ctx context.Context, in *StopDomainRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/StopDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) GetDomain(ctx context.Context, in *GetDomainRequest, opts ...grpc.CallOption) (*Domain, error) {
	out := new(Domain)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/GetDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) ListDomains(ctx context.Context, in *ListDomainsRequest, opts ...grpc.CallOption) (*ListDomainsResponse, error) {
	out := new(ListDomainsResponse)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/ListDomains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*Domain, error) {
	out := new(Domain)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/CreateDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) DeleteDomain(ctx context.Context, in *DeleteDomainRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/DeleteDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) DownloadImage(ctx context.Context, in *DownloadImageRequest, opts ...grpc.CallOption) (DaemonService_DownloadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &DaemonService_ServiceDesc.Streams[0], "/internal.DaemonService/DownloadImage", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonServiceDownloadImageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DaemonService_DownloadImageClient interface {
	Recv() (*ImageChunk, error)
	grpc.ClientStream
}

type daemonServiceDownloadImageClient struct {
	grpc.ClientStream
}

func (x *daemonServiceDownloadImageClient) Recv() (*ImageChunk, error) {
	m := new(ImageChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonServiceClient) GetVolume(ctx context.Context, in *GetVolumeRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/GetVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) ListVolumes(ctx context.Context, in *ListVolumesRequest, opts ...grpc.CallOption) (*ListVolumesResponse, error) {
	out := new(ListVolumesResponse)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/ListVolumes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) CreateVolume(ctx context.Context, in *CreateVolumeRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/CreateVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) UpdateVolume(ctx context.Context, in *UpdateVolumeRequest, opts ...grpc.CallOption) (*Volume, error) {
	out := new(Volume)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/UpdateVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) DeleteVolume(ctx context.Context, in *DeleteVolumeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/DeleteVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) ListVolumeAttachments(ctx context.Context, in *ListVolumeAttachmentsRequest, opts ...grpc.CallOption) (*ListVolumeAttachmentsResponse, error) {
	out := new(ListVolumeAttachmentsResponse)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/ListVolumeAttachments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) GetVolumeAttachment(ctx context.Context, in *VolumeAttachmentIdentifier, opts ...grpc.CallOption) (*VolumeAttachment, error) {
	out := new(VolumeAttachment)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/GetVolumeAttachment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) AttachVolume(ctx context.Context, in *VolumeAttachmentIdentifier, opts ...grpc.CallOption) (*VolumeAttachment, error) {
	out := new(VolumeAttachment)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/AttachVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) DetachVolume(ctx context.Context, in *VolumeAttachmentIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/DetachVolume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) GetPortForwarding(ctx context.Context, in *PortForwardingIdentifier, opts ...grpc.CallOption) (*PortForwarding, error) {
	out := new(PortForwarding)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/GetPortForwarding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) ListPortForwardings(ctx context.Context, in *ListPortForwardingsRequest, opts ...grpc.CallOption) (*ListPortForwardingsResponse, error) {
	out := new(ListPortForwardingsResponse)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/ListPortForwardings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) PutPortForwarding(ctx context.Context, in *PutPortForwardingRequest, opts ...grpc.CallOption) (*PortForwarding, error) {
	out := new(PortForwarding)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/PutPortForwarding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) DeletePortForwarding(ctx context.Context, in *PortForwardingIdentifier, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/DeletePortForwarding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonServiceClient) SyncRoutes(ctx context.Context, in *SyncRoutesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/internal.DaemonService/SyncRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonServiceServer is the server API for DaemonService service.
// All implementations must embed UnimplementedDaemonServiceServer
// for forward compatibility
type DaemonServiceServer interface {
	GetNetwork(context.Context, *GetNetworkRequest) (*Network, error)
	ListNetworks(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error)
	CreateNetwork(context.Context, *CreateNetworkRequest) (*Network, error)
	DeleteNetwork(context.Context, *DeleteNetworkRequest) (*emptypb.Empty, error)
	StartDomain(context.Context, *StartDomainRequest) (*emptypb.Empty, error)
	StopDomain(context.Context, *StopDomainRequest) (*emptypb.Empty, error)
	GetDomain(context.Context, *GetDomainRequest) (*Domain, error)
	ListDomains(context.Context, *ListDomainsRequest) (*ListDomainsResponse, error)
	CreateDomain(context.Context, *CreateDomainRequest) (*Domain, error)
	DeleteDomain(context.Context, *DeleteDomainRequest) (*emptypb.Empty, error)
	DownloadImage(*DownloadImageRequest, DaemonService_DownloadImageServer) error
	GetVolume(context.Context, *GetVolumeRequest) (*Volume, error)
	ListVolumes(context.Context, *ListVolumesRequest) (*ListVolumesResponse, error)
	CreateVolume(context.Context, *CreateVolumeRequest) (*Volume, error)
	UpdateVolume(context.Context, *UpdateVolumeRequest) (*Volume, error)
	DeleteVolume(context.Context, *DeleteVolumeRequest) (*emptypb.Empty, error)
	ListVolumeAttachments(context.Context, *ListVolumeAttachmentsRequest) (*ListVolumeAttachmentsResponse, error)
	GetVolumeAttachment(context.Context, *VolumeAttachmentIdentifier) (*VolumeAttachment, error)
	AttachVolume(context.Context, *VolumeAttachmentIdentifier) (*VolumeAttachment, error)
	DetachVolume(context.Context, *VolumeAttachmentIdentifier) (*emptypb.Empty, error)
	GetPortForwarding(context.Context, *PortForwardingIdentifier) (*PortForwarding, error)
	ListPortForwardings(context.Context, *ListPortForwardingsRequest) (*ListPortForwardingsResponse, error)
	PutPortForwarding(context.Context, *PutPortForwardingRequest) (*PortForwarding, error)
	DeletePortForwarding(context.Context, *PortForwardingIdentifier) (*emptypb.Empty, error)
	SyncRoutes(context.Context, *SyncRoutesRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDaemonServiceServer()
}

// UnimplementedDaemonServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDaemonServiceServer struct {
}

func (UnimplementedDaemonServiceServer) GetNetwork(context.Context, *GetNetworkRequest) (*Network, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNetwork not implemented")
}
func (UnimplementedDaemonServiceServer) ListNetworks(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNetworks not implemented")
}
func (UnimplementedDaemonServiceServer) CreateNetwork(context.Context, *CreateNetworkRequest) (*Network, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNetwork not implemented")
}
func (UnimplementedDaemonServiceServer) DeleteNetwork(context.Context, *DeleteNetworkRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNetwork not implemented")
}
func (UnimplementedDaemonServiceServer) StartDomain(context.Context, *StartDomainRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartDomain not implemented")
}
func (UnimplementedDaemonServiceServer) StopDomain(context.Context, *StopDomainRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopDomain not implemented")
}
func (UnimplementedDaemonServiceServer) GetDomain(context.Context, *GetDomainRequest) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomain not implemented")
}
func (UnimplementedDaemonServiceServer) ListDomains(context.Context, *ListDomainsRequest) (*ListDomainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDomains not implemented")
}
func (UnimplementedDaemonServiceServer) CreateDomain(context.Context, *CreateDomainRequest) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDomain not implemented")
}
func (UnimplementedDaemonServiceServer) DeleteDomain(context.Context, *DeleteDomainRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDomain not implemented")
}
func (UnimplementedDaemonServiceServer) DownloadImage(*DownloadImageRequest, DaemonService_DownloadImageServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadImage not implemented")
}
func (UnimplementedDaemonServiceServer) GetVolume(context.Context, *GetVolumeRequest) (*Volume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVolume not implemented")
}
func (UnimplementedDaemonServiceServer) ListVolumes(context.Context, *ListVolumesRequest) (*ListVolumesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVolumes not implemented")
}
func (UnimplementedDaemonServiceServer) CreateVolume(context.Context, *CreateVolumeRequest) (*Volume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVolume not implemented")
}
func (UnimplementedDaemonServiceServer) UpdateVolume(context.Context, *UpdateVolumeRequest) (*Volume, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVolume not implemented")
}
func (UnimplementedDaemonServiceServer) DeleteVolume(context.Context, *DeleteVolumeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVolume not implemented")
}
func (UnimplementedDaemonServiceServer) ListVolumeAttachments(context.Context, *ListVolumeAttachmentsRequest) (*ListVolumeAttachmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVolumeAttachments not implemented")
}
func (UnimplementedDaemonServiceServer) GetVolumeAttachment(context.Context, *VolumeAttachmentIdentifier) (*VolumeAttachment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVolumeAttachment not implemented")
}
func (UnimplementedDaemonServiceServer) AttachVolume(context.Context, *VolumeAttachmentIdentifier) (*VolumeAttachment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachVolume not implemented")
}
func (UnimplementedDaemonServiceServer) DetachVolume(context.Context, *VolumeAttachmentIdentifier) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachVolume not implemented")
}
func (UnimplementedDaemonServiceServer) GetPortForwarding(context.Context, *PortForwardingIdentifier) (*PortForwarding, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPortForwarding not implemented")
}
func (UnimplementedDaemonServiceServer) ListPortForwardings(context.Context, *ListPortForwardingsRequest) (*ListPortForwardingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPortForwardings not implemented")
}
func (UnimplementedDaemonServiceServer) PutPortForwarding(context.Context, *PutPortForwardingRequest) (*PortForwarding, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPortForwarding not implemented")
}
func (UnimplementedDaemonServiceServer) DeletePortForwarding(context.Context, *PortForwardingIdentifier) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePortForwarding not implemented")
}
func (UnimplementedDaemonServiceServer) SyncRoutes(context.Context, *SyncRoutesRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncRoutes not implemented")
}
func (UnimplementedDaemonServiceServer) mustEmbedUnimplementedDaemonServiceServer() {}

// UnsafeDaemonServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DaemonServiceServer will
// result in compilation errors.
type UnsafeDaemonServiceServer interface {
	mustEmbedUnimplementedDaemonServiceServer()
}

func RegisterDaemonServiceServer(s grpc.ServiceRegistrar, srv DaemonServiceServer) {
	s.RegisterService(&DaemonService_ServiceDesc, srv)
}

func _DaemonService_GetNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).GetNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/GetNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).GetNetwork(ctx, req.(*GetNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_ListNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).ListNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/ListNetworks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).ListNetworks(ctx, req.(*ListNetworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_CreateNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).CreateNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/CreateNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).CreateNetwork(ctx, req.(*CreateNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_DeleteNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).DeleteNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/DeleteNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).DeleteNetwork(ctx, req.(*DeleteNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_StartDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).StartDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/StartDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).StartDomain(ctx, req.(*StartDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_StopDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).StopDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/StopDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).StopDomain(ctx, req.(*StopDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_GetDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).GetDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/GetDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).GetDomain(ctx, req.(*GetDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_ListDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDomainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).ListDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/ListDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).ListDomains(ctx, req.(*ListDomainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_CreateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).CreateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/CreateDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).CreateDomain(ctx, req.(*CreateDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_DeleteDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).DeleteDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/DeleteDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).DeleteDomain(ctx, req.(*DeleteDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_DownloadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadImageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServiceServer).DownloadImage(m, &daemonServiceDownloadImageServer{stream})
}

type DaemonService_DownloadImageServer interface {
	Send(*ImageChunk) error
	grpc.ServerStream
}

type daemonServiceDownloadImageServer struct {
	grpc.ServerStream
}

func (x *daemonServiceDownloadImageServer) Send(m *ImageChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _DaemonService_GetVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).GetVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/GetVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).GetVolume(ctx, req.(*GetVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_ListVolumes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVolumesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).ListVolumes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/ListVolumes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).ListVolumes(ctx, req.(*ListVolumesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_CreateVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).CreateVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/CreateVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).CreateVolume(ctx, req.(*CreateVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_UpdateVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).UpdateVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/UpdateVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).UpdateVolume(ctx, req.(*UpdateVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_DeleteVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).DeleteVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/DeleteVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).DeleteVolume(ctx, req.(*DeleteVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_ListVolumeAttachments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVolumeAttachmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).ListVolumeAttachments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/ListVolumeAttachments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).ListVolumeAttachments(ctx, req.(*ListVolumeAttachmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_GetVolumeAttachment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeAttachmentIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).GetVolumeAttachment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/GetVolumeAttachment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).GetVolumeAttachment(ctx, req.(*VolumeAttachmentIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_AttachVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeAttachmentIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).AttachVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/AttachVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).AttachVolume(ctx, req.(*VolumeAttachmentIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_DetachVolume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeAttachmentIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).DetachVolume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/DetachVolume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).DetachVolume(ctx, req.(*VolumeAttachmentIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_GetPortForwarding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortForwardingIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).GetPortForwarding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/GetPortForwarding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).GetPortForwarding(ctx, req.(*PortForwardingIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_ListPortForwardings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPortForwardingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).ListPortForwardings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/ListPortForwardings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).ListPortForwardings(ctx, req.(*ListPortForwardingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_PutPortForwarding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutPortForwardingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).PutPortForwarding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/PutPortForwarding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).PutPortForwarding(ctx, req.(*PutPortForwardingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_DeletePortForwarding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PortForwardingIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).DeletePortForwarding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/DeletePortForwarding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).DeletePortForwarding(ctx, req.(*PortForwardingIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

func _DaemonService_SyncRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncRoutesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServiceServer).SyncRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.DaemonService/SyncRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServiceServer).SyncRoutes(ctx, req.(*SyncRoutesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DaemonService_ServiceDesc is the grpc.ServiceDesc for DaemonService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DaemonService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "internal.DaemonService",
	HandlerType: (*DaemonServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNetwork",
			Handler:    _DaemonService_GetNetwork_Handler,
		},
		{
			MethodName: "ListNetworks",
			Handler:    _DaemonService_ListNetworks_Handler,
		},
		{
			MethodName: "CreateNetwork",
			Handler:    _DaemonService_CreateNetwork_Handler,
		},
		{
			MethodName: "DeleteNetwork",
			Handler:    _DaemonService_DeleteNetwork_Handler,
		},
		{
			MethodName: "StartDomain",
			Handler:    _DaemonService_StartDomain_Handler,
		},
		{
			MethodName: "StopDomain",
			Handler:    _DaemonService_StopDomain_Handler,
		},
		{
			MethodName: "GetDomain",
			Handler:    _DaemonService_GetDomain_Handler,
		},
		{
			MethodName: "ListDomains",
			Handler:    _DaemonService_ListDomains_Handler,
		},
		{
			MethodName: "CreateDomain",
			Handler:    _DaemonService_CreateDomain_Handler,
		},
		{
			MethodName: "DeleteDomain",
			Handler:    _DaemonService_DeleteDomain_Handler,
		},
		{
			MethodName: "GetVolume",
			Handler:    _DaemonService_GetVolume_Handler,
		},
		{
			MethodName: "ListVolumes",
			Handler:    _DaemonService_ListVolumes_Handler,
		},
		{
			MethodName: "CreateVolume",
			Handler:    _DaemonService_CreateVolume_Handler,
		},
		{
			MethodName: "UpdateVolume",
			Handler:    _DaemonService_UpdateVolume_Handler,
		},
		{
			MethodName: "DeleteVolume",
			Handler:    _DaemonService_DeleteVolume_Handler,
		},
		{
			MethodName: "ListVolumeAttachments",
			Handler:    _DaemonService_ListVolumeAttachments_Handler,
		},
		{
			MethodName: "GetVolumeAttachment",
			Handler:    _DaemonService_GetVolumeAttachment_Handler,
		},
		{
			MethodName: "AttachVolume",
			Handler:    _DaemonService_AttachVolume_Handler,
		},
		{
			MethodName: "DetachVolume",
			Handler:    _DaemonService_DetachVolume_Handler,
		},
		{
			MethodName: "GetPortForwarding",
			Handler:    _DaemonService_GetPortForwarding_Handler,
		},
		{
			MethodName: "ListPortForwardings",
			Handler:    _DaemonService_ListPortForwardings_Handler,
		},
		{
			MethodName: "PutPortForwarding",
			Handler:    _DaemonService_PutPortForwarding_Handler,
		},
		{
			MethodName: "DeletePortForwarding",
			Handler:    _DaemonService_DeletePortForwarding_Handler,
		},
		{
			MethodName: "SyncRoutes",
			Handler:    _DaemonService_SyncRoutes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DownloadImage",
			Handler:       _DaemonService_DownloadImage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "minivirt/daemon.proto",
}