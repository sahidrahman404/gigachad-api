// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (e *Equipment) Exercises(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ExerciseOrder, where *ExerciseWhereInput,
) (*ExerciseConnection, error) {
	opts := []ExercisePaginateOption{
		WithExerciseOrder(orderBy),
		WithExerciseFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := e.Edges.totalCount[0][alias]
	if nodes, err := e.NamedExercises(alias); err == nil || hasTotalCount {
		pager, err := newExercisePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ExerciseConnection{Edges: []*ExerciseEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return e.QueryExercises().Paginate(ctx, after, first, before, last, opts...)
}

func (e *Exercise) Users(ctx context.Context) (*User, error) {
	result, err := e.Edges.UsersOrErr()
	if IsNotLoaded(err) {
		result, err = e.QueryUsers().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (e *Exercise) Equipment(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *EquipmentOrder, where *EquipmentWhereInput,
) (*EquipmentConnection, error) {
	opts := []EquipmentPaginateOption{
		WithEquipmentOrder(orderBy),
		WithEquipmentFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := e.Edges.totalCount[1][alias]
	if nodes, err := e.NamedEquipment(alias); err == nil || hasTotalCount {
		pager, err := newEquipmentPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &EquipmentConnection{Edges: []*EquipmentEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return e.QueryEquipment().Paginate(ctx, after, first, before, last, opts...)
}

func (e *Exercise) MusclesGroups(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *MusclesGroupOrder, where *MusclesGroupWhereInput,
) (*MusclesGroupConnection, error) {
	opts := []MusclesGroupPaginateOption{
		WithMusclesGroupOrder(orderBy),
		WithMusclesGroupFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := e.Edges.totalCount[2][alias]
	if nodes, err := e.NamedMusclesGroups(alias); err == nil || hasTotalCount {
		pager, err := newMusclesGroupPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &MusclesGroupConnection{Edges: []*MusclesGroupEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return e.QueryMusclesGroups().Paginate(ctx, after, first, before, last, opts...)
}

func (e *Exercise) ExerciseTypes(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ExerciseTypeOrder, where *ExerciseTypeWhereInput,
) (*ExerciseTypeConnection, error) {
	opts := []ExerciseTypePaginateOption{
		WithExerciseTypeOrder(orderBy),
		WithExerciseTypeFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := e.Edges.totalCount[3][alias]
	if nodes, err := e.NamedExerciseTypes(alias); err == nil || hasTotalCount {
		pager, err := newExerciseTypePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ExerciseTypeConnection{Edges: []*ExerciseTypeEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return e.QueryExerciseTypes().Paginate(ctx, after, first, before, last, opts...)
}

func (e *Exercise) Routines(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *RoutineOrder, where *RoutineWhereInput,
) (*RoutineConnection, error) {
	opts := []RoutinePaginateOption{
		WithRoutineOrder(orderBy),
		WithRoutineFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := e.Edges.totalCount[4][alias]
	if nodes, err := e.NamedRoutines(alias); err == nil || hasTotalCount {
		pager, err := newRoutinePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &RoutineConnection{Edges: []*RoutineEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return e.QueryRoutines().Paginate(ctx, after, first, before, last, opts...)
}

func (e *Exercise) Workouts(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *WorkoutOrder, where *WorkoutWhereInput,
) (*WorkoutConnection, error) {
	opts := []WorkoutPaginateOption{
		WithWorkoutOrder(orderBy),
		WithWorkoutFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := e.Edges.totalCount[5][alias]
	if nodes, err := e.NamedWorkouts(alias); err == nil || hasTotalCount {
		pager, err := newWorkoutPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &WorkoutConnection{Edges: []*WorkoutEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return e.QueryWorkouts().Paginate(ctx, after, first, before, last, opts...)
}

func (e *Exercise) RoutineExercises(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *RoutineExerciseOrder, where *RoutineExerciseWhereInput,
) (*RoutineExerciseConnection, error) {
	opts := []RoutineExercisePaginateOption{
		WithRoutineExerciseOrder(orderBy),
		WithRoutineExerciseFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := e.Edges.totalCount[6][alias]
	if nodes, err := e.NamedRoutineExercises(alias); err == nil || hasTotalCount {
		pager, err := newRoutineExercisePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &RoutineExerciseConnection{Edges: []*RoutineExerciseEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return e.QueryRoutineExercises().Paginate(ctx, after, first, before, last, opts...)
}

func (e *Exercise) WorkoutLogs(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *WorkoutLogOrder, where *WorkoutLogWhereInput,
) (*WorkoutLogConnection, error) {
	opts := []WorkoutLogPaginateOption{
		WithWorkoutLogOrder(orderBy),
		WithWorkoutLogFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := e.Edges.totalCount[7][alias]
	if nodes, err := e.NamedWorkoutLogs(alias); err == nil || hasTotalCount {
		pager, err := newWorkoutLogPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &WorkoutLogConnection{Edges: []*WorkoutLogEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return e.QueryWorkoutLogs().Paginate(ctx, after, first, before, last, opts...)
}

func (et *ExerciseType) Exercises(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ExerciseOrder, where *ExerciseWhereInput,
) (*ExerciseConnection, error) {
	opts := []ExercisePaginateOption{
		WithExerciseOrder(orderBy),
		WithExerciseFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := et.Edges.totalCount[0][alias]
	if nodes, err := et.NamedExercises(alias); err == nil || hasTotalCount {
		pager, err := newExercisePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ExerciseConnection{Edges: []*ExerciseEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return et.QueryExercises().Paginate(ctx, after, first, before, last, opts...)
}

func (mg *MusclesGroup) Exercises(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ExerciseOrder, where *ExerciseWhereInput,
) (*ExerciseConnection, error) {
	opts := []ExercisePaginateOption{
		WithExerciseOrder(orderBy),
		WithExerciseFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := mg.Edges.totalCount[0][alias]
	if nodes, err := mg.NamedExercises(alias); err == nil || hasTotalCount {
		pager, err := newExercisePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ExerciseConnection{Edges: []*ExerciseEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return mg.QueryExercises().Paginate(ctx, after, first, before, last, opts...)
}

func (r *Routine) Exercises(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ExerciseOrder, where *ExerciseWhereInput,
) (*ExerciseConnection, error) {
	opts := []ExercisePaginateOption{
		WithExerciseOrder(orderBy),
		WithExerciseFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := r.Edges.totalCount[0][alias]
	if nodes, err := r.NamedExercises(alias); err == nil || hasTotalCount {
		pager, err := newExercisePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ExerciseConnection{Edges: []*ExerciseEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return r.QueryExercises().Paginate(ctx, after, first, before, last, opts...)
}

func (r *Routine) Users(ctx context.Context) (*User, error) {
	result, err := r.Edges.UsersOrErr()
	if IsNotLoaded(err) {
		result, err = r.QueryUsers().Only(ctx)
	}
	return result, err
}

func (r *Routine) RoutineExercises(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *RoutineExerciseOrder, where *RoutineExerciseWhereInput,
) (*RoutineExerciseConnection, error) {
	opts := []RoutineExercisePaginateOption{
		WithRoutineExerciseOrder(orderBy),
		WithRoutineExerciseFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := r.Edges.totalCount[2][alias]
	if nodes, err := r.NamedRoutineExercises(alias); err == nil || hasTotalCount {
		pager, err := newRoutineExercisePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &RoutineExerciseConnection{Edges: []*RoutineExerciseEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return r.QueryRoutineExercises().Paginate(ctx, after, first, before, last, opts...)
}

func (re *RoutineExercise) Routines(ctx context.Context) (*Routine, error) {
	result, err := re.Edges.RoutinesOrErr()
	if IsNotLoaded(err) {
		result, err = re.QueryRoutines().Only(ctx)
	}
	return result, err
}

func (re *RoutineExercise) Exercises(ctx context.Context) (*Exercise, error) {
	result, err := re.Edges.ExercisesOrErr()
	if IsNotLoaded(err) {
		result, err = re.QueryExercises().Only(ctx)
	}
	return result, err
}

func (re *RoutineExercise) Users(ctx context.Context) (*User, error) {
	result, err := re.Edges.UsersOrErr()
	if IsNotLoaded(err) {
		result, err = re.QueryUsers().Only(ctx)
	}
	return result, err
}

func (t *Token) Users(ctx context.Context) (*User, error) {
	result, err := t.Edges.UsersOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryUsers().Only(ctx)
	}
	return result, err
}

func (u *User) Tokens(ctx context.Context) (result []*Token, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedTokens(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.TokensOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryTokens().All(ctx)
	}
	return result, err
}

func (u *User) Exercises(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ExerciseOrder, where *ExerciseWhereInput,
) (*ExerciseConnection, error) {
	opts := []ExercisePaginateOption{
		WithExerciseOrder(orderBy),
		WithExerciseFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[1][alias]
	if nodes, err := u.NamedExercises(alias); err == nil || hasTotalCount {
		pager, err := newExercisePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ExerciseConnection{Edges: []*ExerciseEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryExercises().Paginate(ctx, after, first, before, last, opts...)
}

func (u *User) Routines(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *RoutineOrder, where *RoutineWhereInput,
) (*RoutineConnection, error) {
	opts := []RoutinePaginateOption{
		WithRoutineOrder(orderBy),
		WithRoutineFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[2][alias]
	if nodes, err := u.NamedRoutines(alias); err == nil || hasTotalCount {
		pager, err := newRoutinePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &RoutineConnection{Edges: []*RoutineEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryRoutines().Paginate(ctx, after, first, before, last, opts...)
}

func (u *User) Workouts(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *WorkoutOrder, where *WorkoutWhereInput,
) (*WorkoutConnection, error) {
	opts := []WorkoutPaginateOption{
		WithWorkoutOrder(orderBy),
		WithWorkoutFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[3][alias]
	if nodes, err := u.NamedWorkouts(alias); err == nil || hasTotalCount {
		pager, err := newWorkoutPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &WorkoutConnection{Edges: []*WorkoutEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryWorkouts().Paginate(ctx, after, first, before, last, opts...)
}

func (u *User) WorkoutLogs(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *WorkoutLogOrder, where *WorkoutLogWhereInput,
) (*WorkoutLogConnection, error) {
	opts := []WorkoutLogPaginateOption{
		WithWorkoutLogOrder(orderBy),
		WithWorkoutLogFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[4][alias]
	if nodes, err := u.NamedWorkoutLogs(alias); err == nil || hasTotalCount {
		pager, err := newWorkoutLogPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &WorkoutLogConnection{Edges: []*WorkoutLogEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryWorkoutLogs().Paginate(ctx, after, first, before, last, opts...)
}

func (u *User) RoutineExercises(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *RoutineExerciseOrder, where *RoutineExerciseWhereInput,
) (*RoutineExerciseConnection, error) {
	opts := []RoutineExercisePaginateOption{
		WithRoutineExerciseOrder(orderBy),
		WithRoutineExerciseFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := u.Edges.totalCount[5][alias]
	if nodes, err := u.NamedRoutineExercises(alias); err == nil || hasTotalCount {
		pager, err := newRoutineExercisePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &RoutineExerciseConnection{Edges: []*RoutineExerciseEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return u.QueryRoutineExercises().Paginate(ctx, after, first, before, last, opts...)
}

func (w *Workout) Users(ctx context.Context) (*User, error) {
	result, err := w.Edges.UsersOrErr()
	if IsNotLoaded(err) {
		result, err = w.QueryUsers().Only(ctx)
	}
	return result, err
}

func (w *Workout) Exercises(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *ExerciseOrder, where *ExerciseWhereInput,
) (*ExerciseConnection, error) {
	opts := []ExercisePaginateOption{
		WithExerciseOrder(orderBy),
		WithExerciseFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := w.Edges.totalCount[1][alias]
	if nodes, err := w.NamedExercises(alias); err == nil || hasTotalCount {
		pager, err := newExercisePager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &ExerciseConnection{Edges: []*ExerciseEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return w.QueryExercises().Paginate(ctx, after, first, before, last, opts...)
}

func (w *Workout) WorkoutLogs(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, orderBy *WorkoutLogOrder, where *WorkoutLogWhereInput,
) (*WorkoutLogConnection, error) {
	opts := []WorkoutLogPaginateOption{
		WithWorkoutLogOrder(orderBy),
		WithWorkoutLogFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := w.Edges.totalCount[2][alias]
	if nodes, err := w.NamedWorkoutLogs(alias); err == nil || hasTotalCount {
		pager, err := newWorkoutLogPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &WorkoutLogConnection{Edges: []*WorkoutLogEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return w.QueryWorkoutLogs().Paginate(ctx, after, first, before, last, opts...)
}

func (wl *WorkoutLog) Users(ctx context.Context) (*User, error) {
	result, err := wl.Edges.UsersOrErr()
	if IsNotLoaded(err) {
		result, err = wl.QueryUsers().Only(ctx)
	}
	return result, err
}

func (wl *WorkoutLog) Workouts(ctx context.Context) (*Workout, error) {
	result, err := wl.Edges.WorkoutsOrErr()
	if IsNotLoaded(err) {
		result, err = wl.QueryWorkouts().Only(ctx)
	}
	return result, err
}

func (wl *WorkoutLog) Exercises(ctx context.Context) (*Exercise, error) {
	result, err := wl.Edges.ExercisesOrErr()
	if IsNotLoaded(err) {
		result, err = wl.QueryExercises().Only(ctx)
	}
	return result, err
}
