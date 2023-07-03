// Code generated by ent, DO NOT EDIT.

package exercise

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContainsFold(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldName, v))
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldImage, v))
}

// HowTo applies equality check predicate on the "how_to" field. It's identical to HowToEQ.
func HowTo(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldHowTo, v))
}

// EquipmentID applies equality check predicate on the "equipment_id" field. It's identical to EquipmentIDEQ.
func EquipmentID(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldEquipmentID, v))
}

// MusclesGroupID applies equality check predicate on the "muscles_group_id" field. It's identical to MusclesGroupIDEQ.
func MusclesGroupID(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldMusclesGroupID, v))
}

// ExerciseTypeID applies equality check predicate on the "exercise_type_id" field. It's identical to ExerciseTypeIDEQ.
func ExerciseTypeID(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldExerciseTypeID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldUserID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContainsFold(FieldName, v))
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldImage, v))
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNEQ(FieldImage, v))
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldIn(FieldImage, vs...))
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNotIn(FieldImage, vs...))
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGT(FieldImage, v))
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGTE(FieldImage, v))
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLT(FieldImage, v))
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLTE(FieldImage, v))
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContains(FieldImage, v))
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasPrefix(FieldImage, v))
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasSuffix(FieldImage, v))
}

// ImageIsNil applies the IsNil predicate on the "image" field.
func ImageIsNil() predicate.Exercise {
	return predicate.Exercise(sql.FieldIsNull(FieldImage))
}

// ImageNotNil applies the NotNil predicate on the "image" field.
func ImageNotNil() predicate.Exercise {
	return predicate.Exercise(sql.FieldNotNull(FieldImage))
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEqualFold(FieldImage, v))
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContainsFold(FieldImage, v))
}

// HowToEQ applies the EQ predicate on the "how_to" field.
func HowToEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldHowTo, v))
}

// HowToNEQ applies the NEQ predicate on the "how_to" field.
func HowToNEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNEQ(FieldHowTo, v))
}

// HowToIn applies the In predicate on the "how_to" field.
func HowToIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldIn(FieldHowTo, vs...))
}

// HowToNotIn applies the NotIn predicate on the "how_to" field.
func HowToNotIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNotIn(FieldHowTo, vs...))
}

// HowToGT applies the GT predicate on the "how_to" field.
func HowToGT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGT(FieldHowTo, v))
}

// HowToGTE applies the GTE predicate on the "how_to" field.
func HowToGTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGTE(FieldHowTo, v))
}

// HowToLT applies the LT predicate on the "how_to" field.
func HowToLT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLT(FieldHowTo, v))
}

// HowToLTE applies the LTE predicate on the "how_to" field.
func HowToLTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLTE(FieldHowTo, v))
}

// HowToContains applies the Contains predicate on the "how_to" field.
func HowToContains(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContains(FieldHowTo, v))
}

// HowToHasPrefix applies the HasPrefix predicate on the "how_to" field.
func HowToHasPrefix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasPrefix(FieldHowTo, v))
}

// HowToHasSuffix applies the HasSuffix predicate on the "how_to" field.
func HowToHasSuffix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasSuffix(FieldHowTo, v))
}

// HowToIsNil applies the IsNil predicate on the "how_to" field.
func HowToIsNil() predicate.Exercise {
	return predicate.Exercise(sql.FieldIsNull(FieldHowTo))
}

// HowToNotNil applies the NotNil predicate on the "how_to" field.
func HowToNotNil() predicate.Exercise {
	return predicate.Exercise(sql.FieldNotNull(FieldHowTo))
}

// HowToEqualFold applies the EqualFold predicate on the "how_to" field.
func HowToEqualFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEqualFold(FieldHowTo, v))
}

// HowToContainsFold applies the ContainsFold predicate on the "how_to" field.
func HowToContainsFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContainsFold(FieldHowTo, v))
}

// EquipmentIDEQ applies the EQ predicate on the "equipment_id" field.
func EquipmentIDEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldEquipmentID, v))
}

// EquipmentIDNEQ applies the NEQ predicate on the "equipment_id" field.
func EquipmentIDNEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNEQ(FieldEquipmentID, v))
}

// EquipmentIDIn applies the In predicate on the "equipment_id" field.
func EquipmentIDIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldIn(FieldEquipmentID, vs...))
}

// EquipmentIDNotIn applies the NotIn predicate on the "equipment_id" field.
func EquipmentIDNotIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNotIn(FieldEquipmentID, vs...))
}

// EquipmentIDGT applies the GT predicate on the "equipment_id" field.
func EquipmentIDGT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGT(FieldEquipmentID, v))
}

// EquipmentIDGTE applies the GTE predicate on the "equipment_id" field.
func EquipmentIDGTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGTE(FieldEquipmentID, v))
}

// EquipmentIDLT applies the LT predicate on the "equipment_id" field.
func EquipmentIDLT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLT(FieldEquipmentID, v))
}

