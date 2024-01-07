package main

import (
	"fmt"
	"log"
	"net"
)

const (
	CLIENT_TYPE = "tcp"
	CLIENT_HOST = "127.0.0.1"
	CLIENT_PORT = "9090"
)

func main() {
	buff := make([]byte, 1024)
	var message string
	client, err := net.Dial(CLIENT_TYPE, CLIENT_HOST+":"+CLIENT_PORT)
	defer client.Close()
	if err != nil {
		log.Fatal("Error from Server")
	}
	for {
		fmt.Print("Enter a message: ")
		fmt.Scan(&message)
		client.Write([]byte(message))
		client.Read(buff)
		fmt.Printf("message from server: %s\n", string(buff))
	}
}
