// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/exercisetype"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
)

// ExerciseTypeUpdate is the builder for updating ExerciseType entities.
type ExerciseTypeUpdate struct {
	config
	hooks    []Hook
	mutation *ExerciseTypeMutation
}

// Where appends a list predicates to the ExerciseTypeUpdate builder.
func (etu *ExerciseTypeUpdate) Where(ps ...predicate.ExerciseType) *ExerciseTypeUpdate {
	etu.mutation.Where(ps...)
	return etu
}

// SetName sets the "name" field.
func (etu *ExerciseTypeUpdate) SetName(s string) *ExerciseTypeUpdate {
	etu.mutation.SetName(s)
	return etu
}

// SetProperties sets the "properties" field.
func (etu *ExerciseTypeUpdate) SetProperties(s []string) *ExerciseTypeUpdate {
	etu.mutation.SetProperties(s)
	return etu
}

// AppendProperties appends s to the "properties" field.
func (etu *ExerciseTypeUpdate) AppendProperties(s []string) *ExerciseTypeUpdate {
	etu.mutation.AppendProperties(s)
	return etu
}

// SetDescription sets the "description" field.
func (etu *ExerciseTypeUpdate) SetDescription(s string) *ExerciseTypeUpdate {
	etu.mutation.SetDescription(s)
	return etu
}

// AddExerciseIDs adds the "exercises" edge to the Exercise entity by IDs.
func (etu *ExerciseTypeUpdate) AddExerciseIDs(ids ...string) *ExerciseTypeUpdate {
	etu.mutation.AddExerciseIDs(ids...)
	return etu
}

// AddExercises adds the "exercises" edges to the Exercise entity.
func (etu *ExerciseTypeUpdate) AddExercises(e ...*Exercise) *ExerciseTypeUpdate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return etu.AddExerciseIDs(ids...)
}

// Mutation returns the ExerciseTypeMutation object of the builder.
func (etu *ExerciseTypeUpdate) Mutation() *ExerciseTypeMutation {
	return etu.mutation
}

// ClearExercises clears all "exercises" edges to the Exercise entity.
func (etu *ExerciseTypeUpdate) ClearExercises() *ExerciseTypeUpdate {
	etu.mutation.ClearExercises()
	return etu
}

// RemoveExerciseIDs removes the "exercises" edge to Exercise entities by IDs.
func (etu *ExerciseTypeUpdate) RemoveExerciseIDs(ids ...string) *ExerciseTypeUpdate {
	etu.mutation.RemoveExerciseIDs(ids...)
	return etu
}

