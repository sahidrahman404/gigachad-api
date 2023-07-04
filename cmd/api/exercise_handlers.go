package main

import (
	"context"
	"net/http"
	"time"

	"github.com/sahidrahman404/gigachad-api/internal/request"
	"github.com/sahidrahman404/gigachad-api/internal/response"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

func (app *application) createExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var params types.CreateExerciseParams

	err := request.DecodeJSON(w, r, &params)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	v := validator.NewValidator()

	if params.Validate(v); v.HasErrors() {
		app.failedValidation(w, r, v)
		return
	}

	exercise := types.NewExerciseFromParams(params, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.Exercise.Insert(ctx, exercise)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(
		w,
		http.StatusCreated,
		exercise,
	)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) listExerciseHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	exercises, err := app.storage.Exercise.GetAll(ctx)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, exercises)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) updateExerciseHandler(w http.ResponseWriter, r *http.Request) {
	exerciseID, err := app.readParams("exerciseID", r)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	exerciseParams := types.UpdateExerciseParams{}
	err = request.DecodeJSON(w, r, &exerciseParams)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	exercise, err := app.storage.Exercise.Get(ctx, exerciseID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	types.UpdateExerciseFromParams(&exerciseParams, exercise)

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.Exercise.Update(ctx, exercise)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, exercise)
	if err != nil {
		app.serverError(w, r, err)
	}
}
