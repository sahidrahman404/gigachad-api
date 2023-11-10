package types

import (
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

type UpdateRoutineExerciseParams struct {
	RestTimer *string `json:"restTimer"`
	// Sets      schematype.Sets `json:"set"`
}

type CreateRoutineExerciseParams struct {
	RoutineID        pksuid.ID `json:"routineID"`
	RoutineExercises []RE      `json:"routineExercises"`
}

type RE struct {
	RestTimer *string `json:"restTimer"`
	// Sets       schematype.Sets `json:"set"`
	ExerciseID pksuid.ID `json:"exerciseID"`
}

type RoutineExercise struct {
	Ent *ent.RoutineExercise
}

func UpdateRoutineExerciseFromParams(p UpdateRoutineExerciseParams) *RoutineExercise {
	return &RoutineExercise{
		Ent: &ent.RoutineExercise{
			// RestTimer: *p.RestTimer,
			// Sets:      &p.Sets,
		},
	}
}

func NewRoutineExerciseFromParams(
	p CreateRoutineExerciseParams,
	userID pksuid.ID,
) []*RoutineExercise {
	res := []*RoutineExercise{}
	for _, v := range p.RoutineExercises {
		re := RoutineExercise{
			Ent: &ent.RoutineExercise{
				// Sets:       &v.Sets,
				ExerciseID: v.ExerciseID,
				RoutineID:  p.RoutineID,
				UserID:     userID,
			},
		}
		// if v.RestTimer != nil {
		// 	re.Ent.RestTimer = *v.RestTimer
		// }
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
