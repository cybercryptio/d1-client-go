// Code generated by copy-client.sh. DO NOT EDIT.
// version: v1.0.1
// source: https://github.com/cybercryptio/k1.git
// commit: 37945eed2d43346c0bba218aa27d0f2673df015b

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protobuf

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

// KeyAPIClient is the client API for KeyAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyAPIClient interface {
	// GetKeySet returns a nonce and wrapped key set per given Key Initialization Key ID.
	GetKeySet(ctx context.Context, in *GetKeySetRequest, opts ...grpc.CallOption) (*GetKeySetResponse, error)
}

type keyAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyAPIClient(cc grpc.ClientConnInterface) KeyAPIClient {
	return &keyAPIClient{cc}
}

func (c *keyAPIClient) GetKeySet(ctx context.Context, in *GetKeySetRequest, opts ...grpc.CallOption) (*GetKeySetResponse, error) {
	out := new(GetKeySetResponse)
	err := c.cc.Invoke(ctx, "/k1.KeyAPI/GetKeySet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyAPIServer is the server API for KeyAPI service.
// All implementations must embed UnimplementedKeyAPIServer
// for forward compatibility
type KeyAPIServer interface {
	// GetKeySet returns a nonce and wrapped key set per given Key Initialization Key ID.
	GetKeySet(context.Context, *GetKeySetRequest) (*GetKeySetResponse, error)
	mustEmbedUnimplementedKeyAPIServer()
}

// UnimplementedKeyAPIServer must be embedded to have forward compatible implementations.
type UnimplementedKeyAPIServer struct {
}

func (UnimplementedKeyAPIServer) GetKeySet(context.Context, *GetKeySetRequest) (*GetKeySetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKeySet not implemented")
}
func (UnimplementedKeyAPIServer) mustEmbedUnimplementedKeyAPIServer() {}

// UnsafeKeyAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyAPIServer will
// result in compilation errors.
type UnsafeKeyAPIServer interface {
	mustEmbedUnimplementedKeyAPIServer()
}

func RegisterKeyAPIServer(s grpc.ServiceRegistrar, srv KeyAPIServer) {
	s.RegisterService(&KeyAPI_ServiceDesc, srv)
}

func _KeyAPI_GetKeySet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKeySetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyAPIServer).GetKeySet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/k1.KeyAPI/GetKeySet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyAPIServer).GetKeySet(ctx, req.(*GetKeySetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyAPI_ServiceDesc is the grpc.ServiceDesc for KeyAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "k1.KeyAPI",
	HandlerType: (*KeyAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKeySet",
			Handler:    _KeyAPI_GetKeySet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "key.proto",
}
