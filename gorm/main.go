package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/go-gorm/gorm/dbtools"
	"github.com/go-gorm/gorm/model"
)

type Config struct {
	DataSourceName string `json:"dataSourceName"`
}

func main() {

	file, err := os.Open("configuration/config.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer file.Close()

	conf := new(Config)

	json.NewDecoder(file).Decode(conf)

	dbtools.DBInitializer(conf.DataSourceName)

	dbtools.CreateTable(&model.Student{})

}
