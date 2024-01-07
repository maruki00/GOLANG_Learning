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

	for {
		var message string
		buff := make([]byte, 1024)
		fmt.Print("Enter a message: ")
		fmt.Scan(&message)
		if message == "quit" || message == "exit" {
			break
		}
		client.Write([]byte(message))
		msgLen, err := client.Read(buff)
		if err != nil {
			log.Fatal("Error read buffer from client!")
		}
		fmt.Println("Recieved from server: ", string(buff[:msgLen]))
	}
	defer client.Close()

}
