// Code generated by copy-client.sh. DO NOT EDIT.
// version: v2.0.0-ci.67
// source: https://github.com/cybercryptio/d1-service-generic.git
// commit: 75a3140090e05b8c5736b031f09927e499587d9b

// Copyright 2020-2022 CYBERCRYPT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.4
// source: authz.proto

package authz

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

// Represents a request to get the permissions of an object.
type GetPermissionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the object to get the permission list for.
	ObjectId string `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *GetPermissionsRequest) Reset() {
	*x = GetPermissionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPermissionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionsRequest) ProtoMessage() {}

func (x *GetPermissionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionsRequest.ProtoReflect.Descriptor instead.
func (*GetPermissionsRequest) Descriptor() ([]byte, []int) {
	return file_authz_proto_rawDescGZIP(), []int{0}
}

func (x *GetPermissionsRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

// Represents the result of a request to get permissions for an object.
type GetPermissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of groups with access to the object.
	GroupIds []string `protobuf:"bytes,1,rep,name=group_ids,json=groupIds,proto3" json:"group_ids,omitempty"`
}

func (x *GetPermissionsResponse) Reset() {
	*x = GetPermissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPermissionsResponse) ProtoMessage() {}

func (x *GetPermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPermissionsResponse.ProtoReflect.Descriptor instead.
func (*GetPermissionsResponse) Descriptor() ([]byte, []int) {
	return file_authz_proto_rawDescGZIP(), []int{1}
}

func (x *GetPermissionsResponse) GetGroupIds() []string {
	if x != nil {
		return x.GroupIds
	}
	return nil
}

// Represents a request to add permission to an object.
type AddPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the object to add the permission to.
	ObjectId string `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	// The IDs of the groups to give access.
	GroupIds []string `protobuf:"bytes,2,rep,name=group_ids,json=groupIds,proto3" json:"group_ids,omitempty"`
}

func (x *AddPermissionRequest) Reset() {
	*x = AddPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPermissionRequest) ProtoMessage() {}

func (x *AddPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPermissionRequest.ProtoReflect.Descriptor instead.
func (*AddPermissionRequest) Descriptor() ([]byte, []int) {
	return file_authz_proto_rawDescGZIP(), []int{2}
}

func (x *AddPermissionRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *AddPermissionRequest) GetGroupIds() []string {
	if x != nil {
		return x.GroupIds
	}
	return nil
}

// Represents the result of a request to add permission to an object.
type AddPermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddPermissionResponse) Reset() {
	*x = AddPermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPermissionResponse) ProtoMessage() {}

func (x *AddPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPermissionResponse.ProtoReflect.Descriptor instead.
func (*AddPermissionResponse) Descriptor() ([]byte, []int) {
	return file_authz_proto_rawDescGZIP(), []int{3}
}

// Represents a request to remove permission to an object.
type RemovePermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the object to remove the permission for.
	ObjectId string `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	// The IDs of the groups to revoke permission for.
	GroupIds []string `protobuf:"bytes,2,rep,name=group_ids,json=groupIds,proto3" json:"group_ids,omitempty"`
}

func (x *RemovePermissionRequest) Reset() {
	*x = RemovePermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePermissionRequest) ProtoMessage() {}

func (x *RemovePermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authz_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePermissionRequest.ProtoReflect.Descriptor instead.
func (*RemovePermissionRequest) Descriptor() ([]byte, []int) {
	return file_authz_proto_rawDescGZIP(), []int{4}
}

func (x *RemovePermissionRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *RemovePermissionRequest) GetGroupIds() []string {
	if x != nil {
		return x.GroupIds
	}
	return nil
}

// Represents the result of a request to remove permission to an object.
type RemovePermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemovePermissionResponse) Reset() {
	*x = RemovePermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemovePermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemovePermissionResponse) ProtoMessage() {}

func (x *RemovePermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authz_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemovePermissionResponse.ProtoReflect.Descriptor instead.
func (*RemovePermissionResponse) Descriptor() ([]byte, []int) {
	return file_authz_proto_rawDescGZIP(), []int{5}
}

// Represents a request to check whether the user has permission to an object.
type CheckPermissionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the object to check permission for.
	ObjectId string `protobuf:"bytes,1,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *CheckPermissionRequest) Reset() {
	*x = CheckPermissionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionRequest) ProtoMessage() {}

