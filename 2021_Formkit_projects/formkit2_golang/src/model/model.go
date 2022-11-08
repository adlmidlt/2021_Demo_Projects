package model

// WrapTFolderAndTDoc - Обёртка для папок и документов.
type WrapTFolderAndTDoc struct {
	Folders   []TFolder
	Documents []TDoc
}

// WrapTDocWithVers - Обёртика для документов с версиями.
type WrapTDocWithVers struct {
	DocumentWithVersions []TDocWithVers
}

// TFolder - Папка.
type TFolder struct {
	ObjectID   int    `json:"ObjectID"`
	ObjectName string `json:"ObjectName"`
	Created    string `json:"Created"`
	Modified   string `json:"Modified"`
	AuthorID   int    `json:"AuthorID"`
}

// TDoc - Документ.
type TDoc struct {
	ObjectID      int    `json:"ObjectID"`
	ObjectName    string `json:"ObjectName"`
	Created       string `json:"Created"`
	Modified      string `json:"Modified"`
	AuthorID      string `json:"AuthorID"`
	Signed        string `json:"Signed"`
	SignatureType string `json:"SignatureType"`
}

// TDocWithVers - Документы с версией(ями).
type TDocWithVers struct {
	FileName         string `json:"FileName"`
	VersionNumber    int    `json:"VersionNumber"`
	Note             string `json:"Note"`
	ModifyDate       string `json:"ModifyDate"`
	Size             int    `json:"Size"`
	CRC              string `json:"CRC"`
	AuthorDisplayFld string `json:"AuthorDisplayFld"`
	EditorDisplayFld string `json:"EditorDisplayFld"`
	TypeVersionData  string `json:"TypeVersionData"`
}

// TVersDoc - Версия(и) документа.
type TVersDoc struct {
	VersionData     string `json:"VersionData"`
	Size            int    `json:"Size"`
	TypeVersionData string `json:"TypeVersionData"`
	Number          int    `json:"Number"`
}
