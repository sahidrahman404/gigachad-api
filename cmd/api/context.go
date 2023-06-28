package main

import (
	"context"
	"net/http"

	"github.com/sahidrahman404/gigachad-api/internal/types"
)

type contextKey string

const userContextKey = contextKey("user")

func (app *application) contextSetUser(r *http.Request, user *types.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func (app *application) contextGetUser(r *http.Request) *types.User {
	user, ok := r.Context().Value(userContextKey).(*types.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
