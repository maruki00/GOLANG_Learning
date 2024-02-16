package main

import (
	"fmt"
)

func add(a int, b int) int {
	return int(a + b)
}

func procAdd(a int, b int, result *int) {
	*result = int(a + b)
}

var lambda = func(a int, b int) int { return int(a + b) }

func multiRet(a int, b int) (int, int) {
	k := a
	a = b
	b = k
	return a, b
}

func derfaultReturn(xp int, yp int) (x, y int) {
	x = xp + yp
	y = xp
	return
}

func multiArgs(items ...int) {
	for i, v := range items {
		fmt.Println("result: ", i, v)
	}
}
func main() {
	fmt.Println("a+b = ", add(1, 5))
	var result int
	procAdd(1, 5, &result)
	fmt.Println("a+b = ", result)
	fmt.Println("resul;t a+b = ", lambda(1, 5))
	a, b := multiRet(1, 5)
	fmt.Println("reverse 1, 5 = ", a, b)
	x, y := derfaultReturn(1, 5)
	fmt.Println("default return : ", x, y)
	multiArgs(1, 2, 3, 4, 5, 6)
}
