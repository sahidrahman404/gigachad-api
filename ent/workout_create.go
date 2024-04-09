// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/ent/workout"
	"github.com/sahidrahman404/gigachad-api/ent/workoutlog"
)

// WorkoutCreate is the builder for creating a Workout entity.
type WorkoutCreate struct {
	config
	mutation *WorkoutMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (wc *WorkoutCreate) SetName(s string) *WorkoutCreate {
	wc.mutation.SetName(s)
	return wc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (wc *WorkoutCreate) SetNillableName(s *string) *WorkoutCreate {
	if s != nil {
		wc.SetName(*s)
	}
	return wc
}

// SetVolume sets the "volume" field.
func (wc *WorkoutCreate) SetVolume(f float64) *WorkoutCreate {
	wc.mutation.SetVolume(f)
	return wc
}

// SetDuration sets the "duration" field.
func (wc *WorkoutCreate) SetDuration(s string) *WorkoutCreate {
	wc.mutation.SetDuration(s)
	return wc
}

// SetSets sets the "sets" field.
func (wc *WorkoutCreate) SetSets(i int) *WorkoutCreate {
	wc.mutation.SetSets(i)
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WorkoutCreate) SetCreatedAt(t time.Time) *WorkoutCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WorkoutCreate) SetNillableCreatedAt(t *time.Time) *WorkoutCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetImage sets the "image" field.
func (wc *WorkoutCreate) SetImage(s *schematype.Image) *WorkoutCreate {
	wc.mutation.SetImage(s)
	return wc
}

// SetDescription sets the "description" field.
func (wc *WorkoutCreate) SetDescription(s string) *WorkoutCreate {
	wc.mutation.SetDescription(s)
	return wc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (wc *WorkoutCreate) SetNillableDescription(s *string) *WorkoutCreate {
	if s != nil {
		wc.SetDescription(*s)
	}
	return wc
}

// SetUserID sets the "user_id" field.
func (wc *WorkoutCreate) SetUserID(pk pksuid.ID) *WorkoutCreate {
	wc.mutation.SetUserID(pk)
	return wc
}

// SetID sets the "id" field.
func (wc *WorkoutCreate) SetID(pk pksuid.ID) *WorkoutCreate {
	wc.mutation.SetID(pk)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WorkoutCreate) SetNillableID(pk *pksuid.ID) *WorkoutCreate {
	if pk != nil {
		wc.SetID(*pk)
	}
	return wc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (wc *WorkoutCreate) SetUsersID(id pksuid.ID) *WorkoutCreate {
	wc.mutation.SetUsersID(id)
	return wc
}

// SetUsers sets the "users" edge to the User entity.
func (wc *WorkoutCreate) SetUsers(u *User) *WorkoutCreate {
	return wc.SetUsersID(u.ID)
}

// AddExerciseIDs adds the "exercises" edge to the Exercise entity by IDs.
func (wc *WorkoutCreate) AddExerciseIDs(ids ...pksuid.ID) *WorkoutCreate {
	wc.mutation.AddExerciseIDs(ids...)
	return wc
}

// AddExercises adds the "exercises" edges to the Exercise entity.
func (wc *WorkoutCreate) AddExercises(e ...*Exercise) *WorkoutCreate {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return wc.AddExerciseIDs(ids...)
}

// AddWorkoutLogIDs adds the "workout_logs" edge to the WorkoutLog entity by IDs.
func (wc *WorkoutCreate) AddWorkoutLogIDs(ids ...pksuid.ID) *WorkoutCreate {
	wc.mutation.AddWorkoutLogIDs(ids...)
	return wc
}

// AddWorkoutLogs adds the "workout_logs" edges to the WorkoutLog entity.
func (wc *WorkoutCreate) AddWorkoutLogs(w ...*WorkoutLog) *WorkoutCreate {
	ids := make([]pksuid.ID, len(w))
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
	if err := wc.defaults(); err != nil {
		return nil, err
	}
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
func (wc *WorkoutCreate) defaults() error {
	if _, ok := wc.mutation.Name(); !ok {
		v := workout.DefaultName
		wc.mutation.SetName(v)
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		if workout.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized workout.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := workout.DefaultCreatedAt()
		wc.mutation.SetCreatedAt(v)
	}
	if _, ok := wc.mutation.ID(); !ok {
		if workout.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized workout.DefaultID (forgotten import ent/runtime?)")
		}
		v := workout.DefaultID()
		wc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkoutCreate) check() error {
	if _, ok := wc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Workout.name"`)}
	}
	if _, ok := wc.mutation.Volume(); !ok {
		return &ValidationError{Name: "volume", err: errors.New(`ent: missing required field "Workout.volume"`)}
	}
	if _, ok := wc.mutation.Duration(); !ok {
		return &ValidationError{Name: "duration", err: errors.New(`ent: missing required field "Workout.duration"`)}
	}
	if _, ok := wc.mutation.Sets(); !ok {
		return &ValidationError{Name: "sets", err: errors.New(`ent: missing required field "Workout.sets"`)}
	}
	if _, ok := wc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Workout.created_at"`)}
	}
	if _, ok := wc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Workout.user_id"`)}
	}
	if _, ok := wc.mutation.UsersID(); !ok {
		return &ValidationError{Name: "users", err: errors.New(`ent: missing required edge "Workout.users"`)}
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
		if id, ok := _spec.ID.Value.(*pksuid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
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
	_spec.OnConflict = wc.conflict
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wc.mutation.Name(); ok {
		_spec.SetField(workout.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wc.mutation.Volume(); ok {
		_spec.SetField(workout.FieldVolume, field.TypeFloat64, value)
		_node.Volume = value
	}
	if value, ok := wc.mutation.Duration(); ok {
		_spec.SetField(workout.FieldDuration, field.TypeString, value)
		_node.Duration = value
	}
	if value, ok := wc.mutation.Sets(); ok {
		_spec.SetField(workout.FieldSets, field.TypeInt, value)
		_node.Sets = value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.SetField(workout.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wc.mutation.Image(); ok {
		_spec.SetField(workout.FieldImage, field.TypeJSON, value)
		_node.Image = value
	}
	if value, ok := wc.mutation.Description(); ok {
		_spec.SetField(workout.FieldDescription, field.TypeString, value)
		_node.Description = &value
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
	if nodes := wc.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   workout.ExercisesTable,
			Columns: workout.ExercisesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exercise.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &WorkoutLogCreate{config: wc.config, mutation: newWorkoutLogMutation(wc.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.WorkoutLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Workout.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.WorkoutUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (wc *WorkoutCreate) OnConflict(opts ...sql.ConflictOption) *WorkoutUpsertOne {
	wc.conflict = opts
	return &WorkoutUpsertOne{
		create: wc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Workout.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (wc *WorkoutCreate) OnConflictColumns(columns ...string) *WorkoutUpsertOne {
	wc.conflict = append(wc.conflict, sql.ConflictColumns(columns...))
	return &WorkoutUpsertOne{
		create: wc,
	}
}

type (
	// WorkoutUpsertOne is the builder for "upsert"-ing
	//  one Workout node.
	WorkoutUpsertOne struct {
		create *WorkoutCreate
	}

	// WorkoutUpsert is the "OnConflict" setter.
	WorkoutUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *WorkoutUpsert) SetName(v string) *WorkoutUpsert {
	u.Set(workout.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *WorkoutUpsert) UpdateName() *WorkoutUpsert {
	u.SetExcluded(workout.FieldName)
	return u
}

// SetVolume sets the "volume" field.
func (u *WorkoutUpsert) SetVolume(v float64) *WorkoutUpsert {
	u.Set(workout.FieldVolume, v)
	return u
}

// UpdateVolume sets the "volume" field to the value that was provided on create.
func (u *WorkoutUpsert) UpdateVolume() *WorkoutUpsert {
	u.SetExcluded(workout.FieldVolume)
	return u
}

// AddVolume adds v to the "volume" field.
func (u *WorkoutUpsert) AddVolume(v float64) *WorkoutUpsert {
	u.Add(workout.FieldVolume, v)
	return u
}

// SetDuration sets the "duration" field.
func (u *WorkoutUpsert) SetDuration(v string) *WorkoutUpsert {
	u.Set(workout.FieldDuration, v)
	return u
}

// UpdateDuration sets the "duration" field to the value that was provided on create.
func (u *WorkoutUpsert) UpdateDuration() *WorkoutUpsert {
	u.SetExcluded(workout.FieldDuration)
	return u
}

// SetSets sets the "sets" field.
func (u *WorkoutUpsert) SetSets(v int) *WorkoutUpsert {
	u.Set(workout.FieldSets, v)
	return u
}

// UpdateSets sets the "sets" field to the value that was provided on create.
func (u *WorkoutUpsert) UpdateSets() *WorkoutUpsert {
	u.SetExcluded(workout.FieldSets)
	return u
}

// AddSets adds v to the "sets" field.
func (u *WorkoutUpsert) AddSets(v int) *WorkoutUpsert {
	u.Add(workout.FieldSets, v)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *WorkoutUpsert) SetCreatedAt(v time.Time) *WorkoutUpsert {
	u.Set(workout.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *WorkoutUpsert) UpdateCreatedAt() *WorkoutUpsert {
	u.SetExcluded(workout.FieldCreatedAt)
	return u
}

// SetImage sets the "image" field.
func (u *WorkoutUpsert) SetImage(v *schematype.Image) *WorkoutUpsert {
	u.Set(workout.FieldImage, v)
	return u
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *WorkoutUpsert) UpdateImage() *WorkoutUpsert {
	u.SetExcluded(workout.FieldImage)
	return u
}

// ClearImage clears the value of the "image" field.
func (u *WorkoutUpsert) ClearImage() *WorkoutUpsert {
	u.SetNull(workout.FieldImage)
	return u
}

// SetDescription sets the "description" field.
func (u *WorkoutUpsert) SetDescription(v string) *WorkoutUpsert {
	u.Set(workout.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *WorkoutUpsert) UpdateDescription() *WorkoutUpsert {
	u.SetExcluded(workout.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *WorkoutUpsert) ClearDescription() *WorkoutUpsert {
	u.SetNull(workout.FieldDescription)
	return u
}

// SetUserID sets the "user_id" field.
func (u *WorkoutUpsert) SetUserID(v pksuid.ID) *WorkoutUpsert {
	u.Set(workout.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *WorkoutUpsert) UpdateUserID() *WorkoutUpsert {
	u.SetExcluded(workout.FieldUserID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Workout.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(workout.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *WorkoutUpsertOne) UpdateNewValues() *WorkoutUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(workout.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Workout.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *WorkoutUpsertOne) Ignore() *WorkoutUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WorkoutUpsertOne) DoNothing() *WorkoutUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WorkoutCreate.OnConflict
// documentation for more info.
func (u *WorkoutUpsertOne) Update(set func(*WorkoutUpsert)) *WorkoutUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WorkoutUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *WorkoutUpsertOne) SetName(v string) *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *WorkoutUpsertOne) UpdateName() *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateName()
	})
}

// SetVolume sets the "volume" field.
func (u *WorkoutUpsertOne) SetVolume(v float64) *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetVolume(v)
	})
}

// AddVolume adds v to the "volume" field.
func (u *WorkoutUpsertOne) AddVolume(v float64) *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.AddVolume(v)
	})
}

// UpdateVolume sets the "volume" field to the value that was provided on create.
func (u *WorkoutUpsertOne) UpdateVolume() *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateVolume()
	})
}

// SetDuration sets the "duration" field.
func (u *WorkoutUpsertOne) SetDuration(v string) *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetDuration(v)
	})
}

// UpdateDuration sets the "duration" field to the value that was provided on create.
func (u *WorkoutUpsertOne) UpdateDuration() *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateDuration()
	})
}

