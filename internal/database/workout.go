package database

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/workout"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type WorkoutStorer interface {
	Insert(context.Context, *types.Workout) error
	GetAllForUser(context.Context, pksuid.ID) ([]*types.Workout, error)
	GetForUser(ctx context.Context, workoutID pksuid.ID, userID pksuid.ID) (*types.Workout, error)
}

type WorkoutStore struct {
	Client *ent.Client
}

func NewEntWorkoutStore(c *ent.Client) *WorkoutStore {
	return &WorkoutStore{
		Client: c,
	}
}

func (e *WorkoutStore) Insert(ctx context.Context, w *types.Workout) error {
	workout, err := e.Client.Workout.Create().SetName(w.Ent.Name).
		// SetVolume(w.Ent.Volume).SetReps(w.Ent.Reps).
		// SetTime(w.Ent.Time).SetSets(w.Ent.Sets).
		// SetNillableImage(w.Ent.Image).SetDescription(w.Ent.Description).
		SetUserID(w.Ent.UserID).Save(ctx)
	if err != nil {
		return err
	}
	w.Ent.ID = workout.ID
	return nil
}

func (e *WorkoutStore) GetAllForUser(
	ctx context.Context,
	userID pksuid.ID,
) ([]*types.Workout, error) {
	w, err := e.Client.Workout.Query().
		Where(workout.UserID(userID)).
		WithWorkoutLogs(func(wlq *ent.WorkoutLogQuery) {
			wlq.WithExercises(func(eq *ent.ExerciseQuery) {
				eq.WithExerciseTypes()
			})
		}).
		All(ctx)
	if err != nil {
		return nil, err
	}
	workout := []*types.Workout{}
	for _, v := range w {
		workout = append(workout, &types.Workout{
			Ent: v,
		})
	}
	return workout, nil
}

func (e *WorkoutStore) GetForUser(
	ctx context.Context,
	workoutID pksuid.ID,
	userID pksuid.ID,
) (*types.Workout, error) {
	w, err := e.Client.Workout.Query().
		Where(workout.ID(workoutID), workout.UserID(userID)).
		WithWorkoutLogs(func(wlq *ent.WorkoutLogQuery) {
			wlq.WithExercises(func(eq *ent.ExerciseQuery) {
				eq.WithExerciseTypes()
			})
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &types.Workout{
		Ent: w,
	}, nil
}
