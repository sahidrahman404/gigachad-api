package database

import "github.com/sahidrahman404/gigachad-api/ent"

type Storage struct {
	Users           UserStorer
	Tokens          TokenStorer
	Equipment       EquipmentStorer
	MusclesGroup    MusclesGroupStorer
	ExerciseType    ExerciseTypeStorer
	Exercise        ExerciseStorer
	Routine         RoutineStorer
	RoutineExercise RoutineExerciseStorer
	Workout         WorkoutStorer
	WorkoutLog      WorkoutLogStorer
}

func NewStorage(c *ent.Client) *Storage {
	return &Storage{
		Users:           NewEntUserStore(c),
		Tokens:          NewEntTokenStore(c),
		Equipment:       NewEntEquipmentStore(c),
		MusclesGroup:    NewEntMusclesGroupStore(c),
		ExerciseType:    NewEntExerciseTypeStore(c),
		Exercise:        NewEntExerciseStore(c),
		Routine:         NewEntRoutineStore(c),
		RoutineExercise: NewEntRoutineExerciseStore(c),
		Workout:         NewEntWorkoutStore(c),
		WorkoutLog:      NewEntWorkoutLogStore(c),
	}
}
