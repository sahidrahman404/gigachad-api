package database

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type ExerciseStorer interface {
	Insert(ctx context.Context, ex *types.Exercise) error
	Update(ctx context.Context, ex *types.Exercise) error
	Delete(ctx context.Context, ex *types.Exercise) error
	Get(ctx context.Context, id pksuid.ID) (*types.Exercise, error)
	GetAll(ctx context.Context) ([]*types.Exercise, error)
}

type ExerciseStore struct {
	Client *ent.Client
}

func NewEntExerciseStore(e *ent.Client) *ExerciseStore {
	return &ExerciseStore{
		Client: e,
	}
}

func (e *ExerciseStore) Insert(ctx context.Context, ex *types.Exercise) error {
	exercise, err := e.Client.Exercise.Create().
		SetName(ex.Ent.Name).
		SetNillableHowTo(ex.Ent.HowTo).
		SetNillableUserID(ex.Ent.UserID).
		SetNillableImage(ex.Ent.Image).
		// SetEquipmentID(ex.Ent.EquipmentID).
		// SetExerciseTypeID(ex.Ent.ExerciseTypeID).
		// SetMusclesGroupID(ex.Ent.MusclesGroupID).
		Save(ctx)
	if err != nil {
		return err
	}
	ex.Ent.ID = exercise.ID
	return nil
}

func (e *ExerciseStore) Update(ctx context.Context, ex *types.Exercise) error {
	_, err := e.Client.Exercise.UpdateOneID(ex.Ent.ID).
		SetName(ex.Ent.Name).
		SetNillableHowTo(ex.Ent.HowTo).
		SetNillableUserID(ex.Ent.UserID).
		SetNillableImage(ex.Ent.Image).
		// SetEquipmentID(ex.Ent.EquipmentID).
		// SetExerciseTypeID(ex.Ent.ExerciseTypeID).
		// SetMusclesGroupID(ex.Ent.MusclesGroupID).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *ExerciseStore) Delete(ctx context.Context, ex *types.Exercise) error {
	err := e.Client.Exercise.DeleteOneID(ex.Ent.ID).
		Exec(ctx)
	return err
}

func (e *ExerciseStore) Get(ctx context.Context, id pksuid.ID) (*types.Exercise, error) {
	exercise, err := e.Client.Exercise.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &types.Exercise{
		Ent: exercise,
	}, nil
}

func (e *ExerciseStore) GetAll(ctx context.Context) ([]*types.Exercise, error) {
	exercises, err := e.Client.Exercise.Query().Where(exercise.UserIDIsNil()).All(ctx)
	if err != nil {
		return nil, err
	}
	exs := []*types.Exercise{}
	for _, v := range exercises {
		exs = append(exs, &types.Exercise{
			Ent: v,
		})
	}
	return exs, nil
}
