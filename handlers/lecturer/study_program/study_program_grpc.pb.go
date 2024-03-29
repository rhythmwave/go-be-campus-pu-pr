// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: lecturer/study_program.proto

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

// LecturerStudyProgramHandlerClient is the client API for LecturerStudyProgramHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LecturerStudyProgramHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
}

type lecturerStudyProgramHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewLecturerStudyProgramHandlerClient(cc grpc.ClientConnInterface) LecturerStudyProgramHandlerClient {
	return &lecturerStudyProgramHandlerClient{cc}
}

func (c *lecturerStudyProgramHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/lecturer_study_program.LecturerStudyProgramHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LecturerStudyProgramHandlerServer is the server API for LecturerStudyProgramHandler service.
// All implementations must embed UnimplementedLecturerStudyProgramHandlerServer
// for forward compatibility
type LecturerStudyProgramHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	mustEmbedUnimplementedLecturerStudyProgramHandlerServer()
}

// UnimplementedLecturerStudyProgramHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedLecturerStudyProgramHandlerServer struct {
}

func (UnimplementedLecturerStudyProgramHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedLecturerStudyProgramHandlerServer) mustEmbedUnimplementedLecturerStudyProgramHandlerServer() {
}

// UnsafeLecturerStudyProgramHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LecturerStudyProgramHandlerServer will
// result in compilation errors.
type UnsafeLecturerStudyProgramHandlerServer interface {
	mustEmbedUnimplementedLecturerStudyProgramHandlerServer()
}

func RegisterLecturerStudyProgramHandlerServer(s grpc.ServiceRegistrar, srv LecturerStudyProgramHandlerServer) {
	s.RegisterService(&LecturerStudyProgramHandler_ServiceDesc, srv)
}

func _LecturerStudyProgramHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LecturerStudyProgramHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lecturer_study_program.LecturerStudyProgramHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LecturerStudyProgramHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LecturerStudyProgramHandler_ServiceDesc is the grpc.ServiceDesc for LecturerStudyProgramHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LecturerStudyProgramHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lecturer_study_program.LecturerStudyProgramHandler",
	HandlerType: (*LecturerStudyProgramHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _LecturerStudyProgramHandler_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lecturer/study_program.proto",
}
