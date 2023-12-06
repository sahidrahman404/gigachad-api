// Code generated by ent, DO NOT EDIT.

package workout

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

const (
	// Label holds the string label denoting the workout type in the database.
	Label = "workout"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldVolume holds the string denoting the volume field in the database.
	FieldVolume = "volume"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldSets holds the string denoting the sets field in the database.
	FieldSets = "sets"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeExercises holds the string denoting the exercises edge name in mutations.
	EdgeExercises = "exercises"
	// EdgeWorkoutLogs holds the string denoting the workout_logs edge name in mutations.
	EdgeWorkoutLogs = "workout_logs"
	// Table holds the table name of the workout in the database.
	Table = "workouts"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "workouts"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_id"
	// ExercisesTable is the table that holds the exercises relation/edge. The primary key declared below.
	ExercisesTable = "workout_logs"
	// ExercisesInverseTable is the table name for the Exercise entity.
	// It exists in this package in order to avoid circular dependency with the "exercise" package.
	ExercisesInverseTable = "exercises"
	// WorkoutLogsTable is the table that holds the workout_logs relation/edge.
	WorkoutLogsTable = "workout_logs"
	// WorkoutLogsInverseTable is the table name for the WorkoutLog entity.
	// It exists in this package in order to avoid circular dependency with the "workoutlog" package.
	WorkoutLogsInverseTable = "workout_logs"
	// WorkoutLogsColumn is the table column denoting the workout_logs relation/edge.
	WorkoutLogsColumn = "workout_id"
)

// Columns holds all SQL columns for workout fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldVolume,
	FieldDuration,
	FieldSets,
	FieldCreatedAt,
	FieldImage,
	FieldDescription,
	FieldUserID,
}

var (
	// ExercisesPrimaryKey and ExercisesColumn2 are the table columns denoting the
	// primary key for the exercises relation (M2M).
	ExercisesPrimaryKey = []string{"workout_id", "exercise_id"}
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pksuid.ID
)

// OrderOption defines the ordering options for the Workout queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByVolume orders the results by the volume field.
func ByVolume(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVolume, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// BySets orders the results by the sets field.
func BySets(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSets, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByUsersField orders the results by users field.
func ByUsersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), sql.OrderByField(field, opts...))
	}
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

// ByWorkoutLogsCount orders the results by workout_logs count.
func ByWorkoutLogsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkoutLogsStep(), opts...)
	}
}

// ByWorkoutLogs orders the results by workout_logs terms.
func ByWorkoutLogs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkoutLogsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
	)
}
func newExercisesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExercisesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ExercisesTable, ExercisesPrimaryKey...),
	)
}
func newWorkoutLogsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkoutLogsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, WorkoutLogsTable, WorkoutLogsColumn),
	)
}
