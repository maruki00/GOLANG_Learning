package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func main() {
	jsonData := `{"id":1,"name":"abdullah","age":10}{"id":2,"name":"abdullah","age":10}{"id":3,"name":"abdullah","age":10}`

	reader := strings.NewReader(jsonData)

	writer := os.Stdout

	decoder := json.NewDecoder(reader)

	encoder := json.NewEncoder(writer)

	for {
		var v map[string]interface{}
		if err := decoder.Decode(&v); err != nil {
			log.Fatal(err.Error())
			return
		}
		for k := range v {
			if k == "age" {
				// delete(v, k)
			}
		}

		if err := encoder.Encode(&v); err != nil {
			log.Println(err)
		}
	}

}
