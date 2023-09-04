package dbtools

import (
	"fmt"
	"log"

	"github.com/go-gorm/gorm/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dataSourceName string

func DBInitializer(dsn string) {
	dataSourceName = dsn
}

func connect() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func CreateTable(object ...interface{}) {
	db := connect()

	db.AutoMigrate(object...)

}

func Save(object interface{}) {
	db := connect()
	db.Create(object)
}

func UpdateOne(object interface{}, data map[string]interface{}) interface{} {
	db := connect()
	db.Find(object)
	// result := db.Model(object).Updates(data)
	// return result.RowsAffected
	return db
}

func SeletcAll(students []model.Student) {
	db := connect()
	db.Find(&students)
	for _, student := range students {
		fmt.Printf("id: %d, name: %s, age: %d\n", student.Id, student.Name, student.Age)
	}
}
