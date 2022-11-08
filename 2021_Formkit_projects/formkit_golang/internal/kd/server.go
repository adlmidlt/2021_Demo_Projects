package kd

import (
	l "cmd/pkg/logg"
	msg "cmd/pkg/logg/message"
	"net/http"
)

// Server - Сервер.
func Server() {
	l.LogI(msg.SuccessStartServer)
	if err := http.ListenAndServe(":8080", nil); check(msg.ErrStartServer, err) {
		return
	}
}
