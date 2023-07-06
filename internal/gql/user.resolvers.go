package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"time"

	gigachad "github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/database"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
	"github.com/vektah/gqlparser/v2/gqlerror"
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
		errList := gqlerror.List{}
		for _, v := range v.FieldErrors {
			errList = append(errList, gqlerror.Errorf(v))
		}
		return nil, errList
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

// Mutation returns gigachad.MutationResolver implementation.
func (r *Resolver) Mutation() gigachad.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
