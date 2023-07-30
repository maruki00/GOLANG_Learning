package main

import "fmt"

func main() {
	var age int = 10
	var u_age *int = &age

	if u_age != nil {
		println("It's not null")
		fmt.Printf("ur age is %d\n", *u_age)

		// var u_age *int;
		// *u_age = value  -> asm [u_age]
		// u_age  = is the address -> asm u_age
	}
}
