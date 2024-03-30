package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	gigachad "github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/img"
)

// CreateMusclesGroup is the resolver for the CreateMusclesGroup field.
func (r *mutationResolver) CreateMusclesGroup(ctx context.Context, input gigachad.CreateMusclesGroupInput) (*ent.MusclesGroup, error) {
	return r.client.MusclesGroup.Create().
		SetName(input.Name).
		SetImage(img.SetImageField(*input.Image, *r.awsCfg, r.imgproxy)).
		Save(ctx)
}
