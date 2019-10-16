package main

import (
	"context"
	"fmt"
	"io"
	"log"

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
	doServerStreaming(c)

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
