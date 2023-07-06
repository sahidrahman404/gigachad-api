package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"errors"
	"time"

	gigachad "github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/internal/database"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

// CreateAuthenticationToken is the resolver for the createAuthenticationToken field.
func (r *mutationResolver) CreateAuthenticationToken(ctx context.Context, input gigachad.Login) (*string, error) {
	v := validator.NewValidator()

	types.ValidateEmail(v, input.Email)
	types.ValidatePasswordPlaintext(v, input.Password)

	if v.HasErrors() {
		return nil, r.errorMessage(v)
	}

	user, err := r.storage.Users.GetByEmail(input.Email)

	if err != nil {
		switch {
		case errors.Is(err, database.ErrRecordNotFound):
			return nil, r.invalidCredentials()
		default:
			return nil, r.serverError(err)
		}
	}

	match, err := user.Matches(input.Password)
	if err != nil {
		return nil, r.serverError(err)
	}

	if !match {
		return nil, r.invalidCredentials()
	}

	token, err := r.storage.Tokens.New(user.Ent.ID, 24*time.Hour, database.SocpeAuthentication)
	if err != nil {
		return nil, r.serverError(err)
	}
	return &token.Plaintext, nil
}

// CreateActivationToken is the resolver for the createActivationToken field.
func (r *mutationResolver) CreateActivationToken(ctx context.Context, input gigachad.ActivationTokenParams) (*string, error) {
	v := validator.NewValidator()

	if types.ValidateEmail(v, input.Email); v.HasErrors() {
		return nil, r.errorMessage(v)
	}

	user, err := r.storage.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrRecordNotFound):
			return nil, r.invalidCredentials()
		default:
			return nil, r.serverError(err)
		}
	}

	if user.Ent.Activated == 1 {
		v.AddFieldError("email", "user has already been activated")
		return nil, r.errorMessage(v)
	}

	token, err := r.storage.Tokens.New(user.Ent.ID, 3*24*time.Hour, database.ScopeActivation)
	if err != nil {
		return nil, r.serverError(err)
	}

	r.backgroundTask(func() error {
		data := map[string]interface{}{
			"activationToken": token.Plaintext,
		}

		return r.mailer.Send(user.Ent.Email, data, "token_activation.tmpl")
	})

	message := "an email will be sent to you containing activation instructions"

	return &message, nil
}

// CreatePasswordResetToken is the resolver for the createPasswordResetToken field.
func (r *mutationResolver) CreatePasswordResetToken(ctx context.Context, input gigachad.ResetPasswordParams) (*string, error) {
	v := validator.NewValidator()

	if types.ValidateEmail(v, input.Email); v.HasErrors() {
		return nil, r.errorMessage(v)
	}

	user, err := r.storage.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrRecordNotFound):
			v.AddFieldError("email", "no matching email address found")
			return nil, r.errorMessage(v)
		default:
			return nil, r.serverError(err)
		}
	}

	if user.Ent.Activated == 0 {
		v.AddFieldError("email", "user account must be activated")
		return nil, r.errorMessage(v)
	}

	token, err := r.storage.Tokens.New(user.Ent.ID, 45*time.Minute, database.ScopePasswordReset)
	if err != nil {
		return nil, r.serverError(err)
	}

	r.backgroundTask(func() error {
		data := map[string]interface{}{
			"passwordResetToken": token.Plaintext,
		}

		return r.mailer.Send(user.Ent.Email, data, "token_password_reset.tmpl")
	})

	message := "an email will be sent to you containing password instruction"
	return &message, nil
}
