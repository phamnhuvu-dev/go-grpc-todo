package main

import (
	"context"
	"go-grpc-todo/user"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address     = "localhost:7070"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Login(ctx, &user.UserLoginRequest{Username: "phamnhuvu", Password: "123456789"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