// EquipmentIDLTE applies the LTE predicate on the "equipment_id" field.
func EquipmentIDLTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLTE(FieldEquipmentID, v))
}

// EquipmentIDContains applies the Contains predicate on the "equipment_id" field.
func EquipmentIDContains(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContains(FieldEquipmentID, v))
}

// EquipmentIDHasPrefix applies the HasPrefix predicate on the "equipment_id" field.
func EquipmentIDHasPrefix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasPrefix(FieldEquipmentID, v))
}

// EquipmentIDHasSuffix applies the HasSuffix predicate on the "equipment_id" field.
func EquipmentIDHasSuffix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasSuffix(FieldEquipmentID, v))
}

// EquipmentIDIsNil applies the IsNil predicate on the "equipment_id" field.
func EquipmentIDIsNil() predicate.Exercise {
	return predicate.Exercise(sql.FieldIsNull(FieldEquipmentID))
}

// EquipmentIDNotNil applies the NotNil predicate on the "equipment_id" field.
func EquipmentIDNotNil() predicate.Exercise {
	return predicate.Exercise(sql.FieldNotNull(FieldEquipmentID))
}

// EquipmentIDEqualFold applies the EqualFold predicate on the "equipment_id" field.
func EquipmentIDEqualFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEqualFold(FieldEquipmentID, v))
}

// EquipmentIDContainsFold applies the ContainsFold predicate on the "equipment_id" field.
func EquipmentIDContainsFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContainsFold(FieldEquipmentID, v))
}

// MusclesGroupIDEQ applies the EQ predicate on the "muscles_group_id" field.
func MusclesGroupIDEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldMusclesGroupID, v))
}

// MusclesGroupIDNEQ applies the NEQ predicate on the "muscles_group_id" field.
func MusclesGroupIDNEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNEQ(FieldMusclesGroupID, v))
}

// MusclesGroupIDIn applies the In predicate on the "muscles_group_id" field.
func MusclesGroupIDIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldIn(FieldMusclesGroupID, vs...))
}

// MusclesGroupIDNotIn applies the NotIn predicate on the "muscles_group_id" field.
func MusclesGroupIDNotIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNotIn(FieldMusclesGroupID, vs...))
}

// MusclesGroupIDGT applies the GT predicate on the "muscles_group_id" field.
func MusclesGroupIDGT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGT(FieldMusclesGroupID, v))
}

// MusclesGroupIDGTE applies the GTE predicate on the "muscles_group_id" field.
func MusclesGroupIDGTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGTE(FieldMusclesGroupID, v))
}

// MusclesGroupIDLT applies the LT predicate on the "muscles_group_id" field.
func MusclesGroupIDLT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLT(FieldMusclesGroupID, v))
}

// MusclesGroupIDLTE applies the LTE predicate on the "muscles_group_id" field.
func MusclesGroupIDLTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLTE(FieldMusclesGroupID, v))
}

// MusclesGroupIDContains applies the Contains predicate on the "muscles_group_id" field.
func MusclesGroupIDContains(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContains(FieldMusclesGroupID, v))
}

// MusclesGroupIDHasPrefix applies the HasPrefix predicate on the "muscles_group_id" field.
func MusclesGroupIDHasPrefix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasPrefix(FieldMusclesGroupID, v))
}

// MusclesGroupIDHasSuffix applies the HasSuffix predicate on the "muscles_group_id" field.
func MusclesGroupIDHasSuffix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasSuffix(FieldMusclesGroupID, v))
}

// MusclesGroupIDIsNil applies the IsNil predicate on the "muscles_group_id" field.
func MusclesGroupIDIsNil() predicate.Exercise {
	return predicate.Exercise(sql.FieldIsNull(FieldMusclesGroupID))
}

// MusclesGroupIDNotNil applies the NotNil predicate on the "muscles_group_id" field.
func MusclesGroupIDNotNil() predicate.Exercise {
	return predicate.Exercise(sql.FieldNotNull(FieldMusclesGroupID))
}

// MusclesGroupIDEqualFold applies the EqualFold predicate on the "muscles_group_id" field.
func MusclesGroupIDEqualFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEqualFold(FieldMusclesGroupID, v))
}

// MusclesGroupIDContainsFold applies the ContainsFold predicate on the "muscles_group_id" field.
func MusclesGroupIDContainsFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContainsFold(FieldMusclesGroupID, v))
}

// ExerciseTypeIDEQ applies the EQ predicate on the "exercise_type_id" field.
func ExerciseTypeIDEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldExerciseTypeID, v))
}

// ExerciseTypeIDNEQ applies the NEQ predicate on the "exercise_type_id" field.
func ExerciseTypeIDNEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNEQ(FieldExerciseTypeID, v))
}

