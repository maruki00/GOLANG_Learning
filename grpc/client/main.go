package main

import (
	pb "grpc/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type helloServer struct {
	pb.GreetServiceServer
}

const (
	port = ":8080"
)

func main() {
	con, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error Client ....")
	}
	defer con.Close()

	// client := pb.NewGreetServiceClient(con)
	// CallSayHello(client)
}
