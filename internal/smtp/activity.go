package smtp

import (
	"context"
)

type DailyWorkoutReminderInput struct {
	Recipient   string
	Name        string
	WorkoutName string
}

func (m *Mailer) SendDailyWorkoutReminder(ctx context.Context, input DailyWorkoutReminderInput) error {
	data := map[string]interface{}{
		"name":        input.Name,
		"workoutName": input.WorkoutName,
	}
	return m.Send(input.Recipient, data, "daily_workout_reminder.tmpl")
}
