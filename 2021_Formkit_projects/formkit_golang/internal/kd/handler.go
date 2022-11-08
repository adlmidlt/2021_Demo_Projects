package kd

import (
	msg "cmd/pkg/logg/message"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"
)

const docVerNumber = 1 // Подставляется номер версии документа.

// renderFolderContent - Отобразить содержимое папки.
func renderFolderContent(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("../../internal/kd/web/html/index.html")
	check(msg.ErrParseFile, err)

	folders, docs, _ := readFolderContent()

	var dataFolders []Folder
	for _, folder := range folders {
		dataFolders = append(dataFolders, Folder{
			ObjectID:   folder.ObjectID,
			ObjectName: folder.ObjectName,
			Created:    folder.Created,
			Modified:   folder.Modified,
			AuthorID:   folder.AuthorID,
		})
	}
	var dataDocs []Doc
	for _, doc := range docs {
		dataDocs = append(dataDocs, Doc{
			ObjectID:      doc.ObjectID,
			ObjectName:    doc.ObjectName,
			Created:       doc.Created,
			Modified:      doc.Modified,
			AuthorID:      doc.AuthorID,
			Signed:        doc.Signed,
			SignatureType: doc.SignatureType,
		})
	}

	if err = tmpl.Execute(w, WrapFolderAndDoc{
		Folders: dataFolders,
		Docs:    dataDocs,
	}); check(msg.ErrParseTpl, err) {
		return
	}
}

// renderFolderContentByFolderId - Отобразить содержимое папки по идентификатору папки.
func renderFolderContentByFolderId(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../../internal/kd/web/html/index.html")
	check(msg.ErrParseFile, err)

	folderId, errConv := strconv.Atoi(r.URL.Query().Get("idFolder"))
	check(msg.ErrConv, errConv)

	folders, docs, _ := findFolderContentByFolderId(folderId)
	var dataFolders []Folder
	for _, folder := range folders {
		dataFolders = append(dataFolders, Folder{
			ObjectID:   folder.ObjectID,
			ObjectName: folder.ObjectName,
			Created:    folder.Created,
			Modified:   folder.Modified,
			AuthorID:   folder.AuthorID,
		})
	}
	var dataDocs []Doc
	for _, document := range docs {
		dataDocs = append(dataDocs, Doc{
			ObjectID:      document.ObjectID,
			ObjectName:    document.ObjectName,
			Created:       document.Created,
			Modified:      document.Modified,
			AuthorID:      document.AuthorID,
			Signed:        document.Signed,
			SignatureType: document.SignatureType,
		})
	}
	if err = tmpl.Execute(w, WrapFolderAndDoc{
		Folders: dataFolders,
		Docs:    dataDocs,
	}); check(msg.ErrParseTpl, err) {
		return
	}
}

