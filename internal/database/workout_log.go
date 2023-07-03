package database

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type WorkoutLogStorer interface {
	Insert(ctx context.Context, wl *types.WorkoutLog) error
}

type WorkoutLogStore struct {
	Client *ent.Client
}

func NewEntWorkoutLogStore(c *ent.Client) *WorkoutLogStore {
	return &WorkoutLogStore{
		Client: c,
	}
}

func (e *WorkoutLogStore) Insert(ctx context.Context, wl *types.WorkoutLog) error {
	workoutlog, err := e.Client.WorkoutLog.Create().
		SetSets(wl.Ent.Sets).
		SetExerciseID(wl.Ent.ExerciseID).
		SetWorkoutID(wl.Ent.WorkoutID).Save(ctx)
	if err != nil {
		return err
	}
	wl.Ent.ID = workoutlog.ID
	return nil
}
