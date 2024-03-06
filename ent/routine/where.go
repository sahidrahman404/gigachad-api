// Code generated by ent, DO NOT EDIT.

package routine

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

// ID filters vertices based on their ID field.
func ID(id pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Routine {
	return predicate.Routine(sql.FieldEQ(FieldName, v))
}

// ScheduleID applies equality check predicate on the "schedule_id" field. It's identical to ScheduleIDEQ.
func ScheduleID(v string) predicate.Routine {
	return predicate.Routine(sql.FieldEQ(FieldScheduleID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldEQ(FieldUserID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Routine {
	return predicate.Routine(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Routine {
	return predicate.Routine(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Routine {
	return predicate.Routine(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Routine {
	return predicate.Routine(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Routine {
	return predicate.Routine(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Routine {
	return predicate.Routine(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Routine {
	return predicate.Routine(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Routine {
	return predicate.Routine(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Routine {
	return predicate.Routine(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Routine {
	return predicate.Routine(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Routine {
	return predicate.Routine(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Routine {
	return predicate.Routine(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Routine {
	return predicate.Routine(sql.FieldContainsFold(FieldName, v))
}

// ScheduleIDEQ applies the EQ predicate on the "schedule_id" field.
func ScheduleIDEQ(v string) predicate.Routine {
	return predicate.Routine(sql.FieldEQ(FieldScheduleID, v))
}

// ScheduleIDNEQ applies the NEQ predicate on the "schedule_id" field.
func ScheduleIDNEQ(v string) predicate.Routine {
	return predicate.Routine(sql.FieldNEQ(FieldScheduleID, v))
}

// ScheduleIDIn applies the In predicate on the "schedule_id" field.
func ScheduleIDIn(vs ...string) predicate.Routine {
	return predicate.Routine(sql.FieldIn(FieldScheduleID, vs...))
}

// ScheduleIDNotIn applies the NotIn predicate on the "schedule_id" field.
func ScheduleIDNotIn(vs ...string) predicate.Routine {
	return predicate.Routine(sql.FieldNotIn(FieldScheduleID, vs...))
}

// ScheduleIDGT applies the GT predicate on the "schedule_id" field.
func ScheduleIDGT(v string) predicate.Routine {
	return predicate.Routine(sql.FieldGT(FieldScheduleID, v))
}

// ScheduleIDGTE applies the GTE predicate on the "schedule_id" field.
func ScheduleIDGTE(v string) predicate.Routine {
	return predicate.Routine(sql.FieldGTE(FieldScheduleID, v))
}

// ScheduleIDLT applies the LT predicate on the "schedule_id" field.
func ScheduleIDLT(v string) predicate.Routine {
	return predicate.Routine(sql.FieldLT(FieldScheduleID, v))
}

// ScheduleIDLTE applies the LTE predicate on the "schedule_id" field.
func ScheduleIDLTE(v string) predicate.Routine {
	return predicate.Routine(sql.FieldLTE(FieldScheduleID, v))
}

// ScheduleIDContains applies the Contains predicate on the "schedule_id" field.
func ScheduleIDContains(v string) predicate.Routine {
	return predicate.Routine(sql.FieldContains(FieldScheduleID, v))
}

// ScheduleIDHasPrefix applies the HasPrefix predicate on the "schedule_id" field.
func ScheduleIDHasPrefix(v string) predicate.Routine {
	return predicate.Routine(sql.FieldHasPrefix(FieldScheduleID, v))
}

// ScheduleIDHasSuffix applies the HasSuffix predicate on the "schedule_id" field.
func ScheduleIDHasSuffix(v string) predicate.Routine {
	return predicate.Routine(sql.FieldHasSuffix(FieldScheduleID, v))
}

// ScheduleIDIsNil applies the IsNil predicate on the "schedule_id" field.
func ScheduleIDIsNil() predicate.Routine {
	return predicate.Routine(sql.FieldIsNull(FieldScheduleID))
}

// ScheduleIDNotNil applies the NotNil predicate on the "schedule_id" field.
func ScheduleIDNotNil() predicate.Routine {
	return predicate.Routine(sql.FieldNotNull(FieldScheduleID))
}

// ScheduleIDEqualFold applies the EqualFold predicate on the "schedule_id" field.
func ScheduleIDEqualFold(v string) predicate.Routine {
	return predicate.Routine(sql.FieldEqualFold(FieldScheduleID, v))
}

// ScheduleIDContainsFold applies the ContainsFold predicate on the "schedule_id" field.
func ScheduleIDContainsFold(v string) predicate.Routine {
	return predicate.Routine(sql.FieldContainsFold(FieldScheduleID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v pksuid.ID) predicate.Routine {
	return predicate.Routine(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v pksuid.ID) predicate.Routine {
	vc := string(v)
	return predicate.Routine(sql.FieldContains(FieldUserID, vc))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v pksuid.ID) predicate.Routine {
	vc := string(v)
	return predicate.Routine(sql.FieldHasPrefix(FieldUserID, vc))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v pksuid.ID) predicate.Routine {
	vc := string(v)
	return predicate.Routine(sql.FieldHasSuffix(FieldUserID, vc))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v pksuid.ID) predicate.Routine {
	vc := string(v)
	return predicate.Routine(sql.FieldEqualFold(FieldUserID, vc))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v pksuid.ID) predicate.Routine {
	vc := string(v)
	return predicate.Routine(sql.FieldContainsFold(FieldUserID, vc))
}

// HasExercises applies the HasEdge predicate on the "exercises" edge.
func HasExercises() predicate.Routine {
	return predicate.Routine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ExercisesTable, ExercisesPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExercisesWith applies the HasEdge predicate on the "exercises" edge with a given conditions (other predicates).
func HasExercisesWith(preds ...predicate.Exercise) predicate.Routine {
	return predicate.Routine(func(s *sql.Selector) {
		step := newExercisesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Routine {
	return predicate.Routine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Routine {
	return predicate.Routine(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRoutineExercises applies the HasEdge predicate on the "routine_exercises" edge.
func HasRoutineExercises() predicate.Routine {
	return predicate.Routine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, RoutineExercisesTable, RoutineExercisesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoutineExercisesWith applies the HasEdge predicate on the "routine_exercises" edge with a given conditions (other predicates).
func HasRoutineExercisesWith(preds ...predicate.RoutineExercise) predicate.Routine {
	return predicate.Routine(func(s *sql.Selector) {
		step := newRoutineExercisesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Routine) predicate.Routine {
	return predicate.Routine(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Routine) predicate.Routine {
	return predicate.Routine(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Routine) predicate.Routine {
	return predicate.Routine(sql.NotPredicates(p))
}
