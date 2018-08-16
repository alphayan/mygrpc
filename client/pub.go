package client

import (
	"context"
	"log"
	"mygrpc/protobuf"

	"google.golang.org/grpc"
)

func Pub() {
	conn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	sub := protobuf.NewPubsubServerClient(conn)
	ctx := context.Background()
	if _, err := sub.Publish(ctx, &protobuf.String{Value: "golang: hello golang"}); err != nil {
		log.Fatal(err)
	}
	if _, err := sub.Publish(ctx, &protobuf.String{Value: "docker: hello docker"}); err != nil {
		log.Fatal(err)
	}
}
