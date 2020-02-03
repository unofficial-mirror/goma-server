// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache/cache_service.proto

package cache

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("cache/cache_service.proto", fileDescriptor_7933906ce719ddba) }

var fileDescriptor_7933906ce719ddba = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x4e, 0x4c, 0xce,
	0x48, 0xd5, 0x07, 0x93, 0xf1, 0xc5, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0xac, 0x60, 0x41, 0x29, 0x41, 0x24, 0x15, 0x10, 0x19, 0xa3, 0x38, 0x2e, 0x1e,
	0x67, 0x10, 0x37, 0x18, 0xa2, 0x5e, 0x48, 0x8d, 0x8b, 0xd9, 0x3d, 0xb5, 0x44, 0x88, 0x57, 0x0f,
	0xa2, 0xc8, 0x3d, 0xb5, 0x24, 0x28, 0xb5, 0x50, 0x8a, 0x0f, 0x99, 0x5b, 0x5c, 0xa0, 0xc4, 0x00,
	0x52, 0x17, 0x50, 0x8a, 0x50, 0x17, 0x50, 0x8a, 0xa2, 0x0e, 0xcc, 0x05, 0xa9, 0x4b, 0x62, 0x03,
	0x5b, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x88, 0x55, 0xe0, 0x78, 0x9d, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CacheServiceClient is the client API for CacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheServiceClient interface {
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
	Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*PutResp, error)
}

type cacheServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheServiceClient(cc grpc.ClientConnInterface) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, "/cache.CacheService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheServiceClient) Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*PutResp, error) {
	out := new(PutResp)
	err := c.cc.Invoke(ctx, "/cache.CacheService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
type CacheServiceServer interface {
	Get(context.Context, *GetReq) (*GetResp, error)
	Put(context.Context, *PutReq) (*PutResp, error)
}

// UnimplementedCacheServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCacheServiceServer struct {
}

func (*UnimplementedCacheServiceServer) Get(ctx context.Context, req *GetReq) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCacheServiceServer) Put(ctx context.Context, req *PutReq) (*PutResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}

func RegisterCacheServiceServer(s *grpc.Server, srv CacheServiceServer) {
	s.RegisterService(&_CacheService_serviceDesc, srv)
}

func _CacheService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache.CacheService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cache.CacheService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).Put(ctx, req.(*PutReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CacheService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cache.CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CacheService_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _CacheService_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cache/cache_service.proto",
}
