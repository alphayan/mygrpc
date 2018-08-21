package server

import (
	"context"
	"log"
	"net"

	"mygrpc/protobuf"

	"io"

	"strconv"

	"time"

	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"
)

type g int

func (*g) Hello(ctx context.Context, args *protobuf.String) (*protobuf.String, error) {
	reply := &protobuf.String{Value: "hello:" + args.GetValue()}
	time.Sleep(time.Second)
	return reply, nil
}
func (*g) StreamServer(args *protobuf.String, stream protobuf.MyGrpc_StreamServerServer) error {
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
func (*g) ClientStream(stream protobuf.MyGrpc_ClientStreamServer) error {
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
func (*g) Channel(stream protobuf.MyGrpc_ChannelServer) error {
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
	certificate, err := tls.LoadX509KeyPair("key/server.crt", "key/server.key")
	if err != nil {
		log.Fatal(err)
	}
	_ = certificate
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("key/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if !certPool.AppendCertsFromPEM(ca) {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		ClientCAs:    certPool,
	})
	grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(ChainUnaryServer(filter, filter1, filter2)))
	protobuf.RegisterMyGrpcServer(grpcServer, new(g))
	protobuf.RegisterPubsubServerServer(grpcServer, NewPubsubService())
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("启动grpc，端口：1234")
	log.Fatal(grpcServer.Serve(lis))
}
