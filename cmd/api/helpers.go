package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) newEmailData() map[string]any {
	data := map[string]any{
		"BaseURL": app.config.baseURL,
	}

	return data
}

func (app *application) backgroundTask(fn func() error) {
	app.wg.Add(1)

	go func() {
		defer app.wg.Done()

		defer func() {
			err := recover()
			if err != nil {
				app.reportError(fmt.Errorf("%s", err))
			}
		}()

		err := fn()
		if err != nil {
			app.reportError(err)
		}
	}()
}

func (app *application) readParams(param string, r *http.Request) (string, error) {
	params := chi.URLParam(r, param)

	if params == "" {
		return "", fmt.Errorf("invalid %s parameter", param)
	}

	return params, nil
}
