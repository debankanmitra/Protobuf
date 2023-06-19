package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	gtpb "proto/go_grpc/greet/greet_pb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello World")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	var server gtpb.GreetServiceServer

	grpcServer := grpc.NewServer(opts...)
	gtpb.RegisterGreetServiceServer(grpcServer, server)
	grpcServer.Serve(lis)
}
