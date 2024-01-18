package api

import (
	"github.com/evanraisul/book_api/auth"
	"github.com/evanraisul/book_api/handler"
	"github.com/go-chi/chi"
)

func RoutesAddress(router *chi.Mux) {

	//router.Use(MyMiddleware)
	router.Post("/login", auth.LoginHandler)

	router.Group(func(r chi.Router) {
		r.Use(auth.VerifyJWT)
		r.Post("/api/v1/books", handler.CreateBook)
		r.Get("/api/v1/books/{id}", handler.GetBook)
		r.Get("/api/v1/books", handler.ListBooks)
		r.Put("/api/v1/books/{id}", handler.UpdateBook)
		r.Delete("/api/v1/books/{id}", handler.DeleteBook)
	})
}
