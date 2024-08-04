package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/TungE610/message-broker/broker/proto"
)

var addr string = "127.0.0.1:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Fail to connect: %v\n", err)
	}

	defer conn.Close()

	c:= pb.NewBrokerServiceClient(conn)
	s:= pb.NewSumServiceClient(conn)

	// doGreet(c)
	doGreetManyTimes(c)
	sum2Number(s)
}