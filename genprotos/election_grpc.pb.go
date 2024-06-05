// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: election.proto

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

// ElectionServiceClient is the client API for ElectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElectionServiceClient interface {
	CreateElection(ctx context.Context, in *Election, opts ...grpc.CallOption) (*Void, error)
	DeleteElection(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error)
	UpdateElection(ctx context.Context, in *Election, opts ...grpc.CallOption) (*Void, error)
	GetByIdElection(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Election, error)
	GetAllElections(ctx context.Context, in *Void, opts ...grpc.CallOption) (*GetAllElection, error)
}

type electionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewElectionServiceClient(cc grpc.ClientConnInterface) ElectionServiceClient {
	return &electionServiceClient{cc}
}

func (c *electionServiceClient) CreateElection(ctx context.Context, in *Election, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ElectionService/CreateElection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionServiceClient) DeleteElection(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ElectionService/DeleteElection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionServiceClient) UpdateElection(ctx context.Context, in *Election, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ElectionService/UpdateElection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionServiceClient) GetByIdElection(ctx context.Context, in *ById, opts ...grpc.CallOption) (*Election, error) {
	out := new(Election)
	err := c.cc.Invoke(ctx, "/protos.ElectionService/GetByIdElection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionServiceClient) GetAllElections(ctx context.Context, in *Void, opts ...grpc.CallOption) (*GetAllElection, error) {
	out := new(GetAllElection)
	err := c.cc.Invoke(ctx, "/protos.ElectionService/GetAllElections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElectionServiceServer is the server API for ElectionService service.
// All implementations must embed UnimplementedElectionServiceServer
// for forward compatibility
type ElectionServiceServer interface {
	CreateElection(context.Context, *Election) (*Void, error)
	DeleteElection(context.Context, *ById) (*Void, error)
	UpdateElection(context.Context, *Election) (*Void, error)
	GetByIdElection(context.Context, *ById) (*Election, error)
	GetAllElections(context.Context, *Void) (*GetAllElection, error)
	mustEmbedUnimplementedElectionServiceServer()
}

// UnimplementedElectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedElectionServiceServer struct {
}

func (UnimplementedElectionServiceServer) CreateElection(context.Context, *Election) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateElection not implemented")
}
func (UnimplementedElectionServiceServer) DeleteElection(context.Context, *ById) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteElection not implemented")
}
func (UnimplementedElectionServiceServer) UpdateElection(context.Context, *Election) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateElection not implemented")
}
func (UnimplementedElectionServiceServer) GetByIdElection(context.Context, *ById) (*Election, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdElection not implemented")
}
func (UnimplementedElectionServiceServer) GetAllElections(context.Context, *Void) (*GetAllElection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllElections not implemented")
}
func (UnimplementedElectionServiceServer) mustEmbedUnimplementedElectionServiceServer() {}

// UnsafeElectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElectionServiceServer will
// result in compilation errors.
type UnsafeElectionServiceServer interface {
	mustEmbedUnimplementedElectionServiceServer()
}

func RegisterElectionServiceServer(s grpc.ServiceRegistrar, srv ElectionServiceServer) {
	s.RegisterService(&ElectionService_ServiceDesc, srv)
}

func _ElectionService_CreateElection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Election)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).CreateElection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ElectionService/CreateElection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).CreateElection(ctx, req.(*Election))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectionService_DeleteElection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).DeleteElection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ElectionService/DeleteElection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).DeleteElection(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectionService_UpdateElection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Election)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).UpdateElection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ElectionService/UpdateElection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).UpdateElection(ctx, req.(*Election))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectionService_GetByIdElection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ById)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).GetByIdElection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ElectionService/GetByIdElection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).GetByIdElection(ctx, req.(*ById))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectionService_GetAllElections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionServiceServer).GetAllElections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ElectionService/GetAllElections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionServiceServer).GetAllElections(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

// ElectionService_ServiceDesc is the grpc.ServiceDesc for ElectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ElectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ElectionService",
	HandlerType: (*ElectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateElection",
			Handler:    _ElectionService_CreateElection_Handler,
		},
		{
			MethodName: "DeleteElection",
			Handler:    _ElectionService_DeleteElection_Handler,
		},
		{
			MethodName: "UpdateElection",
			Handler:    _ElectionService_UpdateElection_Handler,
		},
		{
			MethodName: "GetByIdElection",
			Handler:    _ElectionService_GetByIdElection_Handler,
		},
		{
			MethodName: "GetAllElections",
			Handler:    _ElectionService_GetAllElections_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "election.proto",
}