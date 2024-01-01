package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.requireAuth(app.healthcheckHandler))
	router.HandlerFunc(http.MethodPost, "/v1/authenticate", app.authenticate)
	router.HandlerFunc(http.MethodGet, "/v1/refreshToken", app.refreshToken)
	router.HandlerFunc(http.MethodGet, "/v1/logout", app.logout)
	router.HandlerFunc(http.MethodPost, "/v1/signup", app.createUserHandler)

	return app.enableCORS(router)

}
