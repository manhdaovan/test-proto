package main

import (
	"context"
	"fmt"

	service "github.com/manhdaovan/test_proto/service/service2"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client1 := service.NewEchoClient(conn)
	in := service.EchoIn{
		Msg:   "message from client 2",
		Order: 2,
	}

	fmt.Println("--- start request client2 --- ")
	out, err := client1.Echo(context.Background(), &in)
	if err != nil {
		panic(err)
	}
	fmt.Printf("out ------ %+v\n", out)
}
