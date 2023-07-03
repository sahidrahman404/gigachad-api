// Code generated by ent, DO NOT EDIT.

package routineexercise

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldContainsFold(FieldID, id))
}

// RestTimer applies equality check predicate on the "rest_timer" field. It's identical to RestTimerEQ.
func RestTimer(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldRestTimer, v))
}

// RoutineID applies equality check predicate on the "routine_id" field. It's identical to RoutineIDEQ.
func RoutineID(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldRoutineID, v))
}

// ExerciseID applies equality check predicate on the "exercise_id" field. It's identical to ExerciseIDEQ.
func ExerciseID(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldExerciseID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldUserID, v))
}

// RestTimerEQ applies the EQ predicate on the "rest_timer" field.
func RestTimerEQ(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldRestTimer, v))
}

// RestTimerNEQ applies the NEQ predicate on the "rest_timer" field.
func RestTimerNEQ(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNEQ(FieldRestTimer, v))
}

// RestTimerIn applies the In predicate on the "rest_timer" field.
func RestTimerIn(vs ...int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIn(FieldRestTimer, vs...))
}

// RestTimerNotIn applies the NotIn predicate on the "rest_timer" field.
func RestTimerNotIn(vs ...int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotIn(FieldRestTimer, vs...))
}

// RestTimerGT applies the GT predicate on the "rest_timer" field.
func RestTimerGT(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGT(FieldRestTimer, v))
}

// RestTimerGTE applies the GTE predicate on the "rest_timer" field.
func RestTimerGTE(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGTE(FieldRestTimer, v))
}

// RestTimerLT applies the LT predicate on the "rest_timer" field.
func RestTimerLT(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLT(FieldRestTimer, v))
}

// RestTimerLTE applies the LTE predicate on the "rest_timer" field.
func RestTimerLTE(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLTE(FieldRestTimer, v))
}

// RestTimerIsNil applies the IsNil predicate on the "rest_timer" field.
func RestTimerIsNil() predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIsNull(FieldRestTimer))
}

// RestTimerNotNil applies the NotNil predicate on the "rest_timer" field.
func RestTimerNotNil() predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotNull(FieldRestTimer))
}

// RoutineIDEQ applies the EQ predicate on the "routine_id" field.
func RoutineIDEQ(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldRoutineID, v))
}

// RoutineIDNEQ applies the NEQ predicate on the "routine_id" field.
func RoutineIDNEQ(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNEQ(FieldRoutineID, v))
}

// RoutineIDIn applies the In predicate on the "routine_id" field.
func RoutineIDIn(vs ...string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIn(FieldRoutineID, vs...))
}

// RoutineIDNotIn applies the NotIn predicate on the "routine_id" field.
func RoutineIDNotIn(vs ...string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotIn(FieldRoutineID, vs...))
}

// RoutineIDGT applies the GT predicate on the "routine_id" field.
func RoutineIDGT(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGT(FieldRoutineID, v))
}

// RoutineIDGTE applies the GTE predicate on the "routine_id" field.
func RoutineIDGTE(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGTE(FieldRoutineID, v))
}

// RoutineIDLT applies the LT predicate on the "routine_id" field.
func RoutineIDLT(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLT(FieldRoutineID, v))
}

// RoutineIDLTE applies the LTE predicate on the "routine_id" field.
func RoutineIDLTE(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLTE(FieldRoutineID, v))
}

// RoutineIDContains applies the Contains predicate on the "routine_id" field.
func RoutineIDContains(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldContains(FieldRoutineID, v))
}

// RoutineIDHasPrefix applies the HasPrefix predicate on the "routine_id" field.
func RoutineIDHasPrefix(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldHasPrefix(FieldRoutineID, v))
}

// RoutineIDHasSuffix applies the HasSuffix predicate on the "routine_id" field.
func RoutineIDHasSuffix(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldHasSuffix(FieldRoutineID, v))
}

// RoutineIDIsNil applies the IsNil predicate on the "routine_id" field.
func RoutineIDIsNil() predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIsNull(FieldRoutineID))
}

// RoutineIDNotNil applies the NotNil predicate on the "routine_id" field.
func RoutineIDNotNil() predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotNull(FieldRoutineID))
}

