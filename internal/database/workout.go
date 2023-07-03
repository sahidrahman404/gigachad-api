package database

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type WorkoutStorer interface {
	Insert(ctx context.Context, w *types.Workout) error
	GetAll(ctx context.Context) (*[]types.Workout, error)
}

type WorkoutStore struct {
	Client *ent.Client
}

func NewWorkoutStore(c *ent.Client) *WorkoutStore {
	return &WorkoutStore{
		Client: c,
	}
}

func (e *WorkoutStore) Insert(ctx context.Context, w *types.Workout) error {
	workout, err := e.Client.Workout.Create().SetName(w.Ent.Name).
		SetVolume(w.Ent.Volume).SetReps(w.Ent.Reps).
		SetTime(w.Ent.Time).SetSets(w.Ent.Sets).
		SetNillableImage(w.Ent.Image).SetDescription(w.Ent.Description).
		SetUserID(w.Ent.UserID).Save(ctx)
	if err != nil {
		return err
	}
	w.Ent.ID = workout.ID
	return nil
}

func (e *WorkoutStore) GetAll(ctx context.Context) (*[]types.Workout, error) {
	w, err := e.Client.Workout.Query().WithWorkoutLogs(func(wlq *ent.WorkoutLogQuery) {
		wlq.WithExercises()
	}).All(ctx)
	if err != nil {
		return nil, err
	}
	var workout *[]types.Workout
	for _, v := range w {
		*workout = append(*workout, types.Workout{
			Ent: v,
		})
	}
	return workout, nil
}
