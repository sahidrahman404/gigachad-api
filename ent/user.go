// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sahidrahman404/gigachad-api/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// HashedPassword holds the value of the "hashed_password" field.
	HashedPassword string `json:"hashed_password,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt string `json:"created_at,omitempty"`
	// Activated holds the value of the "activated" field.
	Activated int `json:"activated,omitempty"`
	// Version holds the value of the "version" field.
	Version int `json:"version,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Tokens holds the value of the tokens edge.
	Tokens []*Token `json:"tokens,omitempty"`
	// Exercises holds the value of the exercises edge.
	Exercises []*Exercise `json:"exercises,omitempty"`
	// Routines holds the value of the routines edge.
	Routines []*Routine `json:"routines,omitempty"`
	// Workouts holds the value of the workouts edge.
	Workouts []*Workout `json:"workouts,omitempty"`
	// WorkoutLogs holds the value of the workout_logs edge.
	WorkoutLogs []*WorkoutLog `json:"workout_logs,omitempty"`
	// RoutineExercises holds the value of the routine_exercises edge.
	RoutineExercises []*RoutineExercise `json:"routine_exercises,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// TokensOrErr returns the Tokens value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TokensOrErr() ([]*Token, error) {
	if e.loadedTypes[0] {
		return e.Tokens, nil
	}
	return nil, &NotLoadedError{edge: "tokens"}
}

// ExercisesOrErr returns the Exercises value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ExercisesOrErr() ([]*Exercise, error) {
	if e.loadedTypes[1] {
		return e.Exercises, nil
	}
	return nil, &NotLoadedError{edge: "exercises"}
}

// RoutinesOrErr returns the Routines value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RoutinesOrErr() ([]*Routine, error) {
	if e.loadedTypes[2] {
		return e.Routines, nil
	}
	return nil, &NotLoadedError{edge: "routines"}
}

// WorkoutsOrErr returns the Workouts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) WorkoutsOrErr() ([]*Workout, error) {
	if e.loadedTypes[3] {
		return e.Workouts, nil
	}
	return nil, &NotLoadedError{edge: "workouts"}
}

// WorkoutLogsOrErr returns the WorkoutLogs value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) WorkoutLogsOrErr() ([]*WorkoutLog, error) {
	if e.loadedTypes[4] {
		return e.WorkoutLogs, nil
	}
	return nil, &NotLoadedError{edge: "workout_logs"}
}

// RoutineExercisesOrErr returns the RoutineExercises value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RoutineExercisesOrErr() ([]*RoutineExercise, error) {
	if e.loadedTypes[5] {
		return e.RoutineExercises, nil
	}
	return nil, &NotLoadedError{edge: "routine_exercises"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldActivated, user.FieldVersion:
			values[i] = new(sql.NullInt64)
		case user.FieldID, user.FieldEmail, user.FieldUsername, user.FieldHashedPassword, user.FieldName, user.FieldCreatedAt:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldHashedPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hashed_password", values[i])
			} else if value.Valid {
				u.HashedPassword = value.String
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.String
			}
		case user.FieldActivated:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field activated", values[i])
			} else if value.Valid {
				u.Activated = int(value.Int64)
			}
		case user.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				u.Version = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryTokens queries the "tokens" edge of the User entity.
func (u *User) QueryTokens() *TokenQuery {
	return NewUserClient(u.config).QueryTokens(u)
}

// QueryExercises queries the "exercises" edge of the User entity.
func (u *User) QueryExercises() *ExerciseQuery {
	return NewUserClient(u.config).QueryExercises(u)
}

// QueryRoutines queries the "routines" edge of the User entity.
func (u *User) QueryRoutines() *RoutineQuery {
	return NewUserClient(u.config).QueryRoutines(u)
}

// QueryWorkouts queries the "workouts" edge of the User entity.
func (u *User) QueryWorkouts() *WorkoutQuery {
	return NewUserClient(u.config).QueryWorkouts(u)
}

// QueryWorkoutLogs queries the "workout_logs" edge of the User entity.
func (u *User) QueryWorkoutLogs() *WorkoutLogQuery {
	return NewUserClient(u.config).QueryWorkoutLogs(u)
}

// QueryRoutineExercises queries the "routine_exercises" edge of the User entity.
func (u *User) QueryRoutineExercises() *RoutineExerciseQuery {
	return NewUserClient(u.config).QueryRoutineExercises(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("hashed_password=")
	builder.WriteString(u.HashedPassword)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt)
	builder.WriteString(", ")
	builder.WriteString("activated=")
	builder.WriteString(fmt.Sprintf("%v", u.Activated))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", u.Version))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
