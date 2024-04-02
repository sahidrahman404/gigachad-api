// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
	"github.com/sahidrahman404/gigachad-api/ent/user"
)

// RoutineExercise is the model entity for the RoutineExercise schema.
type RoutineExercise struct {
	config `json:"-"`
	// ID of the ent.
	ID pksuid.ID `json:"id,omitempty"`
	// RestTime holds the value of the "rest_time" field.
	RestTime *string `json:"rest_time,omitempty"`
	// Sets holds the value of the "sets" field.
	Sets []*schematype.Set `json:"sets,omitempty"`
	// RoutineID holds the value of the "routine_id" field.
	RoutineID pksuid.ID `json:"routine_id,omitempty"`
	// ExerciseID holds the value of the "exercise_id" field.
	ExerciseID pksuid.ID `json:"exercise_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID pksuid.ID `json:"user_id,omitempty"`
	// Order holds the value of the "order" field.
	Order int `json:"order,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoutineExerciseQuery when eager-loading is set.
	Edges        RoutineExerciseEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoutineExerciseEdges holds the relations/edges for other nodes in the graph.
type RoutineExerciseEdges struct {
	// Routines holds the value of the routines edge.
	Routines *Routine `json:"routines,omitempty"`
	// Exercises holds the value of the exercises edge.
	Exercises *Exercise `json:"exercises,omitempty"`
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int
}

// RoutinesOrErr returns the Routines value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoutineExerciseEdges) RoutinesOrErr() (*Routine, error) {
	if e.Routines != nil {
		return e.Routines, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: routine.Label}
	}
	return nil, &NotLoadedError{edge: "routines"}
}

// ExercisesOrErr returns the Exercises value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoutineExerciseEdges) ExercisesOrErr() (*Exercise, error) {
	if e.Exercises != nil {
		return e.Exercises, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: exercise.Label}
	}
	return nil, &NotLoadedError{edge: "exercises"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoutineExerciseEdges) UsersOrErr() (*User, error) {
	if e.Users != nil {
		return e.Users, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoutineExercise) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case routineexercise.FieldSets:
			values[i] = new([]byte)
		case routineexercise.FieldID, routineexercise.FieldRoutineID, routineexercise.FieldExerciseID, routineexercise.FieldUserID:
			values[i] = new(pksuid.ID)
		case routineexercise.FieldOrder:
			values[i] = new(sql.NullInt64)
		case routineexercise.FieldRestTime:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoutineExercise fields.
func (re *RoutineExercise) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case routineexercise.FieldID:
			if value, ok := values[i].(*pksuid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				re.ID = *value
			}
		case routineexercise.FieldRestTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field rest_time", values[i])
			} else if value.Valid {
				re.RestTime = new(string)
				*re.RestTime = value.String
			}
		case routineexercise.FieldSets:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field sets", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &re.Sets); err != nil {
					return fmt.Errorf("unmarshal field sets: %w", err)
				}
			}
		case routineexercise.FieldRoutineID:
			if value, ok := values[i].(*pksuid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field routine_id", values[i])
			} else if value != nil {
				re.RoutineID = *value
			}
		case routineexercise.FieldExerciseID:
			if value, ok := values[i].(*pksuid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field exercise_id", values[i])
			} else if value != nil {
				re.ExerciseID = *value
			}
		case routineexercise.FieldUserID:
			if value, ok := values[i].(*pksuid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				re.UserID = *value
			}
		case routineexercise.FieldOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order", values[i])
			} else if value.Valid {
				re.Order = int(value.Int64)
			}
		default:
			re.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RoutineExercise.
// This includes values selected through modifiers, order, etc.
func (re *RoutineExercise) Value(name string) (ent.Value, error) {
	return re.selectValues.Get(name)
}

// QueryRoutines queries the "routines" edge of the RoutineExercise entity.
func (re *RoutineExercise) QueryRoutines() *RoutineQuery {
	return NewRoutineExerciseClient(re.config).QueryRoutines(re)
}

// QueryExercises queries the "exercises" edge of the RoutineExercise entity.
func (re *RoutineExercise) QueryExercises() *ExerciseQuery {
	return NewRoutineExerciseClient(re.config).QueryExercises(re)
}

// QueryUsers queries the "users" edge of the RoutineExercise entity.
func (re *RoutineExercise) QueryUsers() *UserQuery {
	return NewRoutineExerciseClient(re.config).QueryUsers(re)
}

// Update returns a builder for updating this RoutineExercise.
// Note that you need to call RoutineExercise.Unwrap() before calling this method if this RoutineExercise
// was returned from a transaction, and the transaction was committed or rolled back.
func (re *RoutineExercise) Update() *RoutineExerciseUpdateOne {
	return NewRoutineExerciseClient(re.config).UpdateOne(re)
}

// Unwrap unwraps the RoutineExercise entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (re *RoutineExercise) Unwrap() *RoutineExercise {
	_tx, ok := re.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoutineExercise is not a transactional entity")
	}
	re.config.driver = _tx.drv
	return re
}

// String implements the fmt.Stringer.
func (re *RoutineExercise) String() string {
	var builder strings.Builder
	builder.WriteString("RoutineExercise(")
	builder.WriteString(fmt.Sprintf("id=%v, ", re.ID))
	if v := re.RestTime; v != nil {
		builder.WriteString("rest_time=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("sets=")
	builder.WriteString(fmt.Sprintf("%v", re.Sets))
	builder.WriteString(", ")
	builder.WriteString("routine_id=")
	builder.WriteString(fmt.Sprintf("%v", re.RoutineID))
	builder.WriteString(", ")
	builder.WriteString("exercise_id=")
	builder.WriteString(fmt.Sprintf("%v", re.ExerciseID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", re.UserID))
	builder.WriteString(", ")
	builder.WriteString("order=")
	builder.WriteString(fmt.Sprintf("%v", re.Order))
	builder.WriteByte(')')
	return builder.String()
}

// RoutineExercises is a parsable slice of RoutineExercise.
type RoutineExercises []*RoutineExercise
