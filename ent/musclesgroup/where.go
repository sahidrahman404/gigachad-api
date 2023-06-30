// Code generated by ent, DO NOT EDIT.

package musclesgroup

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldContainsFold(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldEQ(FieldName, v))
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldEQ(FieldImage, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldContainsFold(FieldName, v))
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldEQ(FieldImage, v))
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldNEQ(FieldImage, v))
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldIn(FieldImage, vs...))
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldNotIn(FieldImage, vs...))
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldGT(FieldImage, v))
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldGTE(FieldImage, v))
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldLT(FieldImage, v))
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldLTE(FieldImage, v))
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldContains(FieldImage, v))
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldHasPrefix(FieldImage, v))
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldHasSuffix(FieldImage, v))
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldEqualFold(FieldImage, v))
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.MusclesGroup {
	return predicate.MusclesGroup(sql.FieldContainsFold(FieldImage, v))
}

// HasExercises applies the HasEdge predicate on the "exercises" edge.
func HasExercises() predicate.MusclesGroup {
	return predicate.MusclesGroup(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ExercisesTable, ExercisesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExercisesWith applies the HasEdge predicate on the "exercises" edge with a given conditions (other predicates).
func HasExercisesWith(preds ...predicate.Exercise) predicate.MusclesGroup {
	return predicate.MusclesGroup(func(s *sql.Selector) {
		step := newExercisesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MusclesGroup) predicate.MusclesGroup {
	return predicate.MusclesGroup(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MusclesGroup) predicate.MusclesGroup {
	return predicate.MusclesGroup(func(s *sql.Selector) {
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
func Not(p predicate.MusclesGroup) predicate.MusclesGroup {
	return predicate.MusclesGroup(func(s *sql.Selector) {
		p(s.Not())
	})
}
