// Code generated by ent, DO NOT EDIT.

package user

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldHashedPassword holds the string denoting the hashed_password field in the database.
	FieldHashedPassword = "hashed_password"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldUnit holds the string denoting the unit field in the database.
	FieldUnit = "unit"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldActivated holds the string denoting the activated field in the database.
	FieldActivated = "activated"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// EdgeTokens holds the string denoting the tokens edge name in mutations.
	EdgeTokens = "tokens"
	// EdgeExercises holds the string denoting the exercises edge name in mutations.
	EdgeExercises = "exercises"
	// EdgeRoutines holds the string denoting the routines edge name in mutations.
	EdgeRoutines = "routines"
	// EdgeWorkouts holds the string denoting the workouts edge name in mutations.
	EdgeWorkouts = "workouts"
	// EdgeWorkoutLogs holds the string denoting the workout_logs edge name in mutations.
	EdgeWorkoutLogs = "workout_logs"
	// EdgeRoutineExercises holds the string denoting the routine_exercises edge name in mutations.
	EdgeRoutineExercises = "routine_exercises"
	// Table holds the table name of the user in the database.
	Table = "users"
	// TokensTable is the table that holds the tokens relation/edge.
	TokensTable = "tokens"
	// TokensInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokensInverseTable = "tokens"
	// TokensColumn is the table column denoting the tokens relation/edge.
	TokensColumn = "user_id"
	// ExercisesTable is the table that holds the exercises relation/edge.
	ExercisesTable = "exercises"
	// ExercisesInverseTable is the table name for the Exercise entity.
	// It exists in this package in order to avoid circular dependency with the "exercise" package.
	ExercisesInverseTable = "exercises"
	// ExercisesColumn is the table column denoting the exercises relation/edge.
	ExercisesColumn = "user_id"
	// RoutinesTable is the table that holds the routines relation/edge.
	RoutinesTable = "routines"
	// RoutinesInverseTable is the table name for the Routine entity.
	// It exists in this package in order to avoid circular dependency with the "routine" package.
	RoutinesInverseTable = "routines"
	// RoutinesColumn is the table column denoting the routines relation/edge.
	RoutinesColumn = "user_id"
	// WorkoutsTable is the table that holds the workouts relation/edge.
	WorkoutsTable = "workouts"
	// WorkoutsInverseTable is the table name for the Workout entity.
	// It exists in this package in order to avoid circular dependency with the "workout" package.
	WorkoutsInverseTable = "workouts"
	// WorkoutsColumn is the table column denoting the workouts relation/edge.
	WorkoutsColumn = "user_id"
	// WorkoutLogsTable is the table that holds the workout_logs relation/edge.
	WorkoutLogsTable = "workout_logs"
	// WorkoutLogsInverseTable is the table name for the WorkoutLog entity.
	// It exists in this package in order to avoid circular dependency with the "workoutlog" package.
	WorkoutLogsInverseTable = "workout_logs"
	// WorkoutLogsColumn is the table column denoting the workout_logs relation/edge.
	WorkoutLogsColumn = "user_id"
	// RoutineExercisesTable is the table that holds the routine_exercises relation/edge.
	RoutineExercisesTable = "routine_exercises"
	// RoutineExercisesInverseTable is the table name for the RoutineExercise entity.
	// It exists in this package in order to avoid circular dependency with the "routineexercise" package.
	RoutineExercisesInverseTable = "routine_exercises"
	// RoutineExercisesColumn is the table column denoting the routine_exercises relation/edge.
	RoutineExercisesColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldUsername,
	FieldHashedPassword,
	FieldName,
	FieldUnit,
	FieldCreatedAt,
	FieldActivated,
	FieldVersion,
}

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
	DefaultCreatedAt time.Time
	// DefaultActivated holds the default value on creation for the "activated" field.
	DefaultActivated int
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pksuid.ID
)

// Unit defines the type for the "unit" enum field.
type Unit string

// UnitMETRIC is the default value of the Unit enum.
const DefaultUnit = UnitMETRIC

// Unit values.
const (
	UnitMETRIC   Unit = "METRIC"
	UnitIMPERIAL Unit = "IMPERIAL"
)

func (u Unit) String() string {
	return string(u)
}

// UnitValidator is a validator for the "unit" field enum values. It is called by the builders before save.
func UnitValidator(u Unit) error {
	switch u {
	case UnitMETRIC, UnitIMPERIAL:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for unit field: %q", u)
	}
}

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByHashedPassword orders the results by the hashed_password field.
func ByHashedPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHashedPassword, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByUnit orders the results by the unit field.
func ByUnit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnit, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByActivated orders the results by the activated field.
func ByActivated(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActivated, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByTokensCount orders the results by tokens count.
func ByTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTokensStep(), opts...)
	}
}

// ByTokens orders the results by tokens terms.
func ByTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByRoutinesCount orders the results by routines count.
func ByRoutinesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRoutinesStep(), opts...)
	}
}

// ByRoutines orders the results by routines terms.
func ByRoutines(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoutinesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkoutsCount orders the results by workouts count.
func ByWorkoutsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkoutsStep(), opts...)
	}
}

// ByWorkouts orders the results by workouts terms.
func ByWorkouts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkoutsStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TokensTable, TokensColumn),
	)
}
func newExercisesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExercisesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExercisesTable, ExercisesColumn),
	)
}
func newRoutinesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoutinesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RoutinesTable, RoutinesColumn),
	)
}
func newWorkoutsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkoutsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkoutsTable, WorkoutsColumn),
	)
}
func newWorkoutLogsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkoutLogsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkoutLogsTable, WorkoutLogsColumn),
	)
}
func newRoutineExercisesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoutineExercisesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RoutineExercisesTable, RoutineExercisesColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Unit) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Unit) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Unit(str)
	if err := UnitValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Unit", str)
	}
	return nil
}
