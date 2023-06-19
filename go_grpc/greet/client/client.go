package main

import (
	"context"
	"fmt"
	"log"
	gtpb "proto/go_grpc/greet/greet_pb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i am the client")
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", 50051), grpc.WithInsecure()) //WithInsecure disable SSL for grpc , don't use it in production
	//grpc.WithTransportCredentials(insecure.NewCredentials())
	if err != nil {
		log.Fatal("couldn't connect", err)
	}

	client := gtpb.NewGreetServiceClient(conn) //service client

	fmt.Println("created client", client)

	DoUnary(client)

	defer conn.Close()
}

func DoUnary(client gtpb.GreetServiceClient) {
	fmt.Println("Starting Unary RPC")
	req := &gtpb.GreetRequest{
		Greeting: &gtpb.Greet{
			FirstName: "Debankan",
			LastName:  "Mitra",
		},
	}
	res, err := client.Greet(context.Background(), req) // it looks like we are directly invoking greet() on server
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("REsponse from greet server: %s", res.Result)
}
