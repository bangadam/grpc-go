package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/bangadam/grpc-go/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	doBiDiStreaming(c)
}

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	requests := []*calculatorpb.FindMaximumRequest{
		{
			Number: 1,
		},
		{
			Number: 5,
		},
		{
			Number: 3,
		},
		{
			Number: 6,
		},
		{
			Number: 2,
		},
		{
			Number: 20,
		},
	}

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("could not find maximum: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			fmt.Printf("Sending req: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Printf("Received EOF, close the stream")
				break
			}
			if err != nil {
				log.Fatalf("error while receiving response: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}

		close(waitc)
	}()

	<-waitc
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	requests := []*calculatorpb.ComputeAveragesRequest{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
		{
			Number: 4,
		},
	}

	stream, err := c.ComputeAverages(context.Background())
	if err != nil {
		log.Fatalf("error while calling ComputeAverages: %v", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %v", err)
	}

	fmt.Printf("Response from ComputeAverages: %v\n", res)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		PrimeNumberDecomposition: &calculatorpb.PrimeNumberDecomposition{
			Num: 150,
		},
	}

	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Decomposition RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == nil {
			fmt.Println(msg.GetResult())
		} else {
			break
		}
	}
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		Calculator: &calculatorpb.Calculator{
			Num1: 10,
			Num2: 20,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}

	fmt.Printf("Response from Sum: %v", res.Result)
}
