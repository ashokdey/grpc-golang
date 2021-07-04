package main

import (
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
	log.Printf("Created a client: %f", c)
}
