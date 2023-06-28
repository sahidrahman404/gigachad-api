package database

import "github.com/sahidrahman404/gigachad-api/ent"

type Storage struct {
	Users  UserStorer
	Tokens TokenStorer
}

func NewStorage(c *ent.Client) *Storage {
	return &Storage{
		Users:  NewEntUserStore(c),
		Tokens: NewEntTokenStore(c),
	}
}
