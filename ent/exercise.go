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
)

// Exercise is the model entity for the Exercise schema.
type Exercise struct {
	config `json:"-"`
	// ID of the ent.
	ID pksuid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Image holds the value of the "image" field.
	Image *schematype.Image `json:"image,omitempty"`
	// HowTo holds the value of the "how_to" field.
	HowTo *string `json:"how_to,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID *pksuid.ID `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExerciseQuery when eager-loading is set.
	Edges        ExerciseEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ExerciseEdges holds the relations/edges for other nodes in the graph.
type ExerciseEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// Equipment holds the value of the equipment edge.
	Equipment []*Equipment `json:"equipment,omitempty"`
	// MusclesGroups holds the value of the muscles_groups edge.
	MusclesGroups []*MusclesGroup `json:"muscles_groups,omitempty"`
	// ExerciseTypes holds the value of the exercise_types edge.
	ExerciseTypes []*ExerciseType `json:"exercise_types,omitempty"`
	// Routines holds the value of the routines edge.
	Routines []*Routine `json:"routines,omitempty"`
	// Workouts holds the value of the workouts edge.
	Workouts []*Workout `json:"workouts,omitempty"`
	// RoutineExercises holds the value of the routine_exercises edge.
	RoutineExercises []*RoutineExercise `json:"routine_exercises,omitempty"`
	// WorkoutLogs holds the value of the workout_logs edge.
	WorkoutLogs []*WorkoutLog `json:"workout_logs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
	// totalCount holds the count of the edges above.
	totalCount [8]map[string]int

	namedEquipment        map[string][]*Equipment
	namedMusclesGroups    map[string][]*MusclesGroup
	namedExerciseTypes    map[string][]*ExerciseType
	namedRoutines         map[string][]*Routine
	namedWorkouts         map[string][]*Workout
	namedRoutineExercises map[string][]*RoutineExercise
	namedWorkoutLogs      map[string][]*WorkoutLog
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ExerciseEdges) UsersOrErr() (*User, error) {
	if e.Users != nil {
		return e.Users, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "users"}
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading.
func (e ExerciseEdges) EquipmentOrErr() ([]*Equipment, error) {
	if e.loadedTypes[1] {
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// MusclesGroupsOrErr returns the MusclesGroups value or an error if the edge
// was not loaded in eager-loading.
func (e ExerciseEdges) MusclesGroupsOrErr() ([]*MusclesGroup, error) {
	if e.loadedTypes[2] {
		return e.MusclesGroups, nil
	}
	return nil, &NotLoadedError{edge: "muscles_groups"}
}

// ExerciseTypesOrErr returns the ExerciseTypes value or an error if the edge
// was not loaded in eager-loading.
func (e ExerciseEdges) ExerciseTypesOrErr() ([]*ExerciseType, error) {
	if e.loadedTypes[3] {
		return e.ExerciseTypes, nil
	}
	return nil, &NotLoadedError{edge: "exercise_types"}
}

// RoutinesOrErr returns the Routines value or an error if the edge
// was not loaded in eager-loading.
func (e ExerciseEdges) RoutinesOrErr() ([]*Routine, error) {
	if e.loadedTypes[4] {
		return e.Routines, nil
	}
	return nil, &NotLoadedError{edge: "routines"}
}

// WorkoutsOrErr returns the Workouts value or an error if the edge
// was not loaded in eager-loading.
func (e ExerciseEdges) WorkoutsOrErr() ([]*Workout, error) {
	if e.loadedTypes[5] {
		return e.Workouts, nil
	}
	return nil, &NotLoadedError{edge: "workouts"}
}

// RoutineExercisesOrErr returns the RoutineExercises value or an error if the edge
// was not loaded in eager-loading.
func (e ExerciseEdges) RoutineExercisesOrErr() ([]*RoutineExercise, error) {
	if e.loadedTypes[6] {
		return e.RoutineExercises, nil
	}
	return nil, &NotLoadedError{edge: "routine_exercises"}
}

// WorkoutLogsOrErr returns the WorkoutLogs value or an error if the edge
// was not loaded in eager-loading.
func (e ExerciseEdges) WorkoutLogsOrErr() ([]*WorkoutLog, error) {
	if e.loadedTypes[7] {
		return e.WorkoutLogs, nil
	}
	return nil, &NotLoadedError{edge: "workout_logs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Exercise) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case exercise.FieldUserID:
			values[i] = &sql.NullScanner{S: new(pksuid.ID)}
		case exercise.FieldImage:
			values[i] = new([]byte)
		case exercise.FieldID:
			values[i] = new(pksuid.ID)
		case exercise.FieldName, exercise.FieldHowTo:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Exercise fields.
func (e *Exercise) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case exercise.FieldID:
			if value, ok := values[i].(*pksuid.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case exercise.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case exercise.FieldImage:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field image", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &e.Image); err != nil {
					return fmt.Errorf("unmarshal field image: %w", err)
				}
			}
		case exercise.FieldHowTo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field how_to", values[i])
			} else if value.Valid {
				e.HowTo = new(string)
				*e.HowTo = value.String
			}
		case exercise.FieldUserID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				e.UserID = new(pksuid.ID)
				*e.UserID = *value.S.(*pksuid.ID)
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Exercise.
// This includes values selected through modifiers, order, etc.
func (e *Exercise) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Exercise entity.
func (e *Exercise) QueryUsers() *UserQuery {
	return NewExerciseClient(e.config).QueryUsers(e)
}

// QueryEquipment queries the "equipment" edge of the Exercise entity.
func (e *Exercise) QueryEquipment() *EquipmentQuery {
	return NewExerciseClient(e.config).QueryEquipment(e)
}

// QueryMusclesGroups queries the "muscles_groups" edge of the Exercise entity.
func (e *Exercise) QueryMusclesGroups() *MusclesGroupQuery {
	return NewExerciseClient(e.config).QueryMusclesGroups(e)
}

// QueryExerciseTypes queries the "exercise_types" edge of the Exercise entity.
func (e *Exercise) QueryExerciseTypes() *ExerciseTypeQuery {
	return NewExerciseClient(e.config).QueryExerciseTypes(e)
}

// QueryRoutines queries the "routines" edge of the Exercise entity.
func (e *Exercise) QueryRoutines() *RoutineQuery {
	return NewExerciseClient(e.config).QueryRoutines(e)
}

// QueryWorkouts queries the "workouts" edge of the Exercise entity.
func (e *Exercise) QueryWorkouts() *WorkoutQuery {
	return NewExerciseClient(e.config).QueryWorkouts(e)
}

// QueryRoutineExercises queries the "routine_exercises" edge of the Exercise entity.
func (e *Exercise) QueryRoutineExercises() *RoutineExerciseQuery {
	return NewExerciseClient(e.config).QueryRoutineExercises(e)
}

// QueryWorkoutLogs queries the "workout_logs" edge of the Exercise entity.
func (e *Exercise) QueryWorkoutLogs() *WorkoutLogQuery {
	return NewExerciseClient(e.config).QueryWorkoutLogs(e)
}

// Update returns a builder for updating this Exercise.
// Note that you need to call Exercise.Unwrap() before calling this method if this Exercise
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Exercise) Update() *ExerciseUpdateOne {
	return NewExerciseClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Exercise entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Exercise) Unwrap() *Exercise {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Exercise is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Exercise) String() string {
	var builder strings.Builder
	builder.WriteString("Exercise(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	builder.WriteString("image=")
	builder.WriteString(fmt.Sprintf("%v", e.Image))
	builder.WriteString(", ")
	if v := e.HowTo; v != nil {
		builder.WriteString("how_to=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := e.UserID; v != nil {
		builder.WriteString("user_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// NamedEquipment returns the Equipment named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Exercise) NamedEquipment(name string) ([]*Equipment, error) {
	if e.Edges.namedEquipment == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedEquipment[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Exercise) appendNamedEquipment(name string, edges ...*Equipment) {
	if e.Edges.namedEquipment == nil {
		e.Edges.namedEquipment = make(map[string][]*Equipment)
	}
	if len(edges) == 0 {
		e.Edges.namedEquipment[name] = []*Equipment{}
	} else {
		e.Edges.namedEquipment[name] = append(e.Edges.namedEquipment[name], edges...)
	}
}

// NamedMusclesGroups returns the MusclesGroups named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Exercise) NamedMusclesGroups(name string) ([]*MusclesGroup, error) {
	if e.Edges.namedMusclesGroups == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedMusclesGroups[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Exercise) appendNamedMusclesGroups(name string, edges ...*MusclesGroup) {
	if e.Edges.namedMusclesGroups == nil {
		e.Edges.namedMusclesGroups = make(map[string][]*MusclesGroup)
	}
	if len(edges) == 0 {
		e.Edges.namedMusclesGroups[name] = []*MusclesGroup{}
	} else {
		e.Edges.namedMusclesGroups[name] = append(e.Edges.namedMusclesGroups[name], edges...)
	}
}

// NamedExerciseTypes returns the ExerciseTypes named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Exercise) NamedExerciseTypes(name string) ([]*ExerciseType, error) {
	if e.Edges.namedExerciseTypes == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedExerciseTypes[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Exercise) appendNamedExerciseTypes(name string, edges ...*ExerciseType) {
	if e.Edges.namedExerciseTypes == nil {
		e.Edges.namedExerciseTypes = make(map[string][]*ExerciseType)
	}
	if len(edges) == 0 {
		e.Edges.namedExerciseTypes[name] = []*ExerciseType{}
	} else {
		e.Edges.namedExerciseTypes[name] = append(e.Edges.namedExerciseTypes[name], edges...)
	}
}

// NamedRoutines returns the Routines named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Exercise) NamedRoutines(name string) ([]*Routine, error) {
	if e.Edges.namedRoutines == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedRoutines[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Exercise) appendNamedRoutines(name string, edges ...*Routine) {
	if e.Edges.namedRoutines == nil {
		e.Edges.namedRoutines = make(map[string][]*Routine)
	}
	if len(edges) == 0 {
		e.Edges.namedRoutines[name] = []*Routine{}
	} else {
		e.Edges.namedRoutines[name] = append(e.Edges.namedRoutines[name], edges...)
	}
}

// NamedWorkouts returns the Workouts named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Exercise) NamedWorkouts(name string) ([]*Workout, error) {
	if e.Edges.namedWorkouts == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedWorkouts[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Exercise) appendNamedWorkouts(name string, edges ...*Workout) {
	if e.Edges.namedWorkouts == nil {
		e.Edges.namedWorkouts = make(map[string][]*Workout)
	}
	if len(edges) == 0 {
		e.Edges.namedWorkouts[name] = []*Workout{}
	} else {
		e.Edges.namedWorkouts[name] = append(e.Edges.namedWorkouts[name], edges...)
	}
}

// NamedRoutineExercises returns the RoutineExercises named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Exercise) NamedRoutineExercises(name string) ([]*RoutineExercise, error) {
	if e.Edges.namedRoutineExercises == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedRoutineExercises[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Exercise) appendNamedRoutineExercises(name string, edges ...*RoutineExercise) {
	if e.Edges.namedRoutineExercises == nil {
		e.Edges.namedRoutineExercises = make(map[string][]*RoutineExercise)
	}
	if len(edges) == 0 {
		e.Edges.namedRoutineExercises[name] = []*RoutineExercise{}
	} else {
		e.Edges.namedRoutineExercises[name] = append(e.Edges.namedRoutineExercises[name], edges...)
	}
}

// NamedWorkoutLogs returns the WorkoutLogs named value or an error if the edge was not
// loaded in eager-loading with this name.
func (e *Exercise) NamedWorkoutLogs(name string) ([]*WorkoutLog, error) {
	if e.Edges.namedWorkoutLogs == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := e.Edges.namedWorkoutLogs[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (e *Exercise) appendNamedWorkoutLogs(name string, edges ...*WorkoutLog) {
	if e.Edges.namedWorkoutLogs == nil {
		e.Edges.namedWorkoutLogs = make(map[string][]*WorkoutLog)
	}
	if len(edges) == 0 {
		e.Edges.namedWorkoutLogs[name] = []*WorkoutLog{}
	} else {
		e.Edges.namedWorkoutLogs[name] = append(e.Edges.namedWorkoutLogs[name], edges...)
	}
}

// Exercises is a parsable slice of Exercise.
type Exercises []*Exercise
