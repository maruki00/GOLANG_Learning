package main

import (
	"fmt"
)

type person struct {
	age  int
	name string
}

func main() {

	var st *person = &person{age: 10, name: "helloworld"}
	fmt.Printf("Hello world %d\n", 1)
	print(st.name)
}
