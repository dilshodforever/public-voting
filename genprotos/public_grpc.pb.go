// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: public.proto

package genprotos

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

// PublicServiceClient is the client API for PublicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicServiceClient interface {
	CreatePublic(ctx context.Context, in *Public, opts ...grpc.CallOption) (*Void, error)
	DeletePublic(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	UpdatePublic(ctx context.Context, in *Public, opts ...grpc.CallOption) (*Void, error)
	GetByIdPublic(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Public, error)
	GetAllPublics(ctx context.Context, in *Void, opts ...grpc.CallOption) (*GetAllPublic, error)
}

type publicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicServiceClient(cc grpc.ClientConnInterface) PublicServiceClient {
	return &publicServiceClient{cc}
}

func (c *publicServiceClient) CreatePublic(ctx context.Context, in *Public, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.PublicService/CreatePublic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicServiceClient) DeletePublic(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.PublicService/DeletePublic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicServiceClient) UpdatePublic(ctx context.Context, in *Public, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.PublicService/UpdatePublic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicServiceClient) GetByIdPublic(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Public, error) {
	out := new(Public)
	err := c.cc.Invoke(ctx, "/protos.PublicService/GetByIdPublic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicServiceClient) GetAllPublics(ctx context.Context, in *Void, opts ...grpc.CallOption) (*GetAllPublic, error) {
	out := new(GetAllPublic)
	err := c.cc.Invoke(ctx, "/protos.PublicService/GetAllPublics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicServiceServer is the server API for PublicService service.
// All implementations must embed UnimplementedPublicServiceServer
// for forward compatibility
type PublicServiceServer interface {
	CreatePublic(context.Context, *Public) (*Void, error)
	DeletePublic(context.Context, *ById) (*Void, error)
	UpdatePublic(context.Context, *Public) (*Void, error)
	GetByIdPublic(context.Context, *ById) (*Public, error)
	GetAllPublics(context.Context, *Void) (*GetAllPublic, error)
	mustEmbedUnimplementedPublicServiceServer()
}

// UnimplementedPublicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublicServiceServer struct {
}

func (UnimplementedPublicServiceServer) CreatePublic(context.Context, *Public) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePublic not implemented")
}
func (UnimplementedPublicServiceServer) DeletePublic(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePublic not implemented")
}
func (UnimplementedPublicServiceServer) UpdatePublic(context.Context, *Public) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePublic not implemented")
}
func (UnimplementedPublicServiceServer) GetByIdPublic(context.Context, *ById) (*Public, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdPublic not implemented")
}
func (UnimplementedPublicServiceServer) GetAllPublics(context.Context, *Void) (*GetAllPublic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPublics not implemented")
}
func (UnimplementedPublicServiceServer) mustEmbedUnimplementedPublicServiceServer() {}

// UnsafePublicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicServiceServer will
// result in compilation errors.
type UnsafePublicServiceServer interface {
	mustEmbedUnimplementedPublicServiceServer()
}

func RegisterPublicServiceServer(s grpc.ServiceRegistrar, srv PublicServiceServer) {
	s.RegisterService(&PublicService_ServiceDesc, srv)
}

func _PublicService_CreatePublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Public)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).CreatePublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PublicService/CreatePublic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).CreatePublic(ctx, req.(*Public))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicService_DeletePublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).DeletePublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PublicService/DeletePublic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).DeletePublic(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicService_UpdatePublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Public)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).UpdatePublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PublicService/UpdatePublic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).UpdatePublic(ctx, req.(*Public))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicService_GetByIdPublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).GetByIdPublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PublicService/GetByIdPublic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).GetByIdPublic(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicService_GetAllPublics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).GetAllPublics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.PublicService/GetAllPublics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).GetAllPublics(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// PublicService_ServiceDesc is the grpc.ServiceDesc for PublicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.PublicService",
	HandlerType: (*PublicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePublic",
			Handler:    _PublicService_CreatePublic_Handler,
		},
		{
			MethodName: "DeletePublic",
			Handler:    _PublicService_DeletePublic_Handler,
		},
		{
			MethodName: "UpdatePublic",
			Handler:    _PublicService_UpdatePublic_Handler,
		},
		{
			MethodName: "GetByIdPublic",
			Handler:    _PublicService_GetByIdPublic_Handler,
		},
		{
			MethodName: "GetAllPublics",
			Handler:    _PublicService_GetAllPublics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public.proto",
}
