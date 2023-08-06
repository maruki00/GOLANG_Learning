package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("... This is the Main funmction ...")
	ch := make(chan int)
	go ids(ch)
	go names(ch)
	for {
		result, isValid := <-ch
		if !isValid {
			fmt.Println("The Channel is closed!")
			break
		}
		fmt.Printf("The value %d, is OK!\n", result)
	}
	time.Sleep(1 * time.Second)
}

func names(ch chan int) {
	fmt.Println("... This is the  Names function ...")
	names := []string{"blahblah1", "blahblah2", "blahblah3", "blahblah"}

	for item := range names {
		fmt.Printf("index %d, id: %d, the name is : %s\n", item, <-ch, names[item])
		time.Sleep(1 * time.Millisecond)
	}
}

func ids(ch chan int) {
	fmt.Println("... This is the Ids function ...")
	ids := [4]int{1, 2, 3, 4}
	for id := range ids {
		ch <- ids[id]
		// fmt.Printf("The id is : %d \n", ids[id])
		time.Sleep(1 * time.Millisecond)
	}
	close(ch)
}
