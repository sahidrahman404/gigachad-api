// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/ent/workout"
	"github.com/sahidrahman404/gigachad-api/ent/workoutlog"
)

// WorkoutCreate is the builder for creating a Workout entity.
type WorkoutCreate struct {
	config
	mutation *WorkoutMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (wc *WorkoutCreate) SetName(s string) *WorkoutCreate {
	wc.mutation.SetName(s)
	return wc
}

// SetVolume sets the "volume" field.
func (wc *WorkoutCreate) SetVolume(i int) *WorkoutCreate {
	wc.mutation.SetVolume(i)
	return wc
}

// SetReps sets the "reps" field.
func (wc *WorkoutCreate) SetReps(i int) *WorkoutCreate {
	wc.mutation.SetReps(i)
	return wc
}

// SetTime sets the "time" field.
func (wc *WorkoutCreate) SetTime(s string) *WorkoutCreate {
	wc.mutation.SetTime(s)
	return wc
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (wc *WorkoutCreate) SetNillableTime(s *string) *WorkoutCreate {
	if s != nil {
		wc.SetTime(*s)
	}
	return wc
}

// SetSets sets the "sets" field.
func (wc *WorkoutCreate) SetSets(i int) *WorkoutCreate {
	wc.mutation.SetSets(i)
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WorkoutCreate) SetCreatedAt(s string) *WorkoutCreate {
	wc.mutation.SetCreatedAt(s)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WorkoutCreate) SetNillableCreatedAt(s *string) *WorkoutCreate {
	if s != nil {
		wc.SetCreatedAt(*s)
	}
	return wc
}

// SetImage sets the "image" field.
func (wc *WorkoutCreate) SetImage(s string) *WorkoutCreate {
	wc.mutation.SetImage(s)
	return wc
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (wc *WorkoutCreate) SetNillableImage(s *string) *WorkoutCreate {
	if s != nil {
		wc.SetImage(*s)
	}
	return wc
}

// SetDescription sets the "description" field.
func (wc *WorkoutCreate) SetDescription(s string) *WorkoutCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetUserID sets the "user_id" field.
func (wc *WorkoutCreate) SetUserID(s string) *WorkoutCreate {
	wc.mutation.SetUserID(s)
	return wc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wc *WorkoutCreate) SetNillableUserID(s *string) *WorkoutCreate {
	if s != nil {
		wc.SetUserID(*s)
	}
	return wc
}

// SetID sets the "id" field.
func (wc *WorkoutCreate) SetID(s string) *WorkoutCreate {
	wc.mutation.SetID(s)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WorkoutCreate) SetNillableID(s *string) *WorkoutCreate {
	if s != nil {
		wc.SetID(*s)
	}
	return wc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (wc *WorkoutCreate) SetUsersID(id string) *WorkoutCreate {
	wc.mutation.SetUsersID(id)
	return wc
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (wc *WorkoutCreate) SetNillableUsersID(id *string) *WorkoutCreate {
	if id != nil {
		wc = wc.SetUsersID(*id)
	}
	return wc
}

// SetUsers sets the "users" edge to the User entity.
func (wc *WorkoutCreate) SetUsers(u *User) *WorkoutCreate {
	return wc.SetUsersID(u.ID)
}

// AddWorkoutLogIDs adds the "workout_logs" edge to the WorkoutLog entity by IDs.
func (wc *WorkoutCreate) AddWorkoutLogIDs(ids ...string) *WorkoutCreate {
	wc.mutation.AddWorkoutLogIDs(ids...)
	return wc
}

// AddWorkoutLogs adds the "workout_logs" edges to the WorkoutLog entity.
func (wc *WorkoutCreate) AddWorkoutLogs(w ...*WorkoutLog) *WorkoutCreate {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return wc.AddWorkoutLogIDs(ids...)
}

// Mutation returns the WorkoutMutation object of the builder.
func (wc *WorkoutCreate) Mutation() *WorkoutMutation {
	return wc.mutation
}

// Save creates the Workout in the database.
func (wc *WorkoutCreate) Save(ctx context.Context) (*Workout, error) {
	wc.defaults()
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkoutCreate) SaveX(ctx context.Context) *Workout {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorkoutCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorkoutCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WorkoutCreate) defaults() {
	if _, ok := wc.mutation.CreatedAt(); !ok {
		v := workout.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.ID(); !ok {
		v := workout.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkoutCreate) check() error {
	if _, ok := wc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Workout.name"`)}
	}
	if _, ok := wc.mutation.Volume(); !ok {
		return &ValidationError{Name: "volume", err: errors.New(`ent: missing required field "Workout.volume"`)}
	}
	if _, ok := wc.mutation.Reps(); !ok {
		return &ValidationError{Name: "reps", err: errors.New(`ent: missing required field "Workout.reps"`)}
	}
	if _, ok := wc.mutation.Sets(); !ok {
		return &ValidationError{Name: "sets", err: errors.New(`ent: missing required field "Workout.sets"`)}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Workout.created_at"`)}
	}
	if _, ok := wc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Workout.description"`)}
	}
	return nil
}

func (wc *WorkoutCreate) sqlSave(ctx context.Context) (*Workout, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Workout.ID type: %T", _spec.ID.Value)
		}
	}
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WorkoutCreate) createSpec() (*Workout, *sqlgraph.CreateSpec) {
	var (
		_node = &Workout{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(workout.Table, sqlgraph.NewFieldSpec(workout.FieldID, field.TypeString))
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wc.mutation.Name(); ok {
		_spec.SetField(workout.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wc.mutation.Volume(); ok {
		_spec.SetField(workout.FieldVolume, field.TypeInt, value)
		_node.Volume = value
	}
	if value, ok := wc.mutation.Reps(); ok {
		_spec.SetField(workout.FieldReps, field.TypeInt, value)
		_node.Reps = value
	}
	if value, ok := wc.mutation.Time(); ok {
		_spec.SetField(workout.FieldTime, field.TypeString, value)
		_node.Time = value
	}
	if value, ok := wc.mutation.Sets(); ok {
		_spec.SetField(workout.FieldSets, field.TypeInt, value)
		_node.Sets = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.SetField(workout.FieldCreatedAt, field.TypeString, value)
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.Image(); ok {
		_spec.SetField(workout.FieldImage, field.TypeString, value)
		_node.Image = &value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.SetField(workout.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := wc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workout.UsersTable,
			Columns: []string{workout.UsersColumn},
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
	if nodes := wc.mutation.WorkoutLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workout.WorkoutLogsTable,
			Columns: []string{workout.WorkoutLogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workoutlog.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkoutCreateBulk is the builder for creating many Workout entities in bulk.
type WorkoutCreateBulk struct {
	config
	builders []*WorkoutCreate
}

// Save creates the Workout entities in the database.
func (wcb *WorkoutCreateBulk) Save(ctx context.Context) ([]*Workout, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Workout, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkoutMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkoutCreateBulk) SaveX(ctx context.Context) []*Workout {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorkoutCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorkoutCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}