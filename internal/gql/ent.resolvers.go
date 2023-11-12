package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"entgo.io/contrib/entgql"
	gigachad "github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id pksuid.ID) (ent.Noder, error) {
	return r.client.Noder(ctx, id, ent.WithNodeType(ent.IDToType))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []pksuid.ID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids, ent.WithNodeType(ent.IDToType))
}

// EquipmentSlice is the resolver for the equipmentSlice field.
func (r *queryResolver) EquipmentSlice(ctx context.Context, after *entgql.Cursor[pksuid.ID], first *int, before *entgql.Cursor[pksuid.ID], last *int, where *ent.EquipmentWhereInput) (*ent.EquipmentConnection, error) {
	return r.client.Equipment.Query().Paginate(ctx, after, first, before, last, ent.WithEquipmentFilter(where.Filter))
}

// Exercises is the resolver for the exercises field.
func (r *queryResolver) Exercises(ctx context.Context, after *entgql.Cursor[pksuid.ID], first *int, before *entgql.Cursor[pksuid.ID], last *int, where *ent.ExerciseWhereInput) (*ent.ExerciseConnection, error) {
	return r.client.Exercise.Query().Paginate(ctx, after, first, before, last, ent.WithExerciseFilter(where.Filter))
}

// ExerciseTypes is the resolver for the exerciseTypes field.
func (r *queryResolver) ExerciseTypes(ctx context.Context, after *entgql.Cursor[pksuid.ID], first *int, before *entgql.Cursor[pksuid.ID], last *int, where *ent.ExerciseTypeWhereInput) (*ent.ExerciseTypeConnection, error) {
	return r.client.ExerciseType.Query().Paginate(ctx, after, first, before, last, ent.WithExerciseTypeFilter(where.Filter))
}

// MusclesGroups is the resolver for the musclesGroups field.
func (r *queryResolver) MusclesGroups(ctx context.Context, after *entgql.Cursor[pksuid.ID], first *int, before *entgql.Cursor[pksuid.ID], last *int, where *ent.MusclesGroupWhereInput) (*ent.MusclesGroupConnection, error) {
	return r.client.MusclesGroup.Query().Paginate(ctx, after, first, before, last, ent.WithMusclesGroupFilter(where.Filter))
}

// Routines is the resolver for the routines field.
func (r *queryResolver) Routines(ctx context.Context, after *entgql.Cursor[pksuid.ID], first *int, before *entgql.Cursor[pksuid.ID], last *int, where *ent.RoutineWhereInput) (*ent.RoutineConnection, error) {
	return r.client.Routine.Query().Paginate(ctx, after, first, before, last, ent.WithRoutineFilter(where.Filter))
}

// RoutineExercises is the resolver for the routineExercises field.
func (r *queryResolver) RoutineExercises(ctx context.Context, after *entgql.Cursor[pksuid.ID], first *int, before *entgql.Cursor[pksuid.ID], last *int, where *ent.RoutineExerciseWhereInput) (*ent.RoutineExerciseConnection, error) {
	return r.client.RoutineExercise.Query().Paginate(ctx, after, first, before, last, ent.WithRoutineExerciseFilter(where.Filter))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[pksuid.ID], first *int, before *entgql.Cursor[pksuid.ID], last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.client.User.Query().Paginate(ctx, after, first, before, last, ent.WithUserFilter(where.Filter))
}

// Workouts is the resolver for the workouts field.
func (r *queryResolver) Workouts(ctx context.Context, after *entgql.Cursor[pksuid.ID], first *int, before *entgql.Cursor[pksuid.ID], last *int, where *ent.WorkoutWhereInput) (*ent.WorkoutConnection, error) {
	return r.client.Workout.Query().Paginate(ctx, after, first, before, last, ent.WithWorkoutFilter(where.Filter))
}

// WorkoutLogs is the resolver for the workoutLogs field.
func (r *queryResolver) WorkoutLogs(ctx context.Context, after *entgql.Cursor[pksuid.ID], first *int, before *entgql.Cursor[pksuid.ID], last *int, where *ent.WorkoutLogWhereInput) (*ent.WorkoutLogConnection, error) {
	return r.client.WorkoutLog.Query().Paginate(ctx, after, first, before, last, ent.WithWorkoutLogFilter(where.Filter))
}

// Query returns gigachad.QueryResolver implementation.
func (r *Resolver) Query() gigachad.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
