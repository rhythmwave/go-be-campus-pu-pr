// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: admin/yudicium.proto

package yudicium

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

// AdminYudiciumHandlerClient is the client API for AdminYudiciumHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminYudiciumHandlerClient interface {
	Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error)
	GetListStudent(ctx context.Context, in *GetListStudentRequest, opts ...grpc.CallOption) (*GetListStudentResponse, error)
	Do(ctx context.Context, in *DoRequest, opts ...grpc.CallOption) (*DoResponse, error)
}

type adminYudiciumHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminYudiciumHandlerClient(cc grpc.ClientConnInterface) AdminYudiciumHandlerClient {
	return &adminYudiciumHandlerClient{cc}
}

func (c *adminYudiciumHandlerClient) Apply(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error) {
	out := new(ApplyResponse)
	err := c.cc.Invoke(ctx, "/admin_yudicium.AdminYudiciumHandler/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminYudiciumHandlerClient) GetListStudent(ctx context.Context, in *GetListStudentRequest, opts ...grpc.CallOption) (*GetListStudentResponse, error) {
	out := new(GetListStudentResponse)
	err := c.cc.Invoke(ctx, "/admin_yudicium.AdminYudiciumHandler/GetListStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminYudiciumHandlerClient) Do(ctx context.Context, in *DoRequest, opts ...grpc.CallOption) (*DoResponse, error) {
	out := new(DoResponse)
	err := c.cc.Invoke(ctx, "/admin_yudicium.AdminYudiciumHandler/Do", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminYudiciumHandlerServer is the server API for AdminYudiciumHandler service.
// All implementations must embed UnimplementedAdminYudiciumHandlerServer
// for forward compatibility
type AdminYudiciumHandlerServer interface {
	Apply(context.Context, *ApplyRequest) (*ApplyResponse, error)
	GetListStudent(context.Context, *GetListStudentRequest) (*GetListStudentResponse, error)
	Do(context.Context, *DoRequest) (*DoResponse, error)
	mustEmbedUnimplementedAdminYudiciumHandlerServer()
}

// UnimplementedAdminYudiciumHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAdminYudiciumHandlerServer struct {
}

func (UnimplementedAdminYudiciumHandlerServer) Apply(context.Context, *ApplyRequest) (*ApplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Apply not implemented")
}
func (UnimplementedAdminYudiciumHandlerServer) GetListStudent(context.Context, *GetListStudentRequest) (*GetListStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListStudent not implemented")
}
func (UnimplementedAdminYudiciumHandlerServer) Do(context.Context, *DoRequest) (*DoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Do not implemented")
}
func (UnimplementedAdminYudiciumHandlerServer) mustEmbedUnimplementedAdminYudiciumHandlerServer() {}

// UnsafeAdminYudiciumHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminYudiciumHandlerServer will
// result in compilation errors.
type UnsafeAdminYudiciumHandlerServer interface {
	mustEmbedUnimplementedAdminYudiciumHandlerServer()
}

func RegisterAdminYudiciumHandlerServer(s grpc.ServiceRegistrar, srv AdminYudiciumHandlerServer) {
	s.RegisterService(&AdminYudiciumHandler_ServiceDesc, srv)
}

func _AdminYudiciumHandler_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminYudiciumHandlerServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_yudicium.AdminYudiciumHandler/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminYudiciumHandlerServer).Apply(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminYudiciumHandler_GetListStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminYudiciumHandlerServer).GetListStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_yudicium.AdminYudiciumHandler/GetListStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminYudiciumHandlerServer).GetListStudent(ctx, req.(*GetListStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminYudiciumHandler_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminYudiciumHandlerServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_yudicium.AdminYudiciumHandler/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminYudiciumHandlerServer).Do(ctx, req.(*DoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminYudiciumHandler_ServiceDesc is the grpc.ServiceDesc for AdminYudiciumHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminYudiciumHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_yudicium.AdminYudiciumHandler",
	HandlerType: (*AdminYudiciumHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Apply",
			Handler:    _AdminYudiciumHandler_Apply_Handler,
		},
		{
			MethodName: "GetListStudent",
			Handler:    _AdminYudiciumHandler_GetListStudent_Handler,
		},
		{
			MethodName: "Do",
			Handler:    _AdminYudiciumHandler_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/yudicium.proto",
}
