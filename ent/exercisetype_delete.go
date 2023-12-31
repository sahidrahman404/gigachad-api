// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/exercisetype"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
)

// ExerciseTypeDelete is the builder for deleting a ExerciseType entity.
type ExerciseTypeDelete struct {
	config
	hooks    []Hook
	mutation *ExerciseTypeMutation
}

// Where appends a list predicates to the ExerciseTypeDelete builder.
func (etd *ExerciseTypeDelete) Where(ps ...predicate.ExerciseType) *ExerciseTypeDelete {
	etd.mutation.Where(ps...)
	return etd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (etd *ExerciseTypeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, etd.sqlExec, etd.mutation, etd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (etd *ExerciseTypeDelete) ExecX(ctx context.Context) int {
	n, err := etd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (etd *ExerciseTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(exercisetype.Table, sqlgraph.NewFieldSpec(exercisetype.FieldID, field.TypeString))
	if ps := etd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, etd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	etd.mutation.done = true
	return affected, err
}

// ExerciseTypeDeleteOne is the builder for deleting a single ExerciseType entity.
type ExerciseTypeDeleteOne struct {
	etd *ExerciseTypeDelete
}

// Where appends a list predicates to the ExerciseTypeDelete builder.
func (etdo *ExerciseTypeDeleteOne) Where(ps ...predicate.ExerciseType) *ExerciseTypeDeleteOne {
	etdo.etd.mutation.Where(ps...)
	return etdo
}

// Exec executes the deletion query.
func (etdo *ExerciseTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := etdo.etd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{exercisetype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (etdo *ExerciseTypeDeleteOne) ExecX(ctx context.Context) {
	if err := etdo.Exec(ctx); err != nil {
		panic(err)
	}
}
