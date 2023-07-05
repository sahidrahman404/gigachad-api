package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
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
		field.String("id").DefaultFunc(generateKSUID).Annotations(entgql.OrderField("ID")),
		field.String("name"),
		field.Int("volume"),
		field.Int("reps"),
		field.String("time").Optional(),
		field.Int("sets"),
		field.String("created_at").
			DefaultFunc(generateTime).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("image").Optional().Nillable(),
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

func (Workout) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
