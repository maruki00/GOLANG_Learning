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

	// dbtools.CreateTable(&model.Student{})

	// st := model.Student{
	// 	Id:   2,
	// 	Name: "Abdullah",
	// 	Age:  27,
	// }
	// dbtools.Save(&st)

	// var students []model.Student

	// dbtools.SeletcAll(students)

	st3 := model.Student{
		Id: 2,
	}

	data := map[string]interface{}{
		"Id":   1,
		"Name": "Abdullah oulahyane",
		"Age":  234145,
	}

	dbtools.UpdateOne(st3, data)
	// fmt.Println("rows infected ", rows)

}
