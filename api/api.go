package api

import (
	"github.com/evanraisul/book_api/handler"
	u "github.com/evanraisul/book_api/utils"
)

func API() {

	u.R.Post("/api/v1/books", handler.CreateBook)
	u.R.Get("/api/v1/books/{id}", handler.GetBook)
	u.R.Get("/api/v1/books", handler.ListBooks)
	u.R.Put("/api/v1/books/{id}", handler.UpdateBook)
	u.R.Delete("/api/v1/books/{id}", handler.DeleteBook)
}
