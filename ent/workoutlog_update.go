// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/ent/workout"
	"github.com/sahidrahman404/gigachad-api/ent/workoutlog"
)

// WorkoutLogUpdate is the builder for updating WorkoutLog entities.
type WorkoutLogUpdate struct {
	config
	hooks    []Hook
	mutation *WorkoutLogMutation
}

// Where appends a list predicates to the WorkoutLogUpdate builder.
func (wlu *WorkoutLogUpdate) Where(ps ...predicate.WorkoutLog) *WorkoutLogUpdate {
	wlu.mutation.Where(ps...)
	return wlu
}

// SetSets sets the "sets" field.
func (wlu *WorkoutLogUpdate) SetSets(s []*schematype.Set) *WorkoutLogUpdate {
	wlu.mutation.SetSets(s)
	return wlu
}

// AppendSets appends s to the "sets" field.
func (wlu *WorkoutLogUpdate) AppendSets(s []*schematype.Set) *WorkoutLogUpdate {
	wlu.mutation.AppendSets(s)
	return wlu
}

// SetCreatedAt sets the "created_at" field.
func (wlu *WorkoutLogUpdate) SetCreatedAt(t time.Time) *WorkoutLogUpdate {
	wlu.mutation.SetCreatedAt(t)
	return wlu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wlu *WorkoutLogUpdate) SetNillableCreatedAt(t *time.Time) *WorkoutLogUpdate {
	if t != nil {
		wlu.SetCreatedAt(*t)
	}
	return wlu
}

// SetWorkoutID sets the "workout_id" field.
func (wlu *WorkoutLogUpdate) SetWorkoutID(pk pksuid.ID) *WorkoutLogUpdate {
	wlu.mutation.SetWorkoutID(pk)
	return wlu
}

// SetNillableWorkoutID sets the "workout_id" field if the given value is not nil.
func (wlu *WorkoutLogUpdate) SetNillableWorkoutID(pk *pksuid.ID) *WorkoutLogUpdate {
	if pk != nil {
		wlu.SetWorkoutID(*pk)
	}
	return wlu
}

// SetExerciseID sets the "exercise_id" field.
func (wlu *WorkoutLogUpdate) SetExerciseID(pk pksuid.ID) *WorkoutLogUpdate {
	wlu.mutation.SetExerciseID(pk)
	return wlu
}

// SetNillableExerciseID sets the "exercise_id" field if the given value is not nil.
func (wlu *WorkoutLogUpdate) SetNillableExerciseID(pk *pksuid.ID) *WorkoutLogUpdate {
	if pk != nil {
		wlu.SetExerciseID(*pk)
	}
	return wlu
}

// SetUserID sets the "user_id" field.
func (wlu *WorkoutLogUpdate) SetUserID(pk pksuid.ID) *WorkoutLogUpdate {
	wlu.mutation.SetUserID(pk)
	return wlu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wlu *WorkoutLogUpdate) SetNillableUserID(pk *pksuid.ID) *WorkoutLogUpdate {
	if pk != nil {
		wlu.SetUserID(*pk)
	}
	return wlu
}

