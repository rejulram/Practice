package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rejulram/Practice/gRPCpractice/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("\n Sum Request received in server : %v", *req)
	first := req.GetFirstNumber()
	second := req.GetSecondNumber()
	result := first + second
	res := &calculatorpb.SumResponse{
		Result: result,
	}
	return res, nil
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
