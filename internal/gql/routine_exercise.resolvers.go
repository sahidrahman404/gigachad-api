package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	gigachadv1 "buf.build/gen/go/sahidrahman/gigachadapis/protocolbuffers/go/gigachad/v1"
	"connectrpc.com/connect"
	gigachad "github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/internal/types"
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
	reminderID := types.GetReminderID(input.Reminders)
	reminders := types.GetReminders(input.Reminders)

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		txClient := tx.Client()
		routine, err := txClient.Routine.Create().
			SetID(routineID).
			SetName(input.Name).
			SetNillableReminderID(reminderID).
			SetReminders(reminders).
			SetUserID(user.ID).
			Save(ctx)
		if err != nil {
			return err
		}
		_, err = txClient.RoutineExercise.MapCreateBulk(
			input.RoutineExercises,
			func(c *ent.RoutineExerciseCreate, i int) {
				c.SetNillableRestTime(input.RoutineExercises[i].RestTime).
					SetSets(input.RoutineExercises[i].Sets).
					SetRoutineID(routine.ID).
					SetExerciseID(input.RoutineExercises[i].ExerciseID).
					SetUserID(userID).
					SetOrder(i)
			}).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, r.defaultError(err)
	}

	if reminders == nil {
		return r.client.Routine.Get(ctx, routineID)
	}

	routine, err := r.client.Routine.Query().
		Where(routine.ID(routineID)).
		WithRoutineExercises(func(req *ent.RoutineExerciseQuery) {
			req.WithExercises()
			req.Order(ent.Asc(routineexercise.FieldOrder))
		}).
		Only(ctx)
	if err != nil {
		return r.client.Routine.Get(ctx, routineID)
	}

	exercises := types.GetExercises(routine)

	schedules := types.GetSchedules(input.Reminders)

	client := *r.reminderServiceClient
	client.CreateReminders(ctx, &connect.Request[gigachadv1.CreateRemindersRequest]{
		Msg: &gigachadv1.CreateRemindersRequest{
			AddReminderRequest: &gigachadv1.AddReminderRequest{
				ReminderId:   *reminderID,
				UserLastName: userCtx.GetUserLastName(),
				WorkoutName:  input.Name,
				Email:        user.Email,
				Exercises:    exercises,
			},
			Schedules: schedules,
		},
	})

	return r.client.Routine.Get(ctx, routineID)
}

// UpdateRoutineWithChildren is the resolver for the updateRoutineWithChildren field.
func (r *mutationResolver) UpdateRoutineWithChildren(ctx context.Context, input gigachad.UpdateRoutineWithChildrenInput) (*ent.Routine, error) {
	userCtx := r.getUserFromCtx(ctx)
	user, err := r.requireActivatedUser(userCtx)
	if err != nil {
		return nil, err
	}

	routine, err := r.client.Routine.Query().
		Where().
		WithRoutineExercises(func(req *ent.RoutineExerciseQuery) {
			req.WithExercises()
			req.Order(ent.Asc(routineexercise.FieldOrder))
		}).
		Only(ctx)
	if err != nil {
		return nil, r.defaultError(err)
	}

	reminders := types.UpdateReminders(
		input.Reminders.Reminders,
		routine.Reminders,
		routine.ReminderID,
	)

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		txClient := tx.Client()
		routine, err := txClient.Routine.
			UpdateOneID(input.ID).
			SetName(input.Name).
			SetNillableReminderID(reminders.ID).
			SetReminders(reminders.Reminders).
			SetUserID(user.ID).
			Save(ctx)
		if err != nil {
			return err
		}
		err = txClient.RoutineExercise.MapCreateBulk(
			input.RoutineExercises,
			func(c *ent.RoutineExerciseCreate, i int) {
				c.SetNillableRestTime(input.RoutineExercises[i].RestTime).
					SetSets(input.RoutineExercises[i].Sets).
					SetRoutineID(routine.ID).
					SetExerciseID(input.RoutineExercises[i].ExerciseID).
					SetUserID(user.ID).
					SetID(input.RoutineExercises[i].ID)
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

	if !reminders.Update {
		return r.client.Routine.Get(ctx, input.ID)
	}

	exercises := types.GetExercises(routine)

	schedules := types.GetSchedules(input.Reminders.Reminders)

	client := *r.reminderServiceClient

	client.UpdateReminders(ctx, &connect.Request[gigachadv1.UpdateRemindersRequest]{
		Msg: &gigachadv1.UpdateRemindersRequest{
			OldReminderId: *routine.ReminderID,
			CreateRemindersRequest: &gigachadv1.CreateRemindersRequest{
				ReminderId: *reminders.ID,
				Schedules:  schedules,
				AddReminderRequest: &gigachadv1.AddReminderRequest{
					ReminderId:   *reminders.ID,
					UserLastName: userCtx.GetUserLastName(),
					WorkoutName:  input.Name,
					Email:        user.Email,
					Exercises:    exercises,
				},
			},
		},
	})

	return r.client.Routine.Get(ctx, input.ID)
}
