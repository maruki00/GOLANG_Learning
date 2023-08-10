package model

type Student struct {
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"nvarchar(50);not null"`
	Age  int    `gorm:"not null"`
}
