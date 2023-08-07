package main

import (
	"log"
	"os"
)
type struct config{
	DriverName string `json:"driverName"`,
	DataSource string `json:"dataSourceName"`
}
func main() {
	file, err := os.Open("configs/config.json")
	if err!=nil{
		log.Fatal(err.Error())
	}
	defer file.Close()

}
