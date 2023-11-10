// Code generated by ent, DO NOT EDIT.

package workoutlog

import (
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
func CreatedAt(v string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEQ(FieldCreatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v pksuid.ID) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEQ(FieldUserID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtContains applies the Contains predicate on the "created_at" field.
func CreatedAtContains(v string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldContains(FieldCreatedAt, v))
}

// CreatedAtHasPrefix applies the HasPrefix predicate on the "created_at" field.
func CreatedAtHasPrefix(v string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldHasPrefix(FieldCreatedAt, v))
}

// CreatedAtHasSuffix applies the HasSuffix predicate on the "created_at" field.
func CreatedAtHasSuffix(v string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldHasSuffix(FieldCreatedAt, v))
}

// CreatedAtEqualFold applies the EqualFold predicate on the "created_at" field.
func CreatedAtEqualFold(v string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldEqualFold(FieldCreatedAt, v))
}

// CreatedAtContainsFold applies the ContainsFold predicate on the "created_at" field.
func CreatedAtContainsFold(v string) predicate.WorkoutLog {
	return predicate.WorkoutLog(sql.FieldContainsFold(FieldCreatedAt, v))
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

// HasExercises applies the HasEdge predicate on the "exercises" edge.
func HasExercises() predicate.WorkoutLog {
	return predicate.WorkoutLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ExercisesTable, ExercisesColumn),
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

// HasWorkouts applies the HasEdge predicate on the "workouts" edge.
func HasWorkouts() predicate.WorkoutLog {
	return predicate.WorkoutLog(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, WorkoutsTable, WorkoutsColumn),
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

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WorkoutLog) predicate.WorkoutLog {
	return predicate.WorkoutLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WorkoutLog) predicate.WorkoutLog {
	return predicate.WorkoutLog(func(s *sql.Selector) {
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
func Not(p predicate.WorkoutLog) predicate.WorkoutLog {
	return predicate.WorkoutLog(func(s *sql.Selector) {
		p(s.Not())
	})
}
