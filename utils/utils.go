package utils

import (
	"github.com/go-chi/chi"
)

func Server() (*chi.Mux, string) {
	R := chi.NewRouter()
	Port := "8080"

	return R, Port
}
