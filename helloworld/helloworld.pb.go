// Code generated by protoc-gen-go.
// source: helloworld/helloworld.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	helloworld/helloworld.proto

It has these top-level messages:
	FuckRequest
	FuckReply
	HelloRequest
	HelloReply
*/
package helloworld

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

type FuckRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *FuckRequest) Reset()                    { *m = FuckRequest{} }
func (m *FuckRequest) String() string            { return proto.CompactTextString(m) }
func (*FuckRequest) ProtoMessage()               {}
func (*FuckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FuckRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FuckReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *FuckReply) Reset()                    { *m = FuckReply{} }
func (m *FuckReply) String() string            { return proto.CompactTextString(m) }
func (*FuckReply) ProtoMessage()               {}
func (*FuckReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FuckReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*FuckRequest)(nil), "helloworld.FuckRequest")
	proto.RegisterType((*FuckReply)(nil), "helloworld.FuckReply")
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayFuck(ctx context.Context, in *FuckRequest, opts ...grpc.CallOption) (*FuckReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayFuck(ctx context.Context, in *FuckRequest, opts ...grpc.CallOption) (*FuckReply, error) {
	out := new(FuckReply)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayFuck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayFuck(context.Context, *FuckRequest) (*FuckReply, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayFuck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayFuck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayFuck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayFuck(ctx, req.(*FuckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayFuck",
			Handler:    _Greeter_SayFuck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/helloworld.proto",
}

func init() { proto.RegisterFile("helloworld/helloworld.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x47, 0x30, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0xb8, 0x10, 0x22, 0x4a, 0x8a, 0x5c, 0xdc, 0x6e, 0xa5, 0xc9, 0xd9, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x60, 0xb6, 0x92, 0x2a, 0x17, 0x27, 0x44, 0x49, 0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17, 0x7b,
	0x6e, 0x6a, 0x71, 0x71, 0x62, 0x3a, 0x4c, 0x0d, 0x8c, 0xab, 0xa4, 0xc4, 0xc5, 0xe3, 0x01, 0x32,
	0x17, 0x9f, 0x51, 0x6a, 0x5c, 0x5c, 0x50, 0x35, 0x78, 0xcd, 0x32, 0x6a, 0x63, 0xe4, 0x62, 0x77,
	0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0x12, 0xb2, 0xe3, 0xe2, 0x08, 0x4e, 0xac, 0x04, 0x6b, 0x13,
	0x92, 0xd0, 0x43, 0xf2, 0x0c, 0xb2, 0x6d, 0x52, 0x62, 0x58, 0x64, 0x0a, 0x72, 0x2a, 0x95, 0x18,
	0x84, 0xac, 0xb9, 0xd8, 0x83, 0x13, 0x2b, 0x41, 0x3e, 0x10, 0x12, 0x47, 0x56, 0x84, 0xe4, 0x6d,
	0x29, 0x51, 0x4c, 0x09, 0xb0, 0x66, 0x27, 0x03, 0x2e, 0xe9, 0xcc, 0x7c, 0xbd, 0xf4, 0xa2, 0x82,
	0x64, 0xbd, 0xd4, 0x8a, 0xc4, 0xdc, 0x82, 0x9c, 0xd4, 0x62, 0x24, 0xa5, 0x4e, 0xfc, 0x60, 0x9b,
	0xc2, 0x41, 0xec, 0x00, 0x50, 0xd0, 0x06, 0x30, 0x26, 0xb1, 0x81, 0xc3, 0xd8, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xc0, 0x98, 0xc1, 0x77, 0x82, 0x01, 0x00, 0x00,
}
