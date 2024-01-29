package main

import "fmt"

type person struct {
	name string
	age  int
}

type address struct {
	country string
	city    string
	zipCode int
}

func main() {

	p := person{name: "abdullah", age: 27}
	a := address{}

	a.country = "abcdefg"
	a.city = "abcde"
	a.zipCode = 12344

	fmt.Println("The name is : ", p.name)
	fmt.Println("The age is : ", p.age)

	fmt.Printf("from %s, %s, %d\n", a.country, a.city, a.zipCode)
}
