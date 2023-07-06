package gql

import "fmt"

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
