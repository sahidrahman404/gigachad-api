// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/ent/workout"
	"github.com/sahidrahman404/gigachad-api/ent/workoutlog"
)

// WorkoutLogCreate is the builder for creating a WorkoutLog entity.
type WorkoutLogCreate struct {
	config
	mutation *WorkoutLogMutation
	hooks    []Hook
}

// SetSets sets the "sets" field.
func (wlc *WorkoutLogCreate) SetSets(s []*schematype.Set) *WorkoutLogCreate {
	wlc.mutation.SetSets(s)
	return wlc
}

// SetCreatedAt sets the "created_at" field.
func (wlc *WorkoutLogCreate) SetCreatedAt(s string) *WorkoutLogCreate {
	wlc.mutation.SetCreatedAt(s)
	return wlc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wlc *WorkoutLogCreate) SetNillableCreatedAt(s *string) *WorkoutLogCreate {
	if s != nil {
		wlc.SetCreatedAt(*s)
	}
	return wlc
}

// SetUserID sets the "user_id" field.
func (wlc *WorkoutLogCreate) SetUserID(pk pksuid.ID) *WorkoutLogCreate {
	wlc.mutation.SetUserID(pk)
	return wlc
}

// SetID sets the "id" field.
func (wlc *WorkoutLogCreate) SetID(pk pksuid.ID) *WorkoutLogCreate {
	wlc.mutation.SetID(pk)
	return wlc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wlc *WorkoutLogCreate) SetNillableID(pk *pksuid.ID) *WorkoutLogCreate {
	if pk != nil {
		wlc.SetID(*pk)
	}
	return wlc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (wlc *WorkoutLogCreate) SetUsersID(id pksuid.ID) *WorkoutLogCreate {
	wlc.mutation.SetUsersID(id)
	return wlc
}

// SetUsers sets the "users" edge to the User entity.
func (wlc *WorkoutLogCreate) SetUsers(u *User) *WorkoutLogCreate {
	return wlc.SetUsersID(u.ID)
}

// SetExercisesID sets the "exercises" edge to the Exercise entity by ID.
func (wlc *WorkoutLogCreate) SetExercisesID(id pksuid.ID) *WorkoutLogCreate {
	wlc.mutation.SetExercisesID(id)
	return wlc
}

// SetNillableExercisesID sets the "exercises" edge to the Exercise entity by ID if the given value is not nil.
func (wlc *WorkoutLogCreate) SetNillableExercisesID(id *pksuid.ID) *WorkoutLogCreate {
	if id != nil {
		wlc = wlc.SetExercisesID(*id)
	}
	return wlc
}

// SetExercises sets the "exercises" edge to the Exercise entity.
func (wlc *WorkoutLogCreate) SetExercises(e *Exercise) *WorkoutLogCreate {
	return wlc.SetExercisesID(e.ID)
}

// SetWorkoutsID sets the "workouts" edge to the Workout entity by ID.
func (wlc *WorkoutLogCreate) SetWorkoutsID(id pksuid.ID) *WorkoutLogCreate {
	wlc.mutation.SetWorkoutsID(id)
	return wlc
}

// SetNillableWorkoutsID sets the "workouts" edge to the Workout entity by ID if the given value is not nil.
func (wlc *WorkoutLogCreate) SetNillableWorkoutsID(id *pksuid.ID) *WorkoutLogCreate {
	if id != nil {
		wlc = wlc.SetWorkoutsID(*id)
	}
	return wlc
}

// SetWorkouts sets the "workouts" edge to the Workout entity.
func (wlc *WorkoutLogCreate) SetWorkouts(w *Workout) *WorkoutLogCreate {
	return wlc.SetWorkoutsID(w.ID)
}

// Mutation returns the WorkoutLogMutation object of the builder.
func (wlc *WorkoutLogCreate) Mutation() *WorkoutLogMutation {
	return wlc.mutation
}

// Save creates the WorkoutLog in the database.
func (wlc *WorkoutLogCreate) Save(ctx context.Context) (*WorkoutLog, error) {
	wlc.defaults()
	return withHooks(ctx, wlc.sqlSave, wlc.mutation, wlc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wlc *WorkoutLogCreate) SaveX(ctx context.Context) *WorkoutLog {
	v, err := wlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wlc *WorkoutLogCreate) Exec(ctx context.Context) error {
	_, err := wlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wlc *WorkoutLogCreate) ExecX(ctx context.Context) {
	if err := wlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wlc *WorkoutLogCreate) defaults() {
	if _, ok := wlc.mutation.CreatedAt(); !ok {
		v := workoutlog.DefaultCreatedAt()
		wlc.mutation.SetCreatedAt(v)
	}
	if _, ok := wlc.mutation.ID(); !ok {
		v := workoutlog.DefaultID()
		wlc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wlc *WorkoutLogCreate) check() error {
	if _, ok := wlc.mutation.Sets(); !ok {
		return &ValidationError{Name: "sets", err: errors.New(`ent: missing required field "WorkoutLog.sets"`)}
	}
	if _, ok := wlc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "WorkoutLog.created_at"`)}
	}
	if _, ok := wlc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "WorkoutLog.user_id"`)}
	}
	if _, ok := wlc.mutation.UsersID(); !ok {
		return &ValidationError{Name: "users", err: errors.New(`ent: missing required edge "WorkoutLog.users"`)}
	}
	return nil
}

