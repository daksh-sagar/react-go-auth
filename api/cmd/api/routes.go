package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/authenticate", app.authenticate)
	router.HandlerFunc(http.MethodGet, "/v1/refreshToken", app.refreshToken)
	router.HandlerFunc(http.MethodGet, "/v1/logout", app.logout)

	return router

}
