// Code generated by ent, DO NOT EDIT.

package token

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Token {
	return predicate.Token(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Token {
	return predicate.Token(sql.FieldContainsFold(FieldID, id))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v []byte) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldHash, v))
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldExpiry, v))
}

// Scope applies equality check predicate on the "scope" field. It's identical to ScopeEQ.
func Scope(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldScope, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldUserID, v))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v []byte) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v []byte) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...[]byte) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...[]byte) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v []byte) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v []byte) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v []byte) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v []byte) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldHash, v))
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldExpiry, v))
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldExpiry, v))
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldExpiry, vs...))
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldExpiry, vs...))
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v string) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldExpiry, v))
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v string) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldExpiry, v))
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v string) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldExpiry, v))
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v string) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldExpiry, v))
}

// ExpiryContains applies the Contains predicate on the "expiry" field.
func ExpiryContains(v string) predicate.Token {
	return predicate.Token(sql.FieldContains(FieldExpiry, v))
}

// ExpiryHasPrefix applies the HasPrefix predicate on the "expiry" field.
func ExpiryHasPrefix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasPrefix(FieldExpiry, v))
}

// ExpiryHasSuffix applies the HasSuffix predicate on the "expiry" field.
func ExpiryHasSuffix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasSuffix(FieldExpiry, v))
}

// ExpiryEqualFold applies the EqualFold predicate on the "expiry" field.
func ExpiryEqualFold(v string) predicate.Token {
	return predicate.Token(sql.FieldEqualFold(FieldExpiry, v))
}

// ExpiryContainsFold applies the ContainsFold predicate on the "expiry" field.
func ExpiryContainsFold(v string) predicate.Token {
	return predicate.Token(sql.FieldContainsFold(FieldExpiry, v))
}

// ScopeEQ applies the EQ predicate on the "scope" field.
func ScopeEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldScope, v))
}

// ScopeNEQ applies the NEQ predicate on the "scope" field.
func ScopeNEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldScope, v))
}

// ScopeIn applies the In predicate on the "scope" field.
func ScopeIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldScope, vs...))
}

// ScopeNotIn applies the NotIn predicate on the "scope" field.
func ScopeNotIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldScope, vs...))
}

// ScopeGT applies the GT predicate on the "scope" field.
func ScopeGT(v string) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldScope, v))
}

// ScopeGTE applies the GTE predicate on the "scope" field.
func ScopeGTE(v string) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldScope, v))
}

// ScopeLT applies the LT predicate on the "scope" field.
func ScopeLT(v string) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldScope, v))
}

// ScopeLTE applies the LTE predicate on the "scope" field.
func ScopeLTE(v string) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldScope, v))
}

// ScopeContains applies the Contains predicate on the "scope" field.
func ScopeContains(v string) predicate.Token {
	return predicate.Token(sql.FieldContains(FieldScope, v))
}

// ScopeHasPrefix applies the HasPrefix predicate on the "scope" field.
func ScopeHasPrefix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasPrefix(FieldScope, v))
}

// ScopeHasSuffix applies the HasSuffix predicate on the "scope" field.
func ScopeHasSuffix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasSuffix(FieldScope, v))
}

// ScopeEqualFold applies the EqualFold predicate on the "scope" field.
func ScopeEqualFold(v string) predicate.Token {
	return predicate.Token(sql.FieldEqualFold(FieldScope, v))
}

// ScopeContainsFold applies the ContainsFold predicate on the "scope" field.
func ScopeContainsFold(v string) predicate.Token {
	return predicate.Token(sql.FieldContainsFold(FieldScope, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Token {
	return predicate.Token(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Token {
	return predicate.Token(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Token {
	return predicate.Token(sql.FieldNotNull(FieldUserID))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Token {
	return predicate.Token(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Token {
	return predicate.Token(sql.FieldContainsFold(FieldUserID, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Token) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Token) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
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
func Not(p predicate.Token) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		p(s.Not())
	})
}
