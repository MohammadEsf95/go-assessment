// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.1
// source: proto/service1.proto

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
	GetDataFromService1_GetData_FullMethodName = "/contract.GetDataFromService1/GetData"
)

// GetDataFromService1Client is the client API for GetDataFromService1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GetDataFromService1Client interface {
	GetData(ctx context.Context, in *Service1Request, opts ...grpc.CallOption) (*Service1Response, error)
}

type getDataFromService1Client struct {
	cc grpc.ClientConnInterface
}

func NewGetDataFromService1Client(cc grpc.ClientConnInterface) GetDataFromService1Client {
	return &getDataFromService1Client{cc}
}

func (c *getDataFromService1Client) GetData(ctx context.Context, in *Service1Request, opts ...grpc.CallOption) (*Service1Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Service1Response)
	err := c.cc.Invoke(ctx, GetDataFromService1_GetData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetDataFromService1Server is the server API for GetDataFromService1 service.
// All implementations must embed UnimplementedGetDataFromService1Server
// for forward compatibility.
type GetDataFromService1Server interface {
	GetData(context.Context, *Service1Request) (*Service1Response, error)
	mustEmbedUnimplementedGetDataFromService1Server()
}

// UnimplementedGetDataFromService1Server must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGetDataFromService1Server struct{}

func (UnimplementedGetDataFromService1Server) GetData(context.Context, *Service1Request) (*Service1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}
func (UnimplementedGetDataFromService1Server) mustEmbedUnimplementedGetDataFromService1Server() {}
func (UnimplementedGetDataFromService1Server) testEmbeddedByValue()                             {}

// UnsafeGetDataFromService1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GetDataFromService1Server will
// result in compilation errors.
type UnsafeGetDataFromService1Server interface {
	mustEmbedUnimplementedGetDataFromService1Server()
}

func RegisterGetDataFromService1Server(s grpc.ServiceRegistrar, srv GetDataFromService1Server) {
	// If the following call pancis, it indicates UnimplementedGetDataFromService1Server was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GetDataFromService1_ServiceDesc, srv)
}

func _GetDataFromService1_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Service1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetDataFromService1Server).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GetDataFromService1_GetData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetDataFromService1Server).GetData(ctx, req.(*Service1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// GetDataFromService1_ServiceDesc is the grpc.ServiceDesc for GetDataFromService1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GetDataFromService1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contract.GetDataFromService1",
	HandlerType: (*GetDataFromService1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _GetDataFromService1_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service1.proto",
}
