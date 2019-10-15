package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rejulram/Practice/gRPCpractice/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function is invoked with %v", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello" + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello World")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("Not able to listen on port", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve", err)
	}
}
