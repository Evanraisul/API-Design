package model

import (
	"encoding/json"
	"fmt"
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
	err := json.NewEncoder(w).Encode(e)

	if err != nil {
		fmt.Println(err)
	}
}
