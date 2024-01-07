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
	server, err := net.Listen(SERVER_TYPE, host+":"+port)
	if err != nil {
		log.Fatal("Server Error")
	}
	defer server.Close()
	fmt.Printf("Server Run on %s:%s.\n", host, port)
	for {
		con, err := server.Accept()
		if err != nil {
			log.Fatal("Error Accept Client")
		}
		go handelClient(con)
	}

}

func handelClient(conn net.Conn) {
	var message string
	buff := make([]byte, 1024)
	for {
		msgLen, err := conn.Read(buff)
		if err != nil {
			log.Fatal("Error read buffer from client!")
		}

		fmt.Println("Recieved from client: ", string(buff[:msgLen]))
		fmt.Print("Enter message: ")
		fmt.Scan(&message)
		if message == "exit" || message == "quit" {
			break
		}
		conn.Write([]byte(message))
	}

	conn.Close()
}
