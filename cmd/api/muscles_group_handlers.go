package main

import (
	"context"
	"net/http"
	"time"

	"github.com/sahidrahman404/gigachad-api/ent/schema/pksuid"
	"github.com/sahidrahman404/gigachad-api/internal/request"
	"github.com/sahidrahman404/gigachad-api/internal/response"
	"github.com/sahidrahman404/gigachad-api/internal/types"
	"github.com/sahidrahman404/gigachad-api/internal/validator"
)

func (app *application) createMusclesGroupHandler(w http.ResponseWriter, r *http.Request) {
	var params types.CreateMusclesGroupParams

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

	musclesGroup := types.NewMusclesGroupFromParams(params)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.MusclesGroup.Insert(ctx, musclesGroup)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(
		w,
		http.StatusCreated,
		musclesGroup,
	)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) listMusclesGroupHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	musclesGroup, err := app.storage.MusclesGroup.GetAll(ctx)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	err = response.JSON(w, http.StatusOK, musclesGroup)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) updateMusclesGroupHandler(w http.ResponseWriter, r *http.Request) {
	musclesGroupID, err := app.readParams("musclesgroupID", r)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	musclesGroupParams := types.UpdateMusclesGroupParams{}
	err = request.DecodeJSON(w, r, &musclesGroupParams)
	if err != nil {
		app.badRequest(w, r, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	musclesGroup, err := app.storage.MusclesGroup.Get(ctx, pksuid.ID(musclesGroupID))
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	types.UpdateMusclesGroupFromParams(&musclesGroupParams, musclesGroup)

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = app.storage.MusclesGroup.Update(ctx, musclesGroup)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = response.JSON(w, http.StatusOK, musclesGroup)
	if err != nil {
		app.serverError(w, r, err)
	}
}
