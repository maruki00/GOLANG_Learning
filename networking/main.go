package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type users struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	var dsn = "root:user@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&users{})
	user := &users{
		Id:   110,
		Name: "abdullah",
	}

	fmt.Println(user, db)

	// user2 := User{
	// 	id:   2,
	// 	name: "abdullah",
	// }

	db.Create(user)

	// err2 := db.Create(&user2)
	// if err2 != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(db.Find(&users{}).Where("id=?", 10))

}
