package main

import (
	"context"
	"log"

	pb "github.com/TungE610/message-broker/broker/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.BrokerRequest) (*pb.BrokerResponse, error) {
	log.Printf("Broker function was invoked with %v\n", in)

	return &pb.BrokerResponse {
		Result: "Hello " + in.FirstName,
	}, nil
	
}
