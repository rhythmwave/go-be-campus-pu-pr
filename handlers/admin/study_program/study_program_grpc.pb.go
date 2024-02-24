// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: admin/study_program.proto

package study_program

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

// AdminStudyProgramHandlerClient is the client API for AdminStudyProgramHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminStudyProgramHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc.CallOption) (*GetDetailResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	UpdateDegree(ctx context.Context, in *UpdateDegreeRequest, opts ...grpc.CallOption) (*UpdateDegreeResponse, error)
}

type adminStudyProgramHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminStudyProgramHandlerClient(cc grpc.ClientConnInterface) AdminStudyProgramHandlerClient {
	return &adminStudyProgramHandlerClient{cc}
}

func (c *adminStudyProgramHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/admin_study_program.AdminStudyProgramHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminStudyProgramHandlerClient) GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc.CallOption) (*GetDetailResponse, error) {
	out := new(GetDetailResponse)
	err := c.cc.Invoke(ctx, "/admin_study_program.AdminStudyProgramHandler/GetDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminStudyProgramHandlerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/admin_study_program.AdminStudyProgramHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminStudyProgramHandlerClient) UpdateDegree(ctx context.Context, in *UpdateDegreeRequest, opts ...grpc.CallOption) (*UpdateDegreeResponse, error) {
	out := new(UpdateDegreeResponse)
	err := c.cc.Invoke(ctx, "/admin_study_program.AdminStudyProgramHandler/UpdateDegree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminStudyProgramHandlerServer is the server API for AdminStudyProgramHandler service.
// All implementations must embed UnimplementedAdminStudyProgramHandlerServer
// for forward compatibility
type AdminStudyProgramHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	GetDetail(context.Context, *GetDetailRequest) (*GetDetailResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	UpdateDegree(context.Context, *UpdateDegreeRequest) (*UpdateDegreeResponse, error)
	mustEmbedUnimplementedAdminStudyProgramHandlerServer()
}

// UnimplementedAdminStudyProgramHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAdminStudyProgramHandlerServer struct {
}

func (UnimplementedAdminStudyProgramHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdminStudyProgramHandlerServer) GetDetail(context.Context, *GetDetailRequest) (*GetDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetail not implemented")
}
func (UnimplementedAdminStudyProgramHandlerServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAdminStudyProgramHandlerServer) UpdateDegree(context.Context, *UpdateDegreeRequest) (*UpdateDegreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDegree not implemented")
}
func (UnimplementedAdminStudyProgramHandlerServer) mustEmbedUnimplementedAdminStudyProgramHandlerServer() {
}

// UnsafeAdminStudyProgramHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminStudyProgramHandlerServer will
// result in compilation errors.
type UnsafeAdminStudyProgramHandlerServer interface {
	mustEmbedUnimplementedAdminStudyProgramHandlerServer()
}

func RegisterAdminStudyProgramHandlerServer(s grpc.ServiceRegistrar, srv AdminStudyProgramHandlerServer) {
	s.RegisterService(&AdminStudyProgramHandler_ServiceDesc, srv)
}

func _AdminStudyProgramHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminStudyProgramHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_study_program.AdminStudyProgramHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminStudyProgramHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminStudyProgramHandler_GetDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminStudyProgramHandlerServer).GetDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_study_program.AdminStudyProgramHandler/GetDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminStudyProgramHandlerServer).GetDetail(ctx, req.(*GetDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminStudyProgramHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminStudyProgramHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_study_program.AdminStudyProgramHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminStudyProgramHandlerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminStudyProgramHandler_UpdateDegree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDegreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminStudyProgramHandlerServer).UpdateDegree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_study_program.AdminStudyProgramHandler/UpdateDegree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminStudyProgramHandlerServer).UpdateDegree(ctx, req.(*UpdateDegreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminStudyProgramHandler_ServiceDesc is the grpc.ServiceDesc for AdminStudyProgramHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminStudyProgramHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_study_program.AdminStudyProgramHandler",
	HandlerType: (*AdminStudyProgramHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _AdminStudyProgramHandler_GetList_Handler,
		},
		{
			MethodName: "GetDetail",
			Handler:    _AdminStudyProgramHandler_GetDetail_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AdminStudyProgramHandler_Update_Handler,
		},
		{
			MethodName: "UpdateDegree",
			Handler:    _AdminStudyProgramHandler_UpdateDegree_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/study_program.proto",
}
