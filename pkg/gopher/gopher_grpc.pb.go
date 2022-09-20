// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: pkg/gopher/gopher.proto

package go_gopher_grpc

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

// GopherClient is the client API for Gopher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GopherClient interface {
	// Get Gopher URL
	GetGopher(ctx context.Context, in *GopherRequest, opts ...grpc.CallOption) (*GopherReply, error)
}

type gopherClient struct {
	cc grpc.ClientConnInterface
}

func NewGopherClient(cc grpc.ClientConnInterface) GopherClient {
	return &gopherClient{cc}
}

func (c *gopherClient) GetGopher(ctx context.Context, in *GopherRequest, opts ...grpc.CallOption) (*GopherReply, error) {
	out := new(GopherReply)
	err := c.cc.Invoke(ctx, "/gopher.Gopher/GetGopher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GopherServer is the server API for Gopher service.
// All implementations must embed UnimplementedGopherServer
// for forward compatibility
type GopherServer interface {
	// Get Gopher URL
	GetGopher(context.Context, *GopherRequest) (*GopherReply, error)
	mustEmbedUnimplementedGopherServer()
}

// UnimplementedGopherServer must be embedded to have forward compatible implementations.
type UnimplementedGopherServer struct {
}

func (UnimplementedGopherServer) GetGopher(context.Context, *GopherRequest) (*GopherReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGopher not implemented")
}
func (UnimplementedGopherServer) mustEmbedUnimplementedGopherServer() {}

// UnsafeGopherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GopherServer will
// result in compilation errors.
type UnsafeGopherServer interface {
	mustEmbedUnimplementedGopherServer()
}

func RegisterGopherServer(s grpc.ServiceRegistrar, srv GopherServer) {
	s.RegisterService(&Gopher_ServiceDesc, srv)
}

func _Gopher_GetGopher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GopherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GopherServer).GetGopher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gopher.Gopher/GetGopher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GopherServer).GetGopher(ctx, req.(*GopherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gopher_ServiceDesc is the grpc.ServiceDesc for Gopher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gopher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gopher.Gopher",
	HandlerType: (*GopherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGopher",
			Handler:    _Gopher_GetGopher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/gopher/gopher.proto",
}
