// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: asset/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("asset/query.proto", fileDescriptor_4785ee6dd031b7c8) }

var fileDescriptor_4785ee6dd031b7c8 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0xb1, 0x6e, 0xc2, 0x30,
	0x10, 0x40, 0x93, 0xa1, 0xad, 0x94, 0xad, 0x1d, 0x3a, 0x44, 0x95, 0x3f, 0xa0, 0x52, 0x73, 0x4a,
	0xf9, 0x03, 0x36, 0x46, 0x56, 0xb6, 0x73, 0x64, 0x39, 0x96, 0x88, 0xcf, 0xe4, 0x2e, 0x40, 0xfe,
	0x82, 0xcf, 0x62, 0xcc, 0xc8, 0x88, 0x92, 0x1f, 0x41, 0xd8, 0x03, 0xcb, 0x49, 0xa7, 0x7b, 0xf7,
	0xf4, 0x8a, 0x4f, 0x64, 0x36, 0x02, 0x87, 0xc1, 0xf4, 0x63, 0x15, 0x7a, 0x12, 0xfa, 0xfa, 0x66,
	0xf2, 0xf6, 0x44, 0xde, 0x4a, 0xa8, 0xe2, 0x35, 0xcd, 0xf2, 0xc7, 0x12, 0xd9, 0xbd, 0x01, 0x0c,
	0x0e, 0xd0, 0x7b, 0x12, 0x14, 0x47, 0x9e, 0xd3, 0x57, 0xf9, 0xdb, 0x10, 0x77, 0xc4, 0xa0, 0x91,
	0x4d, 0xd2, 0xc1, 0xb1, 0xd6, 0x46, 0xb0, 0x86, 0x80, 0xd6, 0xf9, 0x08, 0x27, 0xf6, 0xff, 0xa3,
	0x78, 0xdb, 0x3e, 0x89, 0xf5, 0xe6, 0x3a, 0xab, 0x7c, 0x9a, 0x55, 0x7e, 0x9f, 0x55, 0x7e, 0x59,
	0x54, 0x36, 0x2d, 0x2a, 0xbb, 0x2d, 0x2a, 0xdb, 0x81, 0x75, 0xd2, 0x0e, 0xba, 0x6a, 0xa8, 0x83,
	0x57, 0x0f, 0xc4, 0x92, 0xbf, 0xa6, 0x45, 0xe7, 0xe1, 0x9c, 0x36, 0x90, 0x31, 0x18, 0xd6, 0xef,
	0x51, 0xbd, 0x7a, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xc3, 0x52, 0x4b, 0xd1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "songwongtp.asset.asset.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "asset/query.proto",
}
