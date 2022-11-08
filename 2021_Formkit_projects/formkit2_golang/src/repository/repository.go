package repository

import (
	"database/sql"
	"formkit/src/src/config"
	log "formkit/src/src/logger"
	msg "formkit/src/src/logger/message"
	"formkit/src/src/model"
	utils "formkit/src/src/utility"
)

// FindAllFoldersAndDocs - Найти все папки и документы.
func FindAllFoldersAndDocs() ([]model.TFolder, []model.TDoc, error) {
	rows, err := config.ConnectDB().Query("exec web_getFolderContents 119551")
	if err != nil {
		log.LogError(msg.ErrQueryToDB + " FindAllFoldersAndDocs(): " + err.Error())
		return nil, nil, err
	}

	folders, docs, errFill := fillFolderAndDocsList(rows, err)
	if errFill != nil {
		log.LogError(msg.ErrFillList + "fillFolderAndDocsList(): " + errFill.Error())
	}

	return folders, docs, nil
}

// GetFolderWithDocsOnFolderId - Вернуть папку с документами по ИД папки.
func GetFolderWithDocsOnFolderId(folderId int) ([]model.TFolder, []model.TDoc, error) {
	rows, err := config.ConnectDB().Query("exec web_getFolderContents ?", folderId)
	if err != nil {
		log.LogError(msg.ErrQueryToDB + " GetFolderWithDocsOnFolderId(): " + err.Error())
		return nil, nil, err
	}

	folders, docs, errFill := fillFolderAndDocsList(rows, err)
	if errFill != nil {
		log.LogError(msg.ErrFillList + " fillFolderAndDocsList(): " + errFill.Error())
		return nil, nil, errFill
	}

	return folders, docs, nil
}

// OpenDocWithVers - Открыть документ с версиями.
func OpenDocWithVers(docId int, verId int) ([]model.TDocWithVers, error) {
	rows, err := config.ConnectDB().Query("exec web_getFile ?, ?", docId, verId)
	if err != nil {
		log.LogError(msg.ErrQueryToDB + " OpenDocWithVers(): " + err.Error())
		return nil, err
	}

	var versDoc []model.TDocWithVers
	for rows.Next() {
		var verDoc model.TDocWithVers
		if err = rows.Scan(
			&verDoc.FileName,
			&verDoc.VersionNumber,
			&verDoc.Note,
			&verDoc.ModifyDate,
			&verDoc.Size,
			&verDoc.CRC,
			&verDoc.AuthorDisplayFld,
			&verDoc.EditorDisplayFld,
			&verDoc.TypeVersionData,
		); err != nil {
			log.LogError(msg.ErrEmptyTuple + " verDoc: " + err.Error())
			return nil, err
		}
		verDoc.FileName, _ = utils.Win1251ToUTF8(verDoc.FileName)
		verDoc.Note, _ = utils.Win1251ToUTF8(verDoc.Note)
		verDoc.AuthorDisplayFld, _ = utils.Win1251ToUTF8(verDoc.AuthorDisplayFld)
		verDoc.EditorDisplayFld, _ = utils.Win1251ToUTF8(verDoc.EditorDisplayFld)
		verDoc.TypeVersionData, _ = utils.Win1251ToUTF8(verDoc.TypeVersionData)

		versDoc = append(versDoc, verDoc)
	}

	if err = rows.Close(); err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return versDoc, nil
}

// OpenVerDoc - Открыть версию документа.
func OpenVerDoc(docId int, verId int) (model.TVersDoc, error) {
	rows, err := config.ConnectDB().Query("exec web_getFile ?, ?", docId, verId)
	if err != nil {
		log.LogError(msg.ErrQueryToDB + " OpenVerDoc(): " + err.Error())
		return model.TVersDoc{}, err
	}

	var docWithVer model.TVersDoc
	rows.Next()
	if err = rows.Scan(
		&docWithVer.VersionData,
		&docWithVer.Size,
		&docWithVer.TypeVersionData,
		&docWithVer.Number,
	); err != nil {
		log.LogError(msg.ErrEmptyTuple + " docWithVer: " + err.Error())
		return model.TVersDoc{}, err
	}

	if err = rows.Close(); err != nil {
		return model.TVersDoc{}, err
	}
	if err = rows.Err(); err != nil {
		return model.TVersDoc{}, err
	}

	return docWithVer, nil
}

// fillFolderAndDocsList - Заполнить список данными.
func fillFolderAndDocsList(rows *sql.Rows, err error) ([]model.TFolder, []model.TDoc, error) {
	var folders []model.TFolder
	for rows.Next() {
		var folder model.TFolder
		if err = rows.Scan(
			&folder.ObjectID,
			&folder.ObjectName,
			&folder.Created,
			&folder.Modified,
			&folder.AuthorID,
		); err != nil {
			log.LogError(msg.ErrEmptyTuple + " folder: " + err.Error())
		}
		folder.ObjectName, _ = utils.Win1251ToUTF8(folder.ObjectName)

		folders = append(folders, folder)
	}

	rows.NextResultSet()

	var docs []model.TDoc
	for rows.Next() {
		var doc model.TDoc
		if err = rows.Scan(
			&doc.ObjectID,
			&doc.ObjectName,
			&doc.Created,
			&doc.Modified,
			&doc.AuthorID,
			&doc.Signed,
			&doc.SignatureType,
		); err != nil {
			log.LogError(msg.ErrEmptyTuple + "doc: " + err.Error())
		}
		doc.ObjectName, _ = utils.Win1251ToUTF8(doc.ObjectName)
		doc.Signed, _ = utils.Win1251ToUTF8(doc.Signed)
		doc.SignatureType, _ = utils.Win1251ToUTF8(doc.SignatureType)

		docs = append(docs, doc)
	}

	if err = rows.Close(); err != nil {
		return nil, nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, nil, err
	}

	return folders, docs, nil
}
