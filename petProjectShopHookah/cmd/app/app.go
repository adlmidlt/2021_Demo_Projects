package main

import (
	"fmt"
	"net/http"
	"shophookah/cmd/migration_cli"
	"shophookah/internal/db"
	l "shophookah/pkg/logg"
)

func main() {
	l.InitLogg()
	db.PSQL()
	migration_cli.Execute()
	fmt.Println("Hookah Shop")

	http.HandleFunc("/", homePage)

	http.ListenAndServe(":8080", nil)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "dasdasdasdasdasdas")
}
