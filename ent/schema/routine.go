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

// Routine holds the schema definition for the Routine entity.
type Routine struct {
	ent.Schema
}

func (Routine) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.FilterByUserIDRule(),
		},
		Query: privacy.QueryPolicy{
			rule.FilterByUserIDRule(),
		},
	}
}

func (Routine) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pksuid.MixinWithPrefix("RO"),
	}
}

// Fields of the Routine.
func (Routine) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("reminder_id").Optional().Nillable(),
		field.JSON("reminders", []*schematype.Reminder{}).
			Optional().
			Annotations(entgql.Type("[Reminder!]")),
		field.String("user_id").GoType(pksuid.ID("")),
	}
}

// Edges of the Routine.
func (Routine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("exercises", Exercise.Type).
			Through("routine_exercises", RoutineExercise.Type).
			Annotations(
				entgql.RelayConnection(),
				entsql.OnDelete(entsql.Cascade),
			),
		edge.From("users", User.Type).
			Ref("routines").
			Field("user_id").
			Unique().
			Required(),
	}
}

func (Routine) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
