package types

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

type Token struct {
	Ent       *ent.Token
	Plaintext string
}

func GenerateToken(userID pksuid.ID, ttl time.Duration, scope string) (*Token, error) {
	token := &Token{
		Ent: &ent.Token{
			UserID: userID,
			Expiry: time.Now().Add(ttl).Format(time.RFC3339),
			Scope:  scope,
		},
	}

	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Ent.Hash = hash[:]

	return token, nil
}

func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {
	v.CheckField(tokenPlaintext != "", "token", "must be provided")
	v.CheckField(len(tokenPlaintext) != 26, "token", "must be 26 bytes long")
}
