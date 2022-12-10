// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: minivirt/controller.proto

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

type SyncRoutesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SyncRoutesRequest) Reset() {
	*x = SyncRoutesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minivirt_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRoutesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRoutesRequest) ProtoMessage() {}

func (x *SyncRoutesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_minivirt_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRoutesRequest.ProtoReflect.Descriptor instead.
func (*SyncRoutesRequest) Descriptor() ([]byte, []int) {
	return file_minivirt_controller_proto_rawDescGZIP(), []int{0}
}

var File_minivirt_controller_proto protoreflect.FileDescriptor

var file_minivirt_controller_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x69, 0x6e, 0x69, 0x76, 0x69, 0x72, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6d, 0x69, 0x6e, 0x69, 0x76, 0x69,
	0x72, 0x74, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x6d, 0x69, 0x6e, 0x69, 0x76, 0x69, 0x72, 0x74, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x69, 0x6e, 0x69, 0x76, 0x69, 0x72, 0x74,
	0x2f, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6d, 0x69, 0x6e, 0x69, 0x76, 0x69, 0x72, 0x74,
	0x2f, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6d, 0x69, 0x6e, 0x69,
	0x76, 0x69, 0x72, 0x74, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x13, 0x0a, 0x11, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xce, 0x11, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x2e, 0x44, 0x4e,
	0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x1a, 0x0a, 0x2e, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x12, 0x16, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c, 0x50, 0x75, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x2e, 0x50, 0x75, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x44, 0x4e, 0x53,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x2e, 0x44, 0x4e,
	0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x14, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x08, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x15, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3c,
	0x0a, 0x0b, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x13, 0x2e,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a,
	0x53, 0x74, 0x6f, 0x70, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x53, 0x74, 0x6f,
	0x70, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x73, 0x12, 0x13, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x0d, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x15, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x12, 0x29, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x2f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22,
	0x00, 0x12, 0x2f, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x12, 0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x07, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x58, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x1a, 0x11, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0c, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x1a, 0x11, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x44, 0x65, 0x74, 0x61, 0x63,
	0x68, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1b, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x41,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x0f,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x11, 0x50, 0x75, 0x74, 0x50, 0x6f, 0x72, 0x74,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x50, 0x75, 0x74,
	0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x19, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x0b, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12,
	0x17, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x22,
	0x00, 0x12, 0x43, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x12, 0x10, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x1a, 0x06, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x08, 0x50, 0x75, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x10,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x53, 0x79,
	0x6e, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_minivirt_controller_proto_rawDescOnce sync.Once
	file_minivirt_controller_proto_rawDescData = file_minivirt_controller_proto_rawDesc
)

func file_minivirt_controller_proto_rawDescGZIP() []byte {
	file_minivirt_controller_proto_rawDescOnce.Do(func() {
		file_minivirt_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_minivirt_controller_proto_rawDescData)
	})
	return file_minivirt_controller_proto_rawDescData
}

