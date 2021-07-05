package main

import (
	"context"
	"fmt"
	"io"
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
		result := "Hello " + firstName + " number " + strconv.Itoa(i+1) + "\n"
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

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	log.Println("Invoking LongGreet...")

	var result = "Hello "

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			// send the concatenated result
			res := &greetpb.LongGreetResponse{
				Result: result,
			}
			return stream.SendAndClose(res)
		}

		if err != nil {
			log.Fatalf("Error while reading client stream : %v", err.Error())
			return err
		}

		firstName := req.GetGreeting().GetFirstName()
		result += firstName + "! "
	}
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	log.Println("Invoking GreetEveryone...")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while receiving client stream: %v", err.Error())
		}

		firstName := req.GetGreeting().GetFirstName()
		res := "Hello " + firstName + "! "

		err = stream.Send(&greetpb.GreetEveryoneResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error sending stream data to client: %v", err.Error())
			return err
		}
	}
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
