package main

import (
	"fmt"
	"runtime"
	"time"
)

func doSomething(done <-chan interface{}) {
	print("hello go routine")

}
func main() {
	fmt.Println("hello world")
	//randInteger := func() interface{} { return rand.Intn(10000) }
	start := time.Now()
	fmt.Println(runtime.NumCPU(), start)
	go doSomething(make(chan interface{}))
}
