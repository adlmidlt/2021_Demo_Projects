package logger

import (
	"log"
	"os"
)

const logFile = "./formkit.logger"

var (
	Info  *log.Logger
	Error *log.Logger
)

// InitLog - Инициализация логгера, для вывода сообщений о информации и ошибок.
func InitLog() {
	file, errFile := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if errFile != nil {
		log.Fatalf("\033[1;31m[E] Error while opening logger file: %s\033[0m", errFile.Error())
	}

	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.LstdFlags)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.LstdFlags)
}

// LogError - Логгирует сообщения об ошибках.
func LogError(message string) {
	log.Printf("\033[1;31m[E] %s\033[0m", message)
	Error.Printf("%s", message)
}

// LogInfo - Логгирует сообщение с информацией.
func LogInfo(message string) {
	log.Printf("\033[1;34m[I] %s\033[0m", message)
	Info.Printf("%s", message)
}
