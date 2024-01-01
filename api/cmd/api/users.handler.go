package main

import (
	"net/http"

	"github.com/daksh-sagar/react-go-auth/api/internal/data"
)

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	err := app.readJSON(w, r, &params)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	user := &data.User{
		Email:     params.Email,
		Password:  params.Password,
		FirstName: params.FirstName,
		LastName:  params.LastName,
	}

	user, err = app.models.Users.Insert(user)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
