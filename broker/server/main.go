package main

import (
	"log"
	"net"

	pb "github.com/TungE610/message-broker/broker/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

type Server struct {
	pb.BrokerServiceServer
	pb.SumServiceServer
}



func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Panic("Failed to listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterBrokerServiceServer(s, &Server{})
	pb.RegisterSumServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to server %v\n", err)
	}
}