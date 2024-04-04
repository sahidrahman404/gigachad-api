// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"
	"time"

	"github.com/sahidrahman404/gigachad-api/ent/equipment"
	"github.com/sahidrahman404/gigachad-api/ent/exercise"
	"github.com/sahidrahman404/gigachad-api/ent/exercisetype"
	"github.com/sahidrahman404/gigachad-api/ent/musclesgroup"
	"github.com/sahidrahman404/gigachad-api/ent/routine"
	"github.com/sahidrahman404/gigachad-api/ent/routineexercise"
	"github.com/sahidrahman404/gigachad-api/ent/schema"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/token"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/ent/workout"
	"github.com/sahidrahman404/gigachad-api/ent/workoutlog"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	equipmentMixin := schema.Equipment{}.Mixin()
	equipmentMixinFields0 := equipmentMixin[0].Fields()
	_ = equipmentMixinFields0
	equipmentFields := schema.Equipment{}.Fields()
	_ = equipmentFields
	// equipmentDescID is the schema descriptor for id field.
	equipmentDescID := equipmentMixinFields0[0].Descriptor()
	// equipment.DefaultID holds the default value on creation for the id field.
	equipment.DefaultID = equipmentDescID.Default.(func() pksuid.ID)
	exerciseMixin := schema.Exercise{}.Mixin()
	exercise.Policy = privacy.NewPolicies(schema.Exercise{})
	exercise.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := exercise.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	exerciseHooks := schema.Exercise{}.Hooks()

	exercise.Hooks[1] = exerciseHooks[0]
	exerciseMixinFields0 := exerciseMixin[0].Fields()
	_ = exerciseMixinFields0
	exerciseFields := schema.Exercise{}.Fields()
	_ = exerciseFields
	// exerciseDescID is the schema descriptor for id field.
	exerciseDescID := exerciseMixinFields0[0].Descriptor()
	// exercise.DefaultID holds the default value on creation for the id field.
	exercise.DefaultID = exerciseDescID.Default.(func() pksuid.ID)
	exercisetypeMixin := schema.ExerciseType{}.Mixin()
	exercisetypeMixinFields0 := exercisetypeMixin[0].Fields()
	_ = exercisetypeMixinFields0
	exercisetypeFields := schema.ExerciseType{}.Fields()
	_ = exercisetypeFields
	// exercisetypeDescID is the schema descriptor for id field.
	exercisetypeDescID := exercisetypeMixinFields0[0].Descriptor()
	// exercisetype.DefaultID holds the default value on creation for the id field.
	exercisetype.DefaultID = exercisetypeDescID.Default.(func() pksuid.ID)
	musclesgroupMixin := schema.MusclesGroup{}.Mixin()
	musclesgroupMixinFields0 := musclesgroupMixin[0].Fields()
	_ = musclesgroupMixinFields0
	musclesgroupFields := schema.MusclesGroup{}.Fields()
	_ = musclesgroupFields
	// musclesgroupDescID is the schema descriptor for id field.
	musclesgroupDescID := musclesgroupMixinFields0[0].Descriptor()
	// musclesgroup.DefaultID holds the default value on creation for the id field.
	musclesgroup.DefaultID = musclesgroupDescID.Default.(func() pksuid.ID)
	routineMixin := schema.Routine{}.Mixin()
	routine.Policy = privacy.NewPolicies(schema.Routine{})
	routine.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := routine.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	routineMixinFields0 := routineMixin[0].Fields()
	_ = routineMixinFields0
	routineFields := schema.Routine{}.Fields()
	_ = routineFields
	// routineDescID is the schema descriptor for id field.
	routineDescID := routineMixinFields0[0].Descriptor()
	// routine.DefaultID holds the default value on creation for the id field.
	routine.DefaultID = routineDescID.Default.(func() pksuid.ID)
	routineexerciseMixin := schema.RoutineExercise{}.Mixin()
	routineexerciseMixinFields0 := routineexerciseMixin[0].Fields()
	_ = routineexerciseMixinFields0
	routineexerciseFields := schema.RoutineExercise{}.Fields()
	_ = routineexerciseFields
	// routineexerciseDescID is the schema descriptor for id field.
	routineexerciseDescID := routineexerciseMixinFields0[0].Descriptor()
	// routineexercise.DefaultID holds the default value on creation for the id field.
	routineexercise.DefaultID = routineexerciseDescID.Default.(func() pksuid.ID)
	tokenMixin := schema.Token{}.Mixin()
	tokenMixinFields0 := tokenMixin[0].Fields()
	_ = tokenMixinFields0
	tokenFields := schema.Token{}.Fields()
	_ = tokenFields
	// tokenDescID is the schema descriptor for id field.
	tokenDescID := tokenMixinFields0[0].Descriptor()
	// token.DefaultID holds the default value on creation for the id field.
	token.DefaultID = tokenDescID.Default.(func() pksuid.ID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(time.Time)
	// userDescActivated is the schema descriptor for activated field.
	userDescActivated := userFields[6].Descriptor()
	// user.DefaultActivated holds the default value on creation for the activated field.
	user.DefaultActivated = userDescActivated.Default.(int)
	// userDescVersion is the schema descriptor for version field.
	userDescVersion := userFields[7].Descriptor()
	// user.DefaultVersion holds the default value on creation for the version field.
	user.DefaultVersion = userDescVersion.Default.(int)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() pksuid.ID)
	workoutMixin := schema.Workout{}.Mixin()
	workoutMixinFields0 := workoutMixin[0].Fields()
	_ = workoutMixinFields0
	workoutFields := schema.Workout{}.Fields()
	_ = workoutFields
	// workoutDescCreatedAt is the schema descriptor for created_at field.
	workoutDescCreatedAt := workoutFields[3].Descriptor()
	// workout.DefaultCreatedAt holds the default value on creation for the created_at field.
	workout.DefaultCreatedAt = workoutDescCreatedAt.Default.(time.Time)
	// workoutDescID is the schema descriptor for id field.
	workoutDescID := workoutMixinFields0[0].Descriptor()
	// workout.DefaultID holds the default value on creation for the id field.
	workout.DefaultID = workoutDescID.Default.(func() pksuid.ID)
	workoutlogMixin := schema.WorkoutLog{}.Mixin()
	workoutlogMixinFields0 := workoutlogMixin[0].Fields()
	_ = workoutlogMixinFields0
	workoutlogFields := schema.WorkoutLog{}.Fields()
	_ = workoutlogFields
	// workoutlogDescCreatedAt is the schema descriptor for created_at field.
	workoutlogDescCreatedAt := workoutlogFields[1].Descriptor()
	// workoutlog.DefaultCreatedAt holds the default value on creation for the created_at field.
	workoutlog.DefaultCreatedAt = workoutlogDescCreatedAt.Default.(time.Time)
	// workoutlogDescID is the schema descriptor for id field.
	workoutlogDescID := workoutlogMixinFields0[0].Descriptor()
	// workoutlog.DefaultID holds the default value on creation for the id field.
	workoutlog.DefaultID = workoutlogDescID.Default.(func() pksuid.ID)
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)
