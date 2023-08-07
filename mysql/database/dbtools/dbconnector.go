package dbtools

import (
	"database/sql"
	"log"
)

var driverName string
var dataSourceName string

func DBInitializer(dn, dsn string) {

	driverName = dn
	dataSourceName = dsn
}

func connect() (db *sql.DB) {

	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func fetchItems() []module.Employee {

}
