package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/sahidrahman404/gigachad-api/internal/database"
	"github.com/sahidrahman404/gigachad-api/internal/request"
	"github.com/sahidrahman404/gigachad-api/internal/response"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

func (app *application) createPasswordResetTokenHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	v := validator.NewValidator()

	if types.ValidateEmail(v, input.Email); v.HasErrors() {
		app.failedValidation(w, r, v)
		return
	}

	user, err := app.storage.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrRecordNotFound):
			v.AddFieldError("email", "no matching email address found")
			app.failedValidation(w, r, v)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	if !user.Ent.Activated {
		v.AddFieldError("email", "user account must be activated")
		app.failedValidation(w, r, v)
		return
	}

	token, err := app.storage.Tokens.New(user.Ent.ID, 45*time.Minute, database.ScopePasswordReset)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.backgroundTask(func() error {
		data := map[string]interface{}{
			"passwordResetToken": token.Plaintext,
		}

		return app.mailer.Send(user.Ent.Email, data, "token_password_reset.tmpl")
	})

	err = response.JSON(
		w,
		http.StatusAccepted,
		"an email will be sent to you containing password instruction",
	)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) createActivationTokenHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	v := validator.NewValidator()

	if types.ValidateEmail(v, input.Email); v.HasErrors() {
		app.failedValidation(w, r, v)
		return
	}

	user, err := app.storage.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrRecordNotFound):
			app.invalidCredentials(w, r)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	if user.Ent.Activated {
		v.AddFieldError("email", "user has already been activated")
		app.failedValidation(w, r, v)
		return
	}

	token, err := app.storage.Tokens.New(user.Ent.ID, 3*24*time.Hour, database.ScopeActivation)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.backgroundTask(func() error {
		data := map[string]interface{}{
			"activationToken": token.Plaintext,
		}

		return app.mailer.Send(user.Ent.Email, data, "token_activation.tmpl")
	})

	err = response.JSON(
		w,
		http.StatusAccepted,
		"an email will be sent to you containing activation instructions",
	)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) createAuthenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	v := validator.NewValidator()

	types.ValidateEmail(v, input.Email)
	types.ValidatePasswordPlaintext(v, input.Password)

	if !v.HasErrors() {
		app.failedValidation(w, r, v)
		return
	}

	user, err := app.storage.Users.GetByEmail(input.Email)

	if err != nil {
		switch {
		case errors.Is(err, database.ErrRecordNotFound):
			app.invalidCredentials(w, r)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	match, err := user.Matches(input.Password)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	if !match {
		app.invalidCredentials(w, r)
		return
	}

	token, err := app.storage.Tokens.New(user.Ent.ID, 24*time.Hour, database.SocpeAuthentication)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusCreated, token)
	if err != nil {
		app.serverError(w, r, err)
	}
}
