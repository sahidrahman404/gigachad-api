// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
	"github.com/sahidrahman404/gigachad-api/ent/user"
)

// RoutineExerciseUpdate is the builder for updating RoutineExercise entities.
type RoutineExerciseUpdate struct {
	config
	hooks    []Hook
	mutation *RoutineExerciseMutation
}

// Where appends a list predicates to the RoutineExerciseUpdate builder.
func (reu *RoutineExerciseUpdate) Where(ps ...predicate.RoutineExercise) *RoutineExerciseUpdate {
	reu.mutation.Where(ps...)
	return reu
}

// SetRestTimer sets the "rest_timer" field.
func (reu *RoutineExerciseUpdate) SetRestTimer(i int) *RoutineExerciseUpdate {
	reu.mutation.ResetRestTimer()
	reu.mutation.SetRestTimer(i)
	return reu
}

// SetNillableRestTimer sets the "rest_timer" field if the given value is not nil.
func (reu *RoutineExerciseUpdate) SetNillableRestTimer(i *int) *RoutineExerciseUpdate {
	if i != nil {
		reu.SetRestTimer(*i)
	}
	return reu
}

// AddRestTimer adds i to the "rest_timer" field.
func (reu *RoutineExerciseUpdate) AddRestTimer(i int) *RoutineExerciseUpdate {
	reu.mutation.AddRestTimer(i)
	return reu
}

// ClearRestTimer clears the value of the "rest_timer" field.
func (reu *RoutineExerciseUpdate) ClearRestTimer() *RoutineExerciseUpdate {
	reu.mutation.ClearRestTimer()
	return reu
}

// SetSets sets the "sets" field.
func (reu *RoutineExerciseUpdate) SetSets(s *schematype.Sets) *RoutineExerciseUpdate {
	reu.mutation.SetSets(s)
	return reu
}

// SetRoutineID sets the "routine_id" field.
func (reu *RoutineExerciseUpdate) SetRoutineID(s string) *RoutineExerciseUpdate {
	reu.mutation.SetRoutineID(s)
	return reu
}

// SetExerciseID sets the "exercise_id" field.
func (reu *RoutineExerciseUpdate) SetExerciseID(s string) *RoutineExerciseUpdate {
	reu.mutation.SetExerciseID(s)
	return reu
}

