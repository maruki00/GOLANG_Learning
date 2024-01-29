package main

import (
	"fmt"
)

func func1() string {
	return "hello from functions\n"
}
func func2(name string, age int) string {
	return fmt.Sprintf("Hello %s, you are %d years old!\n", name, age)
}

func main() {
	fmt.Print("Start Coding\n")
	var res1 string
	var res2 string = func2("abdullah", 27)
	res1 = func1()

	fmt.Print(res1)
	fmt.Print(res2)

}
