package types

import (
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

type CreateTransformedURLParams struct {
	Width       *int     `json:"width"`
	Height      *float64 `json:"height"`
	Src         string   `json:"src"`
	BreakPoint  []int    `json:"breakPoint"`
	AspectRatio *float64 `json:"aspectRatio"`
}

type CreateOriginalURLParams struct {
	Width  *int     `json:"width"`
	Height *float64 `json:"height"`
	Src    string   `json:"src"`
}

func (o *CreateOriginalURLParams) ValidateOriginalURL(v *validator.Validator) {
	v.CheckField(o.Src != "", "src", "must be provided")
}
