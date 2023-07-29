package main

import "fmt"

func main() {

	var items = [23]int{12, 12, 24, 345, 3453}
	var infarr = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(items[0])
	fmt.Println(items[2])
	fmt.Println(items[3])
	fmt.Println(items[4])
	fmt.Println(items[1])

	items[0] = 1337
	fmt.Println(items[0])

	fmt.Println(infarr)
}
