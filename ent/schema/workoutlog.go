package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
)

// WorkoutLog holds the schema definition for the WorkoutLog entity.
type WorkoutLog struct {
	ent.Schema
}

// Fields of the WorkoutLog.
func (WorkoutLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(generateKSUID),
		field.JSON("sets", &schematype.Sets{}).Annotations(entgql.Type("Map")),
		field.String("created_at").DefaultFunc(generateTime),
		field.String("exercise_id").Optional(),
		field.String("workout_id").Optional(),
		field.String("user_id").Optional(),
	}
}

// Edges of the WorkoutLog.
func (WorkoutLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("workout_logs").Field("user_id").Unique(),
		edge.From("exercises", Exercise.Type).Ref("workout_logs").Field("exercise_id").Unique(),
		edge.From("workouts", Workout.Type).Ref("workout_logs").Field("workout_id").Unique(),
	}
}
