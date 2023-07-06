package gql

import "runtime/debug"

func (r *Resolver) reportError(err error) {
	trace := debug.Stack()

	r.logger.Error(err, trace)
}
