package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
)

// RoutineExercise holds the schema definition for the RoutineExercise entity.
type RoutineExercise struct {
	ent.Schema
}

// Fields of the RoutineExercise.
func (RoutineExercise) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(generateKSUID),
		field.Int("rest_timer").Optional(),
		field.JSON("sets", &schematype.Sets{}),
		field.String("routine_id").Optional(),
		field.String("exercise_id").Optional(),
		field.String("user_id").Optional(),
	}
}

// Edges of the RoutineExercise.
func (RoutineExercise) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("routines", Routine.Type).Ref("routine_exercises").Field("routine_id").Unique(),
		edge.From("exercises", Exercise.Type).Ref("routine_exercises").Field("exercise_id").Unique(),
		edge.From("users", User.Type).Ref("routine_exercises").Field("user_id").Unique(),
	}
}
