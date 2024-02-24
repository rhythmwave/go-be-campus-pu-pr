// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: lecturer/class_discussion.proto

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

// LecturerClassDiscussionHandlerClient is the client API for LecturerClassDiscussionHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LecturerClassDiscussionHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
}

type lecturerClassDiscussionHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewLecturerClassDiscussionHandlerClient(cc grpc.ClientConnInterface) LecturerClassDiscussionHandlerClient {
	return &lecturerClassDiscussionHandlerClient{cc}
}

func (c *lecturerClassDiscussionHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/lecturer_class_discussion.LecturerClassDiscussionHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerClassDiscussionHandlerClient) GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentResponse, error) {
	out := new(GetCommentResponse)
	err := c.cc.Invoke(ctx, "/lecturer_class_discussion.LecturerClassDiscussionHandler/GetComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerClassDiscussionHandlerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/lecturer_class_discussion.LecturerClassDiscussionHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerClassDiscussionHandlerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/lecturer_class_discussion.LecturerClassDiscussionHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerClassDiscussionHandlerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/lecturer_class_discussion.LecturerClassDiscussionHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerClassDiscussionHandlerClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, "/lecturer_class_discussion.LecturerClassDiscussionHandler/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerClassDiscussionHandlerClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, "/lecturer_class_discussion.LecturerClassDiscussionHandler/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LecturerClassDiscussionHandlerServer is the server API for LecturerClassDiscussionHandler service.
// All implementations must embed UnimplementedLecturerClassDiscussionHandlerServer
// for forward compatibility
type LecturerClassDiscussionHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	GetComment(context.Context, *GetCommentRequest) (*GetCommentResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	mustEmbedUnimplementedLecturerClassDiscussionHandlerServer()
}

// UnimplementedLecturerClassDiscussionHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedLecturerClassDiscussionHandlerServer struct {
}

func (UnimplementedLecturerClassDiscussionHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedLecturerClassDiscussionHandlerServer) GetComment(context.Context, *GetCommentRequest) (*GetCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedLecturerClassDiscussionHandlerServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLecturerClassDiscussionHandlerServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLecturerClassDiscussionHandlerServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLecturerClassDiscussionHandlerServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedLecturerClassDiscussionHandlerServer) DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedLecturerClassDiscussionHandlerServer) mustEmbedUnimplementedLecturerClassDiscussionHandlerServer() {
}

// UnsafeLecturerClassDiscussionHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LecturerClassDiscussionHandlerServer will
// result in compilation errors.
type UnsafeLecturerClassDiscussionHandlerServer interface {
	mustEmbedUnimplementedLecturerClassDiscussionHandlerServer()
}

func RegisterLecturerClassDiscussionHandlerServer(s grpc.ServiceRegistrar, srv LecturerClassDiscussionHandlerServer) {
	s.RegisterService(&LecturerClassDiscussionHandler_ServiceDesc, srv)
}

func _LecturerClassDiscussionHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerClassDiscussionHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_class_discussion.LecturerClassDiscussionHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerClassDiscussionHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerClassDiscussionHandler_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerClassDiscussionHandlerServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_class_discussion.LecturerClassDiscussionHandler/GetComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerClassDiscussionHandlerServer).GetComment(ctx, req.(*GetCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerClassDiscussionHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerClassDiscussionHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_class_discussion.LecturerClassDiscussionHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerClassDiscussionHandlerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerClassDiscussionHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerClassDiscussionHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_class_discussion.LecturerClassDiscussionHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerClassDiscussionHandlerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerClassDiscussionHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerClassDiscussionHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_class_discussion.LecturerClassDiscussionHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerClassDiscussionHandlerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerClassDiscussionHandler_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerClassDiscussionHandlerServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_class_discussion.LecturerClassDiscussionHandler/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerClassDiscussionHandlerServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerClassDiscussionHandler_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerClassDiscussionHandlerServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_class_discussion.LecturerClassDiscussionHandler/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerClassDiscussionHandlerServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LecturerClassDiscussionHandler_ServiceDesc is the grpc.ServiceDesc for LecturerClassDiscussionHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LecturerClassDiscussionHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lecturer_class_discussion.LecturerClassDiscussionHandler",
	HandlerType: (*LecturerClassDiscussionHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _LecturerClassDiscussionHandler_GetList_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _LecturerClassDiscussionHandler_GetComment_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LecturerClassDiscussionHandler_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LecturerClassDiscussionHandler_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LecturerClassDiscussionHandler_Delete_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _LecturerClassDiscussionHandler_CreateComment_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _LecturerClassDiscussionHandler_DeleteComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lecturer/class_discussion.proto",
}