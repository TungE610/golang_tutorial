package main

import (
	"fmt"
	"log"

	pb "github.com/TungE610/message-broker/broker/proto"
	"google.golang.org/grpc"
)

func (s *Server) GreetManyTimes(in *pb.BrokerRequest, stream grpc.ServerStreamingServer[pb.BrokerResponse]) error {
	log.Printf("GreetManyTimes function was invoked with %v\n", in)

	for i:= 0; i < 10; i ++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)

		stream.Send(&pb.BrokerResponse{
			Result: res,
		})
		
	}

	return nil
}
