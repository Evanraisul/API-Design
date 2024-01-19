package model

import (
	"encoding/json"
	"fmt"
	"github.com/evanraisul/book_api/utils"
	"net/http"
)

type Error struct {
	ErrorCode int             `json:"errorCode"`
	ErrorType utils.ErrorType `json:"errorType"`
	Message   string          `json:"message"`
}

func (e *Error) GetError(w http.ResponseWriter, code int, typ utils.ErrorType, msg string) {
	e.ErrorCode = code
	e.ErrorType = typ
	e.Message = msg

	w.WriteHeader(http.StatusBadRequest)
	err := json.NewEncoder(w).Encode(e)

	if err != nil {
		fmt.Println(err)
	}
}
