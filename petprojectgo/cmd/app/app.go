package main

import (
	"petprojectgo/internal/kd"
	l "petprojectgo/pkg/logg"
)

// Main - запуск приложения.
func main() {
	l.InitLogg()
	kd.HttpHandler()
	kd.Server()
}
