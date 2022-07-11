// Code generated by copy-client.sh. DO NOT EDIT.
// version: v0.1.47-ci.180
// source: https://github.com/cybercryptio/d1-service-generic.git
// commit: 181d63285cd0a8ae6fa257615c0a1c5a60529e25

// Copyright 2020-2022 CYBERCRYPT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: scopes.proto

package scopes

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

/// Access scopes.
type Scope int32

const (
	/// Read and decrypt data.
	Scope_READ Scope = 0
	/// Store and encrypt data.
	Scope_CREATE Scope = 1
	/// Get permissions to an object.
	Scope_GETACCESS Scope = 2
	/// Modify permissions to an object.
	Scope_MODIFYACCESS Scope = 3
	/// Modify data.
	Scope_UPDATE Scope = 4
	/// Delete data.
	Scope_DELETE Scope = 5
	/// Use secure index for searching in data.
	Scope_INDEX Scope = 6
)

// Enum value maps for Scope.
var (
	Scope_name = map[int32]string{
		0: "READ",
		1: "CREATE",
		2: "GETACCESS",
		3: "MODIFYACCESS",
		4: "UPDATE",
		5: "DELETE",
		6: "INDEX",
	}
	Scope_value = map[string]int32{
		"READ":         0,
		"CREATE":       1,
		"GETACCESS":    2,
		"MODIFYACCESS": 3,
		"UPDATE":       4,
		"DELETE":       5,
		"INDEX":        6,
	}
)

func (x Scope) Enum() *Scope {
	p := new(Scope)
	*p = x
	return p
}

func (x Scope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Scope) Descriptor() protoreflect.EnumDescriptor {
	return file_scopes_proto_enumTypes[0].Descriptor()
}

func (Scope) Type() protoreflect.EnumType {
	return &file_scopes_proto_enumTypes[0]
}

func (x Scope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Scope.Descriptor instead.
func (Scope) EnumDescriptor() ([]byte, []int) {
	return file_scopes_proto_rawDescGZIP(), []int{0}
}

var File_scopes_proto protoreflect.FileDescriptor

var file_scopes_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x64, 0x31, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2a, 0x61, 0x0a, 0x05, 0x53, 0x63, 0x6f,
	0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x45, 0x54, 0x41,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x4d, 0x4f, 0x44, 0x49, 0x46,
	0x59, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44,
	0x41, 0x54, 0x45, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10,
	0x05, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x10, 0x06, 0x42, 0x92, 0x01, 0x0a,
	0x20, 0x69, 0x6f, 0x2e, 0x63, 0x79, 0x62, 0x65, 0x72, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x64,
	0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x73, 0x42, 0x0b, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x79, 0x62, 0x65, 0x72,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x2f, 0x64, 0x31, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0xaa, 0x02, 0x24, 0x43, 0x79, 0x62,
	0x65, 0x72, 0x43, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x44, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scopes_proto_rawDescOnce sync.Once
	file_scopes_proto_rawDescData = file_scopes_proto_rawDesc
)

func file_scopes_proto_rawDescGZIP() []byte {
	file_scopes_proto_rawDescOnce.Do(func() {
		file_scopes_proto_rawDescData = protoimpl.X.CompressGZIP(file_scopes_proto_rawDescData)
	})
	return file_scopes_proto_rawDescData
}

var file_scopes_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_scopes_proto_goTypes = []interface{}{
	(Scope)(0), // 0: d1.scopes.Scope
}
var file_scopes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_scopes_proto_init() }
func file_scopes_proto_init() {
	if File_scopes_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scopes_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scopes_proto_goTypes,
		DependencyIndexes: file_scopes_proto_depIdxs,
		EnumInfos:         file_scopes_proto_enumTypes,
	}.Build()
	File_scopes_proto = out.File
	file_scopes_proto_rawDesc = nil
	file_scopes_proto_goTypes = nil
	file_scopes_proto_depIdxs = nil
}
