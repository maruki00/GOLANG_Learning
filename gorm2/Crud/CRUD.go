package crud

import (
	"fmt"
	global "gorm2/Global"
	model "gorm2/Model"
)

func Create() {
	db := global.GetCnx()
	var person = model.Person{
		Name: "karim",
		Age:  1234,
	}
	db.Create(&person)
}

func SelectOne() {
	db := global.GetCnx()
	var person model.Person
	db.Where("id=?", 1).Find(&person)
	fmt.Println("Result: ", person)
}

func SelectMany() {
	db := global.GetCnx()
	var people []model.Person
	db.Where("id>0").Find(&people)
	fmt.Println("Result: ", people)
}

func SaveOrUpdatePerson() {
	db := global.GetCnx()
	var person model.Person = model.Person{
		Id:   1,
		Name: "Ahmed",
		Age:  123345234,
	}
	db.Save(person)

}

func UpdatePerson() {
	db := global.GetCnx()

	db.Model(model.Person{}).Where("id>?", 0).Update("name", "abdullah")
}

func Delete() {
	db := global.GetCnx()
	var person model.Person
	db.Model(model.Person{}).Where("id=?", 3).Delete(&person)
	fmt.Println("Result: ", person)
}

func Search(name string, age int) {
	db := global.GetCnx()
	var people []model.Person
	db.Model(model.Person{}).Where("name=?", name).Or("age=?", age).Find(&people)
	fmt.Println("Search Result: ", people)
}
