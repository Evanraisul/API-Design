package api

import (
	"github.com/evanraisul/book_api/auth"
	"github.com/evanraisul/book_api/handler"
	"github.com/go-chi/chi"
)

func RoutesAddress(router *chi.Mux) {

	authMiddleware := auth.BasicAuth("realm", map[string]string{"admin": "password"})
	router.Use(authMiddleware)

	router.Post("/api/v1/books", handler.CreateBook)
	router.Get("/api/v1/books/{id}", handler.GetBook)
	router.Get("/api/v1/books", handler.ListBooks)
	router.Put("/api/v1/books/{id}", handler.UpdateBook)
	router.Delete("/api/v1/books/{id}", handler.DeleteBook)
}
