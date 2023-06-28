package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"entgo.io/ent/privacy"
	"github.com/sahidrahman404/gigachad-api/internal/database"
	"github.com/sahidrahman404/gigachad-api/internal/request"
	"github.com/sahidrahman404/gigachad-api/internal/response"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var params types.CreateUserParams

	err := request.DecodeJSONStrict(w, r, &params)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	v := validator.NewValidator()

	if params.ValidateUser(v); v.HasErrors() {
		fmt.Println(v)
		app.failedValidation(w, r, v)
		return
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = app.storage.Users.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrDuplicateUsername):
			v.AddFieldError("username", "a user with this username already exists")
			app.failedValidation(w, r, v)
		case errors.Is(err, database.ErrDuplicateEmail):
			v.AddFieldError("email", "a user with this email address already exists")
			app.failedValidation(w, r, v)
		default:
			app.serverError(w, r, err)
		}
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
			"userID":          user.Ent.ID,
		}

		return app.mailer.Send(user.Ent.Email, data, "user_welcome.tmpl")
	})

	err = response.JSON(w, http.StatusAccepted, user)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		TokenPlainText string `json:"token"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	v := validator.NewValidator()

	if types.ValidateTokenPlaintext(v, input.TokenPlainText); v.HasErrors() {
		app.failedValidation(w, r, v)
		return
	}

	user, err := app.storage.Users.GetForToken(database.ScopeActivation, input.TokenPlainText)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrRecordNotFound):
			v.AddFieldError("token", "invalid or expired activation token")
			app.failedValidation(w, r, v)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	user.Ent.Activated = 1

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	allow := privacy.DecisionContext(ctx, privacy.Allow)
	err = app.storage.Users.Update(user, allow)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrDuplicateUsername):
			v.AddFieldError("username", "a user with this username already exists")
			app.failedValidation(w, r, v)
		case errors.Is(err, database.ErrDuplicateEmail):
			v.AddFieldError("email", "a user with this email address already exists")
			app.failedValidation(w, r, v)
		case errors.Is(err, database.ErrEditConflict):
			app.editConflict(w, r)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	err = app.storage.Tokens.DeleteAllForUser(database.ScopeActivation, user.Ent.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, user)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) updateUserPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Password       string `json:"password"`
		TokenPlainText string `json:"token"`
	}

	err := request.DecodeJSON(w, r, &input)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	v := validator.NewValidator()

	types.ValidatePasswordPlaintext(v, input.Password)
	types.ValidateTokenPlaintext(v, input.TokenPlainText)

	if v.HasErrors() {
		app.failedValidation(w, r, v)
		return
	}

	user, err := app.storage.Users.GetForToken(database.ScopePasswordReset, input.TokenPlainText)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrRecordNotFound):
			v.AddFieldError("token", "invalid or expired password reset token")
			app.failedValidation(w, r, v)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	err = user.SetPassword(input.Password)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	allow := privacy.DecisionContext(ctx, privacy.Allow)
	err = app.storage.Users.Update(user, allow)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrDuplicateUsername):
			v.AddFieldError("username", "a user with this username already exists")
			app.failedValidation(w, r, v)
		case errors.Is(err, database.ErrDuplicateEmail):
			v.AddFieldError("email", "a user with this email address already exists")
			app.failedValidation(w, r, v)
		case errors.Is(err, database.ErrEditConflict):
			app.editConflict(w, r)
		default:
			app.serverError(w, r, err)
		}
		return
	}

	err = app.storage.Tokens.DeleteAllForUser(database.ScopePasswordReset, user.Ent.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, "your password was successfully reset")
	if err != nil {
		app.serverError(w, r, err)
	}
}
