// Copyright (c) 2023 IoTeX
// This source code is provided 'as is' and no warranties are given as to title or non-infringement, merchantability
// or fitness for purpose and, to the extent permitted by law, all liability for your use of the code is disclaimed.
// This source code is governed by Apache License 2.0 that can be found in the LICENSE file.

// To compile the proto, run:
//      protoc --go_out=. *.proto
// option go_package = "../indexpb"; is to specify the package name in the generated go file

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: version.proto

package versionpb

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

type VersionedNamespace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyLen uint32 `protobuf:"varint,2,opt,name=keyLen,proto3" json:"keyLen,omitempty"`
}

func (x *VersionedNamespace) Reset() {
	*x = VersionedNamespace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionedNamespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionedNamespace) ProtoMessage() {}

func (x *VersionedNamespace) ProtoReflect() protoreflect.Message {
	mi := &file_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionedNamespace.ProtoReflect.Descriptor instead.
func (*VersionedNamespace) Descriptor() ([]byte, []int) {
	return file_version_proto_rawDescGZIP(), []int{0}
}

func (x *VersionedNamespace) GetKeyLen() uint32 {
	if x != nil {
		return x.KeyLen
	}
	return 0
}

var File_version_proto protoreflect.FileDescriptor

var file_version_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x22, 0x2c, 0x0a, 0x12, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6b, 0x65, 0x79, 0x4c, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6b, 0x65, 0x79, 0x4c, 0x65, 0x6e, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x69, 0x6f, 0x74, 0x65, 0x78, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64,
	0x62, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_version_proto_rawDescOnce sync.Once
	file_version_proto_rawDescData = file_version_proto_rawDesc
)

func file_version_proto_rawDescGZIP() []byte {
	file_version_proto_rawDescOnce.Do(func() {
		file_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_version_proto_rawDescData)
	})
	return file_version_proto_rawDescData
}

var file_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_version_proto_goTypes = []any{
	(*VersionedNamespace)(nil), // 0: versionpb.VersionedNamespace
}
var file_version_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_version_proto_init() }
func file_version_proto_init() {
	if File_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_version_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*VersionedNamespace); i {
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
			RawDescriptor: file_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_version_proto_goTypes,
		DependencyIndexes: file_version_proto_depIdxs,
		MessageInfos:      file_version_proto_msgTypes,
	}.Build()
	File_version_proto = out.File
	file_version_proto_rawDesc = nil
	file_version_proto_goTypes = nil
	file_version_proto_depIdxs = nil
}
