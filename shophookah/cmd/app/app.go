package main

import (
	"fmt"
	migr "shophookah/cmd/migration_cli"
	"shophookah/internal/sh"
	l "shophookah/pkg/logg"
)

func main() {
	l.InitLogg()
	sh.HttpHandler()
	migr.Execute()
	sh.Server()

	fmt.Println("Hookah Shop")
}
