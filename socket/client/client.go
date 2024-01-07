package client

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
	for {
		con, err := server.Accept()
		if err != nil {
			log.Fatal("Error Accept Client")
		}
		go handelClient(con)
	}

}

func handelClient(conn net.Conn) {
	buff := make([]byte, 1024)
	msgLen, err := conn.Read(buff)
	if err != nil {
		log.Fatal("Error read buffer from client!")
	}
	fmt.Println("Recieved from client: ", string(buff[:msgLen]))
	conn.Close()
}
