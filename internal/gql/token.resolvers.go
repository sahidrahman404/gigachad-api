package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/cschleiden/go-workflows/client"
	gigachad "github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/activity"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/database"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
	"github.com/sahidrahman404/gigachad-api/workflow"
)

// CreateAuthenticationToken is the resolver for the createAuthenticationToken field.
func (r *mutationResolver) CreateAuthenticationToken(ctx context.Context, input gigachad.LoginInput) (*gigachad.AuthenticationToken, error) {
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
	return &gigachad.AuthenticationToken{
		User:           user.Ent,
		TokenPlainText: token.Plaintext,
	}, nil
}

// CreateActivationToken is the resolver for the createActivationToken field.
func (r *mutationResolver) CreateActivationToken(ctx context.Context, input gigachad.ActivationTokenInput) (*ent.User, error) {
	v := validator.NewValidator()

	if types.ValidateEmail(v, input.Email); v.HasErrors() {
		return nil, r.errorMessage(v)
	}

	user, err := r.storage.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrRecordNotFound):
			return nil, r.invalidEmail()
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

	data := map[string]interface{}{
		"activationToken": token.Plaintext,
		"name":            strings.Split(user.Ent.Name, " ")[0],
	}

	m := activity.MailDetails{
		Recipient: user.Ent.Email,
		Data:      data,
		Patterns:  []string{"token_activation.tmpl"},
	}

	options := client.WorkflowInstanceOptions{
		InstanceID: string(user.Ent.ID),
	}

	c := client.New(r.workflowBackend)
	_, _ = c.CreateWorkflowInstance(context.Background(), options, workflow.SendEmail, m)

	return user.Ent, nil
}

// CreatePasswordResetToken is the resolver for the createPasswordResetToken field.
func (r *mutationResolver) CreatePasswordResetToken(ctx context.Context, input gigachad.ResetPasswordInput) (*ent.User, error) {
	v := validator.NewValidator()

	if types.ValidateEmail(v, input.Email); v.HasErrors() {
		return nil, r.errorMessage(v)
	}

	user, err := r.storage.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrRecordNotFound):
			return nil, r.invalidEmail()
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

	return user.Ent, nil
}
