package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/ashokdey/grpc-golang/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Println("Invoking Greet...")
	// extract the req fields
	firstName := req.GetGreeting().GetFirstName()

	// create pointer to greet response
	greet := &greetpb.GreetResponse{
		Result: "Hello " + firstName,
	}

	// return the pointer to the greet response
	return greet, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	log.Println("Invoking GreetManyTimes...")
	firstName := req.GetGreeting().GetFirstName()

	for i := 0; i < 10; i++ {
		log.Printf("I = %v", i)
		result := "Hello " + firstName + " number " + strconv.Itoa(i) + "\n"
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		err := stream.Send(res)
		if err != nil {
			log.Fatalf("Error while streaming: %v", err.Error())
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func main() {
	fmt.Println("Hello gRPC")

	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Printf("Cannot listen : %v", err.Error())
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Printf("Failed to server: %v", err.Error())
	}
}
