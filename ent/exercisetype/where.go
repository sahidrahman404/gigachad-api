// Code generated by ent, DO NOT EDIT.

package exercisetype

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldContainsFold(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldEQ(FieldDescription, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ExerciseType {
	return predicate.ExerciseType(sql.FieldContainsFold(FieldDescription, v))
}

// HasExercises applies the HasEdge predicate on the "exercises" edge.
func HasExercises() predicate.ExerciseType {
	return predicate.ExerciseType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ExercisesTable, ExercisesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExercisesWith applies the HasEdge predicate on the "exercises" edge with a given conditions (other predicates).
func HasExercisesWith(preds ...predicate.Exercise) predicate.ExerciseType {
	return predicate.ExerciseType(func(s *sql.Selector) {
		step := newExercisesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ExerciseType) predicate.ExerciseType {
	return predicate.ExerciseType(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ExerciseType) predicate.ExerciseType {
	return predicate.ExerciseType(func(s *sql.Selector) {
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
func Not(p predicate.ExerciseType) predicate.ExerciseType {
	return predicate.ExerciseType(func(s *sql.Selector) {
		p(s.Not())
	})
}
