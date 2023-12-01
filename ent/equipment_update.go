// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/equipment"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
)

// EquipmentUpdate is the builder for updating Equipment entities.
type EquipmentUpdate struct {
	config
	hooks    []Hook
	mutation *EquipmentMutation
}

// Where appends a list predicates to the EquipmentUpdate builder.
func (eu *EquipmentUpdate) Where(ps ...predicate.Equipment) *EquipmentUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EquipmentUpdate) SetName(s string) *EquipmentUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (eu *EquipmentUpdate) SetNillableName(s *string) *EquipmentUpdate {
	if s != nil {
		eu.SetName(*s)
	}
	return eu
}

// SetImage sets the "image" field.
func (eu *EquipmentUpdate) SetImage(s schematype.Image) *EquipmentUpdate {
	eu.mutation.SetImage(s)
	return eu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (eu *EquipmentUpdate) SetNillableImage(s *schematype.Image) *EquipmentUpdate {
	if s != nil {
		eu.SetImage(*s)
	}
	return eu
}

// ClearImage clears the value of the "image" field.
func (eu *EquipmentUpdate) ClearImage() *EquipmentUpdate {
	eu.mutation.ClearImage()
	return eu
}

// AddExerciseIDs adds the "exercises" edge to the Exercise entity by IDs.
func (eu *EquipmentUpdate) AddExerciseIDs(ids ...pksuid.ID) *EquipmentUpdate {
	eu.mutation.AddExerciseIDs(ids...)
	return eu
}

// AddExercises adds the "exercises" edges to the Exercise entity.
func (eu *EquipmentUpdate) AddExercises(e ...*Exercise) *EquipmentUpdate {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eu.AddExerciseIDs(ids...)
}

// Mutation returns the EquipmentMutation object of the builder.
func (eu *EquipmentUpdate) Mutation() *EquipmentMutation {
	return eu.mutation
}

// ClearExercises clears all "exercises" edges to the Exercise entity.
func (eu *EquipmentUpdate) ClearExercises() *EquipmentUpdate {
	eu.mutation.ClearExercises()
	return eu
}

// RemoveExerciseIDs removes the "exercises" edge to Exercise entities by IDs.
func (eu *EquipmentUpdate) RemoveExerciseIDs(ids ...pksuid.ID) *EquipmentUpdate {
	eu.mutation.RemoveExerciseIDs(ids...)
	return eu
}

// RemoveExercises removes "exercises" edges to Exercise entities.
func (eu *EquipmentUpdate) RemoveExercises(e ...*Exercise) *EquipmentUpdate {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return eu.RemoveExerciseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EquipmentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EquipmentUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EquipmentUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EquipmentUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EquipmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(equipment.Table, equipment.Columns, sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeString))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.SetField(equipment.FieldName, field.TypeString, value)
	}
	if value, ok := eu.mutation.Image(); ok {
		_spec.SetField(equipment.FieldImage, field.TypeJSON, value)
	}
	if eu.mutation.ImageCleared() {
		_spec.ClearField(equipment.FieldImage, field.TypeJSON)
	}
	if eu.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   equipment.ExercisesTable,
			Columns: equipment.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedExercisesIDs(); len(nodes) > 0 && !eu.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   equipment.ExercisesTable,
			Columns: equipment.ExercisesPrimaryKey,
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
	if nodes := eu.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   equipment.ExercisesTable,
			Columns: equipment.ExercisesPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EquipmentUpdateOne is the builder for updating a single Equipment entity.
type EquipmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EquipmentMutation
}

// SetName sets the "name" field.
func (euo *EquipmentUpdateOne) SetName(s string) *EquipmentUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (euo *EquipmentUpdateOne) SetNillableName(s *string) *EquipmentUpdateOne {
	if s != nil {
		euo.SetName(*s)
	}
	return euo
}

// SetImage sets the "image" field.
func (euo *EquipmentUpdateOne) SetImage(s schematype.Image) *EquipmentUpdateOne {
	euo.mutation.SetImage(s)
	return euo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (euo *EquipmentUpdateOne) SetNillableImage(s *schematype.Image) *EquipmentUpdateOne {
	if s != nil {
		euo.SetImage(*s)
	}
	return euo
}

// ClearImage clears the value of the "image" field.
func (euo *EquipmentUpdateOne) ClearImage() *EquipmentUpdateOne {
	euo.mutation.ClearImage()
	return euo
}

// AddExerciseIDs adds the "exercises" edge to the Exercise entity by IDs.
func (euo *EquipmentUpdateOne) AddExerciseIDs(ids ...pksuid.ID) *EquipmentUpdateOne {
	euo.mutation.AddExerciseIDs(ids...)
	return euo
}

// AddExercises adds the "exercises" edges to the Exercise entity.
func (euo *EquipmentUpdateOne) AddExercises(e ...*Exercise) *EquipmentUpdateOne {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return euo.AddExerciseIDs(ids...)
}

// Mutation returns the EquipmentMutation object of the builder.
func (euo *EquipmentUpdateOne) Mutation() *EquipmentMutation {
	return euo.mutation
}

// ClearExercises clears all "exercises" edges to the Exercise entity.
func (euo *EquipmentUpdateOne) ClearExercises() *EquipmentUpdateOne {
	euo.mutation.ClearExercises()
	return euo
}

// RemoveExerciseIDs removes the "exercises" edge to Exercise entities by IDs.
func (euo *EquipmentUpdateOne) RemoveExerciseIDs(ids ...pksuid.ID) *EquipmentUpdateOne {
	euo.mutation.RemoveExerciseIDs(ids...)
	return euo
}

// RemoveExercises removes "exercises" edges to Exercise entities.
func (euo *EquipmentUpdateOne) RemoveExercises(e ...*Exercise) *EquipmentUpdateOne {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return euo.RemoveExerciseIDs(ids...)
}

// Where appends a list predicates to the EquipmentUpdate builder.
func (euo *EquipmentUpdateOne) Where(ps ...predicate.Equipment) *EquipmentUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EquipmentUpdateOne) Select(field string, fields ...string) *EquipmentUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Equipment entity.
func (euo *EquipmentUpdateOne) Save(ctx context.Context) (*Equipment, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EquipmentUpdateOne) SaveX(ctx context.Context) *Equipment {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EquipmentUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EquipmentUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EquipmentUpdateOne) sqlSave(ctx context.Context) (_node *Equipment, err error) {
	_spec := sqlgraph.NewUpdateSpec(equipment.Table, equipment.Columns, sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeString))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Equipment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipment.FieldID)
		for _, f := range fields {
			if !equipment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != equipment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.SetField(equipment.FieldName, field.TypeString, value)
	}
	if value, ok := euo.mutation.Image(); ok {
		_spec.SetField(equipment.FieldImage, field.TypeJSON, value)
	}
	if euo.mutation.ImageCleared() {
		_spec.ClearField(equipment.FieldImage, field.TypeJSON)
	}
	if euo.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   equipment.ExercisesTable,
			Columns: equipment.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedExercisesIDs(); len(nodes) > 0 && !euo.mutation.ExercisesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   equipment.ExercisesTable,
			Columns: equipment.ExercisesPrimaryKey,
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
	if nodes := euo.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   equipment.ExercisesTable,
			Columns: equipment.ExercisesPrimaryKey,
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
	_node = &Equipment{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
