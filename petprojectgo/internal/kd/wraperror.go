package kd

import (
	"database/sql"
	l "petprojectgo/pkg/logg"
	msg "petprojectgo/pkg/logg/message"
)

// Check - Проверить.
func check(msg, err error) bool {
	if err != nil {
		l.LogE(msg, err.Error())
	}
	return false
}

// closeRows - Закрыть строки, предотвращающие дальнейшие перечисления.
func closeRows(rows *sql.Rows) {
	func(rows *sql.Rows) {
		err := rows.Close()
		check(msg.ErrCloseRows, err)
	}(rows)
}
