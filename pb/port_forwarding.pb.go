// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.3
// source: minivirt/port_forwarding.proto

package pb

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

type PortForwardingIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host       string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Protocol   string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	SourcePort uint32 `protobuf:"varint,3,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
}

func (x *PortForwardingIdentifier) Reset() {
	*x = PortForwardingIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minivirt_port_forwarding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortForwardingIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortForwardingIdentifier) ProtoMessage() {}

func (x *PortForwardingIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_minivirt_port_forwarding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortForwardingIdentifier.ProtoReflect.Descriptor instead.
func (*PortForwardingIdentifier) Descriptor() ([]byte, []int) {
	return file_minivirt_port_forwarding_proto_rawDescGZIP(), []int{0}
}

func (x *PortForwardingIdentifier) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *PortForwardingIdentifier) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *PortForwardingIdentifier) GetSourcePort() uint32 {
	if x != nil {
		return x.SourcePort
	}
	return 0
}

type PortForwarding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol   string `protobuf:"bytes,2,opt,name=protocol,proto3" json:"protocol,omitempty"`
	SourcePort uint32 `protobuf:"varint,3,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	TargetIp   string `protobuf:"bytes,4,opt,name=target_ip,json=targetIp,proto3" json:"target_ip,omitempty"`
	TargetPort uint32 `protobuf:"varint,5,opt,name=target_port,json=targetPort,proto3" json:"target_port,omitempty"`
}

func (x *PortForwarding) Reset() {
	*x = PortForwarding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minivirt_port_forwarding_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortForwarding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortForwarding) ProtoMessage() {}

func (x *PortForwarding) ProtoReflect() protoreflect.Message {
	mi := &file_minivirt_port_forwarding_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortForwarding.ProtoReflect.Descriptor instead.
func (*PortForwarding) Descriptor() ([]byte, []int) {
	return file_minivirt_port_forwarding_proto_rawDescGZIP(), []int{1}
}

func (x *PortForwarding) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *PortForwarding) GetSourcePort() uint32 {
	if x != nil {
		return x.SourcePort
	}
	return 0
}

func (x *PortForwarding) GetTargetIp() string {
	if x != nil {
		return x.TargetIp
	}
	return ""
}

func (x *PortForwarding) GetTargetPort() uint32 {
	if x != nil {
		return x.TargetPort
	}
	return 0
}

type ListPortForwardingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *ListPortForwardingsRequest) Reset() {
	*x = ListPortForwardingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minivirt_port_forwarding_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPortForwardingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPortForwardingsRequest) ProtoMessage() {}

