package types

import (
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

type CreateExerciseTypeParams struct {
	Name        string   `json:"name"`
	Properties  []string `json:"properties"`
	Description string   `json:"description"`
}

type UpdateExerciseTypeParams struct {
	Name        *string   `json:"name"`
	Properties  []string `json:"properties"`
	Description *string    `json:"description"`
}

type ExerciseType struct {
	Ent *ent.ExerciseType
}

func NewExerciseTypeFromParmas(p CreateExerciseTypeParams) *ExerciseType {
	return &ExerciseType{
		Ent: &ent.ExerciseType{
			Name:        p.Name,
			Properties:  p.Properties,
			Description: p.Description,
		},
	}
}

func (p CreateExerciseTypeParams) Validate(v *validator.Validator) {
	v.CheckField(p.Name != "", "name", "please provide name")
	v.CheckField(p.Description != "", "image", "please provide description")
	v.CheckField(len(p.Properties) != 0, "properties", "please provide properties")
}

func UpdateExerciseTypeFromParams(etp *UpdateExerciseTypeParams, et *ExerciseType) {
	if etp.Name != nil{
		et.Ent.Name = *etp.Name
	}

	if etp.Description != nil{
		et.Ent.Description = *etp.Description
	}


	if etp.Properties != nil{
		et.Ent.Properties = etp.Properties
	}
}
