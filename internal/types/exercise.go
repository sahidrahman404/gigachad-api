package types

import (
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

type CreateExerciseParams struct {
	Name           string  `json:"name"`
	HowTo          *string `json:"howTo"`
	EquipmentID    string  `json:"equipmentID"`
	MusclesGroupID string  `json:"MusclesGroupID"`
	ExerciseTypeID string  `json:"ExerciseTypeID"`
}

type Exercise struct {
	Ent *ent.Exercise
}

func NewExerciseFromParams(p CreateExerciseParams) *Exercise {
	e := &Exercise{
		Ent: &ent.Exercise{
			Name: p.Name,
			EquipmentID: p.EquipmentID,
			MusclesGroupID: p.MusclesGroupID,
			ExerciseTypeID: p.ExerciseTypeID,
		},
	}
	if p.HowTo != nil{
		h := *p.HowTo
		e.Ent.HowTo = h 
	}
	return e
}

func (p CreateExerciseParams) ValidateExercise(v validator.Validator) {
	v.CheckField(p.Name != "", "name", "name must be provided")
	v.CheckField(p.EquipmentID != "", "equipmentID", "equipmentID must be provided")
	v.CheckField(p.MusclesGroupID != "", "musclesGroupID", "musclesGroupID must be provided")
	v.CheckField(p.ExerciseTypeID != "", "exerciseTypeID", "exerciseTypeID must be provided")

	v.CheckField(len(p.EquipmentID) != 27, "equipmentID", "please add valid equipmentID")
	v.CheckField(len(p.MusclesGroupID) != 27, "musclesGroupID", "please add valid musclesGroupID")
	v.CheckField(len(p.ExerciseTypeID) != 27, "exerciseTypeID", "please add valid exerciseTypeID")
}
