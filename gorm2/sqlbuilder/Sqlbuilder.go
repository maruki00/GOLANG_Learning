package sqlbuilder

import (
	"fmt"
	global "gorm2/Global"
	model "gorm2/Model"

	"gorm.io/gorm"
)

func Initial() *gorm.DB {
	global.Connect()
	db := global.GetCnx()
	return db
}

func RunSqlbuilder() {
	db := Initial()
	db.Exec("insert into person values (NULL, 'abdullah', 12345)")
}

func FetchAll() {
	var person []model.Person
	db := Initial()
	db.Raw("select * from person").Scan(&person)
	for key, value := range person {
		fmt.Println("key: ", key, "  value: ", value)
	}
	// fmt.Println("Fetched Rows: ", person)
}

func Innerjoin() {
	var person []model.Person
	db := Initial()
	db.Raw("select * from person inner join info on person.id = info.person_id").Scan(&person)
	fmt.Println("Person: ", person)
}
