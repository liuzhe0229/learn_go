package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learn_go/grpc/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())

	if err != nil {
		panic("Client failed to connect the service: " + err.Error())
	}

	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "zhe.liu"})

	if err != nil {
		panic("Failed to call the SayHello function: " + err.Error())
	}

	fmt.Println(r.Message)
}
