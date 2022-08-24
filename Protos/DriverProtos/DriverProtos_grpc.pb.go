// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: Protos/DriverProtos/DriverProtos.proto

package DriverProtos

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

// DriverServiceClient is the client API for DriverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DriverServiceClient interface {
	AddDriverLocation(ctx context.Context, in *DriverDetails, opts ...grpc.CallOption) (*DriverResponse, error)
	SearchForDrivers(ctx context.Context, in *DriverSearch, opts ...grpc.CallOption) (*SearchResponse, error)
}

type driverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverServiceClient(cc grpc.ClientConnInterface) DriverServiceClient {
	return &driverServiceClient{cc}
}

func (c *driverServiceClient) AddDriverLocation(ctx context.Context, in *DriverDetails, opts ...grpc.CallOption) (*DriverResponse, error) {
	out := new(DriverResponse)
	err := c.cc.Invoke(ctx, "/DriverService/AddDriverLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverServiceClient) SearchForDrivers(ctx context.Context, in *DriverSearch, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/DriverService/SearchForDrivers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServiceServer is the server API for DriverService service.
// All implementations must embed UnimplementedDriverServiceServer
// for forward compatibility
type DriverServiceServer interface {
	AddDriverLocation(context.Context, *DriverDetails) (*DriverResponse, error)
	SearchForDrivers(context.Context, *DriverSearch) (*SearchResponse, error)
	mustEmbedUnimplementedDriverServiceServer()
}

// UnimplementedDriverServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDriverServiceServer struct {
}

func (UnimplementedDriverServiceServer) AddDriverLocation(context.Context, *DriverDetails) (*DriverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDriverLocation not implemented")
}
func (UnimplementedDriverServiceServer) SearchForDrivers(context.Context, *DriverSearch) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchForDrivers not implemented")
}
func (UnimplementedDriverServiceServer) mustEmbedUnimplementedDriverServiceServer() {}

// UnsafeDriverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverServiceServer will
// result in compilation errors.
type UnsafeDriverServiceServer interface {
	mustEmbedUnimplementedDriverServiceServer()
}

func RegisterDriverServiceServer(s grpc.ServiceRegistrar, srv DriverServiceServer) {
	s.RegisterService(&DriverService_ServiceDesc, srv)
}

func _DriverService_AddDriverLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).AddDriverLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DriverService/AddDriverLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).AddDriverLocation(ctx, req.(*DriverDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverService_SearchForDrivers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverSearch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).SearchForDrivers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DriverService/SearchForDrivers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).SearchForDrivers(ctx, req.(*DriverSearch))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverService_ServiceDesc is the grpc.ServiceDesc for DriverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DriverService",
	HandlerType: (*DriverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDriverLocation",
			Handler:    _DriverService_AddDriverLocation_Handler,
		},
		{
			MethodName: "SearchForDrivers",
			Handler:    _DriverService_SearchForDrivers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Protos/DriverProtos/DriverProtos.proto",
}
