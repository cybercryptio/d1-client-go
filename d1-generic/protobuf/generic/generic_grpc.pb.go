// Code generated by copy-client.sh. DO NOT EDIT.
// version: v0.1.49-ci.10
// source: https://github.com/cybercryptio/d1-service-generic.git
// commit: 75d46a7b02caa465ac3c608fa2316c77869196e8

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package generic

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

// GenericClient is the client API for Generic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenericClient interface {
	// Encrypts data and returns the ciphertext without storing it.
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
	// Authorizes the user for access permissions and if accessible, returns the decrypted content.
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
}

type genericClient struct {
	cc grpc.ClientConnInterface
}

func NewGenericClient(cc grpc.ClientConnInterface) GenericClient {
	return &genericClient{cc}
}

func (c *genericClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	out := new(EncryptResponse)
	err := c.cc.Invoke(ctx, "/d1.generic.Generic/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *genericClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, "/d1.generic.Generic/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenericServer is the server API for Generic service.
// All implementations must embed UnimplementedGenericServer
// for forward compatibility
type GenericServer interface {
	// Encrypts data and returns the ciphertext without storing it.
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
	// Authorizes the user for access permissions and if accessible, returns the decrypted content.
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
	mustEmbedUnimplementedGenericServer()
}

// UnimplementedGenericServer must be embedded to have forward compatible implementations.
type UnimplementedGenericServer struct {
}

func (UnimplementedGenericServer) Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (UnimplementedGenericServer) Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (UnimplementedGenericServer) mustEmbedUnimplementedGenericServer() {}

// UnsafeGenericServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenericServer will
// result in compilation errors.
type UnsafeGenericServer interface {
	mustEmbedUnimplementedGenericServer()
}

func RegisterGenericServer(s grpc.ServiceRegistrar, srv GenericServer) {
	s.RegisterService(&Generic_ServiceDesc, srv)
}

func _Generic_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/d1.generic.Generic/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Generic_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/d1.generic.Generic/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Generic_ServiceDesc is the grpc.ServiceDesc for Generic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Generic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "d1.generic.Generic",
	HandlerType: (*GenericServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encrypt",
			Handler:    _Generic_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _Generic_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generic.proto",
}
