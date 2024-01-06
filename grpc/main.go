package main

import (
	pb "grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer
}

const (
	port = ":8080"
)

func main() {
	listenner, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error Net")
	}
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(listenner); err != nil {
		log.Fatal("GRPC Error")
	}
}
