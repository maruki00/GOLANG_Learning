package main

import (
	myPkg "delivery/golang_learning/v2/fondamentals"
)

func main() {
	//Accessible
	println(myPkg.GlobalVar)
	println(myPkg.GlobalFunc())

	//Not Accessible
	// println(myPkg.localVar)
	// println(myPkg.localFunc())
}
