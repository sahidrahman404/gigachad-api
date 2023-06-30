// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/ent/user"
)

// RoutineUpdate is the builder for updating Routine entities.
type RoutineUpdate struct {
	config
	hooks    []Hook
	mutation *RoutineMutation
}

// Where appends a list predicates to the RoutineUpdate builder.
func (ru *RoutineUpdate) Where(ps ...predicate.Routine) *RoutineUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RoutineUpdate) SetName(s string) *RoutineUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetUserID sets the "user_id" field.
func (ru *RoutineUpdate) SetUserID(s string) *RoutineUpdate {
	ru.mutation.SetUserID(s)
	return ru
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ru *RoutineUpdate) SetNillableUserID(s *string) *RoutineUpdate {
	if s != nil {
		ru.SetUserID(*s)
	}
	return ru
}

// ClearUserID clears the value of the "user_id" field.
func (ru *RoutineUpdate) ClearUserID() *RoutineUpdate {
	ru.mutation.ClearUserID()
	return ru
}

// AddRoutineExerciseIDs adds the "routine_exercises" edge to the RoutineExercise entity by IDs.
func (ru *RoutineUpdate) AddRoutineExerciseIDs(ids ...string) *RoutineUpdate {
	ru.mutation.AddRoutineExerciseIDs(ids...)
	return ru
}

// AddRoutineExercises adds the "routine_exercises" edges to the RoutineExercise entity.
func (ru *RoutineUpdate) AddRoutineExercises(r ...*RoutineExercise) *RoutineUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddRoutineExerciseIDs(ids...)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (ru *RoutineUpdate) SetUsersID(id string) *RoutineUpdate {
	ru.mutation.SetUsersID(id)
	return ru
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (ru *RoutineUpdate) SetNillableUsersID(id *string) *RoutineUpdate {
	if id != nil {
		ru = ru.SetUsersID(*id)
	}
	return ru
}

// SetUsers sets the "users" edge to the User entity.
func (ru *RoutineUpdate) SetUsers(u *User) *RoutineUpdate {
	return ru.SetUsersID(u.ID)
}

// Mutation returns the RoutineMutation object of the builder.
func (ru *RoutineUpdate) Mutation() *RoutineMutation {
	return ru.mutation
}

// ClearRoutineExercises clears all "routine_exercises" edges to the RoutineExercise entity.
func (ru *RoutineUpdate) ClearRoutineExercises() *RoutineUpdate {
	ru.mutation.ClearRoutineExercises()
	return ru
}

// RemoveRoutineExerciseIDs removes the "routine_exercises" edge to RoutineExercise entities by IDs.
func (ru *RoutineUpdate) RemoveRoutineExerciseIDs(ids ...string) *RoutineUpdate {
	ru.mutation.RemoveRoutineExerciseIDs(ids...)
	return ru
}

// RemoveRoutineExercises removes "routine_exercises" edges to RoutineExercise entities.
func (ru *RoutineUpdate) RemoveRoutineExercises(r ...*RoutineExercise) *RoutineUpdate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveRoutineExerciseIDs(ids...)
}