// renderVersDocOrDocVer - Отобразить документ с версиями или версию документа.
func renderVersDocOrDocVer(w http.ResponseWriter, r *http.Request) {
	// Парсим шаблон html файла.
	tmpl, err := template.ParseFiles("../../internal/kd/web/html/doc.html")
	check(msg.ErrParseFile, err)

	// Возвращаем значение по ключу и преобразовываем в число.
	idDoc, errConvIdDoc := strconv.Atoi(r.URL.Query().Get("idDoc"))
	check(msg.ErrConv, errConvIdDoc)

	idVer, errConvIdVer := strconv.Atoi(r.URL.Query().Get("idVer"))
	if errConvIdVer != nil {
		idVer = 0
	}

	// Если версия документа не указана, то она равна 0.
	if idVer == 0 {
		docWithVers, errDocWithVer := findDocWithVer(idDoc, idVer)
		check(msg.ErrRowsEmpty, errDocWithVer)
		// Если у документа в списке одна версия ->
		if len(docWithVers) == 1 {
			// -> то открываем 1 версию.
			openVerDocInPDF(w, idDoc, docVerNumber)
		} else {
			// Иначе возвращаем шаблон документа со списком версий ->
			var dataDocWithVers []DocWithVer
			for _, docVers := range docWithVers {
				dataDocWithVers = append(dataDocWithVers, DocWithVer{
					IdDoc:            idDoc,
					FileName:         docVers.FileName,
					VersionNumber:    docVers.VersionNumber,
					Note:             docVers.Note,
					ModifyDate:       docVers.ModifyDate,
					Size:             docVers.Size,
					CRC:              docVers.CRC,
					AuthorDisplayFld: docVers.AuthorDisplayFld,
					EditorDisplayFld: docVers.EditorDisplayFld,
					TypeVersionData:  docVers.TypeVersionData,
				})
			}
			// Сортирует версии, в первом кортеже все актуальная.
			sort.Slice(dataDocWithVers, func(i, j int) bool {
				return dataDocWithVers[i].VersionNumber > dataDocWithVers[j].VersionNumber
			})
			// -> и отображаем.
			if err = tmpl.Execute(w, WrapDocWithVer{
				DocWithVers: dataDocWithVers,
			}); check(msg.ErrParseTpl, err) {
				return
			}
		}
	} else {
		// Иначе если версия указана, открываем PDF файл.
		openVerDocInPDF(w, idDoc, idVer)
	}
}

// openVerDocInPDF - Открыть версию документа в PDF формате.
func openVerDocInPDF(w http.ResponseWriter, docId int, verId int) {
	verDoc, errVerDoc := findVerDoc(docId, verId)
	check(msg.ErrRowsEmpty, errVerDoc)

	if verDoc.TypeVersionData == "PDF" {
		w.Header().Add("Content-Type", "application/pdf")
	}
	_, err := fmt.Fprintln(w, verDoc.VersionData)
	check(msg.ErrFormat, err)
}

// getJsonWithFolderContent - Получить JSON с содержимым папки.
func getJsonWithFolderContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	folders, docs, err := readFolderContent()
	check(msg.ErrReadFolderContent, err)

	if err = json.NewEncoder(w).Encode(WrapFolderAndDoc{
		Folders: folders,
		Docs:    docs,
	}); check(msg.ErrParseJson, err) {
		return
	}
}

// getJsonWithFolderContentByFolderId - Получить Json с содержимым папки по идентификатору папки.
func getJsonWithFolderContentByFolderId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")

	idFolder, errConv := strconv.Atoi(r.URL.Query().Get("idFolder"))
	check(msg.ErrConv, errConv)

	folders, docs, err := findFolderContentByFolderId(idFolder)
	check(msg.ErrNotFindFolderContentByIdFolder, err)

	if err = json.NewEncoder(w).Encode(WrapFolderAndDoc{
		Folders: folders,
		Docs:    docs,
	}); check(msg.ErrParseJson, err) {
		return
	}
}

// getJsonWithVersDocOrDocVer - Получить JSON с версиями документа или открыть версию документа.
func getJsonWithVersDocOrDocVer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idDoc, errConvIdDoc := strconv.Atoi(r.URL.Query().Get("idDoc"))
	check(msg.ErrConv, errConvIdDoc)

	idVer, errConvIdVer := strconv.Atoi(r.URL.Query().Get("idVer"))
	if errConvIdVer != nil {
		idVer = 0
	}

	if idVer == 0 {
		docWithVers, errDocWithVer := findDocWithVer(idDoc, idVer)
		check(msg.ErrRowsEmpty, errDocWithVer)
		// Если у документа в списке одна версия ->
		if len(docWithVers) == 1 {
			// -> то открываем 1 версию.
			openVerDocInPDF(w, idDoc, docVerNumber)
		} else {
			if err := json.NewEncoder(w).Encode(WrapDocWithVer{
				DocWithVers: docWithVers,
			}); check(msg.ErrParseTpl, err) {
				return
			}
		}
	} else {
		// Иначе если версия указана, открываем PDF файл.
		openVerDocInPDF(w, idDoc, idVer)
	}
}
