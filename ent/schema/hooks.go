package schema

import (
	"context"
	"errors"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/hook"
)

var (
	ErrDeletingExerciseImage = errors.New("failed deleting exercise image from bucket")
	ErrExerciseNotFound      = errors.New("exercise record not found")
	ErrExercise              = errors.New("cannot get exercise record")
	ErrMissingExerciseIDField = errors.New("exercise ID field is missing")
)

func DeleteExerciseImage() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ExerciseFunc(func(ctx context.Context, em *ent.ExerciseMutation) (ent.Value, error) {
			id, exists := em.ID()
			if !exists {
				return nil, ErrMissingExerciseIDField
			}
			ex, err := em.Client().Exercise.Get(ctx, id)
			if err != nil {
				return nil, ErrExerciseNotFound
			}
			if ex.Image == nil {
				return next.Mutate(ctx, em)
			}

			em.DeleteObjectInput.Key = &ex.Image.Filename
			_, err = em.S3Client.DeleteObject(ctx, em.DeleteObjectInput)
			if err != nil {
				return nil, ErrDeletingExerciseImage
			}
			return next.Mutate(ctx, em)
		})
	}
	return hook.On(hk, ent.OpDeleteOne)
}

func DeleteExerciseImageOnUpsert() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ExerciseFunc(func(ctx context.Context, em *ent.ExerciseMutation) (ent.Value, error) {
			_, exists := em.Image()
			if !exists {
				return next.Mutate(ctx, em)
			}
			id, exists := em.ID()
			if !exists {
				return nil, ErrMissingExerciseIDField
			}
			ex, err := em.Client().Exercise.Get(ctx, id)
			if err != nil {
				return next.Mutate(ctx, em)
			}
			if ex.Image == nil {
				return next.Mutate(ctx, em)
			}

			em.DeleteObjectInput.Key = &ex.Image.Filename
			_, err = em.S3Client.DeleteObject(ctx, em.DeleteObjectInput)
			if err != nil {
				return nil, ErrDeletingExerciseImage
			}
			return next.Mutate(ctx, em)
		})
	}
	return hook.On(hk, ent.OpCreate)
}
