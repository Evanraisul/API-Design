package utils

import (
	"github.com/go-chi/chi"
)

var R *chi.Mux = chi.NewRouter()
var Port string = ":8080"
