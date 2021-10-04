package db

import (
	"database/sql"
	"fmt"
	"shophookah/internal"
	l "shophookah/pkg/logg"
	msg "shophookah/pkg/logg/message"
)

// Данные для подключения к БД.
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1212"
	dbname   = "shop_hookah"
	schema   = "hookah"
)

var DB *sql.DB

// PSQL - Настройки для подключения к БД PostgreSQL.
func PSQL() {
	var err error
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		host, port, user, password, dbname, schema)
	DB, err = sql.Open("postgres", psqlConn)
	internal.CheckError(msg.ErrConnToDB, err)
	l.LogI(msg.SuccessConnToDB)
}

// CloseDB - Закрыть соединение с БД.
func CloseDB(err error) {
	func() {
		if err = DB.Close(); err != nil {
			l.LogE(msg.ErrCloseDB, err.Error())
		} else {
			l.LogI(msg.SuccessCloseConnToDB)
		}
	}()
}

// CloseRows - Закрыть строки, предотвращающие дальнейшие перечисления.
func CloseRows(rows *sql.Rows, err error) {
	func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			l.LogE(msg.ErrCloseRows, err.Error())
		}
	}(rows)
}

// ErrRows - Ошибка при явном или неявном закрытии строк.
func ErrRows(rows *sql.Rows, err error) {
	func(rows *sql.Rows) {
		if err = rows.Err(); err != nil {
			l.LogE(msg.ErrExplicitOrImplicitClose, err.Error())
		}
	}(rows)
}
