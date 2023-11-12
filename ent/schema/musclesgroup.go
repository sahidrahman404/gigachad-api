package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
)

// Muscles_Group holds the schema definition for the Muscles_Group entity.
type MusclesGroup struct {
	ent.Schema
}

func (MusclesGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pksuid.MixinWithPrefix("MG"),
	}
}

// Fields of the Muscles_Group.
func (MusclesGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.JSON("image", schematype.Image{}).Annotations(entgql.Type("Image")),
	}
}

// Edges of the Muscles_Group.
func (MusclesGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("exercises", Exercise.Type).
			Annotations(
				entgql.RelayConnection(),
				entsql.OnDelete(entsql.Cascade),
			),
	}
}

func (MusclesGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
