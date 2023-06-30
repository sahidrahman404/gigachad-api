package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Workout holds the schema definition for the Workout entity.
type Workout struct {
	ent.Schema
}

// Fields of the Workout.
func (Workout) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(generateKSUID),
		field.String("name"),
		field.Int("volume"),
		field.Int("reps"),
		field.String("time").Optional(),
		field.Int("sets"),
		field.String("created_at").DefaultFunc(generateTime),
		field.String("image"),
		field.String("description"),
		field.String("user_id").Optional(),
	}
}

// Edges of the Workout.
func (Workout) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("workouts").Field("user_id").Unique(),
		edge.To("workout_logs", WorkoutLog.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
