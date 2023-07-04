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

func (app *application) createExerciseTypeHandler(w http.ResponseWriter, r *http.Request) {
	var params types.CreateExerciseTypeParams

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

	exerciseType := types.NewExerciseTypeFromParmas(params)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.ExerciseType.Insert(ctx, exerciseType)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(
		w,
		http.StatusCreated,
		exerciseType,
	)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) listExerciseTypeHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	exerciseType, err := app.storage.ExerciseType.GetAll(ctx)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, exerciseType)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) updateExerciseTypeHandler(w http.ResponseWriter, r *http.Request) {
	exerciseTypeID, err := app.readParams("exercisetypeID", r)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	exerciseTypeParams := types.UpdateExerciseTypeParams{}
	err = request.DecodeJSON(w, r, &exerciseTypeParams)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	exerciseType, err := app.storage.ExerciseType.Get(ctx, exerciseTypeID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	types.UpdateExerciseTypeFromParams(&exerciseTypeParams, exerciseType)

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.ExerciseType.Update(ctx, exerciseType)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, exerciseType)
	if err != nil {
		app.serverError(w, r, err)
	}
}
