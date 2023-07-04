package types

import (
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

type CreateEquipmentParams struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}


type UpdateEquipmentParams struct {
	Name  *string `json:"name"`
	Image *string `json:"image"`
}

type Equipment struct {
	Ent *ent.Equipment
}

func NewEquipmentFromParams(p CreateEquipmentParams) *Equipment {
	return &Equipment{
		Ent: &ent.Equipment{
			Name:  p.Name,
			Image: p.Image,
		},
	}
}

func (p CreateEquipmentParams) Validate(v *validator.Validator) {
	v.CheckField(p.Name != "", "name", "please provide name")
	v.CheckField(p.Image != "", "image", "please provide image url")

	v.CheckField(validator.IsURL(p.Image), "image", "please provide correct image url")
}

func UpdateEquipmentFromParams(ep *UpdateEquipmentParams, e *Equipment) {
	if ep.Name != nil {
		e.Ent.Name = *ep.Name
	}

	if ep.Image != nil {
		e.Ent.Image = *ep.Image
	}

}
