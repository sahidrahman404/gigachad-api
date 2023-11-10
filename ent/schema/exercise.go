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

// Exercise holds the schema definition for the Exercise entity.
type Exercise struct {
	ent.Schema
}

func (Exercise) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pksuid.MixinWithPrefix("EX"),
	}
}

// Fields of the Exercise.
func (Exercise) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("image").Optional().Nillable(),
		field.String("how_to").Optional().Nillable(),
		field.String("user_id").Optional().Nillable().GoType(pksuid.ID("")),
	}
}

// Edges of the Exercise.
func (Exercise) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("workout_logs", WorkoutLog.Type).
			Annotations(
				entgql.RelayConnection(),
				entsql.OnDelete(entsql.Cascade),
			),
		edge.From("users", User.Type).Ref("exercises").Field("user_id").Unique(),
		edge.From("equipment", Equipment.Type).
			Ref("exercises").
			Annotations(entgql.RelayConnection()),
		edge.From("muscles_groups", MusclesGroup.Type).
			Ref("exercises").
			Annotations(entgql.RelayConnection()),
		edge.From("exercise_types", ExerciseType.Type).
			Ref("exercises").
			Annotations(entgql.RelayConnection()),
		edge.From("routines", Routine.Type).
			Ref("exercises").
			Through("routine_exercises", RoutineExercise.Type).
			Annotations(entgql.RelayConnection()),
	}
}

func (Exercise) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
