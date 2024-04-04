// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID pksuid.ID `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// HashedPassword holds the value of the "hashed_password" field.
	HashedPassword string `json:"-"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
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
	// totalCount holds the count of the edges above.
	totalCount [6]map[string]int

	namedTokens           map[string][]*Token
	namedExercises        map[string][]*Exercise
	namedRoutines         map[string][]*Routine
	namedWorkouts         map[string][]*Workout
	namedWorkoutLogs      map[string][]*WorkoutLog
	namedRoutineExercises map[string][]*RoutineExercise
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
		case user.FieldID:
			values[i] = new(pksuid.ID)
		case user.FieldActivated, user.FieldVersion:
			values[i] = new(sql.NullInt64)
		case user.FieldEmail, user.FieldUsername, user.FieldHashedPassword, user.FieldName:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt:
			values[i] = new(sql.NullTime)
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
			if value, ok := values[i].(*pksuid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
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
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
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
	builder.WriteString("hashed_password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("activated=")
	builder.WriteString(fmt.Sprintf("%v", u.Activated))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", u.Version))
	builder.WriteByte(')')
	return builder.String()
}

// NamedTokens returns the Tokens named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedTokens(name string) ([]*Token, error) {
	if u.Edges.namedTokens == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedTokens[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedTokens(name string, edges ...*Token) {
	if u.Edges.namedTokens == nil {
		u.Edges.namedTokens = make(map[string][]*Token)
	}
	if len(edges) == 0 {
		u.Edges.namedTokens[name] = []*Token{}
	} else {
		u.Edges.namedTokens[name] = append(u.Edges.namedTokens[name], edges...)
	}
}

// NamedExercises returns the Exercises named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedExercises(name string) ([]*Exercise, error) {
	if u.Edges.namedExercises == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedExercises[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedExercises(name string, edges ...*Exercise) {
	if u.Edges.namedExercises == nil {
		u.Edges.namedExercises = make(map[string][]*Exercise)
	}
	if len(edges) == 0 {
		u.Edges.namedExercises[name] = []*Exercise{}
	} else {
		u.Edges.namedExercises[name] = append(u.Edges.namedExercises[name], edges...)
	}
}

// NamedRoutines returns the Routines named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedRoutines(name string) ([]*Routine, error) {
	if u.Edges.namedRoutines == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedRoutines[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedRoutines(name string, edges ...*Routine) {
	if u.Edges.namedRoutines == nil {
		u.Edges.namedRoutines = make(map[string][]*Routine)
	}
	if len(edges) == 0 {
		u.Edges.namedRoutines[name] = []*Routine{}
	} else {
		u.Edges.namedRoutines[name] = append(u.Edges.namedRoutines[name], edges...)
	}
}

// NamedWorkouts returns the Workouts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedWorkouts(name string) ([]*Workout, error) {
	if u.Edges.namedWorkouts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedWorkouts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedWorkouts(name string, edges ...*Workout) {
	if u.Edges.namedWorkouts == nil {
		u.Edges.namedWorkouts = make(map[string][]*Workout)
	}
	if len(edges) == 0 {
		u.Edges.namedWorkouts[name] = []*Workout{}
	} else {
		u.Edges.namedWorkouts[name] = append(u.Edges.namedWorkouts[name], edges...)
	}
}

// NamedWorkoutLogs returns the WorkoutLogs named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedWorkoutLogs(name string) ([]*WorkoutLog, error) {
	if u.Edges.namedWorkoutLogs == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedWorkoutLogs[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedWorkoutLogs(name string, edges ...*WorkoutLog) {
	if u.Edges.namedWorkoutLogs == nil {
		u.Edges.namedWorkoutLogs = make(map[string][]*WorkoutLog)
	}
	if len(edges) == 0 {
		u.Edges.namedWorkoutLogs[name] = []*WorkoutLog{}
	} else {
		u.Edges.namedWorkoutLogs[name] = append(u.Edges.namedWorkoutLogs[name], edges...)
	}
}

// NamedRoutineExercises returns the RoutineExercises named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedRoutineExercises(name string) ([]*RoutineExercise, error) {
	if u.Edges.namedRoutineExercises == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedRoutineExercises[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedRoutineExercises(name string, edges ...*RoutineExercise) {
	if u.Edges.namedRoutineExercises == nil {
		u.Edges.namedRoutineExercises = make(map[string][]*RoutineExercise)
	}
	if len(edges) == 0 {
		u.Edges.namedRoutineExercises[name] = []*RoutineExercise{}
	} else {
		u.Edges.namedRoutineExercises[name] = append(u.Edges.namedRoutineExercises[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User
