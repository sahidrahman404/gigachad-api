package schema

import (
	"context"
	"errors"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/hook"
)

var (
	ErrMissingImageField     = errors.New("image field is missing")
	ErrDeletingExerciseImage = errors.New("failed deleting exercise image from bucket")
	ErrExerciseNotFound      = errors.New("exercise record not found")
	ErrExercise              = errors.New("cannot get exercise record")
)

func DeleteExerciseImage() ent.Hook {
	hk := func(next ent.Mutator) ent.Mutator {
		return hook.ExerciseFunc(func(ctx context.Context, em *ent.ExerciseMutation) (ent.Value, error) {
			id, exists := em.ID()
			if !exists {
				return nil, ErrMissingImageField
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
	return hook.On(hk, ent.OpDeleteOne|ent.OpCreate)
}
