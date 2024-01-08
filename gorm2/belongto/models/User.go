package models

import (
	"belongto/global"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// Id        int    `json:"id" gorm:"primaty key"`
	Name      string `json:"name"`
	CompanyID int    `json:"company_id"`
	Company   Company
}

func init() {
	fmt.Println("init User")
	db := global.GetCnx()
	db.AutoMigrate(User{})
}
