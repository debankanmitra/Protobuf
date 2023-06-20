package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "proto/go_grpc/calculator/protobuf"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalcServiceServer
}

// RPC response to client
/*
*The Calculate method is automatically invoked by the gRPC server when a client makes a request
*to the server's endpoint for the Calculate RPC method. The gRPC server handles the request and
*invokes the corresponding method implemented by the server struct, in this case, the Calculate method.
*The method processes the request, performs the necessary calculations, and returns the response back to the client.

*So, even though the Calculate method is defined outside the main function, it is registered with the gRPC server
*and automatically called when the server receives a request for the Calculate RPC method.
 */
func (*server) Calculate(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	fmt.Printf("In server Calculate function invoked with %v", req)
	sum := req.GetNum1() + req.GetNum2()

	result := &pb.SumResponse{
		Result: sum,
	}

	return result, nil
}

func main() {

	fmt.Println("I am the gRPC calculation server")

	// Starting the server
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalcServiceServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}
