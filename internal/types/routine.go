package types

import (
	"fmt"

	gigachadv1 "buf.build/gen/go/sahidrahman/gigachadapis/protocolbuffers/go/gigachad/v1"
	"github.com/google/uuid"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

var (
	EntRoutineNotFound = `ent: routine not found`
	zero               = 0
	emptyString        = ""
)

type CreateRoutineParams struct {
	Name string `json:"name"`
}

type Routine struct {
	Ent *ent.Routine
}

func NewRoutineFromParams(p CreateRoutineParams, userID pksuid.ID) *Routine {
	r := &Routine{
		Ent: &ent.Routine{
			Name:   p.Name,
			UserID: userID,
		},
	}
	return r
}

func (p CreateRoutineParams) Validate(v *validator.Validator) {
	v.CheckField(p.Name != "", "name", "name must be provided")
}

func GetReminderID(reminders []*schematype.Reminder) *string {
	if reminders == nil {
		return nil
	}
	uuid := uuid.New()
	id := fmt.Sprintf("%s-%s", "weekly-workout-reminder", uuid.String())
	return &id
}

func GetReminders(reminders []*schematype.Reminder) []*schematype.Reminder {
	if reminders == nil {
		return nil
	}
	return reminders
}

type UpdateRemindersResult struct {
	ID        *string
	Reminders []*schematype.Reminder
	Update    bool
}

func UpdateReminders(
	newReminders []*schematype.Reminder,
	oldReminders []*schematype.Reminder,
	oldID *string,
) UpdateRemindersResult {
	uuid := uuid.New()
	newID := fmt.Sprintf("%s-%s", "weekly-workout-reminder", uuid.String())
	if newReminders != nil && oldReminders == nil {
		return UpdateRemindersResult{
			ID:        &newID,
			Reminders: newReminders,
			Update:    true,
		}
	}
	if len(newReminders) != len(oldReminders) {
		return UpdateRemindersResult{
			ID:        &newID,
			Reminders: newReminders,
			Update:    true,
		}
	}
	for i, n := range newReminders {
		if n != oldReminders[i] {
			return UpdateRemindersResult{
				ID:        &newID,
				Reminders: newReminders,
				Update:    true,
			}
		}
	}
	return UpdateRemindersResult{
		ID:        oldID,
		Reminders: oldReminders,
		Update:    false,
	}
}

func GetExercises(r *ent.Routine) []*gigachadv1.Exercise {
	exercises := []*gigachadv1.Exercise{}

	for _, routine := range r.Edges.RoutineExercises {
		sets := []*gigachadv1.Set{}

		for _, set := range routine.Sets {
			setNilToEmptyValue(set)
			sets = append(sets, &gigachadv1.Set{
				Reps:     int32(*set.Reps),
				Kg:       int32(*set.Weight),
				Duration: *set.Duration,
				Km:       int32(*set.Length),
			})
		}

		exercises = append(exercises, &gigachadv1.Exercise{
			Name:     routine.Edges.Exercises.Name,
			RestTime: restTimeMap[*routine.RestTime],
			Sets:     sets,
		})
	}
	return exercises
}

func GetSchedules(reminders []*schematype.Reminder) []*gigachadv1.Schedule {
	schedules := []*gigachadv1.Schedule{}

	for _, v := range reminders {
		schedules = append(schedules, &gigachadv1.Schedule{
			Day:    int32(v.Day),
			Hour:   int32(v.Hour),
			Minute: int32(v.Minute),
			Second: int32(v.Second),
		})
	}

	return schedules
}

func setNilToEmptyValue(s *schematype.Set) {
	if s.Reps == nil {
		s.Reps = &zero
	}

	if s.Weight == nil {
		s.Weight = &zero
	}

	if s.Duration == nil {
		s.Duration = &emptyString
	}

	if s.Length == nil {
		s.Length = &zero
	}
}

func GetOldReminderID(r *ent.Routine) string {
	if r.ReminderID == nil {
		return ""
	}
	return *r.ReminderID
}

var restTimeMap = map[string]string{
	"0":   "0",
	"50":  "50s",
	"60":  "1min",
	"65":  "1min 5s",
	"70":  "1min 10s",
	"75":  "1min 15s",
	"80":  "1min 20s",
	"85":  "1min 25s",
	"90":  "1min 30s",
	"95":  "1min 35s",
	"100": "1min 40s",
	"105": "1min 45s",
	"110": "1min 50s",
	"115": "1min 55s",
	"120": "2min",
	"125": "2min 5s",
	"130": "2min 10s",
	"135": "2min 15s",
	"140": "2min 20s",
	"145": "2min 25s",
	"150": "2min 30s",
	"155": "2min 35s",
	"160": "2min 40s",
	"165": "2min 45s",
	"170": "2min 50s",
	"175": "2min 55s",
	"180": "3min",
	"185": "3min 5s",
	"190": "3min 10s",
	"195": "3min 15s",
	"200": "3min 20s",
	"205": "3min 25s",
	"210": "3min 30s",
	"215": "3min 35s",
	"220": "3min 40s",
	"225": "3min 45s",
	"230": "3min 50s",
	"235": "3min 55s",
	"240": "4min",
	"245": "4min 5s",
	"250": "4min 10s",
	"255": "4min 15s",
	"260": "4min 20s",
	"265": "4min 25s",
	"270": "4min 30s",
	"275": "4min 35s",
	"280": "4min 40s",
	"285": "4min 45s",
	"290": "4min 50s",
	"295": "4min 55s",
	"300": "5min",
}
