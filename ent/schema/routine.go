package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Routine holds the schema definition for the Routine entity.
type Routine struct {
	ent.Schema
}

// Fields of the Routine.
func (Routine) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(generateKSUID),
		field.String("name"),
		field.String("user_id").Optional(),
	}
}

// Edges of the Routine.
func (Routine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("routine_exercises", RoutineExercise.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("users", User.Type).Ref("routines").Field("user_id").Unique(),
	}
}
