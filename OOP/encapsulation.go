package main

import "fmt"

type Employee struct {
	name      string
	age       int
	address   string
	joined_at string
}

func (e *Employee) getName() string {
	return e.name
}

func (e *Employee) getAge() int {
	return e.age
}

func (e *Employee) getAddress() string {
	return e.address
}

func (e *Employee) getJoinedAt() string {
	return e.joined_at
}

func (e *Employee) setName(name string) {
	e.name = name
}

func (e *Employee) setAge(age int) {
	e.age = age
}

func (e *Employee) setAddress(address string) {
	e.address = address
}

func (e *Employee) setJoinedAt(joined_at string) {
	e.joined_at = joined_at
}

func (e *Employee) print() {
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s\nJoined at: %s\n", e.name, e.age, e.address, e.joined_at)
}

func main() {
	emp1 := &Employee{}
	emp1.setName("Ahmed")
	emp1.setAge(25)
	emp1.setAddress("Moscow, Russia")
	emp1.setJoinedAt("2020-01-01")
	emp1.print()
	fmt.Printf("+++++++++++++++++++++++++++++++++++++++++++\nname : %s, age; %d, address: %s, joined the job at: %s\n", emp1.name, emp1.age, emp1.address, emp1.joined_at)
	fmt.Printf("+++++++++++++++++++++++++++++++++++++++++++\nname : %s, age; %d, address: %s, joined the job at: %s\n", emp1.getName(), emp1.getAge(), emp1.getAddress(), emp1.getJoinedAt())
}
