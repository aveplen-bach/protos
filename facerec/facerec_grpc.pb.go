// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: facerec.proto

package face_recognition_service

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

// FaceRecognitionClient is the client API for FaceRecognition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FaceRecognitionClient interface {
	ExtractFFVectorV1(ctx context.Context, in *ExtractFFVectorV1Request, opts ...grpc.CallOption) (*ExtractFFVectorV1Response, error)
}

type faceRecognitionClient struct {
	cc grpc.ClientConnInterface
}

func NewFaceRecognitionClient(cc grpc.ClientConnInterface) FaceRecognitionClient {
	return &faceRecognitionClient{cc}
}

func (c *faceRecognitionClient) ExtractFFVectorV1(ctx context.Context, in *ExtractFFVectorV1Request, opts ...grpc.CallOption) (*ExtractFFVectorV1Response, error) {
	out := new(ExtractFFVectorV1Response)
	err := c.cc.Invoke(ctx, "/aveplen.facerec.FaceRecognition/ExtractFFVectorV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FaceRecognitionServer is the server API for FaceRecognition service.
// All implementations must embed UnimplementedFaceRecognitionServer
// for forward compatibility
type FaceRecognitionServer interface {
	ExtractFFVectorV1(context.Context, *ExtractFFVectorV1Request) (*ExtractFFVectorV1Response, error)
	mustEmbedUnimplementedFaceRecognitionServer()
}

// UnimplementedFaceRecognitionServer must be embedded to have forward compatible implementations.
type UnimplementedFaceRecognitionServer struct {
}

func (UnimplementedFaceRecognitionServer) ExtractFFVectorV1(context.Context, *ExtractFFVectorV1Request) (*ExtractFFVectorV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtractFFVectorV1 not implemented")
}
func (UnimplementedFaceRecognitionServer) mustEmbedUnimplementedFaceRecognitionServer() {}

// UnsafeFaceRecognitionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FaceRecognitionServer will
// result in compilation errors.
type UnsafeFaceRecognitionServer interface {
	mustEmbedUnimplementedFaceRecognitionServer()
}

func RegisterFaceRecognitionServer(s grpc.ServiceRegistrar, srv FaceRecognitionServer) {
	s.RegisterService(&FaceRecognition_ServiceDesc, srv)
}

func _FaceRecognition_ExtractFFVectorV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtractFFVectorV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FaceRecognitionServer).ExtractFFVectorV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aveplen.facerec.FaceRecognition/ExtractFFVectorV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FaceRecognitionServer).ExtractFFVectorV1(ctx, req.(*ExtractFFVectorV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// FaceRecognition_ServiceDesc is the grpc.ServiceDesc for FaceRecognition service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FaceRecognition_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aveplen.facerec.FaceRecognition",
	HandlerType: (*FaceRecognitionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExtractFFVectorV1",
			Handler:    _FaceRecognition_ExtractFFVectorV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "facerec.proto",
}