// SetSets sets the "sets" field.
func (u *WorkoutUpsertOne) SetSets(v int) *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetSets(v)
	})
}

// AddSets adds v to the "sets" field.
func (u *WorkoutUpsertOne) AddSets(v int) *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.AddSets(v)
	})
}

// UpdateSets sets the "sets" field to the value that was provided on create.
func (u *WorkoutUpsertOne) UpdateSets() *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateSets()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *WorkoutUpsertOne) SetCreatedAt(v time.Time) *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *WorkoutUpsertOne) UpdateCreatedAt() *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetImage sets the "image" field.
func (u *WorkoutUpsertOne) SetImage(v *schematype.Image) *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *WorkoutUpsertOne) UpdateImage() *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateImage()
	})
}

// ClearImage clears the value of the "image" field.
func (u *WorkoutUpsertOne) ClearImage() *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.ClearImage()
	})
}

// SetDescription sets the "description" field.
func (u *WorkoutUpsertOne) SetDescription(v string) *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *WorkoutUpsertOne) UpdateDescription() *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *WorkoutUpsertOne) ClearDescription() *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.ClearDescription()
	})
}

// SetUserID sets the "user_id" field.
func (u *WorkoutUpsertOne) SetUserID(v pksuid.ID) *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *WorkoutUpsertOne) UpdateUserID() *WorkoutUpsertOne {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *WorkoutUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WorkoutCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WorkoutUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *WorkoutUpsertOne) ID(ctx context.Context) (id pksuid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: WorkoutUpsertOne.ID is not supported by MySQL driver. Use WorkoutUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *WorkoutUpsertOne) IDX(ctx context.Context) pksuid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// WorkoutCreateBulk is the builder for creating many Workout entities in bulk.
type WorkoutCreateBulk struct {
	config
	err      error
	builders []*WorkoutCreate
	conflict []sql.ConflictOption
}

