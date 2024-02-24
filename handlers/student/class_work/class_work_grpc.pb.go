// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: student/class_work.proto

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

// StudentClassWorkHandlerClient is the client API for StudentClassWorkHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentClassWorkHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error)
}

type studentClassWorkHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentClassWorkHandlerClient(cc grpc.ClientConnInterface) StudentClassWorkHandlerClient {
	return &studentClassWorkHandlerClient{cc}
}

func (c *studentClassWorkHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/student_class_work.StudentClassWorkHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClassWorkHandlerClient) Submit(ctx context.Context, in *SubmitRequest, opts ...grpc.CallOption) (*SubmitResponse, error) {
	out := new(SubmitResponse)
	err := c.cc.Invoke(ctx, "/student_class_work.StudentClassWorkHandler/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentClassWorkHandlerServer is the server API for StudentClassWorkHandler service.
// All implementations must embed UnimplementedStudentClassWorkHandlerServer
// for forward compatibility
type StudentClassWorkHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	Submit(context.Context, *SubmitRequest) (*SubmitResponse, error)
	mustEmbedUnimplementedStudentClassWorkHandlerServer()
}

// UnimplementedStudentClassWorkHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedStudentClassWorkHandlerServer struct {
}

func (UnimplementedStudentClassWorkHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedStudentClassWorkHandlerServer) Submit(context.Context, *SubmitRequest) (*SubmitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedStudentClassWorkHandlerServer) mustEmbedUnimplementedStudentClassWorkHandlerServer() {
}

// UnsafeStudentClassWorkHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentClassWorkHandlerServer will
// result in compilation errors.
type UnsafeStudentClassWorkHandlerServer interface {
	mustEmbedUnimplementedStudentClassWorkHandlerServer()
}

func RegisterStudentClassWorkHandlerServer(s grpc.ServiceRegistrar, srv StudentClassWorkHandlerServer) {
	s.RegisterService(&StudentClassWorkHandler_ServiceDesc, srv)
}

func _StudentClassWorkHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassWorkHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student_class_work.StudentClassWorkHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassWorkHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentClassWorkHandler_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassWorkHandlerServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student_class_work.StudentClassWorkHandler/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassWorkHandlerServer).Submit(ctx, req.(*SubmitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentClassWorkHandler_ServiceDesc is the grpc.ServiceDesc for StudentClassWorkHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentClassWorkHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student_class_work.StudentClassWorkHandler",
	HandlerType: (*StudentClassWorkHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _StudentClassWorkHandler_GetList_Handler,
		},
		{
			MethodName: "Submit",
			Handler:    _StudentClassWorkHandler_Submit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student/class_work.proto",
}