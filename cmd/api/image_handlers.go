package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/sahidrahman404/gigachad-api/internal/request"
	"github.com/sahidrahman404/gigachad-api/internal/response"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

func (app *application) getSignedUrl(w http.ResponseWriter, r *http.Request) {
	var params types.CreateOriginalURLParams

	err := request.DecodeJSONStrict(w, r, &params)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	v := validator.NewValidator()

	if params.ValidateOriginalURL(v); v.HasErrors() {
		app.failedValidation(w, r, v)
		return
	}

	signedURL := signURL(
		app.config.imgproxy.key,
		app.config.imgproxy.salt,
		app.config.imgproxy.imgproxyHost,
		*params.Width,
		params.Height,
		params.Src,
	)

	err = response.JSON(w, http.StatusOK, map[string]string{
		"signedURL": signedURL,
	})
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getTransformedUrls(w http.ResponseWriter, r *http.Request) {
	var params types.CreateTransformedURLParams

	err := request.DecodeJSONStrict(w, r, &params)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	var result []*string
	for _, breakpoint := range params.BreakPoint {
		if params.AspectRatio != nil && params.Height != nil {
			devided := float64(breakpoint) / *params.AspectRatio
			transformedHeight := math.Round(devided)
			transformed := signURL(
				app.config.imgproxy.key,
				app.config.imgproxy.salt,
				app.config.imgproxy.imgproxyHost,
				breakpoint,
				&transformedHeight,
				params.Src,
			)
			stitch := fmt.Sprintf("%s %dw", transformed, params.BreakPoint)
			result = append(result, &stitch)
		}

		if params.Height == nil {
			transformed := signURL(
				app.config.imgproxy.key,
				app.config.imgproxy.salt,
				app.config.imgproxy.imgproxyHost,
				breakpoint,
				nil,
				params.Src,
			)
			stitch := fmt.Sprintf("%s %dw", transformed, params.BreakPoint)
			result = append(result, &stitch)
		}
	}

	err = response.JSON(w, http.StatusOK, result)
	if err != nil {
		app.serverError(w, r, err)
	}
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
