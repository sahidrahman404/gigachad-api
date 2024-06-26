// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/equipment"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/exercisetype"
	"github.com/sahidrahman404/gigachad-api/ent/musclesgroup"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/ent/workout"
	"github.com/sahidrahman404/gigachad-api/ent/workoutlog"
)

// ExerciseCreate is the builder for creating a Exercise entity.
type ExerciseCreate struct {
	config
	mutation *ExerciseMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (ec *ExerciseCreate) SetName(s string) *ExerciseCreate {
	ec.mutation.SetName(s)
	return ec
}

// SetImage sets the "image" field.
func (ec *ExerciseCreate) SetImage(s *schematype.Image) *ExerciseCreate {
	ec.mutation.SetImage(s)
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

// SetUserID sets the "user_id" field.
func (ec *ExerciseCreate) SetUserID(pk pksuid.ID) *ExerciseCreate {
	ec.mutation.SetUserID(pk)
	return ec
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ec *ExerciseCreate) SetNillableUserID(pk *pksuid.ID) *ExerciseCreate {
	if pk != nil {
		ec.SetUserID(*pk)
	}
	return ec
}

// SetID sets the "id" field.
func (ec *ExerciseCreate) SetID(pk pksuid.ID) *ExerciseCreate {
	ec.mutation.SetID(pk)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *ExerciseCreate) SetNillableID(pk *pksuid.ID) *ExerciseCreate {
	if pk != nil {
		ec.SetID(*pk)
	}
	return ec
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (ec *ExerciseCreate) SetUsersID(id pksuid.ID) *ExerciseCreate {
	ec.mutation.SetUsersID(id)
	return ec
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (ec *ExerciseCreate) SetNillableUsersID(id *pksuid.ID) *ExerciseCreate {
	if id != nil {
		ec = ec.SetUsersID(*id)
	}
	return ec
}

// SetUsers sets the "users" edge to the User entity.
func (ec *ExerciseCreate) SetUsers(u *User) *ExerciseCreate {
	return ec.SetUsersID(u.ID)
}

// AddEquipmentIDs adds the "equipment" edge to the Equipment entity by IDs.
func (ec *ExerciseCreate) AddEquipmentIDs(ids ...pksuid.ID) *ExerciseCreate {
	ec.mutation.AddEquipmentIDs(ids...)
	return ec
}

// AddEquipment adds the "equipment" edges to the Equipment entity.
func (ec *ExerciseCreate) AddEquipment(e ...*Equipment) *ExerciseCreate {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddEquipmentIDs(ids...)
}

// AddMusclesGroupIDs adds the "muscles_groups" edge to the MusclesGroup entity by IDs.
func (ec *ExerciseCreate) AddMusclesGroupIDs(ids ...pksuid.ID) *ExerciseCreate {
	ec.mutation.AddMusclesGroupIDs(ids...)
	return ec
}

// AddMusclesGroups adds the "muscles_groups" edges to the MusclesGroup entity.
func (ec *ExerciseCreate) AddMusclesGroups(m ...*MusclesGroup) *ExerciseCreate {
	ids := make([]pksuid.ID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ec.AddMusclesGroupIDs(ids...)
}

// AddExerciseTypeIDs adds the "exercise_types" edge to the ExerciseType entity by IDs.
func (ec *ExerciseCreate) AddExerciseTypeIDs(ids ...pksuid.ID) *ExerciseCreate {
	ec.mutation.AddExerciseTypeIDs(ids...)
	return ec
}

// AddExerciseTypes adds the "exercise_types" edges to the ExerciseType entity.
func (ec *ExerciseCreate) AddExerciseTypes(e ...*ExerciseType) *ExerciseCreate {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ec.AddExerciseTypeIDs(ids...)
}

// AddRoutineIDs adds the "routines" edge to the Routine entity by IDs.
func (ec *ExerciseCreate) AddRoutineIDs(ids ...pksuid.ID) *ExerciseCreate {
	ec.mutation.AddRoutineIDs(ids...)
	return ec
}

// AddRoutines adds the "routines" edges to the Routine entity.
func (ec *ExerciseCreate) AddRoutines(r ...*Routine) *ExerciseCreate {
	ids := make([]pksuid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ec.AddRoutineIDs(ids...)
}

// AddWorkoutIDs adds the "workouts" edge to the Workout entity by IDs.
func (ec *ExerciseCreate) AddWorkoutIDs(ids ...pksuid.ID) *ExerciseCreate {
	ec.mutation.AddWorkoutIDs(ids...)
	return ec
}

// AddWorkouts adds the "workouts" edges to the Workout entity.
func (ec *ExerciseCreate) AddWorkouts(w ...*Workout) *ExerciseCreate {
	ids := make([]pksuid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ec.AddWorkoutIDs(ids...)
}

// AddRoutineExerciseIDs adds the "routine_exercises" edge to the RoutineExercise entity by IDs.
func (ec *ExerciseCreate) AddRoutineExerciseIDs(ids ...pksuid.ID) *ExerciseCreate {
	ec.mutation.AddRoutineExerciseIDs(ids...)
	return ec
}

// AddRoutineExercises adds the "routine_exercises" edges to the RoutineExercise entity.
func (ec *ExerciseCreate) AddRoutineExercises(r ...*RoutineExercise) *ExerciseCreate {
	ids := make([]pksuid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ec.AddRoutineExerciseIDs(ids...)
}

// AddWorkoutLogIDs adds the "workout_logs" edge to the WorkoutLog entity by IDs.
func (ec *ExerciseCreate) AddWorkoutLogIDs(ids ...pksuid.ID) *ExerciseCreate {
	ec.mutation.AddWorkoutLogIDs(ids...)
	return ec
}

// AddWorkoutLogs adds the "workout_logs" edges to the WorkoutLog entity.
func (ec *ExerciseCreate) AddWorkoutLogs(w ...*WorkoutLog) *ExerciseCreate {
	ids := make([]pksuid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return ec.AddWorkoutLogIDs(ids...)
}

// Mutation returns the ExerciseMutation object of the builder.
func (ec *ExerciseCreate) Mutation() *ExerciseMutation {
	return ec.mutation
}

// Save creates the Exercise in the database.
func (ec *ExerciseCreate) Save(ctx context.Context) (*Exercise, error) {
	if err := ec.defaults(); err != nil {
		return nil, err
	}
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
func (ec *ExerciseCreate) defaults() error {
	if _, ok := ec.mutation.ID(); !ok {
		if exercise.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized exercise.DefaultID (forgotten import ent/runtime?)")
		}
		v := exercise.DefaultID()
		ec.mutation.SetID(v)
	}
	return nil
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
		if id, ok := _spec.ID.Value.(*pksuid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
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
	_spec.OnConflict = ec.conflict
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ec.mutation.Name(); ok {
		_spec.SetField(exercise.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ec.mutation.Image(); ok {
		_spec.SetField(exercise.FieldImage, field.TypeJSON, value)
		_node.Image = value
	}
	if value, ok := ec.mutation.HowTo(); ok {
		_spec.SetField(exercise.FieldHowTo, field.TypeString, value)
		_node.HowTo = &value
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
	if nodes := ec.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   exercise.EquipmentTable,
			Columns: exercise.EquipmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.MusclesGroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   exercise.MusclesGroupsTable,
			Columns: exercise.MusclesGroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(musclesgroup.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ExerciseTypesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   exercise.ExerciseTypesTable,
			Columns: exercise.ExerciseTypesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercisetype.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.RoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   exercise.RoutinesTable,
			Columns: exercise.RoutinesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routine.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &RoutineExerciseCreate{config: ec.config, mutation: newRoutineExerciseMutation(ec.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.WorkoutsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   exercise.WorkoutsTable,
			Columns: exercise.WorkoutsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workout.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &WorkoutLogCreate{config: ec.config, mutation: newWorkoutLogMutation(ec.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.RoutineExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
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
			Inverse: true,
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Exercise.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ExerciseUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ec *ExerciseCreate) OnConflict(opts ...sql.ConflictOption) *ExerciseUpsertOne {
	ec.conflict = opts
	return &ExerciseUpsertOne{
		create: ec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Exercise.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ec *ExerciseCreate) OnConflictColumns(columns ...string) *ExerciseUpsertOne {
	ec.conflict = append(ec.conflict, sql.ConflictColumns(columns...))
	return &ExerciseUpsertOne{
		create: ec,
	}
}

type (
	// ExerciseUpsertOne is the builder for "upsert"-ing
	//  one Exercise node.
	ExerciseUpsertOne struct {
		create *ExerciseCreate
	}

	// ExerciseUpsert is the "OnConflict" setter.
	ExerciseUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *ExerciseUpsert) SetName(v string) *ExerciseUpsert {
	u.Set(exercise.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ExerciseUpsert) UpdateName() *ExerciseUpsert {
	u.SetExcluded(exercise.FieldName)
	return u
}

// SetImage sets the "image" field.
func (u *ExerciseUpsert) SetImage(v *schematype.Image) *ExerciseUpsert {
	u.Set(exercise.FieldImage, v)
	return u
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *ExerciseUpsert) UpdateImage() *ExerciseUpsert {
	u.SetExcluded(exercise.FieldImage)
	return u
}

// ClearImage clears the value of the "image" field.
func (u *ExerciseUpsert) ClearImage() *ExerciseUpsert {
	u.SetNull(exercise.FieldImage)
	return u
}

// SetHowTo sets the "how_to" field.
func (u *ExerciseUpsert) SetHowTo(v string) *ExerciseUpsert {
	u.Set(exercise.FieldHowTo, v)
	return u
}

// UpdateHowTo sets the "how_to" field to the value that was provided on create.
func (u *ExerciseUpsert) UpdateHowTo() *ExerciseUpsert {
	u.SetExcluded(exercise.FieldHowTo)
	return u
}

// ClearHowTo clears the value of the "how_to" field.
func (u *ExerciseUpsert) ClearHowTo() *ExerciseUpsert {
	u.SetNull(exercise.FieldHowTo)
	return u
}

// SetUserID sets the "user_id" field.
func (u *ExerciseUpsert) SetUserID(v pksuid.ID) *ExerciseUpsert {
	u.Set(exercise.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ExerciseUpsert) UpdateUserID() *ExerciseUpsert {
	u.SetExcluded(exercise.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *ExerciseUpsert) ClearUserID() *ExerciseUpsert {
	u.SetNull(exercise.FieldUserID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Exercise.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(exercise.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ExerciseUpsertOne) UpdateNewValues() *ExerciseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(exercise.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Exercise.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ExerciseUpsertOne) Ignore() *ExerciseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ExerciseUpsertOne) DoNothing() *ExerciseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ExerciseCreate.OnConflict
// documentation for more info.
func (u *ExerciseUpsertOne) Update(set func(*ExerciseUpsert)) *ExerciseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ExerciseUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ExerciseUpsertOne) SetName(v string) *ExerciseUpsertOne {
	return u.Update(func(s *ExerciseUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ExerciseUpsertOne) UpdateName() *ExerciseUpsertOne {
	return u.Update(func(s *ExerciseUpsert) {
		s.UpdateName()
	})
}

// SetImage sets the "image" field.
func (u *ExerciseUpsertOne) SetImage(v *schematype.Image) *ExerciseUpsertOne {
	return u.Update(func(s *ExerciseUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *ExerciseUpsertOne) UpdateImage() *ExerciseUpsertOne {
	return u.Update(func(s *ExerciseUpsert) {
		s.UpdateImage()
	})
}

// ClearImage clears the value of the "image" field.
func (u *ExerciseUpsertOne) ClearImage() *ExerciseUpsertOne {
	return u.Update(func(s *ExerciseUpsert) {
		s.ClearImage()
	})
}

// SetHowTo sets the "how_to" field.
func (u *ExerciseUpsertOne) SetHowTo(v string) *ExerciseUpsertOne {
	return u.Update(func(s *ExerciseUpsert) {
		s.SetHowTo(v)
	})
}

// UpdateHowTo sets the "how_to" field to the value that was provided on create.
func (u *ExerciseUpsertOne) UpdateHowTo() *ExerciseUpsertOne {
	return u.Update(func(s *ExerciseUpsert) {
		s.UpdateHowTo()
	})
}

// ClearHowTo clears the value of the "how_to" field.
func (u *ExerciseUpsertOne) ClearHowTo() *ExerciseUpsertOne {
	return u.Update(func(s *ExerciseUpsert) {
		s.ClearHowTo()
	})
}

// SetUserID sets the "user_id" field.
func (u *ExerciseUpsertOne) SetUserID(v pksuid.ID) *ExerciseUpsertOne {
	return u.Update(func(s *ExerciseUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ExerciseUpsertOne) UpdateUserID() *ExerciseUpsertOne {
	return u.Update(func(s *ExerciseUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *ExerciseUpsertOne) ClearUserID() *ExerciseUpsertOne {
	return u.Update(func(s *ExerciseUpsert) {
		s.ClearUserID()
	})
}

// Exec executes the query.
func (u *ExerciseUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ExerciseCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ExerciseUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ExerciseUpsertOne) ID(ctx context.Context) (id pksuid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ExerciseUpsertOne.ID is not supported by MySQL driver. Use ExerciseUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ExerciseUpsertOne) IDX(ctx context.Context) pksuid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ExerciseCreateBulk is the builder for creating many Exercise entities in bulk.
type ExerciseCreateBulk struct {
	config
	err      error
	builders []*ExerciseCreate
	conflict []sql.ConflictOption
}

// Save creates the Exercise entities in the database.
func (ecb *ExerciseCreateBulk) Save(ctx context.Context) ([]*Exercise, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
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
					spec.OnConflict = ecb.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Exercise.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ExerciseUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ecb *ExerciseCreateBulk) OnConflict(opts ...sql.ConflictOption) *ExerciseUpsertBulk {
	ecb.conflict = opts
	return &ExerciseUpsertBulk{
		create: ecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Exercise.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ecb *ExerciseCreateBulk) OnConflictColumns(columns ...string) *ExerciseUpsertBulk {
	ecb.conflict = append(ecb.conflict, sql.ConflictColumns(columns...))
	return &ExerciseUpsertBulk{
		create: ecb,
	}
}

// ExerciseUpsertBulk is the builder for "upsert"-ing
// a bulk of Exercise nodes.
type ExerciseUpsertBulk struct {
	create *ExerciseCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Exercise.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(exercise.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ExerciseUpsertBulk) UpdateNewValues() *ExerciseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(exercise.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Exercise.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ExerciseUpsertBulk) Ignore() *ExerciseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ExerciseUpsertBulk) DoNothing() *ExerciseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ExerciseCreateBulk.OnConflict
// documentation for more info.
func (u *ExerciseUpsertBulk) Update(set func(*ExerciseUpsert)) *ExerciseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ExerciseUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ExerciseUpsertBulk) SetName(v string) *ExerciseUpsertBulk {
	return u.Update(func(s *ExerciseUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ExerciseUpsertBulk) UpdateName() *ExerciseUpsertBulk {
	return u.Update(func(s *ExerciseUpsert) {
		s.UpdateName()
	})
}

// SetImage sets the "image" field.
func (u *ExerciseUpsertBulk) SetImage(v *schematype.Image) *ExerciseUpsertBulk {
	return u.Update(func(s *ExerciseUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *ExerciseUpsertBulk) UpdateImage() *ExerciseUpsertBulk {
	return u.Update(func(s *ExerciseUpsert) {
		s.UpdateImage()
	})
}

// ClearImage clears the value of the "image" field.
func (u *ExerciseUpsertBulk) ClearImage() *ExerciseUpsertBulk {
	return u.Update(func(s *ExerciseUpsert) {
		s.ClearImage()
	})
}

// SetHowTo sets the "how_to" field.
func (u *ExerciseUpsertBulk) SetHowTo(v string) *ExerciseUpsertBulk {
	return u.Update(func(s *ExerciseUpsert) {
		s.SetHowTo(v)
	})
}

// UpdateHowTo sets the "how_to" field to the value that was provided on create.
func (u *ExerciseUpsertBulk) UpdateHowTo() *ExerciseUpsertBulk {
	return u.Update(func(s *ExerciseUpsert) {
		s.UpdateHowTo()
	})
}

// ClearHowTo clears the value of the "how_to" field.
func (u *ExerciseUpsertBulk) ClearHowTo() *ExerciseUpsertBulk {
	return u.Update(func(s *ExerciseUpsert) {
		s.ClearHowTo()
	})
}

// SetUserID sets the "user_id" field.
func (u *ExerciseUpsertBulk) SetUserID(v pksuid.ID) *ExerciseUpsertBulk {
	return u.Update(func(s *ExerciseUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ExerciseUpsertBulk) UpdateUserID() *ExerciseUpsertBulk {
	return u.Update(func(s *ExerciseUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *ExerciseUpsertBulk) ClearUserID() *ExerciseUpsertBulk {
	return u.Update(func(s *ExerciseUpsert) {
		s.ClearUserID()
	})
}

// Exec executes the query.
func (u *ExerciseUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ExerciseCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ExerciseCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ExerciseUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
