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
	fmt.Printf("In server greet function invoked with %v", req)
	firstname := req.GetGreeting().GetFirstName()
	result := "Result: Hello " + firstname + "From server"

	response := gtpb.GreetResponse{
		Result: result,
	}

	// returning response back to client
	return &response, nil
}

func main() {
	fmt.Println("Hello i am the Server")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	//var server gtpb.GreetServiceServer
	//s := &server{}

	grpcServer := grpc.NewServer(opts...)
	gtpb.RegisterGreetServiceServer(grpcServer, &server{}) // service server
	grpcServer.Serve(lis)

}
