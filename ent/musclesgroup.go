// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sahidrahman404/gigachad-api/ent/musclesgroup"
)

// MusclesGroup is the model entity for the MusclesGroup schema.
type MusclesGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Image holds the value of the "image" field.
	Image string `json:"image,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MusclesGroupQuery when eager-loading is set.
	Edges        MusclesGroupEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MusclesGroupEdges holds the relations/edges for other nodes in the graph.
type MusclesGroupEdges struct {
	// Exercises holds the value of the exercises edge.
	Exercises []*Exercise `json:"exercises,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ExercisesOrErr returns the Exercises value or an error if the edge
// was not loaded in eager-loading.
func (e MusclesGroupEdges) ExercisesOrErr() ([]*Exercise, error) {
	if e.loadedTypes[0] {
		return e.Exercises, nil
	}
	return nil, &NotLoadedError{edge: "exercises"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MusclesGroup) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case musclesgroup.FieldID, musclesgroup.FieldName, musclesgroup.FieldImage:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MusclesGroup fields.
func (mg *MusclesGroup) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case musclesgroup.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				mg.ID = value.String
			}
		case musclesgroup.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mg.Name = value.String
			}
		case musclesgroup.FieldImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value.Valid {
				mg.Image = value.String
			}
		default:
			mg.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MusclesGroup.
// This includes values selected through modifiers, order, etc.
func (mg *MusclesGroup) Value(name string) (ent.Value, error) {
	return mg.selectValues.Get(name)
}

// QueryExercises queries the "exercises" edge of the MusclesGroup entity.
func (mg *MusclesGroup) QueryExercises() *ExerciseQuery {
	return NewMusclesGroupClient(mg.config).QueryExercises(mg)
}

// Update returns a builder for updating this MusclesGroup.
// Note that you need to call MusclesGroup.Unwrap() before calling this method if this MusclesGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (mg *MusclesGroup) Update() *MusclesGroupUpdateOne {
	return NewMusclesGroupClient(mg.config).UpdateOne(mg)
}

// Unwrap unwraps the MusclesGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mg *MusclesGroup) Unwrap() *MusclesGroup {
	_tx, ok := mg.config.driver.(*txDriver)
	if !ok {
		panic("ent: MusclesGroup is not a transactional entity")
	}
	mg.config.driver = _tx.drv
	return mg
}

// String implements the fmt.Stringer.
func (mg *MusclesGroup) String() string {
	var builder strings.Builder
	builder.WriteString("MusclesGroup(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mg.ID))
	builder.WriteString("name=")
	builder.WriteString(mg.Name)
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(mg.Image)
	builder.WriteByte(')')
	return builder.String()
}

// MusclesGroups is a parsable slice of MusclesGroup.
type MusclesGroups []*MusclesGroup
