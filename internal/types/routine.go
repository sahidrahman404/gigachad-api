package types

import (
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

type CreateRoutineParams struct {
	Name             string `json:"name"`
	RoutineExercises []RE
}

type RE struct {
	RestTimer  *int            `json:"restTimer"`
	Sets       schematype.Sets `json:"set"`
	RoutineID  string          `json:"routineID"`
	ExerciseID string          `json:"exerciseID"`
}

type Routine struct {
	Ent *ent.Routine
}

type RoutineExercise struct {
	Ent *ent.RoutineExercise
}

func NewRoutineExerciseFromParams(p CreateRoutineParams, userID string) (*Routine, *[]RoutineExercise) {
	r := &Routine{
		Ent: &ent.Routine{
			Name:   p.Name,
			UserID: userID,
		},
	}
	res := &[]RoutineExercise{}
	for _, v := range p.RoutineExercises {
		re := RoutineExercise{
			Ent: &ent.RoutineExercise{
				Sets:       &v.Sets,
				ExerciseID: v.ExerciseID,
				RoutineID:  v.RoutineID,
			},
		}
		if v.RestTimer != nil {
			re.Ent.RestTimer = *v.RestTimer
		}
		*res = append(*res, re)
	}
	return r, res
}

func (p CreateRoutineParams) ValidateExercise(v *validator.Validator) {
	v.CheckField(p.Name != "", "name", "name must be provided")
	v.CheckField(len(p.RoutineExercises) != 0, "routineExercises", "routine exercises must be provided")
	for _, val := range p.RoutineExercises {
		v.CheckField(len(val.RoutineID) != 27, "routineID", "please add valid routineID")
		v.CheckField(len(val.ExerciseID) != 27, "exerciseID", "please add valid exerciseID")
	}
}
