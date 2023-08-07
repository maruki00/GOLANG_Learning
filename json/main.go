package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Id   int `json:"id"`
	Name string
	Age  int
}

func main() {
	p := Person{
		Id:   1,
		Name: "Abdullah",
		Age:  27,
	}
	var p2 Person

	jsonData := []byte(`{"id":1, "name":"abdullah", "age":27}`)
	personArray, err1 := json.Marshal(p)
	err2 := json.Unmarshal(jsonData, &p2)

	if err1 != nil {
		log.Fatal(err1.Error())
	}
	fmt.Println(string(personArray))

	if err2 != nil {
		log.Fatal("Could not parse the JSON")
	}
	fmt.Printf("id: %d | name: %s | age: %d \n", p2.Id, p2.Name, p2.Age)

}
