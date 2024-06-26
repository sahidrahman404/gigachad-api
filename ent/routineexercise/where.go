// Code generated by ent, DO NOT EDIT.

package routineexercise

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

// ID filters vertices based on their ID field.
func ID(id pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLTE(FieldID, id))
}

// RestTime applies equality check predicate on the "rest_time" field. It's identical to RestTimeEQ.
func RestTime(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldRestTime, v))
}

// RoutineID applies equality check predicate on the "routine_id" field. It's identical to RoutineIDEQ.
func RoutineID(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldRoutineID, v))
}

// ExerciseID applies equality check predicate on the "exercise_id" field. It's identical to ExerciseIDEQ.
func ExerciseID(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldExerciseID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldUserID, v))
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldOrder, v))
}

// RestTimeEQ applies the EQ predicate on the "rest_time" field.
func RestTimeEQ(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldRestTime, v))
}

// RestTimeNEQ applies the NEQ predicate on the "rest_time" field.
func RestTimeNEQ(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNEQ(FieldRestTime, v))
}

// RestTimeIn applies the In predicate on the "rest_time" field.
func RestTimeIn(vs ...string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIn(FieldRestTime, vs...))
}

// RestTimeNotIn applies the NotIn predicate on the "rest_time" field.
func RestTimeNotIn(vs ...string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotIn(FieldRestTime, vs...))
}

// RestTimeGT applies the GT predicate on the "rest_time" field.
func RestTimeGT(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGT(FieldRestTime, v))
}

// RestTimeGTE applies the GTE predicate on the "rest_time" field.
func RestTimeGTE(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGTE(FieldRestTime, v))
}

// RestTimeLT applies the LT predicate on the "rest_time" field.
func RestTimeLT(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLT(FieldRestTime, v))
}

// RestTimeLTE applies the LTE predicate on the "rest_time" field.
func RestTimeLTE(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLTE(FieldRestTime, v))
}

// RestTimeContains applies the Contains predicate on the "rest_time" field.
func RestTimeContains(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldContains(FieldRestTime, v))
}

// RestTimeHasPrefix applies the HasPrefix predicate on the "rest_time" field.
func RestTimeHasPrefix(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldHasPrefix(FieldRestTime, v))
}

// RestTimeHasSuffix applies the HasSuffix predicate on the "rest_time" field.
func RestTimeHasSuffix(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldHasSuffix(FieldRestTime, v))
}

// RestTimeIsNil applies the IsNil predicate on the "rest_time" field.
func RestTimeIsNil() predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIsNull(FieldRestTime))
}

// RestTimeNotNil applies the NotNil predicate on the "rest_time" field.
func RestTimeNotNil() predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotNull(FieldRestTime))
}

// RestTimeEqualFold applies the EqualFold predicate on the "rest_time" field.
func RestTimeEqualFold(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEqualFold(FieldRestTime, v))
}

// RestTimeContainsFold applies the ContainsFold predicate on the "rest_time" field.
func RestTimeContainsFold(v string) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldContainsFold(FieldRestTime, v))
}

// RoutineIDEQ applies the EQ predicate on the "routine_id" field.
func RoutineIDEQ(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldRoutineID, v))
}

// RoutineIDNEQ applies the NEQ predicate on the "routine_id" field.
func RoutineIDNEQ(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNEQ(FieldRoutineID, v))
}

// RoutineIDIn applies the In predicate on the "routine_id" field.
func RoutineIDIn(vs ...pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIn(FieldRoutineID, vs...))
}

// RoutineIDNotIn applies the NotIn predicate on the "routine_id" field.
func RoutineIDNotIn(vs ...pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotIn(FieldRoutineID, vs...))
}

// RoutineIDGT applies the GT predicate on the "routine_id" field.
func RoutineIDGT(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGT(FieldRoutineID, v))
}

// RoutineIDGTE applies the GTE predicate on the "routine_id" field.
func RoutineIDGTE(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGTE(FieldRoutineID, v))
}

// RoutineIDLT applies the LT predicate on the "routine_id" field.
func RoutineIDLT(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLT(FieldRoutineID, v))
}

// RoutineIDLTE applies the LTE predicate on the "routine_id" field.
func RoutineIDLTE(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLTE(FieldRoutineID, v))
}

