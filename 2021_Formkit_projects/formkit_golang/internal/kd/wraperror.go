package kd

import (
	l "cmd/pkg/logg"
	msg "cmd/pkg/logg/message"
	"database/sql"
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
