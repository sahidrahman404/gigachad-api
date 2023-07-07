// Code generated by ent, DO NOT EDIT.

package equipment

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

// ID filters vertices based on their ID field.
func ID(id pksuid.ID) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pksuid.ID) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pksuid.ID) predicate.Equipment {
	return predicate.Equipment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pksuid.ID) predicate.Equipment {
	return predicate.Equipment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pksuid.ID) predicate.Equipment {
	return predicate.Equipment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pksuid.ID) predicate.Equipment {
	return predicate.Equipment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pksuid.ID) predicate.Equipment {
	return predicate.Equipment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pksuid.ID) predicate.Equipment {
	return predicate.Equipment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pksuid.ID) predicate.Equipment {
	return predicate.Equipment(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldName, v))
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldImage, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Equipment {
	return predicate.Equipment(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Equipment {
	return predicate.Equipment(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldContainsFold(FieldName, v))
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEQ(FieldImage, v))
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldNEQ(FieldImage, v))
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.Equipment {
	return predicate.Equipment(sql.FieldIn(FieldImage, vs...))
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.Equipment {
	return predicate.Equipment(sql.FieldNotIn(FieldImage, vs...))
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldGT(FieldImage, v))
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldGTE(FieldImage, v))
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldLT(FieldImage, v))
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldLTE(FieldImage, v))
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldContains(FieldImage, v))
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldHasPrefix(FieldImage, v))
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldHasSuffix(FieldImage, v))
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldEqualFold(FieldImage, v))
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.Equipment {
	return predicate.Equipment(sql.FieldContainsFold(FieldImage, v))
}

// HasExercises applies the HasEdge predicate on the "exercises" edge.
func HasExercises() predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ExercisesTable, ExercisesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExercisesWith applies the HasEdge predicate on the "exercises" edge with a given conditions (other predicates).
func HasExercisesWith(preds ...predicate.Exercise) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		step := newExercisesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Equipment) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Equipment) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
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
func Not(p predicate.Equipment) predicate.Equipment {
	return predicate.Equipment(func(s *sql.Selector) {
		p(s.Not())
	})
}
