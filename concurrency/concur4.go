package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("... This is the Main funmction ...")
	//anonymous GO routine ...
	go func names() {
		fmt.Println("... This is the  Names function ...")
		names := []string{"blahblah1", "blahblah2", "blahblah3", "blahblah"}
	
		for item := range names {
			fmt.Printf("the name is : %s\n", names[item])
			time.Sleep(500 * time.Millisecond)
		}
	}()
	time.Sleep(1 * time.Second)
}



func ids() {
	fmt.Println("... This is the Ids function ...")
	ids := [4]int{1, 2, 3, 4}
	for id := range ids {
		fmt.Printf("The id is : %d \n", ids[id])
		time.Sleep(500 * time.Millisecond)
	}
}
