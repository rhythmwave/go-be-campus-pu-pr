// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: root/major.proto

package major

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

// RootMajorHandlerClient is the client API for RootMajorHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RootMajorHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc.CallOption) (*GetDetailResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type rootMajorHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewRootMajorHandlerClient(cc grpc.ClientConnInterface) RootMajorHandlerClient {
	return &rootMajorHandlerClient{cc}
}

func (c *rootMajorHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/root_major.RootMajorHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rootMajorHandlerClient) GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc.CallOption) (*GetDetailResponse, error) {
	out := new(GetDetailResponse)
	err := c.cc.Invoke(ctx, "/root_major.RootMajorHandler/GetDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rootMajorHandlerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/root_major.RootMajorHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rootMajorHandlerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/root_major.RootMajorHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rootMajorHandlerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/root_major.RootMajorHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RootMajorHandlerServer is the server API for RootMajorHandler service.
// All implementations must embed UnimplementedRootMajorHandlerServer
// for forward compatibility
type RootMajorHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	GetDetail(context.Context, *GetDetailRequest) (*GetDetailResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedRootMajorHandlerServer()
}

// UnimplementedRootMajorHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedRootMajorHandlerServer struct {
}

func (UnimplementedRootMajorHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedRootMajorHandlerServer) GetDetail(context.Context, *GetDetailRequest) (*GetDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetail not implemented")
}
func (UnimplementedRootMajorHandlerServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRootMajorHandlerServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRootMajorHandlerServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRootMajorHandlerServer) mustEmbedUnimplementedRootMajorHandlerServer() {}

// UnsafeRootMajorHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RootMajorHandlerServer will
// result in compilation errors.
type UnsafeRootMajorHandlerServer interface {
	mustEmbedUnimplementedRootMajorHandlerServer()
}

func RegisterRootMajorHandlerServer(s grpc.ServiceRegistrar, srv RootMajorHandlerServer) {
	s.RegisterService(&RootMajorHandler_ServiceDesc, srv)
}

func _RootMajorHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootMajorHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/root_major.RootMajorHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootMajorHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RootMajorHandler_GetDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootMajorHandlerServer).GetDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/root_major.RootMajorHandler/GetDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootMajorHandlerServer).GetDetail(ctx, req.(*GetDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RootMajorHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootMajorHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/root_major.RootMajorHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootMajorHandlerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RootMajorHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootMajorHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/root_major.RootMajorHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootMajorHandlerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RootMajorHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootMajorHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/root_major.RootMajorHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootMajorHandlerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RootMajorHandler_ServiceDesc is the grpc.ServiceDesc for RootMajorHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RootMajorHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "root_major.RootMajorHandler",
	HandlerType: (*RootMajorHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _RootMajorHandler_GetList_Handler,
		},
		{
			MethodName: "GetDetail",
			Handler:    _RootMajorHandler_GetDetail_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RootMajorHandler_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RootMajorHandler_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RootMajorHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "root/major.proto",
}