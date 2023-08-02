package main

import "fmt"

type Costumer struct {
	name      string
	age       int
	address   string
	joined_at string
}

func (c *Costumer) getName() string {
	return c.name
}

func (c *Costumer) getAge() int {
	return c.age
}

func main() {
	costmer1 := &Costumer{
		name:      "ahmed",
		age:       32,
		address:   "moscow, Russia",
		joined_at: "2023-08-01",
	}
	fmt.Printf("name : %s, age; %d, address: %s, joined the job at: %s\n", costmer1.name, costmer1.age, costmer1.address, costmer1.joined_at)
	fmt.Printf("name : %s\n", costmer1.getName())
	fmt.Printf("age  : %d\n", costmer1.getAge())
}
