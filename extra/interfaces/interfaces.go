package interfaces

import "fmt"

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Animal interface {
	Sleep()
}

func (dog Dog) Sleep() {
	fmt.Println(dog.Name, " Dog slepping")
}

func (cat Cat) Sleep() {
	fmt.Println(cat.Name, " Cat slepping")
}

func PrintInfo(animal Animal) {
	animal.Sleep()
}
