package main

import (
	"go-grpc-todo/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":7070"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user.RegisterUserServer(s, &user.Service{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
