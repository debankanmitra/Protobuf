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
