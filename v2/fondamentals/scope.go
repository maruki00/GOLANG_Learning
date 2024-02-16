package main

import myPackage "fondamentals"

func main() {
	//Accessible
	println(myPackage.GlobalVar)
	println(myPackage.GlobalFunc())

	//Not Accessible
	// println(myPkg.localVar)
	// println(myPkg.localFunc())
}