// SetOrder sets the "order" field.
func (wlu *WorkoutLogUpdate) SetOrder(i int) *WorkoutLogUpdate {
	wlu.mutation.ResetOrder()
	wlu.mutation.SetOrder(i)
	return wlu
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (wlu *WorkoutLogUpdate) SetNillableOrder(i *int) *WorkoutLogUpdate {
	if i != nil {
		wlu.SetOrder(*i)
	}
	return wlu
}

// AddOrder adds i to the "order" field.
func (wlu *WorkoutLogUpdate) AddOrder(i int) *WorkoutLogUpdate {
	wlu.mutation.AddOrder(i)
	return wlu
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (wlu *WorkoutLogUpdate) SetUsersID(id pksuid.ID) *WorkoutLogUpdate {
	wlu.mutation.SetUsersID(id)
	return wlu
}

// SetUsers sets the "users" edge to the User entity.
func (wlu *WorkoutLogUpdate) SetUsers(u *User) *WorkoutLogUpdate {
	return wlu.SetUsersID(u.ID)
}

// SetWorkoutsID sets the "workouts" edge to the Workout entity by ID.
func (wlu *WorkoutLogUpdate) SetWorkoutsID(id pksuid.ID) *WorkoutLogUpdate {
	wlu.mutation.SetWorkoutsID(id)
	return wlu
}

// SetWorkouts sets the "workouts" edge to the Workout entity.
func (wlu *WorkoutLogUpdate) SetWorkouts(w *Workout) *WorkoutLogUpdate {
	return wlu.SetWorkoutsID(w.ID)
}

// SetExercisesID sets the "exercises" edge to the Exercise entity by ID.
func (wlu *WorkoutLogUpdate) SetExercisesID(id pksuid.ID) *WorkoutLogUpdate {
	wlu.mutation.SetExercisesID(id)
	return wlu
}

// SetExercises sets the "exercises" edge to the Exercise entity.
func (wlu *WorkoutLogUpdate) SetExercises(e *Exercise) *WorkoutLogUpdate {
	return wlu.SetExercisesID(e.ID)
}

// Mutation returns the WorkoutLogMutation object of the builder.
func (wlu *WorkoutLogUpdate) Mutation() *WorkoutLogMutation {
	return wlu.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (wlu *WorkoutLogUpdate) ClearUsers() *WorkoutLogUpdate {
	wlu.mutation.ClearUsers()
	return wlu
}

// ClearWorkouts clears the "workouts" edge to the Workout entity.
func (wlu *WorkoutLogUpdate) ClearWorkouts() *WorkoutLogUpdate {
	wlu.mutation.ClearWorkouts()
	return wlu
}

// ClearExercises clears the "exercises" edge to the Exercise entity.
func (wlu *WorkoutLogUpdate) ClearExercises() *WorkoutLogUpdate {
	wlu.mutation.ClearExercises()
	return wlu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wlu *WorkoutLogUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wlu.sqlSave, wlu.mutation, wlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wlu *WorkoutLogUpdate) SaveX(ctx context.Context) int {
	affected, err := wlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wlu *WorkoutLogUpdate) Exec(ctx context.Context) error {
	_, err := wlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wlu *WorkoutLogUpdate) ExecX(ctx context.Context) {
	if err := wlu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wlu *WorkoutLogUpdate) check() error {
	if _, ok := wlu.mutation.UsersID(); wlu.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkoutLog.users"`)
	}
	if _, ok := wlu.mutation.WorkoutsID(); wlu.mutation.WorkoutsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkoutLog.workouts"`)
	}
	if _, ok := wlu.mutation.ExercisesID(); wlu.mutation.ExercisesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkoutLog.exercises"`)
	}
	return nil
}

