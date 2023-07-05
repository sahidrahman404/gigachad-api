package main

import (
	"context"
	"net/http"
	"time"

	"github.com/sahidrahman404/gigachad-api/internal/request"
	"github.com/sahidrahman404/gigachad-api/internal/response"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

func (app *application) createWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	params := types.CreateWorkoutParams{}

	err := request.DecodeJSON(w, r, &params)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	user := app.contextGetUser(r)

	workout := types.NewWorkoutFromParams(params, user.Ent.ID)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.Workout.Insert(ctx, workout)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusCreated, workout)
	if err != nil {
		app.serverError(w, r, err)
	}
}
func (app *application) listWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	workouts, err := app.storage.Workout.GetAllForUser(ctx, user.Ent.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, workouts)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) getWorkoutHandler(w http.ResponseWriter, r *http.Request) {
	workoutID, err := app.readParams("workoutID", r)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	user := app.contextGetUser(r)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	workout, err := app.storage.Workout.GetForUser(ctx, workoutID, user.Ent.ID)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, workout)
	if err != nil {
		app.serverError(w, r, err)
	}
}
