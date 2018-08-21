package client

import (
	"context"
	"fmt"
	"log"

	"mygrpc/protobuf"

	"google.golang.org/grpc"

	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"strconv"
	"time"

	"google.golang.org/grpc/credentials"
)

type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
func Run() {
	conn := Conn()
	defer conn.Close()
	client := protobuf.NewMyGrpcClient(conn)
	//双向直接返回
	s, err := client.Hello(context.Background(), &protobuf.String{Value: "hi"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("直接返回：", s.GetValue())
	//服务端流
	s1, err := client.StreamServer(context.Background(), &protobuf.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	for {
		r1, err := s1.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("服务端流：", r1.GetValue())
		if err != nil {
			log.Fatal(err)
		}
	}
	//客户端流
	s2, err := client.ClientStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		s2.Send(&protobuf.String{Value: "hi" + strconv.Itoa(i)})
	}
	time.Sleep(time.Second)
	reply, err := s2.CloseAndRecv()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("客户端流：", reply.String())
	//双向流
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			if err := stream.Send(&protobuf.String{Value: "hi"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println("双向流：", reply.GetValue())
	}
}

func Conn() *grpc.ClientConn {
	auth := &Authentication{
		User:     "alpha",
		Password: "1234567",
	}
	certificate, err := tls.LoadX509KeyPair("key/client.crt", "key/client.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("key/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "server.io", // NOTE: this is required!
		RootCAs:      certPool,
	})
	conn, err := grpc.Dial(":1234", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(auth))
	if err != nil {
		return Conn()
	}
	return conn
}
