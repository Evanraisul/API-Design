package main

import (
	a "github.com/evanraisul/book_api/api"
	u "github.com/evanraisul/book_api/utils"
	"net/http"
)

func main() {

	a.API()
	http.ListenAndServe(u.Port, u.R)
}
