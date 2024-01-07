package main

import (
	"fmt"
	"log"
	"net"
)

const (
	SERVER_TYPE = "tcp"
	SERVER_HOST = "127.0.0.1"
	SERVER_PORT = "9090"
)

func main() {

	server, err := net.Listen(SERVER_TYPE, fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT))
	if err != nil {
		log.Fatal("Error Server")
	}
	for {
		con, err := server.Accept()
		if err != nil {
			log.Fatal("Error Accepting new Client.")
		}
		go handelClient(con)
	}
}

func handelClient(con net.Conn) {
	var message string
	buff := make([]byte, 1024)
	defer con.Close()
	for {
		len, err := con.Read(buff)
		if err != nil {
			log.Fatal("Error read from client.")
		}
		fmt.Println("From Client: ", string(buff[:len]))
		fmt.Print("Enter a message: ")
		fmt.Scan(&message)
		con.Write([]byte(message))
	}
}
