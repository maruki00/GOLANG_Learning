package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	args := os.Args[1:]

	for _, arg := range args {
		result, err := exec.Command(arg).Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(result))
	}

	fmt.Println("hello world...")
}
