package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/sahidrahman404/gigachad-api/internal/database"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				app.serverError(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			r = app.contextSetUser(r, types.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			app.invalidCredentials(w, r)
			return
		}

		token := headerParts[1]

		v := validator.NewValidator()

		if types.ValidateTokenPlaintext(v, token); v.HasErrors() {
			app.invalidAuthenticationToken(w, r)
			return
		}

		user, err := app.storage.Users.GetForToken(database.SocpeAuthentication, token)
		if err != nil {
			switch {
			case errors.Is(err, database.ErrRecordNotFound):
				app.invalidCredentials(w, r)
			default:
				app.serverError(w, r, err)
			}
			return
		}

		r = app.contextSetUser(r, user)

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuthenticatedUser(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := app.contextGetUser(r)

		if user.IsAnonymous() {
			app.authenticationRequired(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireActivatedUser(next http.HandlerFunc) http.HandlerFunc {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := app.contextGetUser(r)

		if user.Ent.Activated == 1 {
			app.inactiveAccount(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
	return app.requireAuthenticatedUser(fn)
}
