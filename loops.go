package main

import (
	"fmt"
)

func main() {
	fmt.Println("starting Coding ....")
	fmt.Println("For loop")
	var y = 0
	arr := [...]int{1, 2, 3, 5, 6, 7, 4, 7}
	for i := 0; i < 10; i++ {
		//fmt.Printf("Index : %d \n", i)
		var res = i * 123 * 123
		y = res
		res = y
		fmt.Printf(".")
	}

	fmt.Println("\nWhile Loop\n")

	// for true {
	// 	fmt.Println("Hello world")
	// }

	for index, value := range arr {
		fmt.Printf("\nindex %d , value %d", index, value)
	}
	fmt.Print("\n")
}
