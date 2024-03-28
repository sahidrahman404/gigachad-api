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
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
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

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RoutineUpdate) SetNillableName(s *string) *RoutineUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetReminderID sets the "reminder_id" field.
func (ru *RoutineUpdate) SetReminderID(s string) *RoutineUpdate {
	ru.mutation.SetReminderID(s)
	return ru
}

// SetNillableReminderID sets the "reminder_id" field if the given value is not nil.
func (ru *RoutineUpdate) SetNillableReminderID(s *string) *RoutineUpdate {
	if s != nil {
		ru.SetReminderID(*s)
	}
	return ru
}

// ClearReminderID clears the value of the "reminder_id" field.
func (ru *RoutineUpdate) ClearReminderID() *RoutineUpdate {
	ru.mutation.ClearReminderID()
	return ru
}

// SetUserID sets the "user_id" field.
func (ru *RoutineUpdate) SetUserID(pk pksuid.ID) *RoutineUpdate {
	ru.mutation.SetUserID(pk)
	return ru
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ru *RoutineUpdate) SetNillableUserID(pk *pksuid.ID) *RoutineUpdate {
	if pk != nil {
		ru.SetUserID(*pk)
	}
	return ru
}

// AddExerciseIDs adds the "exercises" edge to the Exercise entity by IDs.
func (ru *RoutineUpdate) AddExerciseIDs(ids ...pksuid.ID) *RoutineUpdate {
	ru.mutation.AddExerciseIDs(ids...)
	return ru
}

