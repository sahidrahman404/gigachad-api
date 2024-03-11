package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"buf.build/gen/go/sahidrahman/gigachadapis/connectrpc/go/gigachad/v1/gigachadv1connect"
	gigachadv1 "buf.build/gen/go/sahidrahman/gigachadapis/protocolbuffers/go/gigachad/v1"
	"connectrpc.com/connect"
	gigachad "github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
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
	reminderID := fmt.Sprintf("%s-%s", "weekly-workout-reminder", routineID)

	if err := r.WithTx(ctx, func(tx *ent.Tx) error {
		txClient := tx.Client()
		routine, err := txClient.Routine.Create().
			SetID(routineID).
			SetName(input.Name).
			SetScheduleID(reminderID).
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

	client := gigachadv1connect.NewReminderServiceClient(http.DefaultClient, "http://localhost:8080")

	exercises := []*gigachadv1.Exercise{}

	for _, v := range routine.Edges.RoutineExercises {
		zero := 0
		emptyStr := ""
		sets := []*gigachadv1.Set{}

		for _, v := range v.Sets {
			if v.Reps == nil {
				v.Reps = &zero
			}

			if v.Kg == nil {
				v.Kg = &zero
			}

			if v.Duration == nil {
				v.Duration = &emptyStr
			}

			if v.Km == nil {
				v.Km = &zero
			}
			sets = append(sets, &gigachadv1.Set{
				Reps:     int32(*v.Reps),
				Kg:       int32(*v.Kg),
				Duration: *v.Duration,
				Km:       int32(*v.Km),
			})
		}

		if v.RestTime == nil {
			v.RestTime = &emptyStr
		}

		exercises = append(exercises, &gigachadv1.Exercise{
			Name:     v.Edges.Exercises.Name,
			RestTime: *v.RestTime,
			Sets:     sets,
		})
	}

	schedules := []*gigachadv1.Schedule{}

	for _, v := range input.Reminder {
		schedules = append(schedules, &gigachadv1.Schedule{
			Day:    int32(v.Day),
			Hour:   int32(v.Hour),
			Minute: int32(v.Minute),
			Second: int32(v.Second),
		})
	}

	userName := strings.Split(user.Name, " ")
	userLastName := userName[len(userName)-1]

	client.CreateReminders(ctx, &connect.Request[gigachadv1.CreateRemindersRequest]{
		Msg: &gigachadv1.CreateRemindersRequest{
			AddReminderRequest: &gigachadv1.AddReminderRequest{
				ReminderId:   reminderID,
				UserLastName: userLastName,
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
				c.SetNillableRestTime(input.RoutineExercises[i].RestTime).
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

	return r.client.Routine.Get(ctx, input.ID)
}
