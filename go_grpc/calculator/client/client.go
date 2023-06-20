package main

import (
	"context"
	"fmt"
	"log"

	pb "proto/go_grpc/calculator/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("HELLO I AM THE CLIENT")

	// Creating the client
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("couldn't connect", err)
	}
	defer conn.Close()
	client := pb.NewCalcServiceClient(conn)
	fmt.Println("created client", client)

	// sending RPC Request
	DouUnary(client)
}

func DouUnary(client pb.CalcServiceClient) {

	res, err := client.Calculate(context.Background(), &pb.SumRequest{
		Num1: 60,
		Num2: 10,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("REsponse from greet server: %s", res)
}
