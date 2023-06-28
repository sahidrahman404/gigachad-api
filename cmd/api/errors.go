package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/sahidrahman404/gigachad-api/internal/response"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

func (app *application) reportError(err error) {
	trace := debug.Stack()

	app.logger.Error(err, trace)
}

func (app *application) errorMessage(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	message string,
	headers http.Header,
) {
	message = strings.ToUpper(message[:1]) + message[1:]

	err := response.JSONWithHeaders(w, status, map[string]string{"Error": message}, headers)
	if err != nil {
		app.reportError(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.reportError(err)

	message := "The server encountered a problem and could not process your request"
	app.errorMessage(w, r, http.StatusInternalServerError, message, nil)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	app.errorMessage(w, r, http.StatusNotFound, message, nil)
}

func (app *application) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	app.errorMessage(w, r, http.StatusMethodNotAllowed, message, nil)
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.errorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
}

func (app *application) failedValidation(
	w http.ResponseWriter,
	r *http.Request,
	v *validator.Validator,
) {
	err := response.JSON(w, http.StatusUnprocessableEntity, v)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) editConflict(w http.ResponseWriter, r *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	app.errorMessage(w, r, http.StatusConflict, message, nil)
}

func (app *application) invalidCredentials(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	app.errorMessage(w, r, http.StatusUnauthorized, message, nil)
}

func (app *application) invalidAuthenticationToken(
	w http.ResponseWriter,
	r *http.Request,
) {
	headers := make(http.Header)
	headers.Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	app.errorMessage(w, r, http.StatusUnauthorized, message, headers)
}

func (app *application) authenticationRequired(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	app.errorMessage(w, r, http.StatusUnauthorized, message, nil)
}

func (app *application) inactiveAccount(w http.ResponseWriter, r *http.Request) {
	message := "your user account must be activated to access this resource"
	app.errorMessage(w, r, http.StatusForbidden, message, nil)
}

func (app *application) notPremittedError(w http.ResponseWriter, r *http.Request) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	app.errorMessage(w, r, http.StatusForbidden, message, nil)
}
