package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/evanraisul/book_api/api"
	"github.com/evanraisul/book_api/model"
	"github.com/go-chi/chi"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func executeRequest(req *http.Request, s *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int, req model.Book, res model.Book) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
	reqBook := reflect.ValueOf(req)
	resBook := reflect.ValueOf(res)

	for i := 1; i < reqBook.NumField(); i++ {
		fieldValue1 := reqBook.Field(i)
		fieldValue2 := resBook.Field(i)

		if reflect.DeepEqual(fieldValue1, fieldValue2) {
			t.Errorf("unexpected Error")
		}
	}
}

func TestCreateBook(t *testing.T) {
	router := api.GetNewRoutes()
	api.RoutesAddress(router)

	token, err := getToken(router)
	if err != nil {
		t.Fatal(err)
	}
	reqBody := map[string]interface{}{
		"name":        "TestBook",
		"authorList":  []string{"Author1", "Author2"},
		"publishDate": "2022-01-01",
		"isbn":        "1234567890",
	}
	reqJSON, _ := json.Marshal(reqBody)
	req, er := http.NewRequest("POST", "/api/v1/books", bytes.NewBuffer(reqJSON))
	if er != nil {
		t.Fatal(er)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	response := executeRequest(req, router)

	b, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	var book model.Book
	e := json.Unmarshal([]byte(b), &book)
	if e != nil {
		fmt.Println(e)
	}

	b1, err2 := io.ReadAll(bytes.NewBuffer(reqJSON))
	if err2 != nil {
		t.Fatal(err2)
	}

	var book1 model.Book
	er1 := json.Unmarshal([]byte(b1), &book1)
	if er1 != nil {
		fmt.Println(er1)
	}

	checkResponseCode(t, http.StatusOK, response.Code, book1, book)
}

func getToken(router *chi.Mux) (string, error) {
	reqBody := map[string]string{"username": "admin", "password": "admin"}
	reqJSON, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", err
	}
	response := executeRequest(req, router)
	b, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}
	return string(b), nil
}
