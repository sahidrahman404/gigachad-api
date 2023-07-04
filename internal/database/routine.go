package database

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type RoutineStorer interface {
	Insert(ctx context.Context, r *types.Routine) error
	Get(ctx context.Context, id string) (*types.Routine, error)
	Update(ctx context.Context, r *types.Routine) error
	Delete(ctx context.Context, id string) error
	GetAllForUser(ctx context.Context, userID string) ([]*types.Routine, error)
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

func (e *RoutineStore) Get(ctx context.Context, id string) (*types.Routine, error) {
	r, err := e.Client.Routine.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &types.Routine{
		Ent: r,
	}, nil
}

func (e *RoutineStore) Update(ctx context.Context, r *types.Routine) error {
	_, err := e.Client.Routine.UpdateOneID(r.Ent.ID).SetName(r.Ent.Name).SetUserID(r.Ent.UserID).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *RoutineStore) Delete(ctx context.Context, id string) error {
	err := e.Client.Routine.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *RoutineStore) GetAllForUser(ctx context.Context, userID string) ([]*types.Routine, error) {
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
