package config

import (
	"database/sql"
	log "formkit/src/src/logger"
	msg "formkit/src/src/logger/message"
	_ "github.com/alexbrainman/odbc"
	"net/http"
	"os"
)

const DriverDataBase = "odbc"
const SettingConnectToDB = "Driver=FreeTDS; Server=192.168.0.12,1433; Database=DIRECTUM; Uid=oper; Pwd=oper; ClientCharset=WINDOWS-1251"

// ServerStart - Запуск веб сервера.
func ServerStart() {
	log.LogInfo(msg.SuccessStartServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.LogError(msg.ErrWebServerStart + err.Error())
		os.Exit(1)
	}
}

// ConnectDB - Возвращает настроенное подключение к БД.
func ConnectDB() *sql.DB {
	db, errDB := sql.Open(DriverDataBase, SettingConnectToDB)
	if errDB != nil {
		log.LogError(msg.ErrConnectToDB + errDB.Error())
	} else {
		log.LogInfo(msg.SuccessConnectToDB)
	}
	return db
}
