package utils

const (
	DEFAULT_PORT = "8080"
)

type ErrorType string

const (
	StatusBadRequest ErrorType = "StatusBadRequest"
	StatusNotFound   ErrorType = "StatusNotFound"
)
