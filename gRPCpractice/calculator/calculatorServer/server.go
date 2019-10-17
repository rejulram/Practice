package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/rejulram/Practice/gRPCpractice/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("\n Sum Request received in server : %v", *req)
	first := req.GetFirstNumber()
	second := req.GetSecondNumber()
	result := first + second
	res := &calculatorpb.SumResponse{
		Result: result,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("\n PrimeNumberDecomposition Request received in server : %v", *req)
	number := req.GetNumber()
	divisor := int32(2)
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.PrimeNumberDecompositionResponse{
				PrimeNumber: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Println("Divisor has incremented", divisor)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	fmt.Println("Server Initializing")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("Can't listern on port ", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln("Unable to server now ", err)
	}

}
