package types

import (
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

type CreateRoutineParams struct {
	Name             string `json:"name"`
	RestTimer        *int   `json:"restTimer"`
	RoutineExercises []RE
}

type RE struct {
	Set        []schema.Set `json:"set"`
	RoutineID  string       `json:"routineID"`
	ExerciseID string       `json:"exerciseID"`
}

type Routine struct {
	Ent *ent.Routine
}

type RoutineExercise struct {
	Ent *ent.RoutineExercise
}

func NewRoutineExerciseFromParams(p CreateRoutineParams) (*Routine, *[]RoutineExercise) {
	r := &Routine{
		Ent: &ent.Routine{
			Name: p.Name,
		},
	}
	res := &[]RoutineExercise{}
	for _, v := range p.RoutineExercises {
		re := RoutineExercise{
			Ent: &ent.RoutineExercise{
				Set:        &v.Set,
				ExerciseID: v.ExerciseID,
				RoutineID:  v.RoutineID,
			},
		}
		*res = append(*res, re)
	}
	return r, res
}

func (p CreateRoutineParams) ValidateExercise(v validator.Validator) {
	v.CheckField(p.Name != "", "name", "name must be provided")
	v.CheckField(len(p.RoutineExercises) != 0, "routineExercises", "routine exercises must be provided")
	for _, val := range p.RoutineExercises {
		v.CheckField(len(val.RoutineID) != 27, "routineID", "please add valid routineID")
		v.CheckField(len(val.ExerciseID) != 27, "exerciseID", "please add valid exerciseID")
	}
}
