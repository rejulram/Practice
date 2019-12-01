package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"strconv"

	"github.com/rejulram/Practice/gRPCpractice/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("\nGreet function is invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (s *server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("\nGreetManyTimes function is invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("Long Greet function is invoked ")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//End of Client streaming
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalln("Error while recieving client stream request ", err)
		}
		result += "Hello " + req.GetGreeting().GetFirstName() + "!\n"
	}

}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Println("GreetEveryone function is invoked ")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Println("Error while receiving client request", err)
			return err
		}
		firstname := req.GetGreeting().GetFirstName()
		result := "Hello " + firstname + "!"
		err = stream.Send(&greetpb.GreetEveryoneResponse{
			Result: result,
		})
		if err != nil {
			log.Fatalln("Error while send request to client", err)
			return err
		}

	}

}

func main() {
	fmt.Println("Initializing Server")
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
