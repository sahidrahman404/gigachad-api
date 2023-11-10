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

type ActivationTokenInput struct {
	Email string `json:"email"`
}

type AuthenticationToken struct {
	User           *ent.User `json:"user"`
	TokenPlainText string    `json:"tokenPlainText"`
}

type CreateRoutineExerciseInput struct {
	RestTimer  *string           `json:"restTimer,omitempty"`
	Sets       []*schematype.Set `json:"sets,omitempty"`
	ExerciseID pksuid.ID         `json:"exerciseID"`
	UserID     pksuid.ID         `json:"userID"`
}

type CreateRoutineWithChildrenInput struct {
	Name            string                        `json:"name"`
	RoutineExercise []*CreateRoutineExerciseInput `json:"RoutineExercise,omitempty"`
}

type DeletedID struct {
	ID *pksuid.ID `json:"id,omitempty"`
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
