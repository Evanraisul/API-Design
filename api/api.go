package api

import (
	"github.com/evanraisul/book_api/handler"
	u "github.com/evanraisul/book_api/utils"
)

func API() {

	R, _ := u.Utils()

	R.Post("/api/v1/books", handler.CreateBook)
	R.Get("/api/v1/books/{id}", handler.GetBook)
	R.Get("/api/v1/books", handler.ListBooks)
	R.Put("/api/v1/books/{id}", handler.UpdateBook)
	R.Delete("/api/v1/books/{id}", handler.DeleteBook)
}
