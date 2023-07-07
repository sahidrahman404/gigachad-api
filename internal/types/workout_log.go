package types

import (
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
)

type CreateWorkoutLogParams struct {
	Sets       schematype.Sets `json:"sets"`
	ExerciseID pksuid.ID       `json:"exercise_id"`
	WorkoutID  pksuid.ID       `json:"workout_id"`
}

type WorkoutLog struct {
	Ent *ent.WorkoutLog
}

func NewWorkoutLogFromParams(p CreateWorkoutLogParams, userID pksuid.ID) *WorkoutLog {
	return &WorkoutLog{
		Ent: &ent.WorkoutLog{
			Sets:       &p.Sets,
			ExerciseID: p.ExerciseID,
			WorkoutID:  p.WorkoutID,
			UserID:     userID,
		},
	}
}
