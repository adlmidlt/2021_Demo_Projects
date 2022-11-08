package handler

import (
	"encoding/json"
	"fmt"
	log "formkit/src/src/logger"
	msg "formkit/src/src/logger/message"
	"formkit/src/src/model"
	"formkit/src/src/repository"
	"io/ioutil"
	"net/http"
	"strconv"
)

// HandleFoldersAndDocs - Обработать папки и документы.
func HandleFoldersAndDocs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := ioutil.ReadFile("web/html/index.html")
	if err != nil {
		log.LogError("e: " + err.Error())
	}
	w.Write(tmpl)
}

// HandleFolderWithDocsOnFolderId - Обработать папку с документами по ИД папки.
func HandleFolderWithDocsOnFolderId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := ioutil.ReadFile("web/html/index.html")
	if err != nil {
		log.LogError("e: " + err.Error())
	}
	w.Write(tmpl)
}

// HandleDocWithVer - Обработать документы с версиями.
func HandleDocWithVer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := ioutil.ReadFile("web/html/index.html")
	if err != nil {
		log.LogError("e: " + err.Error())
	}
	w.Write(tmpl)
}

// openVerListOrVerPdfDoc - Открыть список версий или документ в формате PDF.
func openVerListOrVerPdfDoc(w http.ResponseWriter, docId int, verId int) {
	versDoc, err := repository.OpenDocWithVers(docId, verId)
	if err != nil {
		log.LogError(msg.ErrOpenDocWithVers + " versDoc: " + err.Error())
		return
	}

	if len(versDoc) == 1 {
		openVerDocOnTypeFormat(w, docId, 1)
	} else {
		err = json.NewEncoder(w).Encode(model.WrapTDocWithVers{
			DocumentWithVersions: versDoc,
		})
		if err != nil {
			log.LogError(msg.ErrParseData + err.Error())
			return
		}
	}
}

// openVerDocOnTypeFormat - Открыть версию PDF документа.
func openVerDocOnTypeFormat(w http.ResponseWriter, docId int, verId int) {
	docWithVers, errDocWithVers := repository.OpenVerDoc(docId, 1)
	if errDocWithVers != nil {
		log.LogError(msg.ErrOpenDocWithVers + " OpenVerDoc(): " + errDocWithVers.Error())
		return
	}

	switch docWithVers.TypeVersionData {
	case "PDF":
		w.Header().Add("Content-Type", "application/pdf")
		break
	default:
		log.LogInfo(msg.NotExistFormatDoc)
	}

	_, err := fmt.Fprintln(w, docWithVers.VersionData)
	if err != nil {
		log.LogError(msg.ErrDocFormat + err.Error())
	}
}

// XHRJsonFoldersAndDocs - Вывести в Json формате все папки и документы.
func XHRJsonFoldersAndDocs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	folders, docs, _ := repository.FindAllFoldersAndDocs()
	err := json.NewEncoder(w).Encode(model.WrapTFolderAndTDoc{
		Folders:   folders,
		Documents: docs,
	})
	if err != nil {
		log.LogError(msg.ErrParseData + err.Error())
		panic(err)
	}
}

// XHRJsonFolderWithDocsOnFolderId - Вывести в json формате папку с документами по ИД папки.
func XHRJsonFolderWithDocsOnFolderId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	folderId, errConv := strconv.Atoi(r.URL.Query().Get("folderId"))
	if errConv != nil {
		log.LogError(msg.ErrConvertToInt + " folderId: " + errConv.Error())
	}

	folders, docs, _ := repository.GetFolderWithDocsOnFolderId(folderId)
	err := json.NewEncoder(w).Encode(model.WrapTFolderAndTDoc{
		Folders:   folders,
		Documents: docs,
	})
	if err != nil {
		log.LogError(msg.ErrParseData + err.Error())
	}
}

// XHRJsonHandleDocWithVer - Вывести в json формате документы с версиями.
func XHRJsonHandleDocWithVer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	docId, errConvDocId := strconv.Atoi(r.URL.Query().Get("docId"))
	if errConvDocId != nil {
		log.LogError(msg.ErrConvertToInt + " docId: " + errConvDocId.Error())
	}
	verId, errConvVerId := strconv.Atoi(r.URL.Query().Get("verId"))
	if errConvVerId != nil {
		verId = 0
	}

	if verId == 0 {
		openVerListOrVerPdfDoc(w, docId, verId)
	} else {
		openVerDocOnTypeFormat(w, docId, verId)
	}
}
