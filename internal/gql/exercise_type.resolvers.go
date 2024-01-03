package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
)

// CreateExerciseType is the resolver for the createExerciseType field.
func (r *mutationResolver) CreateExerciseType(ctx context.Context, input ent.CreateExerciseTypeInput) (*ent.ExerciseType, error) {
	return r.client.ExerciseType.Create().
		SetName(input.Name).
		SetProperties(input.Properties).
		SetDescription(input.Description).
		Save(ctx)
}
