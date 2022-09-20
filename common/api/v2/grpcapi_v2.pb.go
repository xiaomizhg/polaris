// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpcapi_v2.proto

package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PolarisGRPCClient is the client API for PolarisGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PolarisGRPCClient interface {
	// 统一发现接口
	Discover(ctx context.Context, opts ...grpc.CallOption) (PolarisGRPC_DiscoverClient, error)
}

type polarisGRPCClient struct {
	cc *grpc.ClientConn
}

func NewPolarisGRPCClient(cc *grpc.ClientConn) PolarisGRPCClient {
	return &polarisGRPCClient{cc}
}

func (c *polarisGRPCClient) Discover(ctx context.Context, opts ...grpc.CallOption) (PolarisGRPC_DiscoverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PolarisGRPC_serviceDesc.Streams[0], "/v2.PolarisGRPC/Discover", opts...)
	if err != nil {
		return nil, err
	}
	x := &polarisGRPCDiscoverClient{stream}
	return x, nil
}

type PolarisGRPC_DiscoverClient interface {
	Send(*DiscoverRequest) error
	Recv() (*DiscoverResponse, error)
	grpc.ClientStream
}

type polarisGRPCDiscoverClient struct {
	grpc.ClientStream
}

func (x *polarisGRPCDiscoverClient) Send(m *DiscoverRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *polarisGRPCDiscoverClient) Recv() (*DiscoverResponse, error) {
	m := new(DiscoverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PolarisGRPCServer is the server API for PolarisGRPC service.
type PolarisGRPCServer interface {
	// 统一发现接口
	Discover(PolarisGRPC_DiscoverServer) error
}

func RegisterPolarisGRPCServer(s *grpc.Server, srv PolarisGRPCServer) {
	s.RegisterService(&_PolarisGRPC_serviceDesc, srv)
}

func _PolarisGRPC_Discover_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PolarisGRPCServer).Discover(&polarisGRPCDiscoverServer{stream})
}

type PolarisGRPC_DiscoverServer interface {
	Send(*DiscoverResponse) error
	Recv() (*DiscoverRequest, error)
	grpc.ServerStream
}

type polarisGRPCDiscoverServer struct {
	grpc.ServerStream
}

func (x *polarisGRPCDiscoverServer) Send(m *DiscoverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *polarisGRPCDiscoverServer) Recv() (*DiscoverRequest, error) {
	m := new(DiscoverRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PolarisGRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v2.PolarisGRPC",
	HandlerType: (*PolarisGRPCServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Discover",
			Handler:       _PolarisGRPC_Discover_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpcapi_v2.proto",
}

func init() { proto.RegisterFile("grpcapi_v2.proto", fileDescriptor_grpcapi_v2_99138fa445bc3382) }

var fileDescriptor_grpcapi_v2_99138fa445bc3382 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2f, 0x2a, 0x48,
	0x4e, 0x2c, 0xc8, 0x8c, 0x2f, 0x33, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33,
	0x92, 0x12, 0x28, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x81, 0x8b, 0x4a, 0x09, 0x16, 0xa5, 0x16,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0xc2, 0x85, 0x8c, 0xbc, 0xb8, 0xb8, 0x03, 0xf2, 0x73, 0x12, 0x8b,
	0x32, 0x8b, 0xdd, 0x83, 0x02, 0x9c, 0x85, 0xac, 0xb9, 0x38, 0x5c, 0x32, 0x8b, 0x93, 0xf3, 0xcb,
	0x52, 0x8b, 0x84, 0x84, 0xf5, 0xca, 0x8c, 0xf4, 0x60, 0xbc, 0x20, 0x88, 0x59, 0x52, 0x22, 0xa8,
	0x82, 0x10, 0xe3, 0x94, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x93, 0xd8, 0xc0, 0x46, 0x1a, 0x03, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xe0, 0x0c, 0xae, 0x02, 0x8f, 0x00, 0x00, 0x00,
}