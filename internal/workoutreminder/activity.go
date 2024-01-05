package workoutreminder

import (
	"context"

	"github.com/pborman/uuid"
	"github.com/sahidrahman404/gigachad-api"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/internal/shared"
	"github.com/sahidrahman404/gigachad-api/internal/smtp"
	"go.temporal.io/sdk/client"
)

type WorkoutReminder struct {
	temporalClient *client.Client
	entClient      *ent.Client
}

func NewWorkoutReminder(tc *client.Client, e *ent.Client) *WorkoutReminder {
	return &WorkoutReminder{
		temporalClient: tc,
		entClient:      e,
	}
}

func (wr *WorkoutReminder) GetUserData(ctx context.Context, input pksuid.ID) (*ent.User, error) {
	return wr.entClient.User.Get(ctx, input)
}

type CreateDailyWorkoutScheduleInput struct {
	WorkoutName string
	ScheduleID  string
	User        *ent.User
	Reminders   []*gigachad.CreateRoutineReminderInput
}

func (wr *WorkoutReminder) CreateDailyWorkoutSchedule(ctx context.Context, input CreateDailyWorkoutScheduleInput) error {
	u := input.User

	calendars := []client.ScheduleCalendarSpec{}

	for _, val := range input.Reminders {
		day := val.Day
		seconds := val.Second
		minutes := val.Minute
		hour := val.Hour
		dayOfWeek := []client.ScheduleRange{{
			Start: day,
		}}

		schedlueCalendarSpec := client.ScheduleCalendarSpec{
			Second: []client.ScheduleRange{
				{Start: seconds},
			},
			Minute: []client.ScheduleRange{
				{Start: minutes},
			},
			Hour: []client.ScheduleRange{
				{Start: hour},
			},
			DayOfWeek: dayOfWeek,
		}

		calendars = append(calendars, schedlueCalendarSpec)
	}

	spec := client.ScheduleSpec{
		Calendars: calendars,
	}

	args := smtp.DailyWorkoutReminderInput{
		Recipient:   u.Email,
		Name:        u.Name,
		WorkoutName: input.WorkoutName,
	}

	action := client.ScheduleWorkflowAction{
		ID:        "daily-workout-reminder-email-" + uuid.New(),
		TaskQueue: shared.DailyWorkoutReminderTaskQueueName,
		Workflow:  smtp.SendDailyWorkoutReminderWorkflow,
		Args:      []interface{}{args},
	}

	tc := *wr.temporalClient

	_, err := tc.ScheduleClient().Create(ctx, client.ScheduleOptions{
		ID:             input.ScheduleID,
		Spec:           spec,
		Action:         &action,
		PauseOnFailure: true,
	})
	if err != nil {
		return err
	}

	return nil
}

type UpdateDailyWorkoutScheduleInput struct {
	ScheduleID string
	Schedules  []*gigachad.CreateRoutineReminderInput
}

func (wr *WorkoutReminder) UpdateDailyWorkoutSchedule(ctx context.Context, input UpdateDailyWorkoutScheduleInput) error {
	calendars := []client.ScheduleCalendarSpec{}

	for _, val := range input.Schedules {
		day := val.Day
		seconds := val.Second
		minutes := val.Minute
		hour := val.Hour
		dayOfWeek := []client.ScheduleRange{{
			Start: day,
		}}

		schedlueCalendarSpec := client.ScheduleCalendarSpec{
			Second: []client.ScheduleRange{
				{Start: seconds},
			},
			Minute: []client.ScheduleRange{
				{Start: minutes},
			},
			Hour: []client.ScheduleRange{
				{Start: hour},
			},
			DayOfWeek: dayOfWeek,
		}

		calendars = append(calendars, schedlueCalendarSpec)
	}

	spec := client.ScheduleSpec{
		Calendars: calendars,
	}

	tc := *wr.temporalClient

	scheduleHandler := tc.ScheduleClient().GetHandle(ctx, input.ScheduleID)

	err := scheduleHandler.Update(ctx, client.ScheduleUpdateOptions{
		DoUpdate: func(sui client.ScheduleUpdateInput) (*client.ScheduleUpdate, error) {
			schedule := sui.Description.Schedule
			schedule.Spec = &spec
			return &client.ScheduleUpdate{
				Schedule: &schedule,
			}, nil
		},
	})
	if err != nil {
		return err
	}

	return nil
}
