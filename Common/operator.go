package main

import "fmt"

func main() {
	var seconds int = 10000
	var mod float32 = 1233 % 12
	var lshift int = 1 << 30
	var rshift int = 1 >> 13
	var plus int = 100 + 12
	var minus int = 120 - 30
	var multi = 12 * 32
	var div = 12 / 1

	fmt.Println(seconds, mod, lshift, rshift, plus, minus, multi, div)

}
