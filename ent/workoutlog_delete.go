// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sahidrahman404/gigachad-api/ent/predicate"
	"github.com/sahidrahman404/gigachad-api/ent/workoutlog"
)

// WorkoutLogDelete is the builder for deleting a WorkoutLog entity.
type WorkoutLogDelete struct {
	config
	hooks    []Hook
	mutation *WorkoutLogMutation
}

// Where appends a list predicates to the WorkoutLogDelete builder.
func (wld *WorkoutLogDelete) Where(ps ...predicate.WorkoutLog) *WorkoutLogDelete {
	wld.mutation.Where(ps...)
	return wld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wld *WorkoutLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wld.sqlExec, wld.mutation, wld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wld *WorkoutLogDelete) ExecX(ctx context.Context) int {
	n, err := wld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wld *WorkoutLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(workoutlog.Table, sqlgraph.NewFieldSpec(workoutlog.FieldID, field.TypeString))
	if ps := wld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wld.mutation.done = true
	return affected, err
}

// WorkoutLogDeleteOne is the builder for deleting a single WorkoutLog entity.
type WorkoutLogDeleteOne struct {
	wld *WorkoutLogDelete
}

// Where appends a list predicates to the WorkoutLogDelete builder.
func (wldo *WorkoutLogDeleteOne) Where(ps ...predicate.WorkoutLog) *WorkoutLogDeleteOne {
	wldo.wld.mutation.Where(ps...)
	return wldo
}

// Exec executes the deletion query.
func (wldo *WorkoutLogDeleteOne) Exec(ctx context.Context) error {
	n, err := wldo.wld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{workoutlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wldo *WorkoutLogDeleteOne) ExecX(ctx context.Context) {
	if err := wldo.Exec(ctx); err != nil {
		panic(err)
	}
}
