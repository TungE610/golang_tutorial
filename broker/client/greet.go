package main

import (
	"context"
	"log"

	pb "github.com/TungE610/message-broker/broker/proto"
)

func doGreet(c pb.BrokerServiceClient) {
	log.Printf("doGreet get invoked")

	res, err := c.Greet(context.Background(), &pb.BrokerRequest{
		FirstName: "Tung Bui",
	})

	if err != nil {
		log.Fatalf("Can not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}