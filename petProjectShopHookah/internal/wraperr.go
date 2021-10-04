package internal

import (
	l "shophookah/pkg/logg"
)

// CheckError - Проверить.
func CheckError(msg, err error) bool {
	if err != nil {
		l.LogE(msg, err.Error())
	}
	return false
}
