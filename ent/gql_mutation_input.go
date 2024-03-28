// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

// CreateExerciseTypeInput represents a mutation input for creating exercisetypes.
type CreateExerciseTypeInput struct {
	Name        string
	Properties  []string
	Description string
	ExerciseIDs []pksuid.ID
}

// Mutate applies the CreateExerciseTypeInput on the ExerciseTypeMutation builder.
func (i *CreateExerciseTypeInput) Mutate(m *ExerciseTypeMutation) {
	m.SetName(i.Name)
	if v := i.Properties; v != nil {
		m.SetProperties(v)
	}
	m.SetDescription(i.Description)
	if v := i.ExerciseIDs; len(v) > 0 {
		m.AddExerciseIDs(v...)
	}
}

// SetInput applies the change-set in the CreateExerciseTypeInput on the ExerciseTypeCreate builder.
func (c *ExerciseTypeCreate) SetInput(i CreateExerciseTypeInput) *ExerciseTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateExerciseTypeInput represents a mutation input for updating exercisetypes.
type UpdateExerciseTypeInput struct {
	Name              *string
	Properties        []string
	AppendProperties  []string
	Description       *string
	ClearExercises    bool
	AddExerciseIDs    []pksuid.ID
	RemoveExerciseIDs []pksuid.ID
}

// Mutate applies the UpdateExerciseTypeInput on the ExerciseTypeMutation builder.
func (i *UpdateExerciseTypeInput) Mutate(m *ExerciseTypeMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Properties; v != nil {
		m.SetProperties(v)
	}
	if i.AppendProperties != nil {
		m.AppendProperties(i.Properties)
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	if i.ClearExercises {
		m.ClearExercises()
	}
	if v := i.AddExerciseIDs; len(v) > 0 {
		m.AddExerciseIDs(v...)
	}
	if v := i.RemoveExerciseIDs; len(v) > 0 {
		m.RemoveExerciseIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateExerciseTypeInput on the ExerciseTypeUpdate builder.
func (c *ExerciseTypeUpdate) SetInput(i UpdateExerciseTypeInput) *ExerciseTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateExerciseTypeInput on the ExerciseTypeUpdateOne builder.
func (c *ExerciseTypeUpdateOne) SetInput(i UpdateExerciseTypeInput) *ExerciseTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateRoutineInput represents a mutation input for creating routines.
type CreateRoutineInput struct {
	Name        string
	ReminderID  *string
	ExerciseIDs []pksuid.ID
	UsersID     pksuid.ID
}

// Mutate applies the CreateRoutineInput on the RoutineMutation builder.
func (i *CreateRoutineInput) Mutate(m *RoutineMutation) {
	m.SetName(i.Name)
	if v := i.ReminderID; v != nil {
		m.SetReminderID(*v)
	}
	if v := i.ExerciseIDs; len(v) > 0 {
		m.AddExerciseIDs(v...)
	}
	m.SetUsersID(i.UsersID)
}

// SetInput applies the change-set in the CreateRoutineInput on the RoutineCreate builder.
func (c *RoutineCreate) SetInput(i CreateRoutineInput) *RoutineCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateRoutineInput represents a mutation input for updating routines.
type UpdateRoutineInput struct {
	Name              *string
	ClearReminderID   bool
	ReminderID        *string
	ClearExercises    bool
	AddExerciseIDs    []pksuid.ID
	RemoveExerciseIDs []pksuid.ID
	UsersID           *pksuid.ID
}

// Mutate applies the UpdateRoutineInput on the RoutineMutation builder.
func (i *UpdateRoutineInput) Mutate(m *RoutineMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearReminderID {
		m.ClearReminderID()
	}
	if v := i.ReminderID; v != nil {
		m.SetReminderID(*v)
	}
	if i.ClearExercises {
		m.ClearExercises()
	}
	if v := i.AddExerciseIDs; len(v) > 0 {
		m.AddExerciseIDs(v...)
	}
	if v := i.RemoveExerciseIDs; len(v) > 0 {
		m.RemoveExerciseIDs(v...)
	}
	if v := i.UsersID; v != nil {
		m.SetUsersID(*v)
	}
}

// SetInput applies the change-set in the UpdateRoutineInput on the RoutineUpdate builder.
func (c *RoutineUpdate) SetInput(i UpdateRoutineInput) *RoutineUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateRoutineInput on the RoutineUpdateOne builder.
func (c *RoutineUpdateOne) SetInput(i UpdateRoutineInput) *RoutineUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Email          string
	Username       string
	HashedPassword string
	Name           string
	TokenIDs       []pksuid.ID
	ExerciseIDs    []pksuid.ID
	RoutineIDs     []pksuid.ID
	WorkoutIDs     []pksuid.ID
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetEmail(i.Email)
	m.SetUsername(i.Username)
	m.SetHashedPassword(i.HashedPassword)
	m.SetName(i.Name)
	if v := i.TokenIDs; len(v) > 0 {
		m.AddTokenIDs(v...)
	}
	if v := i.ExerciseIDs; len(v) > 0 {
		m.AddExerciseIDs(v...)
	}
	if v := i.RoutineIDs; len(v) > 0 {
		m.AddRoutineIDs(v...)
	}
	if v := i.WorkoutIDs; len(v) > 0 {
		m.AddWorkoutIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Email             *string
	Username          *string
	HashedPassword    *string
	Name              *string
	ClearTokens       bool
	AddTokenIDs       []pksuid.ID
	RemoveTokenIDs    []pksuid.ID
	ClearExercises    bool
	AddExerciseIDs    []pksuid.ID
	RemoveExerciseIDs []pksuid.ID
	ClearRoutines     bool
	AddRoutineIDs     []pksuid.ID
	RemoveRoutineIDs  []pksuid.ID
	ClearWorkouts     bool
	AddWorkoutIDs     []pksuid.ID
	RemoveWorkoutIDs  []pksuid.ID
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.Username; v != nil {
		m.SetUsername(*v)
	}
	if v := i.HashedPassword; v != nil {
		m.SetHashedPassword(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearTokens {
		m.ClearTokens()
	}
	if v := i.AddTokenIDs; len(v) > 0 {
		m.AddTokenIDs(v...)
	}
	if v := i.RemoveTokenIDs; len(v) > 0 {
		m.RemoveTokenIDs(v...)
	}
	if i.ClearExercises {
		m.ClearExercises()
	}
	if v := i.AddExerciseIDs; len(v) > 0 {
		m.AddExerciseIDs(v...)
	}
	if v := i.RemoveExerciseIDs; len(v) > 0 {
		m.RemoveExerciseIDs(v...)
	}
	if i.ClearRoutines {
		m.ClearRoutines()
	}
	if v := i.AddRoutineIDs; len(v) > 0 {
		m.AddRoutineIDs(v...)
	}
	if v := i.RemoveRoutineIDs; len(v) > 0 {
		m.RemoveRoutineIDs(v...)
	}
	if i.ClearWorkouts {
		m.ClearWorkouts()
	}
	if v := i.AddWorkoutIDs; len(v) > 0 {
		m.AddWorkoutIDs(v...)
	}
	if v := i.RemoveWorkoutIDs; len(v) > 0 {
		m.RemoveWorkoutIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
