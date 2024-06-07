// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: public.proto

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
	PublicService_Create_FullMethodName = "/protos.PublicService/Create"
	PublicService_Get_FullMethodName    = "/protos.PublicService/Get"
	PublicService_GetAll_FullMethodName = "/protos.PublicService/GetAll"
	PublicService_Update_FullMethodName = "/protos.PublicService/Update"
	PublicService_Delete_FullMethodName = "/protos.PublicService/Delete"
)

// PublicServiceClient is the client API for PublicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicServiceClient interface {
	Create(ctx context.Context, in *PublicCreate, opts ...grpc.CallOption) (*Void, error)
	Get(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*PublicRes, error)
	GetAll(ctx context.Context, in *GetAllPublicsRequest, opts ...grpc.CallOption) (*GetAllPublicsResponse, error)
	Update(ctx context.Context, in *PublicUpdate, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Void, error)
}

type publicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicServiceClient(cc grpc.ClientConnInterface) PublicServiceClient {
	return &publicServiceClient{cc}
}

func (c *publicServiceClient) Create(ctx context.Context, in *PublicCreate, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, PublicService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicServiceClient) Get(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*PublicRes, error) {
	out := new(PublicRes)
	err := c.cc.Invoke(ctx, PublicService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicServiceClient) GetAll(ctx context.Context, in *GetAllPublicsRequest, opts ...grpc.CallOption) (*GetAllPublicsResponse, error) {
	out := new(GetAllPublicsResponse)
	err := c.cc.Invoke(ctx, PublicService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicServiceClient) Update(ctx context.Context, in *PublicUpdate, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, PublicService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicServiceClient) Delete(ctx context.Context, in *GetByIdReq, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, PublicService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicServiceServer is the server API for PublicService service.
// All implementations must embed UnimplementedPublicServiceServer
// for forward compatibility
type PublicServiceServer interface {
	Create(context.Context, *PublicCreate) (*Void, error)
	Get(context.Context, *GetByIdReq) (*PublicRes, error)
	GetAll(context.Context, *GetAllPublicsRequest) (*GetAllPublicsResponse, error)
	Update(context.Context, *PublicUpdate) (*Void, error)
	Delete(context.Context, *GetByIdReq) (*Void, error)
	mustEmbedUnimplementedPublicServiceServer()
}

// UnimplementedPublicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublicServiceServer struct {
}

func (UnimplementedPublicServiceServer) Create(context.Context, *PublicCreate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPublicServiceServer) Get(context.Context, *GetByIdReq) (*PublicRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPublicServiceServer) GetAll(context.Context, *GetAllPublicsRequest) (*GetAllPublicsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPublicServiceServer) Update(context.Context, *PublicUpdate) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPublicServiceServer) Delete(context.Context, *GetByIdReq) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
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

func _PublicService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).Create(ctx, req.(*PublicCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).Get(ctx, req.(*GetByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllPublicsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).GetAll(ctx, req.(*GetAllPublicsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).Update(ctx, req.(*PublicUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicServiceServer).Delete(ctx, req.(*GetByIdReq))
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
			MethodName: "Create",
			Handler:    _PublicService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PublicService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PublicService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PublicService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PublicService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "public.proto",
}
