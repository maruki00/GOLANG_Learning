package myPackage

/*
var/func Start with lowercase thats mean that is a local variable like private
other wise they are a global you can access them outside the package
*/
var GlobalVar = "Hello You still can get me outside the package..."

func GlobalFunc() string {
	return "You can access me outside  te package ..."
}

var localVar = "get me outside package if you can "

func localFunc() string {
	return "Aviable only inside the package"
}
