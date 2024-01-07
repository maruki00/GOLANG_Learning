package hasone

import (
	"fmt"
	global "gorm2/Global"
	model "gorm2/Model"
)

func HasOne() {
	var person model.Person
	db := global.GetCnx()
	db.Debug().Model(model.Person{}).Preload("Car").Scan(&person)
	fmt.Println("Person: ", person)
}
