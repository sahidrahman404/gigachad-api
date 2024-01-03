package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"

	gigachad "github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

// CreateRoutineWithChildren is the resolver for the CreateRoutineWithChildren field.
func (r *mutationResolver) CreateRoutineWithChildren(ctx context.Context, input gigachad.CreateRoutineWithChildrenInput) (*ent.Routine, error) {
	userCtx := r.getUserFromCtx(ctx)
	user, err := r.requireActivatedUser(userCtx)
	if err != nil {
		return nil, err
	}

	routineID := pksuid.MustNew("RO")
	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		txClient := tx.Client()
		routine, err := txClient.Routine.Create().
			SetID(routineID).
			SetName(input.Name).
			SetUserID(user.ID).
			Save(ctx)
		if err != nil {
			return err
		}
		_, err = txClient.RoutineExercise.MapCreateBulk(
			input.RoutineExercises,
			func(c *ent.RoutineExerciseCreate, i int) {
				c.SetNillableRestTimer(input.RoutineExercises[i].RestTimer).
					SetSets(input.RoutineExercises[i].Sets).
					SetRoutineID(routine.ID).
					SetExerciseID(input.RoutineExercises[i].ExerciseID).
					SetUserID(user.ID)
			}).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, r.serverError(err)
	}
	return r.client.Routine.Get(ctx, routineID)
}