// RoutineIDContains applies the Contains predicate on the "routine_id" field.
func RoutineIDContains(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldContains(FieldRoutineID, vc))
}

// RoutineIDHasPrefix applies the HasPrefix predicate on the "routine_id" field.
func RoutineIDHasPrefix(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldHasPrefix(FieldRoutineID, vc))
}

// RoutineIDHasSuffix applies the HasSuffix predicate on the "routine_id" field.
func RoutineIDHasSuffix(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldHasSuffix(FieldRoutineID, vc))
}

// RoutineIDEqualFold applies the EqualFold predicate on the "routine_id" field.
func RoutineIDEqualFold(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldEqualFold(FieldRoutineID, vc))
}

// RoutineIDContainsFold applies the ContainsFold predicate on the "routine_id" field.
func RoutineIDContainsFold(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldContainsFold(FieldRoutineID, vc))
}

// ExerciseIDEQ applies the EQ predicate on the "exercise_id" field.
func ExerciseIDEQ(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldExerciseID, v))
}

// ExerciseIDNEQ applies the NEQ predicate on the "exercise_id" field.
func ExerciseIDNEQ(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNEQ(FieldExerciseID, v))
}

// ExerciseIDIn applies the In predicate on the "exercise_id" field.
func ExerciseIDIn(vs ...pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIn(FieldExerciseID, vs...))
}

// ExerciseIDNotIn applies the NotIn predicate on the "exercise_id" field.
func ExerciseIDNotIn(vs ...pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotIn(FieldExerciseID, vs...))
}

// ExerciseIDGT applies the GT predicate on the "exercise_id" field.
func ExerciseIDGT(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGT(FieldExerciseID, v))
}

// ExerciseIDGTE applies the GTE predicate on the "exercise_id" field.
func ExerciseIDGTE(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGTE(FieldExerciseID, v))
}

// ExerciseIDLT applies the LT predicate on the "exercise_id" field.
func ExerciseIDLT(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLT(FieldExerciseID, v))
}

// ExerciseIDLTE applies the LTE predicate on the "exercise_id" field.
func ExerciseIDLTE(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLTE(FieldExerciseID, v))
}

// ExerciseIDContains applies the Contains predicate on the "exercise_id" field.
func ExerciseIDContains(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldContains(FieldExerciseID, vc))
}

// ExerciseIDHasPrefix applies the HasPrefix predicate on the "exercise_id" field.
func ExerciseIDHasPrefix(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldHasPrefix(FieldExerciseID, vc))
}

// ExerciseIDHasSuffix applies the HasSuffix predicate on the "exercise_id" field.
func ExerciseIDHasSuffix(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldHasSuffix(FieldExerciseID, vc))
}

// ExerciseIDEqualFold applies the EqualFold predicate on the "exercise_id" field.
func ExerciseIDEqualFold(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldEqualFold(FieldExerciseID, vc))
}

// ExerciseIDContainsFold applies the ContainsFold predicate on the "exercise_id" field.
func ExerciseIDContainsFold(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldContainsFold(FieldExerciseID, vc))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v pksuid.ID) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldContains(FieldUserID, vc))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldHasPrefix(FieldUserID, vc))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldHasSuffix(FieldUserID, vc))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldEqualFold(FieldUserID, vc))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v pksuid.ID) predicate.RoutineExercise {
	vc := string(v)
	return predicate.RoutineExercise(sql.FieldContainsFold(FieldUserID, vc))
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldEQ(FieldOrder, v))
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNEQ(FieldOrder, v))
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldIn(FieldOrder, vs...))
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldNotIn(FieldOrder, vs...))
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGT(FieldOrder, v))
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldGTE(FieldOrder, v))
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLT(FieldOrder, v))
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v int) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.FieldLTE(FieldOrder, v))
}

// HasRoutines applies the HasEdge predicate on the "routines" edge.
func HasRoutines() predicate.RoutineExercise {
	return predicate.RoutineExercise(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RoutinesTable, RoutinesColumn),
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
			sqlgraph.Edge(sqlgraph.M2O, false, ExercisesTable, ExercisesColumn),
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
	return predicate.RoutineExercise(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RoutineExercise) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RoutineExercise) predicate.RoutineExercise {
	return predicate.RoutineExercise(sql.NotPredicates(p))
}
