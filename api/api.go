package api

import (
	"github.com/evanraisul/book_api/handler/book"
	"github.com/evanraisul/book_api/pkg/auth"
	"github.com/go-chi/chi"
)

func GetNewRoutes() *chi.Mux {
	return chi.NewRouter()
}

func RoutesAddress(router *chi.Mux) {
	router.Post("/login", auth.LoginHandler)
	router.Group(func(r chi.Router) {
		r.Use(auth.VerifyJWT)
		r.Post("/api/v1/books", book.CreateBook)
		r.Get("/api/v1/books/{id}", book.GetBook)
		r.Get("/api/v1/books", book.ListBooks)
		r.Put("/api/v1/books/{id}", book.UpdateBook)
		r.Delete("/api/v1/books/{id}", book.DeleteBook)
	})
}
