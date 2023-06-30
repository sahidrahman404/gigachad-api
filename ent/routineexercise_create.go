// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/ent/schema"
	"github.com/sahidrahman404/gigachad-api/ent/user"
)

// RoutineExerciseCreate is the builder for creating a RoutineExercise entity.
type RoutineExerciseCreate struct {
	config
	mutation *RoutineExerciseMutation
	hooks    []Hook
}

// SetRestTimer sets the "rest_timer" field.
func (rec *RoutineExerciseCreate) SetRestTimer(i int) *RoutineExerciseCreate {
	rec.mutation.SetRestTimer(i)
	return rec
}

// SetNillableRestTimer sets the "rest_timer" field if the given value is not nil.
func (rec *RoutineExerciseCreate) SetNillableRestTimer(i *int) *RoutineExerciseCreate {
	if i != nil {
		rec.SetRestTimer(*i)
	}
	return rec
}

// SetSet sets the "set" field.
func (rec *RoutineExerciseCreate) SetSet(s *[]schema.Set) *RoutineExerciseCreate {
	rec.mutation.SetSet(s)
	return rec
}

// SetRoutineID sets the "routine_id" field.
func (rec *RoutineExerciseCreate) SetRoutineID(s string) *RoutineExerciseCreate {
	rec.mutation.SetRoutineID(s)
	return rec
}

// SetNillableRoutineID sets the "routine_id" field if the given value is not nil.
func (rec *RoutineExerciseCreate) SetNillableRoutineID(s *string) *RoutineExerciseCreate {
	if s != nil {
		rec.SetRoutineID(*s)
	}
	return rec
}

// SetExerciseID sets the "exercise_id" field.
func (rec *RoutineExerciseCreate) SetExerciseID(s string) *RoutineExerciseCreate {
	rec.mutation.SetExerciseID(s)
	return rec
}

// SetNillableExerciseID sets the "exercise_id" field if the given value is not nil.
func (rec *RoutineExerciseCreate) SetNillableExerciseID(s *string) *RoutineExerciseCreate {
	if s != nil {
		rec.SetExerciseID(*s)
	}
	return rec
}

