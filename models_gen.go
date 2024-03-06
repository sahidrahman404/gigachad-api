// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gigachad

import (
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
)

type ActivateUserInput struct {
	TokenPlainText string `json:"tokenPlainText"`
}

type ActivateUserResult struct {
	TokenPlainText string `json:"tokenPlainText"`
}

type ActivationTokenInput struct {
	Email string `json:"email"`
}

type AuthenticationToken struct {
	User           *ent.User `json:"user"`
	TokenPlainText string    `json:"tokenPlainText"`
}

type CreateExerciseInput struct {
	Name            string            `json:"name"`
	Image           *schematype.Image `json:"image,omitempty"`
	HowTo           *string           `json:"howTo,omitempty"`
	UserID          *pksuid.ID        `json:"userID,omitempty"`
	MusclesGroupIDs []pksuid.ID       `json:"musclesGroupIDs,omitempty"`
	ExerciseTypeIDs []pksuid.ID       `json:"exerciseTypeIDs,omitempty"`
}

type CreateMusclesGroupInput struct {
	Name  string            `json:"name"`
	Image *schematype.Image `json:"image"`
}

type CreateRoutineExerciseInput struct {
	RestTimer  *string           `json:"restTimer,omitempty"`
	Sets       []*schematype.Set `json:"sets,omitempty"`
	ExerciseID pksuid.ID         `json:"exerciseID"`
}

type CreateRoutineReminderInput struct {
	Day    int `json:"day"`
	Second int `json:"second"`
	Minute int `json:"minute"`
	Hour   int `json:"hour"`
}

type CreateRoutineWithChildrenInput struct {
	Name             string                        `json:"name"`
	Reminder         []*CreateRoutineReminderInput `json:"reminder,omitempty"`
	RoutineExercises []*CreateRoutineExerciseInput `json:"routineExercises,omitempty"`
}

type CreateWorkoutLogInput struct {
	Sets       []*schematype.Set `json:"sets"`
	ExerciseID pksuid.ID         `json:"exerciseID"`
}

type CreateWorkoutWithChildrenInput struct {
	Volume      int                      `json:"volume"`
	Duration    string                   `json:"duration"`
	Sets        int                      `json:"sets"`
	Image       *schematype.Image        `json:"image,omitempty"`
	Description *string                  `json:"description,omitempty"`
	WorkoutLogs []*CreateWorkoutLogInput `json:"workoutLogs,omitempty"`
}

type DeleteExerciseInput struct {
	ID pksuid.ID `json:"id"`
}

type DeleteRoutineInput struct {
	ID pksuid.ID `json:"id"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPasswordInput struct {
	Email string `json:"email"`
}

type ResetUserPasswordInput struct {
	Password       string `json:"password"`
	TokenPlainText string `json:"tokenPlainText"`
}

type ResetUserPasswordResult struct {
	Password       string `json:"password"`
	TokenPlainText string `json:"tokenPlainText"`
}

type Style struct {
	Width       *string `json:"width,omitempty"`
	Height      *string `json:"height,omitempty"`
	MaxWidth    *string `json:"maxWidth,omitempty"`
	MaxHeight   *string `json:"maxHeight,omitempty"`
	AspectRatio *string `json:"aspectRatio,omitempty"`
}

type UpdateRoutineExerciseInput struct {
	ID         *pksuid.ID        `json:"id,omitempty"`
	RestTimer  *string           `json:"restTimer,omitempty"`
	Sets       []*schematype.Set `json:"sets,omitempty"`
	ExerciseID pksuid.ID         `json:"exerciseID"`
}

type UpdateRoutineSchedulesInput struct {
	ID        *string                       `json:"id,omitempty"`
	Schedules []*CreateRoutineReminderInput `json:"schedules,omitempty"`
}

type UpdateRoutineWithChildrenInput struct {
	ID               pksuid.ID                     `json:"id"`
	Name             string                        `json:"name"`
	Reminder         *UpdateRoutineSchedulesInput  `json:"reminder,omitempty"`
	RoutineExercises []*UpdateRoutineExerciseInput `json:"routineExercises,omitempty"`
}
