package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Exercise holds the schema definition for the Exercise entity.
type Exercise struct {
	ent.Schema
}

// Fields of the Exercise.
func (Exercise) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(generateKSUID),
		field.String("name"),
		field.String("how_to"),
		field.String("equipment_id").Optional(),
		field.String("muscles_group_id").Optional(),
		field.String("exercise_type_id").Optional(),
		field.String("user_id").Optional(),
	}
}

// Edges of the Exercise.
func (Exercise) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("routine_exercises", RoutineExercise.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.To("workout_logs", WorkoutLog.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
		edge.From("users", User.Type).Ref("exercises").Field("user_id").Unique(),
		edge.From("equipments", Equipment.Type).Ref("exercises").Field("muscles_group_id").Unique(),
		edge.From("muscles_groups", MusclesGroup.Type).Ref("exercises").Field("muscles_group_id").Unique(),
		edge.From("exercise_types", ExerciseType.Type).Ref("exercises").Field("exercise_type_id").Unique(),
	}
}