// ExerciseTypeIDIn applies the In predicate on the "exercise_type_id" field.
func ExerciseTypeIDIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldIn(FieldExerciseTypeID, vs...))
}

// ExerciseTypeIDNotIn applies the NotIn predicate on the "exercise_type_id" field.
func ExerciseTypeIDNotIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNotIn(FieldExerciseTypeID, vs...))
}

// ExerciseTypeIDGT applies the GT predicate on the "exercise_type_id" field.
func ExerciseTypeIDGT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGT(FieldExerciseTypeID, v))
}

// ExerciseTypeIDGTE applies the GTE predicate on the "exercise_type_id" field.
func ExerciseTypeIDGTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGTE(FieldExerciseTypeID, v))
}

// ExerciseTypeIDLT applies the LT predicate on the "exercise_type_id" field.
func ExerciseTypeIDLT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLT(FieldExerciseTypeID, v))
}

// ExerciseTypeIDLTE applies the LTE predicate on the "exercise_type_id" field.
func ExerciseTypeIDLTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLTE(FieldExerciseTypeID, v))
}

// ExerciseTypeIDContains applies the Contains predicate on the "exercise_type_id" field.
func ExerciseTypeIDContains(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContains(FieldExerciseTypeID, v))
}

// ExerciseTypeIDHasPrefix applies the HasPrefix predicate on the "exercise_type_id" field.
func ExerciseTypeIDHasPrefix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasPrefix(FieldExerciseTypeID, v))
}

// ExerciseTypeIDHasSuffix applies the HasSuffix predicate on the "exercise_type_id" field.
func ExerciseTypeIDHasSuffix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasSuffix(FieldExerciseTypeID, v))
}

// ExerciseTypeIDIsNil applies the IsNil predicate on the "exercise_type_id" field.
func ExerciseTypeIDIsNil() predicate.Exercise {
	return predicate.Exercise(sql.FieldIsNull(FieldExerciseTypeID))
}

// ExerciseTypeIDNotNil applies the NotNil predicate on the "exercise_type_id" field.
func ExerciseTypeIDNotNil() predicate.Exercise {
	return predicate.Exercise(sql.FieldNotNull(FieldExerciseTypeID))
}

// ExerciseTypeIDEqualFold applies the EqualFold predicate on the "exercise_type_id" field.
func ExerciseTypeIDEqualFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEqualFold(FieldExerciseTypeID, v))
}

// ExerciseTypeIDContainsFold applies the ContainsFold predicate on the "exercise_type_id" field.
func ExerciseTypeIDContainsFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContainsFold(FieldExerciseTypeID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Exercise {
	return predicate.Exercise(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Exercise {
	return predicate.Exercise(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Exercise {
	return predicate.Exercise(sql.FieldNotNull(FieldUserID))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Exercise {
	return predicate.Exercise(sql.FieldContainsFold(FieldUserID, v))
}

// HasRoutineExercises applies the HasEdge predicate on the "routine_exercises" edge.
func HasRoutineExercises() predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RoutineExercisesTable, RoutineExercisesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoutineExercisesWith applies the HasEdge predicate on the "routine_exercises" edge with a given conditions (other predicates).
func HasRoutineExercisesWith(preds ...predicate.RoutineExercise) predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		step := newRoutineExercisesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWorkoutLogs applies the HasEdge predicate on the "workout_logs" edge.
func HasWorkoutLogs() predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WorkoutLogsTable, WorkoutLogsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWorkoutLogsWith applies the HasEdge predicate on the "workout_logs" edge with a given conditions (other predicates).
func HasWorkoutLogsWith(preds ...predicate.WorkoutLog) predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		step := newWorkoutLogsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEquipments applies the HasEdge predicate on the "equipments" edge.
func HasEquipments() predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EquipmentsTable, EquipmentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEquipmentsWith applies the HasEdge predicate on the "equipments" edge with a given conditions (other predicates).
func HasEquipmentsWith(preds ...predicate.Equipment) predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		step := newEquipmentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMusclesGroups applies the HasEdge predicate on the "muscles_groups" edge.
func HasMusclesGroups() predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MusclesGroupsTable, MusclesGroupsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMusclesGroupsWith applies the HasEdge predicate on the "muscles_groups" edge with a given conditions (other predicates).
func HasMusclesGroupsWith(preds ...predicate.MusclesGroup) predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		step := newMusclesGroupsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasExerciseTypes applies the HasEdge predicate on the "exercise_types" edge.
func HasExerciseTypes() predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ExerciseTypesTable, ExerciseTypesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasExerciseTypesWith applies the HasEdge predicate on the "exercise_types" edge with a given conditions (other predicates).
func HasExerciseTypesWith(preds ...predicate.ExerciseType) predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		step := newExerciseTypesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Exercise) predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Exercise) predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
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
func Not(p predicate.Exercise) predicate.Exercise {
	return predicate.Exercise(func(s *sql.Selector) {
		p(s.Not())
	})
}