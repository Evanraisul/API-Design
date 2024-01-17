package model

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	ErrorCode int
	ErrorType string
	Message   string
}

func (e *Error) GetError(w http.ResponseWriter, code int, typ string, msg string) {
	e.ErrorCode = code
	e.ErrorType = typ
	e.Message = msg

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(e)
}
