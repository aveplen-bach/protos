// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: s3g.proto

package s3_grpc_gateway

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

// S3GatewayClient is the client API for S3Gateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type S3GatewayClient interface {
	CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketResponse, error)
	DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*DeleteBucketResponse, error)
	ListBuckets(ctx context.Context, in *ListBucketsRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error)
	PutObject(ctx context.Context, in *PutObjectRequest, opts ...grpc.CallOption) (*PutObjectResponse, error)
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error)
}

type s3GatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewS3GatewayClient(cc grpc.ClientConnInterface) S3GatewayClient {
	return &s3GatewayClient{cc}
}

func (c *s3GatewayClient) CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketResponse, error) {
	out := new(CreateBucketResponse)
	err := c.cc.Invoke(ctx, "/aveplen.s3g.S3Gateway/CreateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3GatewayClient) DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*DeleteBucketResponse, error) {
	out := new(DeleteBucketResponse)
	err := c.cc.Invoke(ctx, "/aveplen.s3g.S3Gateway/DeleteBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3GatewayClient) ListBuckets(ctx context.Context, in *ListBucketsRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error) {
	out := new(ListBucketsResponse)
	err := c.cc.Invoke(ctx, "/aveplen.s3g.S3Gateway/ListBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3GatewayClient) PutObject(ctx context.Context, in *PutObjectRequest, opts ...grpc.CallOption) (*PutObjectResponse, error) {
	out := new(PutObjectResponse)
	err := c.cc.Invoke(ctx, "/aveplen.s3g.S3Gateway/PutObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3GatewayClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error) {
	out := new(GetObjectResponse)
	err := c.cc.Invoke(ctx, "/aveplen.s3g.S3Gateway/GetObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S3GatewayServer is the server API for S3Gateway service.
// All implementations must embed UnimplementedS3GatewayServer
// for forward compatibility
type S3GatewayServer interface {
	CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketResponse, error)
	DeleteBucket(context.Context, *DeleteBucketRequest) (*DeleteBucketResponse, error)
	ListBuckets(context.Context, *ListBucketsRequest) (*ListBucketsResponse, error)
	PutObject(context.Context, *PutObjectRequest) (*PutObjectResponse, error)
	GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error)
	mustEmbedUnimplementedS3GatewayServer()
}

// UnimplementedS3GatewayServer must be embedded to have forward compatible implementations.
type UnimplementedS3GatewayServer struct {
}

func (UnimplementedS3GatewayServer) CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBucket not implemented")
}
func (UnimplementedS3GatewayServer) DeleteBucket(context.Context, *DeleteBucketRequest) (*DeleteBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBucket not implemented")
}
func (UnimplementedS3GatewayServer) ListBuckets(context.Context, *ListBucketsRequest) (*ListBucketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuckets not implemented")
}
func (UnimplementedS3GatewayServer) PutObject(context.Context, *PutObjectRequest) (*PutObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutObject not implemented")
}
func (UnimplementedS3GatewayServer) GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}
func (UnimplementedS3GatewayServer) mustEmbedUnimplementedS3GatewayServer() {}

// UnsafeS3GatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to S3GatewayServer will
// result in compilation errors.
type UnsafeS3GatewayServer interface {
	mustEmbedUnimplementedS3GatewayServer()
}

func RegisterS3GatewayServer(s grpc.ServiceRegistrar, srv S3GatewayServer) {
	s.RegisterService(&S3Gateway_ServiceDesc, srv)
}

func _S3Gateway_CreateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3GatewayServer).CreateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aveplen.s3g.S3Gateway/CreateBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3GatewayServer).CreateBucket(ctx, req.(*CreateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3Gateway_DeleteBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3GatewayServer).DeleteBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aveplen.s3g.S3Gateway/DeleteBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3GatewayServer).DeleteBucket(ctx, req.(*DeleteBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3Gateway_ListBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBucketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3GatewayServer).ListBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aveplen.s3g.S3Gateway/ListBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3GatewayServer).ListBuckets(ctx, req.(*ListBucketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3Gateway_PutObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3GatewayServer).PutObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aveplen.s3g.S3Gateway/PutObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3GatewayServer).PutObject(ctx, req.(*PutObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3Gateway_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3GatewayServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aveplen.s3g.S3Gateway/GetObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3GatewayServer).GetObject(ctx, req.(*GetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// S3Gateway_ServiceDesc is the grpc.ServiceDesc for S3Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var S3Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aveplen.s3g.S3Gateway",
	HandlerType: (*S3GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBucket",
			Handler:    _S3Gateway_CreateBucket_Handler,
		},
		{
			MethodName: "DeleteBucket",
			Handler:    _S3Gateway_DeleteBucket_Handler,
		},
		{
			MethodName: "ListBuckets",
			Handler:    _S3Gateway_ListBuckets_Handler,
		},
		{
			MethodName: "PutObject",
			Handler:    _S3Gateway_PutObject_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _S3Gateway_GetObject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s3g.proto",
}