// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: routing.proto

package routingpb

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

// RoutingServiceClient is the client API for RoutingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoutingServiceClient interface {
	GetAvailability(ctx context.Context, opts ...grpc.CallOption) (RoutingService_GetAvailabilityClient, error)
}

type routingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoutingServiceClient(cc grpc.ClientConnInterface) RoutingServiceClient {
	return &routingServiceClient{cc}
}

func (c *routingServiceClient) GetAvailability(ctx context.Context, opts ...grpc.CallOption) (RoutingService_GetAvailabilityClient, error) {
	stream, err := c.cc.NewStream(ctx, &RoutingService_ServiceDesc.Streams[0], "/routingpb.RoutingService/GetAvailability", opts...)
	if err != nil {
		return nil, err
	}
	x := &routingServiceGetAvailabilityClient{stream}
	return x, nil
}

type RoutingService_GetAvailabilityClient interface {
	Send(*RoutingAvailabilityRequest) error
	Recv() (*RoutingAvailabilityResponse, error)
	grpc.ClientStream
}

type routingServiceGetAvailabilityClient struct {
	grpc.ClientStream
}

func (x *routingServiceGetAvailabilityClient) Send(m *RoutingAvailabilityRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routingServiceGetAvailabilityClient) Recv() (*RoutingAvailabilityResponse, error) {
	m := new(RoutingAvailabilityResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RoutingServiceServer is the server API for RoutingService service.
// All implementations must embed UnimplementedRoutingServiceServer
// for forward compatibility
type RoutingServiceServer interface {
	GetAvailability(RoutingService_GetAvailabilityServer) error
	mustEmbedUnimplementedRoutingServiceServer()
}

// UnimplementedRoutingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoutingServiceServer struct {
}

func (UnimplementedRoutingServiceServer) GetAvailability(RoutingService_GetAvailabilityServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAvailability not implemented")
}
func (UnimplementedRoutingServiceServer) mustEmbedUnimplementedRoutingServiceServer() {}

// UnsafeRoutingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoutingServiceServer will
// result in compilation errors.
type UnsafeRoutingServiceServer interface {
	mustEmbedUnimplementedRoutingServiceServer()
}

func RegisterRoutingServiceServer(s grpc.ServiceRegistrar, srv RoutingServiceServer) {
	s.RegisterService(&RoutingService_ServiceDesc, srv)
}

func _RoutingService_GetAvailability_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RoutingServiceServer).GetAvailability(&routingServiceGetAvailabilityServer{stream})
}

type RoutingService_GetAvailabilityServer interface {
	Send(*RoutingAvailabilityResponse) error
	Recv() (*RoutingAvailabilityRequest, error)
	grpc.ServerStream
}

type routingServiceGetAvailabilityServer struct {
	grpc.ServerStream
}

func (x *routingServiceGetAvailabilityServer) Send(m *RoutingAvailabilityResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routingServiceGetAvailabilityServer) Recv() (*RoutingAvailabilityRequest, error) {
	m := new(RoutingAvailabilityRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RoutingService_ServiceDesc is the grpc.ServiceDesc for RoutingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoutingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "routingpb.RoutingService",
	HandlerType: (*RoutingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAvailability",
			Handler:       _RoutingService_GetAvailability_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "routing.proto",
}
