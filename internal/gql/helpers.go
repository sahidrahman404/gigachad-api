package gql

import (
	"context"
	"fmt"

	"github.com/sahidrahman404/gigachad-api/ent"
)

func (r *Resolver) backgroundTask(fn func() error) {
	r.wg.Add(1)

	go func() {
		defer r.wg.Done()

		defer func() {
			err := recover()
			if err != nil {
				r.reportError(fmt.Errorf("%s", err))
			}
		}()

		err := fn()
		if err != nil {
			r.reportError(err)
		}
	}()
}

func (r *Resolver) WithTx(ctx context.Context, fn func(tx *ent.Tx) error) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
