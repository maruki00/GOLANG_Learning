package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  int
}

func JsonEncode(obj interface{}) string {
	jsonData, err := json.Marshal(obj)
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	return string(jsonData)
}

func JsonDecode(data []byte) []map[string]interface{} {
	d := make([]map[string]interface{}, 1, 1)
	tmp := make(map[string]interface{})
	err := json.Unmarshal(data, &tmp)

	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	d[0] = tmp
	return d
}

func main() {
	abdo := Person{Name: "abdullah", Age: 2345}
	jsonData := []byte(`{"Name":"karim", "Age":1344}`)
	fmt.Println(abdo, jsonData)
	fmt.Println("Result Decoding: ", JsonDecode(jsonData))
	fmt.Println("Result Encoding: ", JsonEncode(abdo))

}
