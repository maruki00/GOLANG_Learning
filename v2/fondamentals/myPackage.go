package myPkg

var GlobalVar = "Hello You still can get me outside the package..."

func GlobalFunc() string {
	return "You can access me outside  te package ..."
}

var localVar = "get me outside package if you can "

func localFunc() string {
	return "Aviable only inside the package"
}
