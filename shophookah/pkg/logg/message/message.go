package message

import "errors"

// SUCCESS
// ---DB
var (
	// SuccessStartServer - Сервер успешно запущен.
	SuccessStartServer = "Сервер успешно запущен"
	// SuccessConnToDB - Успешно подключён к БД.
	SuccessConnToDB = "Соединение с БД установлено"
	// SuccessCloseConnToDB - Соединение с БД успешно закрыто.
	SuccessCloseConnToDB = "Соединение с БД успешно закрыто"
)

// INFO
// ---MIGRATE
var (
	// InfoCurrentVersMigr - Информация о текущей версии миграции.
	InfoCurrentVersMigr = "Текущая версия миграции: %d"
)

// WARNING
// ---MIGRATE
var (
	// WarnMigrVersNotExist - Версия миграции не существует.
	WarnMigrVersNotExist = "Версия миграции: %d не существует или проверьте версию миграции БД"
	// WarnMigrFailedCommandDown - Команда "down" не выполнена.
	WarnMigrFailedCommandDown = "Команда \"down\" не выполнена"
	// WarnMigrFailedCommandUp - Команда "up" не выполнена.
	WarnMigrFailedCommandUp = "Команда \"up\" не выполнена"
)

// ERROR
// ---OTHER
var (
	// ErrConv - Ошибка конвертации.
	ErrConv = errors.New("ошибка конвертации")
)

// ---DB
var (
	// ErrConnToDB - Ошибка подключения к БД.
	ErrConnToDB = errors.New("ошибка подключения к БД")
	// ErrCloseDB - Ошибка закрытия БД.
	ErrCloseDB = errors.New("ошибка закрытия БД")
	// ErrExplicitOrImplicitClose - Ошибка после явного или неявного закрытия строк.
	ErrExplicitOrImplicitClose = errors.New("ошибка после явного или неявного закрытия строк")
	// ErrCloseRows - Ошибка закрытия строк, не удалось предотвратить дальнейшее перечисление.
	ErrCloseRows = errors.New("не удалось предотвратить дальнейшее перечисление")
	// ErrCreateInstanceDB - Ошибка создания экземпляра БД.
	ErrCreateInstanceDB = errors.New("ошибка создания экземпляра БД")
	// ErrNoRow - Не выбраны строки.
	ErrNoRow = errors.New("ошибка, строки не выбраны")
	// ErrRowsEmpty - Ошибка пустые строки.
	ErrRowsEmpty = errors.New("ошибка: пустые строки")
	// ErrQueryToDB - Ошибка запроса к БД.
	ErrQueryToDB = errors.New("не удалось выполнить запрос к БД")
)

// ---MIGRATE
var (
	// ErrFileWithMigrNotFound - Ошибка, файлы с миграциями не найдены, проверьте путь к файлам.
	ErrFileWithMigrNotFound = errors.New("ошибка, файлы с миграциями не найдены, проверьте путь к файлам")
	// ErrCreateNewMigrInstance - Ошибка создания нового экземпляра миграции.
	ErrCreateNewMigrInstance = errors.New("ошибка создания нового экземпляра миграции")
	// ErrCommandExecute - Ошибка выполнения команды.
	ErrCommandExecute = errors.New("не удалось запустить дерево команд, нет соответствующих совпадений для команд")
)

// ---SERVER
var (
	// ErrStartServer - Ошибка запуска сервера.
	ErrStartServer = errors.New("не удалось запустить сервер")
)

// ---TEMPLATE
var (
	// ErrParseTmpl = Ошибка парсинга шаблона.
	ErrParseTmpl = errors.New("ошибка парсинга шаблона")
	// ErrRegistrFailed -  Ошибка, регистрация не удалась.
	ErrRegistrFailed = errors.New("ошибка, регистрация не удалась")
)
