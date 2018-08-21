package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) Auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var appid string
	var appkey string

	if val, ok := md["user"]; ok {
		appid = val[0]
	}
	if val, ok := md["password"]; ok {
		appkey = val[0]
	}

	if appid != a.User || appkey != a.Password {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}

func NewAuth() []*Authentication {
	return []*Authentication{{"alpha", "123456"}, {"beta", "123456"}}
}
