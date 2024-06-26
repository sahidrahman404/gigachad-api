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
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/token"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/ent/workout"
	"github.com/sahidrahman404/gigachad-api/ent/workoutlog"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetHashedPassword sets the "hashed_password" field.
func (uc *UserCreate) SetHashedPassword(s string) *UserCreate {
	uc.mutation.SetHashedPassword(s)
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetUnit sets the "unit" field.
func (uc *UserCreate) SetUnit(u user.Unit) *UserCreate {
	uc.mutation.SetUnit(u)
	return uc
}

// SetNillableUnit sets the "unit" field if the given value is not nil.
func (uc *UserCreate) SetNillableUnit(u *user.Unit) *UserCreate {
	if u != nil {
		uc.SetUnit(*u)
	}
	return uc
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetActivated sets the "activated" field.
func (uc *UserCreate) SetActivated(i int) *UserCreate {
	uc.mutation.SetActivated(i)
	return uc
}

// SetNillableActivated sets the "activated" field if the given value is not nil.
func (uc *UserCreate) SetNillableActivated(i *int) *UserCreate {
	if i != nil {
		uc.SetActivated(*i)
	}
	return uc
}

// SetVersion sets the "version" field.
func (uc *UserCreate) SetVersion(i int) *UserCreate {
	uc.mutation.SetVersion(i)
	return uc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (uc *UserCreate) SetNillableVersion(i *int) *UserCreate {
	if i != nil {
		uc.SetVersion(*i)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(pk pksuid.ID) *UserCreate {
	uc.mutation.SetID(pk)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(pk *pksuid.ID) *UserCreate {
	if pk != nil {
		uc.SetID(*pk)
	}
	return uc
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (uc *UserCreate) AddTokenIDs(ids ...pksuid.ID) *UserCreate {
	uc.mutation.AddTokenIDs(ids...)
	return uc
}

// AddTokens adds the "tokens" edges to the Token entity.
func (uc *UserCreate) AddTokens(t ...*Token) *UserCreate {
	ids := make([]pksuid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uc.AddTokenIDs(ids...)
}

// AddExerciseIDs adds the "exercises" edge to the Exercise entity by IDs.
func (uc *UserCreate) AddExerciseIDs(ids ...pksuid.ID) *UserCreate {
	uc.mutation.AddExerciseIDs(ids...)
	return uc
}

// AddExercises adds the "exercises" edges to the Exercise entity.
func (uc *UserCreate) AddExercises(e ...*Exercise) *UserCreate {
	ids := make([]pksuid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uc.AddExerciseIDs(ids...)
}

// AddRoutineIDs adds the "routines" edge to the Routine entity by IDs.
func (uc *UserCreate) AddRoutineIDs(ids ...pksuid.ID) *UserCreate {
	uc.mutation.AddRoutineIDs(ids...)
	return uc
}

// AddRoutines adds the "routines" edges to the Routine entity.
func (uc *UserCreate) AddRoutines(r ...*Routine) *UserCreate {
	ids := make([]pksuid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRoutineIDs(ids...)
}

// AddWorkoutIDs adds the "workouts" edge to the Workout entity by IDs.
func (uc *UserCreate) AddWorkoutIDs(ids ...pksuid.ID) *UserCreate {
	uc.mutation.AddWorkoutIDs(ids...)
	return uc
}

// AddWorkouts adds the "workouts" edges to the Workout entity.
func (uc *UserCreate) AddWorkouts(w ...*Workout) *UserCreate {
	ids := make([]pksuid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uc.AddWorkoutIDs(ids...)
}

// AddWorkoutLogIDs adds the "workout_logs" edge to the WorkoutLog entity by IDs.
func (uc *UserCreate) AddWorkoutLogIDs(ids ...pksuid.ID) *UserCreate {
	uc.mutation.AddWorkoutLogIDs(ids...)
	return uc
}

// AddWorkoutLogs adds the "workout_logs" edges to the WorkoutLog entity.
func (uc *UserCreate) AddWorkoutLogs(w ...*WorkoutLog) *UserCreate {
	ids := make([]pksuid.ID, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uc.AddWorkoutLogIDs(ids...)
}

// AddRoutineExerciseIDs adds the "routine_exercises" edge to the RoutineExercise entity by IDs.
func (uc *UserCreate) AddRoutineExerciseIDs(ids ...pksuid.ID) *UserCreate {
	uc.mutation.AddRoutineExerciseIDs(ids...)
	return uc
}

// AddRoutineExercises adds the "routine_exercises" edges to the RoutineExercise entity.
func (uc *UserCreate) AddRoutineExercises(r ...*RoutineExercise) *UserCreate {
	ids := make([]pksuid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uc.AddRoutineExerciseIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.Unit(); !ok {
		v := user.DefaultUnit
		uc.mutation.SetUnit(v)
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.Activated(); !ok {
		v := user.DefaultActivated
		uc.mutation.SetActivated(v)
	}
	if _, ok := uc.mutation.Version(); !ok {
		v := user.DefaultVersion
		uc.mutation.SetVersion(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "User.username"`)}
	}
	if _, ok := uc.mutation.HashedPassword(); !ok {
		return &ValidationError{Name: "hashed_password", err: errors.New(`ent: missing required field "User.hashed_password"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if _, ok := uc.mutation.Unit(); !ok {
		return &ValidationError{Name: "unit", err: errors.New(`ent: missing required field "User.unit"`)}
	}
	if v, ok := uc.mutation.Unit(); ok {
		if err := user.UnitValidator(v); err != nil {
			return &ValidationError{Name: "unit", err: fmt.Errorf(`ent: validator failed for field "User.unit": %w`, err)}
		}
	}
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.Activated(); !ok {
		return &ValidationError{Name: "activated", err: errors.New(`ent: missing required field "User.activated"`)}
	}
	if _, ok := uc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "User.version"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
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
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	_spec.OnConflict = uc.conflict
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.HashedPassword(); ok {
		_spec.SetField(user.FieldHashedPassword, field.TypeString, value)
		_node.HashedPassword = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Unit(); ok {
		_spec.SetField(user.FieldUnit, field.TypeEnum, value)
		_node.Unit = value
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.Activated(); ok {
		_spec.SetField(user.FieldActivated, field.TypeInt, value)
		_node.Activated = value
	}
	if value, ok := uc.mutation.Version(); ok {
		_spec.SetField(user.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if nodes := uc.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ExercisesTable,
			Columns: []string{user.ExercisesColumn},
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
	if nodes := uc.mutation.RoutinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RoutinesTable,
			Columns: []string{user.RoutinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(routine.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.WorkoutsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WorkoutsTable,
			Columns: []string{user.WorkoutsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(workout.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.WorkoutLogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WorkoutLogsTable,
			Columns: []string{user.WorkoutLogsColumn},
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
	if nodes := uc.mutation.RoutineExercisesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RoutineExercisesTable,
			Columns: []string{user.RoutineExercisesColumn},
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetEmail(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetEmail(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetEmail sets the "email" field.
func (u *UserUpsert) SetEmail(v string) *UserUpsert {
	u.Set(user.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsert) UpdateEmail() *UserUpsert {
	u.SetExcluded(user.FieldEmail)
	return u
}

// SetUsername sets the "username" field.
func (u *UserUpsert) SetUsername(v string) *UserUpsert {
	u.Set(user.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UserUpsert) UpdateUsername() *UserUpsert {
	u.SetExcluded(user.FieldUsername)
	return u
}

// SetHashedPassword sets the "hashed_password" field.
func (u *UserUpsert) SetHashedPassword(v string) *UserUpsert {
	u.Set(user.FieldHashedPassword, v)
	return u
}

// UpdateHashedPassword sets the "hashed_password" field to the value that was provided on create.
func (u *UserUpsert) UpdateHashedPassword() *UserUpsert {
	u.SetExcluded(user.FieldHashedPassword)
	return u
}

// SetName sets the "name" field.
func (u *UserUpsert) SetName(v string) *UserUpsert {
	u.Set(user.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsert) UpdateName() *UserUpsert {
	u.SetExcluded(user.FieldName)
	return u
}

// SetUnit sets the "unit" field.
func (u *UserUpsert) SetUnit(v user.Unit) *UserUpsert {
	u.Set(user.FieldUnit, v)
	return u
}

// UpdateUnit sets the "unit" field to the value that was provided on create.
func (u *UserUpsert) UpdateUnit() *UserUpsert {
	u.SetExcluded(user.FieldUnit)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserUpsert) SetCreatedAt(v time.Time) *UserUpsert {
	u.Set(user.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserUpsert) UpdateCreatedAt() *UserUpsert {
	u.SetExcluded(user.FieldCreatedAt)
	return u
}

// SetActivated sets the "activated" field.
func (u *UserUpsert) SetActivated(v int) *UserUpsert {
	u.Set(user.FieldActivated, v)
	return u
}

// UpdateActivated sets the "activated" field to the value that was provided on create.
func (u *UserUpsert) UpdateActivated() *UserUpsert {
	u.SetExcluded(user.FieldActivated)
	return u
}

// AddActivated adds v to the "activated" field.
func (u *UserUpsert) AddActivated(v int) *UserUpsert {
	u.Add(user.FieldActivated, v)
	return u
}

// SetVersion sets the "version" field.
func (u *UserUpsert) SetVersion(v int) *UserUpsert {
	u.Set(user.FieldVersion, v)
	return u
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *UserUpsert) UpdateVersion() *UserUpsert {
	u.SetExcluded(user.FieldVersion)
	return u
}

// AddVersion adds v to the "version" field.
func (u *UserUpsert) AddVersion(v int) *UserUpsert {
	u.Add(user.FieldVersion, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(user.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *UserUpsertOne) SetEmail(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// SetUsername sets the "username" field.
func (u *UserUpsertOne) SetUsername(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUsername() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUsername()
	})
}

// SetHashedPassword sets the "hashed_password" field.
func (u *UserUpsertOne) SetHashedPassword(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetHashedPassword(v)
	})
}

// UpdateHashedPassword sets the "hashed_password" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateHashedPassword() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateHashedPassword()
	})
}

// SetName sets the "name" field.
func (u *UserUpsertOne) SetName(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// SetUnit sets the "unit" field.
func (u *UserUpsertOne) SetUnit(v user.Unit) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUnit(v)
	})
}

// UpdateUnit sets the "unit" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUnit() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUnit()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *UserUpsertOne) SetCreatedAt(v time.Time) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateCreatedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetActivated sets the "activated" field.
func (u *UserUpsertOne) SetActivated(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetActivated(v)
	})
}

// AddActivated adds v to the "activated" field.
func (u *UserUpsertOne) AddActivated(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddActivated(v)
	})
}

// UpdateActivated sets the "activated" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateActivated() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateActivated()
	})
}

// SetVersion sets the "version" field.
func (u *UserUpsertOne) SetVersion(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetVersion(v)
	})
}

// AddVersion adds v to the "version" field.
func (u *UserUpsertOne) AddVersion(v int) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateVersion() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateVersion()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id pksuid.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: UserUpsertOne.ID is not supported by MySQL driver. Use UserUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) pksuid.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetEmail(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(user.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetEmail sets the "email" field.
func (u *UserUpsertBulk) SetEmail(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// SetUsername sets the "username" field.
func (u *UserUpsertBulk) SetUsername(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUsername() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUsername()
	})
}

// SetHashedPassword sets the "hashed_password" field.
func (u *UserUpsertBulk) SetHashedPassword(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetHashedPassword(v)
	})
}

// UpdateHashedPassword sets the "hashed_password" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateHashedPassword() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateHashedPassword()
	})
}

// SetName sets the "name" field.
func (u *UserUpsertBulk) SetName(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// SetUnit sets the "unit" field.
func (u *UserUpsertBulk) SetUnit(v user.Unit) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUnit(v)
	})
}

// UpdateUnit sets the "unit" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUnit() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUnit()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *UserUpsertBulk) SetCreatedAt(v time.Time) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateCreatedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetActivated sets the "activated" field.
func (u *UserUpsertBulk) SetActivated(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetActivated(v)
	})
}

// AddActivated adds v to the "activated" field.
func (u *UserUpsertBulk) AddActivated(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddActivated(v)
	})
}

// UpdateActivated sets the "activated" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateActivated() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateActivated()
	})
}

// SetVersion sets the "version" field.
func (u *UserUpsertBulk) SetVersion(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetVersion(v)
	})
}

// AddVersion adds v to the "version" field.
func (u *UserUpsertBulk) AddVersion(v int) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddVersion(v)
	})
}

// UpdateVersion sets the "version" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateVersion() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateVersion()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
