package handler

import (
	"encoding/json"
	m "github.com/evanraisul/book_api/model"
	"github.com/go-chi/chi"
	"github.com/oklog/ulid/v2"
	"net/http"
	"reflect"
)

var Books = make(map[string]m.Book)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook m.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		(&m.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Resource not found")
		return
	}

	newBook.UUID = ulid.Make().String()

	curBook := reflect.ValueOf(newBook)

	for i := 0; i < curBook.NumField(); i++ {
		fieldValue := curBook.Field(i)

		if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
			(&m.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Invalid Request")
			return
		}
	}

	Books[newBook.UUID] = newBook

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := ulid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		(&m.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Parsing Error")
		return
	}

	var deletedBook m.Book
	if err := json.NewDecoder(r.Body).Decode(&deletedBook); err != nil {
		(&m.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Invalid Request Body")
		return
	}

	_, exists := Books[bookID.String()]
	if !exists {
		(&m.Error{}).GetError(w, http.StatusNotFound, "StatusNotFound", "Book Not Found")
		return
	}

	deletedBook.UUID = bookID.String()
	Books[bookID.String()] = deletedBook

	delete(Books, bookID.String())

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deletedBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := ulid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		(&m.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Parsing Error")
		return
	}

	var updatedBook m.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		(&m.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Invalid request body")
		return
	}

	_, exists := Books[bookID.String()]
	if !exists {
		(&m.Error{}).GetError(w, http.StatusNotFound, "StatusNotFound", "Book not found")
		return
	}

	updatedBook.UUID = bookID.String()

	curBook := reflect.ValueOf(updatedBook)

	for i := 0; i < curBook.NumField(); i++ {
		fieldValue := curBook.Field(i)

		if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
			(&m.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Invalid Request")
			return
		}
	}
	Books[bookID.String()] = updatedBook

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedBook)
}

func ListBooks(w http.ResponseWriter, r *http.Request) {
	bookList := make([]m.Book, 0, len(Books))
	for _, book := range Books {
		bookList = append(bookList, book)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookList)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := ulid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		(&m.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Parsing Error")
		return
	}

	book, exists := Books[bookID.String()]
	if !exists {
		(&m.Error{}).GetError(w, http.StatusNotFound, "StatusNotFound", "Invalid handler ID")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
