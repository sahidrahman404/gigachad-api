package types

import (
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

type CreateExerciseParams struct {
	Name           string    `json:"name"`
	HowTo          *string   `json:"howTo"`
	Image          *string   `json:"image"`
	EquipmentID    pksuid.ID `json:"equipmentID"`
	MusclesGroupID pksuid.ID `json:"MusclesGroupID"`
	ExerciseTypeID pksuid.ID `json:"ExerciseTypeID"`
}

type UpdateExerciseParams struct {
	Name           *string    `json:"name"`
	HowTo          *string    `json:"howTo"`
	Image          *string    `json:"image"`
	EquipmentID    *pksuid.ID `json:"equipmentID"`
	MusclesGroupID *pksuid.ID `json:"MusclesGroupID"`
	ExerciseTypeID *pksuid.ID `json:"ExerciseTypeID"`
}

type Exercise struct {
	Ent *ent.Exercise
}

func NewExerciseFromParams(p CreateExerciseParams, userID *pksuid.ID) *Exercise {
	e := &Exercise{
		Ent: &ent.Exercise{
			Name:           p.Name,
			Image:          p.Image,
			EquipmentID:    p.EquipmentID,
			MusclesGroupID: p.MusclesGroupID,
			ExerciseTypeID: p.ExerciseTypeID,
			UserID:         userID,
		},
	}
	if p.HowTo != nil {
		e.Ent.HowTo = p.HowTo
	}
	return e
}

func (p CreateExerciseParams) Validate(v *validator.Validator) {
	v.CheckField(p.Name != "", "name", "name must be provided")
	v.CheckField(p.EquipmentID != "", "equipmentID", "equipmentID must be provided")
	v.CheckField(p.MusclesGroupID != "", "musclesGroupID", "musclesGroupID must be provided")
	v.CheckField(p.ExerciseTypeID != "", "exerciseTypeID", "exerciseTypeID must be provided")

	v.CheckField(len(p.EquipmentID) == 27, "equipmentID", "please add valid equipmentID")
	v.CheckField(len(p.MusclesGroupID) == 27, "musclesGroupID", "please add valid musclesGroupID")
	v.CheckField(len(p.ExerciseTypeID) == 27, "exerciseTypeID", "please add valid exerciseTypeID")
}

func UpdateExerciseFromParams(e *UpdateExerciseParams, ex *Exercise) {
	if e.Name != nil {
		ex.Ent.Name = *e.Name
	}

	if e.HowTo != nil {
		ex.Ent.HowTo = e.HowTo
	}

	if e.Image != nil {
		ex.Ent.Image = e.Image
	}

	if e.EquipmentID != nil {
		ex.Ent.EquipmentID = *e.EquipmentID
	}

	if e.MusclesGroupID != nil {
		ex.Ent.MusclesGroupID = *e.MusclesGroupID
	}

	if e.ExerciseTypeID != nil {
		ex.Ent.ExerciseTypeID = *e.ExerciseTypeID
	}
}
