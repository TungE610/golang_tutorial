package main

import (
	"context"
	"log"

	pb "github.com/TungE610/message-broker/broker/proto"
)

func sum2Number(c pb.SumServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber: 5,
		SecondNumber: 10,
	})

	if err != nil {
		log.Fatalf("Can not sum up %v\n", err)
	}

	log.Printf("Sum of 2 numbers: %d\n", res.Result)
}