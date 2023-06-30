package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Exercise_Type holds the schema definition for the Exercise_Type entity.
type ExerciseType struct {
	ent.Schema
}

// Fields of the Exercise_Type.
func (ExerciseType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(generateKSUID),
		field.String("name"),
		field.Strings("properties"),
		field.String("description"),
	}
}

// Edges of the Exercise_Type.
func (ExerciseType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("exercises", Exercise.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

