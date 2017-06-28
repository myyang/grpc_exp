// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

/*
Package grpc_exp is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	HelloMsg
	HelloReply
*/
package grpc_exp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HelloMsg is simple message
type HelloMsg struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *HelloMsg) Reset()                    { *m = HelloMsg{} }
func (m *HelloMsg) String() string            { return proto.CompactTextString(m) }
func (*HelloMsg) ProtoMessage()               {}
func (*HelloMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloMsg) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type HelloReply struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloMsg)(nil), "grpc_exp.HelloMsg")
	proto.RegisterType((*HelloReply)(nil), "grpc_exp.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HelloService service

type HelloServiceClient interface {
	SayHello(ctx context.Context, in *HelloMsg, opts ...grpc.CallOption) (*HelloReply, error)
	SayBiStreamHello(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayBiStreamHelloClient, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *HelloMsg, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/grpc_exp.HelloService/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) SayBiStreamHello(ctx context.Context, opts ...grpc.CallOption) (HelloService_SayBiStreamHelloClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_HelloService_serviceDesc.Streams[0], c.cc, "/grpc_exp.HelloService/SayBiStreamHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceSayBiStreamHelloClient{stream}
	return x, nil
}

type HelloService_SayBiStreamHelloClient interface {
	Send(*HelloMsg) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type helloServiceSayBiStreamHelloClient struct {
	grpc.ClientStream
}

func (x *helloServiceSayBiStreamHelloClient) Send(m *HelloMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceSayBiStreamHelloClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for HelloService service

type HelloServiceServer interface {
	SayHello(context.Context, *HelloMsg) (*HelloReply, error)
	SayBiStreamHello(HelloService_SayBiStreamHelloServer) error
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_exp.HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*HelloMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_SayBiStreamHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).SayBiStreamHello(&helloServiceSayBiStreamHelloServer{stream})
}

type HelloService_SayBiStreamHelloServer interface {
	Send(*HelloReply) error
	Recv() (*HelloMsg, error)
	grpc.ServerStream
}

type helloServiceSayBiStreamHelloServer struct {
	grpc.ServerStream
}

func (x *helloServiceSayBiStreamHelloServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceSayBiStreamHelloServer) Recv() (*HelloMsg, error) {
	m := new(HelloMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_exp.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayBiStreamHello",
			Handler:       _HelloService_SayBiStreamHello_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x2f, 0x2a, 0x48, 0x8e, 0x4f, 0xad,
	0x28, 0x50, 0x92, 0xe1, 0xe2, 0xf0, 0x00, 0x49, 0xf8, 0x16, 0xa7, 0x0b, 0x09, 0x70, 0x31, 0xe7,
	0x16, 0xa7, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x4a, 0x2a, 0x5c, 0x5c, 0x60,
	0xd9, 0xa0, 0xd4, 0x82, 0x9c, 0x4a, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12,
	0xa8, 0x12, 0x28, 0xcf, 0xa8, 0x8b, 0x91, 0x8b, 0x07, 0xac, 0x2c, 0x38, 0xb5, 0xa8, 0x2c, 0x33,
	0x39, 0x55, 0xc8, 0x8c, 0x8b, 0x23, 0x38, 0xb1, 0x12, 0x2c, 0x24, 0x24, 0xa4, 0x07, 0xb3, 0x4b,
	0x0f, 0x66, 0x91, 0x94, 0x08, 0x9a, 0x18, 0xd8, 0x78, 0x25, 0x06, 0x21, 0x27, 0x2e, 0x81, 0xe0,
	0xc4, 0x4a, 0xa7, 0xcc, 0xe0, 0x92, 0xa2, 0xd4, 0xc4, 0x5c, 0x92, 0xf5, 0x6b, 0x30, 0x1a, 0x30,
	0x26, 0xb1, 0x81, 0x7d, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x36, 0x7f, 0x9a, 0xf0,
	0x00, 0x00, 0x00,
}
