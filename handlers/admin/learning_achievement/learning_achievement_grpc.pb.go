// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: admin/learning_achievement.proto

package learning_achievement

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

// AdminLearningAchievementHandlerClient is the client API for AdminLearningAchievementHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminLearningAchievementHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type adminLearningAchievementHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminLearningAchievementHandlerClient(cc grpc.ClientConnInterface) AdminLearningAchievementHandlerClient {
	return &adminLearningAchievementHandlerClient{cc}
}

func (c *adminLearningAchievementHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/admin_learning_achievement.AdminLearningAchievementHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminLearningAchievementHandlerClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/admin_learning_achievement.AdminLearningAchievementHandler/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminLearningAchievementHandlerClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/admin_learning_achievement.AdminLearningAchievementHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminLearningAchievementHandlerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/admin_learning_achievement.AdminLearningAchievementHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminLearningAchievementHandlerServer is the server API for AdminLearningAchievementHandler service.
// All implementations must embed UnimplementedAdminLearningAchievementHandlerServer
// for forward compatibility
type AdminLearningAchievementHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedAdminLearningAchievementHandlerServer()
}

// UnimplementedAdminLearningAchievementHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAdminLearningAchievementHandlerServer struct {
}

func (UnimplementedAdminLearningAchievementHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdminLearningAchievementHandlerServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAdminLearningAchievementHandlerServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAdminLearningAchievementHandlerServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAdminLearningAchievementHandlerServer) mustEmbedUnimplementedAdminLearningAchievementHandlerServer() {
}

// UnsafeAdminLearningAchievementHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminLearningAchievementHandlerServer will
// result in compilation errors.
type UnsafeAdminLearningAchievementHandlerServer interface {
	mustEmbedUnimplementedAdminLearningAchievementHandlerServer()
}

func RegisterAdminLearningAchievementHandlerServer(s grpc.ServiceRegistrar, srv AdminLearningAchievementHandlerServer) {
	s.RegisterService(&AdminLearningAchievementHandler_ServiceDesc, srv)
}

func _AdminLearningAchievementHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminLearningAchievementHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_learning_achievement.AdminLearningAchievementHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminLearningAchievementHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminLearningAchievementHandler_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminLearningAchievementHandlerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_learning_achievement.AdminLearningAchievementHandler/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminLearningAchievementHandlerServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminLearningAchievementHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminLearningAchievementHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_learning_achievement.AdminLearningAchievementHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminLearningAchievementHandlerServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminLearningAchievementHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminLearningAchievementHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_learning_achievement.AdminLearningAchievementHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminLearningAchievementHandlerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminLearningAchievementHandler_ServiceDesc is the grpc.ServiceDesc for AdminLearningAchievementHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminLearningAchievementHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_learning_achievement.AdminLearningAchievementHandler",
	HandlerType: (*AdminLearningAchievementHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _AdminLearningAchievementHandler_GetList_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AdminLearningAchievementHandler_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AdminLearningAchievementHandler_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AdminLearningAchievementHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/learning_achievement.proto",
}
