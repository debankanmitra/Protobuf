package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	gtpb "proto/go_grpc/greet/greet_pb"

	"google.golang.org/grpc"
)

type server struct {
	gtpb.UnimplementedGreetServiceServer
}

//var s gtpb.GreetServiceServer

func (*server) Greet(ctx context.Context, req *gtpb.GreetRequest) (*gtpb.GreetResponse, error) {
	//return &gtpb.GreetResponse{Message: "Hello again " + in.GetName()}, nil
	firstname := req.GetGreeting().GetFirstName()
	result := "Hello" + firstname

	response := gtpb.GreetResponse{
		Result: result,
	}

	return &response, nil
}

func main() {
	fmt.Println("Hello World")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	//var server gtpb.GreetServiceServer
	//s := &server{}

	grpcServer := grpc.NewServer(opts...)
	gtpb.RegisterGreetServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)

}
