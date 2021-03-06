// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.2
// source: invoice_service.proto

package src_core

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type AuthInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId     string `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	AppSecret string `protobuf:"bytes,2,opt,name=appSecret,proto3" json:"appSecret,omitempty"`
}

func (x *AuthInfo) Reset() {
	*x = AuthInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthInfo) ProtoMessage() {}

func (x *AuthInfo) ProtoReflect() protoreflect.Message {
	mi := &file_invoice_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthInfo.ProtoReflect.Descriptor instead.
func (*AuthInfo) Descriptor() ([]byte, []int) {
	return file_invoice_service_proto_rawDescGZIP(), []int{0}
}

func (x *AuthInfo) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *AuthInfo) GetAppSecret() string {
	if x != nil {
		return x.AppSecret
	}
	return ""
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_invoice_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_invoice_service_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Result) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_invoice_service_proto protoreflect.FileDescriptor

var file_invoice_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x22, 0x3e, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x70, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x22, 0x2e, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x32, 0x47, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x2e, 0x73, 0x72, 0x63, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x10, 0x2e, 0x73, 0x72, 0x63, 0x5f, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x40, 0x0a, 0x1f, 0x69, 0x6f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0f, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x72, 0x63, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_invoice_service_proto_rawDescOnce sync.Once
	file_invoice_service_proto_rawDescData = file_invoice_service_proto_rawDesc
)

func file_invoice_service_proto_rawDescGZIP() []byte {
	file_invoice_service_proto_rawDescOnce.Do(func() {
		file_invoice_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_invoice_service_proto_rawDescData)
	})
	return file_invoice_service_proto_rawDescData
}

var file_invoice_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_invoice_service_proto_goTypes = []interface{}{
	(*AuthInfo)(nil), // 0: src_core.AuthInfo
	(*Result)(nil),   // 1: src_core.Result
}
var file_invoice_service_proto_depIdxs = []int32{
	0, // 0: src_core.InvoiceService.InitInvoice:input_type -> src_core.AuthInfo
	1, // 1: src_core.InvoiceService.InitInvoice:output_type -> src_core.Result
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_invoice_service_proto_init() }
func file_invoice_service_proto_init() {
	if File_invoice_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_invoice_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthInfo); i {
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
		file_invoice_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_invoice_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_invoice_service_proto_goTypes,
		DependencyIndexes: file_invoice_service_proto_depIdxs,
		MessageInfos:      file_invoice_service_proto_msgTypes,
	}.Build()
	File_invoice_service_proto = out.File
	file_invoice_service_proto_rawDesc = nil
	file_invoice_service_proto_goTypes = nil
	file_invoice_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// InvoiceServiceClient is the client API for InvoiceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InvoiceServiceClient interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	InitInvoice(ctx context.Context, in *AuthInfo, opts ...grpc.CallOption) (*Result, error)
}

type invoiceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInvoiceServiceClient(cc grpc.ClientConnInterface) InvoiceServiceClient {
	return &invoiceServiceClient{cc}
}

func (c *invoiceServiceClient) InitInvoice(ctx context.Context, in *AuthInfo, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/src_core.InvoiceService/InitInvoice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InvoiceServiceServer is the server API for InvoiceService service.
type InvoiceServiceServer interface {
	// A simple RPC.
	//
	// Obtains the feature at a given position.
	//
	// A feature with an empty name is returned if there's no feature at the given
	// position.
	InitInvoice(context.Context, *AuthInfo) (*Result, error)
}

// UnimplementedInvoiceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInvoiceServiceServer struct {
}

func (*UnimplementedInvoiceServiceServer) InitInvoice(context.Context, *AuthInfo) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitInvoice not implemented")
}

func RegisterInvoiceServiceServer(s *grpc.Server, srv InvoiceServiceServer) {
	s.RegisterService(&_InvoiceService_serviceDesc, srv)
}

func _InvoiceService_InitInvoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServiceServer).InitInvoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/src_core.InvoiceService/InitInvoice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServiceServer).InitInvoice(ctx, req.(*AuthInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _InvoiceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "src_core.InvoiceService",
	HandlerType: (*InvoiceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitInvoice",
			Handler:    _InvoiceService_InitInvoice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "invoice_service.proto",
}
