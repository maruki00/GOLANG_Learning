package sqlbuilder

import global "gorm2/Global"

func sqlbuilder() {
	global.Connect()
	db := global.GetCnx()
	db.Exec("insert into person values (1, 'abdullah', 12345)")
}
