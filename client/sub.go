package client

import (
	"context"
	"fmt"
	"io"
	"log"
	"mygrpc/protobuf"
)

func Sub() {
	client := protobuf.NewPubsubServerClient(Conn())
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
