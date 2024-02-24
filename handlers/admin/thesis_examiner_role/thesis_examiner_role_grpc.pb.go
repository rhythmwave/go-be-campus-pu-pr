// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: admin/thesis_examiner_role.proto

package thesis_examiner_role

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

// AdminThesisExaminerRoleHandlerClient is the client API for AdminThesisExaminerRoleHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminThesisExaminerRoleHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type adminThesisExaminerRoleHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminThesisExaminerRoleHandlerClient(cc grpc.ClientConnInterface) AdminThesisExaminerRoleHandlerClient {
	return &adminThesisExaminerRoleHandlerClient{cc}
}

func (c *adminThesisExaminerRoleHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/admin_thesis_examiner_role.AdminThesisExaminerRoleHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminThesisExaminerRoleHandlerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/admin_thesis_examiner_role.AdminThesisExaminerRoleHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminThesisExaminerRoleHandlerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/admin_thesis_examiner_role.AdminThesisExaminerRoleHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminThesisExaminerRoleHandlerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/admin_thesis_examiner_role.AdminThesisExaminerRoleHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminThesisExaminerRoleHandlerServer is the server API for AdminThesisExaminerRoleHandler service.
// All implementations must embed UnimplementedAdminThesisExaminerRoleHandlerServer
// for forward compatibility
type AdminThesisExaminerRoleHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedAdminThesisExaminerRoleHandlerServer()
}

// UnimplementedAdminThesisExaminerRoleHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAdminThesisExaminerRoleHandlerServer struct {
}

func (UnimplementedAdminThesisExaminerRoleHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdminThesisExaminerRoleHandlerServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAdminThesisExaminerRoleHandlerServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAdminThesisExaminerRoleHandlerServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAdminThesisExaminerRoleHandlerServer) mustEmbedUnimplementedAdminThesisExaminerRoleHandlerServer() {
}

// UnsafeAdminThesisExaminerRoleHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminThesisExaminerRoleHandlerServer will
// result in compilation errors.
type UnsafeAdminThesisExaminerRoleHandlerServer interface {
	mustEmbedUnimplementedAdminThesisExaminerRoleHandlerServer()
}

func RegisterAdminThesisExaminerRoleHandlerServer(s grpc.ServiceRegistrar, srv AdminThesisExaminerRoleHandlerServer) {
	s.RegisterService(&AdminThesisExaminerRoleHandler_ServiceDesc, srv)
}

func _AdminThesisExaminerRoleHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminThesisExaminerRoleHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_thesis_examiner_role.AdminThesisExaminerRoleHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminThesisExaminerRoleHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminThesisExaminerRoleHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminThesisExaminerRoleHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_thesis_examiner_role.AdminThesisExaminerRoleHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminThesisExaminerRoleHandlerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminThesisExaminerRoleHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminThesisExaminerRoleHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_thesis_examiner_role.AdminThesisExaminerRoleHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminThesisExaminerRoleHandlerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminThesisExaminerRoleHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminThesisExaminerRoleHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_thesis_examiner_role.AdminThesisExaminerRoleHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminThesisExaminerRoleHandlerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminThesisExaminerRoleHandler_ServiceDesc is the grpc.ServiceDesc for AdminThesisExaminerRoleHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminThesisExaminerRoleHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_thesis_examiner_role.AdminThesisExaminerRoleHandler",
	HandlerType: (*AdminThesisExaminerRoleHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _AdminThesisExaminerRoleHandler_GetList_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AdminThesisExaminerRoleHandler_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AdminThesisExaminerRoleHandler_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AdminThesisExaminerRoleHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/thesis_examiner_role.proto",
}
