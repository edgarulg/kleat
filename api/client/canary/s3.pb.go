// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: canary/s3.proto

package canary

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

// Configuration for the S3 canary integration. If enabled, you must
// also configure at least one canary.AwsAccount with a list of
// supportedTypes that includes canary.SupportedType.CONFIGURATION_STORE and/or
// canary.SupportedType.OBJECT_STORE.
type S3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the integration is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *S3) Reset() {
	*x = S3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_canary_s3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3) ProtoMessage() {}

func (x *S3) ProtoReflect() protoreflect.Message {
	mi := &file_canary_s3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3.ProtoReflect.Descriptor instead.
func (*S3) Descriptor() ([]byte, []int) {
	return file_canary_s3_proto_rawDescGZIP(), []int{0}
}

func (x *S3) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

var File_canary_s3_proto protoreflect.FileDescriptor

var file_canary_s3_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x22,
	0x1e, 0x0a, 0x02, 0x53, 0x33, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70,
	0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_canary_s3_proto_rawDescOnce sync.Once
	file_canary_s3_proto_rawDescData = file_canary_s3_proto_rawDesc
)

func file_canary_s3_proto_rawDescGZIP() []byte {
	file_canary_s3_proto_rawDescOnce.Do(func() {
		file_canary_s3_proto_rawDescData = protoimpl.X.CompressGZIP(file_canary_s3_proto_rawDescData)
	})
	return file_canary_s3_proto_rawDescData
}

var file_canary_s3_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_canary_s3_proto_goTypes = []interface{}{
	(*S3)(nil), // 0: proto.canary.S3
}
var file_canary_s3_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_canary_s3_proto_init() }
func file_canary_s3_proto_init() {
	if File_canary_s3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_canary_s3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3); i {
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
			RawDescriptor: file_canary_s3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_canary_s3_proto_goTypes,
		DependencyIndexes: file_canary_s3_proto_depIdxs,
		MessageInfos:      file_canary_s3_proto_msgTypes,
	}.Build()
	File_canary_s3_proto = out.File
	file_canary_s3_proto_rawDesc = nil
	file_canary_s3_proto_goTypes = nil
	file_canary_s3_proto_depIdxs = nil
}