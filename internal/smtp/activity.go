package smtp

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/ent"
)

type DailyWorkoutReminderInput struct {
	Recipient   string
	RoutineData RoutineData
}

type RoutineData struct {
	Routine  ent.Routine
	UserName string
}

func (m *Mailer) SendDailyWorkoutReminder(ctx context.Context, input DailyWorkoutReminderInput) error {
	return m.Send(input.Recipient, input.RoutineData, "daily_workout_reminder.tmpl")
}
