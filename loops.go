package main

import (
	"fmt"
)

func main() {
	fmt.Println("starting Coding ....")
	fmt.Println("For loop")
	var y = 0
	for i := 0; i < 10000000; i++ {
		//fmt.Printf("Index : %d \n", i)
		var res = i * 123 * 123
		y = res
		res = y
		fmt.Printf(".")
	}

	fmt.Println("While Loop")

	// for true {
	// 	fmt.Println("Hello world")
	// }
}
