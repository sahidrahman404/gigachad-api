package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"errors"
	"time"

	gigachad "github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/privacy"
	"github.com/sahidrahman404/gigachad-api/internal/database"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	params := types.CreateUserParams{
		Username: input.Username,
		Email:    input.Email,
		Password: input.HashedPassword,
		Name:     input.Name,
	}

	v := validator.NewValidator()

	if params.ValidateUser(v); v.HasErrors() {
		return nil, r.errorMessage(v)
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return nil, err
	}

	err = r.storage.Users.Insert(user)

	if err != nil {
		return nil, err
	}

	token, err := r.storage.Tokens.New(user.Ent.ID, 3*24*time.Hour, database.ScopeActivation)

	if err != nil {
		return nil, err
	}

	r.backgroundTask(func() error {
		data := map[string]interface{}{
			"activationToken": token.Plaintext,
			"userID":          user.Ent.ID,
		}

		return r.mailer.Send(user.Ent.Email, data, "user_welcome.tmpl")
	})

	return user.Ent, nil
}

// ActivateUser is the resolver for the activateUser field.
func (r *mutationResolver) ActivateUser(ctx context.Context, input gigachad.ActivateUser) (*ent.User, error) {
	v := validator.NewValidator()

	if types.ValidateTokenPlaintext(v, input.TokenPlainText); v.HasErrors() {
		return nil, r.errorMessage(v)
	}

	user, err := r.storage.Users.GetForToken(database.ScopeActivation, input.TokenPlainText)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrRecordNotFound):
			v.AddFieldError("token", "invalid or expired activation token")
			return nil, r.errorMessage(v)
		default:
			return nil, r.serverError(err)
		}
	}

	user.Ent.Activated = 1

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	allow := privacy.DecisionContext(ctx, privacy.Allow)
	err = r.storage.Users.Update(user, allow)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrDuplicateUsername):
			v.AddFieldError("username", "a user with this username already exists")
			return nil, r.errorMessage(v)
		case errors.Is(err, database.ErrDuplicateEmail):
			v.AddFieldError("email", "a user with this email address already exists")
			return nil, r.errorMessage(v)
		case errors.Is(err, database.ErrEditConflict):
			return nil, r.editConflict()
		default:
			return nil, r.serverError(err)
		}
	}

	err = r.storage.Tokens.DeleteAllForUser(database.ScopeActivation, user.Ent.ID)
	if err != nil {
		return nil, r.serverError(err)
	}

	return user.Ent, nil
}

// UpdateUserPassword is the resolver for the updateUserPassword field.
func (r *mutationResolver) UpdateUserPassword(ctx context.Context, input gigachad.ResetUserPasswordParams) (*string, error) {
	v := validator.NewValidator()

	types.ValidatePasswordPlaintext(v, input.Password)
	types.ValidateTokenPlaintext(v, input.TokenPlainText)

	if v.HasErrors() {
		return nil, r.errorMessage(v)
	}

	user, err := r.storage.Users.GetForToken(database.ScopePasswordReset, input.TokenPlainText)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrRecordNotFound):
			v.AddFieldError("token", "invalid or expired password reset token")
			return nil, r.errorMessage(v)
		default:
			return nil, r.serverError(err)
		}
	}

	err = user.SetPassword(input.Password)
	if err != nil {
		return nil, r.serverError(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	allow := privacy.DecisionContext(ctx, privacy.Allow)
	err = r.storage.Users.Update(user, allow)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrDuplicateUsername):
			v.AddFieldError("username", "a user with this username already exists")
			return nil, r.errorMessage(v)
		case errors.Is(err, database.ErrDuplicateEmail):
			v.AddFieldError("email", "a user with this email address already exists")
			return nil, r.errorMessage(v)
		case errors.Is(err, database.ErrEditConflict):
			return nil, r.editConflict()
		default:
			return nil, r.serverError(err)
		}
	}

	err = r.storage.Tokens.DeleteAllForUser(database.ScopePasswordReset, user.Ent.ID)
	if err != nil {
		return nil, r.serverError(err)
	}

	message := "your password was successfully reset"
	return &message, nil
}

// Mutation returns gigachad.MutationResolver implementation.
func (r *Resolver) Mutation() gigachad.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
