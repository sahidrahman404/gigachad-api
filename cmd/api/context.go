package main

import (
	"context"
	"net/http"

	"github.com/sahidrahman404/gigachad-api/internal/gql"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

func (app *application) contextSetUser(r *http.Request, user *types.User) *http.Request {
	ctx := context.WithValue(r.Context(), gql.UserContextKey, user)
	return r.WithContext(ctx)
}

func (app *application) contextGetUser(r *http.Request) *types.User {
	user, ok := r.Context().Value(gql.UserContextKey).(*types.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