// AddExercises adds the "exercises" edges to the Exercise entity.
func (ru *RoutineUpdate) AddExercises(e ...*Exercise) *RoutineUpdate {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ru.AddExerciseIDs(ids...)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (ru *RoutineUpdate) SetUsersID(id pksuid.ID) *RoutineUpdate {
	ru.mutation.SetUsersID(id)
	return ru
}

// SetUsers sets the "users" edge to the User entity.
func (ru *RoutineUpdate) SetUsers(u *User) *RoutineUpdate {
	return ru.SetUsersID(u.ID)
}

// AddRoutineExerciseIDs adds the "routine_exercises" edge to the RoutineExercise entity by IDs.
func (ru *RoutineUpdate) AddRoutineExerciseIDs(ids ...pksuid.ID) *RoutineUpdate {
	ru.mutation.AddRoutineExerciseIDs(ids...)
	return ru
}

// AddRoutineExercises adds the "routine_exercises" edges to the RoutineExercise entity.
func (ru *RoutineUpdate) AddRoutineExercises(r ...*RoutineExercise) *RoutineUpdate {
	ids := make([]pksuid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddRoutineExerciseIDs(ids...)
}

// Mutation returns the RoutineMutation object of the builder.
func (ru *RoutineUpdate) Mutation() *RoutineMutation {
	return ru.mutation
}

// ClearExercises clears all "exercises" edges to the Exercise entity.
func (ru *RoutineUpdate) ClearExercises() *RoutineUpdate {
	ru.mutation.ClearExercises()
	return ru
}

// RemoveExerciseIDs removes the "exercises" edge to Exercise entities by IDs.
func (ru *RoutineUpdate) RemoveExerciseIDs(ids ...pksuid.ID) *RoutineUpdate {
	ru.mutation.RemoveExerciseIDs(ids...)
	return ru
}

// RemoveExercises removes "exercises" edges to Exercise entities.
func (ru *RoutineUpdate) RemoveExercises(e ...*Exercise) *RoutineUpdate {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ru.RemoveExerciseIDs(ids...)
}

// ClearUsers clears the "users" edge to the User entity.
func (ru *RoutineUpdate) ClearUsers() *RoutineUpdate {
	ru.mutation.ClearUsers()
	return ru
}

// ClearRoutineExercises clears all "routine_exercises" edges to the RoutineExercise entity.
func (ru *RoutineUpdate) ClearRoutineExercises() *RoutineUpdate {
	ru.mutation.ClearRoutineExercises()
	return ru
}

// RemoveRoutineExerciseIDs removes the "routine_exercises" edge to RoutineExercise entities by IDs.
func (ru *RoutineUpdate) RemoveRoutineExerciseIDs(ids ...pksuid.ID) *RoutineUpdate {
	ru.mutation.RemoveRoutineExerciseIDs(ids...)
	return ru
}

// RemoveRoutineExercises removes "routine_exercises" edges to RoutineExercise entities.
func (ru *RoutineUpdate) RemoveRoutineExercises(r ...*RoutineExercise) *RoutineUpdate {
	ids := make([]pksuid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveRoutineExerciseIDs(ids...)
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

// check runs all checks and user-defined validators on the builder.
func (ru *RoutineUpdate) check() error {
	if _, ok := ru.mutation.UsersID(); ru.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Routine.users"`)
	}
	return nil
}

func (ru *RoutineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
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
	if value, ok := ru.mutation.ReminderID(); ok {
		_spec.SetField(routine.FieldReminderID, field.TypeString, value)
	}
	if ru.mutation.ReminderIDCleared() {
		_spec.ClearField(routine.FieldReminderID, field.TypeString)
	}
	if ru.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   routine.ExercisesTable,
			Columns: routine.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		createE := &RoutineExerciseCreate{config: ru.config, mutation: newRoutineExerciseMutation(ru.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedExercisesIDs(); len(nodes) > 0 && !ru.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   routine.ExercisesTable,
			Columns: routine.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &RoutineExerciseCreate{config: ru.config, mutation: newRoutineExerciseMutation(ru.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   routine.ExercisesTable,
			Columns: routine.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &RoutineExerciseCreate{config: ru.config, mutation: newRoutineExerciseMutation(ru.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
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
	if ru.mutation.RoutineExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
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
			Inverse: true,
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
			Inverse: true,
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

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RoutineUpdateOne) SetNillableName(s *string) *RoutineUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetReminderID sets the "reminder_id" field.
func (ruo *RoutineUpdateOne) SetReminderID(s string) *RoutineUpdateOne {
	ruo.mutation.SetReminderID(s)
	return ruo
}

// SetNillableReminderID sets the "reminder_id" field if the given value is not nil.
func (ruo *RoutineUpdateOne) SetNillableReminderID(s *string) *RoutineUpdateOne {
	if s != nil {
		ruo.SetReminderID(*s)
	}
	return ruo
}

// ClearReminderID clears the value of the "reminder_id" field.
func (ruo *RoutineUpdateOne) ClearReminderID() *RoutineUpdateOne {
	ruo.mutation.ClearReminderID()
	return ruo
}

// SetUserID sets the "user_id" field.
func (ruo *RoutineUpdateOne) SetUserID(pk pksuid.ID) *RoutineUpdateOne {
	ruo.mutation.SetUserID(pk)
	return ruo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ruo *RoutineUpdateOne) SetNillableUserID(pk *pksuid.ID) *RoutineUpdateOne {
	if pk != nil {
		ruo.SetUserID(*pk)
	}
	return ruo
}

// AddExerciseIDs adds the "exercises" edge to the Exercise entity by IDs.
func (ruo *RoutineUpdateOne) AddExerciseIDs(ids ...pksuid.ID) *RoutineUpdateOne {
	ruo.mutation.AddExerciseIDs(ids...)
	return ruo
}

// AddExercises adds the "exercises" edges to the Exercise entity.
func (ruo *RoutineUpdateOne) AddExercises(e ...*Exercise) *RoutineUpdateOne {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ruo.AddExerciseIDs(ids...)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (ruo *RoutineUpdateOne) SetUsersID(id pksuid.ID) *RoutineUpdateOne {
	ruo.mutation.SetUsersID(id)
	return ruo
}

// SetUsers sets the "users" edge to the User entity.
func (ruo *RoutineUpdateOne) SetUsers(u *User) *RoutineUpdateOne {
	return ruo.SetUsersID(u.ID)
}

// AddRoutineExerciseIDs adds the "routine_exercises" edge to the RoutineExercise entity by IDs.
func (ruo *RoutineUpdateOne) AddRoutineExerciseIDs(ids ...pksuid.ID) *RoutineUpdateOne {
	ruo.mutation.AddRoutineExerciseIDs(ids...)
	return ruo
}

// AddRoutineExercises adds the "routine_exercises" edges to the RoutineExercise entity.
func (ruo *RoutineUpdateOne) AddRoutineExercises(r ...*RoutineExercise) *RoutineUpdateOne {
	ids := make([]pksuid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddRoutineExerciseIDs(ids...)
}

// Mutation returns the RoutineMutation object of the builder.
func (ruo *RoutineUpdateOne) Mutation() *RoutineMutation {
	return ruo.mutation
}

// ClearExercises clears all "exercises" edges to the Exercise entity.
func (ruo *RoutineUpdateOne) ClearExercises() *RoutineUpdateOne {
	ruo.mutation.ClearExercises()
	return ruo
}

// RemoveExerciseIDs removes the "exercises" edge to Exercise entities by IDs.
func (ruo *RoutineUpdateOne) RemoveExerciseIDs(ids ...pksuid.ID) *RoutineUpdateOne {
	ruo.mutation.RemoveExerciseIDs(ids...)
	return ruo
}

// RemoveExercises removes "exercises" edges to Exercise entities.
func (ruo *RoutineUpdateOne) RemoveExercises(e ...*Exercise) *RoutineUpdateOne {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ruo.RemoveExerciseIDs(ids...)
}

// ClearUsers clears the "users" edge to the User entity.
func (ruo *RoutineUpdateOne) ClearUsers() *RoutineUpdateOne {
	ruo.mutation.ClearUsers()
	return ruo
}

// ClearRoutineExercises clears all "routine_exercises" edges to the RoutineExercise entity.
func (ruo *RoutineUpdateOne) ClearRoutineExercises() *RoutineUpdateOne {
	ruo.mutation.ClearRoutineExercises()
	return ruo
}

// RemoveRoutineExerciseIDs removes the "routine_exercises" edge to RoutineExercise entities by IDs.
func (ruo *RoutineUpdateOne) RemoveRoutineExerciseIDs(ids ...pksuid.ID) *RoutineUpdateOne {
	ruo.mutation.RemoveRoutineExerciseIDs(ids...)
	return ruo
}

// RemoveRoutineExercises removes "routine_exercises" edges to RoutineExercise entities.
func (ruo *RoutineUpdateOne) RemoveRoutineExercises(r ...*RoutineExercise) *RoutineUpdateOne {
	ids := make([]pksuid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveRoutineExerciseIDs(ids...)
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

// check runs all checks and user-defined validators on the builder.
func (ruo *RoutineUpdateOne) check() error {
	if _, ok := ruo.mutation.UsersID(); ruo.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Routine.users"`)
	}
	return nil
}

func (ruo *RoutineUpdateOne) sqlSave(ctx context.Context) (_node *Routine, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
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
	if value, ok := ruo.mutation.ReminderID(); ok {
		_spec.SetField(routine.FieldReminderID, field.TypeString, value)
	}
	if ruo.mutation.ReminderIDCleared() {
		_spec.ClearField(routine.FieldReminderID, field.TypeString)
	}
	if ruo.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   routine.ExercisesTable,
			Columns: routine.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		createE := &RoutineExerciseCreate{config: ruo.config, mutation: newRoutineExerciseMutation(ruo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedExercisesIDs(); len(nodes) > 0 && !ruo.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   routine.ExercisesTable,
			Columns: routine.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &RoutineExerciseCreate{config: ruo.config, mutation: newRoutineExerciseMutation(ruo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   routine.ExercisesTable,
			Columns: routine.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &RoutineExerciseCreate{config: ruo.config, mutation: newRoutineExerciseMutation(ruo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
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
	if ruo.mutation.RoutineExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
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
			Inverse: true,
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
			Inverse: true,
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
