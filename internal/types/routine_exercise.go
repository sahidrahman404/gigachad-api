package types

import (
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

type CreateRoutineExerciseParams struct {
	RoutineID        string `json:"routineID"`
	RoutineExercises []RE   `json:"routineExercises"`
}

type RE struct {
	RestTimer  *int            `json:"restTimer"`
	Sets       schematype.Sets `json:"set"`
	ExerciseID string          `json:"exerciseID"`
}

type RoutineExercise struct {
	Ent *ent.RoutineExercise
}

func NewRoutineExerciseFromParams(p CreateRoutineExerciseParams, userID string) []*RoutineExercise {
	res := []*RoutineExercise{}
	for _, v := range p.RoutineExercises {
		re := RoutineExercise{
			Ent: &ent.RoutineExercise{
				Sets:       &v.Sets,
				ExerciseID: v.ExerciseID,
				RoutineID:  p.RoutineID,
				UserID:     userID,
			},
		}
		if v.RestTimer != nil {
			re.Ent.RestTimer = *v.RestTimer
		}
		res = append(res, &re)
	}
	return res
}

func (p CreateRoutineExerciseParams) Validate(v *validator.Validator) {
	v.CheckField(len(p.RoutineID) == 27, "routineID", "please add valid routineID")
	v.CheckField(
		len(p.RoutineExercises) != 0,
		"routineExercises",
		"routine exercises must be provided",
	)
	for _, val := range p.RoutineExercises {
		v.CheckField(len(val.ExerciseID) == 27, "exerciseID", "please add valid exerciseID")
	}
}
