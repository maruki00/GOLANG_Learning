package main

import (
	"belongto/global"
	"belongto/models"
	"encoding/json"
	"fmt"
)

func main() {
	db := global.GetCnx()
	var company = models.Company{Name: "AZ Company"}
	db.Create(&company)
	// next ceate a user of the company
	var user = models.User{Name: "john", CompanyID: int(company.ID)}
	db.Create(&user)

	//////////// 2- find a user and load his company (eager loading) ////////////
	var dbUser models.User
	db.Model(models.User{}).Preload("Company").Where("name = ?", "john").First(&dbUser)
	// fmt.Println("User ID: ", dbUser.ID)
	// fmt.Println("User Name: ", dbUser.Name)
	// fmt.Println("User CompanyID: ", dbUser.CompanyID)
	// fmt.Println("User Company ID: ", dbUser.Company.ID)
	// fmt.Println("User Company Name: ", dbUser.Company.Name)

	var company1 models.Company
	db.Model(models.Company{}).Preload("Users").Where("ID=?", 2).Find(&company1)
	jsonData, _ := json.Marshal(company1)
	fmt.Println("Json Data: ", string(jsonData))
	// fmt.Println("Fetch Company Users: ", company1.Name)
	// for i, user := range company1.Users {
	// 	fmt.Println("User Name ", i, " : ", user.Name)
	// }
	//////////// 3- update user's Company ////////////
	dbUser.Company.Name = "A TO Z Company" // update the name of the company
	// when updating a relationship you must use `db.Session`
	// db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&dbUser)
}
