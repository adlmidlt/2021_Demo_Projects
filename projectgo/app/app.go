package main

import (
	"projectgo/app/handler"
	"projectgo/app/logg"
)

func main() {
	handler.HttpHandler()
	logg.InitLogg()
	logg.LogI("Hello World!!!")
}
