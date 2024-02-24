// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: admin/class_work.proto

package class_work

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

// AdminClassWorkHandlerClient is the client API for AdminClassWorkHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClassWorkHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*GetSubmissionResponse, error)
}

type adminClassWorkHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClassWorkHandlerClient(cc grpc.ClientConnInterface) AdminClassWorkHandlerClient {
	return &adminClassWorkHandlerClient{cc}
}

func (c *adminClassWorkHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/admin_class_work.AdminClassWorkHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClassWorkHandlerClient) GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*GetSubmissionResponse, error) {
	out := new(GetSubmissionResponse)
	err := c.cc.Invoke(ctx, "/admin_class_work.AdminClassWorkHandler/GetSubmission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminClassWorkHandlerServer is the server API for AdminClassWorkHandler service.
// All implementations must embed UnimplementedAdminClassWorkHandlerServer
// for forward compatibility
type AdminClassWorkHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	GetSubmission(context.Context, *GetSubmissionRequest) (*GetSubmissionResponse, error)
	mustEmbedUnimplementedAdminClassWorkHandlerServer()
}

// UnimplementedAdminClassWorkHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAdminClassWorkHandlerServer struct {
}

func (UnimplementedAdminClassWorkHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdminClassWorkHandlerServer) GetSubmission(context.Context, *GetSubmissionRequest) (*GetSubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmission not implemented")
}
func (UnimplementedAdminClassWorkHandlerServer) mustEmbedUnimplementedAdminClassWorkHandlerServer() {}

// UnsafeAdminClassWorkHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminClassWorkHandlerServer will
// result in compilation errors.
type UnsafeAdminClassWorkHandlerServer interface {
	mustEmbedUnimplementedAdminClassWorkHandlerServer()
}

func RegisterAdminClassWorkHandlerServer(s grpc.ServiceRegistrar, srv AdminClassWorkHandlerServer) {
	s.RegisterService(&AdminClassWorkHandler_ServiceDesc, srv)
}

func _AdminClassWorkHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminClassWorkHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_class_work.AdminClassWorkHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminClassWorkHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminClassWorkHandler_GetSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminClassWorkHandlerServer).GetSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_class_work.AdminClassWorkHandler/GetSubmission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminClassWorkHandlerServer).GetSubmission(ctx, req.(*GetSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminClassWorkHandler_ServiceDesc is the grpc.ServiceDesc for AdminClassWorkHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminClassWorkHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_class_work.AdminClassWorkHandler",
	HandlerType: (*AdminClassWorkHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _AdminClassWorkHandler_GetList_Handler,
		},
		{
			MethodName: "GetSubmission",
			Handler:    _AdminClassWorkHandler_GetSubmission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/class_work.proto",
}
