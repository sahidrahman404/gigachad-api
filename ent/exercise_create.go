// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/equipment"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/exercisetype"
	"github.com/sahidrahman404/gigachad-api/ent/musclesgroup"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/ent/workoutlog"
)

// ExerciseCreate is the builder for creating a Exercise entity.
type ExerciseCreate struct {
	config
	mutation *ExerciseMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ec *ExerciseCreate) SetName(s string) *ExerciseCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetImage sets the "image" field.
func (ec *ExerciseCreate) SetImage(s string) *ExerciseCreate {
	ec.mutation.SetImage(s)
	return ec
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (ec *ExerciseCreate) SetNillableImage(s *string) *ExerciseCreate {
	if s != nil {
		ec.SetImage(*s)
	}
	return ec
}

// SetHowTo sets the "how_to" field.
func (ec *ExerciseCreate) SetHowTo(s string) *ExerciseCreate {
	ec.mutation.SetHowTo(s)
	return ec
}

// SetNillableHowTo sets the "how_to" field if the given value is not nil.
func (ec *ExerciseCreate) SetNillableHowTo(s *string) *ExerciseCreate {
	if s != nil {
		ec.SetHowTo(*s)
	}
	return ec
}

// SetEquipmentID sets the "equipment_id" field.
func (ec *ExerciseCreate) SetEquipmentID(s string) *ExerciseCreate {
	ec.mutation.SetEquipmentID(s)
	return ec
}

// SetNillableEquipmentID sets the "equipment_id" field if the given value is not nil.
func (ec *ExerciseCreate) SetNillableEquipmentID(s *string) *ExerciseCreate {
	if s != nil {
		ec.SetEquipmentID(*s)
	}
	return ec
}

// SetMusclesGroupID sets the "muscles_group_id" field.
func (ec *ExerciseCreate) SetMusclesGroupID(s string) *ExerciseCreate {
	ec.mutation.SetMusclesGroupID(s)
	return ec
}

// SetNillableMusclesGroupID sets the "muscles_group_id" field if the given value is not nil.
func (ec *ExerciseCreate) SetNillableMusclesGroupID(s *string) *ExerciseCreate {
	if s != nil {
		ec.SetMusclesGroupID(*s)
	}
	return ec
}

// SetExerciseTypeID sets the "exercise_type_id" field.
func (ec *ExerciseCreate) SetExerciseTypeID(s string) *ExerciseCreate {
	ec.mutation.SetExerciseTypeID(s)
	return ec
}

// SetNillableExerciseTypeID sets the "exercise_type_id" field if the given value is not nil.
func (ec *ExerciseCreate) SetNillableExerciseTypeID(s *string) *ExerciseCreate {
	if s != nil {
		ec.SetExerciseTypeID(*s)
	}
	return ec
}

// SetUserID sets the "user_id" field.
func (ec *ExerciseCreate) SetUserID(s string) *ExerciseCreate {
	ec.mutation.SetUserID(s)
	return ec
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ec *ExerciseCreate) SetNillableUserID(s *string) *ExerciseCreate {
	if s != nil {
		ec.SetUserID(*s)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *ExerciseCreate) SetID(s string) *ExerciseCreate {
	ec.mutation.SetID(s)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *ExerciseCreate) SetNillableID(s *string) *ExerciseCreate {
	if s != nil {
		ec.SetID(*s)
	}
	return ec
}

// AddRoutineExerciseIDs adds the "routine_exercises" edge to the RoutineExercise entity by IDs.
func (ec *ExerciseCreate) AddRoutineExerciseIDs(ids ...string) *ExerciseCreate {
	ec.mutation.AddRoutineExerciseIDs(ids...)
	return ec
}

// AddRoutineExercises adds the "routine_exercises" edges to the RoutineExercise entity.
func (ec *ExerciseCreate) AddRoutineExercises(r ...*RoutineExercise) *ExerciseCreate {
	ids := make([]string, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ec.AddRoutineExerciseIDs(ids...)
}

// AddWorkoutLogIDs adds the "workout_logs" edge to the WorkoutLog entity by IDs.
func (ec *ExerciseCreate) AddWorkoutLogIDs(ids ...string) *ExerciseCreate {
	ec.mutation.AddWorkoutLogIDs(ids...)
	return ec
}

// AddWorkoutLogs adds the "workout_logs" edges to the WorkoutLog entity.
func (ec *ExerciseCreate) AddWorkoutLogs(w ...*WorkoutLog) *ExerciseCreate {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ec.AddWorkoutLogIDs(ids...)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (ec *ExerciseCreate) SetUsersID(id string) *ExerciseCreate {
	ec.mutation.SetUsersID(id)
	return ec
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (ec *ExerciseCreate) SetNillableUsersID(id *string) *ExerciseCreate {
	if id != nil {
		ec = ec.SetUsersID(*id)
	}
	return ec
}

// SetUsers sets the "users" edge to the User entity.
func (ec *ExerciseCreate) SetUsers(u *User) *ExerciseCreate {
	return ec.SetUsersID(u.ID)
}

// SetEquipmentsID sets the "equipments" edge to the Equipment entity by ID.
func (ec *ExerciseCreate) SetEquipmentsID(id string) *ExerciseCreate {
	ec.mutation.SetEquipmentsID(id)
	return ec
}

// SetNillableEquipmentsID sets the "equipments" edge to the Equipment entity by ID if the given value is not nil.
func (ec *ExerciseCreate) SetNillableEquipmentsID(id *string) *ExerciseCreate {
	if id != nil {
		ec = ec.SetEquipmentsID(*id)
	}
	return ec
}

// SetEquipments sets the "equipments" edge to the Equipment entity.
func (ec *ExerciseCreate) SetEquipments(e *Equipment) *ExerciseCreate {
	return ec.SetEquipmentsID(e.ID)
}

// SetMusclesGroupsID sets the "muscles_groups" edge to the MusclesGroup entity by ID.
func (ec *ExerciseCreate) SetMusclesGroupsID(id string) *ExerciseCreate {
	ec.mutation.SetMusclesGroupsID(id)
	return ec
}

// SetNillableMusclesGroupsID sets the "muscles_groups" edge to the MusclesGroup entity by ID if the given value is not nil.
func (ec *ExerciseCreate) SetNillableMusclesGroupsID(id *string) *ExerciseCreate {
	if id != nil {
		ec = ec.SetMusclesGroupsID(*id)
	}
	return ec
}

// SetMusclesGroups sets the "muscles_groups" edge to the MusclesGroup entity.
func (ec *ExerciseCreate) SetMusclesGroups(m *MusclesGroup) *ExerciseCreate {
	return ec.SetMusclesGroupsID(m.ID)
}

// SetExerciseTypesID sets the "exercise_types" edge to the ExerciseType entity by ID.
func (ec *ExerciseCreate) SetExerciseTypesID(id string) *ExerciseCreate {
	ec.mutation.SetExerciseTypesID(id)
	return ec
}

// SetNillableExerciseTypesID sets the "exercise_types" edge to the ExerciseType entity by ID if the given value is not nil.
func (ec *ExerciseCreate) SetNillableExerciseTypesID(id *string) *ExerciseCreate {
	if id != nil {
		ec = ec.SetExerciseTypesID(*id)
	}
	return ec
}

// SetExerciseTypes sets the "exercise_types" edge to the ExerciseType entity.
func (ec *ExerciseCreate) SetExerciseTypes(e *ExerciseType) *ExerciseCreate {
	return ec.SetExerciseTypesID(e.ID)
}

// Mutation returns the ExerciseMutation object of the builder.
func (ec *ExerciseCreate) Mutation() *ExerciseMutation {
	return ec.mutation
}

// Save creates the Exercise in the database.
func (ec *ExerciseCreate) Save(ctx context.Context) (*Exercise, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *ExerciseCreate) SaveX(ctx context.Context) *Exercise {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *ExerciseCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *ExerciseCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *ExerciseCreate) defaults() {
	if _, ok := ec.mutation.ID(); !ok {
		v := exercise.DefaultID()
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *ExerciseCreate) check() error {
	if _, ok := ec.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Exercise.name"`)}
	}
	return nil
}

func (ec *ExerciseCreate) sqlSave(ctx context.Context) (*Exercise, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Exercise.ID type: %T", _spec.ID.Value)
		}
	}
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *ExerciseCreate) createSpec() (*Exercise, *sqlgraph.CreateSpec) {
	var (
		_node = &Exercise{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(exercise.Table, sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString))
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ec.mutation.Name(); ok {
		_spec.SetField(exercise.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ec.mutation.Image(); ok {
		_spec.SetField(exercise.FieldImage, field.TypeString, value)
		_node.Image = &value
	}
	if value, ok := ec.mutation.HowTo(); ok {
		_spec.SetField(exercise.FieldHowTo, field.TypeString, value)
		_node.HowTo = &value
	}
	if value, ok := ec.mutation.EquipmentID(); ok {
		_spec.SetField(exercise.FieldEquipmentID, field.TypeString, value)
		_node.EquipmentID = value
	}
	if nodes := ec.mutation.RoutineExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exercise.RoutineExercisesTable,
			Columns: []string{exercise.RoutineExercisesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routineexercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.WorkoutLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   exercise.WorkoutLogsTable,
			Columns: []string{exercise.WorkoutLogsColumn},
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
	if nodes := ec.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   exercise.UsersTable,
			Columns: []string{exercise.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.EquipmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   exercise.EquipmentsTable,
			Columns: []string{exercise.EquipmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MusclesGroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.MusclesGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   exercise.MusclesGroupsTable,
			Columns: []string{exercise.MusclesGroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musclesgroup.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MusclesGroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ExerciseTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   exercise.ExerciseTypesTable,
			Columns: []string{exercise.ExerciseTypesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercisetype.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ExerciseTypeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ExerciseCreateBulk is the builder for creating many Exercise entities in bulk.
type ExerciseCreateBulk struct {
	config
	builders []*ExerciseCreate
}

// Save creates the Exercise entities in the database.
func (ecb *ExerciseCreateBulk) Save(ctx context.Context) ([]*Exercise, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Exercise, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExerciseMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *ExerciseCreateBulk) SaveX(ctx context.Context) []*Exercise {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *ExerciseCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *ExerciseCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}