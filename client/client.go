package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ashokdey/grpc-golang/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("gRPC client is up.")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect: %v", err.Error())
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	// make a unary req
	MakeUnaryRequest(c)
}

func MakeUnaryRequest(c greetpb.GreetServiceClient) {
	log.Println("Starting Unary rpc...")

	// create the request
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FisrtName: "John",
			LastName:  "Doe",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Greet rpc : %v", err.Error())
	}
	log.Printf("Response from Greet rpc: %v", res.Result)
}
