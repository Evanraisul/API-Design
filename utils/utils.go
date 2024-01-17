package utils

import (
	"github.com/go-chi/chi"
)

func Utils() (R *chi.Mux, Port string) {
	R = chi.NewRouter()
	Port = ":8080"

	return R, Port
}