// Save creates the Workout entities in the database.
func (wcb *WorkoutCreateBulk) Save(ctx context.Context) ([]*Workout, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
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
					spec.OnConflict = wcb.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Workout.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.WorkoutUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (wcb *WorkoutCreateBulk) OnConflict(opts ...sql.ConflictOption) *WorkoutUpsertBulk {
	wcb.conflict = opts
	return &WorkoutUpsertBulk{
		create: wcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Workout.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (wcb *WorkoutCreateBulk) OnConflictColumns(columns ...string) *WorkoutUpsertBulk {
	wcb.conflict = append(wcb.conflict, sql.ConflictColumns(columns...))
	return &WorkoutUpsertBulk{
		create: wcb,
	}
}

// WorkoutUpsertBulk is the builder for "upsert"-ing
// a bulk of Workout nodes.
type WorkoutUpsertBulk struct {
	create *WorkoutCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Workout.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(workout.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *WorkoutUpsertBulk) UpdateNewValues() *WorkoutUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(workout.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Workout.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *WorkoutUpsertBulk) Ignore() *WorkoutUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WorkoutUpsertBulk) DoNothing() *WorkoutUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WorkoutCreateBulk.OnConflict
// documentation for more info.
func (u *WorkoutUpsertBulk) Update(set func(*WorkoutUpsert)) *WorkoutUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WorkoutUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *WorkoutUpsertBulk) SetName(v string) *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *WorkoutUpsertBulk) UpdateName() *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateName()
	})
}

