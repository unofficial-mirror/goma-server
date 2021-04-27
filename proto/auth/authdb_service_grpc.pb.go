// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// AuthDBServiceClient is the client API for AuthDBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthDBServiceClient interface {
	CheckMembership(ctx context.Context, in *CheckMembershipReq, opts ...grpc.CallOption) (*CheckMembershipResp, error)
}

type authDBServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthDBServiceClient(cc grpc.ClientConnInterface) AuthDBServiceClient {
	return &authDBServiceClient{cc}
}

func (c *authDBServiceClient) CheckMembership(ctx context.Context, in *CheckMembershipReq, opts ...grpc.CallOption) (*CheckMembershipResp, error) {
	out := new(CheckMembershipResp)
	err := c.cc.Invoke(ctx, "/auth.AuthDBService/CheckMembership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthDBServiceServer is the server API for AuthDBService service.
// All implementations must embed UnimplementedAuthDBServiceServer
// for forward compatibility
type AuthDBServiceServer interface {
	CheckMembership(context.Context, *CheckMembershipReq) (*CheckMembershipResp, error)
	mustEmbedUnimplementedAuthDBServiceServer()
}

// UnimplementedAuthDBServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthDBServiceServer struct {
}

func (UnimplementedAuthDBServiceServer) CheckMembership(context.Context, *CheckMembershipReq) (*CheckMembershipResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckMembership not implemented")
}
func (UnimplementedAuthDBServiceServer) mustEmbedUnimplementedAuthDBServiceServer() {}

// UnsafeAuthDBServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthDBServiceServer will
// result in compilation errors.
type UnsafeAuthDBServiceServer interface {
	mustEmbedUnimplementedAuthDBServiceServer()
}

func RegisterAuthDBServiceServer(s grpc.ServiceRegistrar, srv AuthDBServiceServer) {
	s.RegisterService(&AuthDBService_ServiceDesc, srv)
}

func _AuthDBService_CheckMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckMembershipReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthDBServiceServer).CheckMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthDBService/CheckMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthDBServiceServer).CheckMembership(ctx, req.(*CheckMembershipReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthDBService_ServiceDesc is the grpc.ServiceDesc for AuthDBService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthDBService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthDBService",
	HandlerType: (*AuthDBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckMembership",
			Handler:    _AuthDBService_CheckMembership_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/authdb_service.proto",
}