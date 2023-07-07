package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

// Workout holds the schema definition for the Workout entity.
type Workout struct {
	ent.Schema
}

func (Workout) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pksuid.MixinWithPrefix("WO"),
	}
}

// Fields of the Workout.
func (Workout) Fields() []ent.Field {
	return []ent.Field{
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
		field.String("user_id").Optional().GoType(pksuid.ID("")),
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
