// Code generated by ent, DO NOT EDIT.

package equipment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

const (
	// Label holds the string label denoting the equipment type in the database.
	Label = "equipment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// EdgeExercises holds the string denoting the exercises edge name in mutations.
	EdgeExercises = "exercises"
	// Table holds the table name of the equipment in the database.
	Table = "equipment"
	// ExercisesTable is the table that holds the exercises relation/edge. The primary key declared below.
	ExercisesTable = "equipment_exercises"
	// ExercisesInverseTable is the table name for the Exercise entity.
	// It exists in this package in order to avoid circular dependency with the "exercise" package.
	ExercisesInverseTable = "exercises"
)

// Columns holds all SQL columns for equipment fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldImage,
}

var (
	// ExercisesPrimaryKey and ExercisesColumn2 are the table columns denoting the
	// primary key for the exercises relation (M2M).
	ExercisesPrimaryKey = []string{"equipment_id", "exercise_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pksuid.ID
)

// OrderOption defines the ordering options for the Equipment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByExercisesCount orders the results by exercises count.
func ByExercisesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExercisesStep(), opts...)
	}
}

// ByExercises orders the results by exercises terms.
func ByExercises(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExercisesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newExercisesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExercisesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ExercisesTable, ExercisesPrimaryKey...),
	)
}
