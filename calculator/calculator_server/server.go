package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/bangadam/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	// Start the server
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Print("Server started")
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v", req)
	firstNumber := req.GetCalculator().GetNum1()
	secondNumber := req.GetCalculator().GetNum2()
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		Result: sum,
	}
	return res, nil
}
