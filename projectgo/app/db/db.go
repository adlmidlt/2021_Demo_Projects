package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"projectgo/app/logg"
)

// Данные для подключения к БД.
const (
	host       = "localhost"
	port       = 5432
	user       = "postgres"
	password   = "1212"
	dbName     = "shop_hookah"
	schema     = "hookah"
	driverPSQL = "postgres"
)

// ConnToDB - Подключение к БД.
func ConnToDB() *sql.DB {
	connPSQL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		host, port, user, password, dbName, schema)
	db, err := sql.Open(driverPSQL, connPSQL)
	if err != nil {
		logg.LogE(errors.New("ошибка подключения к БД"), err.Error())
	}
	return db
}

// CloseDB - Закрыть соединение к БД, для освобождения ресурсов.
func CloseDB(db *sql.DB) {
	func(db *sql.DB) {
		if err := db.Close(); err != nil {
			logg.LogE(errors.New("ошибка, не удалось закрыть соединение с БД"), err.Error())
		}
	}(db)
}
