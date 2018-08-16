package server

import (
	"context"
	"mygrpc/protobuf"
	"strings"
	"time"

	"fmt"

	"github.com/docker/docker/pkg/pubsub"
)

type pubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *pubsubService {
	return &pubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *pubsubService) Publish(ctx context.Context, arg *protobuf.String) (*protobuf.String, error) {
	fmt.Println(arg.Value)
	p.pub.Publish(arg.GetValue())
	return &protobuf.String{}, nil
}

func (p *pubsubService) Subscribe(arg *protobuf.String, stream protobuf.PubsubServer_SubscribeServer) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&protobuf.String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}
