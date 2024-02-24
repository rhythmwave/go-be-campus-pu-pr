// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: admin/academic_guidance.proto

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

// AdminAcademicGuidanceHandlerClient is the client API for AdminAcademicGuidanceHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminAcademicGuidanceHandlerClient interface {
	GetListStudent(ctx context.Context, in *GetListStudentRequest, opts ...grpc.CallOption) (*GetListStudentResponse, error)
	Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*UpsertResponse, error)
	UpsertDecision(ctx context.Context, in *UpsertDecisionRequest, opts ...grpc.CallOption) (*UpsertDecisionResponse, error)
	GetSessionList(ctx context.Context, in *GetSessionListRequest, opts ...grpc.CallOption) (*GetSessionListResponse, error)
	CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error)
	UpdateSession(ctx context.Context, in *UpdateSessionRequest, opts ...grpc.CallOption) (*UpdateSessionResponse, error)
	DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error)
}

type adminAcademicGuidanceHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminAcademicGuidanceHandlerClient(cc grpc.ClientConnInterface) AdminAcademicGuidanceHandlerClient {
	return &adminAcademicGuidanceHandlerClient{cc}
}

func (c *adminAcademicGuidanceHandlerClient) GetListStudent(ctx context.Context, in *GetListStudentRequest, opts ...grpc.CallOption) (*GetListStudentResponse, error) {
	out := new(GetListStudentResponse)
	err := c.cc.Invoke(ctx, "/admin_academic_guidance.AdminAcademicGuidanceHandler/GetListStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAcademicGuidanceHandlerClient) Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*UpsertResponse, error) {
	out := new(UpsertResponse)
	err := c.cc.Invoke(ctx, "/admin_academic_guidance.AdminAcademicGuidanceHandler/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAcademicGuidanceHandlerClient) UpsertDecision(ctx context.Context, in *UpsertDecisionRequest, opts ...grpc.CallOption) (*UpsertDecisionResponse, error) {
	out := new(UpsertDecisionResponse)
	err := c.cc.Invoke(ctx, "/admin_academic_guidance.AdminAcademicGuidanceHandler/UpsertDecision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAcademicGuidanceHandlerClient) GetSessionList(ctx context.Context, in *GetSessionListRequest, opts ...grpc.CallOption) (*GetSessionListResponse, error) {
	out := new(GetSessionListResponse)
	err := c.cc.Invoke(ctx, "/admin_academic_guidance.AdminAcademicGuidanceHandler/GetSessionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAcademicGuidanceHandlerClient) CreateSession(ctx context.Context, in *CreateSessionRequest, opts ...grpc.CallOption) (*CreateSessionResponse, error) {
	out := new(CreateSessionResponse)
	err := c.cc.Invoke(ctx, "/admin_academic_guidance.AdminAcademicGuidanceHandler/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAcademicGuidanceHandlerClient) UpdateSession(ctx context.Context, in *UpdateSessionRequest, opts ...grpc.CallOption) (*UpdateSessionResponse, error) {
	out := new(UpdateSessionResponse)
	err := c.cc.Invoke(ctx, "/admin_academic_guidance.AdminAcademicGuidanceHandler/UpdateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminAcademicGuidanceHandlerClient) DeleteSession(ctx context.Context, in *DeleteSessionRequest, opts ...grpc.CallOption) (*DeleteSessionResponse, error) {
	out := new(DeleteSessionResponse)
	err := c.cc.Invoke(ctx, "/admin_academic_guidance.AdminAcademicGuidanceHandler/DeleteSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminAcademicGuidanceHandlerServer is the server API for AdminAcademicGuidanceHandler service.
// All implementations must embed UnimplementedAdminAcademicGuidanceHandlerServer
// for forward compatibility
type AdminAcademicGuidanceHandlerServer interface {
	GetListStudent(context.Context, *GetListStudentRequest) (*GetListStudentResponse, error)
	Upsert(context.Context, *UpsertRequest) (*UpsertResponse, error)
	UpsertDecision(context.Context, *UpsertDecisionRequest) (*UpsertDecisionResponse, error)
	GetSessionList(context.Context, *GetSessionListRequest) (*GetSessionListResponse, error)
	CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error)
	UpdateSession(context.Context, *UpdateSessionRequest) (*UpdateSessionResponse, error)
	DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionResponse, error)
	mustEmbedUnimplementedAdminAcademicGuidanceHandlerServer()
}

// UnimplementedAdminAcademicGuidanceHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAdminAcademicGuidanceHandlerServer struct {
}

func (UnimplementedAdminAcademicGuidanceHandlerServer) GetListStudent(context.Context, *GetListStudentRequest) (*GetListStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListStudent not implemented")
}
func (UnimplementedAdminAcademicGuidanceHandlerServer) Upsert(context.Context, *UpsertRequest) (*UpsertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedAdminAcademicGuidanceHandlerServer) UpsertDecision(context.Context, *UpsertDecisionRequest) (*UpsertDecisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertDecision not implemented")
}
func (UnimplementedAdminAcademicGuidanceHandlerServer) GetSessionList(context.Context, *GetSessionListRequest) (*GetSessionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessionList not implemented")
}
func (UnimplementedAdminAcademicGuidanceHandlerServer) CreateSession(context.Context, *CreateSessionRequest) (*CreateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedAdminAcademicGuidanceHandlerServer) UpdateSession(context.Context, *UpdateSessionRequest) (*UpdateSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSession not implemented")
}
func (UnimplementedAdminAcademicGuidanceHandlerServer) DeleteSession(context.Context, *DeleteSessionRequest) (*DeleteSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSession not implemented")
}
func (UnimplementedAdminAcademicGuidanceHandlerServer) mustEmbedUnimplementedAdminAcademicGuidanceHandlerServer() {
}

// UnsafeAdminAcademicGuidanceHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminAcademicGuidanceHandlerServer will
// result in compilation errors.
type UnsafeAdminAcademicGuidanceHandlerServer interface {
	mustEmbedUnimplementedAdminAcademicGuidanceHandlerServer()
}

func RegisterAdminAcademicGuidanceHandlerServer(s grpc.ServiceRegistrar, srv AdminAcademicGuidanceHandlerServer) {
	s.RegisterService(&AdminAcademicGuidanceHandler_ServiceDesc, srv)
}

func _AdminAcademicGuidanceHandler_GetListStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAcademicGuidanceHandlerServer).GetListStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_academic_guidance.AdminAcademicGuidanceHandler/GetListStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAcademicGuidanceHandlerServer).GetListStudent(ctx, req.(*GetListStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAcademicGuidanceHandler_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAcademicGuidanceHandlerServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_academic_guidance.AdminAcademicGuidanceHandler/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAcademicGuidanceHandlerServer).Upsert(ctx, req.(*UpsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAcademicGuidanceHandler_UpsertDecision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertDecisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAcademicGuidanceHandlerServer).UpsertDecision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_academic_guidance.AdminAcademicGuidanceHandler/UpsertDecision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAcademicGuidanceHandlerServer).UpsertDecision(ctx, req.(*UpsertDecisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAcademicGuidanceHandler_GetSessionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAcademicGuidanceHandlerServer).GetSessionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_academic_guidance.AdminAcademicGuidanceHandler/GetSessionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAcademicGuidanceHandlerServer).GetSessionList(ctx, req.(*GetSessionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAcademicGuidanceHandler_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAcademicGuidanceHandlerServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_academic_guidance.AdminAcademicGuidanceHandler/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAcademicGuidanceHandlerServer).CreateSession(ctx, req.(*CreateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAcademicGuidanceHandler_UpdateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAcademicGuidanceHandlerServer).UpdateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_academic_guidance.AdminAcademicGuidanceHandler/UpdateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAcademicGuidanceHandlerServer).UpdateSession(ctx, req.(*UpdateSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminAcademicGuidanceHandler_DeleteSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminAcademicGuidanceHandlerServer).DeleteSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_academic_guidance.AdminAcademicGuidanceHandler/DeleteSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminAcademicGuidanceHandlerServer).DeleteSession(ctx, req.(*DeleteSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminAcademicGuidanceHandler_ServiceDesc is the grpc.ServiceDesc for AdminAcademicGuidanceHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminAcademicGuidanceHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_academic_guidance.AdminAcademicGuidanceHandler",
	HandlerType: (*AdminAcademicGuidanceHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListStudent",
			Handler:    _AdminAcademicGuidanceHandler_GetListStudent_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _AdminAcademicGuidanceHandler_Upsert_Handler,
		},
		{
			MethodName: "UpsertDecision",
			Handler:    _AdminAcademicGuidanceHandler_UpsertDecision_Handler,
		},
		{
			MethodName: "GetSessionList",
			Handler:    _AdminAcademicGuidanceHandler_GetSessionList_Handler,
		},
		{
			MethodName: "CreateSession",
			Handler:    _AdminAcademicGuidanceHandler_CreateSession_Handler,
		},
		{
			MethodName: "UpdateSession",
			Handler:    _AdminAcademicGuidanceHandler_UpdateSession_Handler,
		},
		{
			MethodName: "DeleteSession",
			Handler:    _AdminAcademicGuidanceHandler_DeleteSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/academic_guidance.proto",
}
