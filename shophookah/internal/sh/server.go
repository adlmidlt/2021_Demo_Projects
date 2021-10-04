package sh

import (
	"net/http"
	l "shophookah/pkg/logg"
	msg "shophookah/pkg/logg/message"
)

// Server - Сервер.
func Server() {
	l.LogI(msg.SuccessStartServer)
	if err := http.ListenAndServe(":8080", nil); CheckError(msg.ErrStartServer, err) {
		return
	}
}
