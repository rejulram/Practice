package main

import (
	"context"
	"fmt"
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
	fmt.Printf("Created Client: %f", c)
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
	log.Printf("Response from server greet : %v", res.Result)
}
