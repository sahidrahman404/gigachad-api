// Code generated by ent, DO NOT EDIT.

package workout

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
)

// ID filters vertices based on their ID field.
func ID(id pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldName, v))
}

// Volume applies equality check predicate on the "volume" field. It's identical to VolumeEQ.
func Volume(v int) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldVolume, v))
}

// Reps applies equality check predicate on the "reps" field. It's identical to RepsEQ.
func Reps(v int) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldReps, v))
}

// Time applies equality check predicate on the "time" field. It's identical to TimeEQ.
func Time(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldTime, v))
}

// Sets applies equality check predicate on the "sets" field. It's identical to SetsEQ.
func Sets(v int) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldSets, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldCreatedAt, v))
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldImage, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldDescription, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldUserID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Workout {
	return predicate.Workout(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Workout {
	return predicate.Workout(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Workout {
	return predicate.Workout(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Workout {
	return predicate.Workout(sql.FieldContainsFold(FieldName, v))
}

// VolumeEQ applies the EQ predicate on the "volume" field.
func VolumeEQ(v int) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldVolume, v))
}

// VolumeNEQ applies the NEQ predicate on the "volume" field.
func VolumeNEQ(v int) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldVolume, v))
}

// VolumeIn applies the In predicate on the "volume" field.
func VolumeIn(vs ...int) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldVolume, vs...))
}

// VolumeNotIn applies the NotIn predicate on the "volume" field.
func VolumeNotIn(vs ...int) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldVolume, vs...))
}

// VolumeGT applies the GT predicate on the "volume" field.
func VolumeGT(v int) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldVolume, v))
}

// VolumeGTE applies the GTE predicate on the "volume" field.
func VolumeGTE(v int) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldVolume, v))
}

// VolumeLT applies the LT predicate on the "volume" field.
func VolumeLT(v int) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldVolume, v))
}

// VolumeLTE applies the LTE predicate on the "volume" field.
func VolumeLTE(v int) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldVolume, v))
}

// RepsEQ applies the EQ predicate on the "reps" field.
func RepsEQ(v int) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldReps, v))
}

// RepsNEQ applies the NEQ predicate on the "reps" field.
func RepsNEQ(v int) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldReps, v))
}

// RepsIn applies the In predicate on the "reps" field.
func RepsIn(vs ...int) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldReps, vs...))
}

// RepsNotIn applies the NotIn predicate on the "reps" field.
func RepsNotIn(vs ...int) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldReps, vs...))
}

// RepsGT applies the GT predicate on the "reps" field.
func RepsGT(v int) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldReps, v))
}

// RepsGTE applies the GTE predicate on the "reps" field.
func RepsGTE(v int) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldReps, v))
}

// RepsLT applies the LT predicate on the "reps" field.
func RepsLT(v int) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldReps, v))
}

// RepsLTE applies the LTE predicate on the "reps" field.
func RepsLTE(v int) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldReps, v))
}

// TimeEQ applies the EQ predicate on the "time" field.
func TimeEQ(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldTime, v))
}

// TimeNEQ applies the NEQ predicate on the "time" field.
func TimeNEQ(v string) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldTime, v))
}

// TimeIn applies the In predicate on the "time" field.
func TimeIn(vs ...string) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldTime, vs...))
}

// TimeNotIn applies the NotIn predicate on the "time" field.
func TimeNotIn(vs ...string) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldTime, vs...))
}

// TimeGT applies the GT predicate on the "time" field.
func TimeGT(v string) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldTime, v))
}

// TimeGTE applies the GTE predicate on the "time" field.
func TimeGTE(v string) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldTime, v))
}

// TimeLT applies the LT predicate on the "time" field.
func TimeLT(v string) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldTime, v))
}

// TimeLTE applies the LTE predicate on the "time" field.
func TimeLTE(v string) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldTime, v))
}

// TimeContains applies the Contains predicate on the "time" field.
func TimeContains(v string) predicate.Workout {
	return predicate.Workout(sql.FieldContains(FieldTime, v))
}

// TimeHasPrefix applies the HasPrefix predicate on the "time" field.
func TimeHasPrefix(v string) predicate.Workout {
	return predicate.Workout(sql.FieldHasPrefix(FieldTime, v))
}

// TimeHasSuffix applies the HasSuffix predicate on the "time" field.
func TimeHasSuffix(v string) predicate.Workout {
	return predicate.Workout(sql.FieldHasSuffix(FieldTime, v))
}

// TimeIsNil applies the IsNil predicate on the "time" field.
func TimeIsNil() predicate.Workout {
	return predicate.Workout(sql.FieldIsNull(FieldTime))
}

// TimeNotNil applies the NotNil predicate on the "time" field.
func TimeNotNil() predicate.Workout {
	return predicate.Workout(sql.FieldNotNull(FieldTime))
}

// TimeEqualFold applies the EqualFold predicate on the "time" field.
func TimeEqualFold(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEqualFold(FieldTime, v))
}

