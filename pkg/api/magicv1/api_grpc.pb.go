// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: iamnande/cardmod/magic/v1/api.proto

package magicv1

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

// MagicAPIClient is the client API for MagicAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MagicAPIClient interface {
	// ListMagics lists all available magic resources.
	ListMagics(ctx context.Context, in *ListMagicsRequest, opts ...grpc.CallOption) (*ListMagicsResponse, error)
	// DescribeMagic describes a single magic.
	DescribeMagic(ctx context.Context, in *DescribeMagicRequest, opts ...grpc.CallOption) (*DescribeMagicResponse, error)
}

type magicAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewMagicAPIClient(cc grpc.ClientConnInterface) MagicAPIClient {
	return &magicAPIClient{cc}
}

func (c *magicAPIClient) ListMagics(ctx context.Context, in *ListMagicsRequest, opts ...grpc.CallOption) (*ListMagicsResponse, error) {
	out := new(ListMagicsResponse)
	err := c.cc.Invoke(ctx, "/iamnande.cardmod.magic.v1.MagicAPI/ListMagics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *magicAPIClient) DescribeMagic(ctx context.Context, in *DescribeMagicRequest, opts ...grpc.CallOption) (*DescribeMagicResponse, error) {
	out := new(DescribeMagicResponse)
	err := c.cc.Invoke(ctx, "/iamnande.cardmod.magic.v1.MagicAPI/DescribeMagic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MagicAPIServer is the server API for MagicAPI service.
// All implementations must embed UnimplementedMagicAPIServer
// for forward compatibility
type MagicAPIServer interface {
	// ListMagics lists all available magic resources.
	ListMagics(context.Context, *ListMagicsRequest) (*ListMagicsResponse, error)
	// DescribeMagic describes a single magic.
	DescribeMagic(context.Context, *DescribeMagicRequest) (*DescribeMagicResponse, error)
	mustEmbedUnimplementedMagicAPIServer()
}

// UnimplementedMagicAPIServer must be embedded to have forward compatible implementations.
type UnimplementedMagicAPIServer struct {
}

func (UnimplementedMagicAPIServer) ListMagics(context.Context, *ListMagicsRequest) (*ListMagicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMagics not implemented")
}
func (UnimplementedMagicAPIServer) DescribeMagic(context.Context, *DescribeMagicRequest) (*DescribeMagicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeMagic not implemented")
}
func (UnimplementedMagicAPIServer) mustEmbedUnimplementedMagicAPIServer() {}

// UnsafeMagicAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MagicAPIServer will
// result in compilation errors.
type UnsafeMagicAPIServer interface {
	mustEmbedUnimplementedMagicAPIServer()
}

func RegisterMagicAPIServer(s grpc.ServiceRegistrar, srv MagicAPIServer) {
	s.RegisterService(&MagicAPI_ServiceDesc, srv)
}

func _MagicAPI_ListMagics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMagicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagicAPIServer).ListMagics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iamnande.cardmod.magic.v1.MagicAPI/ListMagics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagicAPIServer).ListMagics(ctx, req.(*ListMagicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MagicAPI_DescribeMagic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeMagicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MagicAPIServer).DescribeMagic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iamnande.cardmod.magic.v1.MagicAPI/DescribeMagic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MagicAPIServer).DescribeMagic(ctx, req.(*DescribeMagicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MagicAPI_ServiceDesc is the grpc.ServiceDesc for MagicAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MagicAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iamnande.cardmod.magic.v1.MagicAPI",
	HandlerType: (*MagicAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMagics",
			Handler:    _MagicAPI_ListMagics_Handler,
		},
		{
			MethodName: "DescribeMagic",
			Handler:    _MagicAPI_DescribeMagic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iamnande/cardmod/magic/v1/api.proto",
}
