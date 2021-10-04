package kd

import (
	"database/sql"
	_ "github.com/alexbrainman/odbc"
	l "petprojectgo/pkg/logg"
	msg "petprojectgo/pkg/logg/message"
)

// Драйвер и настройки для подключения к БД.
const (
	driverDB        = "************"
	settingConnToDB = "Driver=***********; " +
		"Server=*************; " +
		"Database=**********; " +
		"Uid=*********; " +
		"Pwd=*********; " +
		"ClientCharset=**********"
)

var db *sql.DB

// ConnToDB - Подключиться к БД.
func connToDB() {
	var err error
	db, err = sql.Open(driverDB, settingConnToDB)
	check(msg.ErrConnToDB, err)

	l.LogI(msg.SuccessConnToDB)
}
