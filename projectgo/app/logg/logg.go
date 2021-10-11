package logg

import (
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const loggFile = "logg.log"

var (
	Info  *log.Logger
	Error *log.Logger
	Warn  *log.Logger
)

// InitLogg - Инициализация лога.
func InitLogg() {
	file, err := os.OpenFile(loggFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("\033[1;31m[E] Error while opening logger file: %s\033[0m", err.Error())
	}

	Info = log.New(file, " INFO ---", log.Lmsgprefix|log.LstdFlags)
	Error = log.New(file, "ERROR --- ", log.Lmsgprefix|log.LstdFlags)
	Warn = log.New(file, " WARN --- ", log.Lmsgprefix|log.LstdFlags)
}

// LogI - Лог информаций.
func LogI(message string) {
	log.Printf("\033[1;34m[I] %s\033[0m", message)
	Info.Printf(" %s - %s", fileWithFuncAndLineNum(), message)
}

// LogE - Лог ошибок.
func LogE(message error, err string) {
	log.Printf("%s \033[1;31m[E] %s: %s\033[0m", fileWithLineNum(), message, err)
	Error.Fatalf("%s - %s: %s", fileWithFuncAndLineNum(), message, err)
}

// LogW - Лог предупреждений.
func LogW(message string, err string) {
	log.Printf("%s \033[1;33m[W] %s: %s\033[0m", fileWithLineNum(), message, err)
	Warn.Printf("%s - %s: %s", fileWithFuncAndLineNum(), message, err)
}

/* Количество кадров стека, которые необходимо пропустить перед записью на ПК, где 0 идентифицирует
кадр для самих вызывающих абонентов, а 1 идентифицирует вызывающего абонента. Возвращает количество
записей, записанных на компьютер.*/
const skipNumOfStackFrame = 3

// fileWithLineNum Возвращает имя файла и номер строки текущего файла.
func fileWithLineNum() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skipNumOfStackFrame, pc)
	frame, _ := runtime.CallersFrames(pc[:n]).Next()
	idxFile := strings.LastIndexByte(frame.File, '/')

	return frame.File[idxFile+1:] + ":" + strconv.FormatInt(int64(frame.Line), 10)
}

// fileWithFuncAndLineNum Возвращает имя файла с функцией и номер строки текущего файла.
func fileWithFuncAndLineNum() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(skipNumOfStackFrame, pc)
	frame, _ := runtime.CallersFrames(pc[:n]).Next()
	idxFile := strings.LastIndexByte(strconv.Itoa(frame.Line), '/')

	return "[" + frame.Function + "] - " + frame.File[idxFile+1:] + ":" + strconv.FormatInt(int64(frame.Line), 10)
}
