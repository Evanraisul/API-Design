package handler

import (
	"encoding/json"
	"fmt"
	"github.com/evanraisul/book_api/model"
	"github.com/go-chi/chi"
	"github.com/oklog/ulid/v2"
	"net/http"
	"reflect"
)

var Books = make(map[string]model.Book)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook model.Book

	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Resource not found")
		return
	}

	newBook.UUID = ulid.Make().String()

	curBook := reflect.ValueOf(newBook)

	for i := 0; i < curBook.NumField(); i++ {
		fieldValue := curBook.Field(i)

		if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
			(&model.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Invalid Request")
			return
		}
	}

	Books[newBook.UUID] = newBook

	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w).Encode(newBook)
	if e != nil {
		fmt.Println(e)
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := ulid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Parsing Error")
		return
	}

	var deletedBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&deletedBook); err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Invalid Request Body")
		return
	}

	_, exists := Books[bookID.String()]
	if !exists {
		(&model.Error{}).GetError(w, http.StatusNotFound, "StatusNotFound", "Book Not Found")
		return
	}

	deletedBook.UUID = bookID.String()
	Books[bookID.String()] = deletedBook

	delete(Books, bookID.String())

	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w).Encode(deletedBook)
	if e != nil {
		fmt.Println(e)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := ulid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Parsing Error")
		return
	}

	var updatedBook model.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Invalid request body")
		return
	}

	_, exists := Books[bookID.String()]
	if !exists {
		(&model.Error{}).GetError(w, http.StatusNotFound, "StatusNotFound", "Book not found")
		return
	}

	updatedBook.UUID = bookID.String()

	curBook := reflect.ValueOf(updatedBook)

	for i := 0; i < curBook.NumField(); i++ {
		fieldValue := curBook.Field(i)

		if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
			(&model.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Invalid Request")
			return
		}
	}
	Books[bookID.String()] = updatedBook

	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w).Encode(updatedBook)
	if e != nil {
		fmt.Println(e)
	}
}

func ListBooks(w http.ResponseWriter, r *http.Request) {
	bookList := make([]model.Book, 0, len(Books))
	for _, book := range Books {
		bookList = append(bookList, book)
	}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(bookList)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := ulid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		(&model.Error{}).GetError(w, http.StatusBadRequest, "StatusBadRequest", "Parsing Error")
		return
	}

	book, exists := Books[bookID.String()]
	if !exists {
		(&model.Error{}).GetError(w, http.StatusNotFound, "StatusNotFound", "Invalid handler ID")
		return
	}

	w.WriteHeader(http.StatusOK)
	e := json.NewEncoder(w).Encode(book)
	if e != nil {
		fmt.Println(e)
	}
}
