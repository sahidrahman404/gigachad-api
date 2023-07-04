package database

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type RoutineExerciseStorer interface {
	Insert(ctx context.Context, r *types.Routine) error
	Get(ctx context.Context, id string) (*types.Routine, error)
	Update(ctx context.Context, r *types.Routine) error
	Delete(ctx context.Context, id string) error
	GetAllForRoutine(ctx context.Context, routineID string) ([]*types.RoutineExercise, error)
}

type RoutineExerciseStore struct {
	Client *ent.Client
}

func NewEntRoutineExerciseStore(c *ent.Client) *RoutineExerciseStore {
	return &RoutineExerciseStore{
		Client: c,
	}
}

func (e *RoutineExerciseStore) Insert(ctx context.Context, r *types.Routine) error {
	routine, err := e.Client.Routine.Create().SetName(r.Ent.Name).SetUserID(r.Ent.UserID).Save(ctx)
	if err != nil {
		return err
	}
	r.Ent.ID = routine.ID
	return nil
}

func (e *RoutineExerciseStore) Get(ctx context.Context, id string) (*types.Routine, error) {
	r, err := e.Client.Routine.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &types.Routine{
		Ent: r,
	}, nil
}

func (e *RoutineExerciseStore) Update(ctx context.Context, r *types.Routine) error {
	_, err := e.Client.Routine.UpdateOneID(r.Ent.ID).SetName(r.Ent.Name).SetUserID(r.Ent.UserID).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *RoutineExerciseStore) Delete(ctx context.Context, id string) error {
	err := e.Client.Routine.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *RoutineExerciseStore) GetAllForRoutine(ctx context.Context, routineID string) ([]*types.RoutineExercise, error) {
	routine, err := e.Client.RoutineExercise.Query().
		Where(routineexercise.HasRoutinesWith(routine.ID(routineID))).
		WithExercises().
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