// SetUserID sets the "user_id" field.
func (rec *RoutineExerciseCreate) SetUserID(s string) *RoutineExerciseCreate {
	rec.mutation.SetUserID(s)
	return rec
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (rec *RoutineExerciseCreate) SetNillableUserID(s *string) *RoutineExerciseCreate {
	if s != nil {
		rec.SetUserID(*s)
	}
	return rec
}

// SetID sets the "id" field.
func (rec *RoutineExerciseCreate) SetID(s string) *RoutineExerciseCreate {
	rec.mutation.SetID(s)
	return rec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rec *RoutineExerciseCreate) SetNillableID(s *string) *RoutineExerciseCreate {
	if s != nil {
		rec.SetID(*s)
	}
	return rec
}

// SetRoutinesID sets the "routines" edge to the Routine entity by ID.
func (rec *RoutineExerciseCreate) SetRoutinesID(id string) *RoutineExerciseCreate {
	rec.mutation.SetRoutinesID(id)
	return rec
}

// SetNillableRoutinesID sets the "routines" edge to the Routine entity by ID if the given value is not nil.
func (rec *RoutineExerciseCreate) SetNillableRoutinesID(id *string) *RoutineExerciseCreate {
	if id != nil {
		rec = rec.SetRoutinesID(*id)
	}
	return rec
}

// SetRoutines sets the "routines" edge to the Routine entity.
func (rec *RoutineExerciseCreate) SetRoutines(r *Routine) *RoutineExerciseCreate {
	return rec.SetRoutinesID(r.ID)
}

// SetExercisesID sets the "exercises" edge to the Exercise entity by ID.
func (rec *RoutineExerciseCreate) SetExercisesID(id string) *RoutineExerciseCreate {
	rec.mutation.SetExercisesID(id)
	return rec
}

// SetNillableExercisesID sets the "exercises" edge to the Exercise entity by ID if the given value is not nil.
func (rec *RoutineExerciseCreate) SetNillableExercisesID(id *string) *RoutineExerciseCreate {
	if id != nil {
		rec = rec.SetExercisesID(*id)
	}
	return rec
}

// SetExercises sets the "exercises" edge to the Exercise entity.
func (rec *RoutineExerciseCreate) SetExercises(e *Exercise) *RoutineExerciseCreate {
	return rec.SetExercisesID(e.ID)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (rec *RoutineExerciseCreate) SetUsersID(id string) *RoutineExerciseCreate {
	rec.mutation.SetUsersID(id)
	return rec
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (rec *RoutineExerciseCreate) SetNillableUsersID(id *string) *RoutineExerciseCreate {
	if id != nil {
		rec = rec.SetUsersID(*id)
	}
	return rec
}

// SetUsers sets the "users" edge to the User entity.
func (rec *RoutineExerciseCreate) SetUsers(u *User) *RoutineExerciseCreate {
	return rec.SetUsersID(u.ID)
}

// Mutation returns the RoutineExerciseMutation object of the builder.
func (rec *RoutineExerciseCreate) Mutation() *RoutineExerciseMutation {
	return rec.mutation
}

// Save creates the RoutineExercise in the database.
func (rec *RoutineExerciseCreate) Save(ctx context.Context) (*RoutineExercise, error) {
	rec.defaults()
	return withHooks(ctx, rec.sqlSave, rec.mutation, rec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rec *RoutineExerciseCreate) SaveX(ctx context.Context) *RoutineExercise {
	v, err := rec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rec *RoutineExerciseCreate) Exec(ctx context.Context) error {
	_, err := rec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rec *RoutineExerciseCreate) ExecX(ctx context.Context) {
	if err := rec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rec *RoutineExerciseCreate) defaults() {
	if _, ok := rec.mutation.ID(); !ok {
		v := routineexercise.DefaultID()
		rec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rec *RoutineExerciseCreate) check() error {
	if _, ok := rec.mutation.Set(); !ok {
		return &ValidationError{Name: "set", err: errors.New(`ent: missing required field "RoutineExercise.set"`)}
	}
	return nil
}

func (rec *RoutineExerciseCreate) sqlSave(ctx context.Context) (*RoutineExercise, error) {
	if err := rec.check(); err != nil {
		return nil, err
	}
	_node, _spec := rec.createSpec()
	if err := sqlgraph.CreateNode(ctx, rec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected RoutineExercise.ID type: %T", _spec.ID.Value)
		}
	}
	rec.mutation.id = &_node.ID
	rec.mutation.done = true
	return _node, nil
}

func (rec *RoutineExerciseCreate) createSpec() (*RoutineExercise, *sqlgraph.CreateSpec) {
	var (
		_node = &RoutineExercise{config: rec.config}
		_spec = sqlgraph.NewCreateSpec(routineexercise.Table, sqlgraph.NewFieldSpec(routineexercise.FieldID, field.TypeString))
	)
	if id, ok := rec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rec.mutation.RestTimer(); ok {
		_spec.SetField(routineexercise.FieldRestTimer, field.TypeInt, value)
		_node.RestTimer = value
	}
	if value, ok := rec.mutation.Set(); ok {
		_spec.SetField(routineexercise.FieldSet, field.TypeJSON, value)
		_node.Set = value
	}
	if nodes := rec.mutation.RoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
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
		_node.RoutineID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rec.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
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
		_node.ExerciseID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rec.mutation.UsersIDs(); len(nodes) > 0 {
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoutineExerciseCreateBulk is the builder for creating many RoutineExercise entities in bulk.
type RoutineExerciseCreateBulk struct {
	config
	builders []*RoutineExerciseCreate
}

// Save creates the RoutineExercise entities in the database.
func (recb *RoutineExerciseCreateBulk) Save(ctx context.Context) ([]*RoutineExercise, error) {
	specs := make([]*sqlgraph.CreateSpec, len(recb.builders))
	nodes := make([]*RoutineExercise, len(recb.builders))
	mutators := make([]Mutator, len(recb.builders))
	for i := range recb.builders {
		func(i int, root context.Context) {
			builder := recb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoutineExerciseMutation)
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
					_, err = mutators[i+1].Mutate(root, recb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, recb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, recb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (recb *RoutineExerciseCreateBulk) SaveX(ctx context.Context) []*RoutineExercise {
	v, err := recb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (recb *RoutineExerciseCreateBulk) Exec(ctx context.Context) error {
	_, err := recb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (recb *RoutineExerciseCreateBulk) ExecX(ctx context.Context) {
	if err := recb.Exec(ctx); err != nil {
		panic(err)
	}
}
