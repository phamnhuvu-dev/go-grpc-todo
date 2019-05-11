package main

import (
	"context"
	"go-grpc-todo/user"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const (
	port = ":7070"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))

	user.RegisterUserGRPCServer(server, &user.Service{})
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func unaryInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	method := info.FullMethod
	if method == "/UserGRPC/Login" || method == "/UserGRPC/Register" {
		log.Println("Don't check")
	} else {
		log.Println("Check")
	}

	h, err := handler(ctx, req)

	log.Printf("Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}
