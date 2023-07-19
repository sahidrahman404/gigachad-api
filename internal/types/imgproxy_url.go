package types

import (
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

type CreateOriginalURLParams struct {
	Width  string `json:"width"`
	Height string `json:"height"`
	Src    string `json:"src"`
}

func (o *CreateOriginalURLParams) validateOriginalURL(v *validator.Validator) {
	v.CheckField(o.Width != "", "width", "must be provided")
	v.CheckField(o.Height != "", "height", "must be provided")
	v.CheckField(o.Src != "", "src", "must be provided")
}
