package sh

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	l "shophookah/pkg/logg"
	msg "shophookah/pkg/logg/message"
)

// Данные для подключения к БД.
const (
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	password   = "1212"
	dbname     = "shop_hookah"
	schema     = "hookah"
	driverPSQL = "postgres"
)

// CreateConnToDB - Настройки для подключения к БД PostgreSQL.
func CreateConnToDB() *sql.DB {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		host, port, user, password, dbname, schema)
	db, err := sql.Open(driverPSQL, psqlConn)
	CheckError(msg.ErrConnToDB, err)
	l.LogI(msg.SuccessConnToDB)
	if err = db.Ping(); err != nil {
		panic(err)
	}
	return db
}

func CloseDB(db *sql.DB) {
	func(db *sql.DB) {
		if err := db.Close(); err != nil {
			l.LogE(msg.ErrCloseDB, err.Error())
		} else {
			l.LogI(msg.SuccessCloseConnToDB)
		}
	}(db)
}

// CloseRows - Закрыть строки, предотвращающие дальнейшие перечисления.
func CloseRows(rows *sql.Rows, err error) {
	func(rows *sql.Rows) {
		if err = rows.Close(); err != nil {
			l.LogE(msg.ErrCloseRows, err.Error())
		}
	}(rows)
}

// ErrRows - Ошибка строк при явном или неявном закрытии.
func ErrRows(rows *sql.Rows, err error) {
	func(rows *sql.Rows) {
		if err = rows.Err(); err != nil {
			l.LogE(msg.ErrExplicitOrImplicitClose, err.Error())
		}
	}(rows)
}
