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
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
)

// RoutineExercise holds the schema definition for the RoutineExercise entity.
type RoutineExercise struct {
	ent.Schema
}

func (RoutineExercise) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pksuid.MixinWithPrefix("RE"),
	}
}

// Fields of the RoutineExercise.
func (RoutineExercise) Fields() []ent.Field {
	return []ent.Field{
		field.String("rest_time").Optional().Nillable(),
		field.JSON("sets", []*schematype.Set{}).Annotations(entgql.Type("[Set!]")),
		field.String("routine_id").GoType(pksuid.ID("")),
		field.String("exercise_id").GoType(pksuid.ID("")),
		field.String("user_id").GoType(pksuid.ID("")),
	}
}

// Edges of the RoutineExercise.
func (RoutineExercise) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("routines", Routine.Type).
			Field("routine_id").
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
		edge.From("users", User.Type).
			Ref("routine_exercises").
			Field("user_id").
			Unique().
			Required(),
	}
}

func (RoutineExercise) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("routine_id"),
	}
}

func (RoutineExercise) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
