// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: lecturer/academic_guidance.proto

package academic_guidance

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

// LecturerAcademicGuidanceHandlerClient is the client API for LecturerAcademicGuidanceHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LecturerAcademicGuidanceHandlerClient interface {
	GetListStudent(ctx context.Context, in *GetListStudentRequest, opts ...grpc.CallOption) (*GetListStudentResponse, error)
	GetSessionList(ctx context.Context, in *GetSessionListRequest, opts ...grpc.CallOption) (*GetSessionListResponse, error)
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	UpdateSession(ctx context.Context, in *UpdateSessionRequest, opts ...grpc.CallOption) (*UpdateSessionResponse, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error)
}

type lecturerAcademicGuidanceHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewLecturerAcademicGuidanceHandlerClient(cc grpc.ClientConnInterface) LecturerAcademicGuidanceHandlerClient {
	return &lecturerAcademicGuidanceHandlerClient{cc}
}

func (c *lecturerAcademicGuidanceHandlerClient) GetListStudent(ctx context.Context, in *GetListStudentRequest, opts ...grpc.CallOption) (*GetListStudentResponse, error) {
	out := new(GetListStudentResponse)
	err := c.cc.Invoke(ctx, "/lecturer_academic_guidance.LecturerAcademicGuidanceHandler/GetListStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerAcademicGuidanceHandlerClient) GetSessionList(ctx context.Context, in *GetSessionListRequest, opts ...grpc.CallOption) (*GetSessionListResponse, error) {
	out := new(GetSessionListResponse)
	err := c.cc.Invoke(ctx, "/lecturer_academic_guidance.LecturerAcademicGuidanceHandler/GetSessionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerAcademicGuidanceHandlerClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, "/lecturer_academic_guidance.LecturerAcademicGuidanceHandler/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerAcademicGuidanceHandlerClient) UpdateSession(ctx context.Context, in *UpdateSessionRequest, opts ...grpc.CallOption) (*UpdateSessionResponse, error) {
	out := new(UpdateSessionResponse)
	err := c.cc.Invoke(ctx, "/lecturer_academic_guidance.LecturerAcademicGuidanceHandler/UpdateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lecturerAcademicGuidanceHandlerClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error) {
	out := new(DeleteSessionResponse)
	err := c.cc.Invoke(ctx, "/lecturer_academic_guidance.LecturerAcademicGuidanceHandler/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LecturerAcademicGuidanceHandlerServer is the server API for LecturerAcademicGuidanceHandler service.
// All implementations must embed UnimplementedLecturerAcademicGuidanceHandlerServer
// for forward compatibility
type LecturerAcademicGuidanceHandlerServer interface {
	GetListStudent(context.Context, *GetListStudentRequest) (*GetListStudentResponse, error)
	GetSessionList(context.Context, *GetSessionListRequest) (*GetSessionListResponse, error)
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	UpdateSession(context.Context, *UpdateSessionRequest) (*UpdateSessionResponse, error)
	DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionResponse, error)
	mustEmbedUnimplementedLecturerAcademicGuidanceHandlerServer()
}

// UnimplementedLecturerAcademicGuidanceHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedLecturerAcademicGuidanceHandlerServer struct {
}

func (UnimplementedLecturerAcademicGuidanceHandlerServer) GetListStudent(context.Context, *GetListStudentRequest) (*GetListStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListStudent not implemented")
}
func (UnimplementedLecturerAcademicGuidanceHandlerServer) GetSessionList(context.Context, *GetSessionListRequest) (*GetSessionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionList not implemented")
}
func (UnimplementedLecturerAcademicGuidanceHandlerServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedLecturerAcademicGuidanceHandlerServer) UpdateSession(context.Context, *UpdateSessionRequest) (*UpdateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSession not implemented")
}
func (UnimplementedLecturerAcademicGuidanceHandlerServer) DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedLecturerAcademicGuidanceHandlerServer) mustEmbedUnimplementedLecturerAcademicGuidanceHandlerServer() {
}

// UnsafeLecturerAcademicGuidanceHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LecturerAcademicGuidanceHandlerServer will
// result in compilation errors.
type UnsafeLecturerAcademicGuidanceHandlerServer interface {
	mustEmbedUnimplementedLecturerAcademicGuidanceHandlerServer()
}

func RegisterLecturerAcademicGuidanceHandlerServer(s grpc.ServiceRegistrar, srv LecturerAcademicGuidanceHandlerServer) {
	s.RegisterService(&LecturerAcademicGuidanceHandler_ServiceDesc, srv)
}

func _LecturerAcademicGuidanceHandler_GetListStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerAcademicGuidanceHandlerServer).GetListStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_academic_guidance.LecturerAcademicGuidanceHandler/GetListStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerAcademicGuidanceHandlerServer).GetListStudent(ctx, req.(*GetListStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerAcademicGuidanceHandler_GetSessionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerAcademicGuidanceHandlerServer).GetSessionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_academic_guidance.LecturerAcademicGuidanceHandler/GetSessionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerAcademicGuidanceHandlerServer).GetSessionList(ctx, req.(*GetSessionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerAcademicGuidanceHandler_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerAcademicGuidanceHandlerServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_academic_guidance.LecturerAcademicGuidanceHandler/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerAcademicGuidanceHandlerServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerAcademicGuidanceHandler_UpdateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerAcademicGuidanceHandlerServer).UpdateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_academic_guidance.LecturerAcademicGuidanceHandler/UpdateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerAcademicGuidanceHandlerServer).UpdateSession(ctx, req.(*UpdateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LecturerAcademicGuidanceHandler_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerAcademicGuidanceHandlerServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_academic_guidance.LecturerAcademicGuidanceHandler/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerAcademicGuidanceHandlerServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LecturerAcademicGuidanceHandler_ServiceDesc is the grpc.ServiceDesc for LecturerAcademicGuidanceHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LecturerAcademicGuidanceHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lecturer_academic_guidance.LecturerAcademicGuidanceHandler",
	HandlerType: (*LecturerAcademicGuidanceHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListStudent",
			Handler:    _LecturerAcademicGuidanceHandler_GetListStudent_Handler,
		},
		{
			MethodName: "GetSessionList",
			Handler:    _LecturerAcademicGuidanceHandler_GetSessionList_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _LecturerAcademicGuidanceHandler_CreateSession_Handler,
		},
		{
			MethodName: "UpdateSession",
			Handler:    _LecturerAcademicGuidanceHandler_UpdateSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _LecturerAcademicGuidanceHandler_DeleteSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lecturer/academic_guidance.proto",
}
