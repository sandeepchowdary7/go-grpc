package main

import (
	"context"
	"log"

	pb "github.com/sandeepchowdary7/go-grpc/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Printf("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		X: 2,
		Y: 4,
	})

	if err != nil {
		log.Fatalf("Something wrong")
	}
	log.Printf("Sum: %d", res.Z)
}
