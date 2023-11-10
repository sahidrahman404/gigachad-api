package types

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/sahidrahman404/gigachad-api/internal/validator"
	"golang.org/x/exp/slices"
)

// Common screen widths. These will be filtered
// according to the image size and layout
var DEFAULT_RESOLUTIONS = []int{
	6016, // 6K
	5120, // 5K
	4480, // 4.5K
	3840, // 4K
	3200, // QHD+
	2560, // WQXGA
	2048, // QXGA
	1920, // 1080p
	1668, // Various iPads
	1280, // 720p
	1080, // iPhone 6-8 Plus
	960,  // older horizontal phones
	828,  // iPhone XR/11
	750,  // iPhone 6-8
	640,  // older and lower-end phones
}

// Width of the blurred, low-res image
const LOW_RES_WIDTH = 24

type CreateGetSrcSetParams struct {
	Width       *int     `json:"width"`
	Height      *float64 `json:"height"`
	AspectRatio *float64 `json:"aspectRatio"`
	Src         string   `json:"src"`
	BreakPoint  []int    `json:"breakPoint"`
	layout      Layout
}

type CreateTransformedURLParams struct {
	Width       *int     `json:"width"`
	Height      *float64 `json:"height"`
	AspectRatio *float64 `json:"aspectRatio"`
	Src         string   `json:"src"`
	BreakPoint  []int    `json:"breakPoint"`
}

type CreateOriginalURLParams struct {
	Width  *int     `json:"width"`
	Height *float64 `json:"height"`
	Src    string   `json:"src"`
}

func (o *CreateOriginalURLParams) ValidateOriginalURL(v *validator.Validator) {
	v.CheckField(o.Src != "", "src", "must be provided")
}

type Layout int

const (
	fixed Layout = iota
	constrained
	fullWidth
)

func (l Layout) String() string {
	return []string{"fixed", "constrained", "fullwidth"}[l]
}

func getBreakpoints(width *int, l Layout) []int {
	if l.String() == "fullWidth" {
		return DEFAULT_RESOLUTIONS
	}

	if width == nil {
		return []int{}
	}

	doubleWidth := *width * 2

	if l.String() == "fixed" {
		return []int{*width, doubleWidth}
	}

	if l.String() == "constrained" {
		filteredDefaultResolutions := []int{
			// Always include the image at 1x and 2x the specified width
			*width,
			doubleWidth,
		}
		for _, w := range DEFAULT_RESOLUTIONS {
			// Filter out any resolutions that are larger than the double-res image
			if w < doubleWidth {
				filteredDefaultResolutions = append(filteredDefaultResolutions, w)
			}
		}
		return filteredDefaultResolutions
	}

	return []int{}
}

func GetSrcSet(
	p CreateGetSrcSetParams,
	key, salt, imgproxyHost string,
) string {
	if p.BreakPoint == nil {
		breakpoints := getBreakpoints(p.Width, p.layout)
		slices.Sort(breakpoints)
		p.BreakPoint = breakpoints
		return fetchUrls(p, breakpoints, key, salt, imgproxyHost)
	}
	return fetchUrls(p, p.BreakPoint, key, salt, imgproxyHost)
}

func getTransformedUrls(
	params CreateTransformedURLParams,
	key, salt, imgproxyHost string,
) []string {
	var result []string
	for _, breakpoint := range params.BreakPoint {
		if params.AspectRatio != nil && params.Height != nil {
			devided := float64(breakpoint) / *params.AspectRatio
			transformedHeight := math.Round(devided)
			transformed := signURL(
				key,
				salt,
				imgproxyHost,
				breakpoint,
				&transformedHeight,
				params.Src,
			)
			stitch := fmt.Sprintf("%s %dw", transformed, breakpoint)
			result = append(result, stitch)
		}

		if params.Height == nil {
			transformed := signURL(
				key,
				salt,
				imgproxyHost,
				breakpoint,
				nil,
				params.Src,
			)
			stitch := fmt.Sprintf("%s %dw", transformed, breakpoint)
			result = append(result, stitch)
		}
	}

	return result
}

func signURL(
	key, salt, imgproxyURL string,
	width int,
	height *float64,
	src string,
) string {
	var keyBin, saltBin []byte
	var err error

	if keyBin, err = hex.DecodeString(key); err != nil {
		log.Fatal(err)
	}

	if saltBin, err = hex.DecodeString(salt); err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf(
		"/resize:auto:%d:%d/plain/%s@webp",
		width,
		height,
		src,
	)

	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("%s/%s%s", imgproxyURL, signature, path)
}

func fetchUrls(
	p CreateGetSrcSetParams,
	breakpoint []int,
	key, salt, imgproxyHost string,
) string {
	response := getTransformedUrls(
		CreateTransformedURLParams{
			Width:       p.Width,
			Height:      p.Height,
			Src:         p.Src,
			BreakPoint:  breakpoint,
			AspectRatio: p.AspectRatio,
		},
		key,
		salt,
		imgproxyHost,
	)
	return strings.Join(response, ",\n")
}