// ClearUsers clears the "users" edge to the User entity.
func (ru *RoutineUpdate) ClearUsers() *RoutineUpdate {
	ru.mutation.ClearUsers()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RoutineUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RoutineUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RoutineUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RoutineUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RoutineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(routine.Table, routine.Columns, sqlgraph.NewFieldSpec(routine.FieldID, field.TypeString))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(routine.FieldName, field.TypeString, value)
	}
	if ru.mutation.RoutineExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   routine.RoutineExercisesTable,
			Columns: []string{routine.RoutineExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineexercise.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedRoutineExercisesIDs(); len(nodes) > 0 && !ru.mutation.RoutineExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   routine.RoutineExercisesTable,
			Columns: []string{routine.RoutineExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineexercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RoutineExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   routine.RoutineExercisesTable,
			Columns: []string{routine.RoutineExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineexercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routine.UsersTable,
			Columns: []string{routine.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routine.UsersTable,
			Columns: []string{routine.UsersColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RoutineUpdateOne is the builder for updating a single Routine entity.
type RoutineUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoutineMutation
}

// SetName sets the "name" field.
func (ruo *RoutineUpdateOne) SetName(s string) *RoutineUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetUserID sets the "user_id" field.
func (ruo *RoutineUpdateOne) SetUserID(s string) *RoutineUpdateOne {
	ruo.mutation.SetUserID(s)
	return ruo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ruo *RoutineUpdateOne) SetNillableUserID(s *string) *RoutineUpdateOne {
	if s != nil {
		ruo.SetUserID(*s)
	}
	return ruo
}

// ClearUserID clears the value of the "user_id" field.
func (ruo *RoutineUpdateOne) ClearUserID() *RoutineUpdateOne {
	ruo.mutation.ClearUserID()
	return ruo
}

// AddRoutineExerciseIDs adds the "routine_exercises" edge to the RoutineExercise entity by IDs.
func (ruo *RoutineUpdateOne) AddRoutineExerciseIDs(ids ...string) *RoutineUpdateOne {
	ruo.mutation.AddRoutineExerciseIDs(ids...)
	return ruo
}

// AddRoutineExercises adds the "routine_exercises" edges to the RoutineExercise entity.
func (ruo *RoutineUpdateOne) AddRoutineExercises(r ...*RoutineExercise) *RoutineUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddRoutineExerciseIDs(ids...)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (ruo *RoutineUpdateOne) SetUsersID(id string) *RoutineUpdateOne {
	ruo.mutation.SetUsersID(id)
	return ruo
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (ruo *RoutineUpdateOne) SetNillableUsersID(id *string) *RoutineUpdateOne {
	if id != nil {
		ruo = ruo.SetUsersID(*id)
	}
	return ruo
}

// SetUsers sets the "users" edge to the User entity.
func (ruo *RoutineUpdateOne) SetUsers(u *User) *RoutineUpdateOne {
	return ruo.SetUsersID(u.ID)
}

// Mutation returns the RoutineMutation object of the builder.
func (ruo *RoutineUpdateOne) Mutation() *RoutineMutation {
	return ruo.mutation
}

// ClearRoutineExercises clears all "routine_exercises" edges to the RoutineExercise entity.
func (ruo *RoutineUpdateOne) ClearRoutineExercises() *RoutineUpdateOne {
	ruo.mutation.ClearRoutineExercises()
	return ruo
}

// RemoveRoutineExerciseIDs removes the "routine_exercises" edge to RoutineExercise entities by IDs.
func (ruo *RoutineUpdateOne) RemoveRoutineExerciseIDs(ids ...string) *RoutineUpdateOne {
	ruo.mutation.RemoveRoutineExerciseIDs(ids...)
	return ruo
}

// RemoveRoutineExercises removes "routine_exercises" edges to RoutineExercise entities.
func (ruo *RoutineUpdateOne) RemoveRoutineExercises(r ...*RoutineExercise) *RoutineUpdateOne {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveRoutineExerciseIDs(ids...)
}

// ClearUsers clears the "users" edge to the User entity.
func (ruo *RoutineUpdateOne) ClearUsers() *RoutineUpdateOne {
	ruo.mutation.ClearUsers()
	return ruo
}

// Where appends a list predicates to the RoutineUpdate builder.
func (ruo *RoutineUpdateOne) Where(ps ...predicate.Routine) *RoutineUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RoutineUpdateOne) Select(field string, fields ...string) *RoutineUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Routine entity.
func (ruo *RoutineUpdateOne) Save(ctx context.Context) (*Routine, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RoutineUpdateOne) SaveX(ctx context.Context) *Routine {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RoutineUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RoutineUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RoutineUpdateOne) sqlSave(ctx context.Context) (_node *Routine, err error) {
	_spec := sqlgraph.NewUpdateSpec(routine.Table, routine.Columns, sqlgraph.NewFieldSpec(routine.FieldID, field.TypeString))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Routine.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, routine.FieldID)
		for _, f := range fields {
			if !routine.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != routine.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(routine.FieldName, field.TypeString, value)
	}
	if ruo.mutation.RoutineExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   routine.RoutineExercisesTable,
			Columns: []string{routine.RoutineExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineexercise.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedRoutineExercisesIDs(); len(nodes) > 0 && !ruo.mutation.RoutineExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   routine.RoutineExercisesTable,
			Columns: []string{routine.RoutineExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineexercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RoutineExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   routine.RoutineExercisesTable,
			Columns: []string{routine.RoutineExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineexercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routine.UsersTable,
			Columns: []string{routine.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   routine.UsersTable,
			Columns: []string{routine.UsersColumn},
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
	_node = &Routine{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{routine.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
