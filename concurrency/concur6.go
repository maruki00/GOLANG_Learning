package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("... This is the Main funmction ...")
	ch := make(chan int, 4)
	ch <- 17
	ch <- 11
	ch <- 13
	ch <- 10
	close(ch)
	go sendToProcess(ch)

	time.Sleep(1 * time.Second)
}

func sendToProcess(ch chan int) {
	sum := 0
	for item := range ch {
		sum += item
	}
	fmt.Printf("%d \n", sum)
	time.Sleep(1 * time.Second)
}
