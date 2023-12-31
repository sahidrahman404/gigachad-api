package database

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type RoutineStorer interface {
	Insert(context.Context, *types.Routine) error
	GetForUser(context.Context, pksuid.ID, pksuid.ID) (*types.Routine, error)
	Update(context.Context, *types.Routine, pksuid.ID) error
	Delete(context.Context, pksuid.ID) error
	GetAllForUser(context.Context, pksuid.ID) ([]*types.Routine, error)
}

type RoutineStore struct {
	Client *ent.Client
}

func NewEntRoutineStore(e *ent.Client) *RoutineStore {
	return &RoutineStore{
		Client: e,
	}
}

func (e *RoutineStore) Insert(ctx context.Context, r *types.Routine) error {
	routine, err := e.Client.Routine.Create().SetName(r.Ent.Name).SetUserID(r.Ent.UserID).Save(ctx)
	if err != nil {
		return err
	}
	r.Ent.ID = routine.ID
	return nil
}

func (e *RoutineStore) GetForUser(
	ctx context.Context,
	routineID pksuid.ID,
	userID pksuid.ID,
) (*types.Routine, error) {
	r, err := e.Client.Routine.Query().
		Where(routine.ID(routineID), routine.UserID(userID)).
		WithRoutineExercises(func(req *ent.RoutineExerciseQuery) {
			req.WithExercises(func(eq *ent.ExerciseQuery) {
				eq.WithExerciseTypes()
			})
		}).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &types.Routine{
		Ent: r,
	}, nil
}

func (e *RoutineStore) Update(ctx context.Context, r *types.Routine, routineID pksuid.ID) error {
	_, err := e.Client.Routine.UpdateOneID(routineID).
		SetName(r.Ent.Name).
		SetUserID(r.Ent.UserID).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *RoutineStore) Delete(ctx context.Context, routineID pksuid.ID) error {
	err := e.Client.Routine.DeleteOneID(routineID).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *RoutineStore) GetAllForUser(
	ctx context.Context,
	userID pksuid.ID,
) ([]*types.Routine, error) {
	routine, err := e.Client.Routine.Query().
		Where(routine.HasUsersWith(user.ID(userID))).
		All(ctx)
	if err != nil {
		return nil, err
	}

	r := []*types.Routine{}
	for _, v := range routine {
		r = append(r, &types.Routine{
			Ent: v,
		})

	}
	return r, nil
}
