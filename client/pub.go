package client

import (
	"context"
	"log"
	"mygrpc/protobuf"
)

func Pub() {
	sub := protobuf.NewPubsubServerClient(Conn())
	ctx := context.Background()
	if _, err := sub.Publish(ctx, &protobuf.String{Value: "golang: hello golang"}); err != nil {
		log.Fatal(err)
	}
	if _, err := sub.Publish(ctx, &protobuf.String{Value: "docker: hello docker"}); err != nil {
		log.Fatal(err)
	}
}
