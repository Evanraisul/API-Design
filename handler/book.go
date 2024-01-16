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
		e := m.Error{
			http.StatusBadRequest,
			"StatusBadRequest",
			"Invalid request body",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		//http.Error(w, "", http.StatusBadRequest)
		return
	}

	newBook.UUID = ulid.Make().String()

	curBook := reflect.ValueOf(newBook)

	for i := 0; i < curBook.NumField(); i++ {
		fieldValue := curBook.Field(i)

		if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
			e := m.Error{
				http.StatusBadRequest,
				"StatusBadRequest",
				"Invalid request body",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
			//http.Error(w, "", http.StatusBadRequest)
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
		e := m.Error{
			http.StatusBadRequest,
			"StatusBadRequest",
			"Invalid handler ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	var deletedBook m.Book
	if err := json.NewDecoder(r.Body).Decode(&deletedBook); err != nil {
		e := m.Error{
			http.StatusBadRequest,
			"StatusBadRequest",
			"Invalid request body",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	_, exists := Books[bookID.String()]
	if !exists {
		e := m.Error{
			http.StatusNotFound,
			"StatusBadRequest",
			"Book not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(e)

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
		e := m.Error{
			http.StatusBadRequest,
			"StatusBadRequest",
			"Invalid handler ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	var updatedBook m.Book
	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		e := m.Error{
			http.StatusBadRequest,
			"StatusBadRequest",
			"Invalid request body",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		return
	}

	_, exists := Books[bookID.String()]
	if !exists {
		e := m.Error{
			http.StatusNotFound,
			"StatusNotFound",
			"Book not found",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(e)
		return
	}

	updatedBook.UUID = bookID.String()

	curBook := reflect.ValueOf(updatedBook)

	for i := 0; i < curBook.NumField(); i++ {
		fieldValue := curBook.Field(i)

		if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
			e := m.Error{
				http.StatusBadRequest,
				"StatusBadRequest",
				"Invalid request body",
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
			//http.Error(w, "", http.StatusBadRequest)
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
		///
		e := m.Error{
			http.StatusBadRequest,
			"StatusBadRequest",
			"Invalid handler ID",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
		//http.Error(w, "Invalid handler ID", http.StatusBadRequest)
		return
	}

	book, exists := Books[bookID.String()]
	if !exists {
		e := m.Error{
			http.StatusNotFound,
			"StatusNotFound",
			"Invalid handler ID",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(e)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
