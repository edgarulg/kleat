// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: storage/persistent_storage.proto

package storage

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration of Spinnaker's persistent storage.
type PersistentStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gcs    *Gcs    `protobuf:"bytes,1,opt,name=gcs,proto3" json:"gcs,omitempty"`
	Azs    *Azs    `protobuf:"bytes,2,opt,name=azs,proto3" json:"azs,omitempty"`
	Oracle *Oracle `protobuf:"bytes,3,opt,name=oracle,proto3" json:"oracle,omitempty"`
	S3     *S3     `protobuf:"bytes,4,opt,name=s3,proto3" json:"s3,omitempty"`
}

func (x *PersistentStorage) Reset() {
	*x = PersistentStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_persistent_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersistentStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersistentStorage) ProtoMessage() {}

func (x *PersistentStorage) ProtoReflect() protoreflect.Message {
	mi := &file_storage_persistent_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersistentStorage.ProtoReflect.Descriptor instead.
func (*PersistentStorage) Descriptor() ([]byte, []int) {
	return file_storage_persistent_storage_proto_rawDescGZIP(), []int{0}
}

func (x *PersistentStorage) GetGcs() *Gcs {
	if x != nil {
		return x.Gcs
	}
	return nil
}

func (x *PersistentStorage) GetAzs() *Azs {
	if x != nil {
		return x.Azs
	}
	return nil
}

func (x *PersistentStorage) GetOracle() *Oracle {
	if x != nil {
		return x.Oracle
	}
	return nil
}

func (x *PersistentStorage) GetS3() *S3 {
	if x != nil {
		return x.S3
	}
	return nil
}

var File_storage_persistent_storage_proto protoreflect.FileDescriptor

var file_storage_persistent_storage_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x1a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x7a, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x67, 0x63,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb1, 0x01, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x67, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x47, 0x63, 0x73, 0x52, 0x03, 0x67, 0x63, 0x73, 0x12, 0x24, 0x0a, 0x03, 0x61,
	0x7a, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x7a, 0x73, 0x52, 0x03, 0x61, 0x7a,
	0x73, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x12, 0x21, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x33, 0x52,
	0x02, 0x73, 0x33, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_persistent_storage_proto_rawDescOnce sync.Once
	file_storage_persistent_storage_proto_rawDescData = file_storage_persistent_storage_proto_rawDesc
)

func file_storage_persistent_storage_proto_rawDescGZIP() []byte {
	file_storage_persistent_storage_proto_rawDescOnce.Do(func() {
		file_storage_persistent_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_persistent_storage_proto_rawDescData)
	})
	return file_storage_persistent_storage_proto_rawDescData
}

var file_storage_persistent_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_persistent_storage_proto_goTypes = []interface{}{
	(*PersistentStorage)(nil), // 0: proto.storage.PersistentStorage
	(*Gcs)(nil),               // 1: proto.storage.Gcs
	(*Azs)(nil),               // 2: proto.storage.Azs
	(*Oracle)(nil),            // 3: proto.storage.Oracle
	(*S3)(nil),                // 4: proto.storage.S3
}
var file_storage_persistent_storage_proto_depIdxs = []int32{
	1, // 0: proto.storage.PersistentStorage.gcs:type_name -> proto.storage.Gcs
	2, // 1: proto.storage.PersistentStorage.azs:type_name -> proto.storage.Azs
	3, // 2: proto.storage.PersistentStorage.oracle:type_name -> proto.storage.Oracle
	4, // 3: proto.storage.PersistentStorage.s3:type_name -> proto.storage.S3
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_storage_persistent_storage_proto_init() }
func file_storage_persistent_storage_proto_init() {
	if File_storage_persistent_storage_proto != nil {
		return
	}
	file_storage_azs_proto_init()
	file_storage_gcs_proto_init()
	file_storage_oracle_proto_init()
	file_storage_s3_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_storage_persistent_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersistentStorage); i {
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
			RawDescriptor: file_storage_persistent_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_persistent_storage_proto_goTypes,
		DependencyIndexes: file_storage_persistent_storage_proto_depIdxs,
		MessageInfos:      file_storage_persistent_storage_proto_msgTypes,
	}.Build()
	File_storage_persistent_storage_proto = out.File
	file_storage_persistent_storage_proto_rawDesc = nil
	file_storage_persistent_storage_proto_goTypes = nil
	file_storage_persistent_storage_proto_depIdxs = nil
}
