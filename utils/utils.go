package utils

import (
	"github.com/go-chi/chi"
)

func Utils() (*chi.Mux, string) {
	R := chi.NewRouter()
	Port := "8080"

	return R, Port
}
