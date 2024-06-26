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

// Workout holds the schema definition for the Workout entity.
type Workout struct {
	ent.Schema
}

func (Workout) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.FilterByUserIDRule(),
		},
		Query: privacy.QueryPolicy{
			rule.FilterByUserIDRule(),
		},
	}
}

func (Workout) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pksuid.MixinWithPrefix("WO"),
	}
}

// Fields of the Workout.
func (Workout) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default(""),
		field.Float("volume"),
		field.String("duration"),
		field.Int("sets"),
		field.Time("created_at").
			Default(generateTime).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.JSON("image", &schematype.Image{}).Annotations(entgql.Type("Image")).Optional(),
		field.String("description").Optional().Nillable(),
		field.String("user_id").GoType(pksuid.ID("")),
	}
}

// Edges of the Workout.
func (Workout) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("workouts").
			Field("user_id").
			Unique().
			Required(),
		edge.To("exercises", Exercise.Type).
			Through("workout_logs", WorkoutLog.Type).
			Annotations(
				entgql.RelayConnection(),
				entsql.OnDelete(entsql.Cascade),
			),
	}
}

func (Workout) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
