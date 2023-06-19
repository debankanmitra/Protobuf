package main

import (
	"fmt"
	"log"
	gtpb "proto/go_grpc/greet/greet_pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("Hello i am the client")
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", 50051), grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(nil, ""))) //WithInsecure disable SSL for grpc , don't use it in production
	//grpc.WithTransportCredentials(insecure.NewCredentials())
	if err != nil {
		log.Fatal("couldn't connect", err)
	}

	client := gtpb.NewGreetServiceClient(conn)

	fmt.Println("created client", client)

	defer conn.Close()
}
