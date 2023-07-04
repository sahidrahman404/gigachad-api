package types

import (
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

type CreateRoutineParams struct {
	Name string `json:"name"`
}

type Routine struct {
	Ent *ent.Routine
}

func NewRoutineFromParams(p CreateRoutineParams, userID string) *Routine {
	r := &Routine{
		Ent: &ent.Routine{
			Name:   p.Name,
			UserID: userID,
		},
	}
	return r
}

func (p CreateRoutineParams) Validate(v *validator.Validator) {
	v.CheckField(p.Name != "", "name", "name must be provided")
}
