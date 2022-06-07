// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: facerec.proto

package facerec

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExtractFFVectorV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ExtractFFVectorV1Request) Reset() {
	*x = ExtractFFVectorV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_facerec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractFFVectorV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractFFVectorV1Request) ProtoMessage() {}

func (x *ExtractFFVectorV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_facerec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractFFVectorV1Request.ProtoReflect.Descriptor instead.
func (*ExtractFFVectorV1Request) Descriptor() ([]byte, []int) {
	return file_facerec_proto_rawDescGZIP(), []int{0}
}

func (x *ExtractFFVectorV1Request) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ExtractFFVectorV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ffvc []float64 `protobuf:"fixed64,1,rep,packed,name=ffvc,proto3" json:"ffvc,omitempty"`
}

func (x *ExtractFFVectorV1Response) Reset() {
	*x = ExtractFFVectorV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_facerec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractFFVectorV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractFFVectorV1Response) ProtoMessage() {}

func (x *ExtractFFVectorV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_facerec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractFFVectorV1Response.ProtoReflect.Descriptor instead.
func (*ExtractFFVectorV1Response) Descriptor() ([]byte, []int) {
	return file_facerec_proto_rawDescGZIP(), []int{1}
}

func (x *ExtractFFVectorV1Response) GetFfvc() []float64 {
	if x != nil {
		return x.Ffvc
	}
	return nil
}

type FFVectorDistanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ffvca []float64 `protobuf:"fixed64,1,rep,packed,name=ffvca,proto3" json:"ffvca,omitempty"`
	Ffvcb []float64 `protobuf:"fixed64,2,rep,packed,name=ffvcb,proto3" json:"ffvcb,omitempty"`
}

func (x *FFVectorDistanceRequest) Reset() {
	*x = FFVectorDistanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_facerec_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FFVectorDistanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FFVectorDistanceRequest) ProtoMessage() {}

func (x *FFVectorDistanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_facerec_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FFVectorDistanceRequest.ProtoReflect.Descriptor instead.
func (*FFVectorDistanceRequest) Descriptor() ([]byte, []int) {
	return file_facerec_proto_rawDescGZIP(), []int{2}
}

func (x *FFVectorDistanceRequest) GetFfvca() []float64 {
	if x != nil {
		return x.Ffvca
	}
	return nil
}

func (x *FFVectorDistanceRequest) GetFfvcb() []float64 {
	if x != nil {
		return x.Ffvcb
	}
	return nil
}

type FFVectorDistanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Distance float64 `protobuf:"fixed64,1,opt,name=distance,proto3" json:"distance,omitempty"`
}

func (x *FFVectorDistanceResponse) Reset() {
	*x = FFVectorDistanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_facerec_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FFVectorDistanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FFVectorDistanceResponse) ProtoMessage() {}

func (x *FFVectorDistanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_facerec_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FFVectorDistanceResponse.ProtoReflect.Descriptor instead.
func (*FFVectorDistanceResponse) Descriptor() ([]byte, []int) {
	return file_facerec_proto_rawDescGZIP(), []int{3}
}

func (x *FFVectorDistanceResponse) GetDistance() float64 {
	if x != nil {
		return x.Distance
	}
	return 0
}

var File_facerec_proto protoreflect.FileDescriptor

var file_facerec_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x61, 0x63, 0x65, 0x72, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x61, 0x76, 0x65, 0x70, 0x6c, 0x65, 0x6e, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x72, 0x65, 0x63,
	0x22, 0x2a, 0x0a, 0x18, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x46, 0x46, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x19,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x46, 0x46, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x66, 0x76,
	0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x04, 0x66, 0x66, 0x76, 0x63, 0x22, 0x45, 0x0a,
	0x17, 0x46, 0x46, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x66, 0x76, 0x63,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x05, 0x66, 0x66, 0x76, 0x63, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x66, 0x66, 0x76, 0x63, 0x62, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52, 0x05, 0x66,
	0x66, 0x76, 0x63, 0x62, 0x22, 0x36, 0x0a, 0x18, 0x46, 0x46, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x32, 0xe6, 0x01, 0x0a,
	0x0f, 0x46, 0x61, 0x63, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x6a, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x46, 0x46, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x56, 0x31, 0x12, 0x29, 0x2e, 0x61, 0x76, 0x65, 0x70, 0x6c, 0x65, 0x6e, 0x2e,
	0x66, 0x61, 0x63, 0x65, 0x72, 0x65, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x46,
	0x46, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x61, 0x76, 0x65, 0x70, 0x6c, 0x65, 0x6e, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x72,
	0x65, 0x63, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x46, 0x46, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x10,
	0x46, 0x46, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x28, 0x2e, 0x61, 0x76, 0x65, 0x70, 0x6c, 0x65, 0x6e, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x72,
	0x65, 0x63, 0x2e, 0x46, 0x46, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x76, 0x65,
	0x70, 0x6c, 0x65, 0x6e, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x72, 0x65, 0x63, 0x2e, 0x46, 0x46, 0x56,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x44, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x76, 0x65, 0x70, 0x6c, 0x65, 0x6e, 0x2d, 0x62, 0x61, 0x63, 0x68,
	0x2f, 0x66, 0x61, 0x63, 0x65, 0x72, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_facerec_proto_rawDescOnce sync.Once
	file_facerec_proto_rawDescData = file_facerec_proto_rawDesc
)

func file_facerec_proto_rawDescGZIP() []byte {
	file_facerec_proto_rawDescOnce.Do(func() {
		file_facerec_proto_rawDescData = protoimpl.X.CompressGZIP(file_facerec_proto_rawDescData)
	})
	return file_facerec_proto_rawDescData
}

var file_facerec_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_facerec_proto_goTypes = []interface{}{
	(*ExtractFFVectorV1Request)(nil),  // 0: aveplen.facerec.ExtractFFVectorV1Request
	(*ExtractFFVectorV1Response)(nil), // 1: aveplen.facerec.ExtractFFVectorV1Response
	(*FFVectorDistanceRequest)(nil),   // 2: aveplen.facerec.FFVectorDistanceRequest
	(*FFVectorDistanceResponse)(nil),  // 3: aveplen.facerec.FFVectorDistanceResponse
}
var file_facerec_proto_depIdxs = []int32{
	0, // 0: aveplen.facerec.FaceRecognition.ExtractFFVectorV1:input_type -> aveplen.facerec.ExtractFFVectorV1Request
	2, // 1: aveplen.facerec.FaceRecognition.FFVectorDistance:input_type -> aveplen.facerec.FFVectorDistanceRequest
	1, // 2: aveplen.facerec.FaceRecognition.ExtractFFVectorV1:output_type -> aveplen.facerec.ExtractFFVectorV1Response
	3, // 3: aveplen.facerec.FaceRecognition.FFVectorDistance:output_type -> aveplen.facerec.FFVectorDistanceResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_facerec_proto_init() }
func file_facerec_proto_init() {
	if File_facerec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_facerec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractFFVectorV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_facerec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractFFVectorV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_facerec_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FFVectorDistanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_facerec_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FFVectorDistanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_facerec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_facerec_proto_goTypes,
		DependencyIndexes: file_facerec_proto_depIdxs,
		MessageInfos:      file_facerec_proto_msgTypes,
	}.Build()
	File_facerec_proto = out.File
	file_facerec_proto_rawDesc = nil
	file_facerec_proto_goTypes = nil
	file_facerec_proto_depIdxs = nil
}
