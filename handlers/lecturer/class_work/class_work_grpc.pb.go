// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: lecturer/class_work.proto

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

// LecturerClassWorkHandlerClient is the client API for LecturerClassWorkHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LecturerClassWorkHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*GetSubmissionResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GradeSubmission(ctx context.Context, in *GradeSubmissionRequest, opts ...grpc.CallOption) (*GradeSubmissionResponse, error)
}

type lecturerClassWorkHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewLecturerClassWorkHandlerClient(cc grpc.ClientConnInterface) LecturerClassWorkHandlerClient {
	return &lecturerClassWorkHandlerClient{cc}
}

func (c *lecturerClassWorkHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/lecturer_class_work.LecturerClassWorkHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerClassWorkHandlerClient) GetSubmission(ctx context.Context, in *GetSubmissionRequest, opts ...grpc.CallOption) (*GetSubmissionResponse, error) {
	out := new(GetSubmissionResponse)
	err := c.cc.Invoke(ctx, "/lecturer_class_work.LecturerClassWorkHandler/GetSubmission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerClassWorkHandlerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/lecturer_class_work.LecturerClassWorkHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerClassWorkHandlerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/lecturer_class_work.LecturerClassWorkHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerClassWorkHandlerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/lecturer_class_work.LecturerClassWorkHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerClassWorkHandlerClient) GradeSubmission(ctx context.Context, in *GradeSubmissionRequest, opts ...grpc.CallOption) (*GradeSubmissionResponse, error) {
	out := new(GradeSubmissionResponse)
	err := c.cc.Invoke(ctx, "/lecturer_class_work.LecturerClassWorkHandler/GradeSubmission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LecturerClassWorkHandlerServer is the server API for LecturerClassWorkHandler service.
// All implementations must embed UnimplementedLecturerClassWorkHandlerServer
// for forward compatibility
type LecturerClassWorkHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	GetSubmission(context.Context, *GetSubmissionRequest) (*GetSubmissionResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	GradeSubmission(context.Context, *GradeSubmissionRequest) (*GradeSubmissionResponse, error)
	mustEmbedUnimplementedLecturerClassWorkHandlerServer()
}

// UnimplementedLecturerClassWorkHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedLecturerClassWorkHandlerServer struct {
}

func (UnimplementedLecturerClassWorkHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedLecturerClassWorkHandlerServer) GetSubmission(context.Context, *GetSubmissionRequest) (*GetSubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubmission not implemented")
}
func (UnimplementedLecturerClassWorkHandlerServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLecturerClassWorkHandlerServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLecturerClassWorkHandlerServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLecturerClassWorkHandlerServer) GradeSubmission(context.Context, *GradeSubmissionRequest) (*GradeSubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GradeSubmission not implemented")
}
func (UnimplementedLecturerClassWorkHandlerServer) mustEmbedUnimplementedLecturerClassWorkHandlerServer() {
}

// UnsafeLecturerClassWorkHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LecturerClassWorkHandlerServer will
// result in compilation errors.
type UnsafeLecturerClassWorkHandlerServer interface {
	mustEmbedUnimplementedLecturerClassWorkHandlerServer()
}

func RegisterLecturerClassWorkHandlerServer(s grpc.ServiceRegistrar, srv LecturerClassWorkHandlerServer) {
	s.RegisterService(&LecturerClassWorkHandler_ServiceDesc, srv)
}

func _LecturerClassWorkHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerClassWorkHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_class_work.LecturerClassWorkHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerClassWorkHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerClassWorkHandler_GetSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerClassWorkHandlerServer).GetSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_class_work.LecturerClassWorkHandler/GetSubmission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerClassWorkHandlerServer).GetSubmission(ctx, req.(*GetSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerClassWorkHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerClassWorkHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_class_work.LecturerClassWorkHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerClassWorkHandlerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerClassWorkHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerClassWorkHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_class_work.LecturerClassWorkHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerClassWorkHandlerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerClassWorkHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerClassWorkHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_class_work.LecturerClassWorkHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerClassWorkHandlerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerClassWorkHandler_GradeSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GradeSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerClassWorkHandlerServer).GradeSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_class_work.LecturerClassWorkHandler/GradeSubmission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerClassWorkHandlerServer).GradeSubmission(ctx, req.(*GradeSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LecturerClassWorkHandler_ServiceDesc is the grpc.ServiceDesc for LecturerClassWorkHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LecturerClassWorkHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lecturer_class_work.LecturerClassWorkHandler",
	HandlerType: (*LecturerClassWorkHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _LecturerClassWorkHandler_GetList_Handler,
		},
		{
			MethodName: "GetSubmission",
			Handler:    _LecturerClassWorkHandler_GetSubmission_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LecturerClassWorkHandler_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LecturerClassWorkHandler_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LecturerClassWorkHandler_Delete_Handler,
		},
		{
			MethodName: "GradeSubmission",
			Handler:    _LecturerClassWorkHandler_GradeSubmission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lecturer/class_work.proto",
}
