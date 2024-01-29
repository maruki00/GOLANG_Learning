package main

import (
	"fmt"
)

func main() {
	var dict = make(map[int]string)
	var dict2 = map[int]string{}
	var dict3 = map[int]string{
		1: "line 1",
		2: "line 2",
		3: "line 3",
	}
	dict2[0] = "line 0"
	dict2[1] = "line 1"

	fmt.Println(dict)
	fmt.Println(dict2)
	fmt.Println(dict3)

	for key, value := range dict3 {
		fmt.Printf("%d <-> %s\n", key, value)
	}

	fmt.Print(len(dict3))
}
