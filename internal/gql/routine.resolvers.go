package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	gigachad "github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

// CreateRoutine is the resolver for the createRoutine field.
func (r *mutationResolver) CreateRoutine(ctx context.Context, input ent.CreateRoutineInput) (*ent.Routine, error) {
	return r.client.Routine.Create().SetInput(input).Save(ctx)
}

// DeleteRoutine is the resolver for the deleteRoutine field.
func (r *mutationResolver) DeleteRoutine(ctx context.Context, id pksuid.ID) (*gigachad.DeletedID, error) {
	v := validator.NewValidator()

	err := r.client.Routine.DeleteOneID(id).Exec(ctx)
	if err != nil {
		switch {
		case err.Error() == types.EntRoutineNotFound:
			v.AddFieldError("routine", "a routine with this ID doesn't exist")
			return nil, r.errorMessage(v)
		default:
			return nil, r.serverError(err)
		}
	}
	return &gigachad.DeletedID{
		ID: &id,
	}, nil
}
