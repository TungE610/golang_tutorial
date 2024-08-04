package main

import (
	"context"
	"io"
	"log"

	pb "github.com/TungE610/message-broker/broker/proto"
)

func doGreetManyTimes(c pb.BrokerServiceClient) {
	log.Printf("doGreetManyTimes was invoked")

	req := &pb.BrokerRequest{
		FirstName: "Tung Bui",

	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while reading the stream")
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break;
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}