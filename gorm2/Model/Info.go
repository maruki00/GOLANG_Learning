package model

import "gorm.io/gorm"

type Info struct {
	gorm.Model
	Id      int    `json:"id"`
	Address string `json:"address"`
	Street  string `json:"street"`
}
