// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: admin/accreditation.proto

package accreditation

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

// AdminAccreditationHandlerClient is the client API for AdminAccreditationHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminAccreditationHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type adminAccreditationHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminAccreditationHandlerClient(cc grpc.ClientConnInterface) AdminAccreditationHandlerClient {
	return &adminAccreditationHandlerClient{cc}
}

func (c *adminAccreditationHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/admin_accreditation.AdminAccreditationHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAccreditationHandlerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/admin_accreditation.AdminAccreditationHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAccreditationHandlerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/admin_accreditation.AdminAccreditationHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAccreditationHandlerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/admin_accreditation.AdminAccreditationHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminAccreditationHandlerServer is the server API for AdminAccreditationHandler service.
// All implementations must embed UnimplementedAdminAccreditationHandlerServer
// for forward compatibility
type AdminAccreditationHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedAdminAccreditationHandlerServer()
}

// UnimplementedAdminAccreditationHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAdminAccreditationHandlerServer struct {
}

func (UnimplementedAdminAccreditationHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdminAccreditationHandlerServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAdminAccreditationHandlerServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAdminAccreditationHandlerServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAdminAccreditationHandlerServer) mustEmbedUnimplementedAdminAccreditationHandlerServer() {
}

// UnsafeAdminAccreditationHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminAccreditationHandlerServer will
// result in compilation errors.
type UnsafeAdminAccreditationHandlerServer interface {
	mustEmbedUnimplementedAdminAccreditationHandlerServer()
}

func RegisterAdminAccreditationHandlerServer(s grpc.ServiceRegistrar, srv AdminAccreditationHandlerServer) {
	s.RegisterService(&AdminAccreditationHandler_ServiceDesc, srv)
}

func _AdminAccreditationHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAccreditationHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_accreditation.AdminAccreditationHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAccreditationHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAccreditationHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAccreditationHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_accreditation.AdminAccreditationHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAccreditationHandlerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAccreditationHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAccreditationHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_accreditation.AdminAccreditationHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAccreditationHandlerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAccreditationHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAccreditationHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_accreditation.AdminAccreditationHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAccreditationHandlerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminAccreditationHandler_ServiceDesc is the grpc.ServiceDesc for AdminAccreditationHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminAccreditationHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_accreditation.AdminAccreditationHandler",
	HandlerType: (*AdminAccreditationHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _AdminAccreditationHandler_GetList_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AdminAccreditationHandler_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AdminAccreditationHandler_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AdminAccreditationHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/accreditation.proto",
}
