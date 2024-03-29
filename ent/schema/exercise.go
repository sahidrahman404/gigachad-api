package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/privacy"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
	"github.com/sahidrahman404/gigachad-api/rule"
)

// Exercise holds the schema definition for the Exercise entity.
type Exercise struct {
	ent.Schema
}

func (Exercise) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.AllowEditExerciseIfOwner(),
		},
		Query: privacy.QueryPolicy{
			rule.FilterExerciseRule(),
		},
	}
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
		field.JSON("image", &schematype.Image{}).Annotations(entgql.Type("Image")).Optional(),
		field.String("how_to").Optional().Nillable(),
		field.String("user_id").Optional().Nillable().GoType(pksuid.ID("")),
	}
}

// Edges of the Exercise.
func (Exercise) Edges() []ent.Edge {
	return []ent.Edge{
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
		edge.From("workouts", Workout.Type).
			Ref("exercises").
			Through("workout_logs", WorkoutLog.Type).
			Annotations(entgql.RelayConnection()),
	}
}

func (Exercise) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
