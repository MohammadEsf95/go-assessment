// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.1
// source: proto/service2.proto

package contract

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GetDataFromService2_GetData_FullMethodName = "/contract.GetDataFromService2/GetData"
)

// GetDataFromService2Client is the client API for GetDataFromService2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetDataFromService2Client interface {
	GetData(ctx context.Context, in *Service2Request, opts ...grpc.CallOption) (*Service2Response, error)
}

type getDataFromService2Client struct {
	cc grpc.ClientConnInterface
}

func NewGetDataFromService2Client(cc grpc.ClientConnInterface) GetDataFromService2Client {
	return &getDataFromService2Client{cc}
}

func (c *getDataFromService2Client) GetData(ctx context.Context, in *Service2Request, opts ...grpc.CallOption) (*Service2Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Service2Response)
	err := c.cc.Invoke(ctx, GetDataFromService2_GetData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetDataFromService2Server is the cmd API for GetDataFromService2 service.
// All implementations must embed UnimplementedGetDataFromService2Server
// for forward compatibility.
type GetDataFromService2Server interface {
	GetData(context.Context, *Service2Request) (*Service2Response, error)
	mustEmbedUnimplementedGetDataFromService2Server()
}

// UnimplementedGetDataFromService2Server must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGetDataFromService2Server struct{}

func (UnimplementedGetDataFromService2Server) GetData(context.Context, *Service2Request) (*Service2Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedGetDataFromService2Server) mustEmbedUnimplementedGetDataFromService2Server() {}
func (UnimplementedGetDataFromService2Server) testEmbeddedByValue()                             {}

// UnsafeGetDataFromService2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetDataFromService2Server will
// result in compilation errors.
type UnsafeGetDataFromService2Server interface {
	mustEmbedUnimplementedGetDataFromService2Server()
}

func RegisterGetDataFromService2Server(s grpc.ServiceRegistrar, srv GetDataFromService2Server) {
	// If the following call pancis, it indicates UnimplementedGetDataFromService2Server was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GetDataFromService2_ServiceDesc, srv)
}

func _GetDataFromService2_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetDataFromService2Server).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetDataFromService2_GetData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetDataFromService2Server).GetData(ctx, req.(*Service2Request))
	}
	return interceptor(ctx, in, info, handler)
}

// GetDataFromService2_ServiceDesc is the grpc.ServiceDesc for GetDataFromService2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetDataFromService2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contract.GetDataFromService2",
	HandlerType: (*GetDataFromService2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _GetDataFromService2_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service2.proto",
}