// SetVolume sets the "volume" field.
func (u *WorkoutUpsertBulk) SetVolume(v float64) *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetVolume(v)
	})
}

// AddVolume adds v to the "volume" field.
func (u *WorkoutUpsertBulk) AddVolume(v float64) *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.AddVolume(v)
	})
}

// UpdateVolume sets the "volume" field to the value that was provided on create.
func (u *WorkoutUpsertBulk) UpdateVolume() *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateVolume()
	})
}

// SetDuration sets the "duration" field.
func (u *WorkoutUpsertBulk) SetDuration(v string) *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetDuration(v)
	})
}

// UpdateDuration sets the "duration" field to the value that was provided on create.
func (u *WorkoutUpsertBulk) UpdateDuration() *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateDuration()
	})
}

// SetSets sets the "sets" field.
func (u *WorkoutUpsertBulk) SetSets(v int) *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetSets(v)
	})
}

// AddSets adds v to the "sets" field.
func (u *WorkoutUpsertBulk) AddSets(v int) *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.AddSets(v)
	})
}

// UpdateSets sets the "sets" field to the value that was provided on create.
func (u *WorkoutUpsertBulk) UpdateSets() *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateSets()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *WorkoutUpsertBulk) SetCreatedAt(v time.Time) *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *WorkoutUpsertBulk) UpdateCreatedAt() *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetImage sets the "image" field.
func (u *WorkoutUpsertBulk) SetImage(v *schematype.Image) *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *WorkoutUpsertBulk) UpdateImage() *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateImage()
	})
}

// ClearImage clears the value of the "image" field.
func (u *WorkoutUpsertBulk) ClearImage() *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.ClearImage()
	})
}

// SetDescription sets the "description" field.
func (u *WorkoutUpsertBulk) SetDescription(v string) *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *WorkoutUpsertBulk) UpdateDescription() *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *WorkoutUpsertBulk) ClearDescription() *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.ClearDescription()
	})
}

// SetUserID sets the "user_id" field.
func (u *WorkoutUpsertBulk) SetUserID(v pksuid.ID) *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *WorkoutUpsertBulk) UpdateUserID() *WorkoutUpsertBulk {
	return u.Update(func(s *WorkoutUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *WorkoutUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the WorkoutCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WorkoutCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WorkoutUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
