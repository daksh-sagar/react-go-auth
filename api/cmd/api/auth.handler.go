package main

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	user, err := app.models.Users.GetUserByEmail(payload.Email)
	if err != nil {
		app.badRequestError(w, r, err)
		return
	}

	isValidPwd, err := user.ValidatePassword(payload.Password)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if !isValidPwd {
		app.badRequestError(w, r, err)
		return
	}

	u := jwtUser{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	http.SetCookie(w, refreshCookie)

	app.writeJSON(w, http.StatusAccepted, envelope{"tokens": tokens}, nil)
}

func (app *application) refreshToken(w http.ResponseWriter, r *http.Request) {
	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			refreshToken := cookie.Value
			claims := &Claims{}

			_, err := jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (any, error) {
				return []byte(app.auth.Secret), nil
			})
			if err != nil {
				app.unauthorizedError(w, r)
				return
			}

			userId, err := strconv.ParseInt(claims.Subject, 10, 64)

			user, err := app.models.Users.GetUserById(userId)

			if err != nil {
				app.unauthorizedError(w, r)
				return
			}

			u := jwtUser{
				Id:        user.Id,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}

			tokenPairs, err := app.auth.GenerateTokenPair(&u)
			if err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}

			http.SetCookie(w, app.auth.GetRefreshCookie(tokenPairs.RefreshToken))
			app.writeJSON(w, http.StatusAccepted, envelope{"tokens": tokenPairs}, nil)
		}

	}
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
	w.WriteHeader(http.StatusAccepted)
}
