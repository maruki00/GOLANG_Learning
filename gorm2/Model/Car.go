package model

type Car struct {
	// gorm.Model
	Id        int    `gorm:"primary key" json:"id"`
	Name      string `json:"name"`
	Person_id int    `json:"person_id"`
}
