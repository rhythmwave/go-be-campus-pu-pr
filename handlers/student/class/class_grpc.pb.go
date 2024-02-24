// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: student/class.proto

package class

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

// StudentClassHandlerClient is the client API for StudentClassHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentClassHandlerClient interface {
	GetOfferedClassList(ctx context.Context, in *GetOfferedClassListRequest, opts ...grpc.CallOption) (*GetOfferedClassListResponse, error)
	GetOfferedSchedule(ctx context.Context, in *GetOfferedScheduleRequest, opts ...grpc.CallOption) (*GetOfferedScheduleResponse, error)
	GetTakenClass(ctx context.Context, in *GetTakenClassRequest, opts ...grpc.CallOption) (*GetTakenClassResponse, error)
}

type studentClassHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentClassHandlerClient(cc grpc.ClientConnInterface) StudentClassHandlerClient {
	return &studentClassHandlerClient{cc}
}

func (c *studentClassHandlerClient) GetOfferedClassList(ctx context.Context, in *GetOfferedClassListRequest, opts ...grpc.CallOption) (*GetOfferedClassListResponse, error) {
	out := new(GetOfferedClassListResponse)
	err := c.cc.Invoke(ctx, "/student_class.StudentClassHandler/GetOfferedClassList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClassHandlerClient) GetOfferedSchedule(ctx context.Context, in *GetOfferedScheduleRequest, opts ...grpc.CallOption) (*GetOfferedScheduleResponse, error) {
	out := new(GetOfferedScheduleResponse)
	err := c.cc.Invoke(ctx, "/student_class.StudentClassHandler/GetOfferedSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClassHandlerClient) GetTakenClass(ctx context.Context, in *GetTakenClassRequest, opts ...grpc.CallOption) (*GetTakenClassResponse, error) {
	out := new(GetTakenClassResponse)
	err := c.cc.Invoke(ctx, "/student_class.StudentClassHandler/GetTakenClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentClassHandlerServer is the server API for StudentClassHandler service.
// All implementations must embed UnimplementedStudentClassHandlerServer
// for forward compatibility
type StudentClassHandlerServer interface {
	GetOfferedClassList(context.Context, *GetOfferedClassListRequest) (*GetOfferedClassListResponse, error)
	GetOfferedSchedule(context.Context, *GetOfferedScheduleRequest) (*GetOfferedScheduleResponse, error)
	GetTakenClass(context.Context, *GetTakenClassRequest) (*GetTakenClassResponse, error)
	mustEmbedUnimplementedStudentClassHandlerServer()
}

// UnimplementedStudentClassHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedStudentClassHandlerServer struct {
}

func (UnimplementedStudentClassHandlerServer) GetOfferedClassList(context.Context, *GetOfferedClassListRequest) (*GetOfferedClassListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOfferedClassList not implemented")
}
func (UnimplementedStudentClassHandlerServer) GetOfferedSchedule(context.Context, *GetOfferedScheduleRequest) (*GetOfferedScheduleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOfferedSchedule not implemented")
}
func (UnimplementedStudentClassHandlerServer) GetTakenClass(context.Context, *GetTakenClassRequest) (*GetTakenClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTakenClass not implemented")
}
func (UnimplementedStudentClassHandlerServer) mustEmbedUnimplementedStudentClassHandlerServer() {}

// UnsafeStudentClassHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentClassHandlerServer will
// result in compilation errors.
type UnsafeStudentClassHandlerServer interface {
	mustEmbedUnimplementedStudentClassHandlerServer()
}

func RegisterStudentClassHandlerServer(s grpc.ServiceRegistrar, srv StudentClassHandlerServer) {
	s.RegisterService(&StudentClassHandler_ServiceDesc, srv)
}

func _StudentClassHandler_GetOfferedClassList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOfferedClassListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassHandlerServer).GetOfferedClassList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student_class.StudentClassHandler/GetOfferedClassList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassHandlerServer).GetOfferedClassList(ctx, req.(*GetOfferedClassListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentClassHandler_GetOfferedSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOfferedScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassHandlerServer).GetOfferedSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student_class.StudentClassHandler/GetOfferedSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassHandlerServer).GetOfferedSchedule(ctx, req.(*GetOfferedScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentClassHandler_GetTakenClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTakenClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassHandlerServer).GetTakenClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student_class.StudentClassHandler/GetTakenClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassHandlerServer).GetTakenClass(ctx, req.(*GetTakenClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentClassHandler_ServiceDesc is the grpc.ServiceDesc for StudentClassHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentClassHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student_class.StudentClassHandler",
	HandlerType: (*StudentClassHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOfferedClassList",
			Handler:    _StudentClassHandler_GetOfferedClassList_Handler,
		},
		{
			MethodName: "GetOfferedSchedule",
			Handler:    _StudentClassHandler_GetOfferedSchedule_Handler,
		},
		{
			MethodName: "GetTakenClass",
			Handler:    _StudentClassHandler_GetTakenClass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student/class.proto",
}