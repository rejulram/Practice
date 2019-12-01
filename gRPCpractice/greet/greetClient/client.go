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
	//doClientStreaming(c)

	//Bi Directional streaming
	doBiDiStreaming(c)

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

func doBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting GreetEvery One Client RPC..")
	// Create a stream invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalln("Error while creating client stream", err)
	}
	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Rejul",
				LastName:  "Ramakrishnan",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ramesh",
				LastName:  "Sundaram",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Rupak",
				LastName:  "Behra",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Aniket",
				LastName:  "Mani",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Ananth",
				LastName:  "G",
			},
		},
	}

	waitc := make(chan int)
	//Send Streaming data
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending : %v \n", req.GetGreeting().GetFirstName())
			err := stream.Send(req)
			if err != nil {
				log.Fatalln("Error straming data from Client", err)
				break
			}
		}
		stream.CloseSend()
	}()
	// Recieve Streaming Data
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln("Error while receiving streaming data from server")
				break
			}
			fmt.Println("Received : ", res.GetResult())
		}
		close(waitc)
	}()
	//Block Until send receive complete
	<-waitc
}
