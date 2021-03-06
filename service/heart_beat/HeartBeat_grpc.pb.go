// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package heart_beat

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

// HeartBeatClient is the client API for HeartBeat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeartBeatClient interface {
	HeartBeat(ctx context.Context, in *HeartBeatMes_Req, opts ...grpc.CallOption) (*HeartBeatMes_Response, error)
}

type heartBeatClient struct {
	cc grpc.ClientConnInterface
}

func NewHeartBeatClient(cc grpc.ClientConnInterface) HeartBeatClient {
	return &heartBeatClient{cc}
}

func (c *heartBeatClient) HeartBeat(ctx context.Context, in *HeartBeatMes_Req, opts ...grpc.CallOption) (*HeartBeatMes_Response, error) {
	out := new(HeartBeatMes_Response)
	err := c.cc.Invoke(ctx, "/CompetitionPlatform.proto.api.HeartBeat/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartBeatServer is the server API for HeartBeat service.
// All implementations must embed UnimplementedHeartBeatServer
// for forward compatibility
type HeartBeatServer interface {
	HeartBeat(context.Context, *HeartBeatMes_Req) (*HeartBeatMes_Response, error)
	mustEmbedUnimplementedHeartBeatServer()
}

// UnimplementedHeartBeatServer must be embedded to have forward compatible implementations.
type UnimplementedHeartBeatServer struct {
}

func (UnimplementedHeartBeatServer) HeartBeat(context.Context, *HeartBeatMes_Req) (*HeartBeatMes_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (UnimplementedHeartBeatServer) mustEmbedUnimplementedHeartBeatServer() {}

// UnsafeHeartBeatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeartBeatServer will
// result in compilation errors.
type UnsafeHeartBeatServer interface {
	mustEmbedUnimplementedHeartBeatServer()
}

func RegisterHeartBeatServer(s grpc.ServiceRegistrar, srv HeartBeatServer) {
	s.RegisterService(&HeartBeat_ServiceDesc, srv)
}

func _HeartBeat_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatMes_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartBeatServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CompetitionPlatform.proto.api.HeartBeat/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartBeatServer).HeartBeat(ctx, req.(*HeartBeatMes_Req))
	}
	return interceptor(ctx, in, info, handler)
}

// HeartBeat_ServiceDesc is the grpc.ServiceDesc for HeartBeat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeartBeat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CompetitionPlatform.proto.api.HeartBeat",
	HandlerType: (*HeartBeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HeartBeat",
			Handler:    _HeartBeat_HeartBeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api/HeartBeat.proto",
}
