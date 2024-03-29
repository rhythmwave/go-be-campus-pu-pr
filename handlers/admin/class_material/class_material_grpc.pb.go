// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: admin/class_material.proto

package class_material

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

// AdminClassMaterialHandlerClient is the client API for AdminClassMaterialHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminClassMaterialHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
}

type adminClassMaterialHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminClassMaterialHandlerClient(cc grpc.ClientConnInterface) AdminClassMaterialHandlerClient {
	return &adminClassMaterialHandlerClient{cc}
}

func (c *adminClassMaterialHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/admin_class_material.AdminClassMaterialHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminClassMaterialHandlerServer is the server API for AdminClassMaterialHandler service.
// All implementations must embed UnimplementedAdminClassMaterialHandlerServer
// for forward compatibility
type AdminClassMaterialHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	mustEmbedUnimplementedAdminClassMaterialHandlerServer()
}

// UnimplementedAdminClassMaterialHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAdminClassMaterialHandlerServer struct {
}

func (UnimplementedAdminClassMaterialHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdminClassMaterialHandlerServer) mustEmbedUnimplementedAdminClassMaterialHandlerServer() {
}

// UnsafeAdminClassMaterialHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminClassMaterialHandlerServer will
// result in compilation errors.
type UnsafeAdminClassMaterialHandlerServer interface {
	mustEmbedUnimplementedAdminClassMaterialHandlerServer()
}

func RegisterAdminClassMaterialHandlerServer(s grpc.ServiceRegistrar, srv AdminClassMaterialHandlerServer) {
	s.RegisterService(&AdminClassMaterialHandler_ServiceDesc, srv)
}

func _AdminClassMaterialHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminClassMaterialHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_class_material.AdminClassMaterialHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminClassMaterialHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminClassMaterialHandler_ServiceDesc is the grpc.ServiceDesc for AdminClassMaterialHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminClassMaterialHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_class_material.AdminClassMaterialHandler",
	HandlerType: (*AdminClassMaterialHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _AdminClassMaterialHandler_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/class_material.proto",
}
