package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Scan(&name)
	if name == "john" {
		fmt.Println("You are granted")
	} else if name == "John" {
		fmt.Println("You are granted")
	} else {
		fmt.Println("You are not allowed")
	}

	switch name {
	case "john":
		fmt.Println("Right.")
	default:
		fmt.Println("You are not allowed")
	}
}
