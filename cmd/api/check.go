package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Application) CheckHandler() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/student/test", app.testHandler)
	return router
}