// TimeContainsFold applies the ContainsFold predicate on the "time" field.
func TimeContainsFold(v string) predicate.Workout {
	return predicate.Workout(sql.FieldContainsFold(FieldTime, v))
}

// SetsEQ applies the EQ predicate on the "sets" field.
func SetsEQ(v int) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldSets, v))
}

// SetsNEQ applies the NEQ predicate on the "sets" field.
func SetsNEQ(v int) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldSets, v))
}

// SetsIn applies the In predicate on the "sets" field.
func SetsIn(vs ...int) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldSets, vs...))
}

// SetsNotIn applies the NotIn predicate on the "sets" field.
func SetsNotIn(vs ...int) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldSets, vs...))
}

// SetsGT applies the GT predicate on the "sets" field.
func SetsGT(v int) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldSets, v))
}

// SetsGTE applies the GTE predicate on the "sets" field.
func SetsGTE(v int) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldSets, v))
}

// SetsLT applies the LT predicate on the "sets" field.
func SetsLT(v int) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldSets, v))
}

// SetsLTE applies the LTE predicate on the "sets" field.
func SetsLTE(v int) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldSets, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v string) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...string) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...string) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v string) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v string) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v string) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v string) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtContains applies the Contains predicate on the "created_at" field.
func CreatedAtContains(v string) predicate.Workout {
	return predicate.Workout(sql.FieldContains(FieldCreatedAt, v))
}

// CreatedAtHasPrefix applies the HasPrefix predicate on the "created_at" field.
func CreatedAtHasPrefix(v string) predicate.Workout {
	return predicate.Workout(sql.FieldHasPrefix(FieldCreatedAt, v))
}

// CreatedAtHasSuffix applies the HasSuffix predicate on the "created_at" field.
func CreatedAtHasSuffix(v string) predicate.Workout {
	return predicate.Workout(sql.FieldHasSuffix(FieldCreatedAt, v))
}

// CreatedAtEqualFold applies the EqualFold predicate on the "created_at" field.
func CreatedAtEqualFold(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEqualFold(FieldCreatedAt, v))
}

// CreatedAtContainsFold applies the ContainsFold predicate on the "created_at" field.
func CreatedAtContainsFold(v string) predicate.Workout {
	return predicate.Workout(sql.FieldContainsFold(FieldCreatedAt, v))
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldImage, v))
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldImage, v))
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldImage, vs...))
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldImage, vs...))
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldImage, v))
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldImage, v))
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldImage, v))
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldImage, v))
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.Workout {
	return predicate.Workout(sql.FieldContains(FieldImage, v))
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.Workout {
	return predicate.Workout(sql.FieldHasPrefix(FieldImage, v))
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.Workout {
	return predicate.Workout(sql.FieldHasSuffix(FieldImage, v))
}

// ImageIsNil applies the IsNil predicate on the "image" field.
func ImageIsNil() predicate.Workout {
	return predicate.Workout(sql.FieldIsNull(FieldImage))
}

// ImageNotNil applies the NotNil predicate on the "image" field.
func ImageNotNil() predicate.Workout {
	return predicate.Workout(sql.FieldNotNull(FieldImage))
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEqualFold(FieldImage, v))
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.Workout {
	return predicate.Workout(sql.FieldContainsFold(FieldImage, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Workout {
	return predicate.Workout(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Workout {
	return predicate.Workout(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Workout {
	return predicate.Workout(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Workout {
	return predicate.Workout(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Workout {
	return predicate.Workout(sql.FieldContainsFold(FieldDescription, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v pksuid.ID) predicate.Workout {
	return predicate.Workout(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v pksuid.ID) predicate.Workout {
	vc := string(v)
	return predicate.Workout(sql.FieldContains(FieldUserID, vc))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v pksuid.ID) predicate.Workout {
	vc := string(v)
	return predicate.Workout(sql.FieldHasPrefix(FieldUserID, vc))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v pksuid.ID) predicate.Workout {
	vc := string(v)
	return predicate.Workout(sql.FieldHasSuffix(FieldUserID, vc))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Workout {
	return predicate.Workout(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Workout {
	return predicate.Workout(sql.FieldNotNull(FieldUserID))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v pksuid.ID) predicate.Workout {
	vc := string(v)
	return predicate.Workout(sql.FieldEqualFold(FieldUserID, vc))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v pksuid.ID) predicate.Workout {
	vc := string(v)
	return predicate.Workout(sql.FieldContainsFold(FieldUserID, vc))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkoutLogs applies the HasEdge predicate on the "workout_logs" edge.
func HasWorkoutLogs() predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WorkoutLogsTable, WorkoutLogsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkoutLogsWith applies the HasEdge predicate on the "workout_logs" edge with a given conditions (other predicates).
func HasWorkoutLogsWith(preds ...predicate.WorkoutLog) predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		step := newWorkoutLogsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Workout) predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Workout) predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
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
func Not(p predicate.Workout) predicate.Workout {
	return predicate.Workout(func(s *sql.Selector) {
		p(s.Not())
	})
}
