package server

import (
	"context"

	"fmt"

	"log"

	"errors"

	"google.golang.org/grpc"
)

func filter(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("fileter:", info)
	return handler(ctx, req)
}

func filter1(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	return handler(ctx, req)
}

func filter2(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	as := NewAuth()
	for _, v := range as {
		if err := v.Auth(ctx); err == nil {
			return handler(ctx, req)
		}
	}

	return nil, errors.New("token 验证失败")
}

func filterStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
	log.Println("fileter:", info)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	as := NewAuth()
	for _, v := range as {
		if err := v.Auth(stream.Context()); err == nil {
			return handler(srv, stream)
		}
	}
	return errors.New("token 验证失败")
}
