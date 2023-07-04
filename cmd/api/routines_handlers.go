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

func (app *application) createRoutineHandler(w http.ResponseWriter, r *http.Request) {
	var params types.CreateRoutineParams

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

	routine := types.NewRoutineFromParams(params, user.Ent.ID)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.Routine.Insert(ctx, routine)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(
		w,
		http.StatusCreated,
		routine,
	)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) listRoutineHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	user := app.contextGetUser(r)
	routine, err := app.storage.Routine.GetAllForUser(ctx, user.Ent.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, routine)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) updateRoutineHandler(w http.ResponseWriter, r *http.Request) {
	routineID, err := app.readParams("routineID", r)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	routineParams := types.CreateRoutineParams{}
	err = request.DecodeJSON(w, r, &routineParams)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	v := validator.NewValidator()

	if routineParams.Validate(v); v.HasErrors() {
		app.failedValidation(w, r, v)
	}

	user := app.contextGetUser(r)

	routine := types.NewRoutineFromParams(routineParams, user.Ent.ID)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.Routine.Update(ctx, routine, routineID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, routine)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) deleteRoutineHandler(w http.ResponseWriter, r *http.Request) {
	routineID, err := app.readParams("routineID", r)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.Routine.Delete(ctx, routineID)
	if err != nil {
		app.serverError(w, r, err)
	}

	err = response.JSON(w, http.StatusOK, "Routine has been successfully deleted")
	if err != nil {
		app.serverError(w, r, err)
	}
}
