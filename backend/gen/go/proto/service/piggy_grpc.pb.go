// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/service/piggy.proto

package service

import (
	context "context"
	v1 "github.com/Exca-DK/pegism/gen/go/proto/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PiggyService_UpdatePiggyName_FullMethodName     = "/pegism.service.PiggyService/UpdatePiggyName"
	PiggyService_GetPiggy_FullMethodName            = "/pegism.service.PiggyService/GetPiggy"
	PiggyService_GetPiggyFromProfile_FullMethodName = "/pegism.service.PiggyService/GetPiggyFromProfile"
	PiggyService_GetPiggyFromName_FullMethodName    = "/pegism.service.PiggyService/GetPiggyFromName"
)

// PiggyServiceClient is the client API for PiggyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PiggyServiceClient interface {
	UpdatePiggyName(ctx context.Context, in *v1.UpdatePiggyNameRequest, opts ...grpc.CallOption) (*v1.UpdatePiggyNameResponse, error)
	GetPiggy(ctx context.Context, in *v1.GetPiggyRequest, opts ...grpc.CallOption) (*v1.GetPiggyResponse, error)
	GetPiggyFromProfile(ctx context.Context, in *v1.GetPiggyRequest, opts ...grpc.CallOption) (*v1.GetPiggyResponse, error)
	GetPiggyFromName(ctx context.Context, in *v1.GetPiggyFromNameRequest, opts ...grpc.CallOption) (*v1.GetPiggyFromNameResponse, error)
}

type piggyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPiggyServiceClient(cc grpc.ClientConnInterface) PiggyServiceClient {
	return &piggyServiceClient{cc}
}

func (c *piggyServiceClient) UpdatePiggyName(ctx context.Context, in *v1.UpdatePiggyNameRequest, opts ...grpc.CallOption) (*v1.UpdatePiggyNameResponse, error) {
	out := new(v1.UpdatePiggyNameResponse)
	err := c.cc.Invoke(ctx, PiggyService_UpdatePiggyName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piggyServiceClient) GetPiggy(ctx context.Context, in *v1.GetPiggyRequest, opts ...grpc.CallOption) (*v1.GetPiggyResponse, error) {
	out := new(v1.GetPiggyResponse)
	err := c.cc.Invoke(ctx, PiggyService_GetPiggy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piggyServiceClient) GetPiggyFromProfile(ctx context.Context, in *v1.GetPiggyRequest, opts ...grpc.CallOption) (*v1.GetPiggyResponse, error) {
	out := new(v1.GetPiggyResponse)
	err := c.cc.Invoke(ctx, PiggyService_GetPiggyFromProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *piggyServiceClient) GetPiggyFromName(ctx context.Context, in *v1.GetPiggyFromNameRequest, opts ...grpc.CallOption) (*v1.GetPiggyFromNameResponse, error) {
	out := new(v1.GetPiggyFromNameResponse)
	err := c.cc.Invoke(ctx, PiggyService_GetPiggyFromName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PiggyServiceServer is the server API for PiggyService service.
// All implementations must embed UnimplementedPiggyServiceServer
// for forward compatibility
type PiggyServiceServer interface {
	UpdatePiggyName(context.Context, *v1.UpdatePiggyNameRequest) (*v1.UpdatePiggyNameResponse, error)
	GetPiggy(context.Context, *v1.GetPiggyRequest) (*v1.GetPiggyResponse, error)
	GetPiggyFromProfile(context.Context, *v1.GetPiggyRequest) (*v1.GetPiggyResponse, error)
	GetPiggyFromName(context.Context, *v1.GetPiggyFromNameRequest) (*v1.GetPiggyFromNameResponse, error)
	mustEmbedUnimplementedPiggyServiceServer()
}

// UnimplementedPiggyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPiggyServiceServer struct {
}

func (UnimplementedPiggyServiceServer) UpdatePiggyName(context.Context, *v1.UpdatePiggyNameRequest) (*v1.UpdatePiggyNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePiggyName not implemented")
}
func (UnimplementedPiggyServiceServer) GetPiggy(context.Context, *v1.GetPiggyRequest) (*v1.GetPiggyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPiggy not implemented")
}
func (UnimplementedPiggyServiceServer) GetPiggyFromProfile(context.Context, *v1.GetPiggyRequest) (*v1.GetPiggyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPiggyFromProfile not implemented")
}
func (UnimplementedPiggyServiceServer) GetPiggyFromName(context.Context, *v1.GetPiggyFromNameRequest) (*v1.GetPiggyFromNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPiggyFromName not implemented")
}
func (UnimplementedPiggyServiceServer) mustEmbedUnimplementedPiggyServiceServer() {}

// UnsafePiggyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PiggyServiceServer will
// result in compilation errors.
type UnsafePiggyServiceServer interface {
	mustEmbedUnimplementedPiggyServiceServer()
}

func RegisterPiggyServiceServer(s grpc.ServiceRegistrar, srv PiggyServiceServer) {
	s.RegisterService(&PiggyService_ServiceDesc, srv)
}

func _PiggyService_UpdatePiggyName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.UpdatePiggyNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiggyServiceServer).UpdatePiggyName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PiggyService_UpdatePiggyName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiggyServiceServer).UpdatePiggyName(ctx, req.(*v1.UpdatePiggyNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiggyService_GetPiggy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetPiggyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiggyServiceServer).GetPiggy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PiggyService_GetPiggy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiggyServiceServer).GetPiggy(ctx, req.(*v1.GetPiggyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiggyService_GetPiggyFromProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetPiggyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiggyServiceServer).GetPiggyFromProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PiggyService_GetPiggyFromProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiggyServiceServer).GetPiggyFromProfile(ctx, req.(*v1.GetPiggyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PiggyService_GetPiggyFromName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetPiggyFromNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PiggyServiceServer).GetPiggyFromName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PiggyService_GetPiggyFromName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PiggyServiceServer).GetPiggyFromName(ctx, req.(*v1.GetPiggyFromNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PiggyService_ServiceDesc is the grpc.ServiceDesc for PiggyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PiggyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pegism.service.PiggyService",
	HandlerType: (*PiggyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdatePiggyName",
			Handler:    _PiggyService_UpdatePiggyName_Handler,
		},
		{
			MethodName: "GetPiggy",
			Handler:    _PiggyService_GetPiggy_Handler,
		},
		{
			MethodName: "GetPiggyFromProfile",
			Handler:    _PiggyService_GetPiggyFromProfile_Handler,
		},
		{
			MethodName: "GetPiggyFromName",
			Handler:    _PiggyService_GetPiggyFromName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service/piggy.proto",
}