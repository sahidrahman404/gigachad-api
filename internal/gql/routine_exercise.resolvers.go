package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"fmt"

	"github.com/pborman/uuid"
	gigachad "github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/internal/shared"
	"github.com/sahidrahman404/gigachad-api/internal/smtp"
	"github.com/sahidrahman404/gigachad-api/internal/workoutreminder"
	"go.temporal.io/sdk/client"
)

// CreateRoutineWithChildren is the resolver for the CreateRoutineWithChildren field.
func (r *mutationResolver) CreateRoutineWithChildren(ctx context.Context, input gigachad.CreateRoutineWithChildrenInput) (*ent.Routine, error) {
	userCtx := r.getUserFromCtx(ctx)
	user, err := r.requireActivatedUser(userCtx)
	if err != nil {
		return nil, err
	}

	userID := user.ID
	routineID := pksuid.MustNew("RO")
	scheduleID := fmt.Sprintf("%s-%s", "daily-workout-reminder-schedule-", routineID)

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		txClient := tx.Client()
		routine, err := txClient.Routine.Create().
			SetID(routineID).
			SetName(input.Name).
			SetScheduleID(scheduleID).
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
					SetUserID(userID)
			}).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, r.serverError(err)
	}

	if input.Reminder == nil {
		return r.client.Routine.Get(ctx, routineID)
	}

	routine, err := r.client.Routine.Query().Where(routine.ID(routineID)).WithRoutineExercises(func(req *ent.RoutineExerciseQuery) {
		req.WithExercises()
	}).Only(ctx)
	if err != nil {
		return r.client.Routine.Get(ctx, routineID)
	}
	tc := *r.temporalClient

	options := client.StartWorkflowOptions{
		ID:        "create-daily-workout-reminder-schedule" + uuid.New(),
		TaskQueue: shared.DailyWorkoutReminderTaskQueueName,
	}

	cdwsi := workoutreminder.CreateDailyWorkoutScheduleInput{
		User: user,
		RoutineData: smtp.RoutineData{
			UserName: user.Name,
			Routine:  *routine,
		},
		Reminders:  input.Reminder,
		ScheduleID: scheduleID,
	}

	_, err = tc.ExecuteWorkflow(ctx, options, workoutreminder.CreateDailyWorkoutScheduleWorkflow, cdwsi)

	return r.client.Routine.Get(ctx, routineID)
}

// UpdateRoutineWithChildren is the resolver for the updateRoutineWithChildren field.
func (r *mutationResolver) UpdateRoutineWithChildren(ctx context.Context, input gigachad.UpdateRoutineWithChildrenInput) (*ent.Routine, error) {
	userCtx := r.getUserFromCtx(ctx)
	user, err := r.requireActivatedUser(userCtx)
	if err != nil {
		return nil, err
	}

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		txClient := tx.Client()
		routine, err := txClient.Routine.UpdateOneID(input.ID).
			SetName(input.Name).
			SetNillableScheduleID(input.Reminder.ID).
			Save(ctx)
		if err != nil {
			return err
		}
		err = txClient.RoutineExercise.MapCreateBulk(
			input.RoutineExercises,
			func(c *ent.RoutineExerciseCreate, i int) {
				c.SetNillableRestTimer(input.RoutineExercises[i].RestTimer).
					SetSets(input.RoutineExercises[i].Sets).
					SetRoutineID(routine.ID).
					SetExerciseID(input.RoutineExercises[i].ExerciseID).
					SetUserID(user.ID).
					SetNillableID(input.RoutineExercises[i].ID)
			}).
			OnConflict().
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, r.serverError(err)
	}

	if input.Reminder == nil {
		return r.client.Routine.Get(ctx, input.ID)
	}

	r.backgroundTask(func() error {
		tc := *r.temporalClient

		options := client.StartWorkflowOptions{
			ID:        "update-daily-workout-reminder-schedule" + uuid.New(),
			TaskQueue: shared.DailyWorkoutReminderTaskQueueName,
		}

		udwsi := workoutreminder.UpdateDailyWorkoutScheduleInput{
			ScheduleID: *input.Reminder.ID,
			Schedules:  input.Reminder.Schedules,
		}

		_, err := tc.ExecuteWorkflow(ctx, options, workoutreminder.UpdateDailyWorkoutScheduleWorkflow, udwsi)
		return err
	})

	return r.client.Routine.Get(ctx, input.ID)
}
