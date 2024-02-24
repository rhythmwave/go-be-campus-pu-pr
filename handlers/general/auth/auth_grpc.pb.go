// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: general/auth.proto

package auth

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

// GeneralAuthHandlerClient is the client API for GeneralAuthHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeneralAuthHandlerClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	GetSsoAuth(ctx context.Context, in *GetSsoAuthRequest, opts ...grpc.CallOption) (*GetSsoAuthResponse, error)
	LoginWithSso(ctx context.Context, in *LoginWithSsoRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type generalAuthHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewGeneralAuthHandlerClient(cc grpc.ClientConnInterface) GeneralAuthHandlerClient {
	return &generalAuthHandlerClient{cc}
}

func (c *generalAuthHandlerClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/general_auth.GeneralAuthHandler/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generalAuthHandlerClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/general_auth.GeneralAuthHandler/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generalAuthHandlerClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error) {
	out := new(UpdatePasswordResponse)
	err := c.cc.Invoke(ctx, "/general_auth.GeneralAuthHandler/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generalAuthHandlerClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/general_auth.GeneralAuthHandler/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generalAuthHandlerClient) GetSsoAuth(ctx context.Context, in *GetSsoAuthRequest, opts ...grpc.CallOption) (*GetSsoAuthResponse, error) {
	out := new(GetSsoAuthResponse)
	err := c.cc.Invoke(ctx, "/general_auth.GeneralAuthHandler/GetSsoAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *generalAuthHandlerClient) LoginWithSso(ctx context.Context, in *LoginWithSsoRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/general_auth.GeneralAuthHandler/LoginWithSso", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeneralAuthHandlerServer is the server API for GeneralAuthHandler service.
// All implementations must embed UnimplementedGeneralAuthHandlerServer
// for forward compatibility
type GeneralAuthHandlerServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*LoginResponse, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	GetSsoAuth(context.Context, *GetSsoAuthRequest) (*GetSsoAuthResponse, error)
	LoginWithSso(context.Context, *LoginWithSsoRequest) (*LoginResponse, error)
	mustEmbedUnimplementedGeneralAuthHandlerServer()
}

// UnimplementedGeneralAuthHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedGeneralAuthHandlerServer struct {
}

func (UnimplementedGeneralAuthHandlerServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGeneralAuthHandlerServer) RefreshToken(context.Context, *RefreshTokenRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedGeneralAuthHandlerServer) UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePassword not implemented")
}
func (UnimplementedGeneralAuthHandlerServer) Logout(context.Context, *LogoutRequest) (*LogoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedGeneralAuthHandlerServer) GetSsoAuth(context.Context, *GetSsoAuthRequest) (*GetSsoAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSsoAuth not implemented")
}
func (UnimplementedGeneralAuthHandlerServer) LoginWithSso(context.Context, *LoginWithSsoRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginWithSso not implemented")
}
func (UnimplementedGeneralAuthHandlerServer) mustEmbedUnimplementedGeneralAuthHandlerServer() {}

// UnsafeGeneralAuthHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeneralAuthHandlerServer will
// result in compilation errors.
type UnsafeGeneralAuthHandlerServer interface {
	mustEmbedUnimplementedGeneralAuthHandlerServer()
}

func RegisterGeneralAuthHandlerServer(s grpc.ServiceRegistrar, srv GeneralAuthHandlerServer) {
	s.RegisterService(&GeneralAuthHandler_ServiceDesc, srv)
}

func _GeneralAuthHandler_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneralAuthHandlerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/general_auth.GeneralAuthHandler/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneralAuthHandlerServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeneralAuthHandler_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneralAuthHandlerServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/general_auth.GeneralAuthHandler/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneralAuthHandlerServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeneralAuthHandler_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneralAuthHandlerServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/general_auth.GeneralAuthHandler/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneralAuthHandlerServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeneralAuthHandler_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneralAuthHandlerServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/general_auth.GeneralAuthHandler/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneralAuthHandlerServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeneralAuthHandler_GetSsoAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSsoAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneralAuthHandlerServer).GetSsoAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/general_auth.GeneralAuthHandler/GetSsoAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneralAuthHandlerServer).GetSsoAuth(ctx, req.(*GetSsoAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeneralAuthHandler_LoginWithSso_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginWithSsoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeneralAuthHandlerServer).LoginWithSso(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/general_auth.GeneralAuthHandler/LoginWithSso",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeneralAuthHandlerServer).LoginWithSso(ctx, req.(*LoginWithSsoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GeneralAuthHandler_ServiceDesc is the grpc.ServiceDesc for GeneralAuthHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeneralAuthHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "general_auth.GeneralAuthHandler",
	HandlerType: (*GeneralAuthHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _GeneralAuthHandler_Login_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _GeneralAuthHandler_RefreshToken_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _GeneralAuthHandler_UpdatePassword_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _GeneralAuthHandler_Logout_Handler,
		},
		{
			MethodName: "GetSsoAuth",
			Handler:    _GeneralAuthHandler_GetSsoAuth_Handler,
		},
		{
			MethodName: "LoginWithSso",
			Handler:    _GeneralAuthHandler_LoginWithSso_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "general/auth.proto",
}