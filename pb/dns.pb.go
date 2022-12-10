// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: minivirt/dns.proto

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

type DNSRecordIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *DNSRecordIdentifier) Reset() {
	*x = DNSRecordIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minivirt_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSRecordIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSRecordIdentifier) ProtoMessage() {}

func (x *DNSRecordIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_minivirt_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSRecordIdentifier.ProtoReflect.Descriptor instead.
func (*DNSRecordIdentifier) Descriptor() ([]byte, []int) {
	return file_minivirt_dns_proto_rawDescGZIP(), []int{0}
}

func (x *DNSRecordIdentifier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DNSRecordIdentifier) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type DNSRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type    string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Ttl     uint64   `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Records []string `protobuf:"bytes,4,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *DNSRecord) Reset() {
	*x = DNSRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minivirt_dns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSRecord) ProtoMessage() {}

func (x *DNSRecord) ProtoReflect() protoreflect.Message {
	mi := &file_minivirt_dns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSRecord.ProtoReflect.Descriptor instead.
func (*DNSRecord) Descriptor() ([]byte, []int) {
	return file_minivirt_dns_proto_rawDescGZIP(), []int{1}
}

func (x *DNSRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DNSRecord) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *DNSRecord) GetTtl() uint64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *DNSRecord) GetRecords() []string {
	if x != nil {
		return x.Records
	}
	return nil
}

type ListDNSRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListDNSRecordsRequest) Reset() {
	*x = ListDNSRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minivirt_dns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDNSRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDNSRecordsRequest) ProtoMessage() {}

func (x *ListDNSRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_minivirt_dns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDNSRecordsRequest.ProtoReflect.Descriptor instead.
func (*ListDNSRecordsRequest) Descriptor() ([]byte, []int) {
	return file_minivirt_dns_proto_rawDescGZIP(), []int{2}
}

type ListDNSRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsRecords []*DNSRecord `protobuf:"bytes,1,rep,name=dns_records,json=dnsRecords,proto3" json:"dns_records,omitempty"`
}

func (x *ListDNSRecordsResponse) Reset() {
	*x = ListDNSRecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minivirt_dns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDNSRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDNSRecordsResponse) ProtoMessage() {}

func (x *ListDNSRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_minivirt_dns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDNSRecordsResponse.ProtoReflect.Descriptor instead.
func (*ListDNSRecordsResponse) Descriptor() ([]byte, []int) {
	return file_minivirt_dns_proto_rawDescGZIP(), []int{3}
}

func (x *ListDNSRecordsResponse) GetDnsRecords() []*DNSRecord {
	if x != nil {
		return x.DnsRecords
	}
	return nil
}

type PutDNSRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsRecord *DNSRecord `protobuf:"bytes,1,opt,name=dns_record,json=dnsRecord,proto3" json:"dns_record,omitempty"`
}

func (x *PutDNSRecordRequest) Reset() {
	*x = PutDNSRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minivirt_dns_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutDNSRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutDNSRecordRequest) ProtoMessage() {}

func (x *PutDNSRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_minivirt_dns_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutDNSRecordRequest.ProtoReflect.Descriptor instead.
func (*PutDNSRecordRequest) Descriptor() ([]byte, []int) {
	return file_minivirt_dns_proto_rawDescGZIP(), []int{4}
}

func (x *PutDNSRecordRequest) GetDnsRecord() *DNSRecord {
	if x != nil {
		return x.DnsRecord
	}
	return nil
}

var File_minivirt_dns_proto protoreflect.FileDescriptor

var file_minivirt_dns_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x69, 0x6e, 0x69, 0x76, 0x69, 0x72, 0x74, 0x2f, 0x64, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3d, 0x0a, 0x13, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x5f, 0x0a, 0x09, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x22, 0x17, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x16, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0b, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x44, 0x4e, 0x53, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0a, 0x64, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x22, 0x40, 0x0a, 0x13, 0x50, 0x75, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0a, 0x64, 0x6e, 0x73, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x44,
	0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x09, 0x64, 0x6e, 0x73, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x32, 0xf5, 0x01, 0x0a, 0x03, 0x44, 0x4e, 0x53, 0x12, 0x32, 0x0a, 0x0c, 0x47,
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
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_minivirt_dns_proto_rawDescOnce sync.Once
	file_minivirt_dns_proto_rawDescData = file_minivirt_dns_proto_rawDesc
)

func file_minivirt_dns_proto_rawDescGZIP() []byte {
	file_minivirt_dns_proto_rawDescOnce.Do(func() {
		file_minivirt_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_minivirt_dns_proto_rawDescData)
	})
	return file_minivirt_dns_proto_rawDescData
}

var file_minivirt_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_minivirt_dns_proto_goTypes = []interface{}{
	(*DNSRecordIdentifier)(nil),    // 0: DNSRecordIdentifier
	(*DNSRecord)(nil),              // 1: DNSRecord
	(*ListDNSRecordsRequest)(nil),  // 2: ListDNSRecordsRequest
	(*ListDNSRecordsResponse)(nil), // 3: ListDNSRecordsResponse
	(*PutDNSRecordRequest)(nil),    // 4: PutDNSRecordRequest
	(*emptypb.Empty)(nil),          // 5: google.protobuf.Empty
}
var file_minivirt_dns_proto_depIdxs = []int32{
	1, // 0: ListDNSRecordsResponse.dns_records:type_name -> DNSRecord
	1, // 1: PutDNSRecordRequest.dns_record:type_name -> DNSRecord
	0, // 2: DNS.GetDNSRecord:input_type -> DNSRecordIdentifier
	2, // 3: DNS.ListDNSRecords:input_type -> ListDNSRecordsRequest
	4, // 4: DNS.PutDNSRecord:input_type -> PutDNSRecordRequest
	0, // 5: DNS.DeleteDNSRecord:input_type -> DNSRecordIdentifier
	1, // 6: DNS.GetDNSRecord:output_type -> DNSRecord
	3, // 7: DNS.ListDNSRecords:output_type -> ListDNSRecordsResponse
	1, // 8: DNS.PutDNSRecord:output_type -> DNSRecord
	5, // 9: DNS.DeleteDNSRecord:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_minivirt_dns_proto_init() }
func file_minivirt_dns_proto_init() {
	if File_minivirt_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_minivirt_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSRecordIdentifier); i {
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
		file_minivirt_dns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSRecord); i {
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
		file_minivirt_dns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDNSRecordsRequest); i {
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
		file_minivirt_dns_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDNSRecordsResponse); i {
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
		file_minivirt_dns_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutDNSRecordRequest); i {
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
			RawDescriptor: file_minivirt_dns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_minivirt_dns_proto_goTypes,
		DependencyIndexes: file_minivirt_dns_proto_depIdxs,
		MessageInfos:      file_minivirt_dns_proto_msgTypes,
	}.Build()
	File_minivirt_dns_proto = out.File
	file_minivirt_dns_proto_rawDesc = nil
	file_minivirt_dns_proto_goTypes = nil
	file_minivirt_dns_proto_depIdxs = nil
}
