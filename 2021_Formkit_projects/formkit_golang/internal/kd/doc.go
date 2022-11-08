package kd

import (
	msg "cmd/pkg/logg/message"
	"cmd/pkg/util"
)

// WrapDocWithVer - Обертка для документа с версиями.
type WrapDocWithVer struct {
	DocWithVers []DocWithVer `json:"docWithVers"`
}

// Doc - Документ.
type Doc struct {
	ObjectID      int    `json:"objectID"`
	ObjectName    string `json:"objectName"`
	Created       string `json:"created"`
	Modified      string `json:"modified"`
	AuthorID      int    `json:"authorID"`
	Signed        string `json:"signed"`
	SignatureType string `json:"signatureType"`
}

// DocWithVer - Документ с версией.
type DocWithVer struct {
	IdDoc            int    `json:"idDoc"`
	FileName         string `json:"fileName"`
	VersionNumber    int    `json:"versionNumber"`
	Note             string `json:"note"`
	ModifyDate       string `json:"modifyDate"`
	Size             int    `json:"size"`
	CRC              string `json:"CRC"`
	AuthorDisplayFld string `json:"authorDisplayFld"`
	EditorDisplayFld string `json:"editorDisplayFld"`
	TypeVersionData  string `json:"typeVersionData"`
}

// VerDoc - Версия документа.
type VerDoc struct {
	VersionData     string `json:"versionData"`
	Size            int    `json:"size"`
	TypeVersionData string `json:"typeVersionData"`
	Number          int    `json:"number"`
}

// findDocWithVer - Найти документ с версией.
func findDocWithVer(idDoc, idVer int) ([]DocWithVer, error) {
	connToDB()
	rows, err := db.Query("exec web_getFile ?, ?", idDoc, idVer)
	check(msg.ErrQueryToDB, err)
	defer closeRows(rows)

	var docWithVers []DocWithVer
	for rows.Next() {
		var docWithVer DocWithVer
		if err = rows.Scan(
			&docWithVer.FileName,
			&docWithVer.VersionNumber,
			&docWithVer.Note,
			&docWithVer.ModifyDate,
			&docWithVer.Size,
			&docWithVer.CRC,
			&docWithVer.AuthorDisplayFld,
			&docWithVer.EditorDisplayFld,
			&docWithVer.TypeVersionData,
		); check(msg.ErrRowsEmpty, err) {
			return nil, err
		}
		docWithVer.FileName, _ = util.Win1251ToUTF8(docWithVer.FileName)
		docWithVer.Note, _ = util.Win1251ToUTF8(docWithVer.Note)
		docWithVer.AuthorDisplayFld, _ = util.Win1251ToUTF8(docWithVer.AuthorDisplayFld)
		docWithVer.EditorDisplayFld, _ = util.Win1251ToUTF8(docWithVer.EditorDisplayFld)
		docWithVer.TypeVersionData, _ = util.Win1251ToUTF8(docWithVer.TypeVersionData)

		docWithVers = append(docWithVers, docWithVer)
	}
	if err = rows.Err(); check(msg.ErrIterRows, err) {
		return nil, err
	}
	return docWithVers, nil
}

// findVerDoc - Найти версию документа.
func findVerDoc(idDoc, idVer int) (VerDoc, error) {
	connToDB()
	rows, err := db.Query("exec web_getFile ?, ?", idDoc, idVer)
	check(msg.ErrQueryToDB, err)
	defer closeRows(rows)

	var verDoc VerDoc
	rows.Next()
	if err = rows.Scan(
		&verDoc.VersionData,
		&verDoc.Size,
		&verDoc.TypeVersionData,
		&verDoc.Number,
	); check(msg.ErrRowsEmpty, err) {
		return VerDoc{}, err
	}
	if err = rows.Err(); check(msg.ErrIterRows, err) {
		return VerDoc{}, err
	}
	return verDoc, nil
}
