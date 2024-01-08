package global

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func connect() {
	d, err := gorm.Open(mysql.Open("gorm_test:user@tcp(127.0.0.1:3306)/gorm_test?parseTime=true"), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Error")
	}
	db = d
}

func GetCnx() *gorm.DB {
	if db == nil {
		connect()
	}
	return db
}
