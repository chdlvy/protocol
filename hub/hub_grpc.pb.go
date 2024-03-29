// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.0--rc3
// source: hub/hub.proto

package hub

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

// HubClient is the client API for Hub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HubClient interface {
	OnlinePushMsg(ctx context.Context, in *OnlinePushMsgReq, opts ...grpc.CallOption) (*OnlinePushMsgResp, error)
}

type hubClient struct {
	cc grpc.ClientConnInterface
}

func NewHubClient(cc grpc.ClientConnInterface) HubClient {
	return &hubClient{cc}
}

func (c *hubClient) OnlinePushMsg(ctx context.Context, in *OnlinePushMsgReq, opts ...grpc.CallOption) (*OnlinePushMsgResp, error) {
	out := new(OnlinePushMsgResp)
	err := c.cc.Invoke(ctx, "/ChatIM.hub.hub/OnlinePushMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HubServer is the server API for Hub service.
// All implementations must embed UnimplementedHubServer
// for forward compatibility
type HubServer interface {
	OnlinePushMsg(context.Context, *OnlinePushMsgReq) (*OnlinePushMsgResp, error)
	mustEmbedUnimplementedHubServer()
}

// UnimplementedHubServer must be embedded to have forward compatible implementations.
type UnimplementedHubServer struct {
}

func (UnimplementedHubServer) OnlinePushMsg(context.Context, *OnlinePushMsgReq) (*OnlinePushMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlinePushMsg not implemented")
}
func (UnimplementedHubServer) mustEmbedUnimplementedHubServer() {}

// UnsafeHubServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HubServer will
// result in compilation errors.
type UnsafeHubServer interface {
	mustEmbedUnimplementedHubServer()
}

func RegisterHubServer(s grpc.ServiceRegistrar, srv HubServer) {
	s.RegisterService(&Hub_ServiceDesc, srv)
}

func _Hub_OnlinePushMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlinePushMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubServer).OnlinePushMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatIM.hub.hub/OnlinePushMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubServer).OnlinePushMsg(ctx, req.(*OnlinePushMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Hub_ServiceDesc is the grpc.ServiceDesc for Hub service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hub_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChatIM.hub.hub",
	HandlerType: (*HubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnlinePushMsg",
			Handler:    _Hub_OnlinePushMsg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hub/hub.proto",
}
