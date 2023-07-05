package main

import (
	"context"
	"net/http"
	"time"

	"github.com/sahidrahman404/gigachad-api/internal/request"
	"github.com/sahidrahman404/gigachad-api/internal/response"
	"github.com/sahidrahman404/gigachad-api/internal/types"
)

func (app *application) createWorkoutLogHandler(w http.ResponseWriter, r *http.Request) {
	params := types.CreateWorkoutLogParams{}

	err := request.DecodeJSON(w, r, &params)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	user := app.contextGetUser(r)

	workoutLog := types.NewWorkoutLogFromParams(params, user.Ent.ID)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.WorkoutLog.Insert(ctx, workoutLog)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusCreated, workoutLog)
	if err != nil {
		app.serverError(w, r, err)
	}
}
