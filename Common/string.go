package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Starting Coding ...")

	const C_NAME string = "Hello : "
	var name string = "abdullah"
	fmt.Print(C_NAME)
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c", name[i])
	}
	fmt.Println("")

	var strHex string = "\x41\x41\x41\x41\x41\x41\x41\x41\x41"
	fmt.Print(strHex)

	var str1 string = "str1 "
	var str2 string = "str2 "
	fmt.Println("\n", strings.Join([]string{str1, str2}, ""))

	name1 := `abdullah oulahyane \n`
	fmt.Print(name1)

}
