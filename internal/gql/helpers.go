package gql

import (
	"context"
	"fmt"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/internal/types"
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

func (r *Resolver) requireActivatedUser(u *types.User) (*ent.User, error) {
	if u.IsAnonymous() {
		return nil, r.authenticationRequired()
	}

	if u.Ent.Activated == 0 {
		return nil, r.inactiveAccount()
	}

	return u.Ent, nil
}
