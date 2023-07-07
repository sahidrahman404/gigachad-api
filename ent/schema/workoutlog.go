package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
)

// WorkoutLog holds the schema definition for the WorkoutLog entity.
type WorkoutLog struct {
	ent.Schema
}

func (WorkoutLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		pksuid.MixinWithPrefix("WL"),
	}
}

// Fields of the WorkoutLog.
func (WorkoutLog) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("sets", &schematype.Sets{}).Annotations(entgql.Type("[Set!]")),
		field.String("created_at").DefaultFunc(generateTime).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
		field.String("exercise_id").Optional().GoType(pksuid.ID("")),
		field.String("workout_id").Optional().GoType(pksuid.ID("")),
		field.String("user_id").Optional().GoType(pksuid.ID("")),
	}
}

// Edges of the WorkoutLog.
func (WorkoutLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("workout_logs").Field("user_id").Unique(),
		edge.From("exercises", Exercise.Type).Ref("workout_logs").Field("exercise_id").Unique(),
		edge.From("workouts", Workout.Type).Ref("workout_logs").Field("workout_id").Unique(),
	}
}

func (WorkoutLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
