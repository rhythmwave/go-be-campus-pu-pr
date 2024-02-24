// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: admin/lecturer_leave.proto

package lecturer_leave

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

// AdminLecturerLeaveHandlerClient is the client API for AdminLecturerLeaveHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminLecturerLeaveHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	End(ctx context.Context, in *EndRequest, opts ...grpc.CallOption) (*EndResponse, error)
}

type adminLecturerLeaveHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminLecturerLeaveHandlerClient(cc grpc.ClientConnInterface) AdminLecturerLeaveHandlerClient {
	return &adminLecturerLeaveHandlerClient{cc}
}

func (c *adminLecturerLeaveHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/admin_lecturer_leave.AdminLecturerLeaveHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminLecturerLeaveHandlerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/admin_lecturer_leave.AdminLecturerLeaveHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminLecturerLeaveHandlerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/admin_lecturer_leave.AdminLecturerLeaveHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminLecturerLeaveHandlerClient) End(ctx context.Context, in *EndRequest, opts ...grpc.CallOption) (*EndResponse, error) {
	out := new(EndResponse)
	err := c.cc.Invoke(ctx, "/admin_lecturer_leave.AdminLecturerLeaveHandler/End", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminLecturerLeaveHandlerServer is the server API for AdminLecturerLeaveHandler service.
// All implementations must embed UnimplementedAdminLecturerLeaveHandlerServer
// for forward compatibility
type AdminLecturerLeaveHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	End(context.Context, *EndRequest) (*EndResponse, error)
	mustEmbedUnimplementedAdminLecturerLeaveHandlerServer()
}

// UnimplementedAdminLecturerLeaveHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAdminLecturerLeaveHandlerServer struct {
}

func (UnimplementedAdminLecturerLeaveHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdminLecturerLeaveHandlerServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAdminLecturerLeaveHandlerServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAdminLecturerLeaveHandlerServer) End(context.Context, *EndRequest) (*EndResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method End not implemented")
}
func (UnimplementedAdminLecturerLeaveHandlerServer) mustEmbedUnimplementedAdminLecturerLeaveHandlerServer() {
}

// UnsafeAdminLecturerLeaveHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminLecturerLeaveHandlerServer will
// result in compilation errors.
type UnsafeAdminLecturerLeaveHandlerServer interface {
	mustEmbedUnimplementedAdminLecturerLeaveHandlerServer()
}

func RegisterAdminLecturerLeaveHandlerServer(s grpc.ServiceRegistrar, srv AdminLecturerLeaveHandlerServer) {
	s.RegisterService(&AdminLecturerLeaveHandler_ServiceDesc, srv)
}

func _AdminLecturerLeaveHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminLecturerLeaveHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_lecturer_leave.AdminLecturerLeaveHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminLecturerLeaveHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminLecturerLeaveHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminLecturerLeaveHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_lecturer_leave.AdminLecturerLeaveHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminLecturerLeaveHandlerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminLecturerLeaveHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminLecturerLeaveHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_lecturer_leave.AdminLecturerLeaveHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminLecturerLeaveHandlerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminLecturerLeaveHandler_End_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminLecturerLeaveHandlerServer).End(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_lecturer_leave.AdminLecturerLeaveHandler/End",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminLecturerLeaveHandlerServer).End(ctx, req.(*EndRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminLecturerLeaveHandler_ServiceDesc is the grpc.ServiceDesc for AdminLecturerLeaveHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminLecturerLeaveHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_lecturer_leave.AdminLecturerLeaveHandler",
	HandlerType: (*AdminLecturerLeaveHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _AdminLecturerLeaveHandler_GetList_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AdminLecturerLeaveHandler_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AdminLecturerLeaveHandler_Update_Handler,
		},
		{
			MethodName: "End",
			Handler:    _AdminLecturerLeaveHandler_End_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/lecturer_leave.proto",
}