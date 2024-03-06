package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"crypto/sha256"
	"time"

	"github.com/cschleiden/go-workflows/client"
	gigachad "github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/activity"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/token"
	userPredicate "github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/internal/database"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	caser := cases.Title(language.AmericanEnglish)
	params := types.CreateUserParams{
		Username: input.Username,
		Email:    input.Email,
		Password: input.HashedPassword,
		Name:     caser.String(input.Name),
	}

	v := validator.NewValidator()

	if params.ValidateUser(v); v.HasErrors() {
		return nil, r.errorMessage(v)
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return nil, err
	}

	err = r.WithTx(ctx, func(tx *ent.Tx) error {
		txClient := tx.Client()
		u, err := txClient.User.Create().
			SetUsername(user.Ent.Username).
			SetEmail(user.Ent.Email).
			SetHashedPassword(user.Ent.HashedPassword).
			SetName(user.Ent.Name).
			Save(ctx)
		if err != nil {
			return err
		}

		t, err := types.GenerateToken(u.ID, 3*24*time.Hour, database.ScopeActivation)
		if err != nil {
			return err
		}

		err = txClient.Token.Create().
			SetExpiry(t.Ent.Expiry).
			SetHash(t.Ent.Hash).
			SetScope(t.Ent.Scope).
			SetUserID(t.Ent.UserID).
			Exec(ctx)

		if err != nil {
			return err
		}

		caser := cases.Title(language.Indonesian)
		name := caser.String(user.Ent.Name)

		r.backgroundTask(func() error {
			data := map[string]interface{}{
				"activationToken": t.Plaintext,
				"name":            name,
			}

			return r.mailer.Send(user.Ent.Email, data, "user_welcome.tmpl")
		})

		return nil
	})

	if err != nil {
		switch {
		case ent.IsConstraintError(err):
			v.AddFieldError("conflict", "a user with this email or username address already exists")
			return nil, r.errorMessage(v)
		default:
			return nil, r.serverError(err)
		}
	}

	return user.Ent, nil
}

// ActivateUser is the resolver for the activateUser field.
func (r *mutationResolver) ActivateUser(ctx context.Context, input gigachad.ActivateUserInput) (*gigachad.ActivateUserResult, error) {
	v := validator.NewValidator()

	if types.ValidateTokenPlaintext(v, input.TokenPlainText); v.HasErrors() {
		return nil, r.errorMessage(v)
	}

	err := r.WithTx(ctx, func(tx *ent.Tx) error {
		txClient := tx.Client()

		tokenHash := sha256.Sum256([]byte(input.TokenPlainText))
		u, err := txClient.User.Query().
			Where(
				userPredicate.HasTokensWith(token.Hash(tokenHash[:]),
					token.Scope(database.ScopeActivation),
					token.ExpiryGT(time.Now().Format(time.RFC3339)),
				),
			).
			Only(ctx)
		if err != nil {
			return err
		}

		u.Activated = 1
		u.Version = u.Version + 1

		_, err = txClient.User.UpdateOneID(u.ID).
			SetUsername(u.Username).
			SetEmail(u.Email).
			SetHashedPassword(u.HashedPassword).
			SetActivated(u.Activated).
			SetName(u.Name).
			SetVersion(u.Version).
			Save(ctx)
		if err != nil {
			return err
		}

		_, err = txClient.Token.Delete().
			Where(token.Scope(database.ScopeActivation), token.HasUsersWith(userPredicate.ID(u.ID))).
			Exec(ctx)

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			v.AddFieldError("token", "invalid or expired activation token")
			return nil, r.errorMessage(v)
		default:
			return nil, r.serverError(err)
		}
	}

	return &gigachad.ActivateUserResult{
		TokenPlainText: input.TokenPlainText,
	}, nil
}

// UpdateUserPassword is the resolver for the updateUserPassword field.
func (r *mutationResolver) UpdateUserPassword(ctx context.Context, input gigachad.ResetUserPasswordInput) (*gigachad.ResetUserPasswordResult, error) {
	v := validator.NewValidator()

	types.ValidatePasswordPlaintext(v, input.Password)
	types.ValidateTokenPlaintext(v, input.TokenPlainText)

	if v.HasErrors() {
		return nil, r.errorMessage(v)
	}

	err := r.WithTx(ctx, func(tx *ent.Tx) error {
		txClient := tx.Client()

		tokenHash := sha256.Sum256([]byte(input.TokenPlainText))

		u, err := txClient.User.Query().
			Where(
				userPredicate.HasTokensWith(token.Hash(tokenHash[:]),
					token.Scope(database.ScopePasswordReset),
					token.ExpiryGT(time.Now().Format(time.RFC3339)),
				),
			).
			Only(ctx)
		if err != nil {
			return err
		}

		user := types.User{
			Ent: u,
		}

		err = user.SetPassword(input.Password)
		if err != nil {
			return err
		}

		_, err = txClient.User.UpdateOneID(user.Ent.ID).
			SetUsername(user.Ent.Username).
			SetEmail(user.Ent.Email).
			SetHashedPassword(user.Ent.HashedPassword).
			SetActivated(user.Ent.Activated).
			SetName(user.Ent.Name).
			SetVersion(user.Ent.Version + 1).
			Save(ctx)

		if err != nil {
			return err
		}

		_, err = txClient.Token.Delete().
			Where(token.Scope(database.ScopePasswordReset), token.HasUsersWith(userPredicate.ID(user.Ent.ID))).
			Exec(ctx)

		if err != nil {
			return err
		}

		_, err = txClient.Token.Delete().
			Where(token.Scope(database.SocpeAuthentication), token.HasUsersWith(userPredicate.ID(user.Ent.ID))).
			Exec(ctx)

		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			v.AddFieldError("token", "invalid or expired password reset token")
			return nil, r.errorMessage(v)
		default:
			return nil, r.serverError(err)
		}
	}

	return &gigachad.ResetUserPasswordResult{
		Password:       input.Password,
		TokenPlainText: input.TokenPlainText,
	}, nil
}

// Viewer is the resolver for the viewer field.
func (r *queryResolver) Viewer(ctx context.Context) (*ent.User, error) {
	user := r.getUserFromCtx(ctx)
	return user.Ent, nil
}

// Mutation returns gigachad.MutationResolver implementation.
func (r *Resolver) Mutation() gigachad.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
