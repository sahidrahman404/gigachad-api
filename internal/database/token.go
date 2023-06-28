package database

import (
	"context"
	"time"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/token"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

const (
	ScopeActivation     = "activation"
	SocpeAuthentication = "authentication"
	ScopePasswordReset  = "password-reset"
)

type TokenStorer interface {
	New(userID string, ttl time.Duration, scope string) (*types.Token, error)
	Insert(token *types.Token) error
	DeleteAllForUser(scope string, userID string) error
}

type EntTokenStore struct {
	Client *ent.Client
}

func NewEntTokenStore(c *ent.Client) *EntTokenStore {
	return &EntTokenStore{Client: c}
}

func (e *EntTokenStore) New(userID string, ttl time.Duration, scope string) (*types.Token, error) {
	token, err := types.GenerateToken(userID, ttl, scope)
	if err != nil {
		return nil, err
	}

	err = e.Insert(token)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (e *EntTokenStore) Insert(t *types.Token) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	err := e.Client.Token.Create().
		SetExpiry(t.Ent.Expiry).
		SetHash(t.Ent.Hash).
		SetScope(t.Ent.Scope).
		SetUserID(t.Ent.UserID).
		Exec(ctx)

	return err
}

func (e *EntTokenStore) DeleteAllForUser(scope string, userID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	_, err := e.Client.Token.Delete().
		Where(token.Scope(scope), token.HasUsersWith(user.ID(userID))).
		Exec(ctx)

	return err
}
