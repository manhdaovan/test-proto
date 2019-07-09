package main

import (
	"context"
	"fmt"

	service "github.com/manhdaovan/test_proto/service/service1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client1 := service.NewEchoClient(conn)
	in := service.EchoIn{
		Msg:   "message from client 1",
		Order: 1,
	}

	fmt.Println("--- start request client1 --- ")
	out, err := client1.Echo(context.Background(), &in)
	if err != nil {
		panic(err)
	}
	fmt.Printf("out ------ %+v\n", out)
}
