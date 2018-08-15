package server

import (
	"context"
	"log"
	"net"

	"mygrpc/protobuf"

	"io"

	"strconv"

	"time"

	"google.golang.org/grpc"
)

type Hel int

func (h *Hel) Hello(ctx context.Context, args *protobuf.String) (*protobuf.String, error) {
	reply := &protobuf.String{Value: "hello:" + args.GetValue()}
	time.Sleep(time.Second)
	return reply, nil
}
func (h *Hel) StreamServer(args *protobuf.String, stream protobuf.MyGrpc_StreamServerServer) error {
	rs := make([]*protobuf.String, 0, 10)
	lenrs := cap(rs)
	if args.GetValue() == "hi" {
		for i := 0; i < lenrs; i++ {
			rs = append(rs, &protobuf.String{Value: "hello" + strconv.Itoa(i) + "hi"})
		}
	}
	if args.GetValue() == "hello" {
		for i := 0; i < lenrs; i++ {
			rs = append(rs, &protobuf.String{Value: "hello" + strconv.Itoa(i) + "world"})
		}
	}
	for k := range rs {
		time.Sleep(time.Second)
		if err := stream.Send(rs[k]); err != nil {
			return err
		}
	}
	return nil
}
func (h *Hel) ClientStream(stream protobuf.MyGrpc_ClientStreamServer) error {
	reply := &protobuf.String{}
	s := ""
	for {
		rec, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(reply)
		}
		if err != nil {
			return err
		}
		s += rec.GetValue() + "+"
		reply = &protobuf.String{Value: "hello:" + s}
	}
	return nil
}
func (h *Hel) Channel(stream protobuf.MyGrpc_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		reply := &protobuf.String{Value: "hello:" + args.GetValue()}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func Run() {
	grpcServer := grpc.NewServer()
	protobuf.RegisterMyGrpcServer(grpcServer, new(Hel))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("启动grpc，端口：1234")
	log.Fatal(grpcServer.Serve(lis))
}