func (x *ListPortForwardingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_minivirt_port_forwarding_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPortForwardingsRequest.ProtoReflect.Descriptor instead.
func (*ListPortForwardingsRequest) Descriptor() ([]byte, []int) {
	return file_minivirt_port_forwarding_proto_rawDescGZIP(), []int{2}
}

func (x *ListPortForwardingsRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type ListPortForwardingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortForwardings []*PortForwarding `protobuf:"bytes,1,rep,name=port_forwardings,json=portForwardings,proto3" json:"port_forwardings,omitempty"`
}

func (x *ListPortForwardingsResponse) Reset() {
	*x = ListPortForwardingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minivirt_port_forwarding_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPortForwardingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPortForwardingsResponse) ProtoMessage() {}

func (x *ListPortForwardingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_minivirt_port_forwarding_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPortForwardingsResponse.ProtoReflect.Descriptor instead.
func (*ListPortForwardingsResponse) Descriptor() ([]byte, []int) {
	return file_minivirt_port_forwarding_proto_rawDescGZIP(), []int{3}
}

func (x *ListPortForwardingsResponse) GetPortForwardings() []*PortForwarding {
	if x != nil {
		return x.PortForwardings
	}
	return nil
}

type PutPortForwardingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PortForwarding *PortForwarding `protobuf:"bytes,1,opt,name=port_forwarding,json=portForwarding,proto3" json:"port_forwarding,omitempty"`
	Host           string          `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *PutPortForwardingRequest) Reset() {
	*x = PutPortForwardingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minivirt_port_forwarding_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutPortForwardingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutPortForwardingRequest) ProtoMessage() {}

func (x *PutPortForwardingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_minivirt_port_forwarding_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutPortForwardingRequest.ProtoReflect.Descriptor instead.
func (*PutPortForwardingRequest) Descriptor() ([]byte, []int) {
	return file_minivirt_port_forwarding_proto_rawDescGZIP(), []int{4}
}

func (x *PutPortForwardingRequest) GetPortForwarding() *PortForwarding {
	if x != nil {
		return x.PortForwarding
	}
	return nil
}

func (x *PutPortForwardingRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

var File_minivirt_port_forwarding_proto protoreflect.FileDescriptor

var file_minivirt_port_forwarding_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x69, 0x6e, 0x69, 0x76, 0x69, 0x72, 0x74, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a,
	0x18, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x0e, 0x50,
	0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x30, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0x59, 0x0a, 0x1b, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x10, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x70, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x68, 0x0a, 0x18, 0x50, 0x75, 0x74, 0x50, 0x6f, 0x72, 0x74,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x38, 0x0a, 0x0f, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x6f, 0x72,
	0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0e, 0x70, 0x6f, 0x72,
	0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x32,
	0xbe, 0x02, 0x0a, 0x15, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x19,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x50, 0x6f, 0x72, 0x74,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x13,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x73, 0x12, 0x1b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x11, 0x50, 0x75, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x50, 0x75, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x46,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x72,
	0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x50, 0x6f,
	0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_minivirt_port_forwarding_proto_rawDescOnce sync.Once
	file_minivirt_port_forwarding_proto_rawDescData = file_minivirt_port_forwarding_proto_rawDesc
)

func file_minivirt_port_forwarding_proto_rawDescGZIP() []byte {
	file_minivirt_port_forwarding_proto_rawDescOnce.Do(func() {
		file_minivirt_port_forwarding_proto_rawDescData = protoimpl.X.CompressGZIP(file_minivirt_port_forwarding_proto_rawDescData)
	})
	return file_minivirt_port_forwarding_proto_rawDescData
}

var file_minivirt_port_forwarding_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_minivirt_port_forwarding_proto_goTypes = []interface{}{
	(*PortForwardingIdentifier)(nil),    // 0: PortForwardingIdentifier
	(*PortForwarding)(nil),              // 1: PortForwarding
	(*ListPortForwardingsRequest)(nil),  // 2: ListPortForwardingsRequest
	(*ListPortForwardingsResponse)(nil), // 3: ListPortForwardingsResponse
	(*PutPortForwardingRequest)(nil),    // 4: PutPortForwardingRequest
	(*emptypb.Empty)(nil),               // 5: google.protobuf.Empty
}
var file_minivirt_port_forwarding_proto_depIdxs = []int32{
	1, // 0: ListPortForwardingsResponse.port_forwardings:type_name -> PortForwarding
	1, // 1: PutPortForwardingRequest.port_forwarding:type_name -> PortForwarding
	0, // 2: PortForwardingService.GetPortForwarding:input_type -> PortForwardingIdentifier
	2, // 3: PortForwardingService.ListPortForwardings:input_type -> ListPortForwardingsRequest
	4, // 4: PortForwardingService.PutPortForwarding:input_type -> PutPortForwardingRequest
	0, // 5: PortForwardingService.DeletePortForwarding:input_type -> PortForwardingIdentifier
	1, // 6: PortForwardingService.GetPortForwarding:output_type -> PortForwarding
	3, // 7: PortForwardingService.ListPortForwardings:output_type -> ListPortForwardingsResponse
	1, // 8: PortForwardingService.PutPortForwarding:output_type -> PortForwarding
	5, // 9: PortForwardingService.DeletePortForwarding:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_minivirt_port_forwarding_proto_init() }
func file_minivirt_port_forwarding_proto_init() {
	if File_minivirt_port_forwarding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_minivirt_port_forwarding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortForwardingIdentifier); i {
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
		file_minivirt_port_forwarding_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortForwarding); i {
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
		file_minivirt_port_forwarding_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPortForwardingsRequest); i {
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
		file_minivirt_port_forwarding_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPortForwardingsResponse); i {
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
		file_minivirt_port_forwarding_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutPortForwardingRequest); i {
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
			RawDescriptor: file_minivirt_port_forwarding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_minivirt_port_forwarding_proto_goTypes,
		DependencyIndexes: file_minivirt_port_forwarding_proto_depIdxs,
		MessageInfos:      file_minivirt_port_forwarding_proto_msgTypes,
	}.Build()
	File_minivirt_port_forwarding_proto = out.File
	file_minivirt_port_forwarding_proto_rawDesc = nil
	file_minivirt_port_forwarding_proto_goTypes = nil
	file_minivirt_port_forwarding_proto_depIdxs = nil
}
