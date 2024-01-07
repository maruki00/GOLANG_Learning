package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const SERVER_TYPE = "tcp"

func main() {

	host := os.Args[1]
	port := os.Args[2]
	client, err := net.Dial(SERVER_TYPE, host+":"+port)
	if err != nil {
		log.Fatal("Server Error")
	}
	defer client.Close()

	var message string
	fmt.Print("Enter a message: ")
	fmt.Scan(&message)
	client.Write([]byte(message))

}
