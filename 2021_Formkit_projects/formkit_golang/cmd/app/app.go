package main

import (
	"cmd/internal/kd"
	l "cmd/pkg/logg"
)

// Main - запуск приложения.
func main() {
	l.InitLogg()
	kd.HttpHandler()
	kd.Server()
}