func (wlu *WorkoutLogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wlu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(workoutlog.Table, workoutlog.Columns, sqlgraph.NewFieldSpec(workoutlog.FieldID, field.TypeString))
	if ps := wlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wlu.mutation.Sets(); ok {
		_spec.SetField(workoutlog.FieldSets, field.TypeJSON, value)
	}
	if value, ok := wlu.mutation.AppendedSets(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workoutlog.FieldSets, value)
		})
	}
	if value, ok := wlu.mutation.CreatedAt(); ok {
		_spec.SetField(workoutlog.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := wlu.mutation.Order(); ok {
		_spec.SetField(workoutlog.FieldOrder, field.TypeInt, value)
	}
	if value, ok := wlu.mutation.AddedOrder(); ok {
		_spec.AddField(workoutlog.FieldOrder, field.TypeInt, value)
	}
	if wlu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workoutlog.UsersTable,
			Columns: []string{workoutlog.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wlu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workoutlog.UsersTable,
			Columns: []string{workoutlog.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wlu.mutation.WorkoutsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutlog.WorkoutsTable,
			Columns: []string{workoutlog.WorkoutsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workout.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wlu.mutation.WorkoutsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutlog.WorkoutsTable,
			Columns: []string{workoutlog.WorkoutsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workout.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wlu.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutlog.ExercisesTable,
			Columns: []string{workoutlog.ExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wlu.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutlog.ExercisesTable,
			Columns: []string{workoutlog.ExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workoutlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wlu.mutation.done = true
	return n, nil
}

// WorkoutLogUpdateOne is the builder for updating a single WorkoutLog entity.
type WorkoutLogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkoutLogMutation
}

// SetSets sets the "sets" field.
func (wluo *WorkoutLogUpdateOne) SetSets(s []*schematype.Set) *WorkoutLogUpdateOne {
	wluo.mutation.SetSets(s)
	return wluo
}

// AppendSets appends s to the "sets" field.
func (wluo *WorkoutLogUpdateOne) AppendSets(s []*schematype.Set) *WorkoutLogUpdateOne {
	wluo.mutation.AppendSets(s)
	return wluo
}

// SetCreatedAt sets the "created_at" field.
func (wluo *WorkoutLogUpdateOne) SetCreatedAt(t time.Time) *WorkoutLogUpdateOne {
	wluo.mutation.SetCreatedAt(t)
	return wluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wluo *WorkoutLogUpdateOne) SetNillableCreatedAt(t *time.Time) *WorkoutLogUpdateOne {
	if t != nil {
		wluo.SetCreatedAt(*t)
	}
	return wluo
}

// SetWorkoutID sets the "workout_id" field.
func (wluo *WorkoutLogUpdateOne) SetWorkoutID(pk pksuid.ID) *WorkoutLogUpdateOne {
	wluo.mutation.SetWorkoutID(pk)
	return wluo
}

// SetNillableWorkoutID sets the "workout_id" field if the given value is not nil.
func (wluo *WorkoutLogUpdateOne) SetNillableWorkoutID(pk *pksuid.ID) *WorkoutLogUpdateOne {
	if pk != nil {
		wluo.SetWorkoutID(*pk)
	}
	return wluo
}

// SetExerciseID sets the "exercise_id" field.
func (wluo *WorkoutLogUpdateOne) SetExerciseID(pk pksuid.ID) *WorkoutLogUpdateOne {
	wluo.mutation.SetExerciseID(pk)
	return wluo
}

// SetNillableExerciseID sets the "exercise_id" field if the given value is not nil.
func (wluo *WorkoutLogUpdateOne) SetNillableExerciseID(pk *pksuid.ID) *WorkoutLogUpdateOne {
	if pk != nil {
		wluo.SetExerciseID(*pk)
	}
	return wluo
}

// SetUserID sets the "user_id" field.
func (wluo *WorkoutLogUpdateOne) SetUserID(pk pksuid.ID) *WorkoutLogUpdateOne {
	wluo.mutation.SetUserID(pk)
	return wluo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wluo *WorkoutLogUpdateOne) SetNillableUserID(pk *pksuid.ID) *WorkoutLogUpdateOne {
	if pk != nil {
		wluo.SetUserID(*pk)
	}
	return wluo
}

// SetOrder sets the "order" field.
func (wluo *WorkoutLogUpdateOne) SetOrder(i int) *WorkoutLogUpdateOne {
	wluo.mutation.ResetOrder()
	wluo.mutation.SetOrder(i)
	return wluo
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (wluo *WorkoutLogUpdateOne) SetNillableOrder(i *int) *WorkoutLogUpdateOne {
	if i != nil {
		wluo.SetOrder(*i)
	}
	return wluo
}

// AddOrder adds i to the "order" field.
func (wluo *WorkoutLogUpdateOne) AddOrder(i int) *WorkoutLogUpdateOne {
	wluo.mutation.AddOrder(i)
	return wluo
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (wluo *WorkoutLogUpdateOne) SetUsersID(id pksuid.ID) *WorkoutLogUpdateOne {
	wluo.mutation.SetUsersID(id)
	return wluo
}

// SetUsers sets the "users" edge to the User entity.
func (wluo *WorkoutLogUpdateOne) SetUsers(u *User) *WorkoutLogUpdateOne {
	return wluo.SetUsersID(u.ID)
}

// SetWorkoutsID sets the "workouts" edge to the Workout entity by ID.
func (wluo *WorkoutLogUpdateOne) SetWorkoutsID(id pksuid.ID) *WorkoutLogUpdateOne {
	wluo.mutation.SetWorkoutsID(id)
	return wluo
}

// SetWorkouts sets the "workouts" edge to the Workout entity.
func (wluo *WorkoutLogUpdateOne) SetWorkouts(w *Workout) *WorkoutLogUpdateOne {
	return wluo.SetWorkoutsID(w.ID)
}

// SetExercisesID sets the "exercises" edge to the Exercise entity by ID.
func (wluo *WorkoutLogUpdateOne) SetExercisesID(id pksuid.ID) *WorkoutLogUpdateOne {
	wluo.mutation.SetExercisesID(id)
	return wluo
}

// SetExercises sets the "exercises" edge to the Exercise entity.
func (wluo *WorkoutLogUpdateOne) SetExercises(e *Exercise) *WorkoutLogUpdateOne {
	return wluo.SetExercisesID(e.ID)
}

// Mutation returns the WorkoutLogMutation object of the builder.
func (wluo *WorkoutLogUpdateOne) Mutation() *WorkoutLogMutation {
	return wluo.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (wluo *WorkoutLogUpdateOne) ClearUsers() *WorkoutLogUpdateOne {
	wluo.mutation.ClearUsers()
	return wluo
}

// ClearWorkouts clears the "workouts" edge to the Workout entity.
func (wluo *WorkoutLogUpdateOne) ClearWorkouts() *WorkoutLogUpdateOne {
	wluo.mutation.ClearWorkouts()
	return wluo
}

// ClearExercises clears the "exercises" edge to the Exercise entity.
func (wluo *WorkoutLogUpdateOne) ClearExercises() *WorkoutLogUpdateOne {
	wluo.mutation.ClearExercises()
	return wluo
}

// Where appends a list predicates to the WorkoutLogUpdate builder.
func (wluo *WorkoutLogUpdateOne) Where(ps ...predicate.WorkoutLog) *WorkoutLogUpdateOne {
	wluo.mutation.Where(ps...)
	return wluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wluo *WorkoutLogUpdateOne) Select(field string, fields ...string) *WorkoutLogUpdateOne {
	wluo.fields = append([]string{field}, fields...)
	return wluo
}

// Save executes the query and returns the updated WorkoutLog entity.
func (wluo *WorkoutLogUpdateOne) Save(ctx context.Context) (*WorkoutLog, error) {
	return withHooks(ctx, wluo.sqlSave, wluo.mutation, wluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wluo *WorkoutLogUpdateOne) SaveX(ctx context.Context) *WorkoutLog {
	node, err := wluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wluo *WorkoutLogUpdateOne) Exec(ctx context.Context) error {
	_, err := wluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wluo *WorkoutLogUpdateOne) ExecX(ctx context.Context) {
	if err := wluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wluo *WorkoutLogUpdateOne) check() error {
	if _, ok := wluo.mutation.UsersID(); wluo.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkoutLog.users"`)
	}
	if _, ok := wluo.mutation.WorkoutsID(); wluo.mutation.WorkoutsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkoutLog.workouts"`)
	}
	if _, ok := wluo.mutation.ExercisesID(); wluo.mutation.ExercisesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "WorkoutLog.exercises"`)
	}
	return nil
}

func (wluo *WorkoutLogUpdateOne) sqlSave(ctx context.Context) (_node *WorkoutLog, err error) {
	if err := wluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(workoutlog.Table, workoutlog.Columns, sqlgraph.NewFieldSpec(workoutlog.FieldID, field.TypeString))
	id, ok := wluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "WorkoutLog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, workoutlog.FieldID)
		for _, f := range fields {
			if !workoutlog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != workoutlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wluo.mutation.Sets(); ok {
		_spec.SetField(workoutlog.FieldSets, field.TypeJSON, value)
	}
	if value, ok := wluo.mutation.AppendedSets(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, workoutlog.FieldSets, value)
		})
	}
	if value, ok := wluo.mutation.CreatedAt(); ok {
		_spec.SetField(workoutlog.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := wluo.mutation.Order(); ok {
		_spec.SetField(workoutlog.FieldOrder, field.TypeInt, value)
	}
	if value, ok := wluo.mutation.AddedOrder(); ok {
		_spec.AddField(workoutlog.FieldOrder, field.TypeInt, value)
	}
	if wluo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workoutlog.UsersTable,
			Columns: []string{workoutlog.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wluo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workoutlog.UsersTable,
			Columns: []string{workoutlog.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wluo.mutation.WorkoutsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutlog.WorkoutsTable,
			Columns: []string{workoutlog.WorkoutsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workout.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wluo.mutation.WorkoutsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutlog.WorkoutsTable,
			Columns: []string{workoutlog.WorkoutsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workout.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if wluo.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutlog.ExercisesTable,
			Columns: []string{workoutlog.ExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wluo.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   workoutlog.ExercisesTable,
			Columns: []string{workoutlog.ExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &WorkoutLog{config: wluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{workoutlog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wluo.mutation.done = true
	return _node, nil
}
