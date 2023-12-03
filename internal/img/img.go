package img

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"slices"
	"strings"

	"github.com/sahidrahman404/gigachad-api/ent/schema/schematype"
	"github.com/sahidrahman404/gigachad-api/internal/aws"
)

// Common screen widths. These will be filtered
// according to the image size and layout
var DEFAULT_RESOLUTIONS = []int{
	// 6016, // 6K
	// 5120, // 5K
	// 4480, // 4.5K
	// 3840, // 4K
	// 3200, // QHD+
	// 2560, // WQXGA
	// 2048, // QXGA
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

type Imgproxy struct {
	Key          string
	Salt         string
	ImgproxyHost string
	OriginCDN    string
}

func SetImageField(p schematype.Image, awsCfg aws.AWSConfig, i *Imgproxy) schematype.Image {
	if p.AspectRatio != nil {
		if p.Width != nil {
			if p.Height == nil {
				height := float64(*p.Width) / *p.AspectRatio
				p.Height = &height
			}
		} else if p.Height != nil {
			width := int(*p.Height) * int(*p.AspectRatio)
			p.Width = &width
		}
	} else if p.Width != nil && p.Height != nil {
		aspectRatio := float64(*p.Width) / *p.Height
		p.AspectRatio = &aspectRatio
	}
	src := fmt.Sprintf("https://%s/%s", i.OriginCDN, p.Filename)
	srcset := getSrcSet(p, awsCfg, i)
	breakpoints := getBreakpoints(p.Width, p.Layout)
	sizes := getSizes(p)
	p.SrcSet = srcset
	p.BreakPoints = breakpoints
	p.Sizes = *sizes
	p.Src = src
	setStyle(&p)
	setPriority(&p)
	setAltRole(&p)

	if p.Layout == "fullWidth" || p.Layout == "constrained" {
		p.Width = nil
		p.Height = nil
	}

	return p
}

func getBreakpoints(width *int, layout string) []int {
	if layout == "fullWidth" {
		return DEFAULT_RESOLUTIONS
	}

	if width == nil {
		return []int{}
	}

	doubleWidth := *width * 2

	if layout == "fixed" {
		return []int{*width, doubleWidth}
	}

	if layout == "constrained" {
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

func getSrcSet(
	p schematype.Image,
	awsCfg aws.AWSConfig,
	i *Imgproxy,
) string {
	if p.BreakPoints == nil {
		breakpoints := getBreakpoints(p.Width, p.Layout)
		slices.Sort(breakpoints)
		return buildSrcSet(p, breakpoints, awsCfg, i.Key, i.Salt, i.ImgproxyHost)
	}
	return buildSrcSet(p, p.BreakPoints, awsCfg, i.Key, i.Salt, i.ImgproxyHost)
}

func getTransformedUrls(
	p schematype.Image,
	awsCfg aws.AWSConfig,
	key, salt, imgproxyHost string,
) []string {
	var result []string
	for _, breakpoint := range p.BreakPoints {
		if p.AspectRatio != nil && p.Height != nil {
			divided := float64(breakpoint) / *p.AspectRatio
			transformedHeight := math.Round(divided)
			transformed := signURL(
				key,
				salt,
				imgproxyHost,
				breakpoint,
				transformedHeight,
				awsCfg,
				p.Filename,
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
	height float64,
	awsCfg aws.AWSConfig,
	filename string,
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
		"/resize:fit:%d:%v/plain/s3://%s/%s@webp",
		width,
		height,
		awsCfg.AWSBucket,
		filename,
	)

	mac := hmac.New(sha256.New, keyBin)
	mac.Write(saltBin)
	mac.Write([]byte(path))
	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return fmt.Sprintf("%s/%s%s", imgproxyURL, signature, path)
}

func buildSrcSet(
	p schematype.Image,
	breakpoint []int,
	awsCfg aws.AWSConfig,
	key, salt, imgproxyHost string,
) string {
	p.BreakPoints = breakpoint
	response := getTransformedUrls(p, awsCfg, key, salt, imgproxyHost)
	return strings.Join(response, ",\n")
}

func setStyle(p *schematype.Image) {
	if p.Layout == "fixed" && p.Width != nil && p.Height != nil {
		width := fmt.Sprintf("%dpx", *p.Width)
		height := fmt.Sprintf("%fpx", *p.Height)
		p.Style.Width = &width
		p.Style.Height = &height
	}
	if p.Layout == "constrained" && p.Width != nil && p.Height != nil {
		width := fmt.Sprintf("%dpx", *p.Width)
		height := fmt.Sprintf("%fpx", *p.Height)
		p.MaxWidth = &width
		p.MaxHeight = &height
		width = "100%"
		p.Style.Width = &width
	}
	if p.Layout == "fullWidth" && p.Height != nil && p.AspectRatio != nil {
		width := "100%"
		p.Style.Width = &width
		aspectRatio := fmt.Sprintf("%f", *p.AspectRatio)
		p.Style.AspectRatio = &aspectRatio
		height := fmt.Sprintf("%fpx", *p.Height)
		p.Style.Height = &height
	}
}

func getSizes(p schematype.Image) *string {
	switch p.Layout {
	// If screen is wider than the max size, image width is the max size,
	// otherwise it's the width of the screen
	case `constrained`:
		result := fmt.Sprintf("(min-width: %dpx) %dpx, 100vw", *p.Width, *p.Width)
		return &result

	// Image is always the same width, whatever the size of the screen
	case `fixed`:
		result := fmt.Sprintf("%dpx", *p.Width)
		return &result

	// Image is always the width of the screen
	case `fullWidth`:
		result := `100vw`
		return &result

	default:
		return nil
	}
}

func setPriority(p *schematype.Image) {
	if p.Priority {
		loading := "eager"
		fetchpriority := "high"
		p.Loading = &loading
		p.FetchPriority = &fetchpriority
	} else {
		// Otherwise use native lazy loading and async decoding
		loading := "lazy"
		fetchpriority := "async"
		p.Loading = &loading
		p.Decoding = &fetchpriority
	}
}

func setAltRole(p *schematype.Image) {
	if p.Alt == nil {
		role := "presentation"
		p.Role = &role
	}
}

func SetNillableImageField(
	p *schematype.Image,
	awsCfg aws.AWSConfig,
	i *Imgproxy,
) *schematype.Image {
	if p == nil {
		return nil
	}
	result := SetImageField(*p, awsCfg, i)
	return &result
}