// RoutineIDEqualFold applies the EqualFold predicate on the "routine_id" field.
func RoutineIDEqualFold(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEqualFold(FieldRoutineID, v))
}

// RoutineIDContainsFold applies the ContainsFold predicate on the "routine_id" field.
func RoutineIDContainsFold(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldContainsFold(FieldRoutineID, v))
}

// ExerciseIDEQ applies the EQ predicate on the "exercise_id" field.
func ExerciseIDEQ(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldExerciseID, v))
}

// ExerciseIDNEQ applies the NEQ predicate on the "exercise_id" field.
func ExerciseIDNEQ(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNEQ(FieldExerciseID, v))
}

// ExerciseIDIn applies the In predicate on the "exercise_id" field.
func ExerciseIDIn(vs ...string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIn(FieldExerciseID, vs...))
}

// ExerciseIDNotIn applies the NotIn predicate on the "exercise_id" field.
func ExerciseIDNotIn(vs ...string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotIn(FieldExerciseID, vs...))
}

// ExerciseIDGT applies the GT predicate on the "exercise_id" field.
func ExerciseIDGT(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGT(FieldExerciseID, v))
}

// ExerciseIDGTE applies the GTE predicate on the "exercise_id" field.
func ExerciseIDGTE(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGTE(FieldExerciseID, v))
}

// ExerciseIDLT applies the LT predicate on the "exercise_id" field.
func ExerciseIDLT(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLT(FieldExerciseID, v))
}

// ExerciseIDLTE applies the LTE predicate on the "exercise_id" field.
func ExerciseIDLTE(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLTE(FieldExerciseID, v))
}

// ExerciseIDContains applies the Contains predicate on the "exercise_id" field.
func ExerciseIDContains(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldContains(FieldExerciseID, v))
}

// ExerciseIDHasPrefix applies the HasPrefix predicate on the "exercise_id" field.
func ExerciseIDHasPrefix(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldHasPrefix(FieldExerciseID, v))
}

// ExerciseIDHasSuffix applies the HasSuffix predicate on the "exercise_id" field.
func ExerciseIDHasSuffix(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldHasSuffix(FieldExerciseID, v))
}

// ExerciseIDIsNil applies the IsNil predicate on the "exercise_id" field.
func ExerciseIDIsNil() predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIsNull(FieldExerciseID))
}

// ExerciseIDNotNil applies the NotNil predicate on the "exercise_id" field.
func ExerciseIDNotNil() predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotNull(FieldExerciseID))
}

// ExerciseIDEqualFold applies the EqualFold predicate on the "exercise_id" field.
func ExerciseIDEqualFold(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEqualFold(FieldExerciseID, v))
}

// ExerciseIDContainsFold applies the ContainsFold predicate on the "exercise_id" field.
func ExerciseIDContainsFold(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldContainsFold(FieldExerciseID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotNull(FieldUserID))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldContainsFold(FieldUserID, v))
}

// HasRoutines applies the HasEdge predicate on the "routines" edge.
func HasRoutines() predicate.RoutineExercise {
	return predicate.RoutineExercise(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoutinesTable, RoutinesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoutinesWith applies the HasEdge predicate on the "routines" edge with a given conditions (other predicates).
func HasRoutinesWith(preds ...predicate.Routine) predicate.RoutineExercise {
	return predicate.RoutineExercise(func(s *sql.Selector) {
		step := newRoutinesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasExercises applies the HasEdge predicate on the "exercises" edge.
func HasExercises() predicate.RoutineExercise {
	return predicate.RoutineExercise(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ExercisesTable, ExercisesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExercisesWith applies the HasEdge predicate on the "exercises" edge with a given conditions (other predicates).
func HasExercisesWith(preds ...predicate.Exercise) predicate.RoutineExercise {
	return predicate.RoutineExercise(func(s *sql.Selector) {
		step := newExercisesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.RoutineExercise {
	return predicate.RoutineExercise(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.RoutineExercise {
	return predicate.RoutineExercise(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RoutineExercise) predicate.RoutineExercise {
	return predicate.RoutineExercise(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RoutineExercise) predicate.RoutineExercise {
	return predicate.RoutineExercise(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RoutineExercise) predicate.RoutineExercise {
	return predicate.RoutineExercise(func(s *sql.Selector) {
		p(s.Not())
	})
}