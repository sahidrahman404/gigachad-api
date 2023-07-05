package database

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type RoutineExerciseStorer interface {
	InsertMany(context.Context, []*types.RoutineExercise, string) error
	Update(context.Context, *types.RoutineExercise, string) error
	Delete(context.Context, string) error
	GetAllForRoutine(context.Context, string) ([]*types.RoutineExercise, error)
}

type RoutineExerciseStore struct {
	Client *ent.Client
}

func NewEntRoutineExerciseStore(c *ent.Client) *RoutineExerciseStore {
	return &RoutineExerciseStore{
		Client: c,
	}
}

func (e *RoutineExerciseStore) InsertMany(
	ctx context.Context,
	re []*types.RoutineExercise,
	userID string,
) error {
	bulk := make([]*ent.RoutineExerciseCreate, len(re))
	for i, v := range re {
		bulk[i] = e.Client.RoutineExercise.Create().
			SetNillableRestTimer(&v.Ent.RestTimer).
			SetSets(v.Ent.Sets).
			SetRoutineID(v.Ent.RoutineID).
			SetExerciseID(v.Ent.ExerciseID).
			SetUserID(userID)
	}
	routine, err := e.Client.RoutineExercise.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return err
	}
	for i, v := range routine {
		re[i].Ent.ID = v.ID
	}
	return nil
}

func (e *RoutineExerciseStore) Update(
	ctx context.Context,
	re *types.RoutineExercise,
	routineExerciseID string,
) error {
	_, err := e.Client.RoutineExercise.UpdateOneID(routineExerciseID).
		SetNillableRestTimer(&re.Ent.RestTimer).
		SetSets(re.Ent.Sets).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *RoutineExerciseStore) Delete(
	ctx context.Context,
	routineExerciseID string,
) error {
	err := e.Client.RoutineExercise.DeleteOneID(routineExerciseID).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *RoutineExerciseStore) GetAllForRoutine(
	ctx context.Context,
	routineID string,
) ([]*types.RoutineExercise, error) {
	routine, err := e.Client.RoutineExercise.Query().
		Where(routineexercise.HasRoutinesWith(routine.ID(routineID))).
		WithExercises(func(eq *ent.ExerciseQuery) {
			eq.WithExerciseTypes()
		}).
		WithRoutines().
		All(ctx)
	if err != nil {
		return nil, err
	}

	r := []*types.RoutineExercise{}
	for _, v := range routine {
		r = append(r, &types.RoutineExercise{
			Ent: v,
		})

	}
	return r, nil
}
