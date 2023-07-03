package database

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type MusclesGroupStorer interface {
	Insert(ctx context.Context, m *types.MusclesGroup) error
	Update(ctx context.Context, m *types.MusclesGroup) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*types.MusclesGroup, error)
}

type MusclesGroupStore struct {
	Client *ent.Client
}

func NewEntMusclesGroupStore(e *ent.Client) *MusclesGroupStore {
	return &MusclesGroupStore{
		Client: e,
	}
}

func (e MusclesGroupStore) Insert(ctx context.Context, m *types.MusclesGroup) error {
	mg, err := e.Client.MusclesGroup.Create().
		SetName(m.Ent.Name).
		SetImage(m.Ent.Image).
		Save(ctx)
	if err != nil {
		return err
	}
	m.Ent.ID = mg.ID
	return nil
}

func (e MusclesGroupStore) Update(ctx context.Context, m *types.MusclesGroup) error {
	_, err := e.Client.MusclesGroup.UpdateOneID(m.Ent.ID).
		SetName(m.Ent.Name).
		SetImage(m.Ent.Image).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e MusclesGroupStore) Delete(ctx context.Context, id string) error {
	err := e.Client.MusclesGroup.DeleteOneID(id).Exec(ctx)
	return err
}

func (e MusclesGroupStore) Get(ctx context.Context, id string) (*types.MusclesGroup, error) {
	mg, err := e.Client.MusclesGroup.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &types.MusclesGroup{
		Ent: mg,
	}, err
}
