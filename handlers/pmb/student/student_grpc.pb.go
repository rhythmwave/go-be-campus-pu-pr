// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pmb/student.proto

package student

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

// PmbStudentHandlerClient is the client API for PmbStudentHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PmbStudentHandlerClient interface {
	BulkCreate(ctx context.Context, in *BulkCreateRequest, opts ...grpc.CallOption) (*BulkCreateResponse, error)
}

type pmbStudentHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewPmbStudentHandlerClient(cc grpc.ClientConnInterface) PmbStudentHandlerClient {
	return &pmbStudentHandlerClient{cc}
}

func (c *pmbStudentHandlerClient) BulkCreate(ctx context.Context, in *BulkCreateRequest, opts ...grpc.CallOption) (*BulkCreateResponse, error) {
	out := new(BulkCreateResponse)
	err := c.cc.Invoke(ctx, "/pmb_student.PmbStudentHandler/BulkCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PmbStudentHandlerServer is the server API for PmbStudentHandler service.
// All implementations must embed UnimplementedPmbStudentHandlerServer
// for forward compatibility
type PmbStudentHandlerServer interface {
	BulkCreate(context.Context, *BulkCreateRequest) (*BulkCreateResponse, error)
	mustEmbedUnimplementedPmbStudentHandlerServer()
}

// UnimplementedPmbStudentHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedPmbStudentHandlerServer struct {
}

func (UnimplementedPmbStudentHandlerServer) BulkCreate(context.Context, *BulkCreateRequest) (*BulkCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkCreate not implemented")
}
func (UnimplementedPmbStudentHandlerServer) mustEmbedUnimplementedPmbStudentHandlerServer() {}

// UnsafePmbStudentHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PmbStudentHandlerServer will
// result in compilation errors.
type UnsafePmbStudentHandlerServer interface {
	mustEmbedUnimplementedPmbStudentHandlerServer()
}

func RegisterPmbStudentHandlerServer(s grpc.ServiceRegistrar, srv PmbStudentHandlerServer) {
	s.RegisterService(&PmbStudentHandler_ServiceDesc, srv)
}

func _PmbStudentHandler_BulkCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PmbStudentHandlerServer).BulkCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pmb_student.PmbStudentHandler/BulkCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PmbStudentHandlerServer).BulkCreate(ctx, req.(*BulkCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PmbStudentHandler_ServiceDesc is the grpc.ServiceDesc for PmbStudentHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PmbStudentHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pmb_student.PmbStudentHandler",
	HandlerType: (*PmbStudentHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BulkCreate",
			Handler:    _PmbStudentHandler_BulkCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pmb/student.proto",
}