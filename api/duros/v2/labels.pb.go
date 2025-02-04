// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: lightbits/api/duros/v2/labels.proto

package v2

import (
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

// label-key and label-value length must be between 1 and 253 characters and
// may contain any of: alphanumeric characters (a-z, A-Z, 0-9), hyphen (-),
// underscore (_) and dot (.).
type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	mi := &file_lightbits_api_duros_v2_labels_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_lightbits_api_duros_v2_labels_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_lightbits_api_duros_v2_labels_proto_rawDescGZIP(), []int{0}
}

func (x *Label) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Label) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_lightbits_api_duros_v2_labels_proto protoreflect.FileDescriptor

var file_lightbits_api_duros_v2_labels_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x64, 0x75, 0x72, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x62, 0x69, 0x74, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x75, 0x72, 0x6f, 0x73, 0x2e, 0x76, 0x32, 0x22, 0x2f, 0x0a,
	0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0d,
	0x5a, 0x0b, 0x64, 0x75, 0x72, 0x6f, 0x73, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x32, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lightbits_api_duros_v2_labels_proto_rawDescOnce sync.Once
	file_lightbits_api_duros_v2_labels_proto_rawDescData = file_lightbits_api_duros_v2_labels_proto_rawDesc
)

func file_lightbits_api_duros_v2_labels_proto_rawDescGZIP() []byte {
	file_lightbits_api_duros_v2_labels_proto_rawDescOnce.Do(func() {
		file_lightbits_api_duros_v2_labels_proto_rawDescData = protoimpl.X.CompressGZIP(file_lightbits_api_duros_v2_labels_proto_rawDescData)
	})
	return file_lightbits_api_duros_v2_labels_proto_rawDescData
}

var file_lightbits_api_duros_v2_labels_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_lightbits_api_duros_v2_labels_proto_goTypes = []any{
	(*Label)(nil), // 0: lightbits.api.duros.v2.Label
}
var file_lightbits_api_duros_v2_labels_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_lightbits_api_duros_v2_labels_proto_init() }
func file_lightbits_api_duros_v2_labels_proto_init() {
	if File_lightbits_api_duros_v2_labels_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lightbits_api_duros_v2_labels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lightbits_api_duros_v2_labels_proto_goTypes,
		DependencyIndexes: file_lightbits_api_duros_v2_labels_proto_depIdxs,
		MessageInfos:      file_lightbits_api_duros_v2_labels_proto_msgTypes,
	}.Build()
	File_lightbits_api_duros_v2_labels_proto = out.File
	file_lightbits_api_duros_v2_labels_proto_rawDesc = nil
	file_lightbits_api_duros_v2_labels_proto_goTypes = nil
	file_lightbits_api_duros_v2_labels_proto_depIdxs = nil
}
