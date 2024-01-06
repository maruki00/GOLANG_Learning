package model

type Person struct {
	Id   int    `json:"id"   gorm:"primary key"`
	Name string `json:"name" `
	Age  int    `json:"age"  `
}
