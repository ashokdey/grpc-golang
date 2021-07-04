package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ashokdey/grpc-golang/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Println("Invoking Greet...")
	// extract the req fields
	firstName := req.GetGreeting().GetFisrtName()

	// create pointer to greet response
	greet := &greetpb.GreetResponse{
		Result: "Hello " + firstName,
	}

	// return the pointer to the greet response
	return greet, nil
}

func main() {
	fmt.Println("Hello gRPC")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Printf("Cannot listen : %v", err.Error())
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Printf("Failed to server: %v", err.Error())
	}
}
