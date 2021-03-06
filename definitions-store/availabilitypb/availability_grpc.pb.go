// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: availability.proto

package availabilitypb

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

// AvailabilityServiceClient is the client API for AvailabilityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AvailabilityServiceClient interface {
	GetAvailability(ctx context.Context, in *AvailabilityRequest, opts ...grpc.CallOption) (AvailabilityService_GetAvailabilityClient, error)
}

type availabilityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAvailabilityServiceClient(cc grpc.ClientConnInterface) AvailabilityServiceClient {
	return &availabilityServiceClient{cc}
}

func (c *availabilityServiceClient) GetAvailability(ctx context.Context, in *AvailabilityRequest, opts ...grpc.CallOption) (AvailabilityService_GetAvailabilityClient, error) {
	stream, err := c.cc.NewStream(ctx, &AvailabilityService_ServiceDesc.Streams[0], "/availabilitypb.AvailabilityService/GetAvailability", opts...)
	if err != nil {
		return nil, err
	}
	x := &availabilityServiceGetAvailabilityClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AvailabilityService_GetAvailabilityClient interface {
	Recv() (*AvailabilityResponse, error)
	grpc.ClientStream
}

type availabilityServiceGetAvailabilityClient struct {
	grpc.ClientStream
}

func (x *availabilityServiceGetAvailabilityClient) Recv() (*AvailabilityResponse, error) {
	m := new(AvailabilityResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AvailabilityServiceServer is the server API for AvailabilityService service.
// All implementations must embed UnimplementedAvailabilityServiceServer
// for forward compatibility
type AvailabilityServiceServer interface {
	GetAvailability(*AvailabilityRequest, AvailabilityService_GetAvailabilityServer) error
	mustEmbedUnimplementedAvailabilityServiceServer()
}

// UnimplementedAvailabilityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAvailabilityServiceServer struct {
}

func (UnimplementedAvailabilityServiceServer) GetAvailability(*AvailabilityRequest, AvailabilityService_GetAvailabilityServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAvailability not implemented")
}
func (UnimplementedAvailabilityServiceServer) mustEmbedUnimplementedAvailabilityServiceServer() {}

// UnsafeAvailabilityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AvailabilityServiceServer will
// result in compilation errors.
type UnsafeAvailabilityServiceServer interface {
	mustEmbedUnimplementedAvailabilityServiceServer()
}

func RegisterAvailabilityServiceServer(s grpc.ServiceRegistrar, srv AvailabilityServiceServer) {
	s.RegisterService(&AvailabilityService_ServiceDesc, srv)
}

func _AvailabilityService_GetAvailability_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AvailabilityRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AvailabilityServiceServer).GetAvailability(m, &availabilityServiceGetAvailabilityServer{stream})
}

type AvailabilityService_GetAvailabilityServer interface {
	Send(*AvailabilityResponse) error
	grpc.ServerStream
}

type availabilityServiceGetAvailabilityServer struct {
	grpc.ServerStream
}

func (x *availabilityServiceGetAvailabilityServer) Send(m *AvailabilityResponse) error {
	return x.ServerStream.SendMsg(m)
}

// AvailabilityService_ServiceDesc is the grpc.ServiceDesc for AvailabilityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AvailabilityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "availabilitypb.AvailabilityService",
	HandlerType: (*AvailabilityServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAvailability",
			Handler:       _AvailabilityService_GetAvailability_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "availability.proto",
}
