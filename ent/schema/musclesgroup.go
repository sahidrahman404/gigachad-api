package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Muscles_Group holds the schema definition for the Muscles_Group entity.
type MusclesGroup struct {
	ent.Schema
}

// Fields of the Muscles_Group.
func (MusclesGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(generateKSUID),
		field.String("name"),
		field.String("image"),
	}
}

// Edges of the Muscles_Group.
func (MusclesGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("exercises", Exercise.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (MusclesGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
