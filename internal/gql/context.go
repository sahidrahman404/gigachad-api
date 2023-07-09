package gql

import (
	"context"

	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type contextKey string

const userContextKey = contextKey("user")

func (r *Resolver) forContext(ctx context.Context) *types.User {
	user, _ := ctx.Value(userContextKey).(*types.User)
	return user
}
