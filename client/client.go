package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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
	// MakeUnaryRequest(c)

	// handle server streaming
	// DoServerStreaming(c)

	// make a client streaming
	DoClientStreaming(c)
}

func MakeUnaryRequest(c greetpb.GreetServiceClient) {
	log.Println("Starting Unary rpc...")

	// create the request
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Greet rpc : %v", err.Error())
	}
	log.Printf("Response from Greet rpc: %v", res.Result)
}

func DoServerStreaming(c greetpb.GreetServiceClient) {
	log.Println("Starting server streaming rpc...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to call read stream: %v", err.Error())
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// end of the stream
			break
		}

		if err != nil {
			log.Fatalf("Error while receiving stream : %v", err.Error())
		}

		log.Printf("Response: %v", msg.GetResult())
	}
}

func DoClientStreaming(c greetpb.GreetServiceClient) {
	log.Println("Starting client streaming rpc...")

	// create array of LongGreetings
	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
				LastName:  "Doe",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Alice",
				LastName:  "Ma",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Cris",
				LastName:  "Evan",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Robert",
				LastName:  "Carpenter",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Zuul",
				LastName:  "Dev",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error in client stream: %v", err.Error())
	}

	// iterate and send LongGreetings
	for _, req := range requests {
		log.Printf("Sending a long greeting")
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving stream response: %v", err.Error())
	}

	log.Printf("Response from stream:")
	log.Println(res)

}