func (x *CheckPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authz_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionRequest.ProtoReflect.Descriptor instead.
func (*CheckPermissionRequest) Descriptor() ([]byte, []int) {
	return file_authz_proto_rawDescGZIP(), []int{6}
}

func (x *CheckPermissionRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

// Represents the result of a request to check whether a user has permission to an object.
type CheckPermissionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates whether the caller has access to the object.
	HasPermission bool `protobuf:"varint,1,opt,name=has_permission,json=hasPermission,proto3" json:"has_permission,omitempty"`
}

func (x *CheckPermissionResponse) Reset() {
	*x = CheckPermissionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionResponse) ProtoMessage() {}

func (x *CheckPermissionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authz_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionResponse.ProtoReflect.Descriptor instead.
func (*CheckPermissionResponse) Descriptor() ([]byte, []int) {
	return file_authz_proto_rawDescGZIP(), []int{7}
}

func (x *CheckPermissionResponse) GetHasPermission() bool {
	if x != nil {
		return x.HasPermission
	}
	return false
}

var File_authz_proto protoreflect.FileDescriptor

var file_authz_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64,
	0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x22, 0x34, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x35, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x73, 0x22, 0x50, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x53, 0x0a, 0x17, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x35, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x68, 0x61, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x68, 0x61, 0x73, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xe9, 0x02, 0x0a, 0x05, 0x41, 0x75,
	0x74, 0x68, 0x7a, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0d, 0x41, 0x64,
	0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x64, 0x31,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x31,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b,
	0x0a, 0x10, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0f, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20,
	0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x64, 0x31, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x91, 0x01, 0x0a, 0x1f, 0x69, 0x6f, 0x2e, 0x63, 0x79, 0x62,
	0x65, 0x72, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x64, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x42, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x7a,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x79, 0x62, 0x65, 0x72, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x2f, 0x64,
	0x31, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69,
	0x63, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x7a, 0xaa, 0x02, 0x23, 0x43, 0x79, 0x62, 0x65, 0x72, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x2e, 0x44, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_authz_proto_rawDescOnce sync.Once
	file_authz_proto_rawDescData = file_authz_proto_rawDesc
)

func file_authz_proto_rawDescGZIP() []byte {
	file_authz_proto_rawDescOnce.Do(func() {
		file_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_authz_proto_rawDescData)
	})
	return file_authz_proto_rawDescData
}

var file_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_authz_proto_goTypes = []interface{}{
	(*GetPermissionsRequest)(nil),    // 0: d1.authz.GetPermissionsRequest
	(*GetPermissionsResponse)(nil),   // 1: d1.authz.GetPermissionsResponse
	(*AddPermissionRequest)(nil),     // 2: d1.authz.AddPermissionRequest
	(*AddPermissionResponse)(nil),    // 3: d1.authz.AddPermissionResponse
	(*RemovePermissionRequest)(nil),  // 4: d1.authz.RemovePermissionRequest
	(*RemovePermissionResponse)(nil), // 5: d1.authz.RemovePermissionResponse
	(*CheckPermissionRequest)(nil),   // 6: d1.authz.CheckPermissionRequest
	(*CheckPermissionResponse)(nil),  // 7: d1.authz.CheckPermissionResponse
}
var file_authz_proto_depIdxs = []int32{
	0, // 0: d1.authz.Authz.GetPermissions:input_type -> d1.authz.GetPermissionsRequest
	2, // 1: d1.authz.Authz.AddPermission:input_type -> d1.authz.AddPermissionRequest
	4, // 2: d1.authz.Authz.RemovePermission:input_type -> d1.authz.RemovePermissionRequest
	6, // 3: d1.authz.Authz.CheckPermission:input_type -> d1.authz.CheckPermissionRequest
	1, // 4: d1.authz.Authz.GetPermissions:output_type -> d1.authz.GetPermissionsResponse
	3, // 5: d1.authz.Authz.AddPermission:output_type -> d1.authz.AddPermissionResponse
	5, // 6: d1.authz.Authz.RemovePermission:output_type -> d1.authz.RemovePermissionResponse
	7, // 7: d1.authz.Authz.CheckPermission:output_type -> d1.authz.CheckPermissionResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_authz_proto_init() }
func file_authz_proto_init() {
	if File_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPermissionsRequest); i {
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
		file_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPermissionsResponse); i {
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
		file_authz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPermissionRequest); i {
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
		file_authz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPermissionResponse); i {
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
		file_authz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePermissionRequest); i {
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
		file_authz_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemovePermissionResponse); i {
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
		file_authz_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionRequest); i {
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
		file_authz_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionResponse); i {
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
			RawDescriptor: file_authz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authz_proto_goTypes,
		DependencyIndexes: file_authz_proto_depIdxs,
		MessageInfos:      file_authz_proto_msgTypes,
	}.Build()
	File_authz_proto = out.File
	file_authz_proto_rawDesc = nil
	file_authz_proto_goTypes = nil
	file_authz_proto_depIdxs = nil
}
