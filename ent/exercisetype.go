// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sahidrahman404/gigachad-api/ent/exercisetype"
)

// ExerciseType is the model entity for the ExerciseType schema.
type ExerciseType struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Properties holds the value of the "properties" field.
	Properties []string `json:"properties,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExerciseTypeQuery when eager-loading is set.
	Edges        ExerciseTypeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ExerciseTypeEdges holds the relations/edges for other nodes in the graph.
type ExerciseTypeEdges struct {
	// Exercises holds the value of the exercises edge.
	Exercises []*Exercise `json:"exercises,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedExercises map[string][]*Exercise
}

// ExercisesOrErr returns the Exercises value or an error if the edge
// was not loaded in eager-loading.
func (e ExerciseTypeEdges) ExercisesOrErr() ([]*Exercise, error) {
	if e.loadedTypes[0] {
		return e.Exercises, nil
	}
	return nil, &NotLoadedError{edge: "exercises"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExerciseType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case exercisetype.FieldProperties:
			values[i] = new([]byte)
		case exercisetype.FieldID, exercisetype.FieldName, exercisetype.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExerciseType fields.
func (et *ExerciseType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case exercisetype.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				et.ID = value.String
			}
		case exercisetype.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				et.Name = value.String
			}
		case exercisetype.FieldProperties:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field properties", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &et.Properties); err != nil {
					return fmt.Errorf("unmarshal field properties: %w", err)
				}
			}
		case exercisetype.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				et.Description = value.String
			}
		default:
			et.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ExerciseType.
// This includes values selected through modifiers, order, etc.
func (et *ExerciseType) Value(name string) (ent.Value, error) {
	return et.selectValues.Get(name)
}

// QueryExercises queries the "exercises" edge of the ExerciseType entity.
func (et *ExerciseType) QueryExercises() *ExerciseQuery {
	return NewExerciseTypeClient(et.config).QueryExercises(et)
}

// Update returns a builder for updating this ExerciseType.
// Note that you need to call ExerciseType.Unwrap() before calling this method if this ExerciseType
// was returned from a transaction, and the transaction was committed or rolled back.
func (et *ExerciseType) Update() *ExerciseTypeUpdateOne {
	return NewExerciseTypeClient(et.config).UpdateOne(et)
}

// Unwrap unwraps the ExerciseType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (et *ExerciseType) Unwrap() *ExerciseType {
	_tx, ok := et.config.driver.(*txDriver)
	if !ok {
		panic("ent: ExerciseType is not a transactional entity")
	}
	et.config.driver = _tx.drv
	return et
}

// String implements the fmt.Stringer.
func (et *ExerciseType) String() string {
	var builder strings.Builder
	builder.WriteString("ExerciseType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", et.ID))
	builder.WriteString("name=")
	builder.WriteString(et.Name)
	builder.WriteString(", ")
	builder.WriteString("properties=")
	builder.WriteString(fmt.Sprintf("%v", et.Properties))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(et.Description)
	builder.WriteByte(')')
	return builder.String()
}

// NamedExercises returns the Exercises named value or an error if the edge was not
// loaded in eager-loading with this name.
func (et *ExerciseType) NamedExercises(name string) ([]*Exercise, error) {
	if et.Edges.namedExercises == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := et.Edges.namedExercises[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (et *ExerciseType) appendNamedExercises(name string, edges ...*Exercise) {
	if et.Edges.namedExercises == nil {
		et.Edges.namedExercises = make(map[string][]*Exercise)
	}
	if len(edges) == 0 {
		et.Edges.namedExercises[name] = []*Exercise{}
	} else {
		et.Edges.namedExercises[name] = append(et.Edges.namedExercises[name], edges...)
	}
}

// ExerciseTypes is a parsable slice of ExerciseType.
type ExerciseTypes []*ExerciseType
