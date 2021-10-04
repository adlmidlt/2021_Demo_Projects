package kd

import (
	"net/http"
	l "petprojectgo/pkg/logg"
	msg "petprojectgo/pkg/logg/message"
)

// Server - Сервер.
func Server() {
	l.LogI(msg.SuccessStartServer)
	if err := http.ListenAndServe(":8080", nil); check(msg.ErrStartServer, err) {
		return
	}
}
