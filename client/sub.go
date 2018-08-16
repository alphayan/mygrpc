package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"mygrpc/protobuf"

	"google.golang.org/grpc"
)

func Sub() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := protobuf.NewPubsubServerClient(conn)
	stream, err := client.Subscribe(context.Background(), &protobuf.String{Value: "golang:"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(reply.GetValue())
	}

}
