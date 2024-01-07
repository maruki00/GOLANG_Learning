package model

type Person struct {
	// gorm.Model
	Id   int    `json:"id"   gorm:"primary key"`
	Name string `json:"name" `
	Age  int    `json:"age"  `
	Car  Car    `json:"car"`
}
