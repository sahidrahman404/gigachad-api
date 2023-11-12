// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/exercisetype"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

// ExerciseTypeCreate is the builder for creating a ExerciseType entity.
type ExerciseTypeCreate struct {
	config
	mutation *ExerciseTypeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (etc *ExerciseTypeCreate) SetName(s string) *ExerciseTypeCreate {
	etc.mutation.SetName(s)
	return etc
}

// SetProperties sets the "properties" field.
func (etc *ExerciseTypeCreate) SetProperties(s []string) *ExerciseTypeCreate {
	etc.mutation.SetProperties(s)
	return etc
}

// SetDescription sets the "description" field.
func (etc *ExerciseTypeCreate) SetDescription(s string) *ExerciseTypeCreate {
	etc.mutation.SetDescription(s)
	return etc
}

// SetID sets the "id" field.
func (etc *ExerciseTypeCreate) SetID(pk pksuid.ID) *ExerciseTypeCreate {
	etc.mutation.SetID(pk)
	return etc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (etc *ExerciseTypeCreate) SetNillableID(pk *pksuid.ID) *ExerciseTypeCreate {
	if pk != nil {
		etc.SetID(*pk)
	}
	return etc
}

// AddExerciseIDs adds the "exercises" edge to the Exercise entity by IDs.
func (etc *ExerciseTypeCreate) AddExerciseIDs(ids ...pksuid.ID) *ExerciseTypeCreate {
	etc.mutation.AddExerciseIDs(ids...)
	return etc
}

// AddExercises adds the "exercises" edges to the Exercise entity.
func (etc *ExerciseTypeCreate) AddExercises(e ...*Exercise) *ExerciseTypeCreate {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return etc.AddExerciseIDs(ids...)
}

// Mutation returns the ExerciseTypeMutation object of the builder.
func (etc *ExerciseTypeCreate) Mutation() *ExerciseTypeMutation {
	return etc.mutation
}

// Save creates the ExerciseType in the database.
func (etc *ExerciseTypeCreate) Save(ctx context.Context) (*ExerciseType, error) {
	etc.defaults()
	return withHooks(ctx, etc.sqlSave, etc.mutation, etc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (etc *ExerciseTypeCreate) SaveX(ctx context.Context) *ExerciseType {
	v, err := etc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (etc *ExerciseTypeCreate) Exec(ctx context.Context) error {
	_, err := etc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etc *ExerciseTypeCreate) ExecX(ctx context.Context) {
	if err := etc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etc *ExerciseTypeCreate) defaults() {
	if _, ok := etc.mutation.ID(); !ok {
		v := exercisetype.DefaultID()
		etc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (etc *ExerciseTypeCreate) check() error {
	if _, ok := etc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ExerciseType.name"`)}
	}
	if _, ok := etc.mutation.Properties(); !ok {
		return &ValidationError{Name: "properties", err: errors.New(`ent: missing required field "ExerciseType.properties"`)}
	}
	if _, ok := etc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "ExerciseType.description"`)}
	}
	return nil
}

func (etc *ExerciseTypeCreate) sqlSave(ctx context.Context) (*ExerciseType, error) {
	if err := etc.check(); err != nil {
		return nil, err
	}
	_node, _spec := etc.createSpec()
	if err := sqlgraph.CreateNode(ctx, etc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*pksuid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	etc.mutation.id = &_node.ID
	etc.mutation.done = true
	return _node, nil
}

func (etc *ExerciseTypeCreate) createSpec() (*ExerciseType, *sqlgraph.CreateSpec) {
	var (
		_node = &ExerciseType{config: etc.config}
		_spec = sqlgraph.NewCreateSpec(exercisetype.Table, sqlgraph.NewFieldSpec(exercisetype.FieldID, field.TypeString))
	)
	if id, ok := etc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := etc.mutation.Name(); ok {
		_spec.SetField(exercisetype.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := etc.mutation.Properties(); ok {
		_spec.SetField(exercisetype.FieldProperties, field.TypeJSON, value)
		_node.Properties = value
	}
	if value, ok := etc.mutation.Description(); ok {
		_spec.SetField(exercisetype.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := etc.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   exercisetype.ExercisesTable,
			Columns: exercisetype.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ExerciseTypeCreateBulk is the builder for creating many ExerciseType entities in bulk.
type ExerciseTypeCreateBulk struct {
	config
	err      error
	builders []*ExerciseTypeCreate
}

// Save creates the ExerciseType entities in the database.
func (etcb *ExerciseTypeCreateBulk) Save(ctx context.Context) ([]*ExerciseType, error) {
	if etcb.err != nil {
		return nil, etcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(etcb.builders))
	nodes := make([]*ExerciseType, len(etcb.builders))
	mutators := make([]Mutator, len(etcb.builders))
	for i := range etcb.builders {
		func(i int, root context.Context) {
			builder := etcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExerciseTypeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, etcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, etcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, etcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (etcb *ExerciseTypeCreateBulk) SaveX(ctx context.Context) []*ExerciseType {
	v, err := etcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (etcb *ExerciseTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := etcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etcb *ExerciseTypeCreateBulk) ExecX(ctx context.Context) {
	if err := etcb.Exec(ctx); err != nil {
		panic(err)
	}
}
