package main

import "fmt"

func main() {
	var items []int = []int{1, 2, 3, 4, 5, 6, 7}

	for index, item := range items {
		fmt.Println("index", index, "value", item)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("index", i)
		continue
	}

	for {
		fmt.Println("Infi loop")
		break
	}
}
