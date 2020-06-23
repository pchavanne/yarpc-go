// Code generated by protoc-gen-go. DO NOT EDIT.
// source: encoding/protobuf/internal/testpb/test.proto

// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package testpb

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

type TestMessage struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestMessage) Reset()         { *m = TestMessage{} }
func (m *TestMessage) String() string { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()    {}
func (*TestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc320162ebaf2b25, []int{0}
}

func (m *TestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMessage.Unmarshal(m, b)
}
func (m *TestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMessage.Marshal(b, m, deterministic)
}
func (m *TestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessage.Merge(m, src)
}
func (m *TestMessage) XXX_Size() int {
	return xxx_messageInfo_TestMessage.Size(m)
}
func (m *TestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessage proto.InternalMessageInfo

func (m *TestMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*TestMessage)(nil), "uber.yarpc.encoding.protobuf.TestMessage")
}

func init() {
	proto.RegisterFile("encoding/protobuf/internal/testpb/test.proto", fileDescriptor_fc320162ebaf2b25)
}

var fileDescriptor_fc320162ebaf2b25 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0xcd, 0x4b, 0xce,
	0x4f, 0xc9, 0xcc, 0x4b, 0xd7, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0xcf, 0xcc,
	0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0x49, 0x2d, 0x2e, 0x29, 0x48, 0x02, 0x53, 0x7a,
	0x60, 0x59, 0x21, 0x99, 0xd2, 0xa4, 0xd4, 0x22, 0xbd, 0xca, 0xc4, 0xa2, 0x82, 0x64, 0x3d, 0x98,
	0x46, 0x3d, 0x98, 0x46, 0x25, 0x65, 0x2e, 0xee, 0x90, 0xd4, 0xe2, 0x12, 0xdf, 0xd4, 0xe2, 0xe2,
	0xc4, 0xf4, 0x54, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x08, 0xc7, 0xe8, 0x24, 0x23, 0x17, 0x0b, 0x48, 0x95, 0x50, 0x2c, 0x17, 0x6b,
	0x68, 0x5e, 0x62, 0x51, 0xa5, 0x90, 0xa6, 0x1e, 0x3e, 0x53, 0xf5, 0x90, 0x8c, 0x94, 0x22, 0x5e,
	0xa9, 0x50, 0x12, 0x17, 0x9b, 0x4b, 0x69, 0x41, 0x4e, 0x6a, 0x05, 0x6d, 0xcc, 0xd7, 0x60, 0x34,
	0x60, 0x74, 0xe2, 0x88, 0x62, 0x83, 0x84, 0x51, 0x12, 0x1b, 0x58, 0x8d, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x2c, 0x33, 0x41, 0xcb, 0x4f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	Unary(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*TestMessage, error)
	Duplex(ctx context.Context, opts ...grpc.CallOption) (Test_DuplexClient, error)
}

type testClient struct {
	cc grpc.ClientConnInterface
}

func NewTestClient(cc grpc.ClientConnInterface) TestClient {
	return &testClient{cc}
}

func (c *testClient) Unary(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*TestMessage, error) {
	out := new(TestMessage)
	err := c.cc.Invoke(ctx, "/uber.yarpc.encoding.protobuf.Test/Unary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Duplex(ctx context.Context, opts ...grpc.CallOption) (Test_DuplexClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Test_serviceDesc.Streams[0], "/uber.yarpc.encoding.protobuf.Test/Duplex", opts...)
	if err != nil {
		return nil, err
	}
	x := &testDuplexClient{stream}
	return x, nil
}

type Test_DuplexClient interface {
	Send(*TestMessage) error
	Recv() (*TestMessage, error)
	grpc.ClientStream
}

type testDuplexClient struct {
	grpc.ClientStream
}

func (x *testDuplexClient) Send(m *TestMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testDuplexClient) Recv() (*TestMessage, error) {
	m := new(TestMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	Unary(context.Context, *TestMessage) (*TestMessage, error)
	Duplex(Test_DuplexServer) error
}

// UnimplementedTestServer can be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (*UnimplementedTestServer) Unary(ctx context.Context, req *TestMessage) (*TestMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary not implemented")
}
func (*UnimplementedTestServer) Duplex(srv Test_DuplexServer) error {
	return status.Errorf(codes.Unimplemented, "method Duplex not implemented")
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_Unary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).Unary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/uber.yarpc.encoding.protobuf.Test/Unary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).Unary(ctx, req.(*TestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Duplex_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServer).Duplex(&testDuplexServer{stream})
}

type Test_DuplexServer interface {
	Send(*TestMessage) error
	Recv() (*TestMessage, error)
	grpc.ServerStream
}

type testDuplexServer struct {
	grpc.ServerStream
}

func (x *testDuplexServer) Send(m *TestMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testDuplexServer) Recv() (*TestMessage, error) {
	m := new(TestMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "uber.yarpc.encoding.protobuf.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unary",
			Handler:    _Test_Unary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Duplex",
			Handler:       _Test_Duplex_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "encoding/protobuf/internal/testpb/test.proto",
}
