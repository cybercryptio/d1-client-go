// Code generated by copy-client.sh. DO NOT EDIT.
// version: v0.1.47-ci.180
// source: https://github.com/cybercryptio/d1-service-generic.git
// commit: 181d63285cd0a8ae6fa257615c0a1c5a60529e25

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authn

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthnClient is the client API for Authn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthnClient interface {
	//*
	// Creates a new user.
	// This call can fail if the auth storage cannot be reached, in which case an error is returned.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	//*
	// Logs in an existing user, returning a User Access Token and an expiry time.
	// This call can fail if the caller provides the wrong credentials or if the auth storage cannot be reached,
	// in which case an error is returned.
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	//*
	// Deletes an existing user.
	// This call can fail if the user does not exist,
	// or if the auth storage cannot be reached, in which case an error is returned.
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error)
	//*
	// Creates a new group with the requested scopes. The caller is added to the group.
	// This call can fail if the auth storage cannot be reached,
	// in which case an error is returned.
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error)
	//*
	// Adds a user to a group.
	// This call can fail if the auth storage cannot be reached,
	// in which case an error is returned.
	AddUserToGroup(ctx context.Context, in *AddUserToGroupRequest, opts ...grpc.CallOption) (*AddUserToGroupResponse, error)
	//*
	// Removes a user from a group.
	// This call can fail if the auth storage cannot be reached,
	// in which case an error is returned.
	RemoveUserFromGroup(ctx context.Context, in *RemoveUserFromGroupRequest, opts ...grpc.CallOption) (*RemoveUserFromGroupResponse, error)
}

type authnClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthnClient(cc grpc.ClientConnInterface) AuthnClient {
	return &authnClient{cc}
}

func (c *authnClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/d1.authn.Authn/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/d1.authn.Authn/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error) {
	out := new(RemoveUserResponse)
	err := c.cc.Invoke(ctx, "/d1.authn.Authn/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*CreateGroupResponse, error) {
	out := new(CreateGroupResponse)
	err := c.cc.Invoke(ctx, "/d1.authn.Authn/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnClient) AddUserToGroup(ctx context.Context, in *AddUserToGroupRequest, opts ...grpc.CallOption) (*AddUserToGroupResponse, error) {
	out := new(AddUserToGroupResponse)
	err := c.cc.Invoke(ctx, "/d1.authn.Authn/AddUserToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnClient) RemoveUserFromGroup(ctx context.Context, in *RemoveUserFromGroupRequest, opts ...grpc.CallOption) (*RemoveUserFromGroupResponse, error) {
	out := new(RemoveUserFromGroupResponse)
	err := c.cc.Invoke(ctx, "/d1.authn.Authn/RemoveUserFromGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthnServer is the server API for Authn service.
// All implementations must embed UnimplementedAuthnServer
// for forward compatibility
type AuthnServer interface {
	//*
	// Creates a new user.
	// This call can fail if the auth storage cannot be reached, in which case an error is returned.
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	//*
	// Logs in an existing user, returning a User Access Token and an expiry time.
	// This call can fail if the caller provides the wrong credentials or if the auth storage cannot be reached,
	// in which case an error is returned.
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	//*
	// Deletes an existing user.
	// This call can fail if the user does not exist,
	// or if the auth storage cannot be reached, in which case an error is returned.
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error)
	//*
	// Creates a new group with the requested scopes. The caller is added to the group.
	// This call can fail if the auth storage cannot be reached,
	// in which case an error is returned.
	CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
	//*
	// Adds a user to a group.
	// This call can fail if the auth storage cannot be reached,
	// in which case an error is returned.
	AddUserToGroup(context.Context, *AddUserToGroupRequest) (*AddUserToGroupResponse, error)
	//*
	// Removes a user from a group.
	// This call can fail if the auth storage cannot be reached,
	// in which case an error is returned.
	RemoveUserFromGroup(context.Context, *RemoveUserFromGroupRequest) (*RemoveUserFromGroupResponse, error)
	mustEmbedUnimplementedAuthnServer()
}

// UnimplementedAuthnServer must be embedded to have forward compatible implementations.
type UnimplementedAuthnServer struct {
}

func (UnimplementedAuthnServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedAuthnServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedAuthnServer) RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedAuthnServer) CreateGroup(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedAuthnServer) AddUserToGroup(context.Context, *AddUserToGroupRequest) (*AddUserToGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToGroup not implemented")
}
func (UnimplementedAuthnServer) RemoveUserFromGroup(context.Context, *RemoveUserFromGroupRequest) (*RemoveUserFromGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserFromGroup not implemented")
}
func (UnimplementedAuthnServer) mustEmbedUnimplementedAuthnServer() {}

// UnsafeAuthnServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthnServer will
// result in compilation errors.
type UnsafeAuthnServer interface {
	mustEmbedUnimplementedAuthnServer()
}

func RegisterAuthnServer(s grpc.ServiceRegistrar, srv AuthnServer) {
	s.RegisterService(&Authn_ServiceDesc, srv)
}

func _Authn_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/d1.authn.Authn/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authn_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/d1.authn.Authn/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authn_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/d1.authn.Authn/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authn_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/d1.authn.Authn/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authn_AddUserToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnServer).AddUserToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/d1.authn.Authn/AddUserToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnServer).AddUserToGroup(ctx, req.(*AddUserToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authn_RemoveUserFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserFromGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnServer).RemoveUserFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/d1.authn.Authn/RemoveUserFromGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnServer).RemoveUserFromGroup(ctx, req.(*RemoveUserFromGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authn_ServiceDesc is the grpc.ServiceDesc for Authn service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authn_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "d1.authn.Authn",
	HandlerType: (*AuthnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Authn_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _Authn_LoginUser_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _Authn_RemoveUser_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _Authn_CreateGroup_Handler,
		},
		{
			MethodName: "AddUserToGroup",
			Handler:    _Authn_AddUserToGroup_Handler,
		},
		{
			MethodName: "RemoveUserFromGroup",
			Handler:    _Authn_RemoveUserFromGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authn.proto",
}
