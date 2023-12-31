// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/ent/workout"
	"github.com/sahidrahman404/gigachad-api/ent/workoutlog"
)

// WorkoutLog is the model entity for the WorkoutLog schema.
type WorkoutLog struct {
	config `json:"-"`
	// ID of the ent.
	ID pksuid.ID `json:"id,omitempty"`
	// Sets holds the value of the "sets" field.
	Sets []*schematype.Set `json:"sets,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt string `json:"created_at,omitempty"`
	// WorkoutID holds the value of the "workout_id" field.
	WorkoutID pksuid.ID `json:"workout_id,omitempty"`
	// ExerciseID holds the value of the "exercise_id" field.
	ExerciseID pksuid.ID `json:"exercise_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID pksuid.ID `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkoutLogQuery when eager-loading is set.
	Edges        WorkoutLogEdges `json:"edges"`
	selectValues sql.SelectValues
}

// WorkoutLogEdges holds the relations/edges for other nodes in the graph.
type WorkoutLogEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// Workouts holds the value of the workouts edge.
	Workouts *Workout `json:"workouts,omitempty"`
	// Exercises holds the value of the exercises edge.
	Exercises *Exercise `json:"exercises,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkoutLogEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Users == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// WorkoutsOrErr returns the Workouts value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkoutLogEdges) WorkoutsOrErr() (*Workout, error) {
	if e.loadedTypes[1] {
		if e.Workouts == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: workout.Label}
		}
		return e.Workouts, nil
	}
	return nil, &NotLoadedError{edge: "workouts"}
}

// ExercisesOrErr returns the Exercises value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkoutLogEdges) ExercisesOrErr() (*Exercise, error) {
	if e.loadedTypes[2] {
		if e.Exercises == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: exercise.Label}
		}
		return e.Exercises, nil
	}
	return nil, &NotLoadedError{edge: "exercises"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkoutLog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workoutlog.FieldSets:
			values[i] = new([]byte)
		case workoutlog.FieldID, workoutlog.FieldWorkoutID, workoutlog.FieldExerciseID, workoutlog.FieldUserID:
			values[i] = new(pksuid.ID)
		case workoutlog.FieldCreatedAt:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkoutLog fields.
func (wl *WorkoutLog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workoutlog.FieldID:
			if value, ok := values[i].(*pksuid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				wl.ID = *value
			}
		case workoutlog.FieldSets:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field sets", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &wl.Sets); err != nil {
					return fmt.Errorf("unmarshal field sets: %w", err)
				}
			}
		case workoutlog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				wl.CreatedAt = value.String
			}
		case workoutlog.FieldWorkoutID:
			if value, ok := values[i].(*pksuid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field workout_id", values[i])
			} else if value != nil {
				wl.WorkoutID = *value
			}
		case workoutlog.FieldExerciseID:
			if value, ok := values[i].(*pksuid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field exercise_id", values[i])
			} else if value != nil {
				wl.ExerciseID = *value
			}
		case workoutlog.FieldUserID:
			if value, ok := values[i].(*pksuid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				wl.UserID = *value
			}
		default:
			wl.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WorkoutLog.
// This includes values selected through modifiers, order, etc.
func (wl *WorkoutLog) Value(name string) (ent.Value, error) {
	return wl.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the WorkoutLog entity.
func (wl *WorkoutLog) QueryUsers() *UserQuery {
	return NewWorkoutLogClient(wl.config).QueryUsers(wl)
}

// QueryWorkouts queries the "workouts" edge of the WorkoutLog entity.
func (wl *WorkoutLog) QueryWorkouts() *WorkoutQuery {
	return NewWorkoutLogClient(wl.config).QueryWorkouts(wl)
}

// QueryExercises queries the "exercises" edge of the WorkoutLog entity.
func (wl *WorkoutLog) QueryExercises() *ExerciseQuery {
	return NewWorkoutLogClient(wl.config).QueryExercises(wl)
}

// Update returns a builder for updating this WorkoutLog.
// Note that you need to call WorkoutLog.Unwrap() before calling this method if this WorkoutLog
// was returned from a transaction, and the transaction was committed or rolled back.
func (wl *WorkoutLog) Update() *WorkoutLogUpdateOne {
	return NewWorkoutLogClient(wl.config).UpdateOne(wl)
}

// Unwrap unwraps the WorkoutLog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wl *WorkoutLog) Unwrap() *WorkoutLog {
	_tx, ok := wl.config.driver.(*txDriver)
	if !ok {
		panic("ent: WorkoutLog is not a transactional entity")
	}
	wl.config.driver = _tx.drv
	return wl
}

// String implements the fmt.Stringer.
func (wl *WorkoutLog) String() string {
	var builder strings.Builder
	builder.WriteString("WorkoutLog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", wl.ID))
	builder.WriteString("sets=")
	builder.WriteString(fmt.Sprintf("%v", wl.Sets))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(wl.CreatedAt)
	builder.WriteString(", ")
	builder.WriteString("workout_id=")
	builder.WriteString(fmt.Sprintf("%v", wl.WorkoutID))
	builder.WriteString(", ")
	builder.WriteString("exercise_id=")
	builder.WriteString(fmt.Sprintf("%v", wl.ExerciseID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", wl.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// WorkoutLogs is a parsable slice of WorkoutLog.
type WorkoutLogs []*WorkoutLog
