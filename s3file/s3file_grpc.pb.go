// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: s3file.proto

package s3file

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
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
	PutImageObject(ctx context.Context, in *ImageObject, opts ...grpc.CallOption) (*empty.Empty, error)
	GetImageObject(ctx context.Context, in *GetImageObjectRequest, opts ...grpc.CallOption) (*ImageObject, error)
}

type s3GatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewS3GatewayClient(cc grpc.ClientConnInterface) S3GatewayClient {
	return &s3GatewayClient{cc}
}

func (c *s3GatewayClient) PutImageObject(ctx context.Context, in *ImageObject, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/aveplen.s3file.S3Gateway/PutImageObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s3GatewayClient) GetImageObject(ctx context.Context, in *GetImageObjectRequest, opts ...grpc.CallOption) (*ImageObject, error) {
	out := new(ImageObject)
	err := c.cc.Invoke(ctx, "/aveplen.s3file.S3Gateway/GetImageObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// S3GatewayServer is the server API for S3Gateway service.
// All implementations must embed UnimplementedS3GatewayServer
// for forward compatibility
type S3GatewayServer interface {
	PutImageObject(context.Context, *ImageObject) (*empty.Empty, error)
	GetImageObject(context.Context, *GetImageObjectRequest) (*ImageObject, error)
	mustEmbedUnimplementedS3GatewayServer()
}

// UnimplementedS3GatewayServer must be embedded to have forward compatible implementations.
type UnimplementedS3GatewayServer struct {
}

func (UnimplementedS3GatewayServer) PutImageObject(context.Context, *ImageObject) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutImageObject not implemented")
}
func (UnimplementedS3GatewayServer) GetImageObject(context.Context, *GetImageObjectRequest) (*ImageObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageObject not implemented")
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

func _S3Gateway_PutImageObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3GatewayServer).PutImageObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aveplen.s3file.S3Gateway/PutImageObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3GatewayServer).PutImageObject(ctx, req.(*ImageObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _S3Gateway_GetImageObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S3GatewayServer).GetImageObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aveplen.s3file.S3Gateway/GetImageObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S3GatewayServer).GetImageObject(ctx, req.(*GetImageObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// S3Gateway_ServiceDesc is the grpc.ServiceDesc for S3Gateway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var S3Gateway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aveplen.s3file.S3Gateway",
	HandlerType: (*S3GatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutImageObject",
			Handler:    _S3Gateway_PutImageObject_Handler,
		},
		{
			MethodName: "GetImageObject",
			Handler:    _S3Gateway_GetImageObject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s3file.proto",
}
