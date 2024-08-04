package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/TungE610/message-broker/broker/proto"
	"google.golang.org/grpc"
)

func (s *Server) LongGreet(stream grpc.ClientStreamingServer[*pb.BrokerRequest, *pb.BrokerResponse]) error {

	log.Println("Long Greet function was invoked")

	res := ""

	for {
		res, err := stream.Recv()

		if err != io.EOF {
			return stream.SendAndClose(&pb.BrokerResponse {
				Result: res,
			})
		}

		if err != nil {
			log.Fatal("Error while reading client stream: %v\n", err)
		}

		res += fmt.Sprintf("Hello %s\n", res.FirstName)
	}
}
