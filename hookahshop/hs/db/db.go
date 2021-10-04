package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"hookahshop/hs/logg"
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

	if err = db.Ping(); err != nil {
		logg.LogE(errors.New("не возможно установить соединение с БД"), err.Error())
	}

	logg.LogI("Соединение с БД успешно установлено.")
	return db
}

// CloseDB - Закрыть БД.
func CloseDB() (*sql.DB, bool) {
	if err := ConnToDB().Close(); err != nil {
		logg.LogE(errors.New("не удалось закрыть соединение с БД"), err.Error())
	}
	logg.LogI("Соединение с БД успешно закрыто.")
	return nil, false
}
