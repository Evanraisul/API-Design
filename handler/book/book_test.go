package book_test

import (
	"bytes"
	"encoding/json"
	"github.com/evanraisul/book_api/api"
	"github.com/go-chi/chi"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func executeRequest(req *http.Request, s *chi.Mux) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
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

	checkResponseCode(t, http.StatusOK, response.Code)
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
