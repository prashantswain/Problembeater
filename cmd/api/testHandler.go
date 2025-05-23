package main

import (
	"fmt"
	"net/http"
)

func (app *application) testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test")
	fmt.Println(r.ContentLength)
}
