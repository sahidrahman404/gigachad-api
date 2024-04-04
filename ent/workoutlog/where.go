// Code generated by ent, DO NOT EDIT.

package workoutlog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

// ID filters vertices based on their ID field.
func ID(id pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEQ(FieldCreatedAt, v))
}

// WorkoutID applies equality check predicate on the "workout_id" field. It's identical to WorkoutIDEQ.
func WorkoutID(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEQ(FieldWorkoutID, v))
}

// ExerciseID applies equality check predicate on the "exercise_id" field. It's identical to ExerciseIDEQ.
func ExerciseID(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEQ(FieldExerciseID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEQ(FieldUserID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldLTE(FieldCreatedAt, v))
}

// WorkoutIDEQ applies the EQ predicate on the "workout_id" field.
func WorkoutIDEQ(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEQ(FieldWorkoutID, v))
}

// WorkoutIDNEQ applies the NEQ predicate on the "workout_id" field.
func WorkoutIDNEQ(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldNEQ(FieldWorkoutID, v))
}

// WorkoutIDIn applies the In predicate on the "workout_id" field.
func WorkoutIDIn(vs ...pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldIn(FieldWorkoutID, vs...))
}

// WorkoutIDNotIn applies the NotIn predicate on the "workout_id" field.
func WorkoutIDNotIn(vs ...pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldNotIn(FieldWorkoutID, vs...))
}

// WorkoutIDGT applies the GT predicate on the "workout_id" field.
func WorkoutIDGT(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldGT(FieldWorkoutID, v))
}

// WorkoutIDGTE applies the GTE predicate on the "workout_id" field.
func WorkoutIDGTE(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldGTE(FieldWorkoutID, v))
}

// WorkoutIDLT applies the LT predicate on the "workout_id" field.
func WorkoutIDLT(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldLT(FieldWorkoutID, v))
}

// WorkoutIDLTE applies the LTE predicate on the "workout_id" field.
func WorkoutIDLTE(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldLTE(FieldWorkoutID, v))
}

// WorkoutIDContains applies the Contains predicate on the "workout_id" field.
func WorkoutIDContains(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldContains(FieldWorkoutID, vc))
}

// WorkoutIDHasPrefix applies the HasPrefix predicate on the "workout_id" field.
func WorkoutIDHasPrefix(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldHasPrefix(FieldWorkoutID, vc))
}

// WorkoutIDHasSuffix applies the HasSuffix predicate on the "workout_id" field.
func WorkoutIDHasSuffix(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldHasSuffix(FieldWorkoutID, vc))
}

// WorkoutIDEqualFold applies the EqualFold predicate on the "workout_id" field.
func WorkoutIDEqualFold(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldEqualFold(FieldWorkoutID, vc))
}

// WorkoutIDContainsFold applies the ContainsFold predicate on the "workout_id" field.
func WorkoutIDContainsFold(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldContainsFold(FieldWorkoutID, vc))
}

// ExerciseIDEQ applies the EQ predicate on the "exercise_id" field.
func ExerciseIDEQ(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEQ(FieldExerciseID, v))
}

// ExerciseIDNEQ applies the NEQ predicate on the "exercise_id" field.
func ExerciseIDNEQ(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldNEQ(FieldExerciseID, v))
}

// ExerciseIDIn applies the In predicate on the "exercise_id" field.
func ExerciseIDIn(vs ...pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldIn(FieldExerciseID, vs...))
}

// ExerciseIDNotIn applies the NotIn predicate on the "exercise_id" field.
func ExerciseIDNotIn(vs ...pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldNotIn(FieldExerciseID, vs...))
}

// ExerciseIDGT applies the GT predicate on the "exercise_id" field.
func ExerciseIDGT(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldGT(FieldExerciseID, v))
}

// ExerciseIDGTE applies the GTE predicate on the "exercise_id" field.
func ExerciseIDGTE(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldGTE(FieldExerciseID, v))
}

// ExerciseIDLT applies the LT predicate on the "exercise_id" field.
func ExerciseIDLT(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldLT(FieldExerciseID, v))
}

// ExerciseIDLTE applies the LTE predicate on the "exercise_id" field.
func ExerciseIDLTE(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldLTE(FieldExerciseID, v))
}

// ExerciseIDContains applies the Contains predicate on the "exercise_id" field.
func ExerciseIDContains(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldContains(FieldExerciseID, vc))
}

// ExerciseIDHasPrefix applies the HasPrefix predicate on the "exercise_id" field.
func ExerciseIDHasPrefix(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldHasPrefix(FieldExerciseID, vc))
}

// ExerciseIDHasSuffix applies the HasSuffix predicate on the "exercise_id" field.
func ExerciseIDHasSuffix(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldHasSuffix(FieldExerciseID, vc))
}

// ExerciseIDEqualFold applies the EqualFold predicate on the "exercise_id" field.
func ExerciseIDEqualFold(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldEqualFold(FieldExerciseID, vc))
}

// ExerciseIDContainsFold applies the ContainsFold predicate on the "exercise_id" field.
func ExerciseIDContainsFold(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldContainsFold(FieldExerciseID, vc))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldContains(FieldUserID, vc))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldHasPrefix(FieldUserID, vc))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldHasSuffix(FieldUserID, vc))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldEqualFold(FieldUserID, vc))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v pksuid.ID) predicate.WorkoutLog {
	vc := string(v)
	return predicate.WorkoutLog(sql.FieldContainsFold(FieldUserID, vc))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.WorkoutLog {
	return predicate.WorkoutLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.WorkoutLog {
	return predicate.WorkoutLog(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkouts applies the HasEdge predicate on the "workouts" edge.
func HasWorkouts() predicate.WorkoutLog {
	return predicate.WorkoutLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, WorkoutsTable, WorkoutsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkoutsWith applies the HasEdge predicate on the "workouts" edge with a given conditions (other predicates).
func HasWorkoutsWith(preds ...predicate.Workout) predicate.WorkoutLog {
	return predicate.WorkoutLog(func(s *sql.Selector) {
		step := newWorkoutsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasExercises applies the HasEdge predicate on the "exercises" edge.
func HasExercises() predicate.WorkoutLog {
	return predicate.WorkoutLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ExercisesTable, ExercisesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExercisesWith applies the HasEdge predicate on the "exercises" edge with a given conditions (other predicates).
func HasExercisesWith(preds ...predicate.Exercise) predicate.WorkoutLog {
	return predicate.WorkoutLog(func(s *sql.Selector) {
		step := newExercisesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WorkoutLog) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WorkoutLog) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WorkoutLog) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.NotPredicates(p))
}
