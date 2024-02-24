// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: root/lecturer_student_activity_log.proto

package lecturer_student_activity_log

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

// RootLecturerStudentActivityLogHandlerClient is the client API for RootLecturerStudentActivityLogHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RootLecturerStudentActivityLogHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
}

type rootLecturerStudentActivityLogHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewRootLecturerStudentActivityLogHandlerClient(cc grpc.ClientConnInterface) RootLecturerStudentActivityLogHandlerClient {
	return &rootLecturerStudentActivityLogHandlerClient{cc}
}

func (c *rootLecturerStudentActivityLogHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/root_lecturer_student_activity_log.RootLecturerStudentActivityLogHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RootLecturerStudentActivityLogHandlerServer is the server API for RootLecturerStudentActivityLogHandler service.
// All implementations must embed UnimplementedRootLecturerStudentActivityLogHandlerServer
// for forward compatibility
type RootLecturerStudentActivityLogHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	mustEmbedUnimplementedRootLecturerStudentActivityLogHandlerServer()
}

// UnimplementedRootLecturerStudentActivityLogHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedRootLecturerStudentActivityLogHandlerServer struct {
}

func (UnimplementedRootLecturerStudentActivityLogHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedRootLecturerStudentActivityLogHandlerServer) mustEmbedUnimplementedRootLecturerStudentActivityLogHandlerServer() {
}

// UnsafeRootLecturerStudentActivityLogHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RootLecturerStudentActivityLogHandlerServer will
// result in compilation errors.
type UnsafeRootLecturerStudentActivityLogHandlerServer interface {
	mustEmbedUnimplementedRootLecturerStudentActivityLogHandlerServer()
}

func RegisterRootLecturerStudentActivityLogHandlerServer(s grpc.ServiceRegistrar, srv RootLecturerStudentActivityLogHandlerServer) {
	s.RegisterService(&RootLecturerStudentActivityLogHandler_ServiceDesc, srv)
}

func _RootLecturerStudentActivityLogHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RootLecturerStudentActivityLogHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/root_lecturer_student_activity_log.RootLecturerStudentActivityLogHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RootLecturerStudentActivityLogHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RootLecturerStudentActivityLogHandler_ServiceDesc is the grpc.ServiceDesc for RootLecturerStudentActivityLogHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RootLecturerStudentActivityLogHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "root_lecturer_student_activity_log.RootLecturerStudentActivityLogHandler",
	HandlerType: (*RootLecturerStudentActivityLogHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _RootLecturerStudentActivityLogHandler_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "root/lecturer_student_activity_log.proto",
}