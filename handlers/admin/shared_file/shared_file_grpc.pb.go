// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: admin/shared_file.proto

package shared_file

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

// AdminSharedFileHandlerClient is the client API for AdminSharedFileHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminSharedFileHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	Approve(ctx context.Context, in *ApproveRequest, opts ...grpc.CallOption) (*ApproveResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type adminSharedFileHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminSharedFileHandlerClient(cc grpc.ClientConnInterface) AdminSharedFileHandlerClient {
	return &adminSharedFileHandlerClient{cc}
}

func (c *adminSharedFileHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/admin_shared_file.AdminSharedFileHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminSharedFileHandlerClient) Approve(ctx context.Context, in *ApproveRequest, opts ...grpc.CallOption) (*ApproveResponse, error) {
	out := new(ApproveResponse)
	err := c.cc.Invoke(ctx, "/admin_shared_file.AdminSharedFileHandler/Approve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminSharedFileHandlerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/admin_shared_file.AdminSharedFileHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminSharedFileHandlerServer is the server API for AdminSharedFileHandler service.
// All implementations must embed UnimplementedAdminSharedFileHandlerServer
// for forward compatibility
type AdminSharedFileHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	Approve(context.Context, *ApproveRequest) (*ApproveResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedAdminSharedFileHandlerServer()
}

// UnimplementedAdminSharedFileHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAdminSharedFileHandlerServer struct {
}

func (UnimplementedAdminSharedFileHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdminSharedFileHandlerServer) Approve(context.Context, *ApproveRequest) (*ApproveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Approve not implemented")
}
func (UnimplementedAdminSharedFileHandlerServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAdminSharedFileHandlerServer) mustEmbedUnimplementedAdminSharedFileHandlerServer() {
}

// UnsafeAdminSharedFileHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminSharedFileHandlerServer will
// result in compilation errors.
type UnsafeAdminSharedFileHandlerServer interface {
	mustEmbedUnimplementedAdminSharedFileHandlerServer()
}

func RegisterAdminSharedFileHandlerServer(s grpc.ServiceRegistrar, srv AdminSharedFileHandlerServer) {
	s.RegisterService(&AdminSharedFileHandler_ServiceDesc, srv)
}

func _AdminSharedFileHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminSharedFileHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_shared_file.AdminSharedFileHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminSharedFileHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminSharedFileHandler_Approve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminSharedFileHandlerServer).Approve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_shared_file.AdminSharedFileHandler/Approve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminSharedFileHandlerServer).Approve(ctx, req.(*ApproveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminSharedFileHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminSharedFileHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_shared_file.AdminSharedFileHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminSharedFileHandlerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminSharedFileHandler_ServiceDesc is the grpc.ServiceDesc for AdminSharedFileHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminSharedFileHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_shared_file.AdminSharedFileHandler",
	HandlerType: (*AdminSharedFileHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _AdminSharedFileHandler_GetList_Handler,
		},
		{
			MethodName: "Approve",
			Handler:    _AdminSharedFileHandler_Approve_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AdminSharedFileHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/shared_file.proto",
}