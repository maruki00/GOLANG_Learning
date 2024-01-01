package main

import "fmt"

type Dog struct {
	name string
}

type Cat struct {
	name string
}

type Animal interface {
	sleep()
}

func (dog Dog) sleep() {
	fmt.Println("Dog slepping")
}

func (cat Cat) sleep() {
	fmt.Println("Cat slepping")
}

func printInfo(animal Animal) {
	animal.sleep()
}

func main() {
	dog := Dog{name: "dog1"}
	cat := Cat{name: "Cat1"}
	printInfo(dog)
	printInfo(cat)
}
