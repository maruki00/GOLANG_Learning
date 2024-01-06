package global

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetCnx() *gorm.DB { return db }
func Connect() {
	d, err := gorm.Open(mysql.Open("gorm_test:user@tcp(127.0.0.1:3306)/gorm_test"))
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	db = d
}
