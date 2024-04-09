package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/privacy"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
	"github.com/sahidrahman404/gigachad-api/rule"
)

// WorkoutLog holds the schema definition for the WorkoutLog entity.
type WorkoutLog struct {
	ent.Schema
}

func (WorkoutLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pksuid.MixinWithPrefix("WL"),
	}
}

func (WorkoutLog) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.FilterByUserIDRule(),
		},
		Query: privacy.QueryPolicy{
			rule.FilterByUserIDRule(),
		},
	}
}

// Fields of the WorkoutLog.
func (WorkoutLog) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("sets", []*schematype.Set{}).Annotations(entgql.Type("[Set!]")),
		field.Time("created_at").Default(generateTime).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("workout_id").GoType(pksuid.ID("")),
		field.String("exercise_id").GoType(pksuid.ID("")),
		field.String("user_id").GoType(pksuid.ID("")),
	}
}

// Edges of the WorkoutLog.
func (WorkoutLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("workout_logs").
			Field("user_id").
			Unique().
			Required(),
		edge.To("workouts", Workout.Type).
			Field("workout_id").
			Unique().
			Required().
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("exercises", Exercise.Type).
			Field("exercise_id").
			Unique().
			Required().
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
	}
}

func (WorkoutLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
