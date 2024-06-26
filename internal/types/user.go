package types

import (
	"errors"
	"strings"

	gigachadv1 "buf.build/gen/go/sahidrahman/gigachadapis/protocolbuffers/go/gigachad/v1"
	"github.com/sahidrahman404/gigachad-api/ent"
	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/ent/user"
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
	Unit     *user.Unit `json:"userPreference"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Name     string     `json:"name"`
}

type User struct {
	Ent *ent.User
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := params.HashPassword()
	if err != nil {
		return nil, err
	}

	return &User{
		Ent: &ent.User{
			Username:       strings.ToLower(params.Username),
			Email:          strings.ToLower(params.Email),
			HashedPassword: *encpw,
			Name:           params.NameCaser(),
			Unit:           params.GetUserPreference(),
		},
	}, nil
}

func (p CreateUserParams) HashPassword() (*string, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	result := string(encpw)
	return &result, nil
}

func (p CreateUserParams) NameCaser() string {
	caser := cases.Title(language.AmericanEnglish)
	return caser.String(p.Name)
}

func (p CreateUserParams) GetUserPreference() user.Unit {
	unit := user.DefaultUnit

	if p.Unit != nil {
		unit = *p.Unit
	}

	return unit
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
	return names[len(names)-1]
}

func (u *User) GetUnitEnum() gigachadv1.Unit {
	result := gigachadv1.Unit_UNIT_METRIC_UNSPECIFIED
	if u.Ent.Unit == user.UnitIMPERIAL {
		result = gigachadv1.Unit_UNIT_IMPERIAL
	}
	return result
}
