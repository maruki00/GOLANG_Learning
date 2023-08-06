package main

import (
	"fmt"
	"time"
)

func main() {
	go workerfunc1()
	fmt.Println("Hello world .....")
	// if we comment the following line the gorutine will not work
	// because go will not waiting for workerfunc1 finishing
	time.Sleep(1 * time.Second)
}

func workerfunc1() {
	fmt.Println("..In Go Routine ....")
}
