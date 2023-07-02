package types

import (
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

type CreateMusclesGroupParams struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type MusclesGroup struct {
	Ent *ent.MusclesGroup
}

func NewMusclesGroupFromParams(p CreateMusclesGroupParams) *MusclesGroup {
	return &MusclesGroup{
		Ent: &ent.MusclesGroup{
			Name:  p.Name,
			Image: p.Image,
		},
	}
}

func (p CreateMusclesGroupParams) Validate(v validator.Validator) {
	v.CheckField(p.Name != "", "name", "please provide name")
	v.CheckField(p.Image != "", "image", "please provide image url")

	v.CheckField(validator.IsURL(p.Image), "image", "please provide correct image url")
}
