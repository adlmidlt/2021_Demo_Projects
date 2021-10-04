package sh

import "net/http"

// HttpHandler - Обработчик http запросов.
func HttpHandler() {
	// TEMPLATE
	http.HandleFunc("/create", UserSingUp)
	/*	http.HandleFunc("/auth", UserSingIn)
		http.HandleFunc("/", HomePage)*/

	// JSON
	// http.HandleFunc("/api/all", getJsonWithFolderContent)

	// Подключение static (*.html, *.png/jpg *.css файлов, *.js)
	http.Handle("/web/", http.StripPrefix("/web/",
		http.FileServer(http.Dir("../../internal/sh/web/"))))
}
