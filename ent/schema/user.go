package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pksuid.MixinWithPrefix("UR"),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email"),
		field.String("username"),
		field.String("hashed_password").Sensitive(),
		field.String("name"),
		field.String("created_at").
			DefaultFunc(generateTime).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Int("activated").
			Default(0).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.Int("version").
			Default(1).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username").Unique(),
		index.Fields("email").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tokens", Token.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("exercises", Exercise.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("routines", Routine.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("workouts", Workout.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("workout_logs", WorkoutLog.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("routine_exercises", RoutineExercise.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
