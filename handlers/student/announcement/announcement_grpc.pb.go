// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: student/announcement.proto

package announcement

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

// StudentAnnouncementHandlerClient is the client API for StudentAnnouncementHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentAnnouncementHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
}

type studentAnnouncementHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentAnnouncementHandlerClient(cc grpc.ClientConnInterface) StudentAnnouncementHandlerClient {
	return &studentAnnouncementHandlerClient{cc}
}

func (c *studentAnnouncementHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/student_announcement.StudentAnnouncementHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentAnnouncementHandlerServer is the server API for StudentAnnouncementHandler service.
// All implementations must embed UnimplementedStudentAnnouncementHandlerServer
// for forward compatibility
type StudentAnnouncementHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	mustEmbedUnimplementedStudentAnnouncementHandlerServer()
}

// UnimplementedStudentAnnouncementHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedStudentAnnouncementHandlerServer struct {
}

func (UnimplementedStudentAnnouncementHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedStudentAnnouncementHandlerServer) mustEmbedUnimplementedStudentAnnouncementHandlerServer() {
}

// UnsafeStudentAnnouncementHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentAnnouncementHandlerServer will
// result in compilation errors.
type UnsafeStudentAnnouncementHandlerServer interface {
	mustEmbedUnimplementedStudentAnnouncementHandlerServer()
}

func RegisterStudentAnnouncementHandlerServer(s grpc.ServiceRegistrar, srv StudentAnnouncementHandlerServer) {
	s.RegisterService(&StudentAnnouncementHandler_ServiceDesc, srv)
}

func _StudentAnnouncementHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentAnnouncementHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student_announcement.StudentAnnouncementHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentAnnouncementHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentAnnouncementHandler_ServiceDesc is the grpc.ServiceDesc for StudentAnnouncementHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentAnnouncementHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student_announcement.StudentAnnouncementHandler",
	HandlerType: (*StudentAnnouncementHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _StudentAnnouncementHandler_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student/announcement.proto",
}