package main

import (
	"context"
	"fmt"
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
	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	// Unary call
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
