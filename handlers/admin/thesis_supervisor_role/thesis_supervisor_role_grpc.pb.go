// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: admin/thesis_supervisor_role.proto

package thesis_supervisor_role

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

// AdminThesisSupervisorRoleHandlerClient is the client API for AdminThesisSupervisorRoleHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminThesisSupervisorRoleHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type adminThesisSupervisorRoleHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminThesisSupervisorRoleHandlerClient(cc grpc.ClientConnInterface) AdminThesisSupervisorRoleHandlerClient {
	return &adminThesisSupervisorRoleHandlerClient{cc}
}

func (c *adminThesisSupervisorRoleHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/admin_thesis_supervisor_role.AdminThesisSupervisorRoleHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminThesisSupervisorRoleHandlerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/admin_thesis_supervisor_role.AdminThesisSupervisorRoleHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminThesisSupervisorRoleHandlerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/admin_thesis_supervisor_role.AdminThesisSupervisorRoleHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminThesisSupervisorRoleHandlerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/admin_thesis_supervisor_role.AdminThesisSupervisorRoleHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminThesisSupervisorRoleHandlerServer is the server API for AdminThesisSupervisorRoleHandler service.
// All implementations must embed UnimplementedAdminThesisSupervisorRoleHandlerServer
// for forward compatibility
type AdminThesisSupervisorRoleHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedAdminThesisSupervisorRoleHandlerServer()
}

// UnimplementedAdminThesisSupervisorRoleHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAdminThesisSupervisorRoleHandlerServer struct {
}

func (UnimplementedAdminThesisSupervisorRoleHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdminThesisSupervisorRoleHandlerServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAdminThesisSupervisorRoleHandlerServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAdminThesisSupervisorRoleHandlerServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAdminThesisSupervisorRoleHandlerServer) mustEmbedUnimplementedAdminThesisSupervisorRoleHandlerServer() {
}

// UnsafeAdminThesisSupervisorRoleHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminThesisSupervisorRoleHandlerServer will
// result in compilation errors.
type UnsafeAdminThesisSupervisorRoleHandlerServer interface {
	mustEmbedUnimplementedAdminThesisSupervisorRoleHandlerServer()
}

func RegisterAdminThesisSupervisorRoleHandlerServer(s grpc.ServiceRegistrar, srv AdminThesisSupervisorRoleHandlerServer) {
	s.RegisterService(&AdminThesisSupervisorRoleHandler_ServiceDesc, srv)
}

func _AdminThesisSupervisorRoleHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminThesisSupervisorRoleHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_thesis_supervisor_role.AdminThesisSupervisorRoleHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminThesisSupervisorRoleHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminThesisSupervisorRoleHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminThesisSupervisorRoleHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_thesis_supervisor_role.AdminThesisSupervisorRoleHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminThesisSupervisorRoleHandlerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminThesisSupervisorRoleHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminThesisSupervisorRoleHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_thesis_supervisor_role.AdminThesisSupervisorRoleHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminThesisSupervisorRoleHandlerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminThesisSupervisorRoleHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminThesisSupervisorRoleHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_thesis_supervisor_role.AdminThesisSupervisorRoleHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminThesisSupervisorRoleHandlerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminThesisSupervisorRoleHandler_ServiceDesc is the grpc.ServiceDesc for AdminThesisSupervisorRoleHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminThesisSupervisorRoleHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_thesis_supervisor_role.AdminThesisSupervisorRoleHandler",
	HandlerType: (*AdminThesisSupervisorRoleHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _AdminThesisSupervisorRoleHandler_GetList_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AdminThesisSupervisorRoleHandler_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AdminThesisSupervisorRoleHandler_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AdminThesisSupervisorRoleHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/thesis_supervisor_role.proto",
}
