package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

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

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	fmt.Printf("FindMaximum function was invoked with a streaming request")
	maximum := int64(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		number := req.GetNumber()
		if number > maximum {
			maximum = number
			sendErr := stream.Send(&calculatorpb.FindMaximumResponse{
				Result: maximum,
			})

			if sendErr != nil {
				log.Fatalf("Error while sending client stream: %v", sendErr)
				return sendErr
			}
		}
	}
}

func (*server) ComputeAverages(stream calculatorpb.CalculatorService_ComputeAveragesServer) error {
	fmt.Printf("ComputeAverages function was invoked with a streaming request")
	var avg float32
	var count int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.ComputeAveragesResponse{
				Result: avg / float32(count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		sum := req.GetNumber()
		count++
		avg += float32(sum)
	}
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("Decomposition function was invoked with %v", req)
	number := req.GetPrimeNumberDecomposition().GetNum()

	k := int64(2)
	for number > 1 {
		if number%k == 0 {
			number = number / k
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				Result: k,
			}
			stream.Send(res)
		} else {
			k++
		}

		time.Sleep(1000 * time.Millisecond)
	}

	return nil
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
