package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"

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

type String struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_mygrpc_3026c0e6eb678ef3, []int{0}
}
func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (dst *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(dst, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*String)(nil), "protobuf.String")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MyGrpcClient is the client API for MyGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MyGrpcClient interface {
	Hello(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	StreamServer(ctx context.Context, in *String, opts ...grpc.CallOption) (MyGrpc_StreamServerClient, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (MyGrpc_ClientStreamClient, error)
	Channel(ctx context.Context, opts ...grpc.CallOption) (MyGrpc_ChannelClient, error)
}

type myGrpcClient struct {
	cc *grpc.ClientConn
}

func NewMyGrpcClient(cc *grpc.ClientConn) MyGrpcClient {
	return &myGrpcClient{cc}
}

func (c *myGrpcClient) Hello(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/protobuf.MyGrpc/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myGrpcClient) StreamServer(ctx context.Context, in *String, opts ...grpc.CallOption) (MyGrpc_StreamServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MyGrpc_serviceDesc.Streams[0], "/protobuf.MyGrpc/StreamServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &myGrpcStreamServerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MyGrpc_StreamServerClient interface {
	Recv() (*String, error)
	grpc.ClientStream
}

type myGrpcStreamServerClient struct {
	grpc.ClientStream
}

func (x *myGrpcStreamServerClient) Recv() (*String, error) {
	m := new(String)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *myGrpcClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (MyGrpc_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MyGrpc_serviceDesc.Streams[1], "/protobuf.MyGrpc/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &myGrpcClientStreamClient{stream}
	return x, nil
}

type MyGrpc_ClientStreamClient interface {
	Send(*String) error
	CloseAndRecv() (*String, error)
	grpc.ClientStream
}

type myGrpcClientStreamClient struct {
	grpc.ClientStream
}

func (x *myGrpcClientStreamClient) Send(m *String) error {
	return x.ClientStream.SendMsg(m)
}

func (x *myGrpcClientStreamClient) CloseAndRecv() (*String, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(String)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *myGrpcClient) Channel(ctx context.Context, opts ...grpc.CallOption) (MyGrpc_ChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MyGrpc_serviceDesc.Streams[2], "/protobuf.MyGrpc/Channel", opts...)
	if err != nil {
		return nil, err
	}
	x := &myGrpcChannelClient{stream}
	return x, nil
}

type MyGrpc_ChannelClient interface {
	Send(*String) error
	Recv() (*String, error)
	grpc.ClientStream
}

type myGrpcChannelClient struct {
	grpc.ClientStream
}

func (x *myGrpcChannelClient) Send(m *String) error {
	return x.ClientStream.SendMsg(m)
}

func (x *myGrpcChannelClient) Recv() (*String, error) {
	m := new(String)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyGrpcServer is the server API for MyGrpc service.
type MyGrpcServer interface {
	Hello(context.Context, *String) (*String, error)
	StreamServer(*String, MyGrpc_StreamServerServer) error
	ClientStream(MyGrpc_ClientStreamServer) error
	Channel(MyGrpc_ChannelServer) error
}

func RegisterMyGrpcServer(s *grpc.Server, srv MyGrpcServer) {
	s.RegisterService(&_MyGrpc_serviceDesc, srv)
}

func _MyGrpc_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyGrpcServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.MyGrpc/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyGrpcServer).Hello(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyGrpc_StreamServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(String)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MyGrpcServer).StreamServer(m, &myGrpcStreamServerServer{stream})
}

type MyGrpc_StreamServerServer interface {
	Send(*String) error
	grpc.ServerStream
}

type myGrpcStreamServerServer struct {
	grpc.ServerStream
}

func (x *myGrpcStreamServerServer) Send(m *String) error {
	return x.ServerStream.SendMsg(m)
}

func _MyGrpc_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MyGrpcServer).ClientStream(&myGrpcClientStreamServer{stream})
}

type MyGrpc_ClientStreamServer interface {
	SendAndClose(*String) error
	Recv() (*String, error)
	grpc.ServerStream
}

type myGrpcClientStreamServer struct {
	grpc.ServerStream
}

func (x *myGrpcClientStreamServer) SendAndClose(m *String) error {
	return x.ServerStream.SendMsg(m)
}

func (x *myGrpcClientStreamServer) Recv() (*String, error) {
	m := new(String)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MyGrpc_Channel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MyGrpcServer).Channel(&myGrpcChannelServer{stream})
}

type MyGrpc_ChannelServer interface {
	Send(*String) error
	Recv() (*String, error)
	grpc.ServerStream
}

type myGrpcChannelServer struct {
	grpc.ServerStream
}

func (x *myGrpcChannelServer) Send(m *String) error {
	return x.ServerStream.SendMsg(m)
}

func (x *myGrpcChannelServer) Recv() (*String, error) {
	m := new(String)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MyGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.MyGrpc",
	HandlerType: (*MyGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _MyGrpc_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamServer",
			Handler:       _MyGrpc_StreamServer_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStream",
			Handler:       _MyGrpc_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Channel",
			Handler:       _MyGrpc_Channel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protobuf/mygrpc.proto",
}

func init() { proto.RegisterFile("protobuf/mygrpc.proto", fileDescriptor_mygrpc_3026c0e6eb678ef3) }

var fileDescriptor_mygrpc_3026c0e6eb678ef3 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0xcf, 0xad, 0x4c, 0x2f, 0x2a, 0x48, 0xd6, 0x03, 0xf3, 0x85, 0x38,
	0x60, 0xc2, 0x4a, 0x72, 0x5c, 0x6c, 0xc1, 0x25, 0x45, 0x99, 0x79, 0xe9, 0x42, 0x22, 0x5c, 0xac,
	0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0xd1, 0x15,
	0x46, 0x2e, 0x36, 0xdf, 0x4a, 0xf7, 0xa2, 0x82, 0x64, 0x21, 0x6d, 0x2e, 0x56, 0x8f, 0xd4, 0x9c,
	0x9c, 0x7c, 0x21, 0x01, 0x3d, 0x98, 0x76, 0x3d, 0x88, 0x5e, 0x29, 0x0c, 0x11, 0x21, 0x13, 0x2e,
	0x9e, 0xe0, 0x92, 0xa2, 0xd4, 0xc4, 0xdc, 0xe0, 0xd4, 0xa2, 0xb2, 0xd4, 0x22, 0x62, 0xf4, 0x18,
	0x30, 0x82, 0x74, 0x39, 0xe7, 0x64, 0xa6, 0xe6, 0x95, 0x40, 0xf4, 0x12, 0xa3, 0x4b, 0x83, 0x51,
	0xc8, 0x90, 0x8b, 0xdd, 0x39, 0x23, 0x31, 0x2f, 0x2f, 0x35, 0x87, 0x38, 0x0d, 0x06, 0x8c, 0x49,
	0x6c, 0x60, 0x41, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x83, 0x81, 0xd0, 0x20, 0x01,
	0x00, 0x00,
}
