package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ashokdey/grpc-golang/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

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
