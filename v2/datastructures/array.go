package main

import "fmt"

func main() {
	var ages []int
	var names [5][5]string = [5][5]string{
		{"name", "name", "name", "name", "name"},
		{"name", "name", "name", "name", "name"},
		{"name", "name", "name", "name", "name"},
		{"name", "name", "name", "name", "name"},
		{"name", "name", "name", "name", "name"},
	}

	ages = []int{1, 2, 3, 4, 5, 5}

	fmt.Println("names: ", names)
	fmt.Println("ages: ", ages)
}
