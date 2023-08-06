package main

import (
	"fmt"
	"time"
)

func main() {
	go funcWorker1()
	fmt.Println("Hello world!")
	time.Sleep(1 * time.Second)
}

func funcWorker1() {
	go funcWorker2()
	fmt.Println("This is the funcWroker1")
}

func funcWorker2() {
	fmt.Println("This is the funcWorker2")
}

// Output
// -----------------------
// Hello world!
// This is the funcWroker1
// This is the funcWorker2
