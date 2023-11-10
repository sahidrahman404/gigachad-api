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

// Exercise_Type holds the schema definition for the Exercise_Type entity.
type ExerciseType struct {
	ent.Schema
}

func (ExerciseType) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pksuid.MixinWithPrefix("ET"),
	}
}

// Fields of the Exercise_Type.
func (ExerciseType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Strings("properties"),
		field.String("description"),
	}
}

// Edges of the Exercise_Type.
func (ExerciseType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("exercises", Exercise.Type).
			Annotations(
				entgql.RelayConnection(),
				entsql.OnDelete(entsql.Cascade),
			),
	}
}

func (ExerciseType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
