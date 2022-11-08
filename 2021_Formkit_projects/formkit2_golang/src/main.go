package main

import (
	"formkit/src/src/config"
	"formkit/src/src/handler"
	log "formkit/src/src/logger"
	"net/http"
)

func main() {
	log.InitLog()

	http.HandleFunc("/", handler.HandleFoldersAndDocs)
	http.HandleFunc("/api/get", handler.XHRJsonFoldersAndDocs)

	// http://192.168.2.221:8080/api/getId?folderId=121138
	http.HandleFunc("/folder", handler.HandleFolderWithDocsOnFolderId)
	http.HandleFunc("/api/getId", handler.XHRJsonFolderWithDocsOnFolderId)

	// http://192.168.2.221:8080/all/doc?docId=430322&verId=1
	http.HandleFunc("/doc", handler.HandleDocWithVer)
	http.HandleFunc("/api/getOnVer", handler.XHRJsonHandleDocWithVer)

	// Подключение static (*.html, *.png/jpg *.css файлов, *.js)
	http.Handle("/web/",
		http.StripPrefix("/web/", http.FileServer(http.Dir("web/"))))
	config.ServerStart()
}
