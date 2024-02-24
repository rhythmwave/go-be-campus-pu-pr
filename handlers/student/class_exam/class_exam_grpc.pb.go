// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: student/class_exam.proto

package class_exam

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

// StudentClassExamHandlerClient is the client API for StudentClassExamHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentClassExamHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
}

type studentClassExamHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentClassExamHandlerClient(cc grpc.ClientConnInterface) StudentClassExamHandlerClient {
	return &studentClassExamHandlerClient{cc}
}

func (c *studentClassExamHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/student_class_exam.StudentClassExamHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClassExamHandlerClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, "/student_class_exam.StudentClassExamHandler/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentClassExamHandlerServer is the server API for StudentClassExamHandler service.
// All implementations must embed UnimplementedStudentClassExamHandlerServer
// for forward compatibility
type StudentClassExamHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
	mustEmbedUnimplementedStudentClassExamHandlerServer()
}

// UnimplementedStudentClassExamHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedStudentClassExamHandlerServer struct {
}

func (UnimplementedStudentClassExamHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedStudentClassExamHandlerServer) Submit(context.Context, *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedStudentClassExamHandlerServer) mustEmbedUnimplementedStudentClassExamHandlerServer() {
}

// UnsafeStudentClassExamHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentClassExamHandlerServer will
// result in compilation errors.
type UnsafeStudentClassExamHandlerServer interface {
	mustEmbedUnimplementedStudentClassExamHandlerServer()
}

func RegisterStudentClassExamHandlerServer(s grpc.ServiceRegistrar, srv StudentClassExamHandlerServer) {
	s.RegisterService(&StudentClassExamHandler_ServiceDesc, srv)
}

func _StudentClassExamHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassExamHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student_class_exam.StudentClassExamHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassExamHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentClassExamHandler_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassExamHandlerServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student_class_exam.StudentClassExamHandler/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassExamHandlerServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentClassExamHandler_ServiceDesc is the grpc.ServiceDesc for StudentClassExamHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentClassExamHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student_class_exam.StudentClassExamHandler",
	HandlerType: (*StudentClassExamHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _StudentClassExamHandler_GetList_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _StudentClassExamHandler_Submit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student/class_exam.proto",
}
