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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PubsubServerClient is the client API for PubsubServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubsubServerClient interface {
	Publish(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	Subscribe(ctx context.Context, in *String, opts ...grpc.CallOption) (PubsubServer_SubscribeClient, error)
}

type pubsubServerClient struct {
	cc *grpc.ClientConn
}

func NewPubsubServerClient(cc *grpc.ClientConn) PubsubServerClient {
	return &pubsubServerClient{cc}
}

func (c *pubsubServerClient) Publish(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/protobuf.PubsubServer/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubsubServerClient) Subscribe(ctx context.Context, in *String, opts ...grpc.CallOption) (PubsubServer_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubsubServer_serviceDesc.Streams[0], "/protobuf.PubsubServer/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubsubServerSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubsubServer_SubscribeClient interface {
	Recv() (*String, error)
	grpc.ClientStream
}

type pubsubServerSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubsubServerSubscribeClient) Recv() (*String, error) {
	m := new(String)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubsubServerServer is the server API for PubsubServer service.
type PubsubServerServer interface {
	Publish(context.Context, *String) (*String, error)
	Subscribe(*String, PubsubServer_SubscribeServer) error
}

func RegisterPubsubServerServer(s *grpc.Server, srv PubsubServerServer) {
	s.RegisterService(&_PubsubServer_serviceDesc, srv)
}

func _PubsubServer_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubsubServerServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.PubsubServer/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubsubServerServer).Publish(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubsubServer_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(String)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubsubServerServer).Subscribe(m, &pubsubServerSubscribeServer{stream})
}

type PubsubServer_SubscribeServer interface {
	Send(*String) error
	grpc.ServerStream
}

type pubsubServerSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubsubServerSubscribeServer) Send(m *String) error {
	return x.ServerStream.SendMsg(m)
}

var _PubsubServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.PubsubServer",
	HandlerType: (*PubsubServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _PubsubServer_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubsubServer_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protobuf/pubsubserver.proto",
}

func init() {
	proto.RegisterFile("protobuf/pubsubserver.proto", fileDescriptor_pubsubserver_b3096882132dae2b)
}

var fileDescriptor_pubsubserver_b3096882132dae2b = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x28, 0x4d, 0x2a, 0x2e, 0x4d, 0x2a, 0x4e, 0x2d, 0x2a, 0x4b,
	0x2d, 0xd2, 0x03, 0x8b, 0x0a, 0x71, 0xc0, 0x24, 0xa5, 0x44, 0xe1, 0xca, 0x72, 0x2b, 0xd3, 0x8b,
	0x0a, 0x92, 0x21, 0x0a, 0x8c, 0x0a, 0xb8, 0x78, 0x02, 0xc0, 0xda, 0x82, 0xc1, 0xda, 0x84, 0x74,
	0xb9, 0xd8, 0x03, 0x4a, 0x93, 0x72, 0x32, 0x8b, 0x33, 0x84, 0x04, 0xf4, 0x60, 0x5a, 0xf4, 0x82,
	0x4b, 0x8a, 0x32, 0xf3, 0xd2, 0xa5, 0x30, 0x44, 0x84, 0x0c, 0xb9, 0x38, 0x83, 0x4b, 0x93, 0x8a,
	0x93, 0x8b, 0x32, 0x93, 0x52, 0x89, 0xd1, 0x60, 0xc0, 0x98, 0xc4, 0x06, 0x16, 0x32, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xf4, 0x60, 0x4f, 0xed, 0xb8, 0x00, 0x00, 0x00,
}
