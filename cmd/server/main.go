package main

import (
	"fmt"
	"net"

	service "github.com/manhdaovan/test_proto/service/service1"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	sv := service.EchoService{}
	service.RegisterEchoServer(server, &sv)
	l, _ := net.Listen("tcp", ":8888")

	fmt.Println("--- start serving --- ")
	if err := server.Serve(l); err != nil {
		panic(err)
	}
}