// SetUserID sets the "user_id" field.
func (reu *RoutineExerciseUpdate) SetUserID(s string) *RoutineExerciseUpdate {
	reu.mutation.SetUserID(s)
	return reu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (reu *RoutineExerciseUpdate) SetNillableUserID(s *string) *RoutineExerciseUpdate {
	if s != nil {
		reu.SetUserID(*s)
	}
	return reu
}

// ClearUserID clears the value of the "user_id" field.
func (reu *RoutineExerciseUpdate) ClearUserID() *RoutineExerciseUpdate {
	reu.mutation.ClearUserID()
	return reu
}

// SetRoutinesID sets the "routines" edge to the Routine entity by ID.
func (reu *RoutineExerciseUpdate) SetRoutinesID(id string) *RoutineExerciseUpdate {
	reu.mutation.SetRoutinesID(id)
	return reu
}

// SetRoutines sets the "routines" edge to the Routine entity.
func (reu *RoutineExerciseUpdate) SetRoutines(r *Routine) *RoutineExerciseUpdate {
	return reu.SetRoutinesID(r.ID)
}

// SetExercisesID sets the "exercises" edge to the Exercise entity by ID.
func (reu *RoutineExerciseUpdate) SetExercisesID(id string) *RoutineExerciseUpdate {
	reu.mutation.SetExercisesID(id)
	return reu
}

// SetExercises sets the "exercises" edge to the Exercise entity.
func (reu *RoutineExerciseUpdate) SetExercises(e *Exercise) *RoutineExerciseUpdate {
	return reu.SetExercisesID(e.ID)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (reu *RoutineExerciseUpdate) SetUsersID(id string) *RoutineExerciseUpdate {
	reu.mutation.SetUsersID(id)
	return reu
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (reu *RoutineExerciseUpdate) SetNillableUsersID(id *string) *RoutineExerciseUpdate {
	if id != nil {
		reu = reu.SetUsersID(*id)
	}
	return reu
}

// SetUsers sets the "users" edge to the User entity.
func (reu *RoutineExerciseUpdate) SetUsers(u *User) *RoutineExerciseUpdate {
	return reu.SetUsersID(u.ID)
}

// Mutation returns the RoutineExerciseMutation object of the builder.
func (reu *RoutineExerciseUpdate) Mutation() *RoutineExerciseMutation {
	return reu.mutation
}

// ClearRoutines clears the "routines" edge to the Routine entity.
func (reu *RoutineExerciseUpdate) ClearRoutines() *RoutineExerciseUpdate {
	reu.mutation.ClearRoutines()
	return reu
}

// ClearExercises clears the "exercises" edge to the Exercise entity.
func (reu *RoutineExerciseUpdate) ClearExercises() *RoutineExerciseUpdate {
	reu.mutation.ClearExercises()
	return reu
}

// ClearUsers clears the "users" edge to the User entity.
func (reu *RoutineExerciseUpdate) ClearUsers() *RoutineExerciseUpdate {
	reu.mutation.ClearUsers()
	return reu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (reu *RoutineExerciseUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, reu.sqlSave, reu.mutation, reu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (reu *RoutineExerciseUpdate) SaveX(ctx context.Context) int {
	affected, err := reu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (reu *RoutineExerciseUpdate) Exec(ctx context.Context) error {
	_, err := reu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (reu *RoutineExerciseUpdate) ExecX(ctx context.Context) {
	if err := reu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (reu *RoutineExerciseUpdate) check() error {
	if _, ok := reu.mutation.RoutinesID(); reu.mutation.RoutinesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoutineExercise.routines"`)
	}
	if _, ok := reu.mutation.ExercisesID(); reu.mutation.ExercisesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoutineExercise.exercises"`)
	}
	return nil
}

func (reu *RoutineExerciseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := reu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(routineexercise.Table, routineexercise.Columns, sqlgraph.NewFieldSpec(routineexercise.FieldID, field.TypeString))
	if ps := reu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := reu.mutation.RestTimer(); ok {
		_spec.SetField(routineexercise.FieldRestTimer, field.TypeInt, value)
	}
	if value, ok := reu.mutation.AddedRestTimer(); ok {
		_spec.AddField(routineexercise.FieldRestTimer, field.TypeInt, value)
	}
	if reu.mutation.RestTimerCleared() {
		_spec.ClearField(routineexercise.FieldRestTimer, field.TypeInt)
	}
	if value, ok := reu.mutation.Sets(); ok {
		_spec.SetField(routineexercise.FieldSets, field.TypeJSON, value)
	}
	if reu.mutation.RoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   routineexercise.RoutinesTable,
			Columns: []string{routineexercise.RoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routine.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := reu.mutation.RoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   routineexercise.RoutinesTable,
			Columns: []string{routineexercise.RoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routine.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if reu.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   routineexercise.ExercisesTable,
			Columns: []string{routineexercise.ExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := reu.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   routineexercise.ExercisesTable,
			Columns: []string{routineexercise.ExercisesColumn},
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
	if reu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineexercise.UsersTable,
			Columns: []string{routineexercise.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := reu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineexercise.UsersTable,
			Columns: []string{routineexercise.UsersColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, reu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routineexercise.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	reu.mutation.done = true
	return n, nil
}

// RoutineExerciseUpdateOne is the builder for updating a single RoutineExercise entity.
type RoutineExerciseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoutineExerciseMutation
}

// SetRestTimer sets the "rest_timer" field.
func (reuo *RoutineExerciseUpdateOne) SetRestTimer(i int) *RoutineExerciseUpdateOne {
	reuo.mutation.ResetRestTimer()
	reuo.mutation.SetRestTimer(i)
	return reuo
}

// SetNillableRestTimer sets the "rest_timer" field if the given value is not nil.
func (reuo *RoutineExerciseUpdateOne) SetNillableRestTimer(i *int) *RoutineExerciseUpdateOne {
	if i != nil {
		reuo.SetRestTimer(*i)
	}
	return reuo
}

// AddRestTimer adds i to the "rest_timer" field.
func (reuo *RoutineExerciseUpdateOne) AddRestTimer(i int) *RoutineExerciseUpdateOne {
	reuo.mutation.AddRestTimer(i)
	return reuo
}

// ClearRestTimer clears the value of the "rest_timer" field.
func (reuo *RoutineExerciseUpdateOne) ClearRestTimer() *RoutineExerciseUpdateOne {
	reuo.mutation.ClearRestTimer()
	return reuo
}

// SetSets sets the "sets" field.
func (reuo *RoutineExerciseUpdateOne) SetSets(s *schematype.Sets) *RoutineExerciseUpdateOne {
	reuo.mutation.SetSets(s)
	return reuo
}

// SetRoutineID sets the "routine_id" field.
func (reuo *RoutineExerciseUpdateOne) SetRoutineID(s string) *RoutineExerciseUpdateOne {
	reuo.mutation.SetRoutineID(s)
	return reuo
}

// SetExerciseID sets the "exercise_id" field.
func (reuo *RoutineExerciseUpdateOne) SetExerciseID(s string) *RoutineExerciseUpdateOne {
	reuo.mutation.SetExerciseID(s)
	return reuo
}

// SetUserID sets the "user_id" field.
func (reuo *RoutineExerciseUpdateOne) SetUserID(s string) *RoutineExerciseUpdateOne {
	reuo.mutation.SetUserID(s)
	return reuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (reuo *RoutineExerciseUpdateOne) SetNillableUserID(s *string) *RoutineExerciseUpdateOne {
	if s != nil {
		reuo.SetUserID(*s)
	}
	return reuo
}

// ClearUserID clears the value of the "user_id" field.
func (reuo *RoutineExerciseUpdateOne) ClearUserID() *RoutineExerciseUpdateOne {
	reuo.mutation.ClearUserID()
	return reuo
}

// SetRoutinesID sets the "routines" edge to the Routine entity by ID.
func (reuo *RoutineExerciseUpdateOne) SetRoutinesID(id string) *RoutineExerciseUpdateOne {
	reuo.mutation.SetRoutinesID(id)
	return reuo
}

// SetRoutines sets the "routines" edge to the Routine entity.
func (reuo *RoutineExerciseUpdateOne) SetRoutines(r *Routine) *RoutineExerciseUpdateOne {
	return reuo.SetRoutinesID(r.ID)
}

// SetExercisesID sets the "exercises" edge to the Exercise entity by ID.
func (reuo *RoutineExerciseUpdateOne) SetExercisesID(id string) *RoutineExerciseUpdateOne {
	reuo.mutation.SetExercisesID(id)
	return reuo
}

// SetExercises sets the "exercises" edge to the Exercise entity.
func (reuo *RoutineExerciseUpdateOne) SetExercises(e *Exercise) *RoutineExerciseUpdateOne {
	return reuo.SetExercisesID(e.ID)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (reuo *RoutineExerciseUpdateOne) SetUsersID(id string) *RoutineExerciseUpdateOne {
	reuo.mutation.SetUsersID(id)
	return reuo
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (reuo *RoutineExerciseUpdateOne) SetNillableUsersID(id *string) *RoutineExerciseUpdateOne {
	if id != nil {
		reuo = reuo.SetUsersID(*id)
	}
	return reuo
}

// SetUsers sets the "users" edge to the User entity.
func (reuo *RoutineExerciseUpdateOne) SetUsers(u *User) *RoutineExerciseUpdateOne {
	return reuo.SetUsersID(u.ID)
}

// Mutation returns the RoutineExerciseMutation object of the builder.
func (reuo *RoutineExerciseUpdateOne) Mutation() *RoutineExerciseMutation {
	return reuo.mutation
}

// ClearRoutines clears the "routines" edge to the Routine entity.
func (reuo *RoutineExerciseUpdateOne) ClearRoutines() *RoutineExerciseUpdateOne {
	reuo.mutation.ClearRoutines()
	return reuo
}

// ClearExercises clears the "exercises" edge to the Exercise entity.
func (reuo *RoutineExerciseUpdateOne) ClearExercises() *RoutineExerciseUpdateOne {
	reuo.mutation.ClearExercises()
	return reuo
}

// ClearUsers clears the "users" edge to the User entity.
func (reuo *RoutineExerciseUpdateOne) ClearUsers() *RoutineExerciseUpdateOne {
	reuo.mutation.ClearUsers()
	return reuo
}

// Where appends a list predicates to the RoutineExerciseUpdate builder.
func (reuo *RoutineExerciseUpdateOne) Where(ps ...predicate.RoutineExercise) *RoutineExerciseUpdateOne {
	reuo.mutation.Where(ps...)
	return reuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (reuo *RoutineExerciseUpdateOne) Select(field string, fields ...string) *RoutineExerciseUpdateOne {
	reuo.fields = append([]string{field}, fields...)
	return reuo
}

// Save executes the query and returns the updated RoutineExercise entity.
func (reuo *RoutineExerciseUpdateOne) Save(ctx context.Context) (*RoutineExercise, error) {
	return withHooks(ctx, reuo.sqlSave, reuo.mutation, reuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (reuo *RoutineExerciseUpdateOne) SaveX(ctx context.Context) *RoutineExercise {
	node, err := reuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (reuo *RoutineExerciseUpdateOne) Exec(ctx context.Context) error {
	_, err := reuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (reuo *RoutineExerciseUpdateOne) ExecX(ctx context.Context) {
	if err := reuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (reuo *RoutineExerciseUpdateOne) check() error {
	if _, ok := reuo.mutation.RoutinesID(); reuo.mutation.RoutinesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoutineExercise.routines"`)
	}
	if _, ok := reuo.mutation.ExercisesID(); reuo.mutation.ExercisesCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoutineExercise.exercises"`)
	}
	return nil
}

func (reuo *RoutineExerciseUpdateOne) sqlSave(ctx context.Context) (_node *RoutineExercise, err error) {
	if err := reuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(routineexercise.Table, routineexercise.Columns, sqlgraph.NewFieldSpec(routineexercise.FieldID, field.TypeString))
	id, ok := reuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RoutineExercise.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := reuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, routineexercise.FieldID)
		for _, f := range fields {
			if !routineexercise.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != routineexercise.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := reuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := reuo.mutation.RestTimer(); ok {
		_spec.SetField(routineexercise.FieldRestTimer, field.TypeInt, value)
	}
	if value, ok := reuo.mutation.AddedRestTimer(); ok {
		_spec.AddField(routineexercise.FieldRestTimer, field.TypeInt, value)
	}
	if reuo.mutation.RestTimerCleared() {
		_spec.ClearField(routineexercise.FieldRestTimer, field.TypeInt)
	}
	if value, ok := reuo.mutation.Sets(); ok {
		_spec.SetField(routineexercise.FieldSets, field.TypeJSON, value)
	}
	if reuo.mutation.RoutinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   routineexercise.RoutinesTable,
			Columns: []string{routineexercise.RoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routine.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := reuo.mutation.RoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   routineexercise.RoutinesTable,
			Columns: []string{routineexercise.RoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routine.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if reuo.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   routineexercise.ExercisesTable,
			Columns: []string{routineexercise.ExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := reuo.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   routineexercise.ExercisesTable,
			Columns: []string{routineexercise.ExercisesColumn},
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
	if reuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineexercise.UsersTable,
			Columns: []string{routineexercise.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := reuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routineexercise.UsersTable,
			Columns: []string{routineexercise.UsersColumn},
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
	_node = &RoutineExercise{config: reuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, reuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routineexercise.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	reuo.mutation.done = true
	return _node, nil
}
