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

func (app *application) createRoutineExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var params types.CreateRoutineExerciseParams

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

	user := app.contextGetUser(r)

	routineExercises := types.NewRoutineExerciseFromParams(params, user.Ent.ID)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.RoutineExercise.InsertMany(ctx, routineExercises, user.Ent.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(
		w,
		http.StatusCreated,
		routineExercises,
	)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getRoutineExerciseHandler(w http.ResponseWriter, r *http.Request) {
	routineID, err := app.readParams("routineID", r)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	routineExercises, err := app.storage.RoutineExercise.GetAllForRoutine(ctx, routineID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, routineExercises)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) updateRoutineExerciseHandler(w http.ResponseWriter, r *http.Request) {
	routineExerciseID, err := app.readParams("routineExerciseID", r)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	routineExerciseParams := types.UpdateRoutineExerciseParams{}
	err = request.DecodeJSON(w, r, &routineExerciseParams)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	routineExercise := types.UpdateRoutineExerciseFromParams(routineExerciseParams)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.RoutineExercise.Update(ctx, routineExercise, routineExerciseID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, routineExercise)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) deleteRoutineExerciseHandler(w http.ResponseWriter, r *http.Request) {
	routineExerciseID, err := app.readParams("routineExerciseID", r)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.RoutineExercise.Delete(ctx, routineExerciseID)
	if err != nil {
		app.serverError(w, r, err)
	}

	err = response.JSON(w, http.StatusOK, "Routine exercise has been successfully deleted")
	if err != nil {
		app.serverError(w, r, err)
	}
}
