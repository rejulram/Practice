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
	doServerStraming(c)
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
