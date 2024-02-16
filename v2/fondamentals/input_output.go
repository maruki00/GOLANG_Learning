package main

import (
	"fmt"
	"log"
)

func main() {
	var name string
	fmt.Print("Enter Your name : ")
	if _, err := fmt.Scanf("%s", &name); err != nil {
		log.Fatal("Error reading ....", err)
	}
	fmt.Println("Hello : ", name)
}
