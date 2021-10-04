package kd

import "net/http"

// HttpHandler - Обработчик http запросов.
func HttpHandler() {
	// TEMPLATE
	http.HandleFunc("/", renderFolderContent)
	http.HandleFunc("/folder", renderFolderContentByFolderId)
	http.HandleFunc("/document", renderVersDocOrDocVer)

	// JSON
	http.HandleFunc("/api/all", getJsonWithFolderContent)
	http.HandleFunc("/api/folder", getJsonWithFolderContentByFolderId)
	http.HandleFunc("/api/document", getJsonWithVersDocOrDocVer)

	// Подключение static (*.html, *.png/jpg *.css файлов, *.js)
	http.Handle("/web/", http.StripPrefix("/web/",
		http.FileServer(http.Dir("../../internal/kd/web/"))))
}
