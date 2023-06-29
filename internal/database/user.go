package database

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"time"

	"entgo.io/ent/privacy"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/token"
	"github.com/sahidrahman404/gigachad-api/ent/user"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

var (
	ErrDuplicateEmail       = errors.New("duplicate email")
	ErrDuplicateUsername    = errors.New("duplicate username")
	SqliteDuplicateEmail    = `SQLite error: UNIQUE constraint failed: users.email`
	SqliteDuplicateUsername = `SQLite error: UNIQUE constraint failed: users.username`
	EntUserNotFound         = `ent: user not found`
)

type UserStorer interface {
	Insert(*types.User) error
	GetByUsername(string) (*types.User, error)
	Update(*types.User, context.Context) error
	GetForToken(string, string) (*types.User, error)
	GetByEmail(string) (*types.User, error)
}

type EntUserStore struct {
	Client *ent.Client
}

func NewEntUserStore(client *ent.Client) *EntUserStore {
	return &EntUserStore{Client: client}
}

func (e *EntUserStore) Insert(u *types.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	allow := privacy.DecisionContext(ctx, privacy.Allow)
	user, err := e.Client.User.Create().
		SetUsername(u.Ent.Username).
		SetEmail(u.Ent.Email).
		SetHashedPassword(u.Ent.HashedPassword).
		SetName(u.Ent.Name).
		Save(allow)

	if err != nil {
		switch {
		case parseSqliteError(err.Error()) == SqliteDuplicateEmail:
			return ErrDuplicateEmail
		case parseSqliteError(err.Error()) == SqliteDuplicateUsername:
			return ErrDuplicateUsername
		default:
			return err
		}
	}

	u.Ent.ID = user.ID
	u.Ent.Activated = user.Activated
	u.Ent.CreatedAt = user.CreatedAt
	u.Ent.Version = user.Version

	return nil
}

func (e EntUserStore) GetByUsername(username string) (*types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	user, err := e.Client.User.Query().Where(user.Username(username)).Only(ctx)
	if err != nil {
		switch {
		case err.Error() == EntUserNotFound:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &types.User{
		Ent: user,
	}, nil
}

func (e EntUserStore) Update(u *types.User, c context.Context) error {
	user, err := e.Client.User.UpdateOneID(u.Ent.ID).
		SetUsername(u.Ent.Username).
		SetEmail(u.Ent.Email).
		SetHashedPassword(u.Ent.HashedPassword).
		SetActivated(u.Ent.Activated).
		SetName(u.Ent.Name).
		SetVersion(u.Ent.Version + 1).
		Save(c)

	if err != nil {
		switch {
		case parseSqliteError(err.Error()) == SqliteDuplicateEmail:
			return ErrDuplicateEmail
		case parseSqliteError(err.Error()) == SqliteDuplicateUsername:
			return ErrDuplicateUsername
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	u.Ent.Version = user.Version

	return nil
}

func (e EntUserStore) GetForToken(tokenScope, tokenPlaintext string) (*types.User, error) {
	tokenHash := sha256.Sum256([]byte(tokenPlaintext))

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user, err := e.Client.User.Query().
		Where(
			user.HasTokensWith(token.Hash(tokenHash[:]),
				token.Scope(tokenScope),
				token.ExpiryGT(time.Now().Format(time.RFC3339)),
			),
		).
		Only(ctx)

	if err != nil {
		switch {
		case err.Error() == EntUserNotFound:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &types.User{
		Ent: user,
	}, nil
}

func (e EntUserStore) GetByEmail(email string) (*types.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	user, err := e.Client.User.Query().Where(user.Email(email)).Only(ctx)

	if err != nil {
		switch {
		case err.Error() == EntUserNotFound:
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &types.User{
		Ent: user,
	}, nil
}
