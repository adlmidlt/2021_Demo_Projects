package kd

import (
	l "cmd/pkg/logg"
	msg "cmd/pkg/logg/message"
	"database/sql"
	_ "github.com/alexbrainman/odbc"
)

// Драйвер и настройки для подключения к БД.
const (
	driverDB        = "odbc"
	settingConnToDB = "Driver=FreeTDS; " +
		"Server=192.168.0.12,1433; " +
		"Database=DIRECTUM; " +
		"Uid=oper; " +
		"Pwd=oper; " +
		"ClientCharset=WINDOWS-1251"
)

var db *sql.DB

// ConnToDB - Подключиться к БД.
func connToDB() {
	var err error
	db, err = sql.Open(driverDB, settingConnToDB)
	if err != nil {
		l.LogE(msg.ErrConnToDB, err.Error())
	} else {
		l.LogI(msg.SuccessConnToDB)
	}
}
