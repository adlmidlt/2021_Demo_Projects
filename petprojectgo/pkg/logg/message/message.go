package message

import "errors"

var (
	// SuccessStartServer - Сервер успешно запущен.
	SuccessStartServer = "сервер успешно запущен"
	// SuccessConnToDB - Успешно подключён к БД.
	SuccessConnToDB = "успешно подключён к БД"
)

var (
	// ErrStartServer - Ошибка запуска сервера.
	ErrStartServer = errors.New("не удалось запустить сервер")
	// ErrConnToDB - Ошибка подключения к БД.
	ErrConnToDB = errors.New("не удалось подключиться к БД")
	// ErrQueryToDB - Ошибка запроса к БД.
	ErrQueryToDB = errors.New("не удалось выполнить запрос к БД")
	// ErrRowsEmpty - Ошибка пустые строки.
	ErrRowsEmpty = errors.New("ошибка: пустые строки")
	// ErrParseFile - Ошибка парсинга файла.
	ErrParseFile = errors.New("не удалось распарсить файл")
	// ErrParseTpl - Ошибка парсинга шаблона.
	ErrParseTpl = errors.New("не удалось распарсить шаблон")
	// ErrConv - Ошибка конвертации.
	ErrConv = errors.New("ошибка конвертации")
	// ErrFormat - Ошибка формата документа, ожидался PDF формат.
	ErrFormat = errors.New("ошибка формата документа, ожидался PDF формат")
	// ErrParseJson - Ошибка парсинга Json.
	ErrParseJson = errors.New("не удалось распарсить Json")
	// ErrCloseRows - Ошибка закрытия строк, не удалось предотвратить дальнейшее перечисление.
	ErrCloseRows = errors.New("не удалось предотвратить дальнейшее перечисление")
	// ErrIterRows - Ошибка возникла во время итерации.
	ErrIterRows = errors.New("ошибка возникла во время итерации строк")
	// ErrReadFolderContent - Ошибка чтения содержимое папки.
	ErrReadFolderContent = errors.New("не удалось прочитать содержимое папки")
	// ErrNotFindFolderContentByIdFolder - Ошибка, не удалось найти содержимое папки по ИД папки.
	ErrNotFindFolderContentByIdFolder = errors.New("не удалось найти содержимое папки по ИД папки")
)
