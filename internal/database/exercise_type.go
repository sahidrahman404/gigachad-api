package database

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type ExerciseTypeStorer interface {
	Insert(ctx context.Context, et *types.ExerciseType) error
	Update(ctx context.Context, et *types.ExerciseType) error
	Delete(ctx context.Context, et *types.ExerciseType) error
	Get(ctx context.Context, id string) (*types.ExerciseType, error)
	GetAll(ctx context.Context) ([]*types.ExerciseType, error)
}

type EntExerciseTypeStore struct {
	Client *ent.Client
}

func NewEntExerciseTypeStore(c *ent.Client) *EntExerciseTypeStore {
	return &EntExerciseTypeStore{
		Client: c,
	}
}

func (e *EntExerciseTypeStore) Insert(ctx context.Context, et *types.ExerciseType) error {
	exerciseType, err := e.Client.ExerciseType.Create().
		SetName(et.Ent.Name).
		SetProperties(et.Ent.Properties).
		SetDescription(et.Ent.Description).
		Save(ctx)
	if err != nil {
		return err
	}
	et.Ent.ID = exerciseType.ID
	return nil
}

func (e *EntExerciseTypeStore) Update(ctx context.Context, et *types.ExerciseType) error {
	_, err := e.Client.ExerciseType.UpdateOneID(et.Ent.ID).
		SetName(et.Ent.Name).
		SetProperties(et.Ent.Properties).
		SetDescription(et.Ent.Description).
		Save(ctx)
	return err
}

func (e *EntExerciseTypeStore) Get(ctx context.Context, id string) (*types.ExerciseType, error) {
	et, err := e.Client.ExerciseType.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &types.ExerciseType{
		Ent: et,
	}, nil
}

func (e *EntExerciseTypeStore) Delete(ctx context.Context, et *types.ExerciseType) error {
	err := e.Client.ExerciseType.DeleteOneID(et.Ent.ID).Exec(ctx)
	return err
}

func (e *EntExerciseTypeStore) GetAll(ctx context.Context) ([]*types.ExerciseType, error) {
	et, err := e.Client.ExerciseType.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	exerciseType := []*types.ExerciseType{}
	for _, v := range et {
		exerciseType = append(exerciseType, &types.ExerciseType{
			Ent: v,
		})
	}
	return exerciseType, nil
}
