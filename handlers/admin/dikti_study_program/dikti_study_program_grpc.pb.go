// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: admin/dikti_study_program.proto

package dikti_study_program

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

// AdminDiktiStudyProgramHandlerClient is the client API for AdminDiktiStudyProgramHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminDiktiStudyProgramHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
}

type adminDiktiStudyProgramHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminDiktiStudyProgramHandlerClient(cc grpc.ClientConnInterface) AdminDiktiStudyProgramHandlerClient {
	return &adminDiktiStudyProgramHandlerClient{cc}
}

func (c *adminDiktiStudyProgramHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/admin_dikti_study_program.AdminDiktiStudyProgramHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminDiktiStudyProgramHandlerServer is the server API for AdminDiktiStudyProgramHandler service.
// All implementations must embed UnimplementedAdminDiktiStudyProgramHandlerServer
// for forward compatibility
type AdminDiktiStudyProgramHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	mustEmbedUnimplementedAdminDiktiStudyProgramHandlerServer()
}

// UnimplementedAdminDiktiStudyProgramHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedAdminDiktiStudyProgramHandlerServer struct {
}

func (UnimplementedAdminDiktiStudyProgramHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedAdminDiktiStudyProgramHandlerServer) mustEmbedUnimplementedAdminDiktiStudyProgramHandlerServer() {
}

// UnsafeAdminDiktiStudyProgramHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminDiktiStudyProgramHandlerServer will
// result in compilation errors.
type UnsafeAdminDiktiStudyProgramHandlerServer interface {
	mustEmbedUnimplementedAdminDiktiStudyProgramHandlerServer()
}

func RegisterAdminDiktiStudyProgramHandlerServer(s grpc.ServiceRegistrar, srv AdminDiktiStudyProgramHandlerServer) {
	s.RegisterService(&AdminDiktiStudyProgramHandler_ServiceDesc, srv)
}

func _AdminDiktiStudyProgramHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminDiktiStudyProgramHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/admin_dikti_study_program.AdminDiktiStudyProgramHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminDiktiStudyProgramHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminDiktiStudyProgramHandler_ServiceDesc is the grpc.ServiceDesc for AdminDiktiStudyProgramHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminDiktiStudyProgramHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin_dikti_study_program.AdminDiktiStudyProgramHandler",
	HandlerType: (*AdminDiktiStudyProgramHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _AdminDiktiStudyProgramHandler_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/dikti_study_program.proto",
}
