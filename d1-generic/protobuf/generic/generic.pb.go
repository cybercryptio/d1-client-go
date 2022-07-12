// Code generated by copy-client.sh. DO NOT EDIT.
// version: v0.1.48
// source: https://github.com/cybercryptio/d1-service-generic.git
// commit: 9c3a7885a1231e87272eec459b08a99e00abfe5c

// Copyright 2020-2022 CYBERCRYPT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: generic.proto

package generic

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

/// Represents a request to encrypt data.
type EncryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// Data to encrypt.
	Plaintext []byte `protobuf:"bytes,1,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
	/// Associated data.
	AssociatedData []byte `protobuf:"bytes,2,opt,name=associated_data,json=associatedData,proto3" json:"associated_data,omitempty"`
}

func (x *EncryptRequest) Reset() {
	*x = EncryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptRequest) ProtoMessage() {}

func (x *EncryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_generic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptRequest.ProtoReflect.Descriptor instead.
func (*EncryptRequest) Descriptor() ([]byte, []int) {
	return file_generic_proto_rawDescGZIP(), []int{0}
}

func (x *EncryptRequest) GetPlaintext() []byte {
	if x != nil {
		return x.Plaintext
	}
	return nil
}

func (x *EncryptRequest) GetAssociatedData() []byte {
	if x != nil {
		return x.AssociatedData
	}
	return nil
}

/// Represents a response to an encryption request.
type EncryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// Ciphertext of the provided plaintext.
	Ciphertext []byte `protobuf:"bytes,1,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	/// Associated data.
	AssociatedData []byte `protobuf:"bytes,2,opt,name=associated_data,json=associatedData,proto3" json:"associated_data,omitempty"`
	/// The object ID of the encrypted data.
	ObjectId string `protobuf:"bytes,3,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *EncryptResponse) Reset() {
	*x = EncryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptResponse) ProtoMessage() {}

func (x *EncryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_generic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptResponse.ProtoReflect.Descriptor instead.
func (*EncryptResponse) Descriptor() ([]byte, []int) {
	return file_generic_proto_rawDescGZIP(), []int{1}
}

func (x *EncryptResponse) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

func (x *EncryptResponse) GetAssociatedData() []byte {
	if x != nil {
		return x.AssociatedData
	}
	return nil
}

func (x *EncryptResponse) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

/// Represents a request to decrypt data.
type DecryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// Data to decrypt.
	Ciphertext []byte `protobuf:"bytes,1,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	/// Associated data.
	AssociatedData []byte `protobuf:"bytes,2,opt,name=associated_data,json=associatedData,proto3" json:"associated_data,omitempty"`
	/// The object ID of the data.
	ObjectId string `protobuf:"bytes,3,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *DecryptRequest) Reset() {
	*x = DecryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptRequest) ProtoMessage() {}

func (x *DecryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_generic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptRequest.ProtoReflect.Descriptor instead.
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return file_generic_proto_rawDescGZIP(), []int{2}
}

func (x *DecryptRequest) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

func (x *DecryptRequest) GetAssociatedData() []byte {
	if x != nil {
		return x.AssociatedData
	}
	return nil
}

func (x *DecryptRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

/// Represents a response to a decryption request.
type DecryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	/// Decrypted data.
	Plaintext []byte `protobuf:"bytes,1,opt,name=plaintext,proto3" json:"plaintext,omitempty"`
	/// Associated data.
	AssociatedData []byte `protobuf:"bytes,2,opt,name=associated_data,json=associatedData,proto3" json:"associated_data,omitempty"`
}

func (x *DecryptResponse) Reset() {
	*x = DecryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptResponse) ProtoMessage() {}

func (x *DecryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_generic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptResponse.ProtoReflect.Descriptor instead.
func (*DecryptResponse) Descriptor() ([]byte, []int) {
	return file_generic_proto_rawDescGZIP(), []int{3}
}

func (x *DecryptResponse) GetPlaintext() []byte {
	if x != nil {
		return x.Plaintext
	}
	return nil
}

func (x *DecryptResponse) GetAssociatedData() []byte {
	if x != nil {
		return x.AssociatedData
	}
	return nil
}

var File_generic_proto protoreflect.FileDescriptor

var file_generic_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x64, 0x31, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x22, 0x57, 0x0a, 0x0e, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61,
	0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x77, 0x0a, 0x0f, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65,
	0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x73, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0e, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x76, 0x0a,
	0x0e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x27, 0x0a, 0x0f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x0f, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x69,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0e, 0x61, 0x73, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x32,
	0x95, 0x01, 0x0a, 0x07, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x12, 0x44, 0x0a, 0x07, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x1a, 0x2e, 0x64, 0x31, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x31, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e,
	0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x1a, 0x2e, 0x64,
	0x31, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x31, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x96, 0x01, 0x0a, 0x21, 0x69, 0x6f, 0x2e, 0x63,
	0x79, 0x62, 0x65, 0x72, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x64, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x42, 0x0c, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x79, 0x62, 0x65, 0x72, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x2f, 0x64, 0x31, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0xaa, 0x02, 0x25, 0x43, 0x79, 0x62, 0x65, 0x72,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x44, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_generic_proto_rawDescOnce sync.Once
	file_generic_proto_rawDescData = file_generic_proto_rawDesc
)

func file_generic_proto_rawDescGZIP() []byte {
	file_generic_proto_rawDescOnce.Do(func() {
		file_generic_proto_rawDescData = protoimpl.X.CompressGZIP(file_generic_proto_rawDescData)
	})
	return file_generic_proto_rawDescData
}

var file_generic_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_generic_proto_goTypes = []interface{}{
	(*EncryptRequest)(nil),  // 0: d1.generic.EncryptRequest
	(*EncryptResponse)(nil), // 1: d1.generic.EncryptResponse
	(*DecryptRequest)(nil),  // 2: d1.generic.DecryptRequest
	(*DecryptResponse)(nil), // 3: d1.generic.DecryptResponse
}
var file_generic_proto_depIdxs = []int32{
	0, // 0: d1.generic.Generic.Encrypt:input_type -> d1.generic.EncryptRequest
	2, // 1: d1.generic.Generic.Decrypt:input_type -> d1.generic.DecryptRequest
	1, // 2: d1.generic.Generic.Encrypt:output_type -> d1.generic.EncryptResponse
	3, // 3: d1.generic.Generic.Decrypt:output_type -> d1.generic.DecryptResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_generic_proto_init() }
func file_generic_proto_init() {
	if File_generic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_generic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptRequest); i {
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
		file_generic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptResponse); i {
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
		file_generic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptRequest); i {
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
		file_generic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecryptResponse); i {
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
			RawDescriptor: file_generic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_generic_proto_goTypes,
		DependencyIndexes: file_generic_proto_depIdxs,
		MessageInfos:      file_generic_proto_msgTypes,
	}.Build()
	File_generic_proto = out.File
	file_generic_proto_rawDesc = nil
	file_generic_proto_goTypes = nil
	file_generic_proto_depIdxs = nil
}
