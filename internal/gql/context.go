package gql

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type contextKey string

const UserContextKey = contextKey("user")

func (r *Resolver) forContext(ctx context.Context) *types.User {
	user, ok := ctx.Value(UserContextKey).(*types.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
