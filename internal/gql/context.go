package gql

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/internal/types"
)

func (r *Resolver) getUserFromCtx(ctx context.Context) *types.User {
	user, ok := ctx.Value(types.UserContextKey).(*types.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
