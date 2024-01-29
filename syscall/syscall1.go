package main

import (
	"fmt"
	"log"
	"syscall"
)

func main() {
	fmt.Println("hello world!")
	filePath := "/tmp/syscallOpen.txt"
	fileFlag := syscall.O_RDWR | syscall.O_CREAT
	fileMode := uint32(0666)

	file, err := syscall.Open(filePath, fileFlag, fileMode)
	if err != nil {
		log.Fatal("Could Not Create File.")
	}
	defer syscall.Close(file)
	_, err = syscall.Write(file, []byte("Hello world"))
	if err != nil {
		log.Fatal("Could not write to file")
	}
	var data []byte
	_, err = syscall.Read(file, data)
	if err != nil {
		log.Fatal("Could not read from file")
	}

	fmt.Println(string(data))

}