func (wlc *WorkoutLogCreate) sqlSave(ctx context.Context) (*WorkoutLog, error) {
	if err := wlc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wlc.driver, _spec); err != nil {
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
	wlc.mutation.id = &_node.ID
	wlc.mutation.done = true
	return _node, nil
}

func (wlc *WorkoutLogCreate) createSpec() (*WorkoutLog, *sqlgraph.CreateSpec) {
	var (
		_node = &WorkoutLog{config: wlc.config}
		_spec = sqlgraph.NewCreateSpec(workoutlog.Table, sqlgraph.NewFieldSpec(workoutlog.FieldID, field.TypeString))
	)
	if id, ok := wlc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wlc.mutation.Sets(); ok {
		_spec.SetField(workoutlog.FieldSets, field.TypeJSON, value)
		_node.Sets = value
	}
	if value, ok := wlc.mutation.CreatedAt(); ok {
		_spec.SetField(workoutlog.FieldCreatedAt, field.TypeString, value)
		_node.CreatedAt = value
	}
	if nodes := wlc.mutation.UsersIDs(); len(nodes) > 0 {
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wlc.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
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
		_node.exercise_workout_logs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wlc.mutation.WorkoutsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
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
		_node.workout_workout_logs = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WorkoutLogCreateBulk is the builder for creating many WorkoutLog entities in bulk.
type WorkoutLogCreateBulk struct {
	config
	err      error
	builders []*WorkoutLogCreate
}

// Save creates the WorkoutLog entities in the database.
func (wlcb *WorkoutLogCreateBulk) Save(ctx context.Context) ([]*WorkoutLog, error) {
	if wlcb.err != nil {
		return nil, wlcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wlcb.builders))
	nodes := make([]*WorkoutLog, len(wlcb.builders))
	mutators := make([]Mutator, len(wlcb.builders))
	for i := range wlcb.builders {
		func(i int, root context.Context) {
			builder := wlcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkoutLogMutation)
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
					_, err = mutators[i+1].Mutate(root, wlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wlcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wlcb *WorkoutLogCreateBulk) SaveX(ctx context.Context) []*WorkoutLog {
	v, err := wlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wlcb *WorkoutLogCreateBulk) Exec(ctx context.Context) error {
	_, err := wlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wlcb *WorkoutLogCreateBulk) ExecX(ctx context.Context) {
	if err := wlcb.Exec(ctx); err != nil {
		panic(err)
	}
}
