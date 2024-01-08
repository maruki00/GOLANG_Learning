package models

import (
	"belongto/global"
	"fmt"

	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	// Id   int    `json:"id" gorm:"primary key"`
	Name  string `json:"name"`
	Users []User
}

func init() {
	fmt.Println("init Company")
	db := global.GetCnx()
	db.AutoMigrate(Company{})
}
