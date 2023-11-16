package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"entgo.io/contrib/entgql"
	gigachad "github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

// CreateExercise is the resolver for the createExercise field.
func (r *mutationResolver) CreateExercise(ctx context.Context, input gigachad.CreateExerciseInput) (*ent.ExerciseConnection, error) {
	uCtx := r.getUserFromCtx(ctx)

	if uCtx.Ent != nil {
		ex, err := r.client.Exercise.Create().
			SetName(input.Name).
			SetImage(types.SetImageField(*input.Image, r.imgproxy)).
			SetUserID(uCtx.Ent.ID).
			Save(ctx)
		if err != nil {
			return nil, r.serverError(err)
		}

		return &ent.ExerciseConnection{
			Edges: []*ent.ExerciseEdge{
				{
					Node: ex,
					Cursor: entgql.Cursor[pksuid.ID]{
						ID: ex.ID,
					},
				},
			},
			PageInfo: entgql.PageInfo[pksuid.ID]{
				HasNextPage:     false,
				HasPreviousPage: false,
				StartCursor: &entgql.Cursor[pksuid.ID]{
					ID: ex.ID,
				},
				EndCursor: &entgql.Cursor[pksuid.ID]{
					ID: ex.ID,
				},
			},
			TotalCount: 1,
		}, nil
	}

	ex, err := r.client.Exercise.Create().
		SetName(input.Name).
		SetImage(types.SetImageField(*input.Image, r.imgproxy)).
		Save(ctx)
	if err != nil {
		return nil, r.serverError(err)
	}

	return &ent.ExerciseConnection{
		Edges: []*ent.ExerciseEdge{
			{
				Node: ex,
				Cursor: entgql.Cursor[pksuid.ID]{
					ID: ex.ID,
				},
			},
		},
		PageInfo: entgql.PageInfo[pksuid.ID]{
			HasNextPage:     false,
			HasPreviousPage: false,
			StartCursor: &entgql.Cursor[pksuid.ID]{
				ID: ex.ID,
			},
			EndCursor: &entgql.Cursor[pksuid.ID]{
				ID: ex.ID,
			},
		},
		TotalCount: 1,
	}, nil
}