// RemoveExercises removes "exercises" edges to Exercise entities.
func (etu *ExerciseTypeUpdate) RemoveExercises(e ...*Exercise) *ExerciseTypeUpdate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return etu.RemoveExerciseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (etu *ExerciseTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, etu.sqlSave, etu.mutation, etu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (etu *ExerciseTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := etu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (etu *ExerciseTypeUpdate) Exec(ctx context.Context) error {
	_, err := etu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etu *ExerciseTypeUpdate) ExecX(ctx context.Context) {
	if err := etu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (etu *ExerciseTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(exercisetype.Table, exercisetype.Columns, sqlgraph.NewFieldSpec(exercisetype.FieldID, field.TypeString))
	if ps := etu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etu.mutation.Name(); ok {
		_spec.SetField(exercisetype.FieldName, field.TypeString, value)
	}
	if value, ok := etu.mutation.Properties(); ok {
		_spec.SetField(exercisetype.FieldProperties, field.TypeJSON, value)
	}
	if value, ok := etu.mutation.AppendedProperties(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, exercisetype.FieldProperties, value)
		})
	}
	if value, ok := etu.mutation.Description(); ok {
		_spec.SetField(exercisetype.FieldDescription, field.TypeString, value)
	}
	if etu.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exercisetype.ExercisesTable,
			Columns: []string{exercisetype.ExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etu.mutation.RemovedExercisesIDs(); len(nodes) > 0 && !etu.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exercisetype.ExercisesTable,
			Columns: []string{exercisetype.ExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etu.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exercisetype.ExercisesTable,
			Columns: []string{exercisetype.ExercisesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, etu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{exercisetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	etu.mutation.done = true
	return n, nil
}

// ExerciseTypeUpdateOne is the builder for updating a single ExerciseType entity.
type ExerciseTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ExerciseTypeMutation
}

// SetName sets the "name" field.
func (etuo *ExerciseTypeUpdateOne) SetName(s string) *ExerciseTypeUpdateOne {
	etuo.mutation.SetName(s)
	return etuo
}

// SetProperties sets the "properties" field.
func (etuo *ExerciseTypeUpdateOne) SetProperties(s []string) *ExerciseTypeUpdateOne {
	etuo.mutation.SetProperties(s)
	return etuo
}

// AppendProperties appends s to the "properties" field.
func (etuo *ExerciseTypeUpdateOne) AppendProperties(s []string) *ExerciseTypeUpdateOne {
	etuo.mutation.AppendProperties(s)
	return etuo
}

// SetDescription sets the "description" field.
func (etuo *ExerciseTypeUpdateOne) SetDescription(s string) *ExerciseTypeUpdateOne {
	etuo.mutation.SetDescription(s)
	return etuo
}

// AddExerciseIDs adds the "exercises" edge to the Exercise entity by IDs.
func (etuo *ExerciseTypeUpdateOne) AddExerciseIDs(ids ...string) *ExerciseTypeUpdateOne {
	etuo.mutation.AddExerciseIDs(ids...)
	return etuo
}

// AddExercises adds the "exercises" edges to the Exercise entity.
func (etuo *ExerciseTypeUpdateOne) AddExercises(e ...*Exercise) *ExerciseTypeUpdateOne {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return etuo.AddExerciseIDs(ids...)
}

// Mutation returns the ExerciseTypeMutation object of the builder.
func (etuo *ExerciseTypeUpdateOne) Mutation() *ExerciseTypeMutation {
	return etuo.mutation
}

// ClearExercises clears all "exercises" edges to the Exercise entity.
func (etuo *ExerciseTypeUpdateOne) ClearExercises() *ExerciseTypeUpdateOne {
	etuo.mutation.ClearExercises()
	return etuo
}

// RemoveExerciseIDs removes the "exercises" edge to Exercise entities by IDs.
func (etuo *ExerciseTypeUpdateOne) RemoveExerciseIDs(ids ...string) *ExerciseTypeUpdateOne {
	etuo.mutation.RemoveExerciseIDs(ids...)
	return etuo
}

// RemoveExercises removes "exercises" edges to Exercise entities.
func (etuo *ExerciseTypeUpdateOne) RemoveExercises(e ...*Exercise) *ExerciseTypeUpdateOne {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return etuo.RemoveExerciseIDs(ids...)
}

// Where appends a list predicates to the ExerciseTypeUpdate builder.
func (etuo *ExerciseTypeUpdateOne) Where(ps ...predicate.ExerciseType) *ExerciseTypeUpdateOne {
	etuo.mutation.Where(ps...)
	return etuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (etuo *ExerciseTypeUpdateOne) Select(field string, fields ...string) *ExerciseTypeUpdateOne {
	etuo.fields = append([]string{field}, fields...)
	return etuo
}

// Save executes the query and returns the updated ExerciseType entity.
func (etuo *ExerciseTypeUpdateOne) Save(ctx context.Context) (*ExerciseType, error) {
	return withHooks(ctx, etuo.sqlSave, etuo.mutation, etuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (etuo *ExerciseTypeUpdateOne) SaveX(ctx context.Context) *ExerciseType {
	node, err := etuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (etuo *ExerciseTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := etuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etuo *ExerciseTypeUpdateOne) ExecX(ctx context.Context) {
	if err := etuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (etuo *ExerciseTypeUpdateOne) sqlSave(ctx context.Context) (_node *ExerciseType, err error) {
	_spec := sqlgraph.NewUpdateSpec(exercisetype.Table, exercisetype.Columns, sqlgraph.NewFieldSpec(exercisetype.FieldID, field.TypeString))
	id, ok := etuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ExerciseType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := etuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, exercisetype.FieldID)
		for _, f := range fields {
			if !exercisetype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != exercisetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := etuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := etuo.mutation.Name(); ok {
		_spec.SetField(exercisetype.FieldName, field.TypeString, value)
	}
	if value, ok := etuo.mutation.Properties(); ok {
		_spec.SetField(exercisetype.FieldProperties, field.TypeJSON, value)
	}
	if value, ok := etuo.mutation.AppendedProperties(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, exercisetype.FieldProperties, value)
		})
	}
	if value, ok := etuo.mutation.Description(); ok {
		_spec.SetField(exercisetype.FieldDescription, field.TypeString, value)
	}
	if etuo.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exercisetype.ExercisesTable,
			Columns: []string{exercisetype.ExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etuo.mutation.RemovedExercisesIDs(); len(nodes) > 0 && !etuo.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exercisetype.ExercisesTable,
			Columns: []string{exercisetype.ExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := etuo.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exercisetype.ExercisesTable,
			Columns: []string{exercisetype.ExercisesColumn},
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
	_node = &ExerciseType{config: etuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, etuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{exercisetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	etuo.mutation.done = true
	return _node, nil
}