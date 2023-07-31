package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	p := person{name: "abdullah", age: 27}

	fmt.Println("The name is : ", p.name)
	fmt.Println("The age is : ", p.age)
}
