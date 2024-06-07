// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: condidate.proto

package genproto

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

const (
	CandidateService_Create_FullMethodName = "/protos.CandidateService/Create"
	CandidateService_Get_FullMethodName    = "/protos.CandidateService/Get"
	CandidateService_GetAll_FullMethodName = "/protos.CandidateService/GetAll"
	CandidateService_Delete_FullMethodName = "/protos.CandidateService/Delete"
)

// CandidateServiceClient is the client API for CandidateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CandidateServiceClient interface {
	Create(ctx context.Context, in *CreateCandidateReq, opts ...grpc.CallOption) (*Void, error)
	Get(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*CandidateRes, error)
	GetAll(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*CandidatiesGetAllResp, error)
	Delete(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Void, error)
}

type candidateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCandidateServiceClient(cc grpc.ClientConnInterface) CandidateServiceClient {
	return &candidateServiceClient{cc}
}

func (c *candidateServiceClient) Create(ctx context.Context, in *CreateCandidateReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, CandidateService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *candidateServiceClient) Get(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*CandidateRes, error) {
	out := new(CandidateRes)
	err := c.cc.Invoke(ctx, CandidateService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *candidateServiceClient) GetAll(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*CandidatiesGetAllResp, error) {
	out := new(CandidatiesGetAllResp)
	err := c.cc.Invoke(ctx, CandidateService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *candidateServiceClient) Delete(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, CandidateService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CandidateServiceServer is the server API for CandidateService service.
// All implementations must embed UnimplementedCandidateServiceServer
// for forward compatibility
type CandidateServiceServer interface {
	Create(context.Context, *CreateCandidateReq) (*Void, error)
	Get(context.Context, *GetByIdReq) (*CandidateRes, error)
	GetAll(context.Context, *Filter) (*CandidatiesGetAllResp, error)
	Delete(context.Context, *GetByIdReq) (*Void, error)
	mustEmbedUnimplementedCandidateServiceServer()
}

// UnimplementedCandidateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCandidateServiceServer struct {
}

func (UnimplementedCandidateServiceServer) Create(context.Context, *CreateCandidateReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCandidateServiceServer) Get(context.Context, *GetByIdReq) (*CandidateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCandidateServiceServer) GetAll(context.Context, *Filter) (*CandidatiesGetAllResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedCandidateServiceServer) Delete(context.Context, *GetByIdReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCandidateServiceServer) mustEmbedUnimplementedCandidateServiceServer() {}

// UnsafeCandidateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CandidateServiceServer will
// result in compilation errors.
type UnsafeCandidateServiceServer interface {
	mustEmbedUnimplementedCandidateServiceServer()
}

func RegisterCandidateServiceServer(s grpc.ServiceRegistrar, srv CandidateServiceServer) {
	s.RegisterService(&CandidateService_ServiceDesc, srv)
}

func _CandidateService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCandidateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandidateServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CandidateService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandidateServiceServer).Create(ctx, req.(*CreateCandidateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CandidateService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandidateServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CandidateService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandidateServiceServer).Get(ctx, req.(*GetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CandidateService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandidateServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CandidateService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandidateServiceServer).GetAll(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _CandidateService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CandidateServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CandidateService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CandidateServiceServer).Delete(ctx, req.(*GetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CandidateService_ServiceDesc is the grpc.ServiceDesc for CandidateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CandidateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.CandidateService",
	HandlerType: (*CandidateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CandidateService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CandidateService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _CandidateService_GetAll_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CandidateService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "condidate.proto",
}
