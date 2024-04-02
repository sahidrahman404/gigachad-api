package types

import (
	"errors"
	"strings"

	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type contextKey string

const UserContextKey = contextKey("user")

const (
	bcryptCost = 12
)

var AnonymousUser = &User{}

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type User struct {
	Ent *ent.User
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		Ent: &ent.User{
			Username:       strings.ToLower(params.Username),
			Email:          strings.ToLower(params.Email),
			HashedPassword: string(encpw),
			Name:           params.Name,
		},
	}, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.CheckField(email != "", "email", "email must be provided")
	v.CheckField(
		validator.Matches(email, validator.RgxEmail),
		"email",
		"email must be a valid email address",
	)
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.CheckField(password != "", "password", "password must be provided")
	v.CheckField(len(password) >= 8, "password", "password must be at least 8 bytes long")
	v.CheckField(len(password) <= 72, "password", "password must not be more 72 bytes long")
}

func (p CreateUserParams) ValidateUser(v *validator.Validator) {
	v.CheckField(p.Username != "", "username", "username must be provided")
	v.CheckField(len(p.Username) <= 72, "username", "username must not be more 72 bytes long")

	v.CheckField(p.Name != "", "name", "name must be provided")
	v.CheckField(len(p.Name) <= 500, "name", "name must not be more than 500 bytes long")

	ValidatePasswordPlaintext(v, p.Password)
	ValidateEmail(v, p.Email)
}

func (u *User) IsAnonymous() bool {
	return u == AnonymousUser
}

func (u *User) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Ent.HashedPassword), []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, nil
		}
	}

	return true, nil
}

func (u *User) SetPassword(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}
	u.Ent.HashedPassword = string(hash)
	return nil
}

func (u *User) GetUserID() *pksuid.ID {
	user := u.Ent
	if user == nil {
		return nil
	}
	return &user.ID
}

func (u *User) GetUserLastName() string {
	name := u.Ent.Name
	names := strings.Split(name, " ")
	caser := cases.Title(language.AmericanEnglish)
	return caser.String(names[len(names)-1])
}
