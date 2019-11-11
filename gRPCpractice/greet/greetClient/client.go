package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/rejulram/Practice/gRPCpractice/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client started")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect server", err)
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created Client: %f", c)
	//Unary gRPC call
	//doUnary(c)

	//Server Steaming gRPC call
	//doServerStreaming(c)

	//Client Streaming gRPC call
	doClientStreaming(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	//Unary gRPC Call
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rejul",
			LastName:  "Ramakrishnan",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalln("Error while calling GRPC greet", err)
	}
	log.Printf("\nResponse from server greet : %v\n", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rejul",
			LastName:  "Ramakrishnan",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalln("Error while calling GRPC GreetManyTimes ", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error while calling GRPC GreetManyTimes ", err)
		}
		log.Println("Response from GreetManyTimes : ", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Rejul",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Anusree",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sikha",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Lakshmi",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ramakrishnan",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalln("Error while calling Long greet", err)
	}

	//Iterrate over slice to send each message individually
	for _, req := range requests {
		fmt.Println("Sending request", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error while recieving response for Long Greet request", err)
	}
	fmt.Println("Long Greeting Response Recieved :", res.GetResult())
}
