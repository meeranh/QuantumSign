// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: secure_communication.proto

package kyber_auth

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

const (
	KeyExchangeService_KeyExchange_FullMethodName = "/securecommunication.KeyExchangeService/KeyExchange"
)

// KeyExchangeServiceClient is the client API for KeyExchangeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyExchangeServiceClient interface {
	KeyExchange(ctx context.Context, in *KeyExchangeRequest, opts ...grpc.CallOption) (*KeyExchangeResponse, error)
}

type keyExchangeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyExchangeServiceClient(cc grpc.ClientConnInterface) KeyExchangeServiceClient {
	return &keyExchangeServiceClient{cc}
}

func (c *keyExchangeServiceClient) KeyExchange(ctx context.Context, in *KeyExchangeRequest, opts ...grpc.CallOption) (*KeyExchangeResponse, error) {
	out := new(KeyExchangeResponse)
	err := c.cc.Invoke(ctx, KeyExchangeService_KeyExchange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyExchangeServiceServer is the server API for KeyExchangeService service.
// All implementations must embed UnimplementedKeyExchangeServiceServer
// for forward compatibility
type KeyExchangeServiceServer interface {
	KeyExchange(context.Context, *KeyExchangeRequest) (*KeyExchangeResponse, error)
	mustEmbedUnimplementedKeyExchangeServiceServer()
}

// UnimplementedKeyExchangeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKeyExchangeServiceServer struct {
}

func (UnimplementedKeyExchangeServiceServer) KeyExchange(context.Context, *KeyExchangeRequest) (*KeyExchangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeyExchange not implemented")
}
func (UnimplementedKeyExchangeServiceServer) mustEmbedUnimplementedKeyExchangeServiceServer() {}

// UnsafeKeyExchangeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyExchangeServiceServer will
// result in compilation errors.
type UnsafeKeyExchangeServiceServer interface {
	mustEmbedUnimplementedKeyExchangeServiceServer()
}

func RegisterKeyExchangeServiceServer(s grpc.ServiceRegistrar, srv KeyExchangeServiceServer) {
	s.RegisterService(&KeyExchangeService_ServiceDesc, srv)
}

func _KeyExchangeService_KeyExchange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyExchangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyExchangeServiceServer).KeyExchange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KeyExchangeService_KeyExchange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyExchangeServiceServer).KeyExchange(ctx, req.(*KeyExchangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyExchangeService_ServiceDesc is the grpc.ServiceDesc for KeyExchangeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyExchangeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "securecommunication.KeyExchangeService",
	HandlerType: (*KeyExchangeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "KeyExchange",
			Handler:    _KeyExchangeService_KeyExchange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secure_communication.proto",
}