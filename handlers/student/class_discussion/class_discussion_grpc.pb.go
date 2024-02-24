// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: student/class_discussion.proto

package class_discussion

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

// StudentClassDiscussionHandlerClient is the client API for StudentClassDiscussionHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentClassDiscussionHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error)
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
}

type studentClassDiscussionHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentClassDiscussionHandlerClient(cc grpc.ClientConnInterface) StudentClassDiscussionHandlerClient {
	return &studentClassDiscussionHandlerClient{cc}
}

func (c *studentClassDiscussionHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/student_class_discussion.StudentClassDiscussionHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClassDiscussionHandlerClient) GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error) {
	out := new(GetCommentResponse)
	err := c.cc.Invoke(ctx, "/student_class_discussion.StudentClassDiscussionHandler/GetComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClassDiscussionHandlerClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, "/student_class_discussion.StudentClassDiscussionHandler/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClassDiscussionHandlerClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, "/student_class_discussion.StudentClassDiscussionHandler/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentClassDiscussionHandlerServer is the server API for StudentClassDiscussionHandler service.
// All implementations must embed UnimplementedStudentClassDiscussionHandlerServer
// for forward compatibility
type StudentClassDiscussionHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	GetComment(context.Context, *GetCommentRequest) (*GetCommentResponse, error)
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	mustEmbedUnimplementedStudentClassDiscussionHandlerServer()
}

// UnimplementedStudentClassDiscussionHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedStudentClassDiscussionHandlerServer struct {
}

func (UnimplementedStudentClassDiscussionHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedStudentClassDiscussionHandlerServer) GetComment(context.Context, *GetCommentRequest) (*GetCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedStudentClassDiscussionHandlerServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedStudentClassDiscussionHandlerServer) DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedStudentClassDiscussionHandlerServer) mustEmbedUnimplementedStudentClassDiscussionHandlerServer() {
}

// UnsafeStudentClassDiscussionHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentClassDiscussionHandlerServer will
// result in compilation errors.
type UnsafeStudentClassDiscussionHandlerServer interface {
	mustEmbedUnimplementedStudentClassDiscussionHandlerServer()
}

func RegisterStudentClassDiscussionHandlerServer(s grpc.ServiceRegistrar, srv StudentClassDiscussionHandlerServer) {
	s.RegisterService(&StudentClassDiscussionHandler_ServiceDesc, srv)
}

func _StudentClassDiscussionHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassDiscussionHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student_class_discussion.StudentClassDiscussionHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassDiscussionHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentClassDiscussionHandler_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassDiscussionHandlerServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student_class_discussion.StudentClassDiscussionHandler/GetComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassDiscussionHandlerServer).GetComment(ctx, req.(*GetCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentClassDiscussionHandler_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassDiscussionHandlerServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student_class_discussion.StudentClassDiscussionHandler/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassDiscussionHandlerServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentClassDiscussionHandler_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentClassDiscussionHandlerServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student_class_discussion.StudentClassDiscussionHandler/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentClassDiscussionHandlerServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentClassDiscussionHandler_ServiceDesc is the grpc.ServiceDesc for StudentClassDiscussionHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentClassDiscussionHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student_class_discussion.StudentClassDiscussionHandler",
	HandlerType: (*StudentClassDiscussionHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _StudentClassDiscussionHandler_GetList_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _StudentClassDiscussionHandler_GetComment_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _StudentClassDiscussionHandler_CreateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _StudentClassDiscussionHandler_DeleteComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student/class_discussion.proto",
}
