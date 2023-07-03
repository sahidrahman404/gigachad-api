package database

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type ExerciseStorer interface {
	Insert(ctx context.Context, ex *types.Exercise) error
	Update(ctx context.Context, ex *types.Exercise) error
	Delete(ctx context.Context, ex *types.Exercise) error
}

type ExerciseStore struct {
	Client *ent.Client
}

func NewExerciseStore(e *ent.Client) *ExerciseStore {
	return &ExerciseStore{
		Client: e,
	}
}

func (e *ExerciseStore) Insert(ctx context.Context, ex *types.Exercise) error {
	exercise, err := e.Client.Exercise.Create().
		SetName(ex.Ent.Name).
		SetHowTo(ex.Ent.HowTo).
		SetUserID(ex.Ent.UserID).
		SetEquipmentID(ex.Ent.EquipmentID).
		SetExerciseTypeID(ex.Ent.ExerciseTypeID).
		SetMusclesGroupID(ex.Ent.MusclesGroupID).Save(ctx)
	if err != nil {
		return err
	}
	ex.Ent.ID = exercise.ID
	return nil
}

func (e *ExerciseStore) Update(ctx context.Context, ex *types.Exercise) error {
	_, err := e.Client.Exercise.UpdateOneID(ex.Ent.ID).
		SetName(ex.Ent.Name).
		SetHowTo(ex.Ent.HowTo).
		SetUserID(ex.Ent.UserID).
		SetEquipmentID(ex.Ent.EquipmentID).
		SetExerciseTypeID(ex.Ent.ExerciseTypeID).
		SetMusclesGroupID(ex.Ent.MusclesGroupID).Save(ctx)
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
