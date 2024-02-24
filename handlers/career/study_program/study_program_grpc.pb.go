// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: career/study_program.proto

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

// CareerStudyProgramHandlerClient is the client API for CareerStudyProgramHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CareerStudyProgramHandlerClient interface {
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc.CallOption) (*GetDetailResponse, error)
}

type careerStudyProgramHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewCareerStudyProgramHandlerClient(cc grpc.ClientConnInterface) CareerStudyProgramHandlerClient {
	return &careerStudyProgramHandlerClient{cc}
}

func (c *careerStudyProgramHandlerClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/career_study_program.CareerStudyProgramHandler/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *careerStudyProgramHandlerClient) GetDetail(ctx context.Context, in *GetDetailRequest, opts ...grpc.CallOption) (*GetDetailResponse, error) {
	out := new(GetDetailResponse)
	err := c.cc.Invoke(ctx, "/career_study_program.CareerStudyProgramHandler/GetDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CareerStudyProgramHandlerServer is the server API for CareerStudyProgramHandler service.
// All implementations must embed UnimplementedCareerStudyProgramHandlerServer
// for forward compatibility
type CareerStudyProgramHandlerServer interface {
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	GetDetail(context.Context, *GetDetailRequest) (*GetDetailResponse, error)
	mustEmbedUnimplementedCareerStudyProgramHandlerServer()
}

// UnimplementedCareerStudyProgramHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedCareerStudyProgramHandlerServer struct {
}

func (UnimplementedCareerStudyProgramHandlerServer) GetList(context.Context, *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedCareerStudyProgramHandlerServer) GetDetail(context.Context, *GetDetailRequest) (*GetDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetail not implemented")
}
func (UnimplementedCareerStudyProgramHandlerServer) mustEmbedUnimplementedCareerStudyProgramHandlerServer() {
}

// UnsafeCareerStudyProgramHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CareerStudyProgramHandlerServer will
// result in compilation errors.
type UnsafeCareerStudyProgramHandlerServer interface {
	mustEmbedUnimplementedCareerStudyProgramHandlerServer()
}

func RegisterCareerStudyProgramHandlerServer(s grpc.ServiceRegistrar, srv CareerStudyProgramHandlerServer) {
	s.RegisterService(&CareerStudyProgramHandler_ServiceDesc, srv)
}

func _CareerStudyProgramHandler_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CareerStudyProgramHandlerServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/career_study_program.CareerStudyProgramHandler/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CareerStudyProgramHandlerServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CareerStudyProgramHandler_GetDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CareerStudyProgramHandlerServer).GetDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/career_study_program.CareerStudyProgramHandler/GetDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CareerStudyProgramHandlerServer).GetDetail(ctx, req.(*GetDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CareerStudyProgramHandler_ServiceDesc is the grpc.ServiceDesc for CareerStudyProgramHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CareerStudyProgramHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "career_study_program.CareerStudyProgramHandler",
	HandlerType: (*CareerStudyProgramHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetList",
			Handler:    _CareerStudyProgramHandler_GetList_Handler,
		},
		{
			MethodName: "GetDetail",
			Handler:    _CareerStudyProgramHandler_GetDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "career/study_program.proto",
}