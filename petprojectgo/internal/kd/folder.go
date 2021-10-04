package kd

import (
	msg "petprojectgo/pkg/logg/message"
	"petprojectgo/pkg/util"
)

// WrapFolderAndDoc - Обёртка для папок и документов.
type WrapFolderAndDoc struct {
	Folders []Folder `json:"folders"`
	Docs    []Doc    `json:"docs"`
}

// Folder - Папка
type Folder struct {
	ObjectID   int    `json:"objectId"`
	ObjectName string `json:"objectName"`
	Created    string `json:"created"`
	Modified   string `json:"modified"`
	AuthorID   int    `json:"authorId"`
}

// readFolderContent - Прочитать содержимое папки.
func readFolderContent() ([]Folder, []Doc, error) {
	connToDB()
	rows, err := db.Query("exec web_getFolderContents 119551")
	check(msg.ErrQueryToDB, err)
	defer closeRows(rows)

	var folders []Folder
	for rows.Next() {
		var folder Folder
		if err = rows.Scan(
			&folder.ObjectID,
			&folder.ObjectName,
			&folder.Created,
			&folder.Modified,
			&folder.AuthorID,
		); check(msg.ErrRowsEmpty, err) {
			return nil, nil, err
		}
		folder.ObjectName, _ = util.Win1251ToUTF8(folder.ObjectName)

		folders = append(folders, folder)
	}
	rows.NextResultSet()

	var docs []Doc
	for rows.Next() {
		var doc Doc
		if err = rows.Scan(
			&doc.ObjectID,
			&doc.ObjectName,
			&doc.Created,
			&doc.Modified,
			&doc.AuthorID,
			&doc.Signed,
			&doc.SignatureType,
		); check(msg.ErrRowsEmpty, err) {
			return nil, nil, err
		}
		doc.ObjectName, _ = util.Win1251ToUTF8(doc.ObjectName)
		doc.Signed, _ = util.Win1251ToUTF8(doc.Signed)
		doc.SignatureType, _ = util.Win1251ToUTF8(doc.SignatureType)

		docs = append(docs, doc)
	}
	if err = rows.Err(); check(msg.ErrIterRows, err) {
		return nil, nil, err
	}
	return folders, docs, err
}

// findFolderContentByFolderId - Найти содержимое папки по ИД папки.
func findFolderContentByFolderId(idFolder int) ([]Folder, []Doc, error) {
	connToDB()
	rows, err := db.Query("exec web_getFolderContents ?", idFolder)
	check(msg.ErrQueryToDB, err)
	defer closeRows(rows)

	var folders []Folder
	for rows.Next() {
		var folder Folder
		if err = rows.Scan(
			&folder.ObjectID,
			&folder.ObjectName,
			&folder.Created,
			&folder.Modified,
			&folder.AuthorID,
		); check(msg.ErrRowsEmpty, err) {
			return nil, nil, err
		}
		folder.ObjectName, _ = util.Win1251ToUTF8(folder.ObjectName)

		folders = append(folders, folder)
	}
	rows.NextResultSet()

	var docs []Doc
	for rows.Next() {
		var doc Doc
		if err = rows.Scan(
			&doc.ObjectID,
			&doc.ObjectName,
			&doc.Created,
			&doc.Modified,
			&doc.AuthorID,
			&doc.Signed,
			&doc.SignatureType,
		); check(msg.ErrRowsEmpty, err) {
			return nil, nil, err
		}
		doc.ObjectName, _ = util.Win1251ToUTF8(doc.ObjectName)
		doc.Signed, _ = util.Win1251ToUTF8(doc.Signed)
		doc.SignatureType, _ = util.Win1251ToUTF8(doc.SignatureType)

		docs = append(docs, doc)
	}
	if err = rows.Err(); check(msg.ErrIterRows, err) {
		return nil, nil, err
	}
	return folders, docs, err
}