var file_minivirt_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_minivirt_controller_proto_goTypes = []interface{}{
	(*SyncRoutesRequest)(nil),             // 0: SyncRoutesRequest
	(*DNSRecordIdentifier)(nil),           // 1: DNSRecordIdentifier
	(*ListDNSRecordsRequest)(nil),         // 2: ListDNSRecordsRequest
	(*PutDNSRecordRequest)(nil),           // 3: PutDNSRecordRequest
	(*GetNetworkRequest)(nil),             // 4: GetNetworkRequest
	(*ListNetworksRequest)(nil),           // 5: ListNetworksRequest
	(*CreateNetworkRequest)(nil),          // 6: CreateNetworkRequest
	(*DeleteNetworkRequest)(nil),          // 7: DeleteNetworkRequest
	(*StartDomainRequest)(nil),            // 8: StartDomainRequest
	(*StopDomainRequest)(nil),             // 9: StopDomainRequest
	(*GetDomainRequest)(nil),              // 10: GetDomainRequest
	(*ListDomainsRequest)(nil),            // 11: ListDomainsRequest
	(*CreateDomainRequest)(nil),           // 12: CreateDomainRequest
	(*DeleteDomainRequest)(nil),           // 13: DeleteDomainRequest
	(*DownloadImageRequest)(nil),          // 14: DownloadImageRequest
	(*GetVolumeRequest)(nil),              // 15: GetVolumeRequest
	(*ListVolumesRequest)(nil),            // 16: ListVolumesRequest
	(*CreateVolumeRequest)(nil),           // 17: CreateVolumeRequest
	(*UpdateVolumeRequest)(nil),           // 18: UpdateVolumeRequest
	(*DeleteVolumeRequest)(nil),           // 19: DeleteVolumeRequest
	(*ListVolumeAttachmentsRequest)(nil),  // 20: ListVolumeAttachmentsRequest
	(*VolumeAttachmentIdentifier)(nil),    // 21: VolumeAttachmentIdentifier
	(*PortForwardingIdentifier)(nil),      // 22: PortForwardingIdentifier
	(*ListPortForwardingsRequest)(nil),    // 23: ListPortForwardingsRequest
	(*PutPortForwardingRequest)(nil),      // 24: PutPortForwardingRequest
	(*RouteTableIdentifier)(nil),          // 25: RouteTableIdentifier
	(*ListRouteTablesRequest)(nil),        // 26: ListRouteTablesRequest
	(*CreateRouteTableRequest)(nil),       // 27: CreateRouteTableRequest
	(*RouteIdentifier)(nil),               // 28: RouteIdentifier
	(*ListRoutesRequest)(nil),             // 29: ListRoutesRequest
	(*PutRouteRequest)(nil),               // 30: PutRouteRequest
	(*DNSRecord)(nil),                     // 31: DNSRecord
	(*ListDNSRecordsResponse)(nil),        // 32: ListDNSRecordsResponse
	(*emptypb.Empty)(nil),                 // 33: google.protobuf.Empty
	(*Network)(nil),                       // 34: Network
	(*ListNetworksResponse)(nil),          // 35: ListNetworksResponse
	(*Domain)(nil),                        // 36: Domain
	(*ListDomainsResponse)(nil),           // 37: ListDomainsResponse
	(*ImageChunk)(nil),                    // 38: ImageChunk
	(*Volume)(nil),                        // 39: Volume
	(*ListVolumesResponse)(nil),           // 40: ListVolumesResponse
	(*ListVolumeAttachmentsResponse)(nil), // 41: ListVolumeAttachmentsResponse
	(*VolumeAttachment)(nil),              // 42: VolumeAttachment
	(*PortForwarding)(nil),                // 43: PortForwarding
	(*ListPortForwardingsResponse)(nil),   // 44: ListPortForwardingsResponse
	(*RouteTable)(nil),                    // 45: RouteTable
	(*ListRouteTablesResponse)(nil),       // 46: ListRouteTablesResponse
	(*Route)(nil),                         // 47: Route
	(*ListRoutesResponse)(nil),            // 48: ListRoutesResponse
}
var file_minivirt_controller_proto_depIdxs = []int32{
	1,  // 0: ControllerService.GetDNSRecord:input_type -> DNSRecordIdentifier
	2,  // 1: ControllerService.ListDNSRecords:input_type -> ListDNSRecordsRequest
	3,  // 2: ControllerService.PutDNSRecord:input_type -> PutDNSRecordRequest
	1,  // 3: ControllerService.DeleteDNSRecord:input_type -> DNSRecordIdentifier
	4,  // 4: ControllerService.GetNetwork:input_type -> GetNetworkRequest
	5,  // 5: ControllerService.ListNetworks:input_type -> ListNetworksRequest
	6,  // 6: ControllerService.CreateNetwork:input_type -> CreateNetworkRequest
	7,  // 7: ControllerService.DeleteNetwork:input_type -> DeleteNetworkRequest
	8,  // 8: ControllerService.StartDomain:input_type -> StartDomainRequest
	9,  // 9: ControllerService.StopDomain:input_type -> StopDomainRequest
	10, // 10: ControllerService.GetDomain:input_type -> GetDomainRequest
	11, // 11: ControllerService.ListDomains:input_type -> ListDomainsRequest
	12, // 12: ControllerService.CreateDomain:input_type -> CreateDomainRequest
	13, // 13: ControllerService.DeleteDomain:input_type -> DeleteDomainRequest
	14, // 14: ControllerService.DownloadImage:input_type -> DownloadImageRequest
	15, // 15: ControllerService.GetVolume:input_type -> GetVolumeRequest
	16, // 16: ControllerService.ListVolumes:input_type -> ListVolumesRequest
	17, // 17: ControllerService.CreateVolume:input_type -> CreateVolumeRequest
	18, // 18: ControllerService.UpdateVolume:input_type -> UpdateVolumeRequest
	19, // 19: ControllerService.DeleteVolume:input_type -> DeleteVolumeRequest
	20, // 20: ControllerService.ListVolumeAttachments:input_type -> ListVolumeAttachmentsRequest
	21, // 21: ControllerService.GetVolumeAttachment:input_type -> VolumeAttachmentIdentifier
	21, // 22: ControllerService.AttachVolume:input_type -> VolumeAttachmentIdentifier
	21, // 23: ControllerService.DetachVolume:input_type -> VolumeAttachmentIdentifier
	22, // 24: ControllerService.GetPortForwarding:input_type -> PortForwardingIdentifier
	23, // 25: ControllerService.ListPortForwardings:input_type -> ListPortForwardingsRequest
	24, // 26: ControllerService.PutPortForwarding:input_type -> PutPortForwardingRequest
	22, // 27: ControllerService.DeletePortForwarding:input_type -> PortForwardingIdentifier
	25, // 28: ControllerService.GetRouteTable:input_type -> RouteTableIdentifier
	26, // 29: ControllerService.ListRouteTables:input_type -> ListRouteTablesRequest
	27, // 30: ControllerService.CreateRouteTable:input_type -> CreateRouteTableRequest
	25, // 31: ControllerService.DeleteRouteTable:input_type -> RouteTableIdentifier
	28, // 32: ControllerService.GetRoute:input_type -> RouteIdentifier
	29, // 33: ControllerService.ListRoutes:input_type -> ListRoutesRequest
	30, // 34: ControllerService.PutRoute:input_type -> PutRouteRequest
	28, // 35: ControllerService.DeleteRoute:input_type -> RouteIdentifier
	0,  // 36: ControllerService.SyncRoutes:input_type -> SyncRoutesRequest
	31, // 37: ControllerService.GetDNSRecord:output_type -> DNSRecord
	32, // 38: ControllerService.ListDNSRecords:output_type -> ListDNSRecordsResponse
	31, // 39: ControllerService.PutDNSRecord:output_type -> DNSRecord
	33, // 40: ControllerService.DeleteDNSRecord:output_type -> google.protobuf.Empty
	34, // 41: ControllerService.GetNetwork:output_type -> Network
	35, // 42: ControllerService.ListNetworks:output_type -> ListNetworksResponse
	34, // 43: ControllerService.CreateNetwork:output_type -> Network
	33, // 44: ControllerService.DeleteNetwork:output_type -> google.protobuf.Empty
	33, // 45: ControllerService.StartDomain:output_type -> google.protobuf.Empty
	33, // 46: ControllerService.StopDomain:output_type -> google.protobuf.Empty
	36, // 47: ControllerService.GetDomain:output_type -> Domain
	37, // 48: ControllerService.ListDomains:output_type -> ListDomainsResponse
	36, // 49: ControllerService.CreateDomain:output_type -> Domain
	33, // 50: ControllerService.DeleteDomain:output_type -> google.protobuf.Empty
	38, // 51: ControllerService.DownloadImage:output_type -> ImageChunk
	39, // 52: ControllerService.GetVolume:output_type -> Volume
	40, // 53: ControllerService.ListVolumes:output_type -> ListVolumesResponse
	39, // 54: ControllerService.CreateVolume:output_type -> Volume
	39, // 55: ControllerService.UpdateVolume:output_type -> Volume
	33, // 56: ControllerService.DeleteVolume:output_type -> google.protobuf.Empty
	41, // 57: ControllerService.ListVolumeAttachments:output_type -> ListVolumeAttachmentsResponse
	42, // 58: ControllerService.GetVolumeAttachment:output_type -> VolumeAttachment
	42, // 59: ControllerService.AttachVolume:output_type -> VolumeAttachment
	33, // 60: ControllerService.DetachVolume:output_type -> google.protobuf.Empty
	43, // 61: ControllerService.GetPortForwarding:output_type -> PortForwarding
	44, // 62: ControllerService.ListPortForwardings:output_type -> ListPortForwardingsResponse
	43, // 63: ControllerService.PutPortForwarding:output_type -> PortForwarding
	33, // 64: ControllerService.DeletePortForwarding:output_type -> google.protobuf.Empty
	45, // 65: ControllerService.GetRouteTable:output_type -> RouteTable
	46, // 66: ControllerService.ListRouteTables:output_type -> ListRouteTablesResponse
	45, // 67: ControllerService.CreateRouteTable:output_type -> RouteTable
	33, // 68: ControllerService.DeleteRouteTable:output_type -> google.protobuf.Empty
	47, // 69: ControllerService.GetRoute:output_type -> Route
	48, // 70: ControllerService.ListRoutes:output_type -> ListRoutesResponse
	47, // 71: ControllerService.PutRoute:output_type -> Route
	33, // 72: ControllerService.DeleteRoute:output_type -> google.protobuf.Empty
	33, // 73: ControllerService.SyncRoutes:output_type -> google.protobuf.Empty
	37, // [37:74] is the sub-list for method output_type
	0,  // [0:37] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_minivirt_controller_proto_init() }
func file_minivirt_controller_proto_init() {
	if File_minivirt_controller_proto != nil {
		return
	}
	file_minivirt_domain_proto_init()
	file_minivirt_volume_proto_init()
	file_minivirt_port_forwarding_proto_init()
	file_minivirt_dns_proto_init()
	file_minivirt_route_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_minivirt_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRoutesRequest); i {
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
			RawDescriptor: file_minivirt_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_minivirt_controller_proto_goTypes,
		DependencyIndexes: file_minivirt_controller_proto_depIdxs,
		MessageInfos:      file_minivirt_controller_proto_msgTypes,
	}.Build()
	File_minivirt_controller_proto = out.File
	file_minivirt_controller_proto_rawDesc = nil
	file_minivirt_controller_proto_goTypes = nil
	file_minivirt_controller_proto_depIdxs = nil
}