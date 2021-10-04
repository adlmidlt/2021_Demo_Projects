package main

import (
	"hookahshop/hs/db"
	"hookahshop/hs/handler"
	"hookahshop/hs/logg"
)

func main() {
	logg.InitLogg()
	db.ConnToDB()
	handler.Handler()
}
