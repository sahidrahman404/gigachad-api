// Code generated by ent, DO NOT EDIT.

package routine

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

const (
	// Label holds the string label denoting the routine type in the database.
	Label = "routine"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeExercises holds the string denoting the exercises edge name in mutations.
	EdgeExercises = "exercises"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeRoutineExercises holds the string denoting the routine_exercises edge name in mutations.
	EdgeRoutineExercises = "routine_exercises"
	// Table holds the table name of the routine in the database.
	Table = "routines"
	// ExercisesTable is the table that holds the exercises relation/edge. The primary key declared below.
	ExercisesTable = "routine_exercises"
	// ExercisesInverseTable is the table name for the Exercise entity.
	// It exists in this package in order to avoid circular dependency with the "exercise" package.
	ExercisesInverseTable = "exercises"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "routines"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_id"
	// RoutineExercisesTable is the table that holds the routine_exercises relation/edge.
	RoutineExercisesTable = "routine_exercises"
	// RoutineExercisesInverseTable is the table name for the RoutineExercise entity.
	// It exists in this package in order to avoid circular dependency with the "routineexercise" package.
	RoutineExercisesInverseTable = "routine_exercises"
	// RoutineExercisesColumn is the table column denoting the routine_exercises relation/edge.
	RoutineExercisesColumn = "routine_id"
)

// Columns holds all SQL columns for routine fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldUserID,
}

var (
	// ExercisesPrimaryKey and ExercisesColumn2 are the table columns denoting the
	// primary key for the exercises relation (M2M).
	ExercisesPrimaryKey = []string{"routine_id", "exercise_id"}
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/sahidrahman404/gigachad-api/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pksuid.ID
)

// OrderOption defines the ordering options for the Routine queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
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

// ByUsersField orders the results by users field.
func ByUsersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), sql.OrderByField(field, opts...))
	}
}

// ByRoutineExercisesCount orders the results by routine_exercises count.
func ByRoutineExercisesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRoutineExercisesStep(), opts...)
	}
}

// ByRoutineExercises orders the results by routine_exercises terms.
func ByRoutineExercises(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoutineExercisesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newExercisesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExercisesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ExercisesTable, ExercisesPrimaryKey...),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
	)
}
func newRoutineExercisesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoutineExercisesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, RoutineExercisesTable, RoutineExercisesColumn),
	)
}
