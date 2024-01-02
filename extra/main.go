package main

import (
	"extra/interfaces"
	"fmt"
	"mainPackage"
	"strings"
)

func goPackage() {
	fmt.Println("Go Package: ", mainPackage.SayHello())
	mainPackage.ToArray()
}
func dataTypes() {
	var age int
	var name string
	name = "abdullah"
	age = 12344
	var id int64
	var uuid string
	id = 123456
	uuid = "hell-world-1234-123-14"
	fmt.Println("DataTypes: ", id, name, age, uuid)
}
func arithmics() {
	num1 := 10
	num2 := 27
	fmt.Println("Arithmic: ", num1+num2, num1*num2, num1-num2, num1/num2, num1>>num2, num1<<num2)
	fmt.Println("Arithmic: ", num1 > num2, num1 < num2, num1%num2, num1 >= num2, num1 <= num2, num1 == num2)
	fmt.Println("Logic:    ", num1&num2, num1|num2, num1 == num2, num1^num2, num1 != num2, num1|num2)
}
func contriolFlow() {
	fmt.Print("Flow Control: ")
	age := 100
	if age > 120 {
		print(" greater")
	} else {
		print(" less")
	}

	switch age {
	case 100:
		println(" 100")
		break
	case 123:
		println(" 123")
		break
	default:
		println(" Notfound")
	}
	println()
}
func goFunc(var1 *int, var2 int) (int, int) {
	println("Go Func ")
	result := *var1 + var2
	var1 = &var2
	return *var1, result
}
func goArray() []int {
	println("Single D array:")
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range array {
		print("index: ", i-1, "| value: ", i)
	}
	println("")
	println("Multi D array")
	array2 := [][]int{{1, 2}, {2, 3}, {3, 4}}
	for _, item := range array2 {
		for _, k := range item {
			fmt.Println("item: ", item, " | K: ", k)
		}
	}
	return array[:10]
}
func goSlices() []int {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	slice := array[3:7]
	fmt.Println("Slices: ", slice)
	return slice
}
func goString() string {
	var myName string
	myName = "abdullah"
	println("String : ", strings.Split(myName, ","), strings.Count(myName, "a"), strings.HasPrefix(myName, "ab"))
	return myName
}
func goMap() map[string]any {
	data := make(map[string]any)
	data["name"] = "abdullah"
	data["age"] = 12344
	fmt.Println("Map:", data)
	return data
}
func goStruct() {
	type Address struct {
		street int
	}
	type Person struct {
		name    string
		age     int
		address Address
	}

	p1 := Person{name: "nbababa", age: 1245, address: Address{street: 10}}
	fmt.Println("Struct: ", p1, p1.name)
}
func goInterfaces() {
	println("Interfaces: Done")
}
func goErrors(x int, y int) (int, error) {
	fmt.Println("Errors: ")
	// panic("Error .....")
	defer func() {
		recover()
	}()
	return x / y, nil
}

// Main
func main() {
	fmt.Println("----------[Started]------------")
	goPackage()
	dataTypes()
	arithmics()
	contriolFlow()
	var var1 int = 10
	println(goFunc(&var1, 2))
	goArray()
	goSlices()
	goString()
	goMap()
	goStruct()
	goInterfaces()
	defer goErrors(1.0, 1)
	fmt.Println(goErrors(10.0, 10.0))
	dog := interfaces.Dog{Name: "dog1"}
	interfaces.PrintInfo(&dog)
}
