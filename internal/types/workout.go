package types

import "github.com/sahidrahman404/gigachad-api/ent"

type CreateWorkoutParams struct {
	Name        string  `json:"name"`
	Volume      int     `json:"volume"`
	Reps        int     `json:"reps"`
	Time        string  `json:"time"`
	Sets        int     `json:"sets"`
	Image       *string `json:"image"`
	Description string  `json:"description"`
}

type Workout struct {
	Ent *ent.Workout
}

func NewWorkoutFromParams(p CreateWorkoutParams, userID string) *Workout {
	return &Workout{
		Ent: &ent.Workout{
			Name:        p.Name,
			Volume:      p.Volume,
			Reps:        p.Reps,
			Time:        p.Time,
			Sets:        p.Sets,
			Image:       p.Image,
			Description: p.Description,
		},
	}
}
