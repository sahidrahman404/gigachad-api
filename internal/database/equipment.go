package database

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type EquipmentStorer interface {
	Insert(ctx context.Context, eq *types.Equipment) error
	Get(ctx context.Context, id string) (*types.Equipment, error)
	Update(ctx context.Context, eq *types.Equipment) error
	GetAll(ctx context.Context) ([]*types.Equipment, error)
}

type EquipmentStore struct {
	Client *ent.Client
}

func NewEntEquipmentStore(c *ent.Client) *EquipmentStore {
	return &EquipmentStore{
		Client: c,
	}
}

func (e *EquipmentStore) Insert(ctx context.Context, eq *types.Equipment) error {
	equipment, err := e.Client.Equipment.Create().SetName(eq.Ent.Name).SetImage(eq.Ent.Image).Save(ctx)
	if err != nil {
		return err
	}
	eq.Ent.ID = equipment.ID
	return nil
}

func (e *EquipmentStore) Get(ctx context.Context, id string) (*types.Equipment, error) {
	equipment, err := e.Client.Equipment.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &types.Equipment{
		Ent: equipment,
	}, err
}

func (e *EquipmentStore) Update(ctx context.Context, eq *types.Equipment) error {
	_, err := e.Client.Equipment.UpdateOneID(eq.Ent.ID).
		SetName(eq.Ent.Name).SetImage(eq.Ent.Image).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (e *EquipmentStore) GetAll(ctx context.Context) ([]*types.Equipment, error) {
	equipment, err := e.Client.Equipment.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	eq := []*types.Equipment{}
	for _, v := range equipment {
		if v != nil {
			eq = append(eq, &types.Equipment{
				Ent: v,
			})
		}
	}
	return eq, nil
}
