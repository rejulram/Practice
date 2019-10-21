package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/rejulram/Practice/gRPCpractice/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client starting")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Unable to connect server ", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	//doUnary(c)
	//doServerStraming(c)
	doClientStreaming(c)
}

// Unary call
func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Sum Unary RPC ")
	req := &calculatorpb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalln("Sum gRPC call failed ", err)
	}
	log.Println("Sum =", res.GetResult())
}

func doServerStraming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do PrimeDecomposition server streaming RPC")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: 210,
	}
	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalln("PrimeNumberDecomposition RPC call failed", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error while receiving stream", err)
		}
		log.Println("Prime Decomposition :", msg.GetPrimeNumber())
	}
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Compute Average  client streaming RPC")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalln("Error while calling ComputeAverage RPC", err)
	}

	/*requests := []*calculatorpb.ComputeAverageRequest{
		&calculatorpb.ComputeAverageRequest{
			Number: 10,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 15,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 20,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 25,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 30,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 35,
		},
	}

	for _, req := range requests {
		fmt.Println("Sending request", req)
		if err := stream.Send(req); err != nil {
			log.Fatalln("Error while client streaming", err)
		}
		time.Sleep(1000 * time.Millisecond)
	} */
	numbers := []int32{10, 15, 20, 25, 30, 35}
	for _, number := range numbers {
		fmt.Println("Sending request", number)
		if err := stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		}); err != nil {
			log.Fatalln("Error while sending request to Server", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalln("Error while recieving reponse from from Server", err)
	}
	fmt.Println("Computed Average recived from server is ", res.GetAverage())
}